package edge

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/wapi/wnd"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// WebView 是创建在一个用 wapi 创建的原生窗口里的, 然后原生窗口是被嵌入到炫彩窗口或元素里的.
//
// 对 WebView 的一些视觉相关的更改操作必须在炫彩 UI 线程里执行.
type WebView struct {
	WebView2

	cbDestroy   func(wv *WebView) // 原生窗口销毁时调用
	cbNCDestroy func(wv *WebView) // 原生窗口非客户区销毁时调用

	// 绑定go函数: name : func
	bindings map[string]interface{}
	// 绑定id: name : id
	bindingsid map[string]string
	// 读写锁bindings
	rwxBindings sync.RWMutex

	hWindow int // 炫彩窗口句柄
	hParent int // 原生窗口宿主炫彩窗口或元素句柄

	// WebView 的固定位置与大小
	size WebViewSize

	// 圆角半径
	roundRadius int32
	// 在窗口获得焦点时尝试保持 WebView 的焦点。
	autoFocus bool
	// 填充父
	fillParent bool
	// WebView 在炫彩窗口而不是元素里
	isInWindow bool
}

// NewWebView 创建 WebView 到炫彩窗口或元素.
//
// hParent: 炫彩窗口或元素句柄.
//
// opts: WebView 选项, 使用 edge.WithXXX 系列函数.
//   - 可查看 WebViewOptions.
//   - 默认会启用的 WebView 选项可查看 DefaultEnabledWebViewOptions.
func (e *Edge) NewWebView(hParent int, opts ...WebViewOption) (*WebView, error) {
	// 获取默认 WebView 配置
	options := defaultWebViewOptions()
	// 应用所有选项
	for _, opt := range opts {
		opt(options)
	}

	return newWebView(e, hParent, options)
}

// NewWebViewWithWindow 创建 WebView 到炫彩窗口.
//   - 内部会使用 XML 创建炫彩窗口.
//   - 可在 opts 参数中使用 edge.WithXmlWindowXXX 系列函数设置窗口属性. 可查看 XmlWindowOptions.
//
// opts: WebView 选项, 使用 edge.WithXXX 系列函数.
//   - 可查看 WebViewOptions.
//   - 默认会启用的 WebView 选项可查看 DefaultEnabledWebViewOptions.
func (e *Edge) NewWebViewWithWindow(opts ...WebViewOption) (*window.Window, *WebView, error) {
	// 获取默认 WebView 配置
	options := defaultWebViewOptions()
	// 应用所有选项
	for _, opt := range opts {
		opt(options)
	}

	xmlStr := options.XmlWindowOpts.XmlStr
	if xmlStr == "" {
		xmlStr = xcc.XmlTransparentWindow
		if options.XmlWindowOpts.ShadowAngleSize > 0 {
			xmlStr = strings.Replace(xmlStr, `content=""`, `content="" shadowRightAngle="false" shadowAngleSize="`+xc.Itoa(options.XmlWindowOpts.ShadowAngleSize)+`"`, 1)
		}
		xmlStr = strings.Replace(xmlStr, `className=""`, `className="`+options.XmlWindowOpts.ClassName+`"`, 1)
		xmlStr = strings.Replace(xmlStr, `content=""`, `content="`+options.XmlWindowOpts.Title+`"`, 1)
		xmlStr = strings.Replace(xmlStr, `rect="20,20,500,500"`, fmt.Sprintf(`rect="20,20,%d,%d"`, options.XmlWindowOpts.Width, options.XmlWindowOpts.Height), 1)
		xmlStr = strings.Replace(xmlStr, `shadowColor="#80000000"`, `shadowColor="#`+xc.HexRGBA(xc.ParseRGBA(options.XmlWindowOpts.ShadowColor))+`"`, 1)
		xmlStr = strings.Replace(xmlStr, `shadowDepth="128"`, `shadowDepth="`+xc.Itoa(options.XmlWindowOpts.ShadowDepth)+`"`, 1)
		xmlStr = strings.Replace(xmlStr, `shadowSize="8"`, `shadowSize="`+xc.Itoa(options.XmlWindowOpts.ShadowSize)+`"`, 1)
		xmlStr = strings.Replace(xmlStr, `transparentAlpha="255"`, `transparentAlpha="`+xc.Itoa(int32(options.XmlWindowOpts.TransparentAlpha))+`"`, 1)
	}

	// 创建窗口从xml
	w := window.NewByLayoutStringW(xmlStr, options.XmlWindowOpts.HParent, 0)
	if w == nil {
		return nil, nil, errors.New("窗口创建失败")
	}

	// 创建 WebView
	wv, err := newWebView(e, w.Handle, options)
	if err != nil {
		w.CloseWindow()
		return nil, nil, err
	}

	w.AdjustLayout()
	return w, wv, nil
}

// 创建 WebView 到炫彩窗口或元素.
//
// hParent: 炫彩窗口或元素句柄.
//
// opts: *WebViewOptions.
func newWebView(e *Edge, hParent int, opts *WebViewOptions) (*WebView, error) {
	if !xc.XC_IsHELE(hParent) && !xc.XC_IsHWINDOW(hParent) {
		return nil, errors.New("hParent is not a xcgui window or element")
	}

	w := &WebView{}
	w.bindings = map[string]interface{}{}
	w.bindingsid = map[string]string{}
	w.autoFocus = opts.AutoFocus
	w.size = opts.WebViewSize
	w.fillParent = opts.FillParent
	w.Edge = e

	err := createWebViewWithOptionsByXcgui(w, hParent, opts)
	if err != nil {
		return nil, err
	}

	if w.autoFocus {
		ReportErrorAuto(w.Focus())
	}

	return w, nil
}

// 创建 WebView 宿主窗口.
func createWebViewWithOptionsByXcgui(w *WebView, hParent int, opt *WebViewOptions) error {
	w.hParent = hParent
	// 获取父窗口或元素的HWND, 其实是一个, 就是父窗口的HWND, 炫彩元素没有自己的HWND
	var hWndXC uintptr
	if w.hParent > 0 {
		if xc.XC_IsHWINDOW(w.hParent) {
			w.hWindow = w.hParent
			w.isInWindow = true
			hWndXC = xc.XWnd_GetHWND(w.hParent)
		} else if xc.XC_IsHELE(w.hParent) {
			hWndXC = xc.XWidget_GetHWND(w.hParent)
			w.hWindow = xc.XWidget_GetHWINDOW(w.hParent)
		}
	}

	dpi := int32(96)
	if w.hWindow > 0 {
		dpi = xc.XWnd_GetDPI(w.hWindow)
		// 如果 WebView 启用了自动焦点, 炫彩就关闭自动焦点
		xc.XWnd_EnableAutoFocus(w.hWindow, !w.autoFocus)
	}

	hInstance := wapi.GetModuleHandleEx(0, "")

	var icon uintptr
	if opt.IconId == 0 {
		// load default icon
		icow := wapi.GetSystemMetrics(wapi.SM_CXICON)
		icoh := wapi.GetSystemMetrics(wapi.SM_CYICON)
		icon = wapi.LoadImageW(hInstance, uintptr(32512), wapi.IMAGE_ICON, icow, icoh, wapi.LR_DEFAULTCOLOR)
	} else {
		// load icon from resource
		icon = wapi.LoadImageW(hInstance, uintptr(opt.IconId), wapi.IMAGE_ICON, 0, 0, wapi.LR_DEFAULTSIZE|wapi.LR_SHARED)
	}

	// 注册窗口类名
	if opt.ClassName == "" {
		opt.ClassName = "go_xcgui_WebView"
	}
	wc := wapi.WNDCLASSEX{
		Style:         wapi.CS_HREDRAW | wapi.CS_VREDRAW | wapi.CS_PARENTDC,
		CbSize:        uint32(unsafe.Sizeof(wapi.WNDCLASSEX{})),
		HInstance:     hInstance,
		LpszClassName: common.StrPtr(opt.ClassName),
		HIcon:         icon,
		HIconSm:       icon,
		LpfnWndProc:   syscall.NewCallback(wndproc),
	}
	wapi.RegisterClassEx(&wc)

	// 计算圆角半径
	w.roundRadius = xc.DpiConvRound(dpi, opt.RoundRadius)
	// 创建宿主窗口
	w.hwnd = wapi.CreateWindowEx(0, opt.ClassName, opt.Title, xcc.WS_MINIMIZE, xc.DpiConvRound(dpi, w.size.Left), xc.DpiConvRound(dpi, w.size.Top), xc.DpiConvRound(dpi, w.size.Width), xc.DpiConvRound(dpi, w.size.Height), 0, 0, hInstance, 0)

	// 记录窗口上下文
	hwndContext.SetWindowContext(w.hwnd, w)
	xcContext.SetWindowContext(uintptr(w.hParent), w)

	// 显示窗口, 更新窗口, 设置焦点
	wapi.ShowWindow(w.hwnd, xcc.SW_SHOWMINIMIZED)
	wapi.UpdateWindow(w.hwnd)

	// ------------------------ 创建 WebView2 控制器 ------------------------
	initWebViewEventImpl(w.GetWebViewEventImpl())
	w.WebView2.msgcb_xcgui = w.msgcb_xcgui

	err := newWebView2Controller(w, opt)
	if err != nil {
		wapi.DestroyWindow(w.hwnd) // 销毁原生窗口
		return err
	}
	// ------------------------ 创建 WebView2 控制器 END ------------------------

	// 设置 WebView2 宿主窗口为炫彩窗口或元素的子窗口
	wapi.SetParent(w.hwnd, hWndXC)
	// 设置 WebView2 宿主窗口样式
	wapi.SetWindowLongPtrW(w.hwnd, wapi.GWL_STYLE, int(xcc.WS_CHILD|xcc.WS_VISIBLE))

	// 窗口 调整位置和大小
	xc.XWnd_RemoveEventC(w.hWindow, xcc.XWM_WINDPROC, onWndProc)
	xc.XWnd_RegEventC1(w.hWindow, xcc.XWM_WINDPROC, onWndProc)

	// 元素事件
	if !w.isInWindow {
		// 调整位置和大小
		xc.XEle_RemoveEventC(w.hParent, xcc.XE_SIZE, onEleSize)
		xc.XEle_RegEventC1(w.hParent, xcc.XE_SIZE, onEleSize)

		// 跟随父销毁
		xc.XEle_RemoveEventC(w.hParent, xcc.XE_DESTROY, onEleDestroy)
		xc.XEle_RegEventC1(w.hParent, xcc.XE_DESTROY, onEleDestroy)

		// 跟随父显示或隐藏
		xc.XEle_RemoveEventC(w.hParent, xcc.XE_SHOW, onEleShow)
		xc.XEle_RegEventC1(w.hParent, xcc.XE_SHOW, onEleShow)
	}

	w.UpdateSize()
	return nil
}

// UpdateSize 更新 WebView 宿主原生窗口的大小和尺寸
func (w *WebView) UpdateSize() {
	if !wapi.IsWindow(w.hwnd) {
		return
	}
	var rcClient xc.RECT // 炫彩窗口或元素的客户区矩形
	if w.isInWindow {
		xc.XWnd_GetBodyRect(w.hParent, &rcClient)
	} else {
		xc.XEle_GetWndClientRect(w.hParent, &rcClient)
	}

	var left, top, width, height int32
	// 填充父
	if w.fillParent {
		left = rcClient.Left
		top = rcClient.Top
		height = rcClient.Bottom - rcClient.Top
		if w.isInWindow {
			var rcWindow xc.RECT
			xc.XWnd_GetRect(w.hParent, &rcWindow)
			// 原本应该只用客户区坐标来计算, 但是发现存在一个问题, 就是客户区坐标的 Right 会莫名奇妙的比窗口创建好时多出 1px, 明明窗口 Rect 没有任何改变, 这导致 WebView 的右边超出了炫彩窗口 2px.
			// 目前采用窗口宽度减去客户区左边的两倍来计算客户区宽度也就是 WebView 的宽度, 这等于是假设客户区是左右对称的, 实际上客户区左边和右边可能会有不同的宽度(虽然可能没有这种需求), 但是目前没有更好的办法.
			width = rcWindow.Right - rcWindow.Left - rcClient.Left*2
		} else {
			width = rcClient.Right - rcClient.Left
		}
	} else {
		left = rcClient.Left + w.size.Left
		top = rcClient.Top + w.size.Top
		width = w.size.Width
		height = w.size.Height
	}

	dpi := int32(96)
	if xc.XC_IsHWINDOW(w.hWindow) {
		dpi = xc.XWnd_GetDPI(w.hWindow)
	}
	wapi.SetWindowPos(w.hwnd, 0, xc.DpiConvRound(dpi, left), xc.DpiConvRound(dpi, top), xc.DpiConvRound(dpi, width), xc.DpiConvRound(dpi, height), wapi.SWP_NOACTIVATE|wapi.SWP_NOZORDER)
}

func onEleDestroy(hEle int, pbHandled *bool) int {
	handle := uintptr(hEle)
	if w := xcContext.GetWindowContext(handle); w != nil {
		if wapi.IsWindow(w.hwnd) {
			wapi.DestroyWindow(w.hwnd)
		}
		xcContext.DeleteWindowContext(handle)
	}
	return 0
}

func onEleSize(hEle int, nFlags xcc.AdjustLayout_, nAdjustNo uint32, pbHandled *bool) int {
	if w := xcContext.GetWindowContext(uintptr(hEle)); w != nil {
		w.UpdateSize()
	}
	return 0
}

func onEleShow(hEle int, bShow bool, pbHandled *bool) int {
	if w := xcContext.GetWindowContext(uintptr(hEle)); w != nil {
		if !wapi.IsWindow(w.hwnd) {
			return 0
		}
		nCmdShow := xcc.SW_SHOW
		if !bShow {
			nCmdShow = xcc.SW_HIDE
		}
		wapi.ShowWindow(w.hwnd, nCmdShow)
	}
	return 0
}

func onWndProc(hWindow int, message uint32, wParam, lParam uintptr, pbHandled *bool) int {
	switch message {
	case wapi.WM_SIZE:
		if w := xcContext.GetWindowContext(uintptr(hWindow)); w != nil { // 原生窗口宿主是炫彩窗口
			if wParam != wapi.SIZE_MINIMIZED { // 最小化时不执行 UpdateSize
				w.UpdateSize()
			}
		} else { // 更新每个元素中的 WebView 位置
			hEles := xcContext.GetHEles(hWindow)
			for i := 0; i < len(hEles); i++ {
				if w := xcContext.GetWindowContext(uintptr(hEles[i])); w != nil {
					if wParam != wapi.SIZE_MINIMIZED {
						w.UpdateSize()
					}
				}
			}
		}
	case wapi.WM_DESTROY:
		handle := uintptr(hWindow)
		if w := xcContext.GetWindowContext(handle); w != nil { // 原生窗口宿主是炫彩窗口
			// 当窗口有透明通道时, 关闭的时候会漏出炫彩窗口, 导致闪烁一下, 所以先隐藏炫彩窗口
			xc.XWnd_Show(hWindow, false)
			xcContext.DeleteWindowContext(handle)
		}
	case wapi.WM_ACTIVATE:
		handle := uintptr(hWindow)
		if w := xcContext.GetWindowContext(handle); w != nil && w.autoFocus { // 原生窗口宿主是炫彩窗口
			wapi.PostMessageW(w.hwnd, wapi.WM_ACTIVATE, wParam, lParam)
		} else { // 只发给最后一个元素中的 WebView
			hEles := xcContext.GetHEles(hWindow)
			if len(hEles) > 0 {
				if w := xcContext.GetWindowContext(uintptr(hEles[len(hEles)-1])); w != nil && w.autoFocus && w.Controller != nil {
					if wapi.GetParent(w.hwnd) == xc.XWnd_GetHWND(hWindow) {
						wapi.PostMessageW(w.hwnd, wapi.WM_ACTIVATE, wParam, lParam)
					}
				}
			}
		}
	}
	return 0
}

// SetSize 设置 WebView 中设置的固定位置与大小.
//   - 需 EnableFillParent(false) 才会生效.
func (w *WebView) SetSize(size WebViewSize) {
	w.size = size
}

// GetSize 获取 WebView 中设置的固定位置与大小.
func (w *WebView) GetSize() WebViewSize {
	return w.size
}

// SetRoundRadius 设置原生窗口的圆角半径。
//   - 需注意炫彩窗口的圆角需通过 SetShadowInfo 设置.
//   - WebView 是创建在一个用 wapi 创建的原生窗口里的, 然后原生窗口是被嵌入到炫彩窗口或元素里的.
func (w *WebView) SetRoundRadius(radius int32) error {
	if radius < 1 {
		w.roundRadius = 0
		return wnd.SetWindowRound(w.hwnd, w.roundRadius)
	}
	dpi := int32(96)
	if xc.XC_IsHWINDOW(w.hWindow) {
		dpi = xc.XWnd_GetDPI(w.hWindow)
	}
	w.roundRadius = xc.DpiConvRound(dpi, radius)
	return wnd.SetWindowRound(w.hwnd, w.roundRadius)
}

// GetRoundRadius 获取原生窗口的圆角半径。
func (w *WebView) GetRoundRadius() int32 {
	return w.roundRadius
}

// EnableAutoFocus 启用在窗口获得焦点时尝试保持 WebView 的焦点.
//   - 内部会调用 xc.XWnd_EnableAutoFocus(w.hWindow, !bEnable)禁用/启用炫彩窗口的自动获取焦点
//   - 如果不禁用炫彩窗口的自动获取焦点, 那焦点就会在炫彩窗口上
func (w *WebView) EnableAutoFocus(bEnable bool) {
	xc.XWnd_EnableAutoFocus(w.hWindow, !bEnable)
	w.autoFocus = bEnable
}

// EnableFillParent 启用填充父窗口.
//   - 启用后 WebView 中设置的固定位置与大小会无效.
func (w *WebView) EnableFillParent(bEnable bool) {
	w.fillParent = bEnable
}

// IsAutoFocus 是否在窗口获得焦点时尝试保持 WebView 的焦点.
func (w *WebView) IsAutoFocus() bool {
	return w.autoFocus
}

// IsFillParent 是否填充父窗口.
func (w *WebView) IsFillParent() bool {
	return w.fillParent
}

package edge

import (
	"errors"
	"sync"
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/wapi/wnd"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// WebView 是创建在一个用 wapi 创建的原生窗口里的, 然后原生窗口是被嵌入到炫彩窗口或元素里的.
//
// 对 webview 的一些更改操作必须在炫彩 UI 线程里执行.
type WebView struct {
	WebView2

	// 读写锁bindings
	rwxBindings sync.RWMutex
	// 绑定go函数: name : func
	bindings map[string]interface{}
	// 绑定id: name : id
	bindingsid map[string]string

	updateWebviewSize func()
	hWindow           int // 炫彩窗口句柄
	hParent           int // 原生窗口宿主炫彩窗口或元素句柄

	// 圆角半径
	roundRadius int32

	// 在窗口获得焦点时尝试保持 webview 的焦点。
	autofocus bool
}

// NewWebView 创建 webview 到炫彩窗口或元素.
//
// hParent: 炫彩窗口或元素句柄.
//
// opts: webview 选项, 使用 edge.WithXXX 系列函数.
func (e *Edge) NewWebView(hParent int, opts ...WebViewOption) (*WebView, error) {
	if !xc.XC_IsHELE(hParent) && !xc.XC_IsHWINDOW(hParent) {
		return nil, errors.New("hParent is not a xcgui window or element")
	}

	// 获取默认 WebView 配置
	options := defaultWebViewOptions()
	// 应用所有选项
	for _, opt := range opts {
		opt(options)
	}

	w := &WebView{}
	w.bindings = map[string]interface{}{}
	w.bindingsid = map[string]string{}
	w.autofocus = options.AutoFocus
	w.Edge = e

	err := createWebViewWithOptionsByXcgui(w, hParent, options)
	if err != nil {
		return nil, err
	}
	return w, nil
}

// 创建 webview 宿主窗口.
func createWebViewWithOptionsByXcgui(w *WebView, hParent int, opt *WebViewOptions) error {
	w.hParent = hParent
	// 获取父窗口或元素的HWND, 其实是一个, 就是父窗口的HWND, 炫彩元素没有自己的HWND
	var hWndXC uintptr
	var isInWindow bool
	if w.hParent > 0 {
		if xc.XC_IsHWINDOW(w.hParent) {
			w.hWindow = w.hParent
			isInWindow = true
			hWndXC = xc.XWnd_GetHWND(w.hParent)
		} else if xc.XC_IsHELE(w.hParent) {
			hWndXC = xc.XWidget_GetHWND(w.hParent)
			w.hWindow = xc.XWidget_GetHWINDOW(w.hParent)
		}
	}

	dpi := int32(96)
	if w.hWindow > 0 {
		dpi = xc.XWnd_GetDPI(w.hWindow)
		// 启用自动焦点
		xc.XWnd_EnableAutoFocus(w.hWindow, true)
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
	w.hwnd = wapi.CreateWindowEx(0, opt.ClassName, opt.Title, xcc.WS_MINIMIZE, xc.DpiConvRound(dpi, opt.Left), xc.DpiConvRound(dpi, opt.Top), xc.DpiConvRound(dpi, opt.Width), xc.DpiConvRound(dpi, opt.Height), 0, 0, hInstance, 0)

	// 记录窗口上下文
	hwndContext.SetWindowContext(w.hwnd, w)
	xcContext.SetWindowContext(uintptr(w.hParent), w)

	// 显示窗口, 更新窗口, 设置焦点
	wapi.ShowWindow(w.hwnd, xcc.SW_SHOWMINIMIZED)
	wapi.UpdateWindow(w.hwnd)
	wapi.SetFocus(w.hwnd)

	// ------------------------ 创建 WebView2 控制器 ------------------------
	initWebViewEventImpl(w.GetWebViewEventImpl())
	w.WebView2.msgcb_xcgui = w.msgcb_xcgui

	err := newWebView2Controller(w, opt)
	if err != nil {
		wapi.SendMessageW(w.hwnd, wapi.WM_CLOSE, 0, 0) // 关闭原生窗口
		return err
	}
	// ------------------------ 创建 WebView2 控制器 END ------------------------

	// 设置 WebView2 宿主窗口为炫彩窗口或元素的子窗口
	wapi.SetParent(w.hwnd, hWndXC)
	// 设置 WebView2 宿主窗口样式
	wapi.SetWindowLongPtrW(w.hwnd, wapi.GWL_STYLE, int(xcc.WS_CHILD|xcc.WS_VISIBLE))

	// 更新 WebView2 宿主窗口大小和尺寸
	w.updateWebviewSize = func() {
		if !wapi.IsWindow(w.hwnd) {
			return
		}
		var rcClient xc.RECT // 炫彩窗口或元素的客户区矩形
		if isInWindow {
			xc.XWnd_GetBodyRect(w.hParent, &rcClient)
		} else {
			xc.XEle_GetWndClientRect(w.hParent, &rcClient)
		}

		var left, top, width, height int32
		// 填充父
		if opt.FillParent {
			left = rcClient.Left
			top = rcClient.Top
			height = rcClient.Bottom - rcClient.Top
			if isInWindow {
				var rcWindow xc.RECT
				xc.XWnd_GetRect(w.hParent, &rcWindow)
				// 原本应该只用客户区坐标来计算, 但是发现存在一个问题, 就是客户区坐标的 Right 会莫名奇妙的比窗口创建好时多出 1px, 明明窗口 Rect 没有任何改变, 这导致 webview 的右边超出了炫彩窗口 2px.
				// 目前采用窗口宽度减去客户区左边的两倍来计算客户区宽度也就是 webview 的宽度, 这等于是假设客户区是左右对称的, 实际上客户区左边和右边可能会有不同的宽度(虽然可能没有这种需求), 但是目前没有更好的办法.
				width = rcWindow.Right - rcWindow.Left - rcClient.Left*2
			} else {
				width = rcClient.Right - rcClient.Left
			}
		} else {
			left = rcClient.Left + opt.Left
			top = rcClient.Top + opt.Top
			width = opt.Width
			height = opt.Height
		}
		wapi.SetWindowPos(w.hwnd, 0, xc.DpiConvRound(dpi, left), xc.DpiConvRound(dpi, top), xc.DpiConvRound(dpi, width), xc.DpiConvRound(dpi, height), wapi.SWP_NOACTIVATE|wapi.SWP_NOZORDER)
	}

	// 窗口 调整位置和大小
	xc.XWnd_RemoveEventC(w.hWindow, xcc.XWM_WINDPROC, onWndProc)
	xc.XWnd_RegEventC1(w.hWindow, xcc.XWM_WINDPROC, onWndProc)

	// 元素事件
	if !isInWindow {
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

	w.updateWebviewSize()
	return nil
}

func onEleDestroy(hEle int, pbHandled *bool) int {
	handle := uintptr(hEle)
	if w := xcContext.GetWindowContext(handle); w != nil {
		if wapi.IsWindow(w.hwnd) {
			wapi.SendMessageW(w.hwnd, wapi.WM_CLOSE, 0, 0)
		}
		xcContext.DeleteWindowContext(handle)
	}
	return 0
}

func onEleSize(hEle int, nFlags xcc.AdjustLayout_, nAdjustNo uint32, pbHandled *bool) int {
	if w := xcContext.GetWindowContext(uintptr(hEle)); w != nil {
		w.updateWebviewSize()
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
			if wParam != wapi.SIZE_MINIMIZED { // 最小化时不执行 updateWebviewSize
				w.updateWebviewSize()
			}
		} else { // 更新每个元素中的 webview 位置
			hEles := xcContext.GetHEles(hWindow)
			for i := 0; i < len(hEles); i++ {
				if w := xcContext.GetWindowContext(uintptr(hEles[i])); w != nil {
					if wParam != wapi.SIZE_MINIMIZED {
						w.updateWebviewSize()
					}
				}
			}
		}
	case wapi.WM_CLOSE:
		handle := uintptr(hWindow)
		if w := xcContext.GetWindowContext(handle); w != nil { // 原生窗口宿主是炫彩窗口
			xc.XWnd_Show(hWindow, false) // 当窗口有透明通道时, 关闭的时候会漏出炫彩窗口, 导致闪烁一下, 所以先隐藏炫彩窗口
			if wapi.IsWindow(w.hwnd) {
				wapi.SendMessageW(w.hwnd, wapi.WM_CLOSE, 0, 0)
			}
			xcContext.DeleteWindowContext(handle)
		} else { // 触发元素销毁事件
			hEles := xcContext.GetHEles(hWindow)
			for i := 0; i < len(hEles); i++ {
				xc.XEle_SendEvent(hEles[i], xcc.XE_DESTROY, 0, 0)
			}
		}
	}
	return 0
}

// SetRoundRadius 设置原生窗口的圆角半径。
//   - 需注意炫彩窗口的圆角需通过 SetShadowInfo 设置.
//   - webview 是创建在一个用 wapi 创建的原生窗口里的, 然后原生窗口是被嵌入到炫彩窗口或元素里的.
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

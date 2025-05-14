package edge

import (
	"encoding/json"
	"errors"
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
	"strconv"
	"sync"
	"syscall"
	"unsafe"
)

// WebViewOption 选项.
type WebViewOption struct {
	// webview 宿主窗口标题
	Title string
	// webview 宿主窗口类名
	ClassName string

	// load icon from resource
	IconId uint

	// webview 左边
	Left int32
	// webview 顶边
	Top int32
	// webview 宽度
	Width int32
	// webview 高度
	Height int32

	// 填充父, 如果为 true, 则 webview 会填满父窗口或元素, 固定坐标和尺寸会失效.
	FillParent bool

	// Debug 是否可打开开发者工具.
	Debug bool
	// AppDrag 是否启用非客户区域支持。WebView2 运行时版本大于 123.0.2420.47 时有效。
	//   - 当此属性为 true 时，将启用所有非客户端区域功能：将启用可拖动区域，它们是网页上用 CSS 属性 app-region:drag/no-drag 标记的区域。
	//   - 设置为拖动时，这些区域将被视为窗口的标题栏，支持拖动整个 webview 及其宿主应用程序窗口；
	//   - 系统菜单在右键单击时显示，双击将触发最大化/恢复窗口大小。
	AppDrag bool

	// AutoFocus 将在窗口获得焦点时尝试保持 webView 的焦点。
	AutoFocus bool
}

// WebView 是创建在一个用 wapi 创建的原生窗口里的, 然后原生窗口是被嵌入到炫彩窗口里的.
//
// 对 webview 的一些更改操作必须在炫彩 UI 线程里执行.
type WebView struct {
	webview2

	// 读写锁bindings
	rwxBindings sync.RWMutex
	// 绑定go函数: name : func
	bindings map[string]interface{}
	// 绑定id: name : id
	bindingsid map[string]string

	updateWebviewSize func()
	hWindow           int // 炫彩窗口句柄
	hParent           int

	// 在窗口获得焦点时尝试保持 webView 的焦点。
	autofocus bool
}

// NewWebView 创建 webview 窗口到炫彩窗口或元素.
//
// hParent: 炫彩窗口或元素句柄.
//
// opt: 选项.
func (e *Edge) NewWebView(hParent int, opt WebViewOption) (*WebView, error) {
	if !xc.XC_IsHELE(hParent) && !xc.XC_IsHWINDOW(hParent) {
		return nil, errors.New("hParent is not a xcgui window or element")
	}

	w := &WebView{}
	w.bindings = map[string]interface{}{}
	w.bindingsid = map[string]string{}
	w.autofocus = opt.AutoFocus
	w.Edge = e

	err := w.createWithOptionsByXcgui(hParent, opt)
	if err != nil {
		return nil, err
	}
	return w, nil
}

// 创建 webview 宿主窗口.
func (w *WebView) createWithOptionsByXcgui(hParent int, opt WebViewOption) error {
	w.hParent = hParent
	// 获取父窗口或元素的HWND
	var hWnd uintptr
	var isWindow bool
	if w.hParent > 0 {
		if xc.XC_IsHWINDOW(w.hParent) {
			w.hWindow = w.hParent
			isWindow = true
			hWnd = xc.XWnd_GetHWND(w.hParent)
		} else if xc.XC_IsHELE(w.hParent) {
			hWnd = xc.XWidget_GetHWND(w.hParent)
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

	// 窗口坐标和宽高
	var left, top, width, height = xc.DpiConv(dpi, opt.Left), xc.DpiConv(dpi, opt.Top), xc.DpiConv(dpi, opt.Width), xc.DpiConv(dpi, opt.Height)

	// 创建宿主窗口
	w.hwnd = wapi.CreateWindowEx(0, opt.ClassName, opt.Title, xcc.WS_MINIMIZE, left, top, width, height, 0, 0, hInstance, 0)

	setWindowContext(w.hwnd, w)
	setWindowContext(hWnd, w)

	// 显示窗口, 更新窗口, 设置焦点
	wapi.ShowWindow(w.hwnd, xcc.SW_SHOW)
	wapi.UpdateWindow(w.hwnd)
	wapi.SetFocus(w.hwnd)

	// ------------------------ 创建 WebView2 控制器 ------------------------
	w.init()
	w.webview2.msgcb_xcgui = w.msgcb_xcgui
	w.SetPermission(COREWEBVIEW2_PERMISSION_KIND_CLIPBOARD_READ, COREWEBVIEW2_PERMISSION_STATE_ALLOW)

	err := w.newWebView2Controller()
	if err != nil {
		// 销毁原生窗口
		wapi.PostMessageW(w.hwnd, wapi.WM_CLOSE, 0, 0)
		deleteWindowContext(w.hwnd)
		return err
	}

	// 获取浏览器设置
	settings, _ := w.GetSettings()
	if settings != nil {
		// 设置是否可开启开发人员工具
		_ = settings.PutAreDevToolsEnabled(opt.Debug)
		// 设置是否启用非客户区域支持
		s9, _ := settings.GetICoreWebView2Settings9()
		if s9 != nil {
			_ = s9.PutIsNonClientRegionSupportEnabled(opt.AppDrag)
		}
	}
	w.Resize()
	// ------------------------ 创建 WebView2 控制器 END ------------------------

	// 设置 WebView2 宿主窗口为炫彩父窗口或元素的子窗口
	wapi.SetParent(w.hwnd, hWnd)
	// 设置 WebView2 宿主窗口样式
	wapi.SetWindowLongPtrW(w.hwnd, wapi.GWL_STYLE, int(xcc.WS_CHILD|xcc.WS_VISIBLE))

	// 更新 WebView2 宿主窗口大小和尺寸
	w.updateWebviewSize = func() {
		if !wapi.IsWindow(w.hwnd) {
			return
		}
		var rc xc.RECT
		if isWindow {
			xc.XWnd_GetBodyRect(w.hParent, &rc)
		} else {
			xc.XEle_GetWndClientRect(w.hParent, &rc)
		}
		// 填充父
		if opt.FillParent {
			left = xc.DpiConv(dpi, rc.Left)
			top = xc.DpiConv(dpi, rc.Top)
			width = xc.DpiConv(dpi, rc.Right-rc.Left)
			height = xc.DpiConv(dpi, rc.Bottom-rc.Top)
		}

		wapi.SetWindowPos(w.hwnd, 0, left, top, width, height, 0)
	}

	// 窗口 调整位置和大小
	xc.XWnd_RemoveEventC(w.hWindow, xcc.XWM_WINDPROC, onWndProc)
	xc.XWnd_RegEventC1(w.hWindow, xcc.XWM_WINDPROC, onWndProc)

	// 元素事件
	if !isWindow {
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
	if w, ok := getWindowContext(xc.XWidget_GetHWND(hEle)).(*WebView); ok {
		if wapi.IsWindow(w.hwnd) {
			wapi.PostMessageW(w.hwnd, wapi.WM_CLOSE, 0, 0)
		}
	}
	return 0
}

func onEleSize(hEle int, nFlags xcc.AdjustLayout_, nAdjustNo uint32, pbHandled *bool) int {
	if w, ok := getWindowContext(xc.XWidget_GetHWND(hEle)).(*WebView); ok {
		w.updateWebviewSize()
	}
	return 0
}

func onWndProc(hWindow int, message uint32, wParam, lParam uintptr, pbHandled *bool) int {
	if message == wapi.WM_SIZE {
		if w, ok := getWindowContext(xc.XWnd_GetHWND(hWindow)).(*WebView); ok {
			w.updateWebviewSize()
		}
	}
	return 0
}

func onEleShow(hEle int, bShow bool, pbHandled *bool) int {
	if w, ok := getWindowContext(xc.XWidget_GetHWND(hEle)).(*WebView); ok {
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

func (w *WebView) msgcb_xcgui(msg string) {
	d := rpcMessage{}
	if err := json.Unmarshal(common.String2Bytes(msg), &d); err != nil {
		return
	}

	id := strconv.Itoa(d.ID)
	if res, err := w.callbinding(&d); err != nil {
		xc.XC_CallUT(func() {
			w.Eval("window._rpc[" + id + "].reject(" + jsString(err.Error()) + "); window._rpc[" + id + "] = undefined")
		})
	} else if b, err := json.Marshal(res); err != nil {
		xc.XC_CallUT(func() {
			w.Eval("window._rpc[" + id + "].reject(" + jsString(err.Error()) + "); window._rpc[" + id + "] = undefined")
		})
	} else {
		xc.XC_CallUT(func() {
			w.Eval("window._rpc[" + id + "].resolve(" + string(b) + "); window._rpc[" + id + "] = undefined")
		})
	}
}

// SetTitle 更新原生窗口的标题。
//   - webview 是创建在一个用 wapi 创建的原生窗口里的, 然后原生窗口是被嵌入到炫彩窗口里的.
func (w *WebView) SetTitle(title string) {
	wapi.SetWindowText(w.hwnd, title)
}

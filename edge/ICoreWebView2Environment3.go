package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2Environment3 是 ICoreWebView2Environment2 的延续.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2environment3
type ICoreWebView2Environment3 struct {
	Vtbl *ICoreWebView2Environment3Vtbl
}

type ICoreWebView2Environment3Vtbl struct {
	ICoreWebView2Environment2Vtbl
	CreateCoreWebView2CompositionController ComProc
	CreateCoreWebView2PointerInfo           ComProc
}

func (i *ICoreWebView2Environment3) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Environment3) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Environment3) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// CreateCoreWebView2CompositionController 异步创建一个新的WebView，以供可视化承载使用。
//   - parentWindow 是应用将连接WebView可视化树的窗口句柄（HWND）。
//   - 应用将通过此窗口句柄接收WebView的指针/鼠标输入（并且需要使用 SendMouseInput / SendPointerInput 进行转发）。
//   - 如果应用将WebView可视化树移动到其他窗口下，则需要调用 SetParentWindow 来更新可视化树的新父窗口句柄。
func (i *ICoreWebView2Environment3) CreateCoreWebView2CompositionController(parentWindow uintptr, handler *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler) error {
	r, _, _ := i.Vtbl.CreateCoreWebView2CompositionController.Call(
		uintptr(unsafe.Pointer(i)),
		parentWindow,
		uintptr(unsafe.Pointer(handler)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// CreateCoreWebView2CompositionControllerEx 异步创建一个新的WebView，以供可视化承载使用。
//   - 应用将通过此窗口句柄接收WebView的指针/鼠标输入（并且需要使用 SendMouseInput / SendPointerInput 进行转发）。
//   - 如果应用将WebView可视化树移动到其他窗口下，则需要调用 SetParentWindow 来更新可视化树的新父窗口句柄。
//
// impl: *WebViewEventImpl。
//
// parentWindow: 应用将连接WebView可视化树的窗口句柄（HWND）。
//
// cb: 创建完成执行后的回调函数。
func (i *ICoreWebView2Environment3) CreateCoreWebView2CompositionControllerEx(impl *WebViewEventImpl, parentWindow uintptr, cb func(errorCode syscall.Errno, controller *ICoreWebView2CompositionController) uintptr) error {
	handler := WvEventHandler.GetHandler(impl, "CreateCoreWebView2CompositionControllerCompleted")
	if handler == nil {
		var c interface{}
		if cb == nil {
			c = nil
		} else {
			c = cb
		}
		_, _ = WvEventHandler.AddCallBack(impl, "CreateCoreWebView2CompositionControllerCompleted", c, nil)
		handler = WvEventHandler.GetHandler(impl, "CreateCoreWebView2CompositionControllerCompleted")
	}
	return i.CreateCoreWebView2CompositionController(parentWindow, (*ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler)(handler))
}

// CreateCoreWebView2PointerInfo 创建一个空的 ICoreWebView2PointerInfo。
//   - 返回的 ICoreWebView2PointerInfo 需要在调用 SendPointerInput 之前填充所有相关信息。
func (i *ICoreWebView2Environment3) CreateCoreWebView2PointerInfo() (*ICoreWebView2PointerInfo, error) {
	var pointerInfo *ICoreWebView2PointerInfo
	r, _, _ := i.Vtbl.CreateCoreWebView2PointerInfo.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&pointerInfo)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return pointerInfo, nil
}

// MustCreateCoreWebView2PointerInfo 创建一个空的 ICoreWebView2PointerInfo. 出错时会触发全局错误回调。
func (i *ICoreWebView2Environment3) MustCreateCoreWebView2PointerInfo() *ICoreWebView2PointerInfo {
	pointerInfo, err := i.CreateCoreWebView2PointerInfo()
	ReportErrorAtuo(err)
	return pointerInfo
}

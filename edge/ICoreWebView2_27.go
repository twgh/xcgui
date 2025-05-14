package edge

// ICoreWebView2_27 是 ICoreWebView2_26 接口的延续，支持 屏幕捕获开始 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_27
type ICoreWebView2_27 struct {
	Vtbl *ICoreWebView2_27Vtbl
}

type ICoreWebView2_27Vtbl struct {
	ICoreWebView2_26Vtbl
	AddScreenCaptureStarting    ComProc
	RemoveScreenCaptureStarting ComProc
}

/*TODO:
AddScreenCaptureStarting
RemoveScreenCaptureStarting
*/

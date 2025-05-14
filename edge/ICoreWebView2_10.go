package edge

// ICoreWebView2_10 是 ICoreWebView2_9 接口的延续，支持 请求基本身份验证 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_10
type ICoreWebView2_10 struct {
	Vtbl *ICoreWebView2_10Vtbl
}

type ICoreWebView2_10Vtbl struct {
	ICoreWebView2_9Vtbl
	AddBasicAuthenticationRequested    ComProc
	RemoveBasicAuthenticationRequested ComProc
}

/*TODO:
AddBasicAuthenticationRequested
RemoveBasicAuthenticationRequested
*/

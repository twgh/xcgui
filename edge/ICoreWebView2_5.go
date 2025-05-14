package edge

// ICoreWebView2_5 是 ICoreWebView2_4 接口的延续，以支持 请求客户端证书 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_5
type ICoreWebView2_5 struct {
	Vtbl *ICoreWebView2_5Vtbl
}

type ICoreWebView2_5Vtbl struct {
	ICoreWebView2_4Vtbl
	AddClientCertificateRequested    ComProc
	RemoveClientCertificateRequested ComProc
}

/*TODO:
AddClientCertificateRequested
RemoveClientCertificateRequested
*/

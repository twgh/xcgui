package edge

// ICoreWebView2_22 是 ICoreWebView2_21 接口的延续，允许设置筛选器，以便接收服务工作线程、共享工作线程和不同源 iframe 的 已请求web资源 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_22
type ICoreWebView2_22 struct {
	Vtbl *ICoreWebView2_22Vtbl
}

type ICoreWebView2_22Vtbl struct {
	ICoreWebView2_21Vtbl
	AddWebResourceRequestedFilterWithRequestSourceKinds    ComProc
	RemoveWebResourceRequestedFilterWithRequestSourceKinds ComProc
}

/*TODO:
AddWebResourceRequestedFilterWithRequestSourceKinds
RemoveWebResourceRequestedFilterWithRequestSourceKinds
*/

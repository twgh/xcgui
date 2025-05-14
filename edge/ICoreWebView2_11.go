package edge

// ICoreWebView2_11 是 ICoreWebView2_10 接口的延续，支持 CDP 方法调用的 sessionId 和 已请求上下文菜单 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_11
type ICoreWebView2_11 struct {
	Vtbl *ICoreWebView2_11Vtbl
}

type ICoreWebView2_11Vtbl struct {
	ICoreWebView2_10Vtbl
	CallDevToolsProtocolMethodForSession ComProc
	AddContextMenuRequested              ComProc
	RemoveContextMenuRequested           ComProc
}

/*TODO:
CallDevToolsProtocolMethodForSession
AddContextMenuRequested
RemoveContextMenuRequested
*/

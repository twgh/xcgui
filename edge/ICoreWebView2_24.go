package edge

// ICoreWebView2_24 是 ICoreWebView2_23 接口的延续，支持 通知接收 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_24
type ICoreWebView2_24 struct {
	Vtbl *ICoreWebView2_24Vtbl
}

type ICoreWebView2_24Vtbl struct {
	ICoreWebView2_23Vtbl
	AddNotificationReceived    ComProc
	RemoveNotificationReceived ComProc
}

/*TODO:
AddNotificationReceived
RemoveNotificationReceived
*/

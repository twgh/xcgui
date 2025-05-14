package edge

// ICoreWebView2_23 是 ICoreWebView2_22 接口的延续，支持发送带有附加对象的 JSON Web 消息。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_23
type ICoreWebView2_23 struct {
	Vtbl *ICoreWebView2_23Vtbl
}

type ICoreWebView2_23Vtbl struct {
	ICoreWebView2_22Vtbl
	PostWebMessageAsJsonWithAdditionalObjects ComProc
}

/*TODO:
PostWebMessageAsJsonWithAdditionalObjects
*/

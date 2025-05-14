package edge

// ICoreWebView2_13 是 ICoreWebView2_12 接口的延续，支持获取关联的 ICoreWebView2Profile 对象。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_13
type ICoreWebView2_13 struct {
	Vtbl *ICoreWebView2_13Vtbl
}

type ICoreWebView2_13Vtbl struct {
	ICoreWebView2_12Vtbl
	GetProfile ComProc
}

/*TODO:
GetProfile
*/

package edge

// ICoreWebView2_7 是 ICoreWebView2_6 接口的延续，支持打印为PDF。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_7
type ICoreWebView2_7 struct {
	Vtbl *ICoreWebView2_7Vtbl
}

type ICoreWebView2_7Vtbl struct {
	ICoreWebView2_6Vtbl
	PrintToPdf ComProc
}

/*TODO:
PrintToPdf
*/

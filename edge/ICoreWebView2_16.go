package edge

// ICoreWebView2_16 是 ICoreWebView2_15 接口的延续，支持打印相关功能。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_16
type ICoreWebView2_16 struct {
	Vtbl *ICoreWebView2_16Vtbl
}

type ICoreWebView2_16Vtbl struct {
	ICoreWebView2_15Vtbl
	Print            ComProc
	ShowPrintUI      ComProc
	PrintToPdfStream ComProc
}

/*TODO:
Print
ShowPrintUI
PrintToPdfStream
*/

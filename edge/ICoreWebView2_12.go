package edge

// ICoreWebView2_12 是 ICoreWebView2_11 接口的延续，支持 状态栏文本已更改 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_12
type ICoreWebView2_12 struct {
	Vtbl *ICoreWebView2_12Vtbl
}

type ICoreWebView2_12Vtbl struct {
	ICoreWebView2_11Vtbl
	AddStatusBarTextChanged    ComProc
	RemoveStatusBarTextChanged ComProc
	GetStatusBarText           ComProc
}

/*TODO:
AddStatusBarTextChanged
RemoveStatusBarTextChanged
GetStatusBarText
*/

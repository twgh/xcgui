package edge

// ICoreWebView2_25 是 ICoreWebView2_24 接口的延续，支持"另存为"UI相关功能。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_25
type ICoreWebView2_25 struct {
	Vtbl *ICoreWebView2_25Vtbl
}

type ICoreWebView2_25Vtbl struct {
	ICoreWebView2_24Vtbl
	AddSaveAsUIShowing    ComProc
	RemoveSaveAsUIShowing ComProc
	ShowSaveAsUI          ComProc
}

/*TODO:
AddSaveAsUIShowing
RemoveSaveAsUIShowing
ShowSaveAsUI
*/

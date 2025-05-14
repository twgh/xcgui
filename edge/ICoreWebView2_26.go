package edge

// ICoreWebView2_26 是 ICoreWebView2_25 接口的延续，支持 开始保存文件安全检查 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_26
type ICoreWebView2_26 struct {
	Vtbl *ICoreWebView2_26Vtbl
}

type ICoreWebView2_26Vtbl struct {
	ICoreWebView2_25Vtbl
	AddSaveFileSecurityCheckStarting    ComProc
	RemoveSaveFileSecurityCheckStarting ComProc
}

/*TODO:
AddSaveFileSecurityCheckStarting
RemoveSaveFileSecurityCheckStarting
*/

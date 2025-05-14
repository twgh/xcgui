package edge

// ICoreWebView2_21 是 ICoreWebView2_20 接口的延续，支持执行脚本并获取字符串和异常。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_21
type ICoreWebView2_21 struct {
	Vtbl *ICoreWebView2_21Vtbl
}

type ICoreWebView2_21Vtbl struct {
	ICoreWebView2_20Vtbl
	ExecuteScriptWithResult ComProc
}

/*TODO:
ExecuteScriptWithResult
*/

package edge

// ICoreWebView2_6 是 ICoreWebView2_5 接口的延续，用于管理打开浏览器任务管理器窗口。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_6
type ICoreWebView2_6 struct {
	Vtbl *ICoreWebView2_6Vtbl
}

type ICoreWebView2_6Vtbl struct {
	ICoreWebView2_5Vtbl
	OpenTaskManagerWindow ComProc
}

/*TODO:
OpenTaskManagerWindow
*/

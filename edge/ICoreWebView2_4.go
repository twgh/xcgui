package edge

// ICoreWebView2_4 是 ICoreWebView2_3 接口的延续，以支持 FrameCreated 和 下载开始 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_4
type ICoreWebView2_4 struct {
	Vtbl *ICoreWebView2_4Vtbl
}

type ICoreWebView2_4Vtbl struct {
	ICoreWebView2_3Vtbl
	AddFrameCreated        ComProc
	RemoveFrameCreated     ComProc
	AddDownloadStarting    ComProc
	RemoveDownloadStarting ComProc
}

/*TODO:
AddFrameCreated
RemoveFrameCreated
AddDownloadStarting
RemoveDownloadStarting
*/

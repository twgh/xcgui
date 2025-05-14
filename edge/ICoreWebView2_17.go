package edge

// ICoreWebView2_17 是 ICoreWebView2_16 接口的延续，支持基于文件映射的共享缓冲区。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_17
type ICoreWebView2_17 struct {
	Vtbl *ICoreWebView2_17Vtbl
}

type ICoreWebView2_17Vtbl struct {
	ICoreWebView2_16Vtbl
	PostSharedBufferToScript ComProc
}

/*TODO:
PostSharedBufferToScript
*/

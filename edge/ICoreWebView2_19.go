package edge

// ICoreWebView2_19 是 ICoreWebView2_18 接口的延续，用于管理内存使用目标级别。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_19
type ICoreWebView2_19 struct {
	Vtbl *ICoreWebView2_19Vtbl
}

type ICoreWebView2_19Vtbl struct {
	ICoreWebView2_18Vtbl
	GetMemoryUsageTargetLevel ComProc
	PutMemoryUsageTargetLevel ComProc
}

/*TODO:
GetMemoryUsageTargetLevel
PutMemoryUsageTargetLevel
*/

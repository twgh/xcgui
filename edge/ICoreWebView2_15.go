package edge

// ICoreWebView2_15 是 ICoreWebView2_14 接口的延续，支持网站图标(Favicon)相关功能。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_15
type ICoreWebView2_15 struct {
	Vtbl *ICoreWebView2_15Vtbl
}

type ICoreWebView2_15Vtbl struct {
	ICoreWebView2_14Vtbl
	AddFaviconChanged    ComProc
	RemoveFaviconChanged ComProc
	GetFaviconUri        ComProc
	GetFavicon           ComProc
}

/*TODO:
AddFaviconChanged
RemoveFaviconChanged
GetFaviconUri
GetFavicon
*/

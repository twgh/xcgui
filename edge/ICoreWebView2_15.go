package edge

import (
	"syscall"
	"unsafe"
)

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

func (i *ICoreWebView2_15) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_15) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_15) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

/*TODO:
AddFaviconChanged
RemoveFaviconChanged
GetFaviconUri
GetFavicon
*/

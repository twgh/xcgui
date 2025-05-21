package edge

import (
	"errors"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

// ICoreWebView2_9 是 ICoreWebView2_8 接口的延续，默认下载对话框定位和锚定。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_9
type ICoreWebView2_9 struct {
	Vtbl *ICoreWebView2_9Vtbl
}

type ICoreWebView2_9Vtbl struct {
	ICoreWebView2_8Vtbl
	AddIsDefaultDownloadDialogOpenChanged    ComProc
	RemoveIsDefaultDownloadDialogOpenChanged ComProc
	GetIsDefaultDownloadDialogOpen           ComProc
	OpenDefaultDownloadDialog                ComProc
	CloseDefaultDownloadDialog               ComProc
	GetDefaultDownloadDialogCornerAlignment  ComProc
	PutDefaultDownloadDialogCornerAlignment  ComProc
	GetDefaultDownloadDialogMargin           ComProc
	PutDefaultDownloadDialogMargin           ComProc
}

func (i *ICoreWebView2_9) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_9) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_9) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

/*TODO:
AddIsDefaultDownloadDialogOpenChanged
RemoveIsDefaultDownloadDialogOpenChanged
GetIsDefaultDownloadDialogOpen
OpenDefaultDownloadDialog
CloseDefaultDownloadDialog
GetDefaultDownloadDialogCornerAlignment
PutDefaultDownloadDialogCornerAlignment
GetDefaultDownloadDialogMargin
PutDefaultDownloadDialogMargin
*/

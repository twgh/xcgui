package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2File 是通过 WebMessage 传递的 DOM File 对象的表示形式。
//   - 可以使用此对象获取拖放到 WebView2 上的文件的路径。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2file
type ICoreWebView2File struct {
	Vtbl *ICoreWebView2FileVtbl
}

type ICoreWebView2FileVtbl struct {
	IUnknownVtbl
	GetPath ComProc
}

func (i *ICoreWebView2File) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2File) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2File) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetPath 获取文件的完整路径.
func (i *ICoreWebView2File) GetPath() (string, error) {
	var _path *uint16
	r, _, _ := i.Vtbl.GetPath.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_path)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	path := common.UTF16PtrToString(_path)
	wapi.CoTaskMemFree(unsafe.Pointer(_path))
	return path, nil
}

// MustGetPath 获取文件的完整路径。出错时会触发全局错误回调。
func (i *ICoreWebView2File) MustGetPath() string {
	path, err := i.GetPath()
	ReportErrorAtuo(err)
	return path
}

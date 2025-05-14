package edge

import (
	"errors"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

// ICoreWebView2Settings7 是 ICoreWebView2Settings6 接口的延续, 用于隐藏 Pdf 工具栏项。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2settings7
type ICoreWebView2Settings7 struct {
	Vtbl *ICoreWebView2Settings7Vtbl
}

type ICoreWebView2Settings7Vtbl struct {
	ICoreWebView2Settings6Vtbl
	GetHiddenPdfToolbarItems ComProc
	PutHiddenPdfToolbarItems ComProc
}

// GetHiddenPdfToolbarItems 获取隐藏的 PDF 工具栏项目。
func (i *ICoreWebView2Settings7) GetHiddenPdfToolbarItems() (COREWEBVIEW2_PDF_TOOLBAR_ITEMS, error) {
	var items COREWEBVIEW2_PDF_TOOLBAR_ITEMS
	r, _, err := i.Vtbl.GetHiddenPdfToolbarItems.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&items)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return items, nil
}

// MustGetHiddenPdfToolbarItems 获取 PDF 工具栏项。忽略错误。
func (i *ICoreWebView2Settings7) MustGetHiddenPdfToolbarItems() COREWEBVIEW2_PDF_TOOLBAR_ITEMS {
	items, _ := i.GetHiddenPdfToolbarItems()
	return items
}

// PutHiddenPdfToolbarItems 用于自定义 PDF 工具栏项。
func (i *ICoreWebView2Settings7) PutHiddenPdfToolbarItems(items COREWEBVIEW2_PDF_TOOLBAR_ITEMS) error {
	r, _, err := i.Vtbl.PutHiddenPdfToolbarItems.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(items),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

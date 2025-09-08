package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2Settings7 是 ICoreWebView2Settings6 接口的延续, 用于隐藏 Pdf 工具栏项。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2settings7
type ICoreWebView2Settings7 struct {
	ICoreWebView2Settings6
}

// GetHiddenPdfToolbarItems 获取隐藏的 PDF 工具栏项目。
func (i *ICoreWebView2Settings7) GetHiddenPdfToolbarItems() (COREWEBVIEW2_PDF_TOOLBAR_ITEMS, error) {
	var items COREWEBVIEW2_PDF_TOOLBAR_ITEMS
	r, _, _ := i.Vtbl.GetHiddenPdfToolbarItems.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&items)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return items, nil
}

// SetHiddenPdfToolbarItems 用于自定义 PDF 工具栏项。
func (i *ICoreWebView2Settings7) SetHiddenPdfToolbarItems(items COREWEBVIEW2_PDF_TOOLBAR_ITEMS) error {
	r, _, _ := i.Vtbl.PutHiddenPdfToolbarItems.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(items),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetHiddenPdfToolbarItems 获取 PDF 工具栏项。出错时会触发全局错误回调。
func (i *ICoreWebView2Settings7) MustGetHiddenPdfToolbarItems() COREWEBVIEW2_PDF_TOOLBAR_ITEMS {
	items, err := i.GetHiddenPdfToolbarItems()
	ReportErrorAuto(err)
	return items
}

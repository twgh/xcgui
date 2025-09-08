package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2PrintSettings2 是 ICoreWebView2PrintSettings 的延续。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2printsettings2
type ICoreWebView2PrintSettings2 struct {
	ICoreWebView2PrintSettings
}

// GetPageRanges 获取要打印的页面范围。
// 返回一个字符串，表示要打印的页面范围（如 "1-5,8"）。
func (i *ICoreWebView2PrintSettings2) GetPageRanges() (string, error) {
	var _pageRanges *uint16
	r, _, _ := i.Vtbl.GetPageRanges.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_pageRanges)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	pageRanges := common.UTF16PtrToString(_pageRanges)
	wapi.CoTaskMemFree(unsafe.Pointer(_pageRanges))
	return pageRanges, nil
}

// SetPageRanges 设置要打印的页面范围。
//   - pageRanges: 要打印的页面范围字符串（如 "1-5,8"）。
//   - 空字符串表示打印所有页面。
func (i *ICoreWebView2PrintSettings2) SetPageRanges(pageRanges string) error {
	pageRangesPtr, err := syscall.UTF16PtrFromString(pageRanges)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.PutPageRanges.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(pageRangesPtr)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetPagesPerSide 获取每张纸上的页面数量。
func (i *ICoreWebView2PrintSettings2) GetPagesPerSide() (int32, error) {
	var value int32
	r, _, _ := i.Vtbl.GetPagesPerSide.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// SetPagesPerSide 设置每张纸上的页面数量。
//   - pagesPerSide: 每张纸上的页面数量。有效值为 1、2、4、6、9 或 16。
func (i *ICoreWebView2PrintSettings2) SetPagesPerSide(pagesPerSide int32) error {
	r, _, _ := i.Vtbl.PutPagesPerSide.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(pagesPerSide),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetCopies 获取打印的副本数量。
func (i *ICoreWebView2PrintSettings2) GetCopies() (int32, error) {
	var value int32
	r, _, _ := i.Vtbl.GetCopies.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// SetCopies 设置打印的副本数量。
//   - copies: 打印的副本数量。必须大于 0, 最大为 999。
func (i *ICoreWebView2PrintSettings2) SetCopies(copies int32) error {
	r, _, _ := i.Vtbl.PutCopies.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(copies),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetCollation 获取打印排序模式。
func (i *ICoreWebView2PrintSettings2) GetCollation() (COREWEBVIEW2_PRINT_COLLATION, error) {
	var value COREWEBVIEW2_PRINT_COLLATION
	r, _, _ := i.Vtbl.GetCollation.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// SetCollation 设置打印排序模式。
//   - collation: 排序模式（默认、逐份打印或非逐份打印）。
func (i *ICoreWebView2PrintSettings2) SetCollation(collation COREWEBVIEW2_PRINT_COLLATION) error {
	r, _, _ := i.Vtbl.PutCollation.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(collation),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetColorMode 获取打印颜色模式。
func (i *ICoreWebView2PrintSettings2) GetColorMode() (COREWEBVIEW2_PRINT_COLOR_MODE, error) {
	var value COREWEBVIEW2_PRINT_COLOR_MODE
	r, _, _ := i.Vtbl.GetColorMode.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// SetColorMode 设置打印颜色模式。
//   - colorMode: 颜色模式（默认、彩色或灰度）。
func (i *ICoreWebView2PrintSettings2) SetColorMode(colorMode COREWEBVIEW2_PRINT_COLOR_MODE) error {
	r, _, _ := i.Vtbl.PutColorMode.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(colorMode),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetDuplex 获取双面打印模式。
func (i *ICoreWebView2PrintSettings2) GetDuplex() (COREWEBVIEW2_PRINT_DUPLEX, error) {
	var value COREWEBVIEW2_PRINT_DUPLEX
	r, _, _ := i.Vtbl.GetDuplex.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// SetDuplex 设置双面打印模式。
//   - duplex: 双面打印模式（默认、单面、双面长边装订或双面短边装订）。
func (i *ICoreWebView2PrintSettings2) SetDuplex(duplex COREWEBVIEW2_PRINT_DUPLEX) error {
	r, _, _ := i.Vtbl.PutDuplex.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(duplex),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetMediaSize 获取打印纸张大小。
func (i *ICoreWebView2PrintSettings2) GetMediaSize() (COREWEBVIEW2_PRINT_MEDIA_SIZE, error) {
	var value COREWEBVIEW2_PRINT_MEDIA_SIZE
	r, _, _ := i.Vtbl.GetMediaSize.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// SetMediaSize 设置打印纸张大小。
//   - mediaSize: 纸张大小（默认或自定义）。
func (i *ICoreWebView2PrintSettings2) SetMediaSize(mediaSize COREWEBVIEW2_PRINT_MEDIA_SIZE) error {
	r, _, _ := i.Vtbl.PutMediaSize.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(mediaSize),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetPrinterName 获取打印机名称。
func (i *ICoreWebView2PrintSettings2) GetPrinterName() (string, error) {
	var _printerName *uint16
	r, _, _ := i.Vtbl.GetPrinterName.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_printerName)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	printerName := common.UTF16PtrToString(_printerName)
	wapi.CoTaskMemFree(unsafe.Pointer(_printerName))
	return printerName, nil
}

// SetPrinterName 设置打印机名称。
//   - printerName: 要使用的打印机名称。
//   - 空字符串表示使用默认打印机。
func (i *ICoreWebView2PrintSettings2) SetPrinterName(printerName string) error {
	printerNamePtr, err := syscall.UTF16PtrFromString(printerName)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.PutPrinterName.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(printerNamePtr)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetPageRanges 获取要打印的页面范围。出错时会触发全局错误回调。
func (i *ICoreWebView2PrintSettings2) MustGetPageRanges() string {
	value, err := i.GetPageRanges()
	ReportErrorAuto(err)
	return value
}

// MustGetPagesPerSide 获取每张纸上的页面数量。出错时会触发全局错误回调。
func (i *ICoreWebView2PrintSettings2) MustGetPagesPerSide() int32 {
	value, err := i.GetPagesPerSide()
	ReportErrorAuto(err)
	return value
}

// MustGetCopies 获取打印的副本数量。出错时会触发全局错误回调。
func (i *ICoreWebView2PrintSettings2) MustGetCopies() int32 {
	value, err := i.GetCopies()
	ReportErrorAuto(err)
	return value
}

// MustGetCollation 获取打印排序模式。出错时会触发全局错误回调。
func (i *ICoreWebView2PrintSettings2) MustGetCollation() COREWEBVIEW2_PRINT_COLLATION {
	value, err := i.GetCollation()
	ReportErrorAuto(err)
	return value
}

// MustGetColorMode 获取打印颜色模式。出错时会触发全局错误回调。
func (i *ICoreWebView2PrintSettings2) MustGetColorMode() COREWEBVIEW2_PRINT_COLOR_MODE {
	value, err := i.GetColorMode()
	ReportErrorAuto(err)
	return value
}

// MustGetDuplex 获取双面打印模式。出错时会触发全局错误回调。
func (i *ICoreWebView2PrintSettings2) MustGetDuplex() COREWEBVIEW2_PRINT_DUPLEX {
	value, err := i.GetDuplex()
	ReportErrorAuto(err)
	return value
}

// MustGetMediaSize 获取打印纸张大小。出错时会触发全局错误回调。
func (i *ICoreWebView2PrintSettings2) MustGetMediaSize() COREWEBVIEW2_PRINT_MEDIA_SIZE {
	value, err := i.GetMediaSize()
	ReportErrorAuto(err)
	return value
}

// MustGetPrinterName 获取打印机名称。出错时会触发全局错误回调。
func (i *ICoreWebView2PrintSettings2) MustGetPrinterName() string {
	value, err := i.GetPrinterName()
	ReportErrorAuto(err)
	return value
}

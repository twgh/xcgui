package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2PrintSettings 提供用于配置打印设置的接口。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2printsettings
type ICoreWebView2PrintSettings struct {
	Vtbl *ICoreWebView2PrintSettingsVtbl
}

type ICoreWebView2PrintSettingsVtbl struct {
	IUnknownVtbl
	GetOrientation                ComProc
	PutOrientation                ComProc
	GetScaleFactor                ComProc
	PutScaleFactor                ComProc
	GetPageWidth                  ComProc
	PutPageWidth                  ComProc
	GetPageHeight                 ComProc
	PutPageHeight                 ComProc
	GetMarginTop                  ComProc
	PutMarginTop                  ComProc
	GetMarginBottom               ComProc
	PutMarginBottom               ComProc
	GetMarginLeft                 ComProc
	PutMarginLeft                 ComProc
	GetMarginRight                ComProc
	PutMarginRight                ComProc
	GetShouldPrintBackgrounds     ComProc
	PutShouldPrintBackgrounds     ComProc
	GetShouldPrintSelectionOnly   ComProc
	PutShouldPrintSelectionOnly   ComProc
	GetShouldPrintHeaderAndFooter ComProc
	PutShouldPrintHeaderAndFooter ComProc
	GetHeaderTitle                ComProc
	PutHeaderTitle                ComProc
	GetFooterUri                  ComProc
	PutFooterUri                  ComProc
	// 2
	GetPageRanges   ComProc
	PutPageRanges   ComProc
	GetPagesPerSide ComProc
	PutPagesPerSide ComProc
	GetCopies       ComProc
	PutCopies       ComProc
	GetCollation    ComProc
	PutCollation    ComProc
	GetColorMode    ComProc
	PutColorMode    ComProc
	GetDuplex       ComProc
	PutDuplex       ComProc
	GetMediaSize    ComProc
	PutMediaSize    ComProc
	GetPrinterName  ComProc
	PutPrinterName  ComProc
}

func (i *ICoreWebView2PrintSettings) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2PrintSettings) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2PrintSettings) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetOrientation 获取打印方向。
func (i *ICoreWebView2PrintSettings) GetOrientation() (COREWEBVIEW2_PRINT_ORIENTATION, error) {
	var value COREWEBVIEW2_PRINT_ORIENTATION
	r, _, _ := i.Vtbl.GetOrientation.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// SetOrientation 设置打印方向。
func (i *ICoreWebView2PrintSettings) SetOrientation(orientation COREWEBVIEW2_PRINT_ORIENTATION) error {
	r, _, _ := i.Vtbl.PutOrientation.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(orientation),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetScaleFactor 获取打印比例因子。
func (i *ICoreWebView2PrintSettings) GetScaleFactor() (float64, error) {
	var value float64
	r, _, _ := i.Vtbl.GetScaleFactor.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// SetScaleFactor 设置打印比例因子。
//   - 缩放因子是一个介于 0.1 和 2.0 之间的值。默认值为 1.0。
//   - 如果提供的值无效，则返回 wapi.E_INVALIDARG, 且当前值不会更改。
func (i *ICoreWebView2PrintSettings) SetScaleFactor(scaleFactor float64) error {
	r, _, _ := i.Vtbl.PutScaleFactor.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&scaleFactor)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetPageWidth 获取页面宽度（以英寸为单位）。
func (i *ICoreWebView2PrintSettings) GetPageWidth() (float64, error) {
	var value float64
	r, _, _ := i.Vtbl.GetPageWidth.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// SetPageWidth 设置页面宽度（以英寸为单位）。
//   - 默认宽度为 8.5 英寸。
func (i *ICoreWebView2PrintSettings) SetPageWidth(pageWidth float64) error {
	r, _, _ := i.Vtbl.PutPageWidth.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&pageWidth)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetPageHeight 获取页面高度（以英寸为单位）。
func (i *ICoreWebView2PrintSettings) GetPageHeight() (float64, error) {
	var value float64
	r, _, _ := i.Vtbl.GetPageHeight.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// SetPageHeight 设置页面高度（以英寸为单位）。
//   - 默认高度为11英寸。
//   - 如果页面高度小于或等于零，则返回 wapi.E_INVALIDARG，且当前值不会更改。
func (i *ICoreWebView2PrintSettings) SetPageHeight(pageHeight float64) error {
	r, _, _ := i.Vtbl.PutPageHeight.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&pageHeight)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetMarginTop 获取上边距（以英寸为单位）。
func (i *ICoreWebView2PrintSettings) GetMarginTop() (float64, error) {
	var value float64
	r, _, _ := i.Vtbl.GetMarginTop.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// SetMarginTop 设置上边距（以英寸为单位）。
//   - 默认值为1厘米，约0.4英寸。
//   - 边距不能小于零。如果提供的值无效，将返回 wapi.E_INVALIDARG，且当前值不会更改。
func (i *ICoreWebView2PrintSettings) SetMarginTop(marginTop float64) error {
	r, _, _ := i.Vtbl.PutMarginTop.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&marginTop)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetMarginBottom 获取下边距（以英寸为单位）。
func (i *ICoreWebView2PrintSettings) GetMarginBottom() (float64, error) {
	var value float64
	r, _, _ := i.Vtbl.GetMarginBottom.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// SetMarginBottom 设置下边距（以英寸为单位）。
//   - 默认值为1厘米，约0.4英寸。
//   - 边距不能小于零。如果提供的值无效，将返回 wapi.E_INVALIDARG，且当前值不会更改。
func (i *ICoreWebView2PrintSettings) SetMarginBottom(marginBottom float64) error {
	r, _, _ := i.Vtbl.PutMarginBottom.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&marginBottom)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetMarginLeft 获取左边距（以英寸为单位）。
func (i *ICoreWebView2PrintSettings) GetMarginLeft() (float64, error) {
	var value float64
	r, _, _ := i.Vtbl.GetMarginLeft.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// SetMarginLeft 设置左边距（以英寸为单位）。
//   - 默认值为1厘米，即约0.4英寸。
//   - 边距不能小于零。如果提供的值无效，将返回 wapi.E_INVALIDARG，且当前值不会更改。
func (i *ICoreWebView2PrintSettings) SetMarginLeft(marginLeft float64) error {
	r, _, _ := i.Vtbl.PutMarginLeft.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&marginLeft)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetMarginRight 获取右边距（以英寸为单位）。
func (i *ICoreWebView2PrintSettings) GetMarginRight() (float64, error) {
	var value float64
	r, _, _ := i.Vtbl.GetMarginRight.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// SetMarginRight 设置右边距（以英寸为单位）。
//   - 默认值为1厘米，即约0.4英寸。
//   - 边距不能小于零。如果提供的值无效，将返回 wapi.E_INVALIDARG，且当前值不会更改。
func (i *ICoreWebView2PrintSettings) SetMarginRight(marginRight float64) error {
	r, _, _ := i.Vtbl.PutMarginRight.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&marginRight)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetShouldPrintBackgrounds 获取是否打印背景颜色和图像。
func (i *ICoreWebView2PrintSettings) GetShouldPrintBackgrounds() (bool, error) {
	var value int32
	r, _, _ := i.Vtbl.GetShouldPrintBackgrounds.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value != 0, nil
}

// SetShouldPrintBackgrounds 设置是否打印背景颜色和图像。
//   - 默认值为 false。
func (i *ICoreWebView2PrintSettings) SetShouldPrintBackgrounds(shouldPrintBackgrounds bool) error {
	r, _, _ := i.Vtbl.PutShouldPrintBackgrounds.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(shouldPrintBackgrounds),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetShouldPrintSelectionOnly 获取是否只打印选中内容。
func (i *ICoreWebView2PrintSettings) GetShouldPrintSelectionOnly() (bool, error) {
	var value int32
	r, _, _ := i.Vtbl.GetShouldPrintSelectionOnly.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value != 0, nil
}

// SetShouldPrintSelectionOnly 设置是否只打印选中内容。
//   - 默认值为 false。
func (i *ICoreWebView2PrintSettings) SetShouldPrintSelectionOnly(shouldPrintSelectionOnly bool) error {
	r, _, _ := i.Vtbl.PutShouldPrintSelectionOnly.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(shouldPrintSelectionOnly),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetShouldPrintHeaderAndFooter 获取是否打印页眉和页脚。
func (i *ICoreWebView2PrintSettings) GetShouldPrintHeaderAndFooter() (bool, error) {
	var value int32
	r, _, _ := i.Vtbl.GetShouldPrintHeaderAndFooter.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value != 0, nil
}

// SetShouldPrintHeaderAndFooter 设置是否打印页眉和页脚。
//   - 默认值为 false。
func (i *ICoreWebView2PrintSettings) SetShouldPrintHeaderAndFooter(shouldPrintHeaderAndFooter bool) error {
	r, _, _ := i.Vtbl.PutShouldPrintHeaderAndFooter.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(shouldPrintHeaderAndFooter),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetHeaderTitle 获取页眉标题。
func (i *ICoreWebView2PrintSettings) GetHeaderTitle() (string, error) {
	var value *uint16
	r, _, _ := i.Vtbl.GetHeaderTitle.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	title := common.UTF16PtrToString(value)
	wapi.CoTaskMemFree(unsafe.Pointer(value))
	return title, nil
}

// SetHeaderTitle 设置页眉标题。
//   - 默认值为当前文档的标题。
func (i *ICoreWebView2PrintSettings) SetHeaderTitle(headerTitle string) error {
	headerTitlePtr, err := syscall.UTF16PtrFromString(headerTitle)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.PutHeaderTitle.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(headerTitlePtr)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetFooterUri 获取页脚URI。
func (i *ICoreWebView2PrintSettings) GetFooterUri() (string, error) {
	var value *uint16
	r, _, _ := i.Vtbl.GetFooterUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	uri := common.UTF16PtrToString(value)
	wapi.CoTaskMemFree(unsafe.Pointer(value))
	return uri, nil
}

// SetFooterUri 设置页脚URI。
//   - 如果提供的是空字符串，则页脚中不会显示统一资源标识符（URI）。
//   - 默认值为当前统一资源标识符。
func (i *ICoreWebView2PrintSettings) SetFooterUri(footerUri string) error {
	footerUriPtr, err := syscall.UTF16PtrFromString(footerUri)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.PutFooterUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(footerUriPtr)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetICoreWebView2PrintSettings2 获取 ICoreWebView2PrintSettings2。
func (i *ICoreWebView2PrintSettings) GetICoreWebView2PrintSettings2() (*ICoreWebView2PrintSettings2, error) {
	var result *ICoreWebView2PrintSettings2
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2PrintSettings2)),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetOrientation 获取打印方向。
// 出错时会触发全局错误回调。
func (i *ICoreWebView2PrintSettings) MustGetOrientation() COREWEBVIEW2_PRINT_ORIENTATION {
	value, err := i.GetOrientation()
	ReportErrorAtuo(err)
	return value
}

// MustGetScaleFactor 获取打印比例因子。
// 出错时会触发全局错误回调。
func (i *ICoreWebView2PrintSettings) MustGetScaleFactor() float64 {
	value, err := i.GetScaleFactor()
	ReportErrorAtuo(err)
	return value
}

// MustGetPageWidth 获取页面宽度（以英寸为单位）。
// 出错时会触发全局错误回调。
func (i *ICoreWebView2PrintSettings) MustGetPageWidth() float64 {
	value, err := i.GetPageWidth()
	ReportErrorAtuo(err)
	return value
}

// MustGetPageHeight 获取页面高度（以英寸为单位）。
// 出错时会触发全局错误回调。
func (i *ICoreWebView2PrintSettings) MustGetPageHeight() float64 {
	value, err := i.GetPageHeight()
	ReportErrorAtuo(err)
	return value
}

// MustGetMarginTop 获取上边距（以英寸为单位）。
// 出错时会触发全局错误回调。
func (i *ICoreWebView2PrintSettings) MustGetMarginTop() float64 {
	value, err := i.GetMarginTop()
	ReportErrorAtuo(err)
	return value
}

// MustGetMarginBottom 获取下边距（以英寸为单位）。
// 出错时会触发全局错误回调。
func (i *ICoreWebView2PrintSettings) MustGetMarginBottom() float64 {
	value, err := i.GetMarginBottom()
	ReportErrorAtuo(err)
	return value
}

// MustGetMarginLeft 获取左边距（以英寸为单位）。
// 出错时会触发全局错误回调。
func (i *ICoreWebView2PrintSettings) MustGetMarginLeft() float64 {
	value, err := i.GetMarginLeft()
	ReportErrorAtuo(err)
	return value
}

// MustGetMarginRight 获取右边距（以英寸为单位）。
// 出错时会触发全局错误回调。
func (i *ICoreWebView2PrintSettings) MustGetMarginRight() float64 {
	value, err := i.GetMarginRight()
	ReportErrorAtuo(err)
	return value
}

// MustGetShouldPrintBackgrounds 获取是否打印背景颜色和图像。
// 出错时会触发全局错误回调。
func (i *ICoreWebView2PrintSettings) MustGetShouldPrintBackgrounds() bool {
	value, err := i.GetShouldPrintBackgrounds()
	ReportErrorAtuo(err)
	return value
}

// MustGetShouldPrintSelectionOnly 获取是否只打印选中内容。
// 出错时会触发全局错误回调。
func (i *ICoreWebView2PrintSettings) MustGetShouldPrintSelectionOnly() bool {
	value, err := i.GetShouldPrintSelectionOnly()
	ReportErrorAtuo(err)
	return value
}

// MustGetShouldPrintHeaderAndFooter 获取是否打印页眉和页脚。
// 出错时会触发全局错误回调。
func (i *ICoreWebView2PrintSettings) MustGetShouldPrintHeaderAndFooter() bool {
	value, err := i.GetShouldPrintHeaderAndFooter()
	ReportErrorAtuo(err)
	return value
}

// MustGetHeaderTitle 获取页眉标题。
// 出错时会触发全局错误回调。
func (i *ICoreWebView2PrintSettings) MustGetHeaderTitle() string {
	value, err := i.GetHeaderTitle()
	ReportErrorAtuo(err)
	return value
}

// MustGetFooterUri 获取页脚URI。
// 出错时会触发全局错误回调。
func (i *ICoreWebView2PrintSettings) MustGetFooterUri() string {
	value, err := i.GetFooterUri()
	ReportErrorAtuo(err)
	return value
}

// MustGetICoreWebView2PrintSettings2 获取 ICoreWebView2PrintSettings2。出错时会触发全局错误回调。
func (i *ICoreWebView2PrintSettings) MustGetICoreWebView2PrintSettings2() *ICoreWebView2PrintSettings2 {
	result, err := i.GetICoreWebView2PrintSettings2()
	ReportErrorAtuo(err)
	return result
}

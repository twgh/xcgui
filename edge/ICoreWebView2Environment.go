package edge

import (
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"

	"syscall"
	"unsafe"
)

// ICoreWebView2Environment 表示 WebView2 环境.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2environment
type ICoreWebView2Environment struct {
	Vtbl *ICoreWebView2EnvironmentVtbl
}

type ICoreWebView2EnvironmentVtbl struct {
	IUnknownVtbl
	CreateCoreWebView2Controller     ComProc
	CreateWebResourceResponse        ComProc
	GetBrowserVersionString          ComProc
	AddNewBrowserVersionAvailable    ComProc
	RemoveNewBrowserVersionAvailable ComProc
	// 2
	CreateWebResourceRequest ComProc
	// 3
	CreateCoreWebView2CompositionController ComProc
	CreateCoreWebView2PointerInfo           ComProc
	// 4
	GetAutomationProviderForWindow ComProc
	// 5
	AddBrowserProcessExited    ComProc
	RemoveBrowserProcessExited ComProc
	// 6
	CreatePrintSettings ComProc
	// 7
	GetUserDataFolder ComProc
	// 8
	AddProcessInfosChanged    ComProc
	RemoveProcessInfosChanged ComProc
	GetProcessInfos           ComProc
	// 9
	CreateContextMenuItem ComProc
	// 10
	CreateCoreWebView2ControllerOptions                ComProc
	CreateCoreWebView2ControllerWithOptions            ComProc
	CreateCoreWebView2CompositionControllerWithOptions ComProc
	// 11
	GetFailureReportFolderPath ComProc
	// 12
	CreateSharedBuffer ComProc
	// 13
	GetProcessExtendedInfos ComProc
	// 14
	CreateWebFileSystemFileHandle      ComProc
	CreateWebFileSystemDirectoryHandle ComProc
	CreateObjectCollection             ComProc
	// 15
	CreateFindOptions ComProc
}

func (i *ICoreWebView2Environment) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Environment) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Environment) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// CreateCoreWebView2Controller 异步创建新的 WebView。
func (i *ICoreWebView2Environment) CreateCoreWebView2Controller(parentWindow uintptr, handler *ICoreWebView2CreateCoreWebView2ControllerCompletedHandler) error {
	r, _, _ := i.Vtbl.CreateCoreWebView2Controller.Call(
		uintptr(unsafe.Pointer(i)),
		parentWindow,
		uintptr(unsafe.Pointer(handler)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// CreateWebResourceResponse 创建新的web资源响应对象。
func (i *ICoreWebView2Environment) CreateWebResourceResponse(content *IStream, statusCode int, reasonPhrase string, headers string) (*ICoreWebView2WebResourceResponse, error) {
	_reason, err := syscall.UTF16PtrFromString(reasonPhrase)
	if err != nil {
		return nil, err
	}
	_headers, err := syscall.UTF16PtrFromString(headers)
	if err != nil {
		return nil, err
	}
	var response *ICoreWebView2WebResourceResponse
	r, _, _ := i.Vtbl.CreateWebResourceResponse.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(content)),
		uintptr(statusCode),
		uintptr(unsafe.Pointer(_reason)),
		uintptr(unsafe.Pointer(_headers)),
		uintptr(unsafe.Pointer(&response)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return response, nil
}

// GetBrowserVersionString 获取当前 ICoreWebView2Environment 的浏览器版本信息，如果不是 WebView2 运行时，则包括通道名称。
func (i *ICoreWebView2Environment) GetBrowserVersionString() (string, error) {
	var _version *uint16
	r, _, _ := i.Vtbl.GetBrowserVersionString.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_version)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	version := common.UTF16PtrToString(_version)
	wapi.CoTaskMemFree(unsafe.Pointer(_version))
	return version, nil
}

// AddNewBrowserVersionAvailable 添加新浏览器版本可用事件处理程序。
//   - 当有新版本的 WebView2 运行时可用时触发。
func (i *ICoreWebView2Environment) AddNewBrowserVersionAvailable(eventHandler *ICoreWebView2NewBrowserVersionAvailableEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddNewBrowserVersionAvailable.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveNewBrowserVersionAvailable 移除新浏览器版本可用事件处理程序。
func (i *ICoreWebView2Environment) RemoveNewBrowserVersionAvailable(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveNewBrowserVersionAvailable.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetICoreWebView2Environment2 获取 ICoreWebView2Environment2。
func (i *ICoreWebView2Environment) GetICoreWebView2Environment2() (*ICoreWebView2Environment2, error) {
	var result *ICoreWebView2Environment2
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2Environment2)),
		unsafe.Pointer(&result))
	return result, err
}

// GetICoreWebView2Environment3 获取 ICoreWebView2Environment3。
func (i *ICoreWebView2Environment) GetICoreWebView2Environment3() (*ICoreWebView2Environment3, error) {
	var result *ICoreWebView2Environment3
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2Environment3)),
		unsafe.Pointer(&result))
	return result, err
}

// GetICoreWebView2Environment4 获取 ICoreWebView2Environment4。
func (i *ICoreWebView2Environment) GetICoreWebView2Environment4() (*ICoreWebView2Environment4, error) {
	var result *ICoreWebView2Environment4
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2Environment4)),
		unsafe.Pointer(&result))
	return result, err
}

// GetICoreWebView2Environment5 获取 ICoreWebView2Environment5。
func (i *ICoreWebView2Environment) GetICoreWebView2Environment5() (*ICoreWebView2Environment5, error) {
	var result *ICoreWebView2Environment5
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2Environment5)),
		unsafe.Pointer(&result))
	return result, err
}

// GetICoreWebView2Environment6 获取 ICoreWebView2Environment6。
func (i *ICoreWebView2Environment) GetICoreWebView2Environment6() (*ICoreWebView2Environment6, error) {
	var result *ICoreWebView2Environment6
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2Environment6)),
		unsafe.Pointer(&result))
	return result, err
}

// GetICoreWebView2Environment7 获取 ICoreWebView2Environment7。
func (i *ICoreWebView2Environment) GetICoreWebView2Environment7() (*ICoreWebView2Environment7, error) {
	var result *ICoreWebView2Environment7
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2Environment7)),
		unsafe.Pointer(&result))
	return result, err
}

// GetICoreWebView2Environment8 获取 ICoreWebView2Environment8。
func (i *ICoreWebView2Environment) GetICoreWebView2Environment8() (*ICoreWebView2Environment8, error) {
	var result *ICoreWebView2Environment8
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2Environment8)),
		unsafe.Pointer(&result))
	return result, err
}

// GetICoreWebView2Environment9 获取 ICoreWebView2Environment9。
func (i *ICoreWebView2Environment) GetICoreWebView2Environment9() (*ICoreWebView2Environment9, error) {
	var result *ICoreWebView2Environment9
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2Environment9)),
		unsafe.Pointer(&result))
	return result, err
}

// GetICoreWebView2Environment10 获取 ICoreWebView2Environment10。
func (i *ICoreWebView2Environment) GetICoreWebView2Environment10() (*ICoreWebView2Environment10, error) {
	var result *ICoreWebView2Environment10
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2Environment10)),
		unsafe.Pointer(&result))
	return result, err
}

// GetICoreWebView2Environment11 获取 ICoreWebView2Environment11。
func (i *ICoreWebView2Environment) GetICoreWebView2Environment11() (*ICoreWebView2Environment11, error) {
	var result *ICoreWebView2Environment11
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2Environment11)),
		unsafe.Pointer(&result))
	return result, err
}

// GetICoreWebView2Environment12 获取 ICoreWebView2Environment12。
func (i *ICoreWebView2Environment) GetICoreWebView2Environment12() (*ICoreWebView2Environment12, error) {
	var result *ICoreWebView2Environment12
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2Environment12)),
		unsafe.Pointer(&result))
	return result, err
}

// GetICoreWebView2Environment13 获取 ICoreWebView2Environment13。
func (i *ICoreWebView2Environment) GetICoreWebView2Environment13() (*ICoreWebView2Environment13, error) {
	var result *ICoreWebView2Environment13
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2Environment13)),
		unsafe.Pointer(&result))
	return result, err
}

// GetICoreWebView2Environment14 获取 ICoreWebView2Environment14。
func (i *ICoreWebView2Environment) GetICoreWebView2Environment14() (*ICoreWebView2Environment14, error) {
	var result *ICoreWebView2Environment14
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2Environment14)),
		unsafe.Pointer(&result))
	return result, err
}

// GetICoreWebView2Environment15 获取 ICoreWebView2Environment15。
func (i *ICoreWebView2Environment) GetICoreWebView2Environment15() (*ICoreWebView2Environment15, error) {
	var result *ICoreWebView2Environment15
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2Environment15)),
		unsafe.Pointer(&result))
	return result, err
}

// MustCreateWebResourceResponse 创建新的web资源响应对象。出错时会触发全局错误回调。
func (i *ICoreWebView2Environment) MustCreateWebResourceResponse(content *IStream, statusCode int, reasonPhrase string, headers string) *ICoreWebView2WebResourceResponse {
	response, err := i.CreateWebResourceResponse(content, statusCode, reasonPhrase, headers)
	ReportErrorAuto(err)
	return response
}

// MustGetBrowserVersionString 获取当前 ICoreWebView2Environment 的浏览器版本信息，如果不是WebView2运行时，则包括通道名称。出错时会触发全局错误回调。
func (i *ICoreWebView2Environment) MustGetBrowserVersionString() string {
	version, err := i.GetBrowserVersionString()
	ReportErrorAuto(err)
	return version
}

// MustGetICoreWebView2Environment2 获取 ICoreWebView2Environment2。出错时会触发全局错误回调。
func (i *ICoreWebView2Environment) MustGetICoreWebView2Environment2() *ICoreWebView2Environment2 {
	result, err := i.GetICoreWebView2Environment2()
	ReportErrorAuto(err)
	return result
}

// MustGetICoreWebView2Environment3 获取 ICoreWebView2Environment3。出错时会触发全局错误回调。
func (i *ICoreWebView2Environment) MustGetICoreWebView2Environment3() *ICoreWebView2Environment3 {
	result, err := i.GetICoreWebView2Environment3()
	ReportErrorAuto(err)
	return result
}

// MustGetICoreWebView2Environment4 获取 ICoreWebView2Environment4。出错时会触发全局错误回调。
func (i *ICoreWebView2Environment) MustGetICoreWebView2Environment4() *ICoreWebView2Environment4 {
	result, err := i.GetICoreWebView2Environment4()
	ReportErrorAuto(err)
	return result
}

// MustGetICoreWebView2Environment5 获取 ICoreWebView2Environment5。出错时会触发全局错误回调。
func (i *ICoreWebView2Environment) MustGetICoreWebView2Environment5() *ICoreWebView2Environment5 {
	result, err := i.GetICoreWebView2Environment5()
	ReportErrorAuto(err)
	return result
}

// MustGetICoreWebView2Environment6 获取 ICoreWebView2Environment6。出错时会触发全局错误回调。
func (i *ICoreWebView2Environment) MustGetICoreWebView2Environment6() *ICoreWebView2Environment6 {
	result, err := i.GetICoreWebView2Environment6()
	ReportErrorAuto(err)
	return result
}

// MustGetICoreWebView2Environment7 获取 ICoreWebView2Environment7。出错时会触发全局错误回调。
func (i *ICoreWebView2Environment) MustGetICoreWebView2Environment7() *ICoreWebView2Environment7 {
	result, err := i.GetICoreWebView2Environment7()
	ReportErrorAuto(err)
	return result
}

// MustGetICoreWebView2Environment8 获取 ICoreWebView2Environment8。出错时会触发全局错误回调。
func (i *ICoreWebView2Environment) MustGetICoreWebView2Environment8() *ICoreWebView2Environment8 {
	result, err := i.GetICoreWebView2Environment8()
	ReportErrorAuto(err)
	return result
}

// MustGetICoreWebView2Environment9 获取 ICoreWebView2Environment9。出错时会触发全局错误回调。
func (i *ICoreWebView2Environment) MustGetICoreWebView2Environment9() *ICoreWebView2Environment9 {
	result, err := i.GetICoreWebView2Environment9()
	ReportErrorAuto(err)
	return result
}

// MustGetICoreWebView2Environment10 获取 ICoreWebView2Environment10。出错时会触发全局错误回调。
func (i *ICoreWebView2Environment) MustGetICoreWebView2Environment10() *ICoreWebView2Environment10 {
	result, err := i.GetICoreWebView2Environment10()
	ReportErrorAuto(err)
	return result
}

// MustGetICoreWebView2Environment11 获取 ICoreWebView2Environment11。出错时会触发全局错误回调。
func (i *ICoreWebView2Environment) MustGetICoreWebView2Environment11() *ICoreWebView2Environment11 {
	result, err := i.GetICoreWebView2Environment11()
	ReportErrorAuto(err)
	return result
}

// MustGetICoreWebView2Environment12 获取 ICoreWebView2Environment12。出错时会触发全局错误回调。
func (i *ICoreWebView2Environment) MustGetICoreWebView2Environment12() *ICoreWebView2Environment12 {
	result, err := i.GetICoreWebView2Environment12()
	ReportErrorAuto(err)
	return result
}

// MustGetICoreWebView2Environment13 获取 ICoreWebView2Environment13。出错时会触发全局错误回调。
func (i *ICoreWebView2Environment) MustGetICoreWebView2Environment13() *ICoreWebView2Environment13 {
	result, err := i.GetICoreWebView2Environment13()
	ReportErrorAuto(err)
	return result
}

// MustGetICoreWebView2Environment14 获取 ICoreWebView2Environment14。出错时会触发全局错误回调。
func (i *ICoreWebView2Environment) MustGetICoreWebView2Environment14() *ICoreWebView2Environment14 {
	result, err := i.GetICoreWebView2Environment14()
	ReportErrorAuto(err)
	return result
}

// MustGetICoreWebView2Environment15 获取 ICoreWebView2Environment15。出错时会触发全局错误回调。
func (i *ICoreWebView2Environment) MustGetICoreWebView2Environment15() *ICoreWebView2Environment15 {
	result, err := i.GetICoreWebView2Environment15()
	ReportErrorAuto(err)
	return result
}

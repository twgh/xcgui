package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2Profile 提供一组用于配置 Profile 对象的属性。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2profile
type ICoreWebView2Profile struct {
	Vtbl *ICoreWebView2ProfileVtbl
}

type ICoreWebView2ProfileVtbl struct {
	IUnknownVtbl
	GetProfileName               ComProc
	GetIsInPrivateModeEnabled    ComProc
	GetProfilePath               ComProc
	GetDefaultDownloadFolderPath ComProc
	PutDefaultDownloadFolderPath ComProc
	GetPreferredColorScheme      ComProc
	PutPreferredColorScheme      ComProc
	// 2
	ClearBrowsingData            ComProc
	ClearBrowsingDataInTimeRange ComProc
	ClearBrowsingDataAll         ComProc
	// 3
	GetPreferredTrackingPreventionLevel ComProc
	PutPreferredTrackingPreventionLevel ComProc
	// 4
	SetPermissionState              ComProc
	GetNonDefaultPermissionSettings ComProc
}

func (i *ICoreWebView2Profile) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Profile) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Profile) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetProfileName 获取配置文件的名称。
func (i *ICoreWebView2Profile) GetProfileName() (string, error) {
	var _value *uint16
	r, _, _ := i.Vtbl.GetProfileName.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	value := common.UTF16PtrToString(_value)
	wapi.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

// GetIsInPrivateModeEnabled 获取配置文件是否启用了隐私模式。
func (i *ICoreWebView2Profile) GetIsInPrivateModeEnabled() (bool, error) {
	var _value int32
	r, _, _ := i.Vtbl.GetIsInPrivateModeEnabled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return _value != 0, nil
}

// GetProfilePath 获取配置文件目录的完整路径。
func (i *ICoreWebView2Profile) GetProfilePath() (string, error) {
	var _value *uint16
	r, _, _ := i.Vtbl.GetProfilePath.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	value := common.UTF16PtrToString(_value)
	wapi.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

// GetDefaultDownloadFolderPath 获取默认下载文件夹路径。
func (i *ICoreWebView2Profile) GetDefaultDownloadFolderPath() (string, error) {
	var _value *uint16
	r, _, _ := i.Vtbl.GetDefaultDownloadFolderPath.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	value := common.UTF16PtrToString(_value)
	wapi.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

// SetDefaultDownloadFolderPath 设置默认下载文件夹路径。
func (i *ICoreWebView2Profile) SetDefaultDownloadFolderPath(value string) error {
	_value, err := syscall.UTF16PtrFromString(value)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.PutDefaultDownloadFolderPath.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetPreferredColorScheme 获取与此配置文件关联的 WebView2 的整体配色方案。
func (i *ICoreWebView2Profile) GetPreferredColorScheme() (COREWEBVIEW2_PREFERRED_COLOR_SCHEME, error) {
	var _value COREWEBVIEW2_PREFERRED_COLOR_SCHEME
	r, _, _ := i.Vtbl.GetPreferredColorScheme.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return _value, nil
}

// SetPreferredColorScheme 设置与此配置文件关联的 WebView2 的整体配色方案。
func (i *ICoreWebView2Profile) SetPreferredColorScheme(value COREWEBVIEW2_PREFERRED_COLOR_SCHEME) error {
	r, _, _ := i.Vtbl.PutPreferredColorScheme.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetICoreWebView2Profile2 获取 ICoreWebView2Profile2。
func (i *ICoreWebView2Profile) GetICoreWebView2Profile2() (*ICoreWebView2Profile2, error) {
	var result *ICoreWebView2Profile2
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2Profile2)),
		unsafe.Pointer(&result))
	return result, err
}

// GetICoreWebView2Profile3 获取 ICoreWebView2Profile3。
func (i *ICoreWebView2Profile) GetICoreWebView2Profile3() (*ICoreWebView2Profile3, error) {
	var result *ICoreWebView2Profile3
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2Profile3)),
		unsafe.Pointer(&result))
	return result, err
}

// GetICoreWebView2Profile4 获取 ICoreWebView2Profile4。
func (i *ICoreWebView2Profile) GetICoreWebView2Profile4() (*ICoreWebView2Profile4, error) {
	var result *ICoreWebView2Profile4
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2Profile4)),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetProfileName 获取配置文件的名称。出错时会触发全局错误回调。
func (i *ICoreWebView2Profile) MustGetProfileName() string {
	result, err := i.GetProfileName()
	ReportErrorAtuo(err)
	return result
}

// MustGetIsInPrivateModeEnabled 获取配置文件是否启用了隐私模式。出错时会触发全局错误回调。
func (i *ICoreWebView2Profile) MustGetIsInPrivateModeEnabled() bool {
	result, err := i.GetIsInPrivateModeEnabled()
	ReportErrorAtuo(err)
	return result
}

// MustGetProfilePath 获取配置文件目录的完整路径。出错时会触发全局错误回调。
func (i *ICoreWebView2Profile) MustGetProfilePath() string {
	result, err := i.GetProfilePath()
	ReportErrorAtuo(err)
	return result
}

// MustGetDefaultDownloadFolderPath 获取默认下载文件夹路径。出错时会触发全局错误回调。
func (i *ICoreWebView2Profile) MustGetDefaultDownloadFolderPath() string {
	result, err := i.GetDefaultDownloadFolderPath()
	ReportErrorAtuo(err)
	return result
}

// MustGetPreferredColorScheme 获取与此配置文件关联的 WebView2 的整体配色方案。出错时会触发全局错误回调。
func (i *ICoreWebView2Profile) MustGetPreferredColorScheme() COREWEBVIEW2_PREFERRED_COLOR_SCHEME {
	result, err := i.GetPreferredColorScheme()
	ReportErrorAtuo(err)
	return result
}

// MustGetICoreWebView2Profile2 获取 ICoreWebView2Profile2。出错时会触发全局错误回调。
func (i *ICoreWebView2Profile) MustGetICoreWebView2Profile2() *ICoreWebView2Profile2 {
	result, err := i.GetICoreWebView2Profile2()
	ReportErrorAtuo(err)
	return result
}

// MustGetICoreWebView2Profile3 获取 ICoreWebView2Profile3。出错时会触发全局错误回调。
func (i *ICoreWebView2Profile) MustGetICoreWebView2Profile3() *ICoreWebView2Profile3 {
	result, err := i.GetICoreWebView2Profile3()
	ReportErrorAtuo(err)
	return result
}

// MustGetICoreWebView2Profile4 获取 ICoreWebView2Profile4。出错时会触发全局错误回调。
func (i *ICoreWebView2Profile) MustGetICoreWebView2Profile4() *ICoreWebView2Profile4 {
	result, err := i.GetICoreWebView2Profile4()
	ReportErrorAtuo(err)
	return result
}

/*
// GetICoreWebView2Profile5 获取 ICoreWebView2Profile5。
func (i *ICoreWebView2Profile) GetICoreWebView2Profile5() (*ICoreWebView2Profile5, error) {
	var result *ICoreWebView2Profile5
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2Profile5)),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetICoreWebView2Profile5 获取 ICoreWebView2Profile5。出错时会触发全局错误回调。
func (i *ICoreWebView2Profile) MustGetICoreWebView2Profile5() *ICoreWebView2Profile5 {
	result, err := i.GetICoreWebView2Profile5()
	ReportErrorAtuo(err)
	return result
}

// GetICoreWebView2Profile6 获取 ICoreWebView2Profile6。
func (i *ICoreWebView2Profile) GetICoreWebView2Profile6() (*ICoreWebView2Profile6, error) {
	var result *ICoreWebView2Profile6
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2Profile6)),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetICoreWebView2Profile6 获取 ICoreWebView2Profile6。出错时会触发全局错误回调。
func (i *ICoreWebView2Profile) MustGetICoreWebView2Profile6() *ICoreWebView2Profile6 {
	result, err := i.GetICoreWebView2Profile6()
	ReportErrorAtuo(err)
	return result
}

// GetICoreWebView2Profile7 获取 ICoreWebView2Profile7。
func (i *ICoreWebView2Profile) GetICoreWebView2Profile7() (*ICoreWebView2Profile7, error) {
	var result *ICoreWebView2Profile7
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2Profile7)),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetICoreWebView2Profile7 获取 ICoreWebView2Profile7。出错时会触发全局错误回调。
func (i *ICoreWebView2Profile) MustGetICoreWebView2Profile7() *ICoreWebView2Profile7 {
	result, err := i.GetICoreWebView2Profile7()
	ReportErrorAtuo(err)
	return result
}

// GetICoreWebView2Profile8 获取 ICoreWebView2Profile8。
func (i *ICoreWebView2Profile) GetICoreWebView2Profile8() (*ICoreWebView2Profile8, error) {
	var result *ICoreWebView2Profile8
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2Profile8)),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetICoreWebView2Profile8 获取 ICoreWebView2Profile8。出错时会触发全局错误回调。
func (i *ICoreWebView2Profile) MustGetICoreWebView2Profile8() *ICoreWebView2Profile8 {
	result, err := i.GetICoreWebView2Profile8()
	ReportErrorAtuo(err)
	return result
}
*/

package edge

// ok

import (
	"errors"
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"

	"syscall"
	"unsafe"
)

// ICoreWebView2CustomSchemeRegistration 表示向 ICoreWebView2Environment 注册自定义方案。
//
// 110.0.1587.40
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2customschemeregistration
type ICoreWebView2CustomSchemeRegistration struct {
	Vtbl *ICoreWebView2CustomSchemeRegistrationVtbl
}

type ICoreWebView2CustomSchemeRegistrationVtbl struct {
	IUnknownVtbl
	GetSchemeName            ComProc
	GetTreatAsSecure         ComProc
	PutTreatAsSecure         ComProc
	GetAllowedOrigins        ComProc
	SetAllowedOrigins        ComProc
	GetHasAuthorityComponent ComProc
	PutHasAuthorityComponent ComProc
}

func (i *ICoreWebView2CustomSchemeRegistration) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2CustomSchemeRegistration) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2CustomSchemeRegistration) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetSchemeName 获取自定义方案名称。
func (i *ICoreWebView2CustomSchemeRegistration) GetSchemeName() (string, error) {
	var value *uint16
	r, _, err := i.Vtbl.GetSchemeName.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return "", err
	}
	if r != 0 {
		return "", syscall.Errno(r)
	}
	str := common.UTF16PtrToString(value)
	wapi.CoTaskMemFree(unsafe.Pointer(value))
	return str, nil
}

// MustGetSchemeName 获取自定义方案名称。出错时会触发全局错误回调。
func (i *ICoreWebView2CustomSchemeRegistration) MustGetSchemeName() string {
	value, err := i.GetSchemeName()
	ReportErrorAtuo(err)
	return value
}

// GetTreatAsSecure 获取采用此方案的网站是否会像 HTTPS 网站一样被视为安全上下文。
func (i *ICoreWebView2CustomSchemeRegistration) GetTreatAsSecure() (bool, error) {
	var value bool
	r, _, err := i.Vtbl.GetTreatAsSecure.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value, nil
}

// MustGetTreatAsSecure 获取采用此方案的网站是否会像 HTTPS 网站一样被视为安全上下文。出错时会触发全局错误回调。
func (i *ICoreWebView2CustomSchemeRegistration) MustGetTreatAsSecure() bool {
	value, err := i.GetTreatAsSecure()
	ReportErrorAtuo(err)
	return value
}

// PutTreatAsSecure 设置该方案是否将被视为安全上下文。
func (i *ICoreWebView2CustomSchemeRegistration) PutTreatAsSecure(value bool) error {
	r, _, err := i.Vtbl.PutTreatAsSecure.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(value),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetAllowedOrigins 获取允许使用自定义方案（如带有 Origin 标头的 XHR 请求和子资源请求）发出请求的来源列表。
func (i *ICoreWebView2CustomSchemeRegistration) GetAllowedOrigins() ([]string, error) {
	var count uint32
	var origins **uint16
	r, _, err := i.Vtbl.GetAllowedOrigins.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&count)),
		uintptr(unsafe.Pointer(&origins)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return nil, err
	}
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	defer wapi.CoTaskMemFree(unsafe.Pointer(origins))

	result := make([]string, count)
	slice := unsafe.Slice(origins, count)
	for i := range slice {
		result[i] = common.UTF16PtrToString(slice[i])
		wapi.CoTaskMemFree(unsafe.Pointer(slice[i]))
	}
	return result, nil
}

// MustGetAllowedOrigins 获取允许使用自定义方案（如带有 Origin 标头的 XHR 请求和子资源请求）发出请求的来源列表。出错时会触发全局错误回调。
func (i *ICoreWebView2CustomSchemeRegistration) MustGetAllowedOrigins() []string {
	origins, err := i.GetAllowedOrigins()
	ReportErrorAtuo(err)
	return origins
}

// SetAllowedOrigins 设置被允许使用该方案（协议）的源数组。
func (i *ICoreWebView2CustomSchemeRegistration) SetAllowedOrigins(origins []string) error {
	if len(origins) == 0 {
		return errors.New("origins is empty")
	}
	_origins := make([]*uint16, len(origins))
	for i := range origins {
		_origin, err := syscall.UTF16PtrFromString(origins[i])
		if err != nil {
			return err
		}
		_origins[i] = _origin
	}
	r, _, err := i.Vtbl.SetAllowedOrigins.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(len(origins)),
		uintptr(unsafe.Pointer(&_origins[0])),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetHasAuthorityComponent 获取此方案是否具有权限组件。
func (i *ICoreWebView2CustomSchemeRegistration) GetHasAuthorityComponent() (bool, error) {
	var value bool
	r, _, err := i.Vtbl.GetHasAuthorityComponent.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value, nil
}

// PutHasAuthorityComponent 设置具有此自定义方案的 URI 是否将包含一个授权组件（自定义方案的主机）。
func (i *ICoreWebView2CustomSchemeRegistration) PutHasAuthorityComponent(value bool) error {
	r, _, err := i.Vtbl.PutHasAuthorityComponent.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(value),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

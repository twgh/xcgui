package edge

import (
	"errors"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

// ICoreWebView2_2 是 ICoreWebView2 接口的延续。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_2
type ICoreWebView2_2 struct {
	Vtbl *ICoreWebView2_2Vtbl
}

type ICoreWebView2_2Vtbl struct {
	ICoreWebView2Vtbl
	AddWebResourceResponseReceived    ComProc
	RemoveWebResourceResponseReceived ComProc
	NavigateWithWebResourceRequest    ComProc
	AddDomContentLoaded               ComProc
	RemoveDomContentLoaded            ComProc
	GetCookieManager                  ComProc
	GetEnvironment                    ComProc
}

func (i *ICoreWebView2_2) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_2) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_2) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetEnvironment 获取用于创建此 ICoreWebView2 的 ICoreWebView2Environment.
func (i *ICoreWebView2_2) GetEnvironment() (*ICoreWebView2Environment, error) {
	var environment *ICoreWebView2Environment
	r, _, err := i.Vtbl.GetEnvironment.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&environment)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return nil, err
	}
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return environment, nil
}

// MustGetEnvironment 获取用于创建此 ICoreWebView2 的 ICoreWebView2Environment. 出错时会触发全局错误回调.
func (i *ICoreWebView2_2) MustGetEnvironment() *ICoreWebView2Environment {
	environment, err := i.GetEnvironment()
	ReportError2(err)
	return environment
}

// NavigateWithWebResourceRequest 使用构造的 ICoreWebView2WebResourceRequest 对象进行导航。
func (i *ICoreWebView2_2) NavigateWithWebResourceRequest(request *ICoreWebView2WebResourceRequest) error {
	r, _, err := i.Vtbl.NavigateWithWebResourceRequest.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(request)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetCookieManager 获取与此 ICoreWebView2 关联的 cookie 管理器对象。
func (i *ICoreWebView2_2) GetCookieManager() (*ICoreWebView2CookieManager, error) {
	var cookieManager *ICoreWebView2CookieManager
	r, _, err := i.Vtbl.GetCookieManager.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&cookieManager)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return nil, err
	}
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return cookieManager, nil
}

// MustGetCookieManager 获取与此 ICoreWebView2 关联的 cookie 管理器对象。出错时会触发全局错误回调。
func (i *ICoreWebView2_2) MustGetCookieManager() *ICoreWebView2CookieManager {
	cookieManager, err := i.GetCookieManager()
	ReportError2(err)
	return cookieManager
}

/*TODO
AddWebResourceResponseReceived
RemoveWebResourceResponseReceived

AddDomContentLoaded
RemoveDomContentLoaded

*/

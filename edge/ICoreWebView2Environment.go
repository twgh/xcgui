package edge

import (
	"errors"
	"github.com/twgh/xcgui/wapi"
	"golang.org/x/sys/windows"
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
}

func (i *ICoreWebView2Environment) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Environment) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Environment) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// CreateCoreWebView2Controller 异步创建新的 WebView。
func (e *ICoreWebView2Environment) CreateCoreWebView2Controller(parentWindow uintptr, handler *ICoreWebView2CreateCoreWebView2ControllerCompletedHandler) error {
	r, _, err := e.Vtbl.CreateCoreWebView2Controller.Call(
		uintptr(unsafe.Pointer(e)),
		parentWindow,
		uintptr(unsafe.Pointer(handler)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// CreateWebResourceResponse 创建新的web资源响应对象。
func (e *ICoreWebView2Environment) CreateWebResourceResponse(content []byte, statusCode int, reasonPhrase string, headers string) (*ICoreWebView2WebResourceResponse, error) {
	var err error
	var stream uintptr

	if len(content) > 0 {
		// Create stream for response
		stream, err = wapi.SHCreateMemStream(content)
		if err != nil {
			return nil, err
		}
	}

	_reason, err := windows.UTF16PtrFromString(reasonPhrase)
	if err != nil {
		return nil, err
	}
	_headers, err := windows.UTF16PtrFromString(headers)
	if err != nil {
		return nil, err
	}
	var response *ICoreWebView2WebResourceResponse
	r, _, err := e.Vtbl.CreateWebResourceResponse.Call(
		uintptr(unsafe.Pointer(e)),
		stream,
		uintptr(statusCode),
		uintptr(unsafe.Pointer(_reason)),
		uintptr(unsafe.Pointer(_headers)),
		uintptr(unsafe.Pointer(&response)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return nil, err
	}
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return response, nil
}

// GetBrowserVersionString 获取当前 ICoreWebView2Environment 的浏览器版本信息，如果不是 WebView2 运行时，则包括通道名称。
func (e *ICoreWebView2Environment) GetBrowserVersionString() (string, error) {
	var _version *uint16
	r, _, err := e.Vtbl.GetBrowserVersionString.Call(
		uintptr(unsafe.Pointer(e)),
		uintptr(unsafe.Pointer(&_version)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return "", err
	}
	if r != 0 {
		return "", syscall.Errno(r)
	}
	version := windows.UTF16PtrToString(_version)
	windows.CoTaskMemFree(unsafe.Pointer(_version))
	return version, nil
}

// MustGetBrowserVersionString 获取当前 ICoreWebView2Environment 的浏览器版本信息，如果不是WebView2运行时，则包括通道名称。出错时会触发全局错误回调。
func (e *ICoreWebView2Environment) MustGetBrowserVersionString() string {
	version, _ := e.GetBrowserVersionString()
	return version
}

/*todo
AddNewBrowserVersionAvailable
RemoveNewBrowserVersionAvailable
*/

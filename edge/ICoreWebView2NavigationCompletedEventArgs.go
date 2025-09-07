package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2NavigationCompletedEventArgs 是导航完成事件的参数。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2navigationcompletedeventargs
type ICoreWebView2NavigationCompletedEventArgs struct {
	Vtbl *ICoreWebView2NavigationCompletedEventArgsVtbl
}

type ICoreWebView2NavigationCompletedEventArgsVtbl struct {
	IUnknownVtbl
	GetIsSuccess      ComProc
	GetWebErrorStatus ComProc
	GetNavigationId   ComProc
	// 2
	GetHttpStatusCode ComProc
}

func (i *ICoreWebView2NavigationCompletedEventArgs) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2NavigationCompletedEventArgs) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2NavigationCompletedEventArgs) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetIsSuccess 导航是否成功。
func (i *ICoreWebView2NavigationCompletedEventArgs) GetIsSuccess() (bool, error) {
	var isSuccess bool
	r, _, _ := i.Vtbl.GetIsSuccess.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isSuccess)),
	)
	if r != 0 {
		return isSuccess, syscall.Errno(r)
	}
	return isSuccess, nil
}

// GetWebErrorStatus 获取导航失败时的错误代码。
func (i *ICoreWebView2NavigationCompletedEventArgs) GetWebErrorStatus() (COREWEBVIEW2_WEB_ERROR_STATUS, error) {
	var status COREWEBVIEW2_WEB_ERROR_STATUS
	r, _, _ := i.Vtbl.GetWebErrorStatus.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&status)),
	)
	if r != 0 {
		return status, syscall.Errno(r)
	}
	return status, nil
}

// GetNavigationId 获取导航 ID。
func (i *ICoreWebView2NavigationCompletedEventArgs) GetNavigationId() (uint64, error) {
	var navigationId uint64
	r, _, _ := i.Vtbl.GetNavigationId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&navigationId)),
	)
	if r != 0 {
		return navigationId, syscall.Errno(r)
	}
	return navigationId, nil
}

// GetICoreWebView2NavigationCompletedEventArgs2 获取 ICoreWebView2NavigationCompletedEventArgs2。
func (i *ICoreWebView2NavigationCompletedEventArgs) GetICoreWebView2NavigationCompletedEventArgs2() (*ICoreWebView2NavigationCompletedEventArgs2, error) {
	var result *ICoreWebView2NavigationCompletedEventArgs2
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2NavigationCompletedEventArgs2)),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetIsSuccess 导航是否成功。出错时会触发全局错误回调。
func (i *ICoreWebView2NavigationCompletedEventArgs) MustGetIsSuccess() bool {
	isSuccess, err := i.GetIsSuccess()
	ReportErrorAtuo(err)
	return isSuccess
}

// MustGetWebErrorStatus 获取导航错误状态。出错时会触发全局错误回调。
func (i *ICoreWebView2NavigationCompletedEventArgs) MustGetWebErrorStatus() COREWEBVIEW2_WEB_ERROR_STATUS {
	status, err := i.GetWebErrorStatus()
	ReportErrorAtuo(err)
	return status
}

// MustGetNavigationId 获取导航 ID。出错时会触发全局错误回调。
func (i *ICoreWebView2NavigationCompletedEventArgs) MustGetNavigationId() uint64 {
	id, err := i.GetNavigationId()
	ReportErrorAtuo(err)
	return id
}

// MustGetICoreWebView2NavigationCompletedEventArgs2 获取 ICoreWebView2NavigationCompletedEventArgs2。出错时会触发全局错误回调。
func (i *ICoreWebView2NavigationCompletedEventArgs) MustGetICoreWebView2NavigationCompletedEventArgs2() *ICoreWebView2NavigationCompletedEventArgs2 {
	result, err := i.GetICoreWebView2NavigationCompletedEventArgs2()
	ReportErrorAtuo(err)
	return result
}

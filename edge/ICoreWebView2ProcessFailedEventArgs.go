package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2ProcessFailedEventArgs 是进程失败事件的事件参数。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2processfailedeventargs
type ICoreWebView2ProcessFailedEventArgs struct {
	Vtbl *ICoreWebView2ProcessFailedEventArgsVtbl
}

type ICoreWebView2ProcessFailedEventArgsVtbl struct {
	IUnknownVtbl
	GetProcessFailedKind ComProc
	// 2
	GetReason                     ComProc
	GetExitCode                   ComProc
	GetProcessDescription         ComProc
	GetFrameInfosForFailedProcess ComProc
	// 3
	GetFailureSourceModulePath ComProc
}

func (i *ICoreWebView2ProcessFailedEventArgs) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ProcessFailedEventArgs) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ProcessFailedEventArgs) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetProcessFailedKind 获取进程失败类型。
func (i *ICoreWebView2ProcessFailedEventArgs) GetProcessFailedKind() (COREWEBVIEW2_PROCESS_FAILED_KIND, error) {
	var kind COREWEBVIEW2_PROCESS_FAILED_KIND
	r, _, _ := i.Vtbl.GetProcessFailedKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&kind)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return kind, nil
}

// GetICoreWebView2ProcessFailedEventArgs2 获取 ICoreWebView2ProcessFailedEventArgs2。
func (i *ICoreWebView2ProcessFailedEventArgs) GetICoreWebView2ProcessFailedEventArgs2() (*ICoreWebView2ProcessFailedEventArgs2, error) {
	var result *ICoreWebView2ProcessFailedEventArgs2
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2ProcessFailedEventArgs2)),
		unsafe.Pointer(&result))
	return result, err
}

// GetICoreWebView2ProcessFailedEventArgs3 获取 ICoreWebView2ProcessFailedEventArgs3。
func (i *ICoreWebView2ProcessFailedEventArgs) GetICoreWebView2ProcessFailedEventArgs3() (*ICoreWebView2ProcessFailedEventArgs3, error) {
	var result *ICoreWebView2ProcessFailedEventArgs3
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2ProcessFailedEventArgs3)),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetProcessFailedKind 获取进程失败类型。出错时会触发全局错误回调。
func (i *ICoreWebView2ProcessFailedEventArgs) MustGetProcessFailedKind() COREWEBVIEW2_PROCESS_FAILED_KIND {
	kind, err := i.GetProcessFailedKind()
	ReportErrorAtuo(err)
	return kind
}

// MustGetICoreWebView2ProcessFailedEventArgs2 获取 ICoreWebView2ProcessFailedEventArgs2。出错时会触发全局错误回调。
func (i *ICoreWebView2ProcessFailedEventArgs) MustGetICoreWebView2ProcessFailedEventArgs2() *ICoreWebView2ProcessFailedEventArgs2 {
	result, err := i.GetICoreWebView2ProcessFailedEventArgs2()
	ReportErrorAtuo(err)
	return result
}

// MustGetICoreWebView2ProcessFailedEventArgs3 获取 ICoreWebView2ProcessFailedEventArgs3。出错时会触发全局错误回调。
func (i *ICoreWebView2ProcessFailedEventArgs) MustGetICoreWebView2ProcessFailedEventArgs3() *ICoreWebView2ProcessFailedEventArgs3 {
	result, err := i.GetICoreWebView2ProcessFailedEventArgs3()
	ReportErrorAtuo(err)
	return result
}

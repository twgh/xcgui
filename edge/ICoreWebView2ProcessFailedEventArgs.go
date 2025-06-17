package edge

// ok

import (
	"errors"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

// ICoreWebView2ProcessFailedEventArgs 是 ProcessFailed 事件的事件参数。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2processfailedeventargs
type ICoreWebView2ProcessFailedEventArgs struct {
	Vtbl *ICoreWebView2ProcessFailedEventArgsVtbl
}

type ICoreWebView2ProcessFailedEventArgsVtbl struct {
	IUnknownVtbl
	GetProcessFailedKind ComProc
}

// GetProcessFailedKind 获取进程失败类型。
func (i *ICoreWebView2ProcessFailedEventArgs) GetProcessFailedKind() (COREWEBVIEW2_PROCESS_FAILED_KIND, error) {
	var kind COREWEBVIEW2_PROCESS_FAILED_KIND
	r, _, err := i.Vtbl.GetProcessFailedKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&kind)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return kind, nil
}

// MustGetProcessFailedKind 获取进程失败类型。出错时会触发全局错误回调。
func (i *ICoreWebView2ProcessFailedEventArgs) MustGetProcessFailedKind() COREWEBVIEW2_PROCESS_FAILED_KIND {
	kind, err := i.GetProcessFailedKind()
	ReportErrorAtuo(err)
	return kind
}

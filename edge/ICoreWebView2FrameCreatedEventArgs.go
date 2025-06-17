package edge

// ok

import (
	"errors"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

// ICoreWebView2FrameCreatedEventArgs 是 FrameCreated 事件的事件参数。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2framecreatedeventargs
type ICoreWebView2FrameCreatedEventArgs struct {
	Vtbl *ICoreWebView2FrameCreatedEventArgsVtbl
}

type ICoreWebView2FrameCreatedEventArgsVtbl struct {
	IUnknownVtbl
	GetFrame ComProc
}

func (i *ICoreWebView2FrameCreatedEventArgs) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2FrameCreatedEventArgs) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2FrameCreatedEventArgs) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetFrame 获取创建的框架。
func (i *ICoreWebView2FrameCreatedEventArgs) GetFrame() (*ICoreWebView2Frame, error) {
	var frame *ICoreWebView2Frame
	r, _, err := i.Vtbl.GetFrame.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&frame)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return nil, err
	}
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return frame, nil
}

// MustGetFrame 获取创建的框架。出错时会触发全局错误回调。
func (i *ICoreWebView2FrameCreatedEventArgs) MustGetFrame() *ICoreWebView2Frame {
	frame, err := i.GetFrame()
	ReportErrorAtuo(err)
	return frame
}

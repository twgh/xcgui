package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/xc"
)

// ICoreWebView2DragStartingEventArgs 是拖拽开始事件的事件参数。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2dragstartingeventargs
type ICoreWebView2DragStartingEventArgs struct {
	Vtbl *ICoreWebView2DragStartingEventArgsVtbl
}

type ICoreWebView2DragStartingEventArgsVtbl struct {
	IUnknownVtbl
	GetAllowedDropEffects ComProc
	GetData               ComProc
	GetHandled            ComProc
	PutHandled            ComProc
	GetPosition           ComProc
	GetDeferral           ComProc
}

func (i *ICoreWebView2DragStartingEventArgs) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2DragStartingEventArgs) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2DragStartingEventArgs) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetAllowedDropEffects 获取允许的拖拽效果。
//   - 返回值是 wapi.DROPEFFECT_* 常量的组合。
//   - https://learn.microsoft.com/en-us/windows/win32/com/dropeffect-constants
func (i *ICoreWebView2DragStartingEventArgs) GetAllowedDropEffects() (uint32, error) {
	var value uint32
	r, _, _ := i.Vtbl.GetAllowedDropEffects.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// GetData 获取拖拽操作中包含的数据对象。
//   - 返回的 IDataObject 接口包含拖拽的数据。
func (i *ICoreWebView2DragStartingEventArgs) GetData() (*IUnknown, error) {
	var value *IUnknown
	r, _, _ := i.Vtbl.GetData.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return value, nil
}

// GetHandled 获取是否已处理拖拽操作。
func (i *ICoreWebView2DragStartingEventArgs) GetHandled() (bool, error) {
	var value bool
	r, _, _ := i.Vtbl.GetHandled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value, nil
}

// SetHandled 设置是否已处理拖拽操作。
//   - 如果设置为 true，则不会启动 WebView2 默认的拖拽操作。
func (i *ICoreWebView2DragStartingEventArgs) SetHandled(value bool) error {
	r, _, _ := i.Vtbl.PutHandled.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetPosition 获取拖拽操作的起始位置（相对于 WebView 的坐标）。
func (i *ICoreWebView2DragStartingEventArgs) GetPosition() (xc.POINT, error) {
	var point xc.POINT
	r, _, _ := i.Vtbl.GetPosition.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&point)),
	)
	if r != 0 {
		return xc.POINT{}, syscall.Errno(r)
	}
	return point, nil
}

// GetDeferral 获取延迟对象。
//   - 使用此操作可在稍后时间完成该事件。
//   - 在延迟完成之前，后续尝试在 WebView2 中启动拖放操作将会失败，并且如果光标已作为拖放操作的一部分进行了更改，将无法恢复。
func (i *ICoreWebView2DragStartingEventArgs) GetDeferral() (*ICoreWebView2Deferral, error) {
	var deferral *ICoreWebView2Deferral
	r, _, _ := i.Vtbl.GetDeferral.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&deferral)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return deferral, nil
}

// MustGetAllowedDropEffects 获取允许的拖拽效果。出错时会触发全局错误回调。
func (i *ICoreWebView2DragStartingEventArgs) MustGetAllowedDropEffects() uint32 {
	value, err := i.GetAllowedDropEffects()
	ReportErrorAuto(err)
	return value
}

// MustGetData 获取拖拽操作中包含的数据对象。出错时会触发全局错误回调。
func (i *ICoreWebView2DragStartingEventArgs) MustGetData() *IUnknown {
	value, err := i.GetData()
	ReportErrorAuto(err)
	return value
}

// MustGetHandled 获取是否已处理拖拽操作。出错时会触发全局错误回调。
func (i *ICoreWebView2DragStartingEventArgs) MustGetHandled() bool {
	value, err := i.GetHandled()
	ReportErrorAuto(err)
	return value
}

// MustGetPosition 获取拖拽操作的起始位置。出错时会触发全局错误回调。
func (i *ICoreWebView2DragStartingEventArgs) MustGetPosition() xc.POINT {
	value, err := i.GetPosition()
	ReportErrorAuto(err)
	return value
}

// MustGetDeferral 获取延迟对象。出错时会触发全局错误回调。
func (i *ICoreWebView2DragStartingEventArgs) MustGetDeferral() *ICoreWebView2Deferral {
	value, err := i.GetDeferral()
	ReportErrorAuto(err)
	return value
}

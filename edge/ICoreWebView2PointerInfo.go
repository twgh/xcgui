package edge

// ok

import (
	"errors"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/xc"
	"syscall"
	"unsafe"
)

// ICoreWebView2PointerInfo 这主要表示一个组合的 win32 POINTER_INFO / POINTER_TOUCH_INFO / POINTER_PEN_INFO 对象。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo
type ICoreWebView2PointerInfo struct {
	Vtbl *ICoreWebView2PointerInfoVtbl
}

type ICoreWebView2PointerInfoVtbl struct {
	IUnknownVtbl
	GetPointerKind         ComProc
	PutPointerKind         ComProc
	GetPointerId           ComProc
	PutPointerId           ComProc
	GetFrameId             ComProc
	PutFrameId             ComProc
	GetPointerFlags        ComProc
	PutPointerFlags        ComProc
	GetPointerDeviceRect   ComProc
	PutPointerDeviceRect   ComProc
	GetDisplayRect         ComProc
	PutDisplayRect         ComProc
	GetPixelLocation       ComProc
	PutPixelLocation       ComProc
	GetHimetricLocation    ComProc
	PutHimetricLocation    ComProc
	GetPixelLocationRaw    ComProc
	PutPixelLocationRaw    ComProc
	GetHimetricLocationRaw ComProc
	PutHimetricLocationRaw ComProc
	GetTime                ComProc
	PutTime                ComProc
	GetHistoryCount        ComProc
	PutHistoryCount        ComProc
	GetInputData           ComProc
	PutInputData           ComProc
	GetKeyStates           ComProc
	PutKeyStates           ComProc
	GetPerformanceCount    ComProc
	PutPerformanceCount    ComProc
	GetButtonChangeKind    ComProc
	PutButtonChangeKind    ComProc
	GetPenFlags            ComProc
	PutPenFlags            ComProc
	GetPenMask             ComProc
	PutPenMask             ComProc
	GetPenPressure         ComProc
	PutPenPressure         ComProc
	GetPenRotation         ComProc
	PutPenRotation         ComProc
	GetPenTiltX            ComProc
	PutPenTiltX            ComProc
	GetPenTiltY            ComProc
	PutPenTiltY            ComProc
	GetTouchFlags          ComProc
	PutTouchFlags          ComProc
	GetTouchMask           ComProc
	PutTouchMask           ComProc
	GetTouchContact        ComProc
	PutTouchContact        ComProc
	GetTouchContactRaw     ComProc
	PutTouchContactRaw     ComProc
	GetTouchOrientation    ComProc
	PutTouchOrientation    ComProc
	GetTouchPressure       ComProc
	PutTouchPressure       ComProc
}

func (i *ICoreWebView2PointerInfo) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2PointerInfo) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2PointerInfo) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetPointerKind 获取指针类型.
func (i *ICoreWebView2PointerInfo) GetPointerKind() (wapi.PT_, error) {
	var kind wapi.PT_
	r, _, err := i.Vtbl.GetPointerKind.Call(
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

// MustGetPointerKind 获取指针类型。出错时会触发全局错误回调。
func (i *ICoreWebView2PointerInfo) MustGetPointerKind() wapi.PT_ {
	kind, err := i.GetPointerKind()
	ReportErrorAtuo(err)
	return kind
}

// SetPointerKind 设置指针类型.
func (i *ICoreWebView2PointerInfo) SetPointerKind(kind wapi.PT_) error {
	r, _, err := i.Vtbl.PutPointerKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(kind),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetPointerId 获取指针ID.
func (i *ICoreWebView2PointerInfo) GetPointerId() (uint32, error) {
	var id uint32
	r, _, err := i.Vtbl.GetPointerId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&id)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return id, nil
}

// MustGetPointerId 获取指针ID。出错时会触发全局错误回调。
func (i *ICoreWebView2PointerInfo) MustGetPointerId() uint32 {
	id, err := i.GetPointerId()
	ReportErrorAtuo(err)
	return id
}

// SetPointerId 设置指针ID.
func (i *ICoreWebView2PointerInfo) SetPointerId(id uint32) error {
	r, _, err := i.Vtbl.PutPointerId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(id),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetFrameId 获取帧ID.
func (i *ICoreWebView2PointerInfo) GetFrameId() (uint32, error) {
	var id uint32
	r, _, err := i.Vtbl.GetFrameId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&id)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return id, nil
}

// MustGetFrameId 获取帧ID。出错时会触发全局错误回调。
func (i *ICoreWebView2PointerInfo) MustGetFrameId() uint32 {
	id, err := i.GetFrameId()
	ReportErrorAtuo(err)
	return id
}

// SetFrameId 设置帧ID.
func (i *ICoreWebView2PointerInfo) SetFrameId(id uint32) error {
	r, _, err := i.Vtbl.PutFrameId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(id),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetPointerFlags 获取指针标志.
func (i *ICoreWebView2PointerInfo) GetPointerFlags() (wapi.POINTER_FLAG_, error) {
	var flags wapi.POINTER_FLAG_
	r, _, err := i.Vtbl.GetPointerFlags.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&flags)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return flags, nil
}

// MustGetPointerFlags 获取指针标志。出错时会触发全局错误回调。
func (i *ICoreWebView2PointerInfo) MustGetPointerFlags() wapi.POINTER_FLAG_ {
	flags, err := i.GetPointerFlags()
	ReportErrorAtuo(err)
	return flags
}

// SetPointerFlags 设置指针标志.
func (i *ICoreWebView2PointerInfo) SetPointerFlags(flags wapi.POINTER_FLAG_) error {
	r, _, err := i.Vtbl.PutPointerFlags.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(flags),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetPointerDeviceRect 获取指针设备矩形区域.
func (i *ICoreWebView2PointerInfo) GetPointerDeviceRect() (xc.RECT, error) {
	var rect xc.RECT
	r, _, err := i.Vtbl.GetPointerDeviceRect.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&rect)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return xc.RECT{}, err
	}
	if r != 0 {
		return xc.RECT{}, syscall.Errno(r)
	}
	return rect, nil
}

// MustGetPointerDeviceRect 获取指针设备矩形区域。出错时会触发全局错误回调。
func (i *ICoreWebView2PointerInfo) MustGetPointerDeviceRect() xc.RECT {
	rect, err := i.GetPointerDeviceRect()
	ReportErrorAtuo(err)
	return rect
}

// SetPointerDeviceRect 设置指针设备矩形区域.
func (i *ICoreWebView2PointerInfo) SetPointerDeviceRect(rect xc.RECT) error {
	r, _, err := i.Vtbl.PutPointerDeviceRect.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&rect)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetDisplayRect 获取显示矩形区域.
func (i *ICoreWebView2PointerInfo) GetDisplayRect() (xc.RECT, error) {
	var rect xc.RECT
	r, _, err := i.Vtbl.GetDisplayRect.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&rect)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return xc.RECT{}, err
	}
	if r != 0 {
		return xc.RECT{}, syscall.Errno(r)
	}
	return rect, nil
}

// MustGetDisplayRect 获取显示矩形区域。出错时会触发全局错误回调。
func (i *ICoreWebView2PointerInfo) MustGetDisplayRect() xc.RECT {
	rect, err := i.GetDisplayRect()
	ReportErrorAtuo(err)
	return rect
}

// SetDisplayRect 设置显示矩形区域.
func (i *ICoreWebView2PointerInfo) SetDisplayRect(rect xc.RECT) error {
	r, _, err := i.Vtbl.PutDisplayRect.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&rect)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetPixelLocation 获取像素位置.
func (i *ICoreWebView2PointerInfo) GetPixelLocation() (xc.POINT, error) {
	var point xc.POINT
	r, _, err := i.Vtbl.GetPixelLocation.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&point)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return xc.POINT{}, err
	}
	if r != 0 {
		return xc.POINT{}, syscall.Errno(r)
	}
	return point, nil
}

// MustGetPixelLocation 获取像素位置。出错时会触发全局错误回调。
func (i *ICoreWebView2PointerInfo) MustGetPixelLocation() xc.POINT {
	point, err := i.GetPixelLocation()
	ReportErrorAtuo(err)
	return point
}

// SetPixelLocation 设置像素位置.
func (i *ICoreWebView2PointerInfo) SetPixelLocation(point xc.POINT) error {
	r, _, err := i.Vtbl.PutPixelLocation.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&point)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetHimetricLocation 获取 HIMETRIC 坐标位置.
func (i *ICoreWebView2PointerInfo) GetHimetricLocation() (xc.POINT, error) {
	var point xc.POINT
	r, _, err := i.Vtbl.GetHimetricLocation.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&point)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return xc.POINT{}, err
	}
	if r != 0 {
		return xc.POINT{}, syscall.Errno(r)
	}
	return point, nil
}

// MustGetHimetricLocation 获取 HIMETRIC 坐标位置。出错时会触发全局错误回调。
func (i *ICoreWebView2PointerInfo) MustGetHimetricLocation() xc.POINT {
	point, err := i.GetHimetricLocation()
	ReportErrorAtuo(err)
	return point
}

// SetHimetricLocation 设置 HIMETRIC 坐标位置.
func (i *ICoreWebView2PointerInfo) SetHimetricLocation(point xc.POINT) error {
	r, _, err := i.Vtbl.PutHimetricLocation.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&point)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetPixelLocationRaw 获取原始像素坐标位置.
func (i *ICoreWebView2PointerInfo) GetPixelLocationRaw() (xc.POINT, error) {
	var point xc.POINT
	r, _, err := i.Vtbl.GetPixelLocationRaw.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&point)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return xc.POINT{}, err
	}
	if r != 0 {
		return xc.POINT{}, syscall.Errno(r)
	}
	return point, nil
}

// MustGetPixelLocationRaw 获取原始像素坐标位置。出错时会触发全局错误回调。
func (i *ICoreWebView2PointerInfo) MustGetPixelLocationRaw() xc.POINT {
	point, err := i.GetPixelLocationRaw()
	ReportErrorAtuo(err)
	return point
}

// SetPixelLocationRaw 设置原始像素坐标位置.
func (i *ICoreWebView2PointerInfo) SetPixelLocationRaw(point xc.POINT) error {
	r, _, err := i.Vtbl.PutPixelLocationRaw.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&point)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetHimetricLocationRaw 获取原始 HIMETRIC 坐标位置.
func (i *ICoreWebView2PointerInfo) GetHimetricLocationRaw() (xc.POINT, error) {
	var point xc.POINT
	r, _, err := i.Vtbl.GetHimetricLocationRaw.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&point)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return xc.POINT{}, err
	}
	if r != 0 {
		return xc.POINT{}, syscall.Errno(r)
	}
	return point, nil
}

// MustGetHimetricLocationRaw 获取原始 HIMETRIC 坐标位置。出错时会触发全局错误回调。
func (i *ICoreWebView2PointerInfo) MustGetHimetricLocationRaw() xc.POINT {
	point, err := i.GetHimetricLocationRaw()
	ReportErrorAtuo(err)
	return point
}

// SetHimetricLocationRaw 设置原始 HIMETRIC 坐标位置.
func (i *ICoreWebView2PointerInfo) SetHimetricLocationRaw(point xc.POINT) error {
	r, _, err := i.Vtbl.PutHimetricLocationRaw.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&point)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetTime 获取时间戳.
func (i *ICoreWebView2PointerInfo) GetTime() (uint32, error) {
	var time uint32
	r, _, err := i.Vtbl.GetTime.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&time)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return time, nil
}

// MustGetTime 获取时间戳。出错时会触发全局错误回调。
func (i *ICoreWebView2PointerInfo) MustGetTime() uint32 {
	time, err := i.GetTime()
	ReportErrorAtuo(err)
	return time
}

// SetTime 设置时间戳.
func (i *ICoreWebView2PointerInfo) SetTime(time uint32) error {
	r, _, err := i.Vtbl.PutTime.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(time),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetHistoryCount 获取历史记录数量.
func (i *ICoreWebView2PointerInfo) GetHistoryCount() (uint32, error) {
	var count uint32
	r, _, err := i.Vtbl.GetHistoryCount.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&count)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return count, nil
}

// MustGetHistoryCount 获取历史记录数量。出错时会触发全局错误回调。
func (i *ICoreWebView2PointerInfo) MustGetHistoryCount() uint32 {
	count, err := i.GetHistoryCount()
	ReportErrorAtuo(err)
	return count
}

// SetHistoryCount 设置历史记录数量.
func (i *ICoreWebView2PointerInfo) SetHistoryCount(count uint32) error {
	r, _, err := i.Vtbl.PutHistoryCount.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(count),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetInputData 获取输入数据.
func (i *ICoreWebView2PointerInfo) GetInputData() (int32, error) {
	var data int32
	r, _, err := i.Vtbl.GetInputData.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&data)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return data, nil
}

// MustGetInputData 获取输入数据。出错时会触发全局错误回调。
func (i *ICoreWebView2PointerInfo) MustGetInputData() int32 {
	data, err := i.GetInputData()
	ReportErrorAtuo(err)
	return data
}

// SetInputData 设置输入数据.
func (i *ICoreWebView2PointerInfo) SetInputData(data int32) error {
	r, _, err := i.Vtbl.PutInputData.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(data),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetKeyStates 获取按键状态.
func (i *ICoreWebView2PointerInfo) GetKeyStates() (uint32, error) {
	var states uint32
	r, _, err := i.Vtbl.GetKeyStates.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&states)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return states, nil
}

// MustGetKeyStates 获取按键状态。出错时会触发全局错误回调。
func (i *ICoreWebView2PointerInfo) MustGetKeyStates() uint32 {
	states, err := i.GetKeyStates()
	ReportErrorAtuo(err)
	return states
}

// SetKeyStates 设置按键状态.
func (i *ICoreWebView2PointerInfo) SetKeyStates(states uint32) error {
	r, _, err := i.Vtbl.PutKeyStates.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(states),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetPerformanceCount 获取性能计数器值.
func (i *ICoreWebView2PointerInfo) GetPerformanceCount() (uint64, error) {
	var count uint64
	r, _, err := i.Vtbl.GetPerformanceCount.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&count)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return count, nil
}

// MustGetPerformanceCount 获取性能计数器值。出错时会触发全局错误回调。
func (i *ICoreWebView2PointerInfo) MustGetPerformanceCount() uint64 {
	count, err := i.GetPerformanceCount()
	ReportErrorAtuo(err)
	return count
}

// SetPerformanceCount 设置性能计数器值.
func (i *ICoreWebView2PointerInfo) SetPerformanceCount(count uint64) error {
	r, _, err := i.Vtbl.PutPerformanceCount.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(count),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetButtonChangeKind 获取按钮改变类型.
func (i *ICoreWebView2PointerInfo) GetButtonChangeKind() (wapi.POINTER_CHANGE_, error) {
	var kind wapi.POINTER_CHANGE_
	r, _, err := i.Vtbl.GetButtonChangeKind.Call(
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

// MustGetButtonChangeKind 获取按钮改变类型。出错时会触发全局错误回调。
func (i *ICoreWebView2PointerInfo) MustGetButtonChangeKind() wapi.POINTER_CHANGE_ {
	kind, err := i.GetButtonChangeKind()
	ReportErrorAtuo(err)
	return kind
}

// SetButtonChangeKind 设置按钮改变类型.
func (i *ICoreWebView2PointerInfo) SetButtonChangeKind(kind wapi.POINTER_CHANGE_) error {
	r, _, err := i.Vtbl.PutButtonChangeKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(kind),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetPenFlags 获取笔标志.
func (i *ICoreWebView2PointerInfo) GetPenFlags() (wapi.PEN_FLAG_, error) {
	var flags wapi.PEN_FLAG_
	r, _, err := i.Vtbl.GetPenFlags.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&flags)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return flags, nil
}

// SetPenFlags 设置笔标志.
func (i *ICoreWebView2PointerInfo) SetPenFlags(flags wapi.PEN_FLAG_) error {
	r, _, err := i.Vtbl.PutPenFlags.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(flags),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetPenMask 获取笔掩码.
func (i *ICoreWebView2PointerInfo) GetPenMask() (wapi.PEN_MASK_, error) {
	var mask wapi.PEN_MASK_
	r, _, err := i.Vtbl.GetPenMask.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&mask)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return mask, nil
}

// MustGetPenMask 获取笔掩码。出错时会触发全局错误回调。
func (i *ICoreWebView2PointerInfo) MustGetPenMask() wapi.PEN_MASK_ {
	mask, err := i.GetPenMask()
	ReportErrorAtuo(err)
	return mask
}

// SetPenMask 设置笔掩码.
func (i *ICoreWebView2PointerInfo) SetPenMask(mask wapi.PEN_MASK_) error {
	r, _, err := i.Vtbl.PutPenMask.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(mask),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetPenPressure 获取笔压力值.
func (i *ICoreWebView2PointerInfo) GetPenPressure() (uint32, error) {
	var pressure uint32
	r, _, err := i.Vtbl.GetPenPressure.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&pressure)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return pressure, nil
}

// MustGetPenPressure 获取笔压力值。出错时会触发全局错误回调。
func (i *ICoreWebView2PointerInfo) MustGetPenPressure() uint32 {
	pressure, err := i.GetPenPressure()
	ReportErrorAtuo(err)
	return pressure
}

// SetPenPressure 设置笔压力值.
func (i *ICoreWebView2PointerInfo) SetPenPressure(pressure uint32) error {
	r, _, err := i.Vtbl.PutPenPressure.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(pressure),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetPenRotation 获取笔旋转角度.
func (i *ICoreWebView2PointerInfo) GetPenRotation() (uint32, error) {
	var rotation uint32
	r, _, err := i.Vtbl.GetPenRotation.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&rotation)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return rotation, nil
}

// MustGetPenRotation 获取笔旋转角度。出错时会触发全局错误回调。
func (i *ICoreWebView2PointerInfo) MustGetPenRotation() uint32 {
	rotation, err := i.GetPenRotation()
	ReportErrorAtuo(err)
	return rotation
}

// SetPenRotation 设置笔旋转角度.
func (i *ICoreWebView2PointerInfo) SetPenRotation(rotation uint32) error {
	r, _, err := i.Vtbl.PutPenRotation.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(rotation),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetPenTiltX 获取笔的X轴倾斜角度.
func (i *ICoreWebView2PointerInfo) GetPenTiltX() (int32, error) {
	var tiltX int32
	r, _, err := i.Vtbl.GetPenTiltX.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&tiltX)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return tiltX, nil
}

// MustGetPenTiltX 获取笔的X轴倾斜角度。出错时会触发全局错误回调。
func (i *ICoreWebView2PointerInfo) MustGetPenTiltX() int32 {
	tiltX, err := i.GetPenTiltX()
	ReportErrorAtuo(err)
	return tiltX
}

// SetPenTiltX 设置笔的X轴倾斜角度.
func (i *ICoreWebView2PointerInfo) SetPenTiltX(tiltX int32) error {
	r, _, err := i.Vtbl.PutPenTiltX.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(tiltX),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetPenTiltY 获取笔的Y轴倾斜角度.
func (i *ICoreWebView2PointerInfo) GetPenTiltY() (int32, error) {
	var tiltY int32
	r, _, err := i.Vtbl.GetPenTiltY.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&tiltY)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return tiltY, nil
}

// MustGetPenTiltY 获取笔的Y轴倾斜角度。出错时会触发全局错误回调。
func (i *ICoreWebView2PointerInfo) MustGetPenTiltY() int32 {
	tiltY, err := i.GetPenTiltY()
	ReportErrorAtuo(err)
	return tiltY
}

// SetPenTiltY 设置笔的Y轴倾斜角度.
func (i *ICoreWebView2PointerInfo) SetPenTiltY(tiltY int32) error {
	r, _, err := i.Vtbl.PutPenTiltY.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(tiltY),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetTouchFlags 获取触摸标志.
func (i *ICoreWebView2PointerInfo) GetTouchFlags() (uint32, error) {
	var flags uint32
	r, _, err := i.Vtbl.GetTouchFlags.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&flags)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return flags, nil
}

// MustGetTouchFlags 获取触摸标志。出错时会触发全局错误回调。
func (i *ICoreWebView2PointerInfo) MustGetTouchFlags() uint32 {
	flags, err := i.GetTouchFlags()
	ReportErrorAtuo(err)
	return flags
}

// SetTouchFlags 设置触摸标志.
func (i *ICoreWebView2PointerInfo) SetTouchFlags(flags uint32) error {
	r, _, err := i.Vtbl.PutTouchFlags.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(flags),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetTouchMask 获取触摸掩码.
func (i *ICoreWebView2PointerInfo) GetTouchMask() (wapi.TOUCH_MASK_, error) {
	var mask wapi.TOUCH_MASK_
	r, _, err := i.Vtbl.GetTouchMask.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&mask)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return mask, nil
}

// SetTouchMask 设置触摸掩码.
func (i *ICoreWebView2PointerInfo) SetTouchMask(mask wapi.TOUCH_MASK_) error {
	r, _, err := i.Vtbl.PutTouchMask.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(mask),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetTouchContact 获取触摸接触区域.
func (i *ICoreWebView2PointerInfo) GetTouchContact() (xc.RECT, error) {
	var rect xc.RECT
	r, _, err := i.Vtbl.GetTouchContact.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&rect)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return xc.RECT{}, err
	}
	if r != 0 {
		return xc.RECT{}, syscall.Errno(r)
	}
	return rect, nil
}

// MustGetTouchContact 获取触摸接触区域。出错时会触发全局错误回调。
func (i *ICoreWebView2PointerInfo) MustGetTouchContact() xc.RECT {
	rect, err := i.GetTouchContact()
	ReportErrorAtuo(err)
	return rect
}

// SetTouchContact 设置触摸接触区域.
func (i *ICoreWebView2PointerInfo) SetTouchContact(rect xc.RECT) error {
	r, _, err := i.Vtbl.PutTouchContact.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&rect)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetTouchContactRaw 获取原始触摸接触区域.
func (i *ICoreWebView2PointerInfo) GetTouchContactRaw() (xc.RECT, error) {
	var rect xc.RECT
	r, _, err := i.Vtbl.GetTouchContactRaw.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&rect)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return xc.RECT{}, err
	}
	if r != 0 {
		return xc.RECT{}, syscall.Errno(r)
	}
	return rect, nil
}

// MustGetTouchContactRaw 获取原始触摸接触区域。出错时会触发全局错误回调。
func (i *ICoreWebView2PointerInfo) MustGetTouchContactRaw() xc.RECT {
	rect, err := i.GetTouchContactRaw()
	ReportErrorAtuo(err)
	return rect
}

// SetTouchContactRaw 设置原始触摸接触区域.
func (i *ICoreWebView2PointerInfo) SetTouchContactRaw(rect xc.RECT) error {
	r, _, err := i.Vtbl.PutTouchContactRaw.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&rect)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetTouchOrientation 获取触摸方向.
func (i *ICoreWebView2PointerInfo) GetTouchOrientation() (uint32, error) {
	var orientation uint32
	r, _, err := i.Vtbl.GetTouchOrientation.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&orientation)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return orientation, nil
}

// MustGetTouchOrientation 获取触摸方向。出错时会触发全局错误回调。
func (i *ICoreWebView2PointerInfo) MustGetTouchOrientation() uint32 {
	orientation, err := i.GetTouchOrientation()
	ReportErrorAtuo(err)
	return orientation
}

// SetTouchOrientation 设置触摸方向.
func (i *ICoreWebView2PointerInfo) SetTouchOrientation(orientation uint32) error {
	r, _, err := i.Vtbl.PutTouchOrientation.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(orientation),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetTouchPressure 获取触摸压力值.
func (i *ICoreWebView2PointerInfo) GetTouchPressure() (uint32, error) {
	var pressure uint32
	r, _, err := i.Vtbl.GetTouchPressure.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&pressure)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return pressure, nil
}

// MustGetTouchPressure 获取触摸压力值。出错时会触发全局错误回调。
func (i *ICoreWebView2PointerInfo) MustGetTouchPressure() uint32 {
	pressure, err := i.GetTouchPressure()
	ReportErrorAtuo(err)
	return pressure
}

// SetTouchPressure 设置触摸压力值.
func (i *ICoreWebView2PointerInfo) SetTouchPressure(pressure uint32) error {
	r, _, err := i.Vtbl.PutTouchPressure.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(pressure),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

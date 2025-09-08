package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2FrameInfoCollectionIterator 提供对框架信息集合的迭代访问。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2frameinfocollectioniterator
type ICoreWebView2FrameInfoCollectionIterator struct {
	Vtbl *ICoreWebView2FrameInfoCollectionIteratorVtbl
}

type ICoreWebView2FrameInfoCollectionIteratorVtbl struct {
	IUnknownVtbl
	GetHasCurrent ComProc
	GetCurrent    ComProc
	MoveNext      ComProc
}

func (i *ICoreWebView2FrameInfoCollectionIterator) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2FrameInfoCollectionIterator) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2FrameInfoCollectionIterator) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetHasCurrent 获取迭代器是否有当前项。
//   - 如果迭代器正在迭代的集合为空，或者迭代器已超过集合的末尾，则此值为 FALSE。
func (i *ICoreWebView2FrameInfoCollectionIterator) GetHasCurrent() (bool, error) {
	var value bool
	r, _, _ := i.Vtbl.GetHasCurrent.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value, nil
}

// GetCurrent 获取当前的框架信息。
//   - 如果 GetHasCurrent 为 FALSE，则返回 wapi.ERROR_INVALID_INDEX.
func (i *ICoreWebView2FrameInfoCollectionIterator) GetCurrent() (*ICoreWebView2FrameInfo, error) {
	var value *ICoreWebView2FrameInfo
	r, _, _ := i.Vtbl.GetCurrent.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return value, nil
}

// MoveNext 将迭代器移动到下一项。
func (i *ICoreWebView2FrameInfoCollectionIterator) MoveNext() (bool, error) {
	var value bool
	r, _, _ := i.Vtbl.MoveNext.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return value, nil
}

// MustGetHasCurrent 获取迭代器是否有当前项。出错时会触发全局错误回调。
//   - 如果迭代器正在迭代的集合为空，或者迭代器已超过集合的末尾，则此值为 FALSE。
func (i *ICoreWebView2FrameInfoCollectionIterator) MustGetHasCurrent() bool {
	value, err := i.GetHasCurrent()
	ReportErrorAuto(err)
	return value
}

// MustGetCurrent 获取当前的框架信息。出错时会触发全局错误回调。
//   - 如果 GetHasCurrent 为 FALSE，则返回 wapi.ERROR_INVALID_INDEX.
func (i *ICoreWebView2FrameInfoCollectionIterator) MustGetCurrent() *ICoreWebView2FrameInfo {
	value, err := i.GetCurrent()
	ReportErrorAuto(err)
	return value
}

// MustMoveNext 将迭代器移动到下一项。出错时会触发全局错误回调。
func (i *ICoreWebView2FrameInfoCollectionIterator) MustMoveNext() bool {
	value, err := i.MoveNext()
	ReportErrorAuto(err)
	return value
}

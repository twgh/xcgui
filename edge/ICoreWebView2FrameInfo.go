package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"

	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2FrameInfo 提供对框架信息的访问。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2frameinfo
type ICoreWebView2FrameInfo struct {
	Vtbl *ICoreWebView2FrameInfoVtbl
}

type ICoreWebView2FrameInfoVtbl struct {
	IUnknownVtbl
	GetName   ComProc
	GetSource ComProc
	// 2
	GetParentFrameInfo ComProc
	GetFrameId         ComProc
	GetFrameKind       ComProc
}

func (i *ICoreWebView2FrameInfo) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2FrameInfo) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2FrameInfo) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetName 获取框架的名称。
//   - iframe 的 window.name 属性的值。
//   - 默认值等同于声明它的 iframe HTML 标签，如<iframe name="frame-name">...</iframe>。
//   - 当框架没有 name 属性且 window.name 没有赋值时，返回的字符串为空。
func (i *ICoreWebView2FrameInfo) GetName() (string, error) {
	var value *uint16
	r, _, _ := i.Vtbl.GetName.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	name := common.UTF16PtrToString(value)
	wapi.CoTaskMemFree(unsafe.Pointer(value))
	return name, nil
}

// GetSource 获取框架的源URL。
func (i *ICoreWebView2FrameInfo) GetSource() (string, error) {
	var value *uint16
	r, _, _ := i.Vtbl.GetSource.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	source := common.UTF16PtrToString(value)
	wapi.CoTaskMemFree(unsafe.Pointer(value))
	return source, nil
}

// MustGetName 获取框架的名称。出错时会触发全局错误回调。
//   - iframe 的 window.name 属性的值。
//   - 默认值等同于声明它的 iframe HTML 标签，如<iframe name="frame-name">...</iframe>。
//   - 当框架没有 name 属性且 window.name 没有赋值时，返回的字符串为空。
func (i *ICoreWebView2FrameInfo) MustGetName() string {
	value, err := i.GetName()
	ReportErrorAuto(err)
	return value
}

// MustGetSource 获取框架的源URL。出错时会触发全局错误回调。
func (i *ICoreWebView2FrameInfo) MustGetSource() string {
	value, err := i.GetSource()
	ReportErrorAuto(err)
	return value
}

// GetICoreWebView2FrameInfo2 获取 ICoreWebView2FrameInfo2。
func (i *ICoreWebView2FrameInfo) GetICoreWebView2FrameInfo2() (*ICoreWebView2FrameInfo2, error) {
	var result *ICoreWebView2FrameInfo2
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2FrameInfo2)),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetICoreWebView2FrameInfo2 获取 ICoreWebView2FrameInfo2。出错时会触发全局错误回调。
func (i *ICoreWebView2FrameInfo) MustGetICoreWebView2FrameInfo2() *ICoreWebView2FrameInfo2 {
	result, err := i.GetICoreWebView2FrameInfo2()
	ReportErrorAuto(err)
	return result
}

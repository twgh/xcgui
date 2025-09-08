package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2FindOptions 提供了 WebView2 查找选项的设置接口。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2findoptions
type ICoreWebView2FindOptions struct {
	Vtbl *ICoreWebView2FindOptionsVtbl
}

type ICoreWebView2FindOptionsVtbl struct {
	IUnknownVtbl
	GetFindTerm                  ComProc
	PutFindTerm                  ComProc
	GetIsCaseSensitive           ComProc
	PutIsCaseSensitive           ComProc
	GetShouldHighlightAllMatches ComProc
	PutShouldHighlightAllMatches ComProc
	GetShouldMatchWord           ComProc
	PutShouldMatchWord           ComProc
	GetSuppressDefaultFindDialog ComProc
	PutSuppressDefaultFindDialog ComProc
}

func (i *ICoreWebView2FindOptions) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2FindOptions) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2FindOptions) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetFindTerm 获取查找词。
func (i *ICoreWebView2FindOptions) GetFindTerm() (string, error) {
	var _value *uint16
	r, _, _ := i.Vtbl.GetFindTerm.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	value := common.UTF16PtrToString(_value)
	wapi.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

// SetFindTerm 设置查找词。
//   - 这将在你下次调用 Start 方法时生效。
func (i *ICoreWebView2FindOptions) SetFindTerm(value string) error {
	_value, err := syscall.UTF16PtrFromString(value)
	if err != nil {
		return err
	}
	r, _, _ := i.Vtbl.PutFindTerm.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetIsCaseSensitive 获取是否区分大小写。
func (i *ICoreWebView2FindOptions) GetIsCaseSensitive() (bool, error) {
	var _value bool
	r, _, _ := i.Vtbl.GetIsCaseSensitive.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return _value, nil
}

// SetIsCaseSensitive 设置是否区分大小写。
func (i *ICoreWebView2FindOptions) SetIsCaseSensitive(value bool) error {
	r, _, _ := i.Vtbl.PutIsCaseSensitive.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetShouldMatchWord 获取是否匹配整个单词。
func (i *ICoreWebView2FindOptions) GetShouldMatchWord() (bool, error) {
	var _value bool
	r, _, _ := i.Vtbl.GetShouldMatchWord.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return _value, nil
}

// SetShouldMatchWord 设置是否匹配整个单词。
func (i *ICoreWebView2FindOptions) SetShouldMatchWord(value bool) error {
	r, _, _ := i.Vtbl.PutShouldMatchWord.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetShouldHighlightAllMatches 获取是否高亮显示所有匹配项。
func (i *ICoreWebView2FindOptions) GetShouldHighlightAllMatches() (bool, error) {
	var _value bool
	r, _, _ := i.Vtbl.GetShouldHighlightAllMatches.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return _value, nil
}

// SetShouldHighlightAllMatches 设置是否高亮显示所有匹配项。
//   - 只有调用 Start、FindNext 或 FindPrevious 时，对该属性的更改才会生效。除非在服务器端再次调用 Start 函数，否则无法更新会话的偏好设置。因此，只有调用这些函数之一后，更改才会生效。
func (i *ICoreWebView2FindOptions) SetShouldHighlightAllMatches(value bool) error {
	r, _, _ := i.Vtbl.PutShouldHighlightAllMatches.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetSuppressDefaultFindDialog 获取是否抑制默认查找对话框。
func (i *ICoreWebView2FindOptions) GetSuppressDefaultFindDialog() (bool, error) {
	var _value bool
	r, _, _ := i.Vtbl.GetSuppressDefaultFindDialog.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return _value, nil
}

// SetSuppressDefaultFindDialog 设置是否抑制默认查找对话框。
//   - 只有调用 Start、FindNext 或 FindPrevious 时，对该属性的更改才会生效。除非在服务器端再次调用 Start 函数，否则无法更新会话的偏好设置。因此，只有调用这些函数之一后，更改才会生效。
func (i *ICoreWebView2FindOptions) SetSuppressDefaultFindDialog(value bool) error {
	r, _, _ := i.Vtbl.PutSuppressDefaultFindDialog.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetFindTerm 获取查找词。出错时会触发全局错误回调。
func (i *ICoreWebView2FindOptions) MustGetFindTerm() string {
	value, err := i.GetFindTerm()
	ReportErrorAuto(err)
	return value
}

// MustGetIsCaseSensitive 获取是否区分大小写。出错时会触发全局错误回调。
func (i *ICoreWebView2FindOptions) MustGetIsCaseSensitive() bool {
	value, err := i.GetIsCaseSensitive()
	ReportErrorAuto(err)
	return value
}

// MustGetShouldMatchWord 获取是否匹配整个单词。出错时会触发全局错误回调。
func (i *ICoreWebView2FindOptions) MustGetShouldMatchWord() bool {
	value, err := i.GetShouldMatchWord()
	ReportErrorAuto(err)
	return value
}

// MustGetShouldHighlightAllMatches 获取是否高亮显示所有匹配项。出错时会触发全局错误回调。
func (i *ICoreWebView2FindOptions) MustGetShouldHighlightAllMatches() bool {
	value, err := i.GetShouldHighlightAllMatches()
	ReportErrorAuto(err)
	return value
}

// MustGetSuppressDefaultFindDialog 获取是否抑制默认查找对话框。出错时会触发全局错误回调。
func (i *ICoreWebView2FindOptions) MustGetSuppressDefaultFindDialog() bool {
	value, err := i.GetSuppressDefaultFindDialog()
	ReportErrorAuto(err)
	return value
}

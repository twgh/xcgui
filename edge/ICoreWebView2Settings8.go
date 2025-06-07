package edge

import (
	"errors"
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

// ICoreWebView2Settings8 是 ICoreWebView2Settings7 接口的延续, 支持管理是否需要信誉检查。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2settings8
type ICoreWebView2Settings8 struct {
	Vtbl *ICoreWebView2Settings8Vtbl
}

type ICoreWebView2Settings8Vtbl struct {
	ICoreWebView2Settings7Vtbl
	GetIsReputationCheckingRequired ComProc
	PutIsReputationCheckingRequired ComProc
}

func (i *ICoreWebView2Settings8) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Settings8) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Settings8) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetIsReputationCheckingRequired 获取是否需要进行信誉检查。
func (i *ICoreWebView2Settings8) GetIsReputationCheckingRequired() (bool, error) {
	var required bool
	r, _, err := i.Vtbl.GetIsReputationCheckingRequired.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&required)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return required, nil
}

// MustGetIsReputationCheckingRequired 获取是否需要进行信誉检查。出错时会触发全局错误回调。
func (i *ICoreWebView2Settings8) MustGetIsReputationCheckingRequired() bool {
	required, err := i.GetIsReputationCheckingRequired()
	ReportErrorAtuo(err)
	return required
}

// PutIsReputationCheckingRequired 设置是否需要进行信誉检查。默认值为 true。
func (i *ICoreWebView2Settings8) PutIsReputationCheckingRequired(required bool) error {
	r, _, err := i.Vtbl.PutIsReputationCheckingRequired.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(required),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

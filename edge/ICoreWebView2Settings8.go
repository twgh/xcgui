package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
)

// ICoreWebView2Settings8 是 ICoreWebView2Settings7 接口的延续, 支持管理是否需要信誉检查。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2settings8
type ICoreWebView2Settings8 struct {
	ICoreWebView2Settings7
}

// GetIsReputationCheckingRequired 获取是否需要进行信誉检查。
func (i *ICoreWebView2Settings8) GetIsReputationCheckingRequired() (bool, error) {
	var required bool
	r, _, _ := i.Vtbl.GetIsReputationCheckingRequired.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&required)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return required, nil
}

// SetIsReputationCheckingRequired 设置是否需要进行信誉检查。默认值为 true。
func (i *ICoreWebView2Settings8) SetIsReputationCheckingRequired(required bool) error {
	r, _, _ := i.Vtbl.PutIsReputationCheckingRequired.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(required),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetIsReputationCheckingRequired 获取是否需要进行信誉检查。出错时会触发全局错误回调。
func (i *ICoreWebView2Settings8) MustGetIsReputationCheckingRequired() bool {
	required, err := i.GetIsReputationCheckingRequired()
	ReportErrorAuto(err)
	return required
}

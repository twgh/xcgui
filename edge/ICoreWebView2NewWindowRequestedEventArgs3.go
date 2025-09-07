package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2NewWindowRequestedEventArgs3 是 ICoreWebView2NewWindowRequestedEventArgs2 的延续。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2newwindowrequestedeventargs3
type ICoreWebView2NewWindowRequestedEventArgs3 struct {
	ICoreWebView2NewWindowRequestedEventArgs2
}

// GetOriginalSourceFrameInfo 获取新窗口请求发起所在框架的框架信息快照。
func (i *ICoreWebView2NewWindowRequestedEventArgs3) GetOriginalSourceFrameInfo() (*ICoreWebView2FrameInfo, error) {
	var value *ICoreWebView2FrameInfo
	r, _, _ := i.Vtbl.GetOriginalSourceFrameInfo.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return value, nil
}

// MustGetOriginalSourceFrameInfo 获取新窗口请求发起所在框架的框架信息快照。出错时会触发全局错误回调。
func (i *ICoreWebView2NewWindowRequestedEventArgs3) MustGetOriginalSourceFrameInfo() *ICoreWebView2FrameInfo {
	value, err := i.GetOriginalSourceFrameInfo()
	ReportErrorAtuo(err)
	return value
}

package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2FrameInfo2 是 ICoreWebView2FrameInfo 接口的延续，提供 ParentFrameInfo、FrameId 和 FrameKind 属性。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2frameinfo2
type ICoreWebView2FrameInfo2 struct {
	ICoreWebView2FrameInfo
}

// GetParentFrameInfo 获取此父框架的框架信息。
//   - ParentFrameInfo 仅在通过调用 ICoreWebView2ProcessExtendedInfo.AssociatedFrameInfos 获取时才会被填充。
//   - 通过 ICoreWebView2.ProcessFailed 获取的 ICoreWebView2FrameInfo 对象的 null ParentFrameInfo 始终为 null。
//   - 对于 WebView2 中没有父框架的主框架，此属性也为 null。
//   - 请注意，此 ParentFrameInfo 可能已过时，因为它是一个快照。
func (i *ICoreWebView2FrameInfo2) GetParentFrameInfo() (*ICoreWebView2FrameInfo, error) {
	var frameInfo *ICoreWebView2FrameInfo
	r, _, _ := i.Vtbl.GetParentFrameInfo.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&frameInfo)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return frameInfo, nil
}

// GetFrameId 获取与当前 FrameInfo 关联的框架的唯一标识符。
//   - 它与 ICoreWebView2 中的 FrameId 以及通过 ICoreWebView2Frame 的 FrameId 是同一种类型的 ID。
//   - FrameId 仅在调用 ICoreWebView2ProcessExtendedInfo.AssociatedFrameInfos 获取时才会被填充（非零）。
//   - 通过 ICoreWebView2.ProcessFailed 获取的 ICoreWebView2FrameInfo 对象将始终具有无效的框架 Id 0。
//   - 请注意，此 FrameId 可能已过时，因为它是一个快照。如果在异步调用 ICoreWebView2Environment.GetProcessExtendedInfos 启动后创建或销毁了 WebView2 或发生 FrameCreated/FrameDestroyed 事件，您可能需要再次调用异步方法以获取更新的 FrameInfo。
func (i *ICoreWebView2FrameInfo2) GetFrameId() (uint32, error) {
	var id uint32
	r, _, _ := i.Vtbl.GetFrameId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&id)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return id, nil
}

// GetFrameKind 获取框架的框架类型。
//   - FrameKind 仅在调用 ICoreWebView2ProcessExtendedInfo.AssociatedFrameInfos 获取时才会被填充。
//   - 通过 ICoreWebView2.ProcessFailed 获取的 ICoreWebView2FrameInfo 对象将始终具有默认值 COREWEBVIEW2_FRAME_KIND_UNKNOWN。
//   - 请注意，此 FrameKind 可能已过时，因为它是一个快照。
func (i *ICoreWebView2FrameInfo2) GetFrameKind() (COREWEBVIEW2_FRAME_KIND, error) {
	var kind COREWEBVIEW2_FRAME_KIND
	r, _, _ := i.Vtbl.GetFrameKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&kind)),
	)
	if r != 0 {
		return COREWEBVIEW2_FRAME_KIND_UNKNOWN, syscall.Errno(r)
	}
	return kind, nil
}

// MustGetParentFrameInfo 获取此父框架的框架信息。出错时会触发全局错误回调。
//   - ParentFrameInfo 仅在通过调用 ICoreWebView2ProcessExtendedInfo.AssociatedFrameInfos 获取时才会被填充。
//   - 通过 ICoreWebView2.ProcessFailed 获取的 ICoreWebView2FrameInfo 对象的 null ParentFrameInfo 始终为 null。
//   - 对于 WebView2 中没有父框架的主框架，此属性也为 null。
//   - 请注意，此 ParentFrameInfo 可能已过时，因为它是一个快照。
func (i *ICoreWebView2FrameInfo2) MustGetParentFrameInfo() *ICoreWebView2FrameInfo {
	value, err := i.GetParentFrameInfo()
	ReportErrorAtuo(err)
	return value
}

// MustGetFrameId 获取与当前 FrameInfo 关联的框架的唯一标识符。出错时会触发全局错误回调。
//   - 它与 ICoreWebView2 中的 FrameId 以及通过 ICoreWebView2Frame 的 FrameId 是同一种类型的 ID。
//   - FrameId 仅在调用 ICoreWebView2ProcessExtendedInfo.AssociatedFrameInfos 获取时才会被填充（非零）。
//   - 通过 ICoreWebView2.ProcessFailed 获取的 ICoreWebView2FrameInfo 对象将始终具有无效的框架 Id 0。
//   - 请注意，此 FrameId 可能已过时，因为它是一个快照。如果在异步调用 ICoreWebView2Environment.GetProcessExtendedInfos 启动后创建或销毁了 WebView2 或发生 FrameCreated/FrameDestroyed 事件，您可能需要再次调用异步方法以获取更新的 FrameInfo。
func (i *ICoreWebView2FrameInfo2) MustGetFrameId() uint32 {
	value, err := i.GetFrameId()
	ReportErrorAtuo(err)
	return value
}

// MustGetFrameKind 获取框架的框架类型。出错时会触发全局错误回调。
//   - FrameKind 仅在调用 ICoreWebView2ProcessExtendedInfo.AssociatedFrameInfos 获取时才会被填充。
//   - 通过 ICoreWebView2.ProcessFailed 获取的 ICoreWebView2FrameInfo 对象将始终具有默认值 COREWEBVIEW2_FRAME_KIND_UNKNOWN。
//   - 请注意，此 FrameKind 可能已过时，因为它是一个快照。
func (i *ICoreWebView2FrameInfo2) MustGetFrameKind() COREWEBVIEW2_FRAME_KIND {
	value, err := i.GetFrameKind()
	ReportErrorAtuo(err)
	return value
}

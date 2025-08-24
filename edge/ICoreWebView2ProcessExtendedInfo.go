package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2ProcessExtendedInfo 提供对进程扩展信息的访问。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2processextendedinfo
type ICoreWebView2ProcessExtendedInfo struct {
	Vtbl *ICoreWebView2ProcessExtendedInfoVtbl
}

type ICoreWebView2ProcessExtendedInfoVtbl struct {
	IUnknownVtbl
	GetProcessInfo          ComProc
	GetAssociatedFrameInfos ComProc
}

func (i *ICoreWebView2ProcessExtendedInfo) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ProcessExtendedInfo) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ProcessExtendedInfo) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetProcessInfo 获取当前进程的进程信息。
func (i *ICoreWebView2ProcessExtendedInfo) GetProcessInfo() (*ICoreWebView2ProcessInfo, error) {
	var processInfo *ICoreWebView2ProcessInfo
	r, _, _ := i.Vtbl.GetProcessInfo.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&processInfo)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return processInfo, nil
}

// GetAssociatedFrameInfos 获取在此渲染器进程中正在积极运行（显示或隐藏用户界面元素）的相关 FrameInfo 集合。
//   - 仅当此 CoreWebView2ProcessExtendedInfo 对应于呈现器进程时， AssociatedFrameInfos 才会填充。
//   - 非呈现器进程的 AssociatedFrameInfos 将始终为空。对于没有活动帧的呈现器进程， AssociatedFrameInfos 也可能为空。
func (i *ICoreWebView2ProcessExtendedInfo) GetAssociatedFrameInfos() (*ICoreWebView2FrameInfoCollection, error) {
	var frames *ICoreWebView2FrameInfoCollection
	r, _, _ := i.Vtbl.GetAssociatedFrameInfos.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&frames)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return frames, nil
}

// MustGetProcessInfo 获取当前进程的进程信息。出错时会触发全局错误回调。
func (i *ICoreWebView2ProcessExtendedInfo) MustGetProcessInfo() *ICoreWebView2ProcessInfo {
	processInfo, err := i.GetProcessInfo()
	ReportErrorAtuo(err)
	return processInfo
}

// MustGetAssociatedFrameInfos 获取在此渲染器进程中正在积极运行（显示或隐藏用户界面元素）的相关 FrameInfo 集合。出错时会触发全局错误回调。
//   - 仅当此 CoreWebView2ProcessExtendedInfo 对应于呈现器进程时， AssociatedFrameInfos 才会填充。
//   - 非呈现器进程的 AssociatedFrameInfos 将始终为空。对于没有活动帧的呈现器进程， AssociatedFrameInfos 也可能为空。
func (i *ICoreWebView2ProcessExtendedInfo) MustGetAssociatedFrameInfos() *ICoreWebView2FrameInfoCollection {
	frames, err := i.GetAssociatedFrameInfos()
	ReportErrorAtuo(err)
	return frames
}

package edge

// ok

import (
	"errors"
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

// ICoreWebView2_8 是 ICoreWebView2_7 接口的延续，支持媒体功能。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_8
type ICoreWebView2_8 struct {
	Vtbl *ICoreWebView2_8Vtbl
}

type ICoreWebView2_8Vtbl struct {
	ICoreWebView2_7Vtbl
	AddIsMutedChanged                   ComProc
	RemoveIsMutedChanged                ComProc
	GetIsMuted                          ComProc
	PutIsMuted                          ComProc
	AddIsDocumentPlayingAudioChanged    ComProc
	RemoveIsDocumentPlayingAudioChanged ComProc
	GetIsDocumentPlayingAudio           ComProc
}

func (i *ICoreWebView2_8) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_8) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_8) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetIsMuted 获取 WebView 是否静音。
func (i *ICoreWebView2_8) GetIsMuted() (bool, error) {
	var isMuted bool
	r, _, err := i.Vtbl.GetIsMuted.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isMuted)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return isMuted, nil
}

// SetIsMuted 设置 WebView 静音状态。
func (i *ICoreWebView2_8) SetIsMuted(isMuted bool) error {
	r, _, err := i.Vtbl.PutIsMuted.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(isMuted),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetIsDocumentPlayingAudio 获取当前文档是否正在播放音频。
//   - 如果音频正在播放，即使是静音状态，此属性也将为真。
func (i *ICoreWebView2_8) GetIsDocumentPlayingAudio() (bool, error) {
	var isPlaying bool
	r, _, err := i.Vtbl.GetIsDocumentPlayingAudio.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isPlaying)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return isPlaying, nil
}

// AddIsMutedChanged 添加静音状态改变事件处理程序.
func (i *ICoreWebView2_8) AddIsMutedChanged(eventHandler *ICoreWebView2IsMutedChangedEventHandler, token *EventRegistrationToken) error {
	r, _, err := i.Vtbl.AddIsMutedChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveIsMutedChanged 移除静音状态改变事件处理程序.
func (i *ICoreWebView2_8) RemoveIsMutedChanged(token EventRegistrationToken) error {
	r, _, err := i.Vtbl.RemoveIsMutedChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddIsDocumentPlayingAudioChanged 添加文档音频播放状态改变事件处理程序.
func (i *ICoreWebView2_8) AddIsDocumentPlayingAudioChanged(eventHandler *ICoreWebView2IsDocumentPlayingAudioChangedEventHandler, token *EventRegistrationToken) error {
	r, _, err := i.Vtbl.AddIsDocumentPlayingAudioChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveIsDocumentPlayingAudioChanged 移除文档音频播放状态改变事件处理程序.
func (i *ICoreWebView2_8) RemoveIsDocumentPlayingAudioChanged(token EventRegistrationToken) error {
	r, _, err := i.Vtbl.RemoveIsDocumentPlayingAudioChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetIsMuted 获取 WebView 是否静音。出错时会触发全局错误回调。
func (i *ICoreWebView2_8) MustGetIsMuted() bool {
	result, err := i.GetIsMuted()
	ReportErrorAtuo(err)
	return result
}

// MustGetIsDocumentPlayingAudio 获取当前文档是否正在播放音频。出错时会触发全局错误回调。
//   - 如果音频正在播放，即使是静音状态，此属性也将为真。
func (i *ICoreWebView2_8) MustGetIsDocumentPlayingAudio() bool {
	result, err := i.GetIsDocumentPlayingAudio()
	ReportErrorAtuo(err)
	return result
}

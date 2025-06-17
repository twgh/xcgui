package edge

// ok

import (
	"errors"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

// ICoreWebView2EnvironmentOptions7 提供用于创建 WebView2 环境以管理频道搜索和发布频道的其他选项。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions7
type ICoreWebView2EnvironmentOptions7 struct {
	Vtbl *ICoreWebView2EnvironmentOptions7Vtbl
}

type ICoreWebView2EnvironmentOptions7Vtbl struct {
	IUnknownVtbl
	GetChannelSearchKind ComProc
	PutChannelSearchKind ComProc
	GetReleaseChannels   ComProc
	PutReleaseChannels   ComProc
}

func (i *ICoreWebView2EnvironmentOptions7) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2EnvironmentOptions7) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2EnvironmentOptions7) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetChannelSearchKind 获取频道搜索类型.
func (i *ICoreWebView2EnvironmentOptions7) GetChannelSearchKind() (COREWEBVIEW2_CHANNEL_SEARCH_KIND, error) {
	var value COREWEBVIEW2_CHANNEL_SEARCH_KIND
	r, _, err := i.Vtbl.GetChannelSearchKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// SetChannelSearchKind 设置频道搜索类型.
//   - ChannelSearchKind 属性默认值为 COREWEBVIEW2_CHANNEL_SEARCH_KIND_MOST_STABLE.
//   - 环境创建会按稳定性从高到低的顺序在计算机上搜索发布通道，并使用找到的第一个通道。
//   - 默认搜索顺序为：WebView2 运行时 -> Beta -> Dev -> Canary。
func (i *ICoreWebView2EnvironmentOptions7) SetChannelSearchKind(value COREWEBVIEW2_CHANNEL_SEARCH_KIND) error {
	r, _, err := i.Vtbl.PutChannelSearchKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(value),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetReleaseChannels 获取发布频道.
func (i *ICoreWebView2EnvironmentOptions7) GetReleaseChannels() (COREWEBVIEW2_RELEASE_CHANNELS, error) {
	var value COREWEBVIEW2_RELEASE_CHANNELS
	r, _, err := i.Vtbl.GetReleaseChannels.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// SetReleaseChannels 设置发布频道.
//
// value: 一个或多个 COREWEBVIEW2_RELEASE_CHANNELS 的组合，指示环境创建应搜索哪些频道。
//   - 默认值是所有通道.
func (i *ICoreWebView2EnvironmentOptions7) SetReleaseChannels(value COREWEBVIEW2_RELEASE_CHANNELS) error {
	r, _, err := i.Vtbl.PutReleaseChannels.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(value),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetChannelSearchKind 获取频道搜索类型.
func (i *ICoreWebView2EnvironmentOptions7) MustGetChannelSearchKind() COREWEBVIEW2_CHANNEL_SEARCH_KIND {
	value, err := i.GetChannelSearchKind()
	ReportErrorAtuo(err)
	return value
}

// MustGetReleaseChannels 获取发布频道.
func (i *ICoreWebView2EnvironmentOptions7) MustGetReleaseChannels() COREWEBVIEW2_RELEASE_CHANNELS {
	value, err := i.GetReleaseChannels()
	ReportErrorAtuo(err)
	return value
}

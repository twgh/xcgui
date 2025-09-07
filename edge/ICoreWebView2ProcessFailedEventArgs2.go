package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2ProcessFailedEventArgs2 是 ICoreWebView2ProcessFailedEventArgs 的延续。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2processfailedeventargs2
type ICoreWebView2ProcessFailedEventArgs2 struct {
	ICoreWebView2ProcessFailedEventArgs
}

// GetReason 获取进程失败的原因。
func (i *ICoreWebView2ProcessFailedEventArgs2) GetReason() (COREWEBVIEW2_PROCESS_FAILED_REASON, error) {
	var reason COREWEBVIEW2_PROCESS_FAILED_REASON
	r, _, _ := i.Vtbl.GetReason.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&reason)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return reason, nil
}

// GetExitCode 获取失败进程的退出代码。
//   - 当 ProcessFailedKind 为 COREWEBVIEW2_PROCESS_FAILED_KIND_RENDER_PROCESS_UNRESPONSIVE 时，退出代码始终为 STILL_ACTIVE（259）。
func (i *ICoreWebView2ProcessFailedEventArgs2) GetExitCode() (int32, error) {
	var exitCode int32
	r, _, _ := i.Vtbl.GetExitCode.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&exitCode)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return exitCode, nil
}

// GetProcessDescription 获取失败进程的描述。
//   - 这是一个适用于日志记录或开发目的的技术英语术语，未针对最终用户进行本地化。
//   - 它适用于实用程序进程（例如，“音频服务”“视频捕获”）和插件进程（例如，“Flash”）。
//   - 如果 WebView2 运行时未为进程分配描述，则返回的 processDescription 为空。
func (i *ICoreWebView2ProcessFailedEventArgs2) GetProcessDescription() (string, error) {
	var _processDescription *uint16
	r, _, _ := i.Vtbl.GetProcessDescription.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_processDescription)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	processDescription := common.UTF16PtrToString(_processDescription)
	wapi.CoTaskMemFree(unsafe.Pointer(_processDescription))
	return processDescription, nil
}

// GetFrameInfosForFailedProcess 获取失败进程正在渲染的 ICoreWebView2 中的框架信息集合。
func (i *ICoreWebView2ProcessFailedEventArgs2) GetFrameInfosForFailedProcess() (*ICoreWebView2FrameInfoCollection, error) {
	var frames *ICoreWebView2FrameInfoCollection
	r, _, _ := i.Vtbl.GetFrameInfosForFailedProcess.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&frames)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return frames, nil
}

// MustGetReason 获取进程失败的原因。出错时会触发全局错误回调。
func (i *ICoreWebView2ProcessFailedEventArgs2) MustGetReason() COREWEBVIEW2_PROCESS_FAILED_REASON {
	reason, err := i.GetReason()
	ReportErrorAtuo(err)
	return reason
}

// MustGetExitCode 获取失败进程的退出代码。出错时会触发全局错误回调。
//   - 当 ProcessFailedKind 为 COREWEBVIEW2_PROCESS_FAILED_KIND_RENDER_PROCESS_UNRESPONSIVE 时，退出代码始终为 STILL_ACTIVE（259）。
func (i *ICoreWebView2ProcessFailedEventArgs2) MustGetExitCode() int32 {
	exitCode, err := i.GetExitCode()
	ReportErrorAtuo(err)
	return exitCode
}

// MustGetProcessDescription 获取失败进程的描述。出错时会触发全局错误回调。
//   - 这是一个适用于日志记录或开发目的的技术英语术语，未针对最终用户进行本地化。
//   - 它适用于实用程序进程（例如，“音频服务”“视频捕获”）和插件进程（例如，“Flash”）。
//   - 如果 WebView2 运行时未为进程分配描述，则返回的 processDescription 为空。
func (i *ICoreWebView2ProcessFailedEventArgs2) MustGetProcessDescription() string {
	processDescription, err := i.GetProcessDescription()
	ReportErrorAtuo(err)
	return processDescription
}

// MustGetFrameInfosForFailedProcess 获取失败进程正在渲染的 ICoreWebView2 中的框架信息集合。出错时会触发全局错误回调。
func (i *ICoreWebView2ProcessFailedEventArgs2) MustGetFrameInfosForFailedProcess() *ICoreWebView2FrameInfoCollection {
	frames, err := i.GetFrameInfosForFailedProcess()
	ReportErrorAtuo(err)
	return frames
}

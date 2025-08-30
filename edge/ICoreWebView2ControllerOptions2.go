package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2ControllerOptions2 是 ICoreWebView2ControllerOptions 的延续，提供了设置脚本区域设置的功能。
type ICoreWebView2ControllerOptions2 struct {
	ICoreWebView2ControllerOptions
}

// GetScriptLocale 获取脚本区域设置。
func (i *ICoreWebView2ControllerOptions2) GetScriptLocale() (string, error) {
	var value *uint16
	r, _, _ := i.Vtbl.GetScriptLocale.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	locale := common.UTF16PtrToString(value)
	wapi.CoTaskMemFree(unsafe.Pointer(value))
	return locale, nil
}

// SetScriptLocale 设置脚本区域设置。
//   - 它为所有 Intl JavaScript API 以及依赖于它的其他 JavaScript API 设置默认区域设置，即 Intl.DateTimeFormat()，这会影响字符串格式，例如时间/日期格式。示例：Intl.DateTimeFormat().format(new Date()) 预期的区域设置值采用 BCP 47 语言标记格式。
//   - 对此属性的更改适用于更改后创建的渲染器进程。任何现有的渲染器进程将继续使用之前的 ScriptLocale 值。要确保更改应用于所有渲染器进程，请关闭并重新启动 ICoreWebView2Environment 以及所有关联的 WebView2 对象。
//   - ScriptLocale 的默认值取决于 WebView2 语言和操作系统区域。如果 WebView2 语言和操作系统区域的语言部分匹配，则会使用操作系统区域。否则，将使用 WebView2 语言。
//   - 你可以将 ScriptLocale 设置为空字符串以获取默认的 ScriptLocale 值。
func (i *ICoreWebView2ControllerOptions2) SetScriptLocale(value string) error {
	r, _, _ := i.Vtbl.PutScriptLocale.Call(
		uintptr(unsafe.Pointer(i)),
		common.StrPtr(value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetScriptLocale 获取脚本区域设置。出错时会触发全局错误回调。
func (i *ICoreWebView2ControllerOptions2) MustGetScriptLocale() string {
	result, err := i.GetScriptLocale()
	ReportErrorAtuo(err)
	return result
}

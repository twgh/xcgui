package edge

import (
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

// ICoreWebView2AcceleratorKeyPressedEventArgs 是 AcceleratorKeyPressed 事件的事件参数。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2acceleratorkeypressedeventargs
type ICoreWebView2AcceleratorKeyPressedEventArgs struct {
	Vtbl *ICoreWebView2AcceleratorKeyPressedEventArgsVtbl
}

type ICoreWebView2AcceleratorKeyPressedEventArgsVtbl struct {
	IUnknownVtbl
	GetKeyEventKind      ComProc
	GetVirtualKey        ComProc
	GetKeyEventLParam    ComProc
	GetPhysicalKeyStatus ComProc
	GetHandled           ComProc
	PutHandled           ComProc
	// 2
	GetIsBrowserAcceleratorKeyEnabled ComProc
	PutIsBrowserAcceleratorKeyEnabled ComProc
}

func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetKeyEventKind 获取键事件类型。
func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) GetKeyEventKind() (COREWEBVIEW2_KEY_EVENT_KIND, error) {
	var keyEventKind COREWEBVIEW2_KEY_EVENT_KIND
	r, _, _ := i.Vtbl.GetKeyEventKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&keyEventKind)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return keyEventKind, nil
}

// GetVirtualKey 获取按下或释放的虚拟键代码。
func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) GetVirtualKey() (uint, error) {
	var virtualKey uint
	r, _, _ := i.Vtbl.GetVirtualKey.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&virtualKey)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return virtualKey, nil
}

// COREWEBVIEW2_PHYSICAL_KEY_STATUS 包含打包到发送到 Win32 键事件的 LPARAM 中的信息。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/corewebview2_physical_key_status
type COREWEBVIEW2_PHYSICAL_KEY_STATUS struct {
	RepeatCount   uint32 // 当前消息的重复计数。
	ScanCode      uint32 // 扫描代码。
	IsExtendedKey bool   // 是否为扩展键。
	IsMenuKeyDown bool   // 是否按住菜单键（上下文代码）。
	WasKeyDown    bool   // 表示按键已按下。
	IsKeyReleased bool   // 指示键已释放。
}

// GetPhysicalKeyStatus 获取有关按下或释放的键的物理状态的信息。
func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) GetPhysicalKeyStatus() (COREWEBVIEW2_PHYSICAL_KEY_STATUS, error) {
	var physicalKeyStatus COREWEBVIEW2_PHYSICAL_KEY_STATUS
	r, _, _ := i.Vtbl.GetPhysicalKeyStatus.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&physicalKeyStatus)),
	)
	if r != 0 {
		return physicalKeyStatus, syscall.Errno(r)
	}
	return physicalKeyStatus, nil
}

// SetHandled 设置事件是否已处理。
//   - 如果 Handled 属性设置为 TRUE，则这会阻止 WebView 对此快捷键执行默认操作。否则，WebView 将执行快捷键的默认操作。
func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) SetHandled(handled bool) error {
	r, _, _ := i.Vtbl.PutHandled.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(handled),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetKeyEventLParam 获取键事件的 LPARAM。
func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) GetKeyEventLParam() (int, error) {
	var lParam int
	r, _, _ := i.Vtbl.GetKeyEventLParam.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&lParam)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return lParam, nil
}

// GetHandled 获取事件是否已处理。
func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) GetHandled() (bool, error) {
	var handled int32
	r, _, _ := i.Vtbl.GetHandled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&handled)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return handled != 0, nil
}

// GetICoreWebView2AcceleratorKeyPressedEventArgs2 获取 ICoreWebView2AcceleratorKeyPressedEventArgs2。
func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) GetICoreWebView2AcceleratorKeyPressedEventArgs2() (*ICoreWebView2AcceleratorKeyPressedEventArgs2, error) {
	var result *ICoreWebView2AcceleratorKeyPressedEventArgs2
	err := i.QueryInterface(
		unsafe.Pointer(wapi.NewGUID(IID_ICoreWebView2AcceleratorKeyPressedEventArgs2)),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetKeyEventKind 获取键事件类型。出错时会触发全局错误回调.
func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) MustGetKeyEventKind() COREWEBVIEW2_KEY_EVENT_KIND {
	keyEventKind, err := i.GetKeyEventKind()
	ReportErrorAtuo(err)
	return keyEventKind
}

// MustGetVirtualKey 获取按下或释放的虚拟键代码。出错时会触发全局错误回调.
func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) MustGetVirtualKey() uint {
	virtualKey, err := i.GetVirtualKey()
	ReportErrorAtuo(err)
	return virtualKey
}

// MustGetPhysicalKeyStatus 获取有关按下或释放的键的物理状态的信息。出错时会触发全局错误回调.
func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) MustGetPhysicalKeyStatus() COREWEBVIEW2_PHYSICAL_KEY_STATUS {
	physicalKeyStatus, err := i.GetPhysicalKeyStatus()
	ReportErrorAtuo(err)
	return physicalKeyStatus
}

// MustGetKeyEventLParam 获取键事件的 LPARAM。出错时会触发全局错误回调。
func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) MustGetKeyEventLParam() int {
	lParam, err := i.GetKeyEventLParam()
	ReportErrorAtuo(err)
	return lParam
}

// MustGetHandled 获取事件是否已处理。出错时会触发全局错误回调。
func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) MustGetHandled() bool {
	handled, err := i.GetHandled()
	ReportErrorAtuo(err)
	return handled
}

// MustGetICoreWebView2AcceleratorKeyPressedEventArgs2 获取 ICoreWebView2AcceleratorKeyPressedEventArgs2。出错时会触发全局错误回调。
func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) MustGetICoreWebView2AcceleratorKeyPressedEventArgs2() *ICoreWebView2AcceleratorKeyPressedEventArgs2 {
	result, err := i.GetICoreWebView2AcceleratorKeyPressedEventArgs2()
	ReportErrorAtuo(err)
	return result
}

package edge

// ok

import (
	"errors"
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
}

func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetKeyEventKind 获取键事件类型。
func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) GetKeyEventKind() (COREWEBVIEW2_KEY_EVENT_KIND, error) {
	var keyEventKind COREWEBVIEW2_KEY_EVENT_KIND
	r, _, err := i.Vtbl.GetKeyEventKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&keyEventKind)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return keyEventKind, nil
}

// MustGetKeyEventKind 获取键事件类型。出错时会触发全局错误回调.
func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) MustGetKeyEventKind() COREWEBVIEW2_KEY_EVENT_KIND {
	keyEventKind, err := i.GetKeyEventKind()
	ReportErrorAtuo(err)
	return keyEventKind
}

// GetVirtualKey 获取按下或释放的虚拟键代码。
func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) GetVirtualKey() (uint, error) {
	var virtualKey uint
	r, _, err := i.Vtbl.GetVirtualKey.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&virtualKey)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return virtualKey, nil
}

// MustGetVirtualKey 获取按下或释放的虚拟键代码。出错时会触发全局错误回调.
func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) MustGetVirtualKey() uint {
	virtualKey, err := i.GetVirtualKey()
	ReportErrorAtuo(err)
	return virtualKey
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
	r, _, err := i.Vtbl.GetPhysicalKeyStatus.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&physicalKeyStatus)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return COREWEBVIEW2_PHYSICAL_KEY_STATUS{}, err
	}
	if r != 0 {
		return physicalKeyStatus, syscall.Errno(r)
	}
	return physicalKeyStatus, nil
}

// MustGetPhysicalKeyStatus 获取有关按下或释放的键的物理状态的信息。出错时会触发全局错误回调.
func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) MustGetPhysicalKeyStatus() COREWEBVIEW2_PHYSICAL_KEY_STATUS {
	physicalKeyStatus, err := i.GetPhysicalKeyStatus()
	ReportErrorAtuo(err)
	return physicalKeyStatus
}

// PutHandled 设置事件是否已处理。
//   - 如果 Handled 属性设置为 TRUE，则这会阻止 WebView 对此快捷键执行默认操作。否则，WebView 将执行快捷键的默认操作。
func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) PutHandled(handled bool) error {
	r, _, err := i.Vtbl.PutHandled.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(handled),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetKeyEventLParam 获取键事件的 LPARAM。
func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) GetKeyEventLParam() (int, error) {
	var lParam int
	r, _, err := i.Vtbl.GetKeyEventLParam.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&lParam)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return lParam, nil
}

// MustGetKeyEventLParam 获取键事件的 LPARAM。出错时会触发全局错误回调。
func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) MustGetKeyEventLParam() int {
	lParam, err := i.GetKeyEventLParam()
	ReportErrorAtuo(err)
	return lParam
}

// GetHandled 获取事件是否已处理。
func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) GetHandled() (bool, error) {
	var handled int32
	r, _, err := i.Vtbl.GetHandled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&handled)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return handled != 0, nil
}

// MustGetHandled 获取事件是否已处理。出错时会触发全局错误回调。
func (i *ICoreWebView2AcceleratorKeyPressedEventArgs) MustGetHandled() bool {
	handled, err := i.GetHandled()
	ReportErrorAtuo(err)
	return handled
}

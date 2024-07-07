package wutil

import (
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/xcc"
	"syscall"
	"unsafe"
)

// HookKeyboard 全局键盘钩子.
type HookKeyboard struct {
	Handle uintptr // 钩子句柄.
}

// NewHookKeyboard 创建一个全局键盘钩子, 可用于监听全局键盘消息, 失败返回nil.
//
// KeyboardProc: 回调函数, 可接收键盘消息. 在回调函数中放行键盘消息应该 return wutil.CallNextHookEx_Keyboard(nCode, wParam, lParam). 拦截键盘消息应该 return 1.
func NewHookKeyboard(KeyboardProc wapi.LowLevelKeyboardProc) *HookKeyboard {
	p := &HookKeyboard{}
	p.Handle = wapi.SetWindowsHookExW(wapi.WH_KEYBOARD_LL, syscall.NewCallback(KeyboardProc), wapi.GetModuleHandleW(""), 0)
	if p.Handle == 0 {
		return nil
	}
	return p
}

// Unhook 卸载全局键盘钩子.
func (h *HookKeyboard) Unhook() bool {
	return wapi.UnhookWindowsHookEx(h.Handle)
}

// CallNextHookEx_Keyboard 用于 NewHookKeyboard 的参数 KeyboardProc 回调函数的返回值.
//   - 在回调函数中放行键盘消息应该 return wutil.CallNextHookEx_Keyboard(nCode, wParam, lParam). 拦截键盘消息应该 return 1.
func CallNextHookEx_Keyboard(nCode int32, wParam xcc.WM_, lParam *wapi.KBDLLHOOKSTRUCT) uintptr {
	return wapi.CallNextHookEx(0, nCode, uintptr(wParam), uintptr(unsafe.Pointer(lParam)))
}

// HookMouse 全局鼠标钩子.
type HookMouse struct {
	Handle uintptr // 钩子句柄.
}

// 创建一个全局键盘钩子, 可用于监听全局键盘消息, 失败返回nil.
//
// MouseProc: 回调函数, 可接收鼠标消息. 在回调函数中放行鼠标消息应该 return wutil.CallNextHookEx_Mouse(nCode, wParam, lParam). 拦截鼠标消息应该 return 1.
func NewHookMouse(MouseProc wapi.LowLevelMouseProc) *HookMouse {
	p := &HookMouse{}
	p.Handle = wapi.SetWindowsHookExW(wapi.WH_MOUSE_LL, syscall.NewCallback(MouseProc), wapi.GetModuleHandleW(""), 0)
	if p.Handle == 0 {
		return nil
	}
	return p
}

// Unhook 卸载全局鼠标钩子.
func (h *HookMouse) Unhook() bool {
	return wapi.UnhookWindowsHookEx(h.Handle)
}

// CallNextHookEx_Mouse 用于 NewHookMouse 的参数 MouseProc 回调函数的返回值.
//   - 在回调函数中放行鼠标消息应该 return wutil.CallNextHookEx_Mouse(nCode, wParam, lParam). 拦截鼠标消息应该 return 1.
func CallNextHookEx_Mouse(nCode int32, wParam xcc.WM_, lParam *wapi.MSLLHOOKSTRUCT) uintptr {
	return wapi.CallNextHookEx(0, nCode, uintptr(wParam), uintptr(unsafe.Pointer(lParam)))
}

// GetHigh16Bits 获取 int32 的高16位.
func GetHigh16Bits(value int32) int16 {
	return int16(value >> 16)
}

// GetLow16Bits 获取 int32 的低16位.
func GetLow16Bits(value int32) int16 {
	return int16(value & 0xFFFF)
}

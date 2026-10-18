package xc

// 拦截 xcgui.dll 内部的错误弹窗.
//
// 当传入错误的句柄或句柄已失效时, xcgui.dll 内部会弹出标题为 "XCGUI-错误提示!" 的
// 消息框(窗口类 #32770), 点"是"会直接终止进程(控制台无任何输出), 点"否"才能看到崩溃栈.
//
// xcgui.dll 内部是通过 MessageBoxA/MessageBoxW 弹窗的(已从其导入表确认), 所以这里把
// xcgui.dll 导入表里的这两个槽位替换为 Go 回调, 于是:
//   - 不再弹出任何窗口;
//   - 可以在回调里拿到报错的 API 名和非法句柄(弹窗正文里就有);
//   - 可以在回调里打印 Go 侧的完整调用栈;
//   - 回调返回什么按钮 ID 由 SetErrorHookReturnButton 决定.
//
// 注意: 本钩子只能让错误"可见", 不能把进程救回来. dll 内部在弹窗之后仍会继续使用无效
// 句柄, 从而在 C 代码里触发访问违例(0xc0000005), 这种错误在 Go 里无法 recover.
// 想要不崩, 只能在调用前用 XC_IsHELE/XC_IsHWINDOW/XC_IsHXCGUI 等函数校验句柄.

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
	"sync"
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/wapi"
)

// ErrorHookHandler xcgui.dll 内部错误弹窗的处理函数.
//
//	api: 弹窗来源的 Win32 API, 为 "MessageBoxW" 或 "MessageBoxA".
//	text: 弹窗正文, 例如 "在API:XEle_Redraw()中, 输入句柄可能无效[0x1869F].\n\n\n是否终止程序?"
//	caption: 弹窗标题, 一般为 "XCGUI-错误提示!".
type ErrorHookHandler func(api, text, caption string)

var (
	errorHookMu        sync.Mutex
	errorHookInstalled bool
	errorHookHandler   ErrorHookHandler  = defaultErrorHookHandler
	errorHookReturn                      = uintptr(wapi.ID_NO) // 默认等价于点"否"
	msgBoxWCallback    uintptr                                 // MessageBoxW 的替换函数地址
	msgBoxACallback    uintptr                                 // MessageBoxA 的替换函数地址
	errorHookLogFile   = "xcgui_err.log"                       // 默认处理写日志的文件路径
)

// InstallErrorHook 安装 xcgui.dll 内部错误弹窗的拦截钩子.
//   - 必须在 xcgui.dll 加载之后调用, 即 app.New() 或 xc.XInitXCGUI() 之后.
//   - 可重复调用, 只有第一次会真正修改导入表.
//   - handler 可选, 不填则使用默认处理: 把错误内容和 Go 调用栈追加写入 xcgui_err.log, 同时
//     输出到 OutputDebugString(DebugView / IDE 输出窗口可见); 只有存在控制台时才写 stderr.
//   - 建议在调试时使用, 因为杀软可能对 IAT 改写有意见.
func InstallErrorHook(handler ...ErrorHookHandler) error {
	errorHookMu.Lock()
	defer errorHookMu.Unlock()

	if len(handler) > 0 && handler[0] != nil {
		errorHookHandler = handler[0]
	}
	if errorHookInstalled {
		return nil
	}
	if GetXcguiObject() == nil {
		return errors.New("xcgui.dll 尚未加载, 请先调用 app.New() 或 xc.XInitXCGUI()")
	}
	mod := GetXcguiObject().Handle()
	if mod == 0 {
		return errors.New("xcgui.dll 未加载成功, 本函数调用无效")
	}

	// 两个回调函数只需创建一次.
	msgBoxWCallback = syscall.NewCallback(onXCGUIErrorMessageBoxW)
	msgBoxACallback = syscall.NewCallback(onXCGUIErrorMessageBoxA)

	n := patchIAT(mod, "MessageBoxW", msgBoxWCallback)
	n += patchIAT(mod, "MessageBoxA", msgBoxACallback)
	if n == 0 {
		return errors.New("xcgui.dll 的导入表中未找到 MessageBoxA/MessageBoxW, 可能是 dll 版本已变更")
	}
	errorHookInstalled = true
	return nil
}

// SetErrorHookHandler 设置错误弹窗的处理函数, 可在 InstallErrorHook 之前或之后调用, 传 nil 恢复默认处理.
func SetErrorHookHandler(handler ErrorHookHandler) {
	errorHookMu.Lock()
	defer errorHookMu.Unlock()
	if handler == nil {
		errorHookHandler = defaultErrorHookHandler
		return
	}
	errorHookHandler = handler
}

// SetErrorHookReturnButton 设置拦截弹窗后回传给 dll 的按钮 ID, 默认 wapi.ID_NO.
//   - wapi.ID_NO(点"否"): dll 继续执行, 之后通常会在 dll 内部崩溃, 但能保留 Go 的崩溃栈, 便于排查.
//   - wapi.ID_YES(点"是"): dll 直接终止进程, 控制台不会有任何输出.
func SetErrorHookReturnButton(id wapi.ID_) {
	errorHookMu.Lock()
	defer errorHookMu.Unlock()
	errorHookReturn = uintptr(id)
}

// SetErrorHookLogFile 设置默认处理写日志的文件路径, 默认 "xcgui_err.log"(相对当前工作目录, 该目录不可写时写到临时目录).
//   - 传空字符串则不写文件, 但要注意: 没有控制台的程序完全看不到错误内容.
//   - 正式版建议设置成程序自己的日志目录, 例如 filepath.Join(日志目录, "xcgui_err.log").
func SetErrorHookLogFile(path string) {
	errorHookMu.Lock()
	defer errorHookMu.Unlock()
	errorHookLogFile = path
}

func defaultErrorHookHandler(api, text, caption string) {
	msg := fmt.Sprintf("[XCGUI错误] %s 标题=%q 内容=%q\n%s", api, caption, text, debug.Stack())

	// 正式版用 -H windowsgui 编译, 没有控制台也没有有效的 stderr 句柄, 输出到 stderr 会被直接丢弃,
	// 所以这里写文件为主, OutputDebugString 为辅(可在 DebugView 或 IDE 输出窗口看到).
	appendErrorLog(msg)
	wapi.OutputDebugString(msg)

	// 只有存在控制台(或启动方重定向了标准错误)时才写 stderr.
	if h := wapi.GetStdHandle(wapi.STD_ERROR_HANDLE); h != 0 && h != ^uintptr(0) {
		_, _ = os.Stderr.WriteString(msg)
	}
}

// appendErrorLog 把文本追加到 errorHookLogFile, 该文件不可写时退回到临时目录.
func appendErrorLog(s string) {
	path := getErrorHookLogFile()
	if path == "" {
		return
	}
	if appendFile(path, s) {
		return
	}
	_ = appendFile(filepath.Join(os.TempDir(), filepath.Base(path)), s)
}

func getErrorHookLogFile() string {
	errorHookMu.Lock()
	defer errorHookMu.Unlock()
	return errorHookLogFile
}

func appendFile(path, s string) bool {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return false
	}
	_, err = f.WriteString(s)
	_ = f.Close()
	return err == nil
}

func onXCGUIErrorMessageBoxW(hWnd, text, caption, uType uintptr) uintptr {
	handleErrorBox("MessageBoxW", utf16PtrToString(text), utf16PtrToString(caption))
	return getErrorHookReturn()
}

func onXCGUIErrorMessageBoxA(hWnd, text, caption, uType uintptr) uintptr {
	handleErrorBox("MessageBoxA", ansiPtrToString(text), ansiPtrToString(caption))
	return getErrorHookReturn()
}

func handleErrorBox(api, text, caption string) {
	errorHookMu.Lock()
	handler := errorHookHandler
	errorHookMu.Unlock()
	if handler != nil {
		handler(api, text, caption)
	}
}

func getErrorHookReturn() uintptr {
	errorHookMu.Lock()
	defer errorHookMu.Unlock()
	return errorHookReturn
}

// patchIAT 把已加载模块 modBase 的导入表中名为 funcName 的槽位替换为 newFunc, 返回替换的槽位数量.
func patchIAT(modBase uintptr, funcName string, newFunc uintptr) int {
	if readUint16(modBase) != 0x5A4D { // MZ
		return -1
	}
	nt := modBase + uintptr(readUint32(modBase+0x3C))
	if readUint32(nt) != 0x4550 { // PE\0\0
		return -1
	}
	opt := nt + 24
	ddOffset := uintptr(96) // PE32 的可选头中数据目录的偏移
	if readUint16(opt) == 0x20B {
		ddOffset = 112 // PE32+
	}
	importRVA := readUint32(opt + ddOffset + 8) // 数据目录[1] 为导入表
	if importRVA == 0 {
		return -1
	}
	count := 0
	for desc := modBase + uintptr(importRVA); readUint32(desc+12) != 0; desc += 20 {
		oft := readUint32(desc)     // OriginalFirstThunk, 指向函数名表
		ft := readUint32(desc + 16) // FirstThunk, 指向 IAT
		for i := uintptr(0); ; i++ {
			entry := modBase + uintptr(ft) + i*8
			if oft != 0 {
				entry = modBase + uintptr(oft) + i*8
			}
			nameEntry := readUintptr(entry)
			if nameEntry == 0 {
				break
			}
			if nameEntry&0x80000000 != 0 { // 序号导入, 没有函数名
				continue
			}
			// +2 跳过函数名前面的 Hint 字段.
			if readCString(modBase+uintptr(uint32(nameEntry))+2) != funcName {
				continue
			}
			slot := modBase + uintptr(ft) + i*8
			var oldProtect wapi.PAGE_
			if !wapi.VirtualProtect(slot, unsafe.Sizeof(newFunc), wapi.PAGE_READWRITE, &oldProtect) {
				continue
			}
			*(*uintptr)(unsafe.Pointer(slot)) = newFunc
			_ = wapi.VirtualProtect(slot, unsafe.Sizeof(newFunc), oldProtect, &oldProtect)
			count++
		}
	}
	return count
}

func readUint16(p uintptr) uint16 { return *(*uint16)(unsafe.Pointer(p)) }
func readUint32(p uintptr) uint32 { return *(*uint32)(unsafe.Pointer(p)) }
func readUintptr(p uintptr) uintptr {
	return *(*uintptr)(unsafe.Pointer(p))
}

// readCStringBytes 读取以 0 结尾的 ANSI 字节串, 最多读 4096 字节.
func readCStringBytes(p uintptr) []byte {
	if p == 0 {
		return nil
	}
	var b []byte
	for i := uintptr(0); i < 4096; i++ {
		c := *(*byte)(unsafe.Pointer(p + i))
		if c == 0 {
			break
		}
		b = append(b, c)
	}
	return b
}

func readCString(p uintptr) string {
	return string(readCStringBytes(p))
}

// utf16PtrToString 读取以 0 结尾的 UTF-16 字符串, 最多读 4096 个字符.
func utf16PtrToString(p uintptr) string {
	if p == 0 {
		return ""
	}
	var b []uint16
	for i := uintptr(0); i < 4096; i++ {
		c := *(*uint16)(unsafe.Pointer(p + i*2))
		if c == 0 {
			break
		}
		b = append(b, c)
	}
	return syscall.UTF16ToString(b)
}

// ansiPtrToString 按系统默认 ANSI 代码页读取以 0 结尾的字符串, 并转换为 UTF-8.
func ansiPtrToString(p uintptr) string {
	b := readCStringBytes(p)
	if len(b) == 0 {
		return ""
	}
	buf := make([]uint16, len(b)+1)
	n := wapi.MultiByteToWideChar(wapi.CP_ACP, 0,
		uintptr(unsafe.Pointer(&b[0])), int32(len(b)),
		uintptr(unsafe.Pointer(&buf[0])), int32(len(buf)))
	if n == 0 {
		return string(b)
	}
	return syscall.UTF16ToString(buf[:n])
}

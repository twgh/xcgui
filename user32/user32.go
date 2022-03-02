// user32.dll
package user32

import "syscall"

var (
	// Library.
	user32 *syscall.LazyDLL

	// Functions.
	setWindowPos     *syscall.LazyProc
	getDesktopWindow *syscall.LazyProc
)

func init() {
	// Library.
	user32 = syscall.NewLazyDLL("user32.dll")

	// Functions.
	setWindowPos = user32.NewProc("SetWindowPos")
	getDesktopWindow = user32.NewProc("GetDesktopWindow")
}

const (
	HWND_NOTOPMOST = -2 // 将窗口置于所有非顶层窗口之上（即在所有顶层窗口之后）。如果窗口已经是非顶层窗口则该标志不起作用。
	HWND_TOPMOST   = -1 // 将窗口置于所有非顶层窗口之上。即使窗口未被激活, 窗口也将保持顶级位置。
	HWND_TOP       = 0  // 将窗口置于Z序的顶部。
	HWND_BOTTOM    = 1  // 将窗口置于Z序的底部。如果参数hWnd标识了一个顶层窗口，则窗口失去顶级位置，并且被置在所有其他窗口的底部。
)

// 窗口大小和定位的标志
type SWP_ uint32

const (
	SWP_ASYNCWINDOWPOS SWP_ = 0x4000 // 如果调用线程和拥有窗口的线程连接到不同的输入队列，系统会将请求发布到拥有窗口的线程。这可以防止调用线程在其他线程处理请求时阻塞其执行。
	SWP_DEFERERASE     SWP_ = 0x2000 // 防止生成WM_SYNCPAINT消息。
	SWP_DRAWFRAME      SWP_ = 0x0020 // 在窗口周围绘制一个框架（在窗口的类描述中定义）。
	SWP_FRAMECHANGED   SWP_ = 0x0020 // 应用使用SetWindowLong函数 设置的新框架样式。向窗口发送WM_NCCALCSIZE消息，即使窗口大小没有改变。如果未指定此标志，则仅在更改窗口大小时发送 WM_NCCALCSIZE 。
	SWP_HIDEWINDOW     SWP_ = 0x0080 // 隐藏窗口。
	SWP_NOACTIVATE     SWP_ = 0x0010 // 不激活窗口。如果未设置此标志，则窗口被激活并移动到最顶层或非最顶层组的顶部（取决于hWndInsertAfter参数的设置）。
	SWP_NOCOPYBITS     SWP_ = 0x0100 // 丢弃客户区的全部内容。如果未指定此标志，则在调整窗口大小或重新定位后，将保存客户区的有效内容并将其复制回客户区。
	SWP_NOMOVE         SWP_ = 0x0002 // 保留当前位置（忽略X和Y参数）。
	SWP_NOOWNERZORDER  SWP_ = 0x0200 // 不改变所有者窗口在 Z 顺序中的位置。
	SWP_NOREDRAW       SWP_ = 0x0008 // 不重绘更改。如果设置了此标志，则不会发生任何类型的重新绘制。这适用于客户区、非客户区（包括标题栏和滚动条）以及由于窗口移动而未覆盖的父窗口的任何部分。设置此标志时，应用程序必须显式地使需要重绘的窗口和父窗口的任何部分无效或重绘。
	SWP_NOREPOSITION   SWP_ = 0x0200 // 与SWP_NOOWNERZORDER标志相同。
	SWP_NOSENDCHANGING SWP_ = 0x0400 // 阻止窗口接收WM_WINDOWPOSCHANGING消息。
	SWP_NOSIZE         SWP_ = 0x0001 // 保留当前大小（忽略cx和cy参数）。
	SWP_NOZORDER       SWP_ = 0x0004 // 保留当前 Z 顺序（忽略hWndInsertAfter参数）。
	SWP_SHOWWINDOW     SWP_ = 0x0040 // 显示窗口。
)

// SetWindowPos函数改变一个子窗口，弹出式窗口或顶层窗口的尺寸，位置和Z序。子窗口，弹出式窗口，及顶层窗口根据它们在屏幕上出现的顺序排序、顶层窗口设置的级别最高，并且被设置为Z序的第一个窗口.
//
// hWnd: 欲定位的窗口句柄.
//
// hWndInsertAfter: 在Z序中位于定位窗口之前的窗口句柄. 此参数必须是窗口句柄或以下值之一: user32.HWND_.
//
// x: 窗口新的x坐标。如hwnd是一个子窗口，则x用父窗口的客户区坐标表示.
//
// y: 窗口新的y坐标。如hwnd是一个子窗口，则y用父窗口的客户区坐标表示.
//
// cx: 指定新的窗口宽度.
//
// cy: 指定新的窗口高度.
//
// wFlags: 窗口大小和定位的标志. 该参数可以是以下值的组合: user32.SWP_.
func SetWindowPos(hWnd int, hWndInsertAfter int, x, y, cx, cy int, wFlags SWP_) bool {
	r, _, _ := setWindowPos.Call(uintptr(hWnd), uintptr(hWndInsertAfter), uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(wFlags))
	return int(r) != 0
}

// 获取桌面窗口的句柄.
func GetDesktopWindow() int {
	r, _, _ := getDesktopWindow.Call()
	return int(r)
}

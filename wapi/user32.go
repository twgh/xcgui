package wapi

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"

	"github.com/twgh/xcgui/common"
)

var (
	// User32 负责管理用户界面相关的基本功能，如窗口的创建、显示、消息处理和用户输入（鼠标、键盘）的响应。
	User32 = syscall.NewLazyDLL("user32.dll")

	// Functions.
	setWindowPos                               = User32.NewProc("SetWindowPos")
	getDesktopWindow                           = User32.NewProc("GetDesktopWindow")
	messageBoxW                                = User32.NewProc("MessageBoxW")
	isClipboardFormatAvailable                 = User32.NewProc("IsClipboardFormatAvailable")
	openClipboard                              = User32.NewProc("OpenClipboard")
	closeClipboard                             = User32.NewProc("CloseClipboard")
	emptyClipboard                             = User32.NewProc("EmptyClipboard")
	getClipboardData                           = User32.NewProc("GetClipboardData")
	setClipboardData                           = User32.NewProc("SetClipboardData")
	setForegroundWindow                        = User32.NewProc("SetForegroundWindow")
	findWindowExW                              = User32.NewProc("FindWindowExW")
	getWindowTextLengthW                       = User32.NewProc("GetWindowTextLengthW")
	getWindowTextW                             = User32.NewProc("GetWindowTextW")
	clientToScreen                             = User32.NewProc("ClientToScreen")
	getCursorPos                               = User32.NewProc("GetCursorPos")
	registerHotKey                             = User32.NewProc("RegisterHotKey")
	unregisterHotKey                           = User32.NewProc("UnregisterHotKey")
	getMessageW                                = User32.NewProc("GetMessageW")
	translateMessage                           = User32.NewProc("TranslateMessage")
	dispatchMessageW                           = User32.NewProc("DispatchMessageW")
	postQuitMessage                            = User32.NewProc("PostQuitMessage")
	sendMessageW                               = User32.NewProc("SendMessageW")
	postMessageW                               = User32.NewProc("PostMessageW")
	isWindow                                   = User32.NewProc("IsWindow")
	registerWindowMessageW                     = User32.NewProc("RegisterWindowMessageW")
	findWindowW                                = User32.NewProc("FindWindowW")
	loadImageW                                 = User32.NewProc("LoadImageW")
	createIconFromResource                     = User32.NewProc("CreateIconFromResource")
	destroyIcon                                = User32.NewProc("DestroyIcon")
	setWindowsHookExW                          = User32.NewProc("SetWindowsHookExW")
	unhookWindowsHookEx                        = User32.NewProc("UnhookWindowsHookEx")
	callNextHookEx                             = User32.NewProc("CallNextHookEx")
	procSetWindowLongPtrW                      = User32.NewProc("SetWindowLongPtrW")
	procGetWindowLongPtrW                      = User32.NewProc("GetWindowLongPtrW")
	procPhysicalToLogicalPointForPerMonitorDPI = User32.NewProc("PhysicalToLogicalPointForPerMonitorDPI")
	procMoveWindow                             = User32.NewProc("MoveWindow")
	procSetParent                              = User32.NewProc("SetParent")
	procShowWindow                             = User32.NewProc("ShowWindow")
	procUpdateWindow                           = User32.NewProc("UpdateWindow")
	procSetFocus                               = User32.NewProc("SetFocus")
	procGetClassName                           = User32.NewProc("GetClassNameW")
	procEnumWindows                            = User32.NewProc("EnumWindows")
	procIsWindowVisible                        = User32.NewProc("IsWindowVisible")
	procGetWindowThreadProcessId               = User32.NewProc("GetWindowThreadProcessId")
	procGetParent                              = User32.NewProc("GetParent")
	procPeekMessageW                           = User32.NewProc("PeekMessageW")
	procGetSystemMetrics                       = User32.NewProc("GetSystemMetrics")
	procRegisterClassExW                       = User32.NewProc("RegisterClassExW")
	procDefWindowProcW                         = User32.NewProc("DefWindowProcW")
	procCreateWindowExW                        = User32.NewProc("CreateWindowExW")
	procPostThreadMessageW                     = User32.NewProc("PostThreadMessageW")
	procIsDialogMessageW                       = User32.NewProc("IsDialogMessageW")
	procGetAncestor                            = User32.NewProc("GetAncestor")
	procDestroyWindow                          = User32.NewProc("DestroyWindow")
	procGetClientRect                          = User32.NewProc("GetClientRect")
	procSetWindowTextW                         = User32.NewProc("SetWindowTextW")
	procAdjustWindowRect                       = User32.NewProc("AdjustWindowRect")
	procGetAsyncKeyState                       = User32.NewProc("GetAsyncKeyState")
	procSetCursor                              = User32.NewProc("SetCursor")
	procGetWindowRect                          = User32.NewProc("GetWindowRect")
	procSetWindowRgn                           = User32.NewProc("SetWindowRgn")
	procSetWindowLongW                         = User32.NewProc("SetWindowLongW")
	procGetWindowLongW                         = User32.NewProc("GetWindowLongW")
)

// SetWindowRgn 设置窗口的窗口区域。窗口区域确定窗口中系统允许绘制的区域。系统不显示窗口区域外部的窗口的任何部分。
//
// hWnd: 窗口句柄。
//
// hRgn: 区域句柄（HRGN）。函数将窗口的窗口区域设置为此区域。如果 hRgn 为 NULL，则该函数将窗口区域设置为 NULL。
//
// bRedraw: 是否重绘窗口。通常，如果窗口可见，请将 bRedraw 设置为 TRUE 。
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-setwindowrgn
func SetWindowRgn(hWnd uintptr, hRgn uintptr, bRedraw bool) bool {
	ret, _, _ := procSetWindowRgn.Call(
		hWnd,
		hRgn,
		common.BoolPtr(bRedraw),
	)
	return ret != 0
}

// SetCursor 设置鼠标光标。
//   - 返回值是上一个游标的句柄（如果有）。
//   - 如果没有上一个游标，则返回值为 0。
//
// hCursor: 要设置的鼠标光标的句柄。
//   - 游标必须由 CreateCursor 或 CreateIconIndirect 函数创建，或者由 LoadCursor 或 LoadImage 函数加载。
//   - 如果此参数为 0，则光标将从屏幕中删除。
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-setcursor
func SetCursor(hCursor uintptr) uintptr {
	ret, _, _ := procSetCursor.Call(hCursor)
	return ret
}

// GetAsyncKeyState 确定在调用此函数时是否按下了键，以及自上次调用 GetAsyncKeyState 以来是否按下了该键。
//   - 如果在调用此函数时按下了键，则返回值的最高位位于位置 1。
//   - 如果自上次调用 GetAsyncKeyState 以来按下了该键，则返回值的最低有效位设置为 1。
//   - 如果在调用此函数时未按下该键，则返回值为零。
//   - 如果自上次调用 GetAsyncKeyState 以来未按下该键，则返回值为零。
//
// vKey: 虚拟键码。有关代码的列表，请参阅虚拟键码: xcc.VK_ .
//
// 返回值: 如果函数成功，则返回值指定最高位和最低有效位的状态。如果函数失败，则返回值为零。要获取扩展错误信息，请调用 GetLastError。
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-getasynckeystate
func GetAsyncKeyState(vKey int32) int16 {
	ret, _, _ := procGetAsyncKeyState.Call(uintptr(vKey))
	return int16(ret)
}

// AdjustWindowRect 计算窗口矩形的所需大小，以便客户区域具有指定的大小。返回的窗口矩形用于传递给 CreateWindowEx 函数来创建具有指定客户区域大小的窗口。
//
// lpRect: 指向 xc.RECT 结构的指针，该结构包含客户区域所需的大小。返回时，结构包含窗口所需的大小。
//
// dwStyle: 要创建的窗口的窗口样式。此参数可以是以下值的组合：xcc.WS_ . 不能指定 xcc.WS_OVERLAPPED 样式。
//
// bMenu: 指示窗口是否有菜单。
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-AdjustWindowRect.
func AdjustWindowRect(lpRect *xc.RECT, dwStyle xcc.WS_, bMenu bool) bool {
	ret, _, _ := procAdjustWindowRect.Call(
		uintptr(unsafe.Pointer(lpRect)),
		uintptr(dwStyle),
		common.BoolPtr(bMenu),
	)
	return ret != 0
}

// SetWindowText 改变指定窗口的标题栏文本（如果有）或指定控件的文本内容。无法更改其他应用程序中控件的文本。
//
// hWnd: 要改变文本的窗口或控件的句柄。
//
// str: 要用作窗口或控件新文本的字符串。
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-setwindowtextw.
func SetWindowText(hWnd uintptr, str string) bool {
	ret, _, _ := procSetWindowTextW.Call(
		hWnd,
		common.StrPtr(str),
	)
	return ret != 0
}

// GetClientRect 检索指定窗口客户区的坐标。客户区是整个窗口的工作区，不包括标题栏、菜单栏和滚动条。
// 左上角坐标总是(0,0)。右下角坐标指定客户区的宽度和高度。
//
// hWnd: 要获取其客户区坐标的窗口句柄。
//
// lpRect: 指向 xc.RECT 结构的指针，该结构接收客户区的坐标。左上角坐标为(0,0)，右下角坐标指定宽度和高度。
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-GetClientRect.
func GetClientRect(hWnd uintptr, lpRect *xc.RECT) bool {
	ret, _, _ := procGetClientRect.Call(
		hWnd,
		uintptr(unsafe.Pointer(lpRect)),
	)
	return ret != 0
}

// GetWindowRect 检索指定窗口的边界矩形的尺寸。 尺寸以相对于屏幕左上角的屏幕坐标提供。
//
// hWnd: 窗口句柄。
//
// lpRect: 指向 xc.RECT 结构的指针，该结构接收窗口左上角和右下角的屏幕坐标。
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-GetWindowRect.
func GetWindowRect(hWnd uintptr, lpRect *xc.RECT) bool {
	ret, _, _ := procGetWindowRect.Call(
		hWnd,
		uintptr(unsafe.Pointer(lpRect)),
	)
	return ret != 0
}

// DestroyWindow 销毁指定的窗口。 函数将 WM_DESTROY 和 WM_NCDESTROY 消息发送到窗口，以停用窗口并从窗口中删除键盘焦点。 如果窗口位于查看器链) 的顶部，函数还会销毁窗口的菜单、销毁计时器、删除剪贴板所有权，并中断剪贴板查看器链 。
//   - 如果指定的窗口是父窗口或所有者窗口， 则 DestroyWindow 会在销毁父窗口或所有者窗口时自动销毁关联的子窗口或拥有窗口。 函数首先销毁子窗口或拥有的窗口，然后销毁父窗口或所有者窗口。
//   - DestroyWindow 还会销毁 CreateDialog 函数创建的无模式对话框。
//
// hwnd: 要检索其上级窗口的句柄。 如果此参数是桌面窗口，则该函数返回 NULL。
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-DestroyWindow.
func DestroyWindow(hwnd uintptr) bool {
	ret, _, _ := procDestroyWindow.Call(
		hwnd,
	)
	return ret != 0
}

// GetAncestor 检索指定窗口的上级句柄。
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-GetAncestor.
//
// hwnd: 要检索其上级窗口的句柄。 如果此参数是桌面窗口，则该函数返回 NULL。
//
// gaFlags: 要检索的上级。 此参数的取值可为下列值之一：wapi.GA_ .
func GetAncestor(hwnd uintptr, gaFlags GA_) uintptr {
	ret, _, _ := procGetAncestor.Call(
		hwnd,
		uintptr(gaFlags),
	)
	return ret
}

// GA_ 指定要检索的上级。
type GA_ uint32

const (
	GA_PARENT    GA_ = 1 // 检索父窗口。 这不包括所有者，因为它与 GetParent 函数一样。
	GA_ROOT      GA_ = 2 // 通过遍历父窗口链来检索根窗口。
	GA_ROOTOWNER GA_ = 3 // 通过遍历 GetParent 返回的父窗口和所有者窗口链来检索拥有的根窗口。
)

// IsDialogMessage 确定消息是否适用于指定的对话框，如果是，则处理该消息。
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-IsDialogMessageW.
//
// hDlg: 对话框的句柄。
//
// msg: 指向包含要检查的消息的 MSG 结构的指针。
func IsDialogMessage(hDlg uintptr, lpMsg *MSG) bool {
	ret, _, _ := procIsDialogMessageW.Call(
		hDlg,
		uintptr(unsafe.Pointer(lpMsg)),
	)
	return ret != 0
}

// PostThreadMessage 将消息发布到指定线程的消息队列。它返回时不等待线程处理消息。
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-PostThreadMessageW.
//
// threadID: 要向其发布消息的线程的标识符。
//   - 如果指定的线程没有消息队列，该函数将失败。 当线程首次调用某个用户或 GDI 函数时，系统会创建线程的消息队列。
//   - 消息发布受 UIPI 的约束。 进程的线程只能将消息发布到较低或等于完整性级别的线程的已发布消息队列。
//   - 此线程必须具有 SE_TCB_NAME 特权，才能将消息发布到属于具有相同本地唯一标识符（LUID）但位于其他桌面中的进程的线程。 否则，函数将失败并返回 ERROR_INVALID_THREAD_ID。
//   - 此线程必须与调用线程属于同一桌面，或者属于具有相同 LUID 的进程。 否则，函数将失败并返回 ERROR_INVALID_THREAD_ID。
//
// msg: 要发布的消息。
//
// wParam: 消息的附加参数。
//
// lParam: 消息的附加参数。
func PostThreadMessage(threadID uint32, msg uint32, wParam, lParam uintptr) bool {
	ret, _, _ := procPostThreadMessageW.Call(
		uintptr(threadID),
		uintptr(msg),
		wParam,
		lParam,
	)
	return ret != 0
}

// DefWindowProc 调用默认窗口过程，为应用程序未处理的任何窗口消息提供默认处理。此函数可确保处理每个消息。返回值是消息处理的结果，取决于消息。
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-DefWindowProcW.
//
// hWnd: 接收消息的窗口过程的句柄。
//
// Msg: 消息。
//
// wParam: 消息的附加参数。
//
// lParam: 消息的附加参数。
func DefWindowProc(hWnd uintptr, Msg uint32, wParam, lParam uintptr) uintptr {
	r, _, _ := procDefWindowProcW.Call(
		hWnd,
		uintptr(Msg),
		wParam,
		lParam,
	)
	return r
}

// CreateWindowEx 创建具有扩展窗口样式的重叠、弹出窗口或子窗口; 否则，此函数与 CreateWindow 函数相同。
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-CreateWindowExW.
//
// exStyle: 窗口的扩展样式。 此参数可以是以下值的组合：xcc.WS_EX_ .
//
// className: 类名。
//
// windowName: 窗口名称。
//
// style: 窗口的样式。 此参数可以是以下值的组合：xcc.WS_ .
//
// x: 窗口的初始水平位置。 对于重叠窗口或弹出窗口，x 参数是窗口左上角的初始 x 坐标，以屏幕坐标表示。 对于子窗口，x 是窗口左上角相对于父窗口工作区左上角的 x 坐标。 如果 x 设置为 CW_USEDEFAULT, 系统将选择窗口左上角的默认位置，并忽略 y 参数。 CW_USEDEFAULT 仅适用于重叠窗口; 如果为弹出窗口或子窗口指定，则 x 和 y 参数设置为零。
//
// y: 窗口的初始垂直位置。 对于重叠窗口或弹出窗口，y 参数是窗口左上角的初始 y 坐标，以屏幕坐标表示。 对于子窗口，y 是子窗口左上角相对于父窗口工作区左上角的初始 y 坐标。 对于列表框 y 是列表框工作区左上角相对于父窗口工作区左上角的初始 y 坐标。如果使用 WS_VISIBLE 样式位设置创建重叠窗口，并将 x 参数设置为 CW_USEDEFAULT，则 y 参数将确定窗口的显示方式。 如果 y 参数 CW_USEDEFAULT，则窗口管理器会在创建窗口后使用 SW_SHOW 标志调用 ShowWindow。 如果 y 参数是一些其他值，则窗口管理器会调用 ShowWindow，该值作为 nCmdShow 参数。
//
// nWidth: 窗口的宽度（以设备单位为单位）。 对于重叠窗口，nWidth 是窗口的宽度、屏幕坐标或 CW_USEDEFAULT。 如果 nWidth 为 CW_USEDEFAULT，则系统会选择窗口的默认宽度和高度;默认宽度从初始 x 坐标扩展到屏幕右边缘;默认高度从初始 y 坐标扩展到图标区域的顶部。 CW_USEDEFAULT 仅适用于重叠窗口;如果为弹出窗口或子窗口指定了 CW_USEDEFAULT，则 nWidth，nHeight 参数设置为零。
//
// nHeight: 窗口的高度（以设备单位为单位）。 对于重叠窗口，nHeight 是窗口的高度（以屏幕坐标为单位）。 如果 nWidth 参数设置为 CW_USEDEFAULT，则系统将忽略 nHeight 。
//
// hWndParent: 正在创建的窗口的父窗口或所有者窗口的句柄。 若要创建子窗口或拥有的窗口，请提供有效的窗口句柄。 对于弹出窗口，此参数是可选的。
//
// hMenu; 菜单的句柄，或指定子窗口标识符，具体取决于窗口样式。 对于重叠或弹出窗口，hMenu 标识要与窗口一起使用的菜单;如果要使用类菜单，它可以 NULL。 对于子窗口，hMenu 指定子窗口标识符，这是对话框控件用来通知其父级事件的整数值。 应用程序确定子窗口标识符;对于具有相同父窗口的所有子窗口，它必须是唯一的。
//
// hInstance: 要与窗口关联的模块实例的句柄。
//
// lpParam: 指向通过 CREATESTRUCT 结构（lpCreateParams 成员）指向的值的指针，该参数由 WM_CREATE 消息的 lParam 参数指向。 此消息在返回之前由此函数发送到创建的窗口。如果应用程序调用 CreateWindow 来创建 MDI 客户端窗口，lpParam 应指向 CLIENTCREATESTRUCT 结构。 如果 MDI 客户端窗口调用 CreateWindow 创建 MDI 子窗口，lpParam 应指向 MDICREATESTRUCT 结构。 如果不需要其他数据，lpParam 可能会 NULL。
func CreateWindowEx(exStyle xcc.WS_EX_, className, windowName string, style xcc.WS_, x, y, nWidth, nHeight int32, hWndParent uintptr, hMenu uintptr, hInstance uintptr, lpParam uintptr) uintptr {
	r, _, _ := procCreateWindowExW.Call(
		uintptr(exStyle),
		common.StrPtr(className),
		common.StrPtr(windowName),
		uintptr(style),
		uintptr(x),
		uintptr(y),
		uintptr(nWidth),
		uintptr(nHeight),
		hWndParent,
		hMenu,
		hInstance,
		lpParam,
	)
	return r
}

// RegisterClassEx 注册一个窗口类，以便在调用 CreateWindow 或 CreateWindowEx 函数时使用。
//   - 如果函数成功，则返回值为唯一标识要注册的类的类原子。 此原子只能由 CreateWindow, CreateWindowEx, GetClassInfo, GetClassInfoEx 使用， FindWindow, FindWindowEx 和 UnregisterClass 函数和 IActiveIMMap：：FilterClientWindows 方法。
//   - 如果函数失败，则返回值为零。 若要获取扩展的错误信息，请调用 GetLastError。
//   - 卸载 DLL 时，不会取消注册 DLL 注册的窗口类。 DLL 必须在卸载时显式注销其类。
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-registerclassexw.
//
// wc: 在将结构传递给函数之前，必须用相应的类属性填充结构。
func RegisterClassEx(wc *WNDCLASSEX) uintptr {
	r, _, _ := procRegisterClassExW.Call(uintptr(unsafe.Pointer(wc)))
	return r
}

// WNDCLASSEX 包含窗口类信息。它与 RegisterClassEx 和 GetClassInfoEx 函数一起使用。
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/ns-winuser-wndclassexw.
type WNDCLASSEX struct {
	CbSize      uint32  // 此结构的大小（以字节为单位）。 将此成员设置为 sizeof(WNDCLASSEX)。 在调用 GetClassInfoEx 函数之前，请务必设置此成员。
	Style       uint32  // 类样式（s）。 此成员可以是 类样式的任意组合。
	LpfnWndProc uintptr // 指向窗口过程的指针。 必须使用 CallWindowProc 函数来调用窗口过程。 有关详细信息，请参阅 WindowProc。
	CbClsExtra  int32   // 要按照窗口类结构分配的额外字节数。 系统将字节初始化为零。
	CbWndExtra  int32   // 在窗口实例之后分配的额外字节数。 系统将字节初始化为零。 如果应用程序使用 WNDCLASSEX 注册通过使用资源文件中的 CLASS 指令创建的对话框，则必须将此成员设置为 DLGWINDOWEXTRA。
	HInstance   uintptr // 包含类的窗口过程的实例的句柄。
	HIcon       uintptr // 类图标的句柄。 此成员必须是图标资源的句柄。 如果此成员 NULL，则系统提供默认图标。
	HCursor     uintptr // 类游标的句柄。 此成员必须是游标资源的句柄。 如果此成员 NULL，则每当鼠标移动到应用程序的窗口中时，应用程序都必须显式设置光标形状。

	// 类背景画笔的句柄。 此成员可以是用于绘制背景的画笔的句柄，也可以是颜色值。 颜色值必须是以下标准系统颜色之一（值 1 必须添加到所选颜色中）。 如果提供了颜色值，则必须将其转换为以下 HBRUSH 类型之一：
	//
	// 使用 UnregisterClass 注销类时，系统会自动删除类背景画笔。 应用程序不应删除这些画笔。
	//
	// 当此成员 NULL时，每当请求应用程序在其工作区中绘制时，都必须绘制其自己的背景。 若要确定是否必须绘制背景，应用程序可以处理 WM_ERASEBKGND 消息，也可以测试由 beginPaint 函数 填充的 PAINTSTRUCT 结构的 fErase 成员。
	HbrBackground uintptr
	LpszMenuName  uintptr // 指向以 null 结尾的字符串的指针, 使用 common.StrPtr 生成，该字符串指定类菜单的资源名称，因为名称显示在资源文件中。  如果此成员 NULL，则属于此类的窗口没有默认菜单。

	// 指向以 null 结尾的字符串或原子的指针, 使用 common.StrPtr 生成。 如果此参数是 atom，则它必须是上一次调用 RegisterClass 或 RegisterClassEx 函数创建的类 atom。 原子必须位于 lpszClassName 的低序单词中;高序单词必须为零。
	//
	// 如果 lpszClassName 是字符串，则指定窗口类名。 类名称可以是注册到 RegisterClass 或 RegisterClassEx 的任何名称，也可以是预定义的控件类名称。
	//
	// lpszClassName 的最大长度为 256。 如果 lpszClassName 大于最大长度，则 RegisterClassEx 函数将失败。
	LpszClassName uintptr
	HIconSm       uintptr // 与窗口类关联的小图标的句柄。 如果此成员 NULL，系统将搜索由 hIcon 成员指定的图标资源，以获取要用作小图标的相应大小的图标。
}

// GetSystemMetrics 检索指定的系统指标或系统配置设置。检索的所有维度都以像素为单位。
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-GetSystemMetrics.
//
// nIndex: 要检索的系统指标或配置设置。 此参数的取值可为下列值之一：wapi.SM_ . 请注意: 所有SM_CX* 值为宽度，所有SM_CY* 值为高度。
func GetSystemMetrics(nIndex SM_) int32 {
	ret, _, _ := procGetSystemMetrics.Call(uintptr(nIndex))
	return int32(ret)
}

// SM_ 定义了一些常用的系统度量索引
type SM_ int32

const (
	SM_CXSCREEN          SM_ = 0      // 屏幕宽度（像素）
	SM_CYSCREEN          SM_ = 1      // 屏幕高度（像素）
	SM_CXVSCROLL         SM_ = 2      // 垂直滚动条箭头宽度
	SM_CYHSCROLL         SM_ = 3      // 水平滚动条箭头高度
	SM_CYCAPTION         SM_ = 4      // 标题栏高度
	SM_CXBORDER          SM_ = 5      // 窗口边框宽度
	SM_CYBORDER          SM_ = 6      // 窗口边框高度
	SM_CXDLGFRAME        SM_ = 7      // 对话框边框宽度
	SM_CYDLGFRAME        SM_ = 8      // 对话框边框高度
	SM_CYVTHUMB          SM_ = 9      // 垂直滚动条滑块高度
	SM_CXHTHUMB          SM_ = 10     // 水平滚动条滑块宽度
	SM_CXICON            SM_ = 11     // 默认图标宽度
	SM_CYICON            SM_ = 12     // 默认图标高度
	SM_CXCURSOR          SM_ = 13     // 光标宽度
	SM_CYCURSOR          SM_ = 14     // 光标高度
	SM_CYMENU            SM_ = 15     // 菜单栏高度
	SM_CXFULLSCREEN      SM_ = 16     // 全屏窗口的客户区宽度
	SM_CYFULLSCREEN      SM_ = 17     // 全屏窗口的客户区高度
	SM_CYKANJIWINDOW     SM_ = 18     // 日语汉字窗口高度（已废弃）
	SM_MOUSEPRESENT      SM_ = 19     // 鼠标是否存在（0=无，非零=有）
	SM_CYVSCROLL         SM_ = 20     // 垂直滚动条高度
	SM_CXHSCROLL         SM_ = 21     // 水平滚动条宽度
	SM_DEBUG             SM_ = 22     // 是否启用调试版本
	SM_SWAPBUTTON        SM_ = 23     // 鼠标左右键是否交换
	SM_CXMIN             SM_ = 28     // 窗口最小宽度
	SM_CYMIN             SM_ = 29     // 窗口最小高度
	SM_CXSIZE            SM_ = 30     // 标题栏按钮宽度
	SM_CYSIZE            SM_ = 31     // 标题栏按钮高度
	SM_CXFRAME           SM_ = 32     // 可调整边框宽度（同 SM_CXSIZEFRAME ）
	SM_CYFRAME           SM_ = 33     // 可调整边框高度（同 SM_CYSIZEFRAME ）
	SM_CXMINTRACK        SM_ = 34     // 窗口最小跟踪宽度
	SM_CYMINTRACK        SM_ = 35     // 窗口最小跟踪高度
	SM_CXDOUBLECLK       SM_ = 36     // 双击有效区域宽度
	SM_CYDOUBLECLK       SM_ = 37     // 双击有效区域高度
	SM_CXICONSPACING     SM_ = 38     // 图标排列单元格宽度
	SM_CYICONSPACING     SM_ = 39     // 图标排列单元格高度
	SM_MENUDROPALIGNMENT SM_ = 40     // 菜单弹出方向（0=左对齐，非零=右对齐）
	SM_PENWINDOWS        SM_ = 41     // PenWindows 是否加载
	SM_DBCSENABLED       SM_ = 42     // 是否启用 DBCS 字符集
	SM_CMOUSEBUTTONS     SM_ = 43     // 鼠标按钮数量（0=无鼠标）
	SM_SECURE            SM_ = 44     // 是否启用安全模式
	SM_CXEDGE            SM_ = 45     // 3D边框宽度
	SM_CYEDGE            SM_ = 46     // 3D边框高度
	SM_CXMINSPACING      SM_ = 47     // 图标网格单元格宽度
	SM_CYMINSPACING      SM_ = 48     // 图标网格单元格高度
	SM_CXSMICON          SM_ = 49     // 小图标建议宽度
	SM_CYSMICON          SM_ = 50     // 小图标建议高度
	SM_CYSMCAPTION       SM_ = 51     // 小标题栏高度
	SM_CXSMSIZE          SM_ = 52     // 小标题栏按钮宽度
	SM_CYSMSIZE          SM_ = 53     // 小标题栏按钮高度
	SM_CXMENUSIZE        SM_ = 54     // 菜单栏按钮宽度
	SM_CYMENUSIZE        SM_ = 55     // 菜单栏按钮高度
	SM_ARRANGE           SM_ = 56     // 排列方向标志
	SM_CXMINIMIZED       SM_ = 57     // 最小化窗口宽度
	SM_CYMINIMIZED       SM_ = 58     // 最小化窗口高度
	SM_CXMAXTRACK        SM_ = 59     // 窗口最大可调宽度
	SM_CYMAXTRACK        SM_ = 60     // 窗口最大可调高度
	SM_CXMAXIMIZED       SM_ = 61     // 最大化窗口宽度
	SM_CYMAXIMIZED       SM_ = 62     // 最大化窗口高度
	SM_NETWORK           SM_ = 63     // 网络存在标志（最低位=1表示存在）
	SM_CLEANBOOT         SM_ = 67     // 启动模式（0=正常，1=安全模式，2=带网络的安全模式）
	SM_CXDRAG            SM_ = 68     // 拖动生效区域宽度
	SM_CYDRAG            SM_ = 69     // 拖动生效区域高度
	SM_SHOWSOUNDS        SM_ = 70     // 是否强制视觉提示代替声音
	SM_CXMENUCHECK       SM_ = 71     // 菜单复选框宽度
	SM_CYMENUCHECK       SM_ = 72     // 菜单复选框高度
	SM_SLOWMACHINE       SM_ = 73     // 是否低性能计算机
	SM_MIDEASTENABLED    SM_ = 74     // 是否启用中东语言支持
	SM_MOUSEWHEELPRESENT SM_ = 75     // 是否支持鼠标滚轮
	SM_XVIRTUALSCREEN    SM_ = 76     // 虚拟屏幕左上角X坐标
	SM_YVIRTUALSCREEN    SM_ = 77     // 虚拟屏幕左上角Y坐标
	SM_CXVIRTUALSCREEN   SM_ = 78     // 虚拟屏幕宽度
	SM_CYVIRTUALSCREEN   SM_ = 79     // 虚拟屏幕高度
	SM_CMONITORS         SM_ = 80     // 显示器数量
	SM_SAMEDISPLAYFORMAT SM_ = 81     // 所有显示器颜色格式是否相同
	SM_IMMENABLED        SM_ = 82     // 是否启用输入法
	SM_CXFOCUSBORDER     SM_ = 83     // 焦点边框宽度
	SM_CYFOCUSBORDER     SM_ = 84     // 焦点边框高度
	SM_TABLETPC          SM_ = 86     // 是否是 Tablet PC
	SM_MEDIACENTER       SM_ = 87     // 是否是 Media Center Edition
	SM_STARTER           SM_ = 88     // 是否是 Windows Starter Edition
	SM_SERVERR2          SM_ = 89     // 是否是 Windows Server 2003 R2
	SM_REMOTESESSION     SM_ = 0x1000 // 是否在远程会话中
)

// PeekMessage 调度传入的非排队消息，检查已发布的消息的线程消息队列，并检索消息（如果有）。如果消息可用，则返回值为true。如果没有消息可用，则返回值为false。
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-PeekMessageW.
//
// lpMsg: 指向 MSG 结构的指针，该结构接收消息。
//
// hWnd: 要检索其消息队列的窗口的句柄。窗口必须属于当前线程。如果 hWnd 为 NULL，则函数检索与调用线程关联的窗口的消息队列。
//
// wMsgFilterMin: 指定要检索的最小消息 ID。 如果 wMsgFilterMin 为零，则不检查最小消息 ID。使用 WM_KEYFIRST（0x0100）指定第一条键盘消息或 WM_MOUSEFIRST（0x0200）来指定第一条鼠标消息。如果 wMsgFilterMin 和 wMsgFilterMax 均为零，PeekMessage 将返回所有可用消息（即，不执行范围筛选）。
//
// wMsgFilterMax: 指定要检索的最大消息 ID。 使用 WM_KEYLAST 指定最后一条键盘消息或 WM_MOUSELAST 指定最后一条鼠标消息。如果 wMsgFilterMin 和 wMsgFilterMax 均为零，PeekMessage 将返回所有可用消息（即，不执行范围筛选）。
//
// wRemoveMsg: 指定如何处理消息。此参数可以是以下一个或多个值: wapi.PM_。默认情况下，将处理所有消息类型。 若要指定只应处理某些消息，请指定以下一个或多个值: wapi.PM_QS_。
func PeekMessage(lpMsg *MSG, hWnd uintptr, wMsgFilterMin, wMsgFilterMax, wRemoveMsg uint32) bool {
	ret, _, _ := procPeekMessageW.Call(
		uintptr(unsafe.Pointer(lpMsg)),
		hWnd,
		uintptr(wMsgFilterMin),
		uintptr(wMsgFilterMax),
		uintptr(wRemoveMsg),
	)
	return ret != 0
}

const (
	PM_NOREMOVE uint32 = 0x0000 // PeekMessage 处理后，消息不会从队列中删除。
	PM_REMOVE   uint32 = 0x0001 // PeekMessage 处理后，将从队列中删除消息。
	PM_NOYIELD  uint32 = 0x0002 // 防止系统释放等待调用方空闲的任何线程（请参阅 WaitForInputIdle）。将此值与 PM_NOREMOVE 或 PM_REMOVE 组合在一起。
)

const (
	PM_QS_INPUT       = QS_INPUT << 16                                // 处理鼠标和键盘消息
	PM_QS_PAINT       = QS_PAINT << 16                                // 处理绘制消息
	PM_QS_POSTMESSAGE = (QS_POSTMESSAGE | QS_HOTKEY | QS_TIMER) << 16 // 处理所有已发布的消息，包括计时器和热键
	PM_QS_SENDMESSAGE = QS_SENDMESSAGE << 16                          // 处理所有已发送的消息
)

const (
	QS_KEY                  uint32 = 0x0001                          // 键盘消息
	QS_MOUSEMOVE            uint32 = 0x0002                          // 鼠标移动消息
	QS_MOUSEBUTTON          uint32 = 0x0004                          // 鼠标按钮消息
	QS_POSTMESSAGE          uint32 = 0x0008                          // 已发布的消息
	QS_TIMER                uint32 = 0x0010                          // 计时器消息
	QS_PAINT                uint32 = 0x0020                          // 绘制消息
	QS_SENDMESSAGE          uint32 = 0x0040                          // 已发送的消息
	QS_HOTKEY               uint32 = 0x0080                          // 热键消息
	QS_ALLPOSTMESSAGEuint32 uint32 = 0x0100                          // 所有已发布的消息
	QS_RAWINPUT             uint32 = 0x0400                          // 原始输入消息
	QS_MOUSE                       = QS_MOUSEMOVE | QS_MOUSEBUTTON   // 所有鼠标消息
	QS_INPUT                       = QS_MOUSE | QS_KEY | QS_RAWINPUT // 所有输入消息
)

const (
	CS_BYTEALIGNCLIENT uint32 = 0x1000     // 将窗口的工作区与 x 方向的字节边界对齐. 此样式会影响窗口的宽度及其在显示器上的水平位置。
	CS_BYTEALIGNWINDOW uint32 = 0x2000     // 使窗口在字节边界 (沿 x 方向) 对齐. 此样式会影响窗口的宽度及其在显示器上的水平位置。
	CS_CLASSDC         uint32 = 0x0040     // 分配一个设备上下文，以便类中的所有窗口共享. 由于窗口类特定于进程，因此应用程序的多个线程可以创建同一类的窗口。 线程还可以尝试同时使用设备上下文。 发生这种情况时，系统仅允许一个线程成功完成其绘制操作。
	CS_DBLCLKS         uint32 = 0x0008     // 当用户在光标位于属于类的窗口中时双击鼠标时，将双击消息发送到窗口过程
	CS_DROPSHADOW      uint32 = 0x00020000 // 在窗口上启用投影效果. 通过 SPI_SETDROPSHADOW 打开和关闭效果。 通常，对于小型、生存期较短的窗口（如菜单）启用此功能，以强调其与其他窗口的 Z 顺序关系。 从具有此样式的类创建的 Windows 必须是顶级窗口;它们可能不是子窗口。
	CS_GLOBALCLASS     uint32 = 0x4000     // 指示窗口类是应用程序全局类
	CS_HREDRAW         uint32 = 0x0002     // 如果移动或大小调整更改了工作区的宽度，则重绘整个窗口
	CS_NOCLOSE         uint32 = 0x0200     // 在窗口菜单上禁用“关闭”
	CS_OWNDC           uint32 = 0x0020     // 为类中的每个窗口分配唯一的设备上下文
	CS_PARENTDC        uint32 = 0x0080     // 将子窗口的剪裁矩形设置为父窗口的剪裁矩形，以便子窗口可以在父窗口上绘制. 具有 CS_PARENTDC 样式位的窗口从系统的设备上下文缓存接收常规设备上下文。 它不会为子级提供父级的设备上下文或设备上下文设置。 指定 CS_PARENTDC 可增强应用程序的性能。

	// 保存此类窗口遮盖的屏幕图像部分作为位图。 删除窗口时，系统会使用保存的位图还原屏幕图像，包括被遮盖的其他窗口。 因此，如果位图使用的内存尚未丢弃，并且其他屏幕操作未使存储的图像失效，则系统不会将 WM_PAINT 消息发送到被遮盖的窗口。
	//
	// 此样式适用于小型窗口 (例如菜单或对话框) ，这些菜单或对话框在发生其他屏幕活动之前会短暂显示，然后删除。 此样式会增加显示窗口所需的时间，因为系统必须先分配内存来存储位图。
	CS_SAVEBITS uint32 = 0x0800
	CS_VREDRAW  uint32 = 0x0001 // 如果移动或大小调整更改了工作区的高度，则重新绘制整个窗口
)

// GetParent 检索指定窗口的父级或所有者的句柄。若要检索指定上级的句柄，请使用 GetAncestor 函数。
//   - 如果窗口是子窗口，则返回值是父窗口的句柄。 如果窗口是具有 WS_POPUP 样式的顶级窗口，则返回值是所有者窗口的句柄。
//   - 如果函数失败，则返回值为 NULL。 要获得更多的错误信息，请调用 GetLastError。
//   - 此函数通常由于以下原因之一而失败：1. 该窗口是无所有者或没有 WS_POPUP 样式的顶级窗口。2. 所有者窗口具有 WS_POPUP 样式。
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-GetParent.
//
// hWnd: 要检索其父窗口句柄的窗口的句柄。
func GetParent(hWnd uintptr) uintptr {
	ret, _, _ := procGetParent.Call(hWnd)
	return ret
}

// IsWindowVisible 确定指定窗口的可见性状态。
//   - 如果指定的窗口、其父窗口等具有 WS_VISIBLE 样式，则返回值为非零。 否则返回值为零。
//   - 由于返回值指定窗口是否具有 WS_VISIBLE 样式，因此即使窗口被其他窗口完全遮挡，也可能为非零。
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-IsWindowVisible.
//
// hWnd: 窗口的句柄。
func IsWindowVisible(hWnd uintptr) bool {
	ret, _, _ := procIsWindowVisible.Call(hWnd)
	return ret != 0
}

// GetClassName 检索指定窗口所属的类的名称。
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-GetClassName.
//
// hWnd: 窗口的句柄，以及窗口所属的类的间接句柄。
//
// lpClassName: 类名字符串。
//
// nMaxCount: lpClassName 缓冲区的长度（以字符为单位）。
//   - 缓冲区必须足够大，才能包含终止 null 字符; 否则，类名字符串将被截断为 nMaxCount-1 字符。
//
// 返回值:
//   - 如果函数成功，则返回值是复制到缓冲区的字符数，不包括终止 null 字符。
//   - 如果函数失败，则返回值为零。
func GetClassName(hWnd uintptr, lpClassName *string, nMaxCount int) int {
	buf := make([]uint16, nMaxCount)
	ret, _, _ := procGetClassName.Call(hWnd, uintptr(unsafe.Pointer(&buf[0])), uintptr(nMaxCount))
	*lpClassName = syscall.UTF16ToString(buf[:])
	return int(ret)
}

// EnumWindows 通过将句柄传递到每个窗口，进而将传递给应用程序定义的回调函数，枚举屏幕上的所有顶级窗口。
//   - 枚举窗口 将一直持续到最后一个顶级窗口被枚举或回调函数返回 FALSE。
//   - EnumWindows 函数不枚举子窗口，但系统拥有的几个具有 WS_CHILD 样式的顶级窗口除外。
//   - 此函数比在循环中调用 GetWindow 函数更可靠。 调用 GetWindow 以执行此任务的应用程序有被捕获到无限循环或引用已销毁窗口的句柄的风险。
//   - 如果 EnumWindowsProc 返回零，则返回值也为零。 在这种情况下，回调函数应调用 SetLastError 以获取要返回到 EnumWindows 调用方有意义的错误代码。
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-EnumWindows.
//
// lpEnumFunc: 指向应用程序定义的回调函数的指针。使用 syscall.NewCallback 生成. 格式: func (hwnd, lParam uintptr) uintptr {return 1}, return 0 可立刻终止.
//
// lParam: 要传递给回调函数的应用程序定义值。
func EnumWindows(lpEnumFunc uintptr, lParam uintptr) bool {
	ret, _, _ := procEnumWindows.Call(lpEnumFunc, lParam)
	return ret != 0
}

// EnumWindowsEx 通过将句柄传递到每个窗口，进而将传递给应用程序定义的回调函数，枚举屏幕上的所有顶级窗口。
//   - 枚举窗口 将一直持续到最后一个顶级窗口被枚举或回调函数返回 FALSE。
//   - EnumWindowsEx 函数不枚举子窗口，但系统拥有的几个具有 WS_CHILD 样式的顶级窗口除外。
//   - 此函数比在循环中调用 GetWindow 函数更可靠。 调用 GetWindow 以执行此任务的应用程序有被捕获到无限循环或引用已销毁窗口的句柄的风险。
//   - 如果 EnumWindowsProc 返回零，则返回值也为零。 在这种情况下，回调函数应调用 SetLastError 以获取要返回到 EnumWindowsEx 调用方有意义的错误代码。
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-EnumWindows.
//
// EnumWindowsProc: 回调函数. return 0 可立刻终止. return 1 继续枚举.
//
// lParam: 要传递给回调函数的应用程序定义值。
func EnumWindowsEx(EnumWindowsProc func(hwnd, lParam uintptr) uintptr, lParam uintptr) bool {
	ret, _, _ := procEnumWindows.Call(syscall.NewCallback(EnumWindowsProc), lParam)
	return ret != 0
}

// SetFocus 将键盘焦点设置为指定的窗口。窗口必须附加到调用线程的消息队列。如果函数成功，则返回值是以前具有键盘焦点的窗口的句柄。 如果 hWnd 参数无效或窗口未附加到调用线程的消息队列，则返回值为 NULL。 若要获取扩展错误信息，请调用 GetLastError 函数。
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-SetFocus.
//
// hWnd: 将接收键盘输入的窗口的句柄。 如果此参数为 NULL，则忽略击键。
func SetFocus(hwnd uintptr) uintptr {
	ret, _, _ := procSetFocus.Call(
		hwnd,
	)
	return ret
}

// UpdateWindow 如果窗口的更新区域不为空， UpdateWindow 函数通过向窗口发送 WM_PAINT 消息来更新指定窗口的工作区。 函数绕过应用程序队列，将 WM_PAINT 消息直接发送到指定窗口的窗口过程。 如果更新区域为空，则不发送任何消息。
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-UpdateWindow.
//
// hWnd: 要更新的窗口的句柄。
func UpdateWindow(hWnd uintptr) bool {
	ret, _, _ := procUpdateWindow.Call(
		hWnd,
	)
	return ret != 0
}

// ShowWindow 设置指定窗口的显示状态。
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-ShowWindow.
//
// hWnd: 窗口的句柄。
//
// nCmdShow: 指定窗口如何显示: xcc.SW_ .
func ShowWindow(hWnd uintptr, nCmdShow xcc.SW_) bool {
	ret, _, _ := procShowWindow.Call(
		hWnd,
		uintptr(nCmdShow),
	)
	return ret != 0
}

// SetParent 更改指定子窗口的父窗口。如果函数成功，则返回值是上一个父窗口的句柄。失败则为0.
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-SetParent.
//
// hWndChild: 子窗口的句柄。
//
// hWndNewParent: 新父窗口的句柄。 如果此参数为 NULL，桌面窗口将成为新的父窗口。 如果此参数 HWND_MESSAGE，则子窗口将成为 仅消息窗口。
func SetParent(hWndChild, hWndNewParent uintptr) uintptr {
	ret, _, _ := procSetParent.Call(hWndChild, hWndNewParent)
	return ret
}

// MoveWindow 更改指定窗口的位置和尺寸。 对于顶级窗口，位置和尺寸是相对于屏幕左上角的。 对于子窗口，它们相对于父窗口工作区的左上角。
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-MoveWindow.
//
// hWnd: 窗口句柄。
//
// x: 窗口左上角的新 x 坐标。
//
// y: 窗口左上角的新 y 坐标。
//
// width: 窗口的新宽度。
//
// height: 窗口的新高度。
//
// bRepaint: 是否重绘窗口。这适用于工作区、非工作区 (包括标题栏和滚动条) ，以及由于移动子窗口而发现父窗口的任何部分。
func MoveWindow(hWnd uintptr, x, y, width, height int32, bRepaint bool) bool {
	r, _, _ := procMoveWindow.Call(
		hWnd,
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		common.BoolPtr(bRepaint),
	)
	return r != 0
}

// GetWindowThreadProcessId 检索创建指定窗口的线程的标识符，以及创建该窗口的进程（可选）的标识符。
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-GetWindowThreadProcessId.
//
// hWnd: 窗口句柄。
func GetWindowThreadProcessId(hWnd uintptr) uint32 {
	var processId uint32
	procGetWindowThreadProcessId.Call(hWnd, uintptr(unsafe.Pointer(&processId)))
	return processId
}

// PhysicalToLogicalPointForPerMonitorDPI 将窗口中的点从物理坐标转换为逻辑坐标，而不考虑每英寸点数 (dpi) 对调用者的感知。
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-physicaltologicalpointforpermonitordpi.
//
// hWnd: 窗口句柄。
//
// lpPoint: 指向 POINT 结构的指针，包含物理坐标，函数会将其转换为逻辑坐标。
func PhysicalToLogicalPointForPerMonitorDPI(hWnd uintptr, lpPoint *POINT) bool {
	ret, _, _ := procPhysicalToLogicalPointForPerMonitorDPI.Call(
		hWnd, uintptr(unsafe.Pointer(lpPoint)),
	)
	return ret != 0
}

// 窗口属性索引常量.
const (
	// GWL_EXSTYLE 设置新的扩展窗口样式。
	GWL_EXSTYLE int32 = -20

	// GWLP_HINSTANCE 设置新的应用程序实例句柄。
	GWLP_HINSTANCE int32 = -6

	// GWLP_ID 设置子窗口的新标识符。窗口不能是顶级窗口。
	GWLP_ID int32 = -12

	// GWL_STYLE 设置新的窗口样式。
	GWL_STYLE int32 = -16

	// GWLP_USERDATA 设置与窗口关联的用户数据。
	// 此数据供创建窗口的应用程序使用。其值最初为零。
	GWLP_USERDATA int32 = -21

	// GWLP_WNDPROC 设置窗口过程的新地址。
	GWLP_WNDPROC int32 = -4
)

// 对话框属性索引常量
const (
	// DWLP_DLGPROC 设置指向对话框过程的新指针。
	DWLP_DLGPROC = unsafe.Sizeof(uintptr(0)) * 4 // 通常为 4 或 8，取决于系统架构（32 位或 64 位）

	// DWLP_MSGRESULT 设置在对话框过程中处理的消息的返回值。
	DWLP_MSGRESULT = 0

	// DWLP_USER 设置与对话框关联的用户数据。
	DWLP_USER = DWLP_DLGPROC + unsafe.Sizeof(uintptr(0)) // 通常为 DWLP_DLGPROC + 4 或 +8
)

type WH_ int32

const (
	WH_JOURNALRECORD   WH_ = 0  // 用于记录发布到系统消息队列的输入消息。此挂钩可用于录制宏, win11不支持.
	WH_JOURNALPLAYBACK WH_ = 1  // 该过程发布以前由 WH_JOURNALRECORD 挂钩过程记录的消息, win11不支持.
	WH_KEYBOARD        WH_ = 2  // 安装监视击键消息的挂钩过程.
	WH_GETMESSAGE      WH_ = 3  // 安装用于监视发布到消息队列的消息的挂钩过程.
	WH_CALLWNDPROC     WH_ = 4  // 安装一个挂钩过程，该过程在系统将消息发送到目标窗口过程之前对其进行监视.
	WH_CBT             WH_ = 5  // 安装一个挂钩过程，用于接收对 CBT 应用程序有用的通知.
	WH_SYSMSGFILTER    WH_ = 6  // 安装一个挂钩过程，用于监视由于对话框、消息框、菜单或滚动条中的输入事件而生成的消息。 挂钩过程监视与调用线程相同的桌面中的所有应用程序的消息.
	WH_MOUSE           WH_ = 7  // 安装监视鼠标消息的挂钩过程.
	WH_DEBUG           WH_ = 9  // 安装一个挂钩过程，用于调试其他挂钩过程.
	WH_SHELL           WH_ = 10 // 安装一个挂钩过程，用于接收对 shell 应用程序有用的通知.
	WH_FOREGROUNDIDLE  WH_ = 11 // 安装一个挂钩过程，当应用程序的前景线程即将变为空闲状态时将调用该挂钩过程。 此挂钩可用于在空闲时间执行低优先级任务.
	WH_CALLWNDPROCRET  WH_ = 12 // 安装挂钩过程，该挂钩过程在目标窗口过程处理消息后对其进行监视。 有关详细信息.
	WH_KEYBOARD_LL     WH_ = 13 // 安装监视低级别键盘输入事件的挂钩过程.
	WH_MOUSE_LL        WH_ = 14 // 安装用于监视低级别鼠标输入事件的挂钩过程.
)

// KBDLLHOOKSTRUCT 包含有关低级别键盘输入事件的信息.
type KBDLLHOOKSTRUCT struct {
	VkCode      uint32  // 虚拟按键代码, xcc.VK_ . 详情: https://learn.microsoft.com/zh-cn/windows/win32/inputdev/virtual-key-codes.
	ScanCode    uint32  // 按键代码的硬件扫描代码.
	Flags       uint32  // 扩展键标志、事件注入标志、上下文代码和转换状态标志。此成员指定如下。应用程序可以使用以下值来测试击键标志。测试LLKHF_INJECTED (位 4) 将告知是否已注入事件。如果是，则测试 LLKHF_LOWER_IL_INJECTED (位 1) 会告诉你事件是否是从以较低完整性级别运行的进程注入的.
	Time        uint32  // 此消息的时间戳，相当于 GetMessageTime 为此消息返回的时间戳.
	DwExtraInfo uintptr // 与消息关联的其他信息.
}

// LowLevelKeyboardProc 是一个低级键盘钩子过程，它将接收有关键盘消息的信息.
//
// nCode: 挂钩过程用于确定如何处理消息的代码. 如果 nCode 小于零，则挂钩过程必须将消息传递给 CallNextHookEx 函数，而无需进一步处理，并且应返回 CallNextHookEx 返回的值.
//
// wParam: 键盘消息的标识符. 可以是以下消息之一： xcc.WM_KEYDOWN、xcc.WM_KEYUP、xcc.WM_SYSKEYDOWN 或 xcc.WM_SYSKEYUP.
//
// LPARAM: 指向 KBDLLHOOKSTRUCT 结构的指针.
//
// 返回值: 如果 nCode 小于零，则挂钩过程必须返回 CallNextHookEx 返回的值. 如果 nCode 大于或等于零，并且挂钩过程未处理消息，强烈建议调用 CallNextHookEx 并返回它返回的值;否则，安装 WH_KEYBOARD_LL 挂钩的其他应用程序将不会收到挂钩通知，因此行为可能不正确. 如果挂钩过程处理了消息，它可能会返回非零值，以防止系统将消息传递给挂钩链的其余部分或目标窗口过程.
type LowLevelKeyboardProc func(nCode int32, wParam xcc.WM_, lParam *KBDLLHOOKSTRUCT) uintptr

// WHEEL_DELTA 一次标准滚轮滚动的增量.
const WHEEL_DELTA int16 = 120

type MSLLHOOKSTRUCT struct {
	PT POINT // 鼠标光标的屏幕坐标

	// 说明:
	//  - 如果消息是 xcc.WM_MOUSEWHEEL 鼠标滚轮滚动消息，正值表示滚轮向前/上旋转（远离用户）；负值表示滚轮向后/下旋转（朝向用户）。此成员的高位是滚轮增量, 可使用 wutil.GetHigh16Bits()来取高位值。保留低位。一次标准滚轮滚动的增量定义为 wapi.WHEEL_DELTA，即 120.
	//  - 如果消息是 xcc.WM_XBUTTONDOWN、xcc.WM_XBUTTONUP、xcc.WM_XBUTTONDBLCLK、xcc.WM_NCXBUTTONDOWN、xcc.WM_NCXBUTTONUP 或 xcc.WM_NCXBUTTONDBLCLK，则高位指定按下或释放的 X 按钮, 可使用 wutil.GetHigh16Bits()来取高位值。并且保留低位。此值可以是以下一个或多个值: 1. 按下或释放第一个 X 按钮; 2. 按下或释放第二个 X 按钮。否则，不使用 mouseData.
	MouseData int32

	Flags       uint32  // 事件注入的标志。应用程序可以使用以下值来测试标志。测试 LLMHF_INJECTED (位 0) 将告知是否已注入事件。如果是，则测试 LLMHF_LOWER_IL_INJECTED (位 1) 将告诉你事件是否是从以较低完整性级别运行的进程注入的.
	Time        uint32  // 此消息的时间戳.
	DwExtraInfo uintptr // 与消息关联的其他信息.
}

// LowLevelMouseProc 是一个低级鼠标钩子过程，它将接收有关鼠标消息的信息.
//
// nCode: 挂钩过程用于确定如何处理消息的代码. 如果 nCode 小于零，则挂钩过程必须将消息传递给 CallNextHookEx 函数，而无需进一步处理，并且应返回 CallNextHookEx 返回的值.
//
// wParam: 鼠标消息的标识符. 可以是以下消息之一：xcc.WM_LBUTTONDOWN、xcc.WM_LBUTTONUP、xcc.WM_MOUSEMOVE、xcc.WM_MOUSEWHEEL、xcc.WM_RBUTTONDOWN 或 xcc.WM_RBUTTONUP.
//
// LPARAM: 指向 MSLLHOOKSTRUCT 结构的指针.
//
// 返回值: 如果 nCode 小于零，则挂钩过程必须返回 CallNextHookEx 返回的值. 如果 nCode 大于或等于零，并且挂钩过程未处理消息，强烈建议调用 CallNextHookEx 并返回它返回的值;否则，安装 WH_MOUSE_LL 挂钩的其他应用程序将不会收到挂钩通知，因此行为可能不正确. 如果挂钩过程处理了消息，它可能会返回非零值，以防止系统将消息传递给挂钩链的其余部分或目标窗口过程.
type LowLevelMouseProc func(nCode int32, wParam xcc.WM_, lParam *MSLLHOOKSTRUCT) uintptr

// SetWindowsHookExW 将应用程序定义的挂钩过程安装到挂钩链中。 需要安装挂钩过程来监视系统的某些类型的事件。 这些事件与特定线程或与调用线程位于同一桌面中的所有线程相关联.
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-setWindowsHookExW.
//
// idHook: 要安装的挂钩过程的类型, wapi.WH_ .
//
// lpfn: 回调函数指针, 使用 syscall.NewCallback()生成. 如果 dwThreadId 参数为零或指定由其他进程创建的线程的标识符，则 lpfn 参数必须指向 DLL 中的挂钩过程。否则，lpfn 可以指向与当前进程关联的代码中的挂钩过程.
//
// hmod: DLL 的句柄，包含 lpfn 参数指向的挂钩过程。如果 dwThreadId 参数指定当前进程创建的线程，并且挂钩过程位于与当前进程关联的代码中，则必须将 hMod 参数设置为 0.
//
// dwThreadId: 要与之关联的挂钩过程的线程的标识符。对于桌面应用，如果此参数为零，则挂钩过程与调用线程在同一桌面中运行的所有现有线程相关联.
//
// 返回值: 如果成功，则返回挂钩过程的句柄; 如果失败，则返回 0.
func SetWindowsHookExW(idHook WH_, lpfn uintptr, hmod uintptr, dwThreadId uint32) uintptr {
	r, _, _ := setWindowsHookExW.Call(uintptr(idHook), lpfn, hmod, uintptr(dwThreadId))
	return r
}

// CallNextHookEx 将挂钩信息传递给当前挂钩链中的下一个挂钩过程。挂钩过程可以在处理挂钩信息之前或之后调用此函数.
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-callNextHookEx.
//
// hhk: 忽略此参数. 可填0.
//
// nCode: 传递给当前挂钩过程的挂钩代码。下一个挂钩过程使用此代码来确定如何处理挂钩信息.
//
// wParam: 传递给当前挂钩过程的 wParam 值。此参数的含义取决于与当前挂钩链关联的挂钩类型.
//
// lParam: 传递给当前挂钩过程的 lParam 值。此参数的含义取决于与当前挂钩链关联的挂钩类型.
//
// 返回值: 此值由链中的下一个挂钩过程返回。当前挂钩过程还必须返回此值。返回值的含义取决于挂钩类型。有关详细信息，请参阅各个挂钩过程的说明.
func CallNextHookEx(hhk uintptr, nCode int32, wParam, lParam uintptr) uintptr {
	r, _, _ := callNextHookEx.Call(hhk, uintptr(nCode), wParam, lParam)
	return r
}

// UnhookWindowsHookEx 删除 SetWindowsHookExW 函数安装在挂钩链中的挂钩过程.
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-unhookWindowsHookEx.
//
// hhk: 要移除的挂钩的句柄。此参数是由先前调用 SetWindowsHookExW 获取的挂钩句柄.
func UnhookWindowsHookEx(hhk uintptr) bool {
	r, _, _ := unhookWindowsHookEx.Call(hhk)
	return r != 0
}

// DestroyIcon 销毁图标并释放图标占用的任何内存.
//   - 只需为使用以下函数创建的图标和游标调用 DestroyIcon ： CreateIconFromResourceEx (如果调用时没有 LR_SHARED 标志) 、 CreateIconIndirect 和 CopyIcon。 请勿使用此函数销毁共享图标。 只要从中加载共享图标的模块保留在内存中，共享图标就有效.
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-destroyicon.
//
// hIcon: 要销毁的图标的句柄。 图标不得处于使用中.
func DestroyIcon(hIcon uintptr) bool {
	r, _, _ := destroyIcon.Call(hIcon)
	return r != 0
}

// CreateIconFromResource 从描述图标的资源位创建图标或光标。若要指定所需的高度或宽度，请使用 CreateIconFromResourceEx 函数.
//   - 使用完图标后，请使用 DestroyIcon 函数销毁它.
//
// 详见: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-createiconfromresource.
//
// presbits: 包含图标或游标资源位的 DWORD 对齐缓冲区指针。通常通过调用 LookupIconIdFromDirectory、 LookupIconIdFromDirectoryEx 和 LoadResource 函数来加载这些位.
//
// dwResSize: 参数指向的位集的大小（以字节为单位）.
//
// fIcon: 指示是要创建图标还是游标。 如果此参数为 TRUE，则创建图标。 如果为 FALSE，则创建游标。LOCALHEADER 结构定义游标热点，是从游标资源位读取的第一个数据.
//
// dwVer: presbits 参数指向的资源位的图标或光标格式的版本号。 该值必须大于或等于 0x00020000 且小于或等于 0x00030000。 此参数通常设置为 0x00030000.
//
// 返回值: HICON句柄.
func CreateIconFromResource(presbits uintptr, dwResSize uint32, fIcon bool, dwVer uint32) (uintptr, error) {
	r, _, err := createIconFromResource.Call(presbits, uintptr(dwResSize), common.BoolPtr(fIcon), uintptr(dwVer))
	return r, err
}

// 系统预定义的标准光标标识符, https://learn.microsoft.com/zh-cn/windows/win32/menurc/about-cursors.
const (
	IDC_ARROW       = 32512 // 标准箭头光标
	IDC_IBEAM       = 32513 // 文本输入光标（I 型光标）
	IDC_WAIT        = 32514 // 等待光标（沙漏）
	IDC_CROSS       = 32515 // 十字光标
	IDC_UPARROW     = 32516 // 垂直箭头光标
	IDC_SIZENWSE    = 32642 // 双箭头光标, 对角线调整大小 1（西北-东南）
	IDC_SIZENESW    = 32643 // 双箭头光标, 对角线调整大小 2（东北-西南）
	IDC_SIZEWE      = 32644 // 双箭头光标, 水平调整大小（水平）
	IDC_SIZENS      = 32645 // 双箭头光标, 垂直调整大小（垂直）
	IDC_SIZEALL     = 32646 // 四向箭头光标, 移动
	IDC_NO          = 32648 // 禁止光标（圆圈斜杠）
	IDC_HAND        = 32649 // 手形光标
	IDC_APPSTARTING = 32650 // 应用程序启动光标（箭头+沙漏）
	IDC_HELP        = 32651 // 帮助光标（箭头+问号）
	IDC_PIN         = 32671 // 位置选择（表示固定或定位）
	IDC_PERSON      = 32672 // 人员选择（表示用户或联系人）
)

type IMAGE_ uint32

const (
	IMAGE_BITMAP IMAGE_ = 0 // 加载位图
	IMAGE_ICON   IMAGE_ = 1 // 加载图标
	IMAGE_CURSOR IMAGE_ = 2 // 加载游标
)

type LR_ uint32

const (
	LR_CREATEDIBSECTION LR_ = 0x00002000 // 当 uType 参数指定 IMAGE_BITMAP 时，会导致函数返回 DIB 节位图而不是兼容的位图。 此标志可用于加载位图而不将其映射到显示设备的颜色.
	LR_DEFAULTCOLOR     LR_ = 0          // 默认标志;它不执行任何工作。 它的意思是“不 LR_MONOCHROME ”.
	LR_DEFAULTSIZE      LR_ = 0x00000040 // 如果 cxDesired 或 cyDesired 值设置为零，则使用游标或图标的系统指标值指定的宽度或高度。 如果未指定此标志，并且 cxDesired 和 cyDesired 设置为零，则函数将使用实际资源大小。 如果资源包含多个图像，则 函数使用第一个图像的大小.
	LR_LOADFROMFILE     LR_ = 0x00000010 // 从 名称 (图标、光标或位图文件指定的文件) 加载独立图像.

	// 在颜色表中搜索图像，并将以下灰色底纹替换为相应的三维颜色
	//	- Dk 灰色，RGB (128，128，128) 与 COLOR_3DSHADOW
	//	- 灰色，RGB (192，192，192) ，带 COLOR_3DFACE
	//	- Lt Gray，RGB (223，223，223) 与 COLOR_3DLIGHT
	LR_LOADMAP3DCOLORS LR_ = 0x00001000

	// 检索图像中第一个像素的颜色值，并将颜色表中的相应条目替换为默认窗口颜色 (COLOR_WINDOW) 。 图像中使用该条目的所有像素都将成为默认的窗口颜色。 此值仅适用于具有相应颜色表的图像.
	//
	// 如果要加载颜色深度大于 8bpp 的位图，请不要使用此选项.
	//
	// 如果 fuLoad 同时包含 LR_LOADTRANSPARENT 值和 LR_LOADMAP3DCOLORS 值， LR_LOADTRANSPARENT 优先。 但是，颜色表条目将替换为 COLOR_3DFACE 而不是 COLOR_WINDOW.
	LR_LOADTRANSPARENT LR_ = 0x00000020
	LR_MONOCHROME      LR_ = 0x00000001 // 加载黑白图像.

	// 如果多次加载映像，则共享映像句柄。 如果未设置 LR_SHARED ，则对同一资源的第二次 LoadImageW 调用将再次加载映像并返回不同的句柄.
	//
	// 使用此标志时，系统将在不再需要资源时销毁资源.
	//
	// 对于非标准大小、加载后可能会更改或从文件加载的图像，请勿使用 LR_SHARED .
	//
	// 加载系统图标或光标时，必须使用 LR_SHARED 否则函数将无法加载资源.
	//
	// 无论请求的大小如何，此函数都会查找缓存中具有请求的资源名称的第一个映像.
	LR_SHARED   LR_ = 0x00008000
	LR_VGACOLOR LR_ = 0x00000080 // 使用真正的 VGA 颜色.
)

// LoadImageW 加载图标、光标、动画光标或位图.
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-loadimagew.
//
// hInst: 包含要加载的图像的 DLL 或可执行文件 (.exe) 模块的句柄。 有关详细信息，请参阅 GetModuleHandle。若要 (图标、光标或位图文件) 加载预定义图像或独立资源，请将此参数设置为0.
//
// pName: 要加载的图像。如果 hInst 参数为非0且 fuLoad 参数省略 LR_LOADFROMFILE， 则 pName 指定 hInst 模块中的图像资源。如果要按名称从模块加载图像资源， 则 pName 参数是指向包含映像资源名称的字符串. 如果要传字符串请用: common.StrPtr().
//
// Type: 要加载的图像的类型。 wapi.IMAGE_ .
//
// cx: 图标或光标的宽度（以像素为单位）。 如果此参数为零且 fuLoad 参数 为 LR_DEFAULTSIZE，则函数使用 SM_CXICON 或 SM_CXCURSOR 系统指标值来设置宽度。 如果此参数为零且未使用 LR_DEFAULTSIZE ，则函数使用实际资源宽度.
//
// cy: 图标或光标的高度（以像素为单位）。 如果此参数为零且 fuLoad 参数 为 LR_DEFAULTSIZE，则函数使用 SM_CYICON 或 SM_CYCURSOR 系统指标值来设置高度。 如果此参数为零且未使用 LR_DEFAULTSIZE ，则函数使用实际资源高度.
//
// fuLoad: 此参数可使用以下一个或多个值: wapi.LR_ .
//
// 返回值: 返回 HICON 句柄.
func LoadImageW(hInst uintptr, pName uintptr, Type IMAGE_, cx, cy int32, fuLoad LR_) uintptr {
	r, _, _ := loadImageW.Call(hInst, pName, uintptr(Type), uintptr(cx), uintptr(cy), uintptr(fuLoad))
	return r
}

const (
	NullStr  = "\x00"
	NullStr2 = NullStr + NullStr // 2个 NullStr
)

// FindWindowW 检索顶级窗口的句柄，该窗口的类名称和窗口名称与指定的字符串匹配。 此函数不搜索子窗口。 此函数不执行区分大小写的搜索.
//   - 如果 lpWindowName 参数不为空， FindWindowW 将调用 GetWindowTextW 函数以检索窗口名称进行比较。 有关可能出现的潜在问题的说明，请参阅 GetWindowTextW 的备注.
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-FindWindowW.
//
// lpClassName: 窗口类名, 可为空.
//
// lpWindowName: 窗口名称（窗口的标题）, 可为空.
//
// 返回值: 返回窗口句柄.
func FindWindowW(lpClassName, lpWindowName string) uintptr {
	r, _, _ := findWindowW.Call(common.StrPtr(lpClassName), common.StrPtr(lpWindowName))
	return r
}

// RegisterWindowMessageW 定义保证在整个系统中唯一的新窗口消息。 发送或发布消息时可以使用消息值.
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-RegisterWindowMessageW.
//
// lpString: 要注册的消息.
//
// 返回值: 如果成功注册消息，则返回值是范围0xC000到0xFFFF的消息标识符. 如果函数失败，则返回值为零.
func RegisterWindowMessageW(lpString string) int {
	r, _, _ := registerWindowMessageW.Call(common.StrPtr(lpString))
	return int(r)
}

// IsWindow 判断一个窗口句柄是否有效.
//   - 线程不应将 IsWindow 用于未创建的窗口，因为调用此函数后可能会销毁该窗口。 此外，由于窗口句柄被回收，句柄甚至可以指向其他窗口.
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-iswindow.
//
// hWnd: 要测试的窗口的句柄.
func IsWindow(hWnd uintptr) bool {
	r, _, _ := isWindow.Call(hWnd)
	return r != 0
}

type HWND_ int

const (
	HWND_NOTOPMOST HWND_ = -2 // 将窗口置于所有非顶层窗口之上（即在所有顶层窗口之后）。如果窗口已经是非顶层窗口则该标志不起作用.
	HWND_TOPMOST   HWND_ = -1 // 将窗口置于所有非顶层窗口之上。即使窗口未被激活, 窗口也将保持顶级位置.
	HWND_TOP       HWND_ = 0  // 将窗口置于Z序的顶部.
	HWND_BOTTOM    HWND_ = 1  // 将窗口置于Z序的底部。如果参数hWnd标识了一个顶层窗口，则窗口失去顶级位置，并且被置在所有其他窗口的底部.
)

// SWP_ 是窗口大小和定位的标志.
type SWP_ uint32

const (
	SWP_ASYNCWINDOWPOS SWP_ = 0x4000 // 如果调用线程和拥有窗口的线程连接到不同的输入队列，系统会将请求发布到拥有窗口的线程。这可以防止调用线程在其他线程处理请求时阻塞其执行.
	SWP_DEFERERASE     SWP_ = 0x2000 // 防止生成 WM_SYNCPAINT 消息.
	SWP_DRAWFRAME      SWP_ = 0x0020 // 在窗口周围绘制一个框架（在窗口的类描述中定义）.
	SWP_FRAMECHANGED   SWP_ = 0x0020 // 应用使用 SetWindowLong 函数 设置的新框架样式。向窗口发送WM_NCCALCSIZE 消息，即使窗口大小没有改变。如果未指定此标志，则仅在更改窗口大小时发送 WM_NCCALCSIZE .
	SWP_HIDEWINDOW     SWP_ = 0x0080 // 隐藏窗口.
	SWP_NOACTIVATE     SWP_ = 0x0010 // 不激活窗口。如果未设置此标志，则窗口被激活并移动到最顶层或非最顶层组的顶部（取决于 hWndInsertAfter 参数的设置）.
	SWP_NOCOPYBITS     SWP_ = 0x0100 // 丢弃客户区的全部内容。如果未指定此标志，则在调整窗口大小或重新定位后，将保存客户区的有效内容并将其复制回客户区.
	SWP_NOMOVE         SWP_ = 0x0002 // 保留当前位置（忽略 X 和 Y 参数）.
	SWP_NOOWNERZORDER  SWP_ = 0x0200 // 不改变所有者窗口在 Z 顺序中的位置.
	SWP_NOREDRAW       SWP_ = 0x0008 // 不重绘更改。如果设置了此标志，则不会发生任何类型的重新绘制。这适用于客户区、非客户区（包括标题栏和滚动条）以及由于窗口移动而未覆盖的父窗口的任何部分。设置此标志时，应用程序必须显式地使需要重绘的窗口和父窗口的任何部分无效或重绘.
	SWP_NOREPOSITION   SWP_ = 0x0200 // 与 SWP_NOOWNERZORDER 标志相同.
	SWP_NOSENDCHANGING SWP_ = 0x0400 // 阻止窗口接收 WM_WINDOWPOSCHANGING 消息.
	SWP_NOSIZE         SWP_ = 0x0001 // 保留当前大小（忽略 cx 和 cy 参数）.
	SWP_NOZORDER       SWP_ = 0x0004 // 保留当前 Z 顺序（忽略 hWndInsertAfter 参数）.
	SWP_SHOWWINDOW     SWP_ = 0x0040 // 显示窗口.
)

// SetWindowPos 改变一个子窗口，弹出式窗口或顶层窗口的尺寸，位置和Z序。子窗口，弹出式窗口，及顶层窗口根据它们在屏幕上出现的顺序排序、顶层窗口设置的级别最高，并且被设置为Z序的第一个窗口.
//
// 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-SetWindowPos.
//
// hWnd: 欲定位的窗口句柄.
//
// hWndInsertAfter: 在Z序中位于定位窗口之前的窗口句柄. 此参数必须是窗口句柄或以下值之一: wapi.HWND_.
//
// x: 窗口新的x坐标。如hwnd是一个子窗口，则x用父窗口的客户区坐标表示.
//
// y: 窗口新的y坐标。如hwnd是一个子窗口，则y用父窗口的客户区坐标表示.
//
// cx: 指定新的窗口宽度.
//
// cy: 指定新的窗口高度.
//
// wFlags: 窗口大小和定位的标志. 该参数可以是以下值的组合: wapi.SWP_.
func SetWindowPos(hWnd uintptr, hWndInsertAfter HWND_, x, y, cx, cy int32, wFlags SWP_) bool {
	r, _, _ := setWindowPos.Call(hWnd, uintptr(hWndInsertAfter), uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(wFlags))
	return r != 0
}

// GetDesktopWindow 获取桌面窗口的句柄.
//
// 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-GetDesktopWindow.
func GetDesktopWindow() uintptr {
	r, _, _ := getDesktopWindow.Call()
	return r
}

// MB_ 指示消息框的内容和行为.
type MB_ uint32

// 要指示消息框中显示的按钮，请指定以下值之一.

const (
	MB_AbortRetryIgnore  MB_ = 0x00000002 // 消息框包含三个按钮：失败、重试和忽略.
	MB_CanaelTryContinue MB_ = 0x00000006 // 消息框包含三个按钮：取消、重试、继续。使用此消息框类型而不是 MB_AbortRetryIgnore.
	MB_Help              MB_ = 0x00004000 // 向消息框 添加帮助按钮。当用户单击帮助按钮或按 F1 时，系统会向所有者 发送WM_HELP消息.
	MB_OK                MB_ = 0x00000000 // 消息框包含一个按钮：确认。这是默认设置.
	MB_OkCancel          MB_ = 0x00000001 // 消息框包含两个按钮：确认和取消.
	MB_RetryCancel       MB_ = 0x00000005 // 消息框包含两个按钮：重试和取消.
	MB_YesNo             MB_ = 0x00000004 // 消息框包含两个按钮：是和否.
	MB_YesNoCancel       MB_ = 0x00000003 // 消息框包含三个按钮：是、否和取消.
)

// 要在消息框中显示图标，请指定以下值之一.

const (
	MB_IconExclamation MB_ = 0x00000030 // 消息框中会出现一个感叹号图标.
	MB_IconWarning     MB_ = 0x00000030 // 消息框中会出现一个感叹号图标.
	MB_IconInformation MB_ = 0x00000040 // 一个由圆圈中的小写字母i组成的图标出现在消息框中.
	MB_IconAsterisk    MB_ = 0x00000040 // 一个由圆圈中的小写字母i组成的图标出现在消息框中.
	MB_IconQuestion    MB_ = 0x00000020 // 问号图标出现在消息框中。不再推荐使用问号消息图标，因为它不能清楚地表示特定类型的消息，并且作为问题的消息措辞可能适用于任何消息类型。此外，用户可能会将消息符号问号与帮助信息混淆。因此，请勿在消息框中使用此问号消息符号。系统继续支持它的包含只是为了向后兼容.
	MB_IconStop        MB_ = 0x00000010 // 一个停止标志图标出现在消息框中.
	MB_IconError       MB_ = 0x00000010 // 一个停止标志图标出现在消息框中.
	MB_IconHand        MB_ = 0x00000010 // 一个停止标志图标出现在消息框中.
)

// 要指示默认按钮，请指定以下值之一.

const (
	MB_DefButton1 MB_ = 0x00000000 // 第一个按钮是默认按钮. MB_DefButton1 是默认值, 除非指定了 MB_DefButton2, MB_DefButton3 或 MB_DefButton4.
	MB_DefButton2 MB_ = 0x00000100 // 第二个按钮是默认按钮.
	MB_DefButton3 MB_ = 0x00000200 // 第三个按钮是默认按钮.
	MB_DefButton4 MB_ = 0x00000300 // 第四个按钮是默认按钮.
)

// 要指示对话框的模式，请指定以下值之一.

const (
	MB_ApplModal MB_ = 0x00000000 // 用户必须先响应消息框，然后才能在hWnd参数标识的窗口中继续工作。但是，用户可以移动到其他线程的窗口并在这些窗口中工作。根据应用程序中窗口的层次结构，用户可能能够移动到线程内的其他窗口。消息框父级的所有子窗口都会自动禁用，但弹出窗口不会。如果未指定 MB_SystemModal 或 MB_TaskModal, 则 MB_ApplModal 是默认值.

	MB_SystemModal MB_ = 0x00001000 // 与 MB_ApplModal 相同，只是消息框具有 WS_EX_TOPMOST 样式。使用系统模式消息框来通知用户需要立即注意的严重的、具有潜在破坏性的错误（例如，内存不足）。此标志对用户与除与hWnd关联的窗口之外的窗口进行交互的能力没有影响.

	MB_TaskModal MB_ = 0x00002000 // 与 MB_ApplModal 相同，除了如果hWnd参数为0则禁用所有属于当前线程的顶级窗口。当调用应用程序或库没有可用的窗口句柄但仍需要防止输入到调用线程中的其他窗口而不暂停其他线程时，请使用此标志.
)

// 要指定其他选项，请使用以下一个或多个值.

const (
	MB_Default_Desktop_Only MB_ = 0x00020000 // 与交互式窗口站的桌面相同。有关详细信息，请参阅窗口站。 如果当前输入桌面不是默认桌面，MessageBox不会返回，直到用户切换到默认桌面.

	MB_Right         MB_ = 0x00080000 // 文本右对齐.
	MB_RtlReading    MB_ = 0x00100000 // 在希伯来语和阿拉伯语系统上使用从右到左的阅读顺序显示消息和标题文本.
	MB_SetForeground MB_ = 0x00010000 // 消息框成为前台窗口。在内部，系统为消息框调用 SetForegroundWindow 函数.
	MB_TopMost       MB_ = 0x00040000 // 消息框是使用 WS_EX_TOPMOST 窗口样式创建的.

	MB_Service_Notification MB_ = 0x00200000 // 调用者是通知用户事件的服务。即使没有用户登录到计算机，该功能也会在当前活动桌面上显示一个消息框。终端服务：如果调用线程具有模拟令牌，则该函数将消息框定向到模拟令牌中指定的会话。如果设置了此标志，则hWnd参数必须为0。这是为了使消息框可以出现在与hWnd对应的桌面以外的桌面上。有关使用此标志的安全注意事项的信息，请参阅交互式服务。特别要注意，此标志可以在锁定的桌面上生成交互式内容，因此只能用于非常有限的一组场景，例如资源耗尽.
)

// ID_ 指示 MessageBoxW 的返回值.
type ID_ uint32

const (
	ID_Abort    ID_ = 3  // 失败按钮被单击.
	ID_Cancel   ID_ = 2  // 取消按钮被单击.
	ID_Continue ID_ = 11 // 继续按钮被单击.
	ID_Ignore   ID_ = 5  // 忽略按钮被单击.
	ID_NO       ID_ = 7  // 否按钮被单击.
	ID_OK       ID_ = 1  // 确定按钮被单击.
	ID_Retry    ID_ = 4  // MB_RetryCancel 和 MB_AbortRetryIgnore 里的重试按钮被单击.
	ID_TryAgain ID_ = 10 // MB_CanaelTryContinue 里的重试按钮被单击.
	ID_YES      ID_ = 6  // 是按钮被单击.
)

// MessageBoxW 显示一个模式对话框，其中包含一个系统图标、一组按钮和一条特定于应用程序的简短消息.
//
// 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-MessageBoxW.
//
// hWnd: 要创建的消息框的所有者窗口的句柄。如果此参数为0，则消息框没有所有者窗口.
//
// lpText: 要显示的消息。如果字符串由多行组成，您可以在每行之间使用换行符分隔各行.
//
// lpCaption: 对话框标题。如果此参数为空，则默认标题为Error.
//
// uType: 对话框的内容和行为, 是以下值的组合: wapi.MB_.
//
// 返回值: 如果函数失败，则返回值为0; 成功则返回一个整数，指示用户单击了哪个按钮.
func MessageBoxW(hWnd uintptr, lpText, lpCaption string, uType MB_) ID_ {
	r, _, _ := messageBoxW.Call(hWnd, common.StrPtr(lpText), common.StrPtr(lpCaption), uintptr(uType))
	return ID_(r)
}

// OpenClipboard 打开剪贴板进行检查并防止其他应用程序修改剪贴板内容.
//   - 如果另一个窗口打开了剪贴板，则 OpenClipboard 会失败.
//   - 应用程序应在每次成功调用 OpenClipboard 后调用 CloseClipboard 函数.
//   - 除非调用 EmptyClipboard 函数，否则由hWndNewOwner参数标识的窗口不会成为剪贴板所有者.
//   - 如果应用程序在 hwnd 设置为0的情况下调用 OpenClipboard, EmptyClipboard 会将剪贴板所有者设置为NULL；这会导致 SetClipboardData 失败.
//
// 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-OpenClipboard.
//
// hWnd: 要与打开的剪贴板关联的窗口句柄。如果此参数为0，则打开的剪贴板与当前任务相关联.
func OpenClipboard(hWnd uintptr) bool {
	r, _, _ := openClipboard.Call(hWnd)
	return r != 0
}

// CloseClipboard 关闭剪贴板.
//   - 当窗口完成检查或更改剪贴板时，通过调用 CloseClipboard 关闭剪贴板。这使其他窗口能够访问剪贴板.
//
// 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-CloseClipboard.
func CloseClipboard() bool {
	r, _, _ := closeClipboard.Call()
	return r != 0
}

// EmptyClipboard 清空剪贴板并释放剪贴板中数据的句柄。然后该函数将剪贴板的所有权分配给当前打开剪贴板的窗口.
//   - 在调用 EmptyClipboard 之前，应用程序必须使用 OpenClipboard 函数打开剪贴板.
//   - 如果应用程序在打开剪贴板时指定了0窗口句柄，则 EmptyClipboard 会成功，但会将剪贴板所有者设置为NULL。请注意，这会导致 SetClipboardData 失败.
//
// 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-EmptyClipboard.
func EmptyClipboard() bool {
	r, _, _ := emptyClipboard.Call()
	return r != 0
}

// CF_ 标准剪贴板格式.
type CF_ uint32

const (
	CF_TEXT         CF_ = 1  // 文字格式。每行以回车/换行（CR-LF）组合结束。空字符表示数据的结尾。对ANSI文本使用此格式.
	CF_BITMAP       CF_ = 2  // 位图的句柄（HBITMAP）
	CF_METAFILEPICT CF_ = 3  // 处理由METAFILEPICT结构定义的图元文件图片格式。通过动态数据交换（DDE）传递CF_METAFILEPICT句柄时，负责删除【HMEM】的应用程序也应该释放 CF_METAFILEPICT 句柄引用的元文件.
	CF_SYLK         CF_ = 4  // Microsoft符号链接（SYLK）格式.
	CF_DIF          CF_ = 5  // 软件艺术数据交换格式.
	CF_TIFF         CF_ = 6  // 标记图像文件格式.
	CF_OEMTEXT      CF_ = 7  // 文字格式包含OEM字符集中的字符。每行以回车/换行（CR-LF）组合结束。空字符表示数据的结尾.
	CF_DIB          CF_ = 8  // 一个包含BITMAPINFO结构的内存对象，后跟位图位.
	CF_PALETTE      CF_ = 9  // 处理调色板。每当应用程序将数据放置在依赖于或假定调色板的剪贴板中时，它也应将调色板放在剪贴板上。如果剪贴板包含CF_PALETTE（逻辑调色板）格式的数据，则应用程序应使用 SelectPalette 和 RealizePalette 函数来实现（比较）剪贴板中与该逻辑调色板的任何其他数据。当显示剪贴板数据时，Windows剪贴板始终将剪贴板上的任何对象用作CF_PALETTE格式的当前调色板.
	CF_PENDATA      CF_ = 10 // 用于Pen Computing的Microsoft Windows笔的扩展数据.
	CF_RIFF         CF_ = 11 // 表示音频数据比CF_WAVE标准波形格式更复杂.
	CF_WAVE         CF_ = 12 // 以诸如11 kHz或22 kHz脉冲编码调制（PCM）的标准波形格式之一表示音频数据.
	CF_UNICODETEXT  CF_ = 13 // 仅Windows NT： Unicode文字格式。每行以回车/换行（CR-LF）组合结束。空字符表示数据的结尾.
	CF_ENHMETAFILE  CF_ = 14 // 增强图元文件的句柄（HENHMETAFILE）.
	CF_HDROP        CF_ = 15 // 类型为HDROP的句柄，用于标识文件列表。应用程序可以通过将句柄传递给DragQueryFile函数来检索有关文件的信息.
)

// IsClipboardFormatAvailable 确定剪贴板是否包含指定格式的数据.
//
// 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-IsClipboardFormatAvailable.
//
// uFormat: 标准或注册的剪贴板格式, wapi.CF_ .
func IsClipboardFormatAvailable(uFormat CF_) bool {
	r, _, _ := isClipboardFormatAvailable.Call(uintptr(uFormat))
	return r != 0
}

// GetClipboardData 从剪贴板中检索指定格式的数据。剪贴板必须先前已打开.
//
// 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-GetClipboardData.
//
// uFormat: 剪贴板格式, wapi.CF_ .
//
// 返回值: 如果函数成功，则返回值是指定格式的剪贴板对象的句柄. 如果函数失败，则返回值为0.
func GetClipboardData(uFormat CF_) uintptr {
	r, _, _ := getClipboardData.Call(uintptr(uFormat))
	return r
}

// SetClipboardData 以指定的剪贴板格式将数据放在剪贴板上。该窗口必须是当前剪贴板所有者，并且应用程序必须调用 OpenClipboard 函数.
//
// 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-setclipboarddata.
//
// uFormat: 标准或注册的剪贴板格式, wapi.CF_ .
//
// hMem: 指定格式的数据的句柄。该参数可以为0，表示窗口根据请求提供指定剪贴板格式的数据（渲染格式）.
//
// 返回值: 如果函数成功，则返回值是数据的句柄. 如果函数失败，则返回值为0.
func SetClipboardData(uFormat CF_, hMem uintptr) uintptr {
	r, _, _ := setClipboardData.Call(uintptr(uFormat), hMem)
	return r
}

// SetForegroundWindow 将创建指定窗口的线程带到前台并激活窗口. 键盘输入被定向到窗口, 并且为用户改变了各种视觉提示. 系统为创建前台窗口的线程分配比其他线程稍高的优先级.
//
// 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-SetForegroundWindow.
//
// hWnd: 应激活并置于前台的窗口句柄.
func SetForegroundWindow(hWnd uintptr) bool {
	r, _, _ := setForegroundWindow.Call(hWnd)
	return r != 0
}

// FindWindowExW 检索类名称和窗口名称与指定字符串匹配的窗口的句柄. 该函数搜索子窗口，从指定子窗口后面的那个开始. 此函数不执行区分大小写的搜索.
//
// 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-FindWindowExW.
//
// hWndParent: 要搜索其子窗口的父窗口的句柄. 如果hwndParent为0，该函数使用桌面窗口作为父窗口. 该函数在作为桌面子窗口的窗口中进行搜索.
//
// hWndChildAfter: 子窗口的句柄。搜索从 Z 顺序中的下一个子窗口开始。子窗口必须是hwndParent的直接子窗口，而不仅仅是后代窗口。 如果hwndChildAfter为0，则搜索从hwndParent的第一个子窗口开始。 请注意，如果hwndParent和hwndChildAfter都是0，则该函数将搜索所有顶级和仅消息窗口.
//
// lpszClass: 窗口类名, 可空.
//
// lpszWindow: 窗口名称（窗口的标题）, 可空.
func FindWindowExW(hWndParent, hWndChildAfter uintptr, lpszClass, lpszWindow string) uintptr {
	r, _, _ := findWindowExW.Call(hWndParent, hWndChildAfter, common.StrPtr(lpszClass), common.StrPtr(lpszWindow))
	return r
}

// GetWindowTextLengthW 检索指定窗口标题栏文本的长度（以字符为单位）（如果窗口有标题栏）。如果指定的窗口是控件，则该函数检索控件内文本的长度。但是无法检索另一个应用程序中编辑控件的文本长度.
//
// 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-GetWindowTextLengthW.
//
// hWnd: 窗口或控件的句柄.
//
// 返回值: 如果成功，则返回值是文本的长度（以字符为单位）。在某些情况下，此值可能大于文本的长度。如果窗口没有文本，则返回值为零.
func GetWindowTextLengthW(hWnd uintptr) int {
	r, _, _ := getWindowTextLengthW.Call(hWnd)
	return int(r)
}

// GetWindowTextW 将指定窗口标题栏（如果有）的文本复制到缓冲区中。如果指定的窗口是控件，则复制控件的文本。但是无法检索另一个应用程序中控件的文本.
//
// 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-GetWindowTextW.
//
// hWnd: 包含文本的窗口或控件的句柄.
//
// lpString: 接收文本.
//
// nMaxCount: 复制到缓冲区的最大字符数，包括空字符。如果文本超出此限制，则将其截断.
//
// 返回值:
//   - 如果函数成功，则返回值是复制字符串的长度（以字符为单位），不包括终止空字符。
//   - 如果窗口没有标题栏或文本，如果标题栏为空，或者窗口或控制句柄无效，则返回值为零.
func GetWindowTextW(hWnd uintptr, lpString *string, nMaxCount int) int {
	buf := make([]uint16, nMaxCount)
	r, _, _ := getWindowTextW.Call(hWnd, common.Uint16SliceDataPtr(&buf), uintptr(nMaxCount))
	*lpString = syscall.UTF16ToString(buf[:])
	return int(r)
}

// ClientToScreen 将指定点的客户区坐标转换为屏幕坐标.
//
// 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-ClientToScreen.
//
// hWnd: 窗口真实句柄.
//
// lpPoint: wapi.POINT 指针. 如果函数成功，则将新的屏幕坐标复制到此结构中.
func ClientToScreen(hWnd uintptr, lpPoint *POINT) bool {
	r, _, _ := clientToScreen.Call(hWnd, uintptr(unsafe.Pointer(lpPoint)))
	return r != 0
}

// GetCursorPos 检索鼠标光标的位置，以屏幕坐标表示.
//
// 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-getcursorpos.
//
// lpPoint: 指向接收光标屏幕坐标的 wapi.POINT 结构的指针.
func GetCursorPos(lpPoint *POINT) bool {
	r, _, _ := getCursorPos.Call(uintptr(unsafe.Pointer(lpPoint)))
	return r != 0
}

// type Mod_ uint32

const (
	Mod_Alt      uint32 = 0x0001 // 必须按住任一 ALT 键.
	Mod_Control  uint32 = 0x0002 // 必须按住任一 CTRL 键.
	Mod_Norepeat uint32 = 0x4000 // 更改热键行为，以便键盘自动重复不会产生多个热键通知。Windows Vista：  不支持此标志.
	Mod_Shift    uint32 = 0x0004 // 必须按住任一 SHIFT 键.
	Mod_Win      uint32 = 0x0008 // 任一 WINDOWS 键被按住。这些键标有 Windows 徽标。涉及 WINDOWS 键的键盘快捷键保留供操作系统使用.
)

// RegisterHotKey 注册系统范围的热键.
//
// 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-registerhotkey.
//
// hWnd: 真实窗口句柄。将接收由热键生成的 WM_HOTKEY 消息的窗口句柄。如果此参数为0，则 WM_HOTKEY 消息将发布到调用线程的消息队列中，并且必须在消息循环中进行处理.
//
// id: 热键的标识符。如果hWnd参数为0，则热键与当前线程相关联，而不是与特定窗口相关联。如果已存在具有相同hWnd和id参数的热键，请参阅备注了解所采取的操作.
//
// fsModifiers: 为了生成 WM_HOTKEY 消息，必须与vk参数指定的键组合按下的键 。fsModifiers参数可以是以下值的组合: wapi.Mod_ .
//
// vk: 热键的虚拟键代码: xcc.VK_ . 请参阅虚拟键码: https://docs.microsoft.com/zh-cn/windows/win32/inputdev/virtual-key-codes.
func RegisterHotKey(hWnd uintptr, id int32, fsModifiers, vk uint32) bool {
	r, _, _ := registerHotKey.Call(hWnd, uintptr(id), uintptr(fsModifiers), uintptr(vk))
	return r != 0
}

// UnregisterHotKey 释放先前注册的热键.
//
// 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-unregisterHotKey.
//
// hWnd: 真实窗口句柄。与要释放的热键关联的窗口句柄。如果热键与窗口无关，则此参数应为0.
//
// id: 要释放的热键的标识符.
func UnregisterHotKey(hWnd uintptr, id int32) bool {
	r, _, _ := unregisterHotKey.Call(hWnd, uintptr(id))
	return r != 0
}

// GetMessage 从调用线程的消息队列中检索消息。应用程序通常使用返回值来确定是否结束主消息循环并退出程序。该函数分派传入的已发送消息，直到发布的消息可用于检索。 与 GetMessage 不同， PeekMessage 函数在返回之前不会等待消息发布.
//
// pMsg: 指向从线程的消息队列接收消息信息的 MSG 结构的指针.
//
// hWnd: 要检索其消息的窗口的句柄。窗口必须属于当前线程。
//   - 如果 hWnd 为0， GetMessage 检索属于当前线程的任何窗口的消息，以及当前线程的消息队列中 hwnd 值为0的任何消息（参见 MSG 结构）。因此，如果 hWnd 为0，则同时处理窗口消息和线程消息。
//   - 如果 hWnd 为-1， GetMessage 仅检索当前线程的消息队列中 hwnd 值为0的消息，即 PostMessageW （当 hWnd 参数为0时）或 PostThreadMessage 发布的线程消息.
//
// wMsgFilterMin: 要检索的最低消息值的整数值。使用 WM_KEYFIRST (0x0100) 指定第一条键盘消息或 WM_MOUSEFIRST (0x0200) 指定第一条鼠标消息.
//
// wMsgFilterMax: 要检索的最高消息值的整数值。使用 WM_KEYLAST 指定最后一个键盘消息或 WM_MOUSELAST 指定最后一个鼠标消息.
//
// 返回值: 如果函数检索到 WM_QUIT 以外的消息，则返回值非0。如果函数检索到 WM_QUIT 消息，则返回值为0。如果有错误，返回值为-1.
//
// 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-GetMessageW.
func GetMessage(pMsg *MSG, hWnd uintptr, wMsgFilterMin uint32, wMsgFilterMax uint32) int32 {
	r, _, _ := getMessageW.Call(uintptr(unsafe.Pointer(pMsg)), hWnd, uintptr(wMsgFilterMin), uintptr(wMsgFilterMax))
	return int32(r)
}

// TranslateMessage 将虚拟键消息转换为字符消息。字符消息被发布到调用线程的消息队列中，以便在线程下次调用 GetMessage 或 PeekMessage 函数时读取.
//
// 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-TranslateMessage.
//
// pMsg: 一个指向 MSG 结构的指针，该结构包含使用 GetMessage 或 PeekMessage 函数从调用线程的消息队列中检索到的消息信息.
func TranslateMessage(pMsg *MSG) bool {
	r, _, _ := translateMessage.Call(uintptr(unsafe.Pointer(pMsg)))
	return r != 0
}

// DispatchMessage 向窗口过程发送消息。它通常用于发送由 GetMessage 函数检索到的消息.
//
// 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-DispatchMessageW.
//
// pMsg: 指向包含消息的结构的指针.
//
// 返回值: 返回值指定窗口过程返回的值。尽管它的含义取决于所发送的消息，但返回值通常会被忽略.
func DispatchMessage(pMsg *MSG) int {
	r, _, _ := dispatchMessageW.Call(uintptr(unsafe.Pointer(pMsg)))
	return int(r)
}

// PostQuitMessage 向系统指示线程已请求终止（退出）。它通常用于响应 WM_DESTROY 消息.
//
// 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-PostQuitMessage.
//
// nExitCode: 应用程序退出代码。该值用作 WM_QUIT 消息的wParam参数.
func PostQuitMessage(nExitCode int32) {
	postQuitMessage.Call(uintptr(nExitCode))
}

type MSG struct {
	Hwnd    uintptr
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      POINT
}

type POINT struct {
	X int32
	Y int32
}

// SendMessageW 将指定的消息发送到一个或多个窗口。SendMessage函数调用指定窗口的窗口过程，直到窗口过程处理完消息才返回.
//
// 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-SendMessageW.
//
// hWnd: 窗口句柄，其窗口过程将接收消息。如果该参数为 HWND_BROADCAST ((HWND)0xffff)，则将消息发送到系统中的所有顶层窗口，包括禁用或不可见的无主窗口、重叠窗口和弹出窗口；但消息不会发送到子窗口.
//
// Msg: 要发送的消息。有关系统提供的消息的列表，请参阅: https://docs.microsoft.com/en-us/windows/win32/winmsg/about-messages-and-message-queues.
//
// wParam: 其他特定于消息的信息.
//
// lParam: 其他特定于消息的信息.
//
// 返回值: 返回值指定消息处理的结果；这取决于发送的消息.
func SendMessageW(hWnd uintptr, Msg uint32, wParam, lParam uintptr) int {
	r, _, _ := sendMessageW.Call(hWnd, uintptr(Msg), wParam, lParam)
	return int(r)
}

// PostMessageW 在与创建指定窗口的线程关联的消息队列中放置（发布）一条消息，并在不等待线程处理消息的情况下返回.
//
// 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-postmessagew.
//
// hWnd: 窗口句柄，其窗口过程将接收消息。如果该参数为 HWND_BROADCAST ((HWND)0xffff)，则将消息发送到系统中的所有顶层窗口，包括禁用或不可见的无主窗口、重叠窗口和弹出窗口；但消息不会发送到子窗口.
//
// Msg: 要发送的消息。有关系统提供的消息的列表，请参阅: https://docs.microsoft.com/en-us/windows/win32/winmsg/about-messages-and-message-queues.
//
// wParam: 其他特定于消息的信息.
//
// lParam: 其他特定于消息的信息.
func PostMessageW(hWnd uintptr, Msg uint32, wParam, lParam uintptr) bool {
	r, _, _ := postMessageW.Call(hWnd, uintptr(Msg), wParam, lParam)
	return r != 0
}

const (
	WM_PAINT            uint32 = 15     // 窗口绘制消息
	WM_CLOSE            uint32 = 16     // 窗口关闭消息.
	WM_DESTROY          uint32 = 2      // 窗口销毁消息.
	WM_NCDESTROY        uint32 = 130    // 窗口非客户区销毁消息, 在 WM_DESTROY 之后.
	WM_MOUSEMOVE        uint32 = 512    // 窗口鼠标移动消息.
	WM_LBUTTONDOWN      uint32 = 513    // 窗口鼠标左键按下消息
	WM_LBUTTONUP        uint32 = 514    // 窗口鼠标左键弹起消息.
	WM_RBUTTONDOWN      uint32 = 516    // 窗口鼠标右键按下消息.
	WM_RBUTTONUP        uint32 = 517    // 窗口鼠标右键弹起消息.
	WM_LBUTTONDBLCLK    uint32 = 515    // 窗口鼠标左键双击消息.
	WM_RBUTTONDBLCLK    uint32 = 518    // 窗口鼠标右键双击消息.
	WM_MBUTTONDOWN      uint32 = 519    // 窗口鼠标中键按下消息.
	WM_MBUTTONUP        uint32 = 520    // 窗口鼠标中键弹起消息.
	WM_MOUSEWHEEL       uint32 = 522    // 窗口鼠标滚轮滚动消息.
	WM_XBUTTONDOWN      uint32 = 523    // 鼠标按下第一个或第二个 X 按钮.
	WM_XBUTTONUP        uint32 = 524    // 鼠标弹起第一个或第二个 X 按钮.
	WM_XBUTTONDBLCLK    uint32 = 525    // 鼠标双击第一个或第二个 X 按钮.
	WM_NCXBUTTONDOWN    uint32 = 171    // 鼠标按下第一个或第二个 X 按钮.
	WM_NCXBUTTONUP      uint32 = 172    // 鼠标弹起第一个或第二个 X 按钮.
	WM_NCXBUTTONDBLCLK  uint32 = 173    // 鼠标双击第一个或第二个 X 按钮.
	WM_EXITSIZEMOVE     uint32 = 562    // 窗口退出移动或调整大小模式循环改，详情参见MSDN.
	WM_MOUSEHOVER       uint32 = 673    // 窗口鼠标进入消息
	WM_MOUSELEAVE       uint32 = 675    // 窗口鼠标离开消息.
	WM_SIZE             uint32 = 5      // 窗口大小改变消息.
	WM_TIMER            uint32 = 275    // 窗口定时器消息.
	WM_SETFOCUS         uint32 = 7      // 窗口获得焦点.
	WM_KILLFOCUS        uint32 = 8      // 窗口失去焦点.
	WM_KEYDOWN          uint32 = 256    // 窗口键盘按键消息.
	WM_KEYUP            uint32 = 257    // 窗口键盘按键弹起消息.
	WM_SYSKEYDOWN       uint32 = 260    // 当用户按下F10键（激活菜单栏）或按住ALT键然后按下另一个键时，发布到具有键盘焦点的窗口。当当前没有窗口具有键盘焦点时，也会发生这种情况;在这种情况下， WM_SYSKEYDOWN 消息被发送到活动窗口。接收消息的窗口可以通过检查lParam参数中的上下文代码来区分这两个上下文.
	WM_SYSKEYUP         uint32 = 261    // 当用户释放按住 Alt 键时按下的键时，使用键盘焦点发布到窗口。 当当前没有窗口具有键盘焦点时，也会发生这种情况;在这种情况下， WM_SYSKEYUP 消息将发送到活动窗口。 接收消息的窗口可以通过检查 lParam 中的上下文代码来区分这两个上下文.
	WM_CAPTURECHANGED   uint32 = 533    // 窗口鼠标捕获改变消息.
	WM_SETCURSOR        uint32 = 32     // 窗口设置鼠标光标.
	WM_CHAR             uint32 = 258    // 窗口字符消息.
	WM_DROPFILES        uint32 = 563    // 拖动文件到窗口.
	WM_HOTKEY           uint32 = 0x0312 // 当用户按下 RegisterHotKey 函数注册的热键时发布。消息放置在与注册热键的线程关联的消息队列的顶部.
	WM_SETICON          uint32 = 0x0080 // 设置窗口图标的消息
	WM_ACTIVATE         uint32 = 6      // 窗口激活消息, 发送到正在激活的窗口和正在停用的窗口.
	WM_MOVE             uint32 = 3      // 窗口移动消息.
	WM_MOVING           uint32 = 534    // 当用户正在移动窗口时，该消息会被发送到窗口.
	WM_GETMINMAXINFO    uint32 = 36     // 获取窗口最小最大尺寸信息.
	WM_QUIT             uint32 = 0x0012 // 应用程序退出消息.
	WM_KEYFIRST         uint32 = 0x0100 // 按键消息的第一个值.
	WM_MOUSEFIRST       uint32 = 0x0200 // 鼠标消息的第一个值.
	WM_KEYLAST          uint32 = 0x0109 // 按键消息的最后一个值.
	WM_MOUSELAST        uint32 = 0x020E // 鼠标消息的最后一个值.
	WM_SYSCOMMAND       uint32 = 0x0112 // 系统命令消息.
	WM_WINDOWPOSCHANGED uint32 = 0x0047 // 窗口位置改变消息.
)

// 系统菜单命令值常量
const (
	SC_SIZE               = 0xF000 // 调整窗口大小
	SC_MOVE               = 0xF010 // 移动窗口
	SC_MINIMIZE           = 0xF020 // 最小化窗口
	SC_MAXIMIZE           = 0xF030 // 最大化窗口
	SC_NEXTWINDOW         = 0xF040 // 切换到下一个窗口
	SC_PREVWINDOW         = 0xF050 // 切换到上一个窗口
	SC_CLOSE              = 0xF060 // 关闭窗口
	SC_VSCROLL            = 0xF070 // 垂直滚动
	SC_HSCROLL            = 0xF080 // 水平滚动
	SC_MOUSEMENU          = 0xF090 // 鼠标菜单
	SC_KEYMENU            = 0xF100 // 键盘菜单
	SC_ARRANGE            = 0xF110 // 排列窗口
	SC_RESTORE            = 0xF120 // 恢复窗口
	SC_TASKLIST           = 0xF130 // 任务列表
	SC_SCREENSAVE         = 0xF140 // 屏幕保护程序
	SC_HOTKEY             = 0xF150 // 热键
	SC_DEFAULT            = 0xF160 // 默认操作
	SC_MONITORPOWERuint32 = 0xF170 // 显示器电源设置
	SC_CONTEXTHELP        = 0xF180 // 上下文帮助
	SC_SEPARATOR          = 0xF00F // 分隔符
)

const (
	ICON_BIG   = 1 // 大图标
	ICON_SMALL = 0 // 小图标
)

// WM_SIZE 消息的 wParam 值
const (
	SIZE_RESTORED  = iota // 0: 窗口恢复（正常大小）
	SIZE_MINIMIZED        // 1: 窗口最小化（任务栏）
	SIZE_MAXIMIZED        // 2: 窗口最大化（全屏）
	SIZE_MAXSHOW          // 3: 其他窗口最大化时触发显示（如任务栏自动隐藏）
	SIZE_MAXHIDE          // 4: 其他窗口最大化时触发隐藏（如任务栏自动隐藏）
)

// WM_ACTIVATE 消息的窗口激活状态值（对应 wParam）
const (
	WA_INACTIVE    = 0 // 窗口被停用（非活动状态）
	WA_ACTIVE      = 1 // 窗口被激活（例如通过键盘切换）
	WA_CLICKACTIVE = 2 // 窗口通过鼠标点击激活
)

// WM_KEYUP, WM_KEYDOWN, WM_CHAR 消息的 lParam 高位标志（HIWORD）
const (
	KF_EXTENDED = 0x0100 // 扩展键（如右侧Alt/Ctrl/方向键）
	KF_DLGMODE  = 0x0800 // 对话框激活状态下的按键
	KF_MENUMODE = 0x1000 // 菜单激活状态下的按键
	KF_ALTDOWN  = 0x2000 // Alt键被按住
	KF_REPEAT   = 0x4000 // 重复按键计数标志
	KF_UP       = 0x8000 // 按键释放事件（用于 WM_KEYUP）
)

// 低级键盘钩子标志 (KBDLLHOOKSTRUCT.flags)
const (
	LLKHF_EXTENDED          = 0x0001 // 扩展键标志 (如右Alt/Ctrl)
	LLKHF_INJECTED          = 0x0010 // 事件由其他进程注入
	LLKHF_ALTDOWN           = 0x0020 // Alt键按下状态
	LLKHF_UP                = 0x0080 // 按键释放事件
	LLKHF_LOWER_IL_INJECTED = 0x0002 // 低完整性进程注入事件
)

// PT_ 定义了指针输入的类型
type PT_ uint32

const (
	// PT_POINTER 普通指针输入
	PT_POINTER PT_ = iota + 1
	// PT_TOUCH 触摸输入
	PT_TOUCH
	// PT_PEN 笔输入
	PT_PEN
	// PT_MOUSE 鼠标输入
	PT_MOUSE
	// PT_TOUCHPAD 触摸板输入
	PT_TOUCHPAD
)

// POINTER_FLAG_ 指针输入标志常量
type POINTER_FLAG_ uint32

const (
	POINTER_FLAG_NONE           POINTER_FLAG_ = 0x00000000 // 默认无标志
	POINTER_FLAG_NEW            POINTER_FLAG_ = 0x00000001 // 新指针
	POINTER_FLAG_INRANGE        POINTER_FLAG_ = 0x00000002 // 指针未离开范围
	POINTER_FLAG_INCONTACT      POINTER_FLAG_ = 0x00000004 // 指针处于接触状态
	POINTER_FLAG_FIRSTBUTTON    POINTER_FLAG_ = 0x00000010 // 主按钮动作
	POINTER_FLAG_SECONDBUTTON   POINTER_FLAG_ = 0x00000020 // 副按钮动作
	POINTER_FLAG_THIRDBUTTON    POINTER_FLAG_ = 0x00000040 // 第三按钮
	POINTER_FLAG_FOURTHBUTTON   POINTER_FLAG_ = 0x00000080 // 第四按钮
	POINTER_FLAG_FIFTHBUTTON    POINTER_FLAG_ = 0x00000100 // 第五按钮
	POINTER_FLAG_PRIMARY        POINTER_FLAG_ = 0x00002000 // 系统主指针
	POINTER_FLAG_CONFIDENCE     POINTER_FLAG_ = 0x00004000 // 非意外触发的指针
	POINTER_FLAG_CANCELED       POINTER_FLAG_ = 0x00008000 // 异常离开的指针
	POINTER_FLAG_DOWN           POINTER_FLAG_ = 0x00010000 // 指针按下状态(接触开始)
	POINTER_FLAG_UPDATE         POINTER_FLAG_ = 0x00020000 // 指针更新
	POINTER_FLAG_UP             POINTER_FLAG_ = 0x00040000 // 指针抬起状态(接触结束)
	POINTER_FLAG_WHEEL          POINTER_FLAG_ = 0x00080000 // 垂直滚轮
	POINTER_FLAG_HWHEEL         POINTER_FLAG_ = 0x00100000 // 水平滚轮
	POINTER_FLAG_CAPTURECHANGED POINTER_FLAG_ = 0x00200000 // 捕获丢失
	POINTER_FLAG_HASTRANSFORM   POINTER_FLAG_ = 0x00400000 // 关联变换矩阵的输入
)

// POINTER_CHANGE_ 定义了指针按钮状态变化的类型
type POINTER_CHANGE_ uint32

const (
	POINTER_CHANGE_NONE              POINTER_CHANGE_ = iota // 无按钮状态变化
	POINTER_CHANGE_FIRSTBUTTON_DOWN                         // 第一个按钮按下
	POINTER_CHANGE_FIRSTBUTTON_UP                           // 第一个按钮释放
	POINTER_CHANGE_SECONDBUTTON_DOWN                        // 第二个按钮按下
	POINTER_CHANGE_SECONDBUTTON_UP                          // 第二个按钮释放
	POINTER_CHANGE_THIRDBUTTON_DOWN                         // 第三个按钮按下
	POINTER_CHANGE_THIRDBUTTON_UP                           // 第三个按钮释放
	POINTER_CHANGE_FOURTHBUTTON_DOWN                        // 第四个按钮按下
	POINTER_CHANGE_FOURTHBUTTON_UP                          // 第四个按钮释放
	POINTER_CHANGE_FIFTHBUTTON_DOWN                         // 第五个按钮按下
	POINTER_CHANGE_FIFTHBUTTON_UP                           // 第五个按钮释放
)

// PEN_FLAG_ 定义笔输入标志常量
type PEN_FLAG_ uint32

const (
	PEN_FLAG_NONE     PEN_FLAG_ = 0x00000000 // 默认无标志
	PEN_FLAG_BARREL   PEN_FLAG_ = 0x00000001 // 笔杆按钮按下
	PEN_FLAG_INVERTED PEN_FLAG_ = 0x00000002 // 笔尖倒置
	PEN_FLAG_ERASER   PEN_FLAG_ = 0x00000004 // 橡皮擦按钮按下
)

// PEN_MASK_ 定义笔输入掩码常量
type PEN_MASK_ uint32

const (
	PEN_MASK_NONE     PEN_MASK_ = 0x00000000 // 默认值 - 所有可选字段均无效
	PEN_MASK_PRESSURE PEN_MASK_ = 0x00000001 // 压力字段有效
	PEN_MASK_ROTATION PEN_MASK_ = 0x00000002 // 旋转字段有效
	PEN_MASK_TILT_X   PEN_MASK_ = 0x00000004 // X轴倾斜字段有效
	PEN_MASK_TILT_Y   PEN_MASK_ = 0x00000008 // Y轴倾斜字段有效
)

// TOUCH_MASK_ 定义触摸输入掩码常量
type TOUCH_MASK_ uint32

const (
	TOUCH_MASK_NONE        TOUCH_MASK_ = 0x00000000 // 默认 - 所有可选字段均无效
	TOUCH_MASK_CONTACTAREA TOUCH_MASK_ = 0x00000001 // rcContact 字段有效
	TOUCH_MASK_ORIENTATION TOUCH_MASK_ = 0x00000002 // orientation 字段有效
	TOUCH_MASK_PRESSURE    TOUCH_MASK_ = 0x00000004 // pressure 字段有效
)

const (
	TOUCH_FLAG_NONE uint32 = 0x00000000 // 默认
)

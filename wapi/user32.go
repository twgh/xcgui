package wapi

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
)

var (
	// Library.
	user32 = syscall.NewLazyDLL("user32.dll")

	// Functions.
	setWindowPos               = user32.NewProc("SetWindowPos")
	getDesktopWindow           = user32.NewProc("GetDesktopWindow")
	messageBoxW                = user32.NewProc("MessageBoxW")
	isClipboardFormatAvailable = user32.NewProc("IsClipboardFormatAvailable")
	openClipboard              = user32.NewProc("OpenClipboard")
	closeClipboard             = user32.NewProc("CloseClipboard")
	emptyClipboard             = user32.NewProc("EmptyClipboard")
	getClipboardData           = user32.NewProc("GetClipboardData")
	setClipboardData           = user32.NewProc("SetClipboardData")
	setForegroundWindow        = user32.NewProc("SetForegroundWindow")
	findWindowExW              = user32.NewProc("FindWindowExW")
	getWindowTextLengthW       = user32.NewProc("GetWindowTextLengthW")
	getWindowTextW             = user32.NewProc("GetWindowTextW")
	clientToScreen             = user32.NewProc("ClientToScreen")
	getCursorPos               = user32.NewProc("GetCursorPos")
	registerHotKey             = user32.NewProc("RegisterHotKey")
	unregisterHotKey           = user32.NewProc("UnregisterHotKey")
	getMessageW                = user32.NewProc("GetMessageW")
	translateMessage           = user32.NewProc("TranslateMessage")
	dispatchMessageW           = user32.NewProc("DispatchMessageW")
	postQuitMessage            = user32.NewProc("PostQuitMessage")
	sendMessageW               = user32.NewProc("SendMessageW")
	postMessageW               = user32.NewProc("PostMessageW")
	isWindow                   = user32.NewProc("IsWindow")
	registerWindowMessageW     = user32.NewProc("RegisterWindowMessageW")
	findWindowW                = user32.NewProc("FindWindowW")
	loadImageW                 = user32.NewProc("LoadImageW")
	createIconFromResource     = user32.NewProc("CreateIconFromResource")
	destroyIcon                = user32.NewProc("DestroyIcon")
)

// DestroyIcon 销毁图标并释放图标占用的任何内存。
//
//	@Description 只需为使用以下函数创建的图标和游标调用 DestroyIcon ： CreateIconFromResourceEx (如果调用时没有 LR_SHARED 标志) 、 CreateIconIndirect 和 CopyIcon。 请勿使用此函数销毁共享图标。 只要从中加载共享图标的模块保留在内存中，共享图标就有效。
//	详见: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-destroyicon.
//	@param hIcon 要销毁的图标的句柄。 图标不得处于使用中。
func DestroyIcon(hIcon uintptr) bool {
	r, _, _ := destroyIcon.Call(hIcon)
	return r != 0
}

// CreateIconFromResource 从描述图标的资源位创建图标或光标。若要指定所需的高度或宽度，请使用 CreateIconFromResourceEx 函数。
//
//	@Description 使用完图标后，请使用 DestroyIcon 函数销毁它。
//
// 详见: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-createiconfromresource.
//
//	@param presbits 包含图标或游标资源位的 DWORD 对齐缓冲区指针。通常通过调用 LookupIconIdFromDirectory、 LookupIconIdFromDirectoryEx 和 LoadResource 函数来加载这些位。
//	@param dwResSize 参数指向的位集的大小（以字节为单位）。
//	@param fIcon 指示是要创建图标还是游标。 如果此参数为 TRUE，则创建图标。 如果为 FALSE，则创建游标。LOCALHEADER 结构定义游标热点，是从游标资源位读取的第一个数据。
//	@param dwVer presbits 参数指向的资源位的图标或光标格式的版本号。 该值必须大于或等于 0x00020000 且小于或等于 0x00030000。 此参数通常设置为 0x00030000。
//	@return HICON
func CreateIconFromResource(presbits uintptr, dwResSize uint32, fIcon bool, dwVer uint32) (uintptr, error) {
	r, _, err := createIconFromResource.Call(presbits, uintptr(dwResSize), common.BoolPtr(fIcon), uintptr(dwVer))
	return r, err
}

type IMAGE_ uint32

const (
	IMAGE_BITMAP = 0 // 加载位图
	IMAGE_ICON   = 1 // 加载图标
	IMAGE_CURSOR = 2 // 加载游标
)

type LR_ uint32

const (
	LR_CREATEDIBSECTION = 0x00002000 // 当 uType 参数指定 IMAGE_BITMAP 时，会导致函数返回 DIB 节位图而不是兼容的位图。 此标志可用于加载位图而不将其映射到显示设备的颜色。
	LR_DEFAULTCOLOR     = 0          // 默认标志;它不执行任何工作。 它的意思是“不 LR_MONOCHROME ”。
	LR_DEFAULTSIZE      = 0x00000040 // 如果 cxDesired 或 cyDesired 值设置为零，则使用游标或图标的系统指标值指定的宽度或高度。 如果未指定此标志，并且 cxDesired 和 cyDesired 设置为零，则函数将使用实际资源大小。 如果资源包含多个图像，则 函数使用第一个图像的大小。
	LR_LOADFROMFILE     = 0x00000010 // 从 名称 (图标、光标或位图文件指定的文件) 加载独立图像。

	// 在颜色表中搜索图像，并将以下灰色底纹替换为相应的三维颜色
	//	- Dk 灰色，RGB (128，128，128) 与 COLOR_3DSHADOW
	//	- 灰色，RGB (192，192，192) ，带 COLOR_3DFACE
	//	- Lt Gray，RGB (223，223，223) 与 COLOR_3DLIGHT
	LR_LOADMAP3DCOLORS = 0x00001000

	// 检索图像中第一个像素的颜色值，并将颜色表中的相应条目替换为默认窗口颜色 (COLOR_WINDOW) 。 图像中使用该条目的所有像素都将成为默认的窗口颜色。 此值仅适用于具有相应颜色表的图像。
	//
	// 如果要加载颜色深度大于 8bpp 的位图，请不要使用此选项。
	//
	// 如果 fuLoad 同时包含 LR_LOADTRANSPARENT 值和 LR_LOADMAP3DCOLORS 值， LR_LOADTRANSPARENT 优先。 但是，颜色表条目将替换为 COLOR_3DFACE 而不是 COLOR_WINDOW。
	LR_LOADTRANSPARENT = 0x00000020
	LR_MONOCHROME      = 0x00000001 // 加载黑白图像。

	// 如果多次加载映像，则共享映像句柄。 如果未设置 LR_SHARED ，则对同一资源的第二次 LoadImageW 调用将再次加载映像并返回不同的句柄。
	//
	// 使用此标志时，系统将在不再需要资源时销毁资源。
	//
	// 对于非标准大小、加载后可能会更改或从文件加载的图像，请勿使用 LR_SHARED 。
	//
	// 加载系统图标或光标时，必须使用 LR_SHARED 否则函数将无法加载资源。
	//
	// 无论请求的大小如何，此函数都会查找缓存中具有请求的资源名称的第一个映像。
	LR_SHARED   = 0x00008000
	LR_VGACOLOR = 0x00000080 // 使用真正的 VGA 颜色。
)

// LoadImageW 加载图标、光标、动画光标或位图.
//
//	详见: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-loadimagew.
//	@param hInst 包含要加载的图像的 DLL 或可执行文件 (.exe) 模块的句柄。 有关详细信息，请参阅 GetModuleHandle。若要 (图标、光标或位图文件) 加载预定义图像或独立资源，请将此参数设置为0.
//	@param name 要加载的图像。如果 hInst 参数为非0且 fuLoad 参数省略 LR_LOADFROMFILE， 则 name 指定 hInst 模块中的图像资源。如果要按名称从模块加载图像资源， 则 name 参数是指向包含映像资源名称的字符串。
//	@param Type 要加载的图像的类型。 wapi.IMAGE_ .
//	@param cx 图标或光标的宽度（以像素为单位）。 如果此参数为零且 fuLoad 参数 为LR_DEFAULTSIZE，则函数使用 SM_CXICON 或 SM_CXCURSOR 系统指标值来设置宽度。 如果此参数为零且未使用 LR_DEFAULTSIZE ，则函数使用实际资源宽度。
//	@param cy 图标或光标的高度（以像素为单位）。 如果此参数为零且 fuLoad 参数 为LR_DEFAULTSIZE，则函数使用 SM_CYICON 或 SM_CYCURSOR 系统指标值来设置高度。 如果此参数为零且未使用 LR_DEFAULTSIZE ，则函数使用实际资源高度。
//	@param fuLoad 此参数可使用以下一个或多个值: wapi.LR_ .
//	@return 返回HICON。
func LoadImageW(hInst uintptr, name string, Type IMAGE_, cx, cy int32, fuLoad LR_) uintptr {
	r, _, _ := loadImageW.Call(hInst, common.StrPtr(name), uintptr(Type), uintptr(cx), uintptr(cy), uintptr(fuLoad))
	return r
}

const (
	NULL  = "\x00"
	NULL2 = NULL + NULL // 2个 NULL
)

// FindWindowW 检索顶级窗口的句柄，该窗口的类名称和窗口名称与指定的字符串匹配。 此函数不搜索子窗口。 此函数不执行区分大小写的搜索.
//
//	@Description 如果 lpWindowName 参数不 为 NULL， FindWindowW 将调用 GetWindowTextW 函数以检索窗口名称进行比较。 有关可能出现的潜在问题的说明，请参阅 GetWindowTextW 的备注。
//	详见: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-FindWindowW.
//	@param lpClassName 窗口类名, 可为空.
//	@param lpWindowName 窗口名称（窗口的标题）, 可为空.
//	@return 返回窗口句柄。
func FindWindowW(lpClassName, lpWindowName string) uintptr {
	r, _, _ := findWindowW.Call(common.StrPtr(lpClassName), common.StrPtr(lpWindowName))
	return r
}

// RegisterWindowMessageW 定义保证在整个系统中唯一的新窗口消息。 发送或发布消息时可以使用消息值.
//
//	详见: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-RegisterWindowMessageW.
//	@param lpString 要注册的消息。
//	@return int 如果成功注册消息，则返回值是范围0xC000到0xFFFF的消息标识符. 如果函数失败，则返回值为零.
func RegisterWindowMessageW(lpString string) int {
	r, _, _ := registerWindowMessageW.Call(common.StrPtr(lpString))
	return int(r)
}

// IsWindow 判断一个窗口句柄是否有效.
//
//	@Description 线程不应将 IsWindow 用于未创建的窗口，因为调用此函数后可能会销毁该窗口。 此外，由于窗口句柄被回收，句柄甚至可以指向其他窗口.
//	详见: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-iswindow.
//	@param hWnd 要测试的窗口的句柄。
//	@return bool
func IsWindow(hWnd uintptr) bool {
	r, _, _ := isWindow.Call(hWnd)
	return r != 0
}

type HWND_ int

const (
	HWND_NOTOPMOST HWND_ = -2 // 将窗口置于所有非顶层窗口之上（即在所有顶层窗口之后）。如果窗口已经是非顶层窗口则该标志不起作用。
	HWND_TOPMOST   HWND_ = -1 // 将窗口置于所有非顶层窗口之上。即使窗口未被激活, 窗口也将保持顶级位置。
	HWND_TOP       HWND_ = 0  // 将窗口置于Z序的顶部。
	HWND_BOTTOM    HWND_ = 1  // 将窗口置于Z序的底部。如果参数hWnd标识了一个顶层窗口，则窗口失去顶级位置，并且被置在所有其他窗口的底部。
)

// SWP_ 是窗口大小和定位的标志.
type SWP_ uint32

const (
	SWP_ASYNCWINDOWPOS SWP_ = 0x4000 // 如果调用线程和拥有窗口的线程连接到不同的输入队列，系统会将请求发布到拥有窗口的线程。这可以防止调用线程在其他线程处理请求时阻塞其执行。
	SWP_DEFERERASE     SWP_ = 0x2000 // 防止生成WM_SYNCPAINT消息。
	SWP_DRAWFRAME      SWP_ = 0x0020 // 在窗口周围绘制一个框架（在窗口的类描述中定义）。
	SWP_FRAMECHANGED   SWP_ = 0x0020 // 应用使用 SetWindowLong 函数 设置的新框架样式。向窗口发送WM_NCCALCSIZE消息，即使窗口大小没有改变。如果未指定此标志，则仅在更改窗口大小时发送 WM_NCCALCSIZE 。
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

// SetWindowPos 改变一个子窗口，弹出式窗口或顶层窗口的尺寸，位置和Z序。子窗口，弹出式窗口，及顶层窗口根据它们在屏幕上出现的顺序排序、顶层窗口设置的级别最高，并且被设置为Z序的第一个窗口.
//
//	@Description 详见: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-SetWindowPos.
//	@param hWnd 欲定位的窗口句柄.
//	@param hWndInsertAfter 在Z序中位于定位窗口之前的窗口句柄. 此参数必须是窗口句柄或以下值之一: wapi.HWND_.
//	@param x 窗口新的x坐标。如hwnd是一个子窗口，则x用父窗口的客户区坐标表示.
//	@param y 窗口新的y坐标。如hwnd是一个子窗口，则y用父窗口的客户区坐标表示.
//	@param cx 指定新的窗口宽度.
//	@param cy 指定新的窗口高度.
//	@param wFlags 窗口大小和定位的标志. 该参数可以是以下值的组合: wapi.SWP_.
//	@return bool
func SetWindowPos(hWnd uintptr, hWndInsertAfter HWND_, x, y, cx, cy int32, wFlags SWP_) bool {
	r, _, _ := setWindowPos.Call(hWnd, uintptr(hWndInsertAfter), uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(wFlags))
	return r != 0
}

// GetDesktopWindow 获取桌面窗口的句柄.
//
//	@Description 详见: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-GetDesktopWindow.
//	@return int
func GetDesktopWindow() uintptr {
	r, _, _ := getDesktopWindow.Call()
	return r
}

// MB_ 指示消息框的内容和行为.
type MB_ uint32

// 要指示消息框中显示的按钮，请指定以下值之一.

const (
	MB_AbortRetryIgnore  MB_ = 0x00000002 // 消息框包含三个按钮：失败、重试和忽略。
	MB_CanaelTryContinue MB_ = 0x00000006 // 消息框包含三个按钮：取消、重试、继续。使用此消息框类型而不是 MB_AbortRetryIgnore。
	MB_Help              MB_ = 0x00004000 // 向消息框 添加帮助按钮。当用户单击帮助按钮或按 F1 时，系统会向所有者 发送WM_HELP消息。
	MB_OK                MB_ = 0x00000000 // 消息框包含一个按钮：确认。这是默认设置。
	MB_OkCancel          MB_ = 0x00000001 // 消息框包含两个按钮：确认和取消。
	MB_RetryCancel       MB_ = 0x00000005 // 消息框包含两个按钮：重试和取消。
	MB_YesNo             MB_ = 0x00000004 // 消息框包含两个按钮：是和否。
	MB_YesNoCancel       MB_ = 0x00000003 // 消息框包含三个按钮：是、否和取消。
)

// 要在消息框中显示图标，请指定以下值之一.

const (
	MB_IconExclamation MB_ = 0x00000030 // 消息框中会出现一个感叹号图标。
	MB_IconWarning     MB_ = 0x00000030 // 消息框中会出现一个感叹号图标。
	MB_IconInformation MB_ = 0x00000040 // 一个由圆圈中的小写字母i组成的图标出现在消息框中。
	MB_IconAsterisk    MB_ = 0x00000040 // 一个由圆圈中的小写字母i组成的图标出现在消息框中。
	MB_IconQuestion    MB_ = 0x00000020 // 问号图标出现在消息框中。不再推荐使用问号消息图标，因为它不能清楚地表示特定类型的消息，并且作为问题的消息措辞可能适用于任何消息类型。此外，用户可能会将消息符号问号与帮助信息混淆。因此，请勿在消息框中使用此问号消息符号。系统继续支持它的包含只是为了向后兼容。
	MB_IconStop        MB_ = 0x00000010 // 一个停止标志图标出现在消息框中。
	MB_IconError       MB_ = 0x00000010 // 一个停止标志图标出现在消息框中。
	MB_IconHand        MB_ = 0x00000010 // 一个停止标志图标出现在消息框中。
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
	MB_ApplModal MB_ = 0x00000000 // 用户必须先响应消息框，然后才能在hWnd参数标识的窗口中继续工作。但是，用户可以移动到其他线程的窗口并在这些窗口中工作。根据应用程序中窗口的层次结构，用户可能能够移动到线程内的其他窗口。消息框父级的所有子窗口都会自动禁用，但弹出窗口不会。如果未指定 MB_SystemModal 或 MB_TaskModal, 则 MB_ApplModal 是默认值。

	MB_SystemModal MB_ = 0x00001000 // 与 MB_ApplModal 相同，只是消息框具有 WS_EX_TOPMOST 样式。使用系统模式消息框来通知用户需要立即注意的严重的、具有潜在破坏性的错误（例如，内存不足）。此标志对用户与除与hWnd关联的窗口之外的窗口进行交互的能力没有影响。

	MB_TaskModal MB_ = 0x00002000 // 与 MB_ApplModal 相同，除了如果hWnd参数为0则禁用所有属于当前线程的顶级窗口。当调用应用程序或库没有可用的窗口句柄但仍需要防止输入到调用线程中的其他窗口而不暂停其他线程时，请使用此标志。
)

// 要指定其他选项，请使用以下一个或多个值.

const (
	MB_Default_Desktop_Only MB_ = 0x00020000 // 与交互式窗口站的桌面相同。有关详细信息，请参阅窗口站。 如果当前输入桌面不是默认桌面，MessageBox不会返回，直到用户切换到默认桌面。

	MB_Right         MB_ = 0x00080000 // 文本右对齐。
	MB_RtlReading    MB_ = 0x00100000 // 在希伯来语和阿拉伯语系统上使用从右到左的阅读顺序显示消息和标题文本。
	MB_SetForeground MB_ = 0x00010000 // 消息框成为前台窗口。在内部，系统为消息框调用 SetForegroundWindow 函数。
	MB_TopMost       MB_ = 0x00040000 // 消息框是使用 WS_EX_TOPMOST 窗口样式创建的。

	MB_Service_Notification MB_ = 0x00200000 // 调用者是通知用户事件的服务。即使没有用户登录到计算机，该功能也会在当前活动桌面上显示一个消息框。终端服务：如果调用线程具有模拟令牌，则该函数将消息框定向到模拟令牌中指定的会话。如果设置了此标志，则hWnd参数必须为NULL。这是为了使消息框可以出现在与hWnd对应的桌面以外的桌面上。有关使用此标志的安全注意事项的信息，请参阅交互式服务。特别要注意，此标志可以在锁定的桌面上生成交互式内容，因此只能用于非常有限的一组场景，例如资源耗尽。
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
//	@Description 详见: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-MessageBoxW.
//	@param hWnd 要创建的消息框的所有者窗口的句柄。如果此参数为0，则消息框没有所有者窗口.
//	@param lpText 要显示的消息。如果字符串由多行组成，您可以在每行之间使用换行符分隔各行.
//	@param lpCaption 对话框标题。如果此参数为空，则默认标题为Error.
//	@param uType 对话框的内容和行为, 是以下值的组合: wapi.MB_.
//	@return wapi.ID_ 如果函数失败，则返回值为0; 成功则返回一个整数，指示用户单击了哪个按钮.
func MessageBoxW(hWnd uintptr, lpText, lpCaption string, uType MB_) ID_ {
	r, _, _ := messageBoxW.Call(hWnd, common.StrPtr(lpText), common.StrPtr(lpCaption), uintptr(uType))
	return ID_(r)
}

// OpenClipboard 打开剪贴板进行检查并防止其他应用程序修改剪贴板内容.
//
//	@Description 如果另一个窗口打开了剪贴板，则 OpenClipboard 会失败.
//	应用程序应在每次成功调用 OpenClipboard 后调用 CloseClipboard 函数.
//	除非调用 EmptyClipboard 函数，否则由hWndNewOwner参数标识的窗口不会成为剪贴板所有者.
//	如果应用程序在 hwnd 设置为0的情况下调用 OpenClipboard, EmptyClipboard 会将剪贴板所有者设置为NULL；这会导致 SetClipboardData 失败.
//	详见: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-OpenClipboard.
//	@param hWnd 要与打开的剪贴板关联的窗口句柄。如果此参数为0，则打开的剪贴板与当前任务相关联.
//	@return bool
func OpenClipboard(hWnd uintptr) bool {
	r, _, _ := openClipboard.Call(hWnd)
	return r != 0
}

// CloseClipboard 关闭剪贴板.
//
//	@Description 当窗口完成检查或更改剪贴板时，通过调用 CloseClipboard 关闭剪贴板。这使其他窗口能够访问剪贴板.
//	详见: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-CloseClipboard.
//	@return bool
func CloseClipboard() bool {
	r, _, _ := closeClipboard.Call()
	return r != 0
}

// EmptyClipboard 清空剪贴板并释放剪贴板中数据的句柄。然后该函数将剪贴板的所有权分配给当前打开剪贴板的窗口。
//
//	@Description 在调用 EmptyClipboard 之前，应用程序必须使用 OpenClipboard 函数打开剪贴板。
//	如果应用程序在打开剪贴板时指定了NULL窗口句柄，则 EmptyClipboard 会成功，但会将剪贴板所有者设置为NULL。请注意，这会导致 SetClipboardData 失败。
//	详见: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-EmptyClipboard.
//	@return bool
func EmptyClipboard() bool {
	r, _, _ := emptyClipboard.Call()
	return r != 0
}

// CF_ 标准剪贴板格式.
type CF_ uint32

const (
	CF_TEXT         CF_ = 1  // 文字格式。每行以回车/换行（CR-LF）组合结束。空字符表示数据的结尾。对ANSI文本使用此格式。
	CF_BITMAP       CF_ = 2  // 位图的句柄（HBITMAP）
	CF_METAFILEPICT CF_ = 3  // 处理由METAFILEPICT结构定义的图元文件图片格式。通过动态数据交换（DDE）传递CF_METAFILEPICT句柄时，负责删除【HMEM】的应用程序也应该释放CF_METAFILEPICT句柄引用的元文件。
	CF_SYLK         CF_ = 4  // Microsoft符号链接（SYLK）格式。
	CF_DIF          CF_ = 5  // 软件艺术数据交换格式。
	CF_TIFF         CF_ = 6  // 标记图像文件格式。
	CF_OEMTEXT      CF_ = 7  // 文字格式包含OEM字符集中的字符。每行以回车/换行（CR-LF）组合结束。空字符表示数据的结尾。
	CF_DIB          CF_ = 8  // 一个包含BITMAPINFO结构的内存对象，后跟位图位。
	CF_PALETTE      CF_ = 9  // 处理调色板。每当应用程序将数据放置在依赖于或假定调色板的剪贴板中时，它也应将调色板放在剪贴板上。如果剪贴板包含CF_PALETTE（逻辑调色板）格式的数据，则应用程序应使用SelectPalette和RealizePalette函数来实现（比较）剪贴板中与该逻辑调色板的任何其他数据。当显示剪贴板数据时，Windows剪贴板始终将剪贴板上的任何对象用作CF_PALETTE格式的当前调色板。
	CF_PENDATA      CF_ = 10 // 用于Pen Computing的Microsoft Windows笔的扩展数据。
	CF_RIFF         CF_ = 11 // 表示音频数据比CF_WAVE标准波形格式更复杂。
	CF_WAVE         CF_ = 12 // 以诸如11 kHz或22 kHz脉冲编码调制（PCM）的标准波形格式之一表示音频数据。
	CF_UNICODETEXT  CF_ = 13 // 仅Windows NT： Unicode文字格式。每行以回车/换行（CR-LF）组合结束。空字符表示数据的结尾。
	CF_ENHMETAFILE  CF_ = 14 // 增强图元文件的句柄（HENHMETAFILE）。
	CF_HDROP        CF_ = 15 // 类型为HDROP的句柄，用于标识文件列表。应用程序可以通过将句柄传递给DragQueryFile函数来检索有关文件的信息。
)

// IsClipboardFormatAvailable 确定剪贴板是否包含指定格式的数据.
//
//	@Description 详见: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-IsClipboardFormatAvailable.
//	@param uFormat 标准或注册的剪贴板格式, wapi.CF_ .
//	@return bool
func IsClipboardFormatAvailable(uFormat CF_) bool {
	r, _, _ := isClipboardFormatAvailable.Call(uintptr(uFormat))
	return r != 0
}

// GetClipboardData 从剪贴板中检索指定格式的数据。剪贴板必须先前已打开.
//
//	@Description 详见: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-GetClipboardData.
//	@param uFormat 剪贴板格式, wapi.CF_ .
//	@return uintptr 如果函数成功，则返回值是指定格式的剪贴板对象的句柄. 如果函数失败，则返回值为NULL.
func GetClipboardData(uFormat CF_) uintptr {
	r, _, _ := getClipboardData.Call(uintptr(uFormat))
	return r
}

// SetClipboardData 以指定的剪贴板格式将数据放在剪贴板上。该窗口必须是当前剪贴板所有者，并且应用程序必须调用 OpenClipboard 函数.
//
//	@Description 详见: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-setclipboarddata.
//	@param uFormat 标准或注册的剪贴板格式, wapi.CF_ .
//	@param hMem 指定格式的数据的句柄。该参数可以为0，表示窗口根据请求提供指定剪贴板格式的数据（渲染格式）.
//	@return uintptr 如果函数成功，则返回值是数据的句柄. 如果函数失败，则返回值为NULL.
func SetClipboardData(uFormat CF_, hMem uintptr) uintptr {
	r, _, _ := setClipboardData.Call(uintptr(uFormat), hMem)
	return r
}

// SetForegroundWindow 将创建指定窗口的线程带到前台并激活窗口. 键盘输入被定向到窗口, 并且为用户改变了各种视觉提示. 系统为创建前台窗口的线程分配比其他线程稍高的优先级.
//
//	@Description 详见: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-SetForegroundWindow.
//	@param hWnd 应激活并置于前台的窗口句柄.
//	@return bool
func SetForegroundWindow(hWnd uintptr) bool {
	r, _, _ := setForegroundWindow.Call(hWnd)
	return r != 0
}

// FindWindowExW 检索类名称和窗口名称与指定字符串匹配的窗口的句柄. 该函数搜索子窗口，从指定子窗口后面的那个开始. 此函数不执行区分大小写的搜索.
//
//	@Description 详见: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-FindWindowExW.
//	@param hWndParent 要搜索其子窗口的父窗口的句柄. 如果hwndParent为0，该函数使用桌面窗口作为父窗口. 该函数在作为桌面子窗口的窗口中进行搜索.
//	@param hWndChildAfter 子窗口的句柄。搜索从 Z 顺序中的下一个子窗口开始。子窗口必须是hwndParent的直接子窗口，而不仅仅是后代窗口。 如果hwndChildAfter为0，则搜索从hwndParent的第一个子窗口开始。 请注意，如果hwndParent和hwndChildAfter都是0，则该函数将搜索所有顶级和仅消息窗口。
//	@param lpszClass 窗口类名, 可空.
//	@param lpszWindow 窗口名称（窗口的标题）, 可空.
//	@return uintptr
func FindWindowExW(hWndParent, hWndChildAfter uintptr, lpszClass, lpszWindow string) uintptr {
	r, _, _ := findWindowExW.Call(hWndParent, hWndChildAfter, common.StrPtr(lpszClass), common.StrPtr(lpszWindow))
	return r
}

// GetWindowTextLengthW 检索指定窗口标题栏文本的长度（以字符为单位）（如果窗口有标题栏）。如果指定的窗口是控件，则该函数检索控件内文本的长度。但是无法检索另一个应用程序中编辑控件的文本长度。
//
//	@Description 详见: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-GetWindowTextLengthW.
//	@param hWnd 窗口或控件的句柄。
//	@return int 如果成功，则返回值是文本的长度（以字符为单位）。在某些情况下，此值可能大于文本的长度。如果窗口没有文本，则返回值为零。
func GetWindowTextLengthW(hWnd uintptr) int {
	r, _, _ := getWindowTextLengthW.Call(hWnd)
	return int(r)
}

// GetWindowTextW 将指定窗口标题栏（如果有）的文本复制到缓冲区中。如果指定的窗口是控件，则复制控件的文本。但是无法检索另一个应用程序中控件的文本。
//
//	@Description 详见: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-GetWindowTextW.
//	@param hWnd 包含文本的窗口或控件的句柄。
//	@param lpString 接收文本.
//	@param nMaxCount 复制到缓冲区的最大字符数，包括空字符。如果文本超出此限制，则将其截断.
//	@return int 如果函数成功，则返回值是复制字符串的长度（以字符为单位），不包括终止空字符。如果窗口没有标题栏或文本，如果标题栏为空，或者窗口或控制句柄无效，则返回值为零。
func GetWindowTextW(hWnd uintptr, lpString *string, nMaxCount int) int {
	buf := make([]uint16, nMaxCount)
	r, _, _ := getWindowTextW.Call(hWnd, common.Uint16SliceDataPtr(&buf), uintptr(nMaxCount))
	*lpString = syscall.UTF16ToString(buf[0:])
	return int(r)
}

// ClientToScreen 将指定点的客户区坐标转换为屏幕坐标。
//
//	@Description 详见: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-ClientToScreen.
//	@param hWnd 窗口真实句柄
//	@param lpPoint wapi.POINT 指针. 如果函数成功，则将新的屏幕坐标复制到此结构中.
//	@return bool
func ClientToScreen(hWnd uintptr, lpPoint *POINT) bool {
	r, _, _ := clientToScreen.Call(hWnd, uintptr(unsafe.Pointer(lpPoint)))
	return r != 0
}

// GetCursorPos 检索鼠标光标的位置，以屏幕坐标表示.
//
//	@Description 详见: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-getcursorpos.
//	@param lpPoint 指向接收光标屏幕坐标的 wapi.POINT 结构的指针.
//	@return bool
func GetCursorPos(lpPoint *POINT) bool {
	r, _, _ := getCursorPos.Call(uintptr(unsafe.Pointer(lpPoint)))
	return r != 0
}

type Mod_ uint32

const (
	Mod_Alt      Mod_ = 0x0001 // 必须按住任一 ALT 键。
	Mod_Control  Mod_ = 0x0002 // 必须按住任一 CTRL 键。
	Mod_Norepeat Mod_ = 0x4000 // 更改热键行为，以便键盘自动重复不会产生多个热键通知。Windows Vista：  不支持此标志。
	Mod_Shift    Mod_ = 0x0004 // 必须按住任一 SHIFT 键。
	Mod_Win      Mod_ = 0x0008 // 任一 WINDOWS 键被按住。这些键标有 Windows 徽标。涉及 WINDOWS 键的键盘快捷键保留供操作系统使用。
)

// RegisterHotKey 注册系统范围的热键.
//
//	@Description 详见: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-registerhotkey.
//	@param hWnd 真实窗口句柄。将接收由热键生成的 WM_HOTKEY 消息的窗口句柄。如果此参数为0，则 WM_HOTKEY 消息将发布到调用线程的消息队列中，并且必须在消息循环中进行处理。
//	@param id 热键的标识符。如果hWnd参数为0，则热键与当前线程相关联，而不是与特定窗口相关联。如果已存在具有相同hWnd和id参数的热键，请参阅备注了解所采取的操作。
//	@param fsModifiers 为了生成 WM_HOTKEY 消息，必须与vk参数指定的键组合按下的键 。fsModifiers参数可以是以下值的组合: xcc.Mod_ .
//	@param vk 热键的虚拟键代码: xcc.VK_ . 请参阅虚拟键码: https://docs.microsoft.com/zh-cn/windows/win32/inputdev/virtual-key-codes.
//	@return bool
func RegisterHotKey(hWnd uintptr, id int32, fsModifiers, vk uint32) bool {
	r, _, _ := registerHotKey.Call(hWnd, uintptr(id), uintptr(fsModifiers), uintptr(vk))
	return r != 0
}

// UnregisterHotKey 释放先前注册的热键.
//
//	@Description 详见: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-unregisterHotKey.
//	@param hWnd 真实窗口句柄。与要释放的热键关联的窗口句柄。如果热键与窗口无关，则此参数应为0.
//	@param id 要释放的热键的标识符.
//	@return bool
func UnregisterHotKey(hWnd uintptr, id int32) bool {
	r, _, _ := unregisterHotKey.Call(hWnd, uintptr(id))
	return r != 0
}

// GetMessage 从调用线程的消息队列中检索消息。应用程序通常使用返回值来确定是否结束主消息循环并退出程序。该函数分派传入的已发送消息，直到发布的消息可用于检索。 与 GetMessage 不同， PeekMessage 函数在返回之前不会等待消息发布。
//
//	@Description: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-GetMessageW.
//	@param pMsg 指向从线程的消息队列接收消息信息的 MSG 结构的指针。
//	@param hWnd 要检索其消息的窗口的句柄。窗口必须属于当前线程。如果hWnd为0， GetMessage 检索属于当前线程的任何窗口的消息，以及当前线程的消息队列中hwnd值为0的任何消息（参见 MSG 结构）。因此，如果hWnd为0，则同时处理窗口消息和线程消息。如果hWnd为-1， GetMessage 仅检索当前线程的消息队列中hwnd值为0的消息，即 PostMessage （当hWnd参数为0时）或 PostThreadMessage 发布的线程消息。
//	@param wMsgFilterMin 要检索的最低消息值的整数值。使用WM_KEYFIRST (0x0100) 指定第一条键盘消息或WM_MOUSEFIRST (0x0200) 指定第一条鼠标消息。
//	@param wMsgFilterMax 要检索的最高消息值的整数值。使用WM_KEYLAST指定最后一个键盘消息或WM_MOUSELAST指定最后一个鼠标消息。
//	@return int32 如果函数检索到 WM_QUIT 以外的消息，则返回值非零。如果函数检索到 WM_QUIT 消息，则返回值为零。如果有错误，返回值为-1。
func GetMessage(pMsg *MSG, hWnd uintptr, wMsgFilterMin uint32, wMsgFilterMax uint32) int32 {
	r, _, _ := getMessageW.Call(uintptr(unsafe.Pointer(pMsg)), hWnd, uintptr(wMsgFilterMin), uintptr(wMsgFilterMax))
	return int32(r)
}

// TranslateMessage 将虚拟键消息转换为字符消息。字符消息被发布到调用线程的消息队列中，以便在线程下次调用 GetMessage 或 PeekMessage 函数时读取。
//
//	@Description: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-TranslateMessage.
//	@param pMsg 一个指向 MSG 结构的指针，该结构包含使用 GetMessage 或 PeekMessage 函数从调用线程的消息队列中检索到的消息信息。
//	@return bool
func TranslateMessage(pMsg *MSG) bool {
	r, _, _ := translateMessage.Call(uintptr(unsafe.Pointer(pMsg)))
	return r != 0
}

// DispatchMessage 向窗口过程发送消息。它通常用于发送由 GetMessage 函数检索到的消息。
//
//	@Description: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-DispatchMessageW.
//	@param pMsg 指向包含消息的结构的指针。
//	@return int 返回值指定窗口过程返回的值。尽管它的含义取决于所发送的消息，但返回值通常会被忽略。
func DispatchMessage(pMsg *MSG) int {
	r, _, _ := dispatchMessageW.Call(uintptr(unsafe.Pointer(pMsg)))
	return int(r)
}

// PostQuitMessage 向系统指示线程已请求终止（退出）。它通常用于响应 WM_DESTROY 消息。
//
//	@Description: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-PostQuitMessage.
//	@param nExitCode 应用程序退出代码。该值用作 WM_QUIT 消息的wParam参数。
func PostQuitMessage(nExitCode int32) {
	postQuitMessage.Call(uintptr(nExitCode))
}

type MSG struct {
	Hwnd    uintptr
	Message uint32
	WParam  uint
	LParam  uint
	Time    uint32
	Pt      POINT
}

type POINT struct {
	X int32
	Y int32
}

// SendMessageW 将指定的消息发送到一个或多个窗口。SendMessage函数调用指定窗口的窗口过程，直到窗口过程处理完消息才返回。
//
//	@Description 详见: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-SendMessageW.
//	@param hWnd 窗口句柄，其窗口过程将接收消息。如果该参数为 HWND_BROADCAST ((HWND)0xffff)，则将消息发送到系统中的所有顶层窗口，包括禁用或不可见的无主窗口、重叠窗口和弹出窗口；但消息不会发送到子窗口。
//	@param Msg 要发送的消息。有关系统提供的消息的列表，请参阅: https://docs.microsoft.com/en-us/windows/win32/winmsg/about-messages-and-message-queues.
//	@param wParam 其他特定于消息的信息。
//	@param lParam 其他特定于消息的信息。
//	@return int 返回值指定消息处理的结果；这取决于发送的消息。
func SendMessageW(hWnd uintptr, Msg int32, wParam, lParam uint) int {
	r, _, _ := sendMessageW.Call(hWnd, uintptr(Msg), uintptr(wParam), uintptr(lParam))
	return int(r)
}

// PostMessageW 在与创建指定窗口的线程关联的消息队列中放置（发布）一条消息，并在不等待线程处理消息的情况下返回。
//
//	@Description 详见: https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-postmessagew.
//	@param hWnd 窗口句柄，其窗口过程将接收消息。如果该参数为 HWND_BROADCAST ((HWND)0xffff)，则将消息发送到系统中的所有顶层窗口，包括禁用或不可见的无主窗口、重叠窗口和弹出窗口；但消息不会发送到子窗口。
//	@param Msg 要发送的消息。有关系统提供的消息的列表，请参阅: https://docs.microsoft.com/en-us/windows/win32/winmsg/about-messages-and-message-queues.
//	@param wParam 其他特定于消息的信息。
//	@param lParam 其他特定于消息的信息。
//	@return bool
func PostMessageW(hWnd uintptr, Msg int32, wParam, lParam uint) bool {
	r, _, _ := postMessageW.Call(hWnd, uintptr(Msg), uintptr(wParam), uintptr(lParam))
	return r != 0
}

package app

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// App 程序. 封装了炫彩的全局API.
type App struct {
}

// New 炫彩_初始化, 失败返回nil.
//
//	@Description 默认会在程序运行目录和系统system32目录寻找并加载xcgui.dll.
//	如果你想要更改xcgui.dll的路径, 那么请在调用本函数之前调用 xc.SetXcguiPath().
//	@param bD2D 是否启用D2D.
//	@return *App
func New(bD2D bool) *App {
	moudle := xc.LoadXCGUI()
	if moudle.Handle() == 0 {
		return nil
	}

	p := &App{}
	if !xc.XInitXCGUI(bD2D) {
		return nil
	}
	return p
}

// Run 炫彩_运行. 运行消息循环, 当炫彩窗口数量为0时退出.
//
//	@return int
func (a *App) Run() int {
	return xc.XRunXCGUI()
}

// Exit 炫彩_退出, 退出界面库释放资源.
//
//	@return int
func (a *App) Exit() int {
	return xc.XExitXCGUI()
}

// ShowAndRun 显示窗口并调用炫彩_运行.
//
//	@param hWindow 炫彩窗口句柄.
func (a *App) ShowAndRun(hWindow int) {
	xc.XWnd_ShowWindow(hWindow, xcc.SW_SHOW)
	xc.XRunXCGUI()
}

// DebugToFileInfo 炫彩_输出调试信息到文件, 打印调试信息到文件xcgui_debug.txt.
//
//	@param pInfo 文本.
func (a *App) DebugToFileInfo(pInfo string) {
	xc.XC_DebugToFileInfo(pInfo)
}

// SetActivateTopWindow 炫彩_激活窗口, 激活当前进程最上层窗口.
//
//	@return bool
func (a *App) SetActivateTopWindow() bool {
	return xc.XC_SetActivateTopWindow()
}

// GetDefaultFont 炫彩_取默认字体.
//
//	@return int 字体句柄.
func (a *App) GetDefaultFont() int {
	return xc.XC_GetDefaultFont()
}

// MessageBox 炫彩_消息框.
//
//	@param pTitle 标题.
//	@param pText 内容文本.
//	@param nFlags 标识: xcc.MessageBox_Flag_.
//	@param hWndParent 父窗口句柄(真实的窗口句柄).
//	@param XCStyle xcc.Window_Style_.
//	@return xcc.MessageBox_Flag_. 返回: xcc.MessageBox_Flag_Ok: 点击确定按钮退出. xcc.MessageBox_Flag_Cancel: 点击取消按钮退出. xcc.MessageBox_Flag_Other: 其他方式退出.
func (a *App) MessageBox(pTitle, pText string, nFlags xcc.MessageBox_Flag_, hWndParent uintptr, XCStyle xcc.Window_Style_) xcc.MessageBox_Flag_ {
	return xc.XC_MessageBox(pTitle, pText, nFlags, hWndParent, XCStyle)
}

// Msg_Create 消息框_创建, 此窗口是一个模态窗口, 弹出窗口请调用 xc.XModalWnd_DoModal().
//
//	@param pTitle 标题.
//	@param pText 内容文本.
//	@param nFlags 标识: xcc.MessageBox_Flag_.
//	@param hWndParent 父窗口句柄(真实的窗口句柄).
//	@param XCStyle xcc.Window_Style_.
//	@return int 返回消息框窗口句柄.
func (a *App) Msg_Create(pTitle, pText string, nFlags xcc.MessageBox_Flag_, hWndParent uintptr, XCStyle xcc.Window_Style_) int {
	return xc.XMsg_Create(pTitle, pText, nFlags, hWndParent, XCStyle)
}

// Msg_CreateEx 消息框_创建扩展, 此窗口是一个模态窗口, 弹出窗口请调用 xc.XModalWnd_DoModal().
//
//	@param dwExStyle 窗口扩展样式.
//	@param dwStyle 窗口样式.
//	@param lpClassName 窗口类名.
//	@param pTitle 标题.
//	@param pText 内容文本.
//	@param nFlags 标识: xcc.MessageBox_Flag_.
//	@param hWndParent 父窗口句柄(真实的窗口句柄).
//	@param XCStyle xcc.Window_Style_.
//	@return int 消息框窗口句柄.
func (a *App) Msg_CreateEx(dwExStyle int, dwStyle int, lpClassName string, pTitle, pText string, nFlags xcc.MessageBox_Flag_, hWndParent uintptr, XCStyle xcc.Window_Style_) int {
	return xc.XMsg_CreateEx(dwExStyle, dwStyle, lpClassName, pTitle, pText, nFlags, hWndParent, XCStyle)
}

// 炫彩_发送窗口消息.
//
// hWindow: 窗口句柄.
//
// msg:.
//
// wParam:.
//
// lParam:.
func (a *App) SendMessage(hWindow int, msg uint32, wParam, lParam uint) uint {
	return xc.XC_SendMessage(hWindow, msg, wParam, lParam)
}

// 炫彩_投递窗口消息.
//
// hWindow: 窗口句柄.
//
// msg:.
//
// wParam:.
//
// lParam:.
func (a *App) PostMessage(hWindow int, msg uint32, wParam int32, lParam int32) bool {
	return xc.XC_PostMessage(hWindow, msg, wParam, lParam)
}

// CallUiThread 炫彩_调用界面线程, 调用UI线程, 设置回调函数, 在回调函数里操作UI.
//
//	@Description: 回调函数尽量不要使用匿名函数, 使用匿名函数意味着你每次都在创建1个新的回调, 超过2000个时, 程序必将panic.
//	如果使用 CallUiThreadEx 和 CallUiThreader 则没有此限制.
//	@param pCall 回调函数.
//	@param data 传进回调函数的用户自定义数据.
//	@return int
func (a *App) CallUiThread(pCall func(data int) int, data int) int {
	return xc.XC_CallUiThread(pCall, data)
}

// CallUiThreadEx 炫彩_调用界面线程, 调用UI线程, 设置回调函数, 在回调函数里操作UI.
//
//	@Description: 与 CallUiThread 的区别是: 本函数没有2000个回调上限的限制, 回调函数可以直接使用匿名函数.
//	@param f 回调函数.
//	@param data 传进回调函数的用户自定义数据.
//	@return int
func (a *App) CallUiThreadEx(pCall func(data int) int, data int) int {
	return xc.XC_CallUiThreadEx(pCall, data)
}

// CallUT 炫彩_调用界面线程, 调用UI线程, 设置回调函数, 在回调函数里操作UI.
//
//	@Description: 与 CallUiThread 的区别是: 本函数没有2000个回调上限的限制, 回调函数可以直接使用匿名函数. 回调函数没有参数也没有返回值.
//	@param f 回调函数, 没有参数也没有返回值, 可以直接使用匿名函数.
func (a *App) CallUT(f func()) {
	xc.XC_CallUT(f)
}

// CallUiThreader 炫彩_调用界面线程, 调用UI线程, 设置回调函数, 在回调函数里操作UI.
//
//	@Description: 与 CallUiThread 的区别是: 本函数没有2000个回调上限的限制, 回调函数可以直接使用匿名函数.
//	@param u xc.UiThreader.
//	@param data 传进回调函数的用户自定义数据.
//	@return int
func (a *App) CallUiThreader(u xc.UiThreader, data int) int {
	return xc.XC_CallUiThreader(u, data)
}

// 炫彩_判断元素, 判断是否为元素句柄.
//
// hEle: 元素句柄.
func (a *App) IsHELE(hEle int) bool {
	return xc.XC_IsHELE(hEle)
}

// 炫彩_判断窗口, 判断是否为窗口句柄.
//
// hWindow: 窗口句柄.
func (a *App) IsHWINDOW(hWindow int) bool {
	return xc.XC_IsHWINDOW(hWindow)
}

// 炫彩_判断形状对象, 判断是否为形状对象.
//
// hShape: 形状对象句柄.
func (a *App) IsShape(hShape int) bool {
	return xc.XC_IsShape(hShape)
}

// 炫彩_判断句柄包含类型, 判断句柄是否拥有该类型.
//
// hXCGUI: 炫彩句柄.
//
// nType: 句柄类型, XC_OBJECT_TYPE , 以XC_开头的常量.
func (a *App) IsHXCGUI(hXCGUI int, nType xcc.XC_OBJECT_TYPE) bool {
	return xc.XC_IsHXCGUI(hXCGUI, nType)
}

// 炫彩_转换HWND到HWINDOW, 通过窗口HWND句柄获取HWINDOW句柄.
//
// hWnd: 窗口真实句柄HWND.
func (a *App) HWindowFromHWnd(hWnd uintptr) int {
	return xc.XC_hWindowFromHWnd(hWnd)
}

// 炫彩_置属性, 设置对象属性.
//
// hXCGUI: 对象句柄.
//
// pName: 属性名.
//
// pValue: 属性值.
func (a *App) SetProperty(hXCGUI int, pName string, pValue string) bool {
	return xc.XC_SetProperty(hXCGUI, pName, pValue)
}

// 炫彩_取属性, 获取对象属性, 返回属性值.
//
// hXCGUI: 对象句柄.
//
// pName: 属性名.
func (a *App) GetProperty(hXCGUI int, pName string) string {
	return xc.XC_GetProperty(hXCGUI, pName)
}

// 炫彩_注册窗口类名, 如果是在DLL中使用, 那么DLL卸载时需要注销窗口类名, 否则DLL卸载后, 类名所指向的窗口过程地址失效.
//
// pClassName: 类名.
func (a *App) RegisterWindowClassName(pClassName string) bool {
	return xc.XC_RegisterWindowClassName(pClassName)
}

// 炫彩_判断滚动视图扩展元素, 判断元素是否从滚动视图元素扩展的新元素, 包含滚动视图元素.
//
// hEle: 元素句柄.
func (a *App) IsSViewExtend(hEle int) bool {
	return xc.XC_IsSViewExtend(hEle)
}

// 炫彩_取对象类型, 获取句柄类型, 返回: xcc.XC_OBJECT_TYPE, 以XC_开头的常量.
//
// hXCGUI: 炫彩对象句柄.
func (a *App) GetObjectType(hXCGUI int) xcc.XC_OBJECT_TYPE {
	return xc.XC_GetObjectType(hXCGUI)
}

// 炫彩_取对象从ID, 通过ID获取对象句柄, 不包括窗口对象.
//
// hWindow: 所属窗口句柄.
//
// nID: ID值.
func (a *App) GetObjectByID(hWindow int, nID int) int {
	return xc.XC_GetObjectByID(hWindow, nID)
}

// 炫彩_取对象从ID名称, 通过ID名称获取对象句柄.
//
// hWindow: 所属窗口句柄.
//
// pName: ID名称.
func (a *App) GetObjectByIDName(hWindow int, pName string) int {
	return xc.XC_GetObjectByIDName(hWindow, pName)
}

// 炫彩_取对象从UID, 通过UID获取对象句柄, 不包括窗口对象.
//
// nUID: UID值.
func (a *App) GetObjectByUID(nUID int) int {
	return xc.XC_GetObjectByUID(nUID)
}

// 炫彩_取对象从UID名称, 通过UID名称获取对象句柄.
//
// pName: UID名称.
func (a *App) GetObjectByUIDName(pName string) int {
	return xc.XC_GetObjectByUIDName(pName)
}

// 炫彩_取对象从名称, 通过name获取对象句柄.
//
// pName: name名称.
func (a *App) GetObjectByName(pName string) int {
	return xc.XC_GetObjectByName(pName)
}

// 炫彩_置绘制频率, 设置UI的最小重绘频率.
//
// nMilliseconds: 重绘最小时间间隔, 单位毫秒.
func (a *App) SetPaintFrequency(nMilliseconds int) {
	xc.XC_SetPaintFrequency(nMilliseconds)
}

// 炫彩_置文本渲染质量, 设置文本渲染质量.
//
// nType: 参见GDI+ TextRenderingHint 定义.
func (a *App) SetTextRenderingHint(nType int) {
	xc.XC_SetTextRenderingHint(nType)
}

// 炫彩_启用GDI绘制文本, 将影响到以下函数: XDraw_TextOut, XDraw_TextOutEx, XDraw_TextOutA.
//
// bEnable: 是否启用.
func (a *App) EnableGdiDrawText(bEnable bool) {
	xc.XC_EnableGdiDrawText(bEnable)
}

// 炫彩_判断矩形相交, 判断两个矩形是否相交及重叠.
//
// pRect1: 矩形1.
//
// pRect2: 矩形2.
func (a *App) RectInRect(pRect1 *xc.RECT, pRect2 *xc.RECT) bool {
	return xc.XC_RectInRect(pRect1, pRect2)
}

// 炫彩_组合矩形, 组合两个矩形区域.
//
// pDest: 新的矩形区域.
//
// pSrc1: 源矩形1.
//
// pSrc2: 源矩形2.
func (a *App) CombineRect(pDest *xc.RECT, pSrc1 *xc.RECT, pSrc2 *xc.RECT) {
	xc.XC_CombineRect(pDest, pSrc1, pSrc2)
}

// 炫彩_显示布局边界, 显示布局对象边界.
//
// bShow: 是否显示.
func (a *App) ShowLayoutFrame(bShow bool) {
	xc.XC_ShowLayoutFrame(bShow)
}

// 炫彩_启用debug文件.
//
// bEnable: 是否启用.
func (a *App) EnableDebugFile(bEnable bool) {
	xc.XC_EnableDebugFile(bEnable)
}

// 炫彩_启用资源监视器.
//
// bEnable: 是否启用.
func (a *App) EnableResMonitor(bEnable bool) {
	xc.XC_EnableResMonitor(bEnable)
}

// 炫彩_置布局边界颜色.
//
// color: ABGR 颜色值.
func (a *App) SetLayoutFrameColor(color int) int {
	return xc.XC_SetLayoutFrameColor(color)
}

// 炫彩_启用错误弹窗, 启用错误弹出, 通过该接口可以设置遇到严重错误时不弹出消息提示框.
//
// bEnabel: 是否启用.
func (a *App) EnableErrorMessageBox(bEnabel bool) int {
	return xc.XC_EnableErrorMessageBox(bEnabel)
}

// 炫彩_启用自动退出程序, 启动或禁用自动退出程序, 当检测到所有用户创建的窗口都关闭时, 自动退出程序; 可调用 XC_PostQuitMessage() 手动退出程序.
//
// bEnabel: 是否启用.
func (a *App) EnableAutoExitApp(bEnabel bool) int {
	return xc.XC_EnableAutoExitApp(bEnabel)
}

// 炫彩_取文本绘制大小.
//
// pString: 字符串.
//
// length: 字符串长度.
//
// hFontX: 字体.
//
// pOutSize: 接收返回大小.
func (a *App) GetTextSize(pString string, length int, hFontX int, pOutSize *xc.SIZE) int {
	return xc.XC_GetTextSize(pString, length, hFontX, pOutSize)
}

// 炫彩_取文本显示大小.
//
// pString: 字符串.
//
// length: 字符串长度.
//
// hFontX: 字体.
//
// pOutSize: 接收返回大小.
func (a *App) GetTextShowSize(pString string, length int, hFontX int, pOutSize *xc.SIZE) int {
	return xc.XC_GetTextShowSize(pString, length, hFontX, pOutSize)
}

// 炫彩_取文本显示大小扩展.
//
// pString: 字符串.
//
// length: 字符串长度.
//
// hFontX: 字体.
//
// nTextAlign: 文本对齐方式, TextFormatFlag_, TextAlignFlag_, TextTrimming_.
//
// pOutSize: 接收返回大小.
func (a *App) GetTextShowSizeEx(pString string, length int, hFontX int, nTextAlign xcc.TextFormatFlag_, pOutSize *xc.SIZE) int {
	return xc.XC_GetTextShowSizeEx(pString, length, hFontX, nTextAlign, pOutSize)
}

// GetTextShowRect 炫彩_取文本显示矩形.
//
//	@param pString 字符串.
//	@param length 字符串长度.
//	@param hFontX 字体.
//	@param nTextAlign 文本对齐: xcc.TextFormatFlag_.
//	@param width 最大宽度.
//	@param pOutSize 接收返回大小.
//	@return int
func (a *App) GetTextShowRect(pString string, length int, hFontX int, nTextAlign xcc.TextFormatFlag_, width int, pOutSize *xc.SIZE) int {
	return xc.XC_GetTextShowRect(pString, length, hFontX, nTextAlign, width, pOutSize)
}

// 炫彩_置默认字体.
//
// hFontX: 炫彩字体句柄.
func (a *App) SetDefaultFont(hFontX int) int {
	return xc.XC_SetDefaultFont(hFontX)
}

// 炫彩_添加搜索路径, 添加文件搜索路径, 默认路径为exe目录和程序当前运行目录.
//
// pPath: 文件夹.
func (a *App) AddFileSearchPath(pPath string) int {
	return xc.XC_AddFileSearchPath(pPath)
}

// 炫彩_初始化字体, 初始化LOGFONTW结构体.
//
// pFont: LOGFONTW结构体指针.
//
// pName: 字体名称.
//
// size: 字体大小.
//
// bBold: 是否为粗体.
//
// bItalic: 是否为斜体.
//
// bUnderline: 是否有下划线.
//
// bStrikeOut: 是否有删除线.
func (a *App) InitFont(pFont *xc.LOGFONTW, pName string, size int, bBold bool, bItalic bool, bUnderline bool, bStrikeOut bool) int {
	return xc.XC_InitFont(pFont, pName, size, bBold, bItalic, bUnderline, bStrikeOut)
}

// 炫彩_分配内存, 在UI库中申请内存, 返回: 内存首地址.
//
// size: 大小, 字节为单位.
func (a *App) Malloc(size int) int {
	return xc.XC_Malloc(size)
}

// 炫彩_释放内存, 在UI库中释放内存.
//
// p: 内存首地址.
func (a *App) Free(p int) int {
	return xc.XC_Free(p)
}

// 炫彩_弹框, 弹出提示框.
//
// pTitle: 提示框标题.
//
// pText: 提示内容.
func (a *App) Alert(pTitle string, pText string) int {
	return xc.XC_Alert(pTitle, pText)
}

// 对指定文件执行操作. 如果函数成功，则返回大于 32 的值。如果函数失败，则返回指示失败原因的错误值.
//
// hwnd: 用于显示 UI 或错误消息的父窗口的句柄。如果操作与窗口无关，则此值可以为0.
//
// lpOperation: 填“open”则打开lpFlie文档. 其它操作详见: https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shellexecutew.
//
// lpFile: 想用关联的程序打印或打开的一个程序名或文件名.
//
// lpParameters: 如果lpFile是一个可执行文件，则这个字串包含了传递给执行程序的参数.
//
// lpDirectory: 想使用的默认路径完整路径.
//
// nShowCmd: 定义了如何显示启动程序的常数值: xcc.SW_.
func (a *App) Sys_ShellExecute(hwnd uintptr, lpOperation string, lpFile string, lpParameters string, lpDirectory string, nShowCmd xcc.SW_) uintptr {
	return xc.XC_Sys_ShellExecute(hwnd, lpOperation, lpFile, lpParameters, lpDirectory, nShowCmd)
}

// 炫彩_载入动态库, 系统API LoadLibrary, 返回动态库模块句柄.
//
// lpFileName: 文件名.
func (a *App) LoadLibrary(lpFileName string) uintptr {
	return xc.XC_LoadLibrary(lpFileName)
}

// 炫彩_取动态库中函数地址, 系统API GetProcAddress, 返回函数地址.
//
// hModule: 动态库模块句柄.
//
// lpProcName: 函数名.
func (a *App) GetProcAddress(hModule uintptr, lpProcName string) uintptr {
	return xc.XC_GetProcAddress(hModule, lpProcName)
}

// 炫彩_释放动态库, 系统API FreeLibrary.
//
// hModule: 动态库模块句柄.
func (a *App) FreeLibrary(hModule uintptr) bool {
	return xc.XC_FreeLibrary(hModule)
}

// 炫彩_加载DLL, 返回DLL模块句柄. 加载指定DLL, 并且调用DLL中函数LoadDll(), DLL中导出函数格式: int WINAPI LoadDll().
//
// pDllFileName: DLL文件名.
func (a *App) LoadDll(pDllFileName string) uintptr {
	return xc.XC_LoadDll(pDllFileName)
}

// 炫彩_PostQuitMessage, 发送WM_QUIT消息退出消息循环.
//
// nExitCode: 退出码.
func (a *App) PostQuitMessage(nExitCode int) int {
	return xc.XC_PostQuitMessage(nExitCode)
}

// 炫彩_加载布局文件, 返回窗口句柄或布局句柄或元素句柄.
//
// pFileName: 布局文件名.
//
// hParent: 父对象句柄.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func (a *App) LoadLayout(pFileName string, hParent int, hAttachWnd uintptr) int {
	return xc.XC_LoadLayout(pFileName, hParent, hAttachWnd)
}

// 炫彩_加载布局文件ZIP, 加载布局文件从zip压缩包中, 返回窗口句柄或布局句柄或元素句柄.
//
// pZipFileName: zip文件名.
//
// pFileName: 布局文件名.
//
// pPassword: zip密码.
//
// hParent: 父对象句柄.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func (a *App) LoadLayoutZip(pZipFileName string, pFileName string, pPassword string, hParent int, hAttachWnd uintptr) int {
	return xc.XC_LoadLayoutZip(pZipFileName, pFileName, pPassword, hParent, hAttachWnd)
}

// 炫彩_加载布局文件内存ZIP, 加载布局文件从zip压缩包中, 返回窗口句柄或布局句柄或元素句柄.
//
// data: 内存块指针.
//
// pFileName: 布局文件名.
//
// pPassword: zip密码.
//
// hParent: 父对象句柄.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func (a *App) LoadLayoutZipMem(data []byte, pFileName string, pPassword string, hParent int, hAttachWnd uintptr) int {
	return xc.XC_LoadLayoutZipMem(data, pFileName, pPassword, hParent, hAttachWnd)
}

// 炫彩_加载布局文件资源ZIP扩展. 加载布局文件从RC资源zip压缩包中, 返回窗口句柄或元素句柄.
//
// id: RC资源ID.
//
// pFileName: 布局文件名.
//
// pPassword: zip密码.
//
// pPrefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄, 窗口句柄或UI元素句柄, 可填0.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用, 可填0.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
//
// hModule: 模块句柄, 可填0.
func (a *App) LoadLayoutZipResEx(id int32, pFileName string, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd, hModule uintptr) int {
	return xc.XC_LoadLayoutZipResEx(id, pFileName, pPassword, pPrefixName, hParent, hParentWnd, hAttachWnd, hModule)
}

// 炫彩_加载资源文件资源ZIP. 加载资源文件从RC资源zip压缩包中.
//
// id: RC资源ID.
//
// pFileName: 资源文件名.
//
// pPassword: zip压缩包密码.
//
// hModule: 模块句柄, 可填0.
func (a *App) LoadResourceZipRes(id int, pFileName string, pPassword string, hModule uintptr) bool {
	return xc.XC_LoadResourceZipRes(id, pFileName, pPassword, hModule)
}

// 炫彩_加载样式文件从资源ZIP. 从RC资源中的ZIP包中, 加载样式文件.
//
// id: RC资源ID.
//
// pFileName: 文件名.
//
// pPassword: 密码.
//
// hModule: 模块句柄, 可填0.
func (a *App) LoadStyleZipRes(id int, pFileName string, pPassword string, hModule uintptr) bool {
	return xc.XC_LoadStyleZipRes(id, pFileName, pPassword, hModule)
}

// 炫彩_加载布局文件从字符串W, 加载布局文件从内存字符串, 返回窗口句柄或布局句柄或元素句柄.
//
// pStringXML: 字符串.
//
// hParent: 父对象.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func (a *App) LoadLayoutFromStringW(pStringXML string, hParent int, hAttachWnd uintptr) int {
	return xc.XC_LoadLayoutFromStringW(pStringXML, hParent, hAttachWnd)
}

// 炫彩_加载布局文件Ex, 返回窗口句柄或布局句柄或元素句柄.
//
// pFileName: 布局文件名.
//
// pPrefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄, 窗口句柄或UI元素句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func (a *App) LoadLayoutEx(pFileName, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) int {
	return xc.XC_LoadLayoutEx(pFileName, pPrefixName, hParent, hParentWnd, hAttachWnd)
}

// 炫彩_加载布局文件ZIPEx, 加载布局文件从zip压缩包中, 返回窗口句柄或布局句柄或元素句柄.
//
// pZipFileName: zip文件名.
//
// pFileName: 布局文件名.
//
// pPassword: zip密码.
//
// pPrefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄, 窗口句柄或UI元素句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func (a *App) LoadLayoutZipEx(pZipFileName string, pFileName string, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) int {
	return xc.XC_LoadLayoutZipEx(pZipFileName, pFileName, pPassword, pPrefixName, hParent, hParentWnd, hAttachWnd)
}

// 炫彩_加载布局文件内存ZIPEx, 加载布局文件从zip压缩包中, 返回窗口句柄或布局句柄或元素句柄.
//
// data: 布局文件数据.
//
// pFileName: 布局文件名.
//
// pPassword: zip密码.
//
// pPrefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄, 窗口句柄或UI元素句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func (a *App) LoadLayoutZipMemEx(data []byte, pFileName string, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) int {
	return xc.XC_LoadLayoutZipMemEx(data, pFileName, pPassword, pPrefixName, hParent, hParentWnd, hAttachWnd)
}

// 炫彩_加载布局文件从字符串WEx, 加载布局文件从内存字符串, 返回窗口句柄或布局句柄或元素句柄.
//
// pStringXML: 字符串.
//
// pPrefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄, 窗口句柄或UI元素句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func (a *App) LoadLayoutFromStringWEx(pStringXML, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) int {
	return xc.XC_LoadLayoutFromStringWEx(pStringXML, pPrefixName, hParent, hParentWnd, hAttachWnd)
}

/*
// 炫彩_加载布局文件从字符串, 加载布局文件从内存字符串, 返回窗口句柄或布局句柄或元素句柄.
//
// pStringXML: 字符串.
//
// hParent: 父对象.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func (a *App) LoadLayoutFromString(pStringXML string, hParent, hAttachWnd int) int {
	return xc.XC_LoadLayoutFromString(pStringXML, hParent, hAttachWnd)
}
*/

// 炫彩_加载样式文件.
//
// pFileName: 样式文件名称.
func (a *App) LoadStyle(pFileName string) bool {
	return xc.XC_LoadStyle(pFileName)
}

// 炫彩_加载样式文件ZIP.
//
// pZipFile: ZIP文件名.
//
// pFileName: 文件名.
//
// pPassword: 密码.
func (a *App) LoadStyleZip(pZipFile string, pFileName string, pPassword string) bool {
	return xc.XC_LoadStyleZip(pZipFile, pFileName, pPassword)
}

// 炫彩_加载样式文件从内存ZIP.
//
// data: 样式文件数据.
//
// pFileName: 文件名.
//
// pPassword: 密码.
func (a *App) LoadStyleZipMem(data []byte, pFileName string, pPassword string) bool {
	return xc.XC_LoadStyleZipMem(data, pFileName, pPassword)
}

// 炫彩_加载资源文件.
//
// pFileName: 资源文件名.
func (a *App) LoadResource(pFileName string) bool {
	return xc.XC_LoadResource(pFileName)
}

// 炫彩_加载资源文件ZIP.
//
// pZipFileName: zip文件名.
//
// pFileName: 资源文件名.
//
// pPassword: zip压缩包密码.
func (a *App) LoadResourceZip(pZipFileName string, pFileName string, pPassword string) bool {
	return xc.XC_LoadResourceZip(pZipFileName, pFileName, pPassword)
}

// 炫彩_加载资源文件内存ZIP.
//
// data: 资源文件数据.
//
// pFileName: 资源文件名.
//
// pPassword: zip压缩包密码.
func (a *App) LoadResourceZipMem(data []byte, pFileName string, pPassword string) bool {
	return xc.XC_LoadResourceZipMem(data, pFileName, pPassword)
}

// 炫彩_加载资源文件从字符串W.
//
// pStringXML: 字符串.
//
// pFileName: 资源文件名.
func (a *App) LoadResourceFromStringW(pStringXML string, pFileName string) bool {
	return xc.XC_LoadResourceFromStringW(pStringXML, pFileName)
}

// 炫彩_W2A.
//
// pValue: 参数.
func (a *App) WtoA(pValue string) uintptr {
	return xc.XC_wtoa(pValue)
}

// 炫彩_A2W.
//
// pValue: 参数.
func (a *App) AtoW(pValue uintptr) string {
	return xc.XC_atow(pValue)
}

// 炫彩_UTF8到文本W.
//
// pUtf8: 参数.
func (a *App) Utf8toW(pUtf8 uintptr) string {
	return xc.XC_utf8tow(pUtf8)
}

// 炫彩_UTF8到文本W扩展.
//
// pUtf8: utf8字符串指针.
//
// length: utf8字符串长度.
func (a *App) Utf8toWEx(pUtf8 uintptr, length int) string {
	return xc.XC_utf8towEx(pUtf8, length)
}

// 炫彩_UTF8到文本A.
//
// pUtf8: utf8字符串指针.
func (a *App) Utf8toA(pUtf8 uintptr) uintptr {
	return xc.XC_utf8toa(pUtf8)
}

// 炫彩_文本A到UTF8.
//
// pValue: 参数.
func (a *App) AtoUtf8(pValue uintptr) uintptr {
	return xc.XC_atoutf8(pValue)
}

// 炫彩_文本W到UTF8.
//
// pValue: 字符串.
func (a *App) WtoUtf8(pValue string) uintptr {
	return xc.XC_wtoutf8(pValue)
}

// 炫彩_文本W到UTF8扩展.
//
// pValue: 字符串.
//
// length: 字符串长度.
func (a *App) WtoUtf8Ex(pValue string, length int) uintptr {
	return xc.XC_wtoutf8Ex(pValue, length)
}

// 炫彩_U2A, 返回写入接收缓冲区字节数量.
//
// pIn: 待转换的Unicode字符串.
//
// inLen: pIn字符数量.
//
// pOut: 指向接收转换后的Ansi字符串缓冲区指针.
//
// outLen: pOut缓冲区大小, 字节单位.
func (a *App) UnicodeToAnsi(pIn string, inLen int, pOut uintptr, outLen int) int {
	return xc.XC_UnicodeToAnsi(pIn, inLen, pOut, outLen)
}

// 炫彩_A2U, 返回写入接收缓冲区字符数量.
//
// pIn: 指向待转换的Ansi字符串指针.
//
// inLen: pIn字符数量.
//
// pOut: 指向接收转换后的Unicode字符串缓冲区指针.
//
// outLen: pOut缓冲区大小,字符wchar_t单位.
func (a *App) AnsiToUnicode(pIn uintptr, inLen int, pOut *string, outLen int) int {
	return xc.XC_AnsiToUnicode(pIn, inLen, pOut, outLen)
}

// 炫彩_打印调试信息, 打印调试信息到文件xcgui_debug.txt.
//
// level: 级别.
//
// pInfo: 信息.
func (a *App) Print(level int, pInfo string) int {
	return xc.XDebug_Print(level, pInfo)
}

/*
// 炫彩_加载资源文件从字符串.
//
// pStringXML: 字符串.
//
// pFileName: 资源文件名.
func (a *App) LoadResourceFromString(pStringXML string, pFileName string) bool {
	return xc.XC_LoadResourceFromString(pStringXML, pFileName)
}

// 炫彩_加载样式文件从字符串.
//
// pFileName: 样式文件名, 用于打印错误文件和定位关联资源文件位置.
//
// pString: 字符串.
func (a *App) LoadStyleFromString(pFileName string, pString string) bool {
	return xc.XC_LoadStyleFromString(pFileName, pString)
} */

// 炫彩_取D2D工厂, 开启D2D有效, 返回 ID2D1Factory* .
func (a *App) GetD2dFactory() int {
	return xc.XC_GetD2dFactory()
}

// 炫彩_取DWrite工厂, 开启D2D有效, 返回 IDWriteFactory* .
func (a *App) GetDWriteFactory() int {
	return xc.XC_GetDWriteFactory()
}

// 炫彩_取WIC工厂, 开启D2D有效, 返回 IWICImagingFactory* .
func (a *App) GetWicFactory() int {
	return xc.XC_GetWicFactory()
}

// 炫彩_置D2D文本渲染模式.
//
// mode: 渲染模式, XC_DWRITE_RENDERING_MODE_ .
func (a *App) SetD2dTextRenderingMode(mode xcc.XC_DWRITE_RENDERING_MODE_) int {
	return xc.XC_SetD2dTextRenderingMode(mode)
}

// 炫彩_是否启用了D2D.
func (a *App) IsEnableD2D() bool {
	return xc.XC_IsEnableD2D()
}

// 炫彩_加载样式文件从字符串W.
//
// pFileName: 样式文件名.
//
// pString: 字符串.
func (a *App) LoadStyleFromStringW(pFileName string, pString string) bool {
	return xc.XC_LoadStyleFromStringW(pFileName, pString)
}

// 炫彩_显示边界.
//
// bShow: 是否显示.
func (a *App) ShowSvgFrame(bShow bool) int {
	return xc.XC_ShowSvgFrame(bShow)
}

// 通知消息_弹出, 未实现, 预留接口.
//
// position: 位置, Position_Flag_.
//
// pTitle: 标题.
//
// pText: 内容.
//
// hIcon: 图标.
//
// skin: 外观类型, NotifyMsg_Skin_.
func (a *App) NotifyMsg_Popup(position xcc.Position_Flag_, pTitle, pText string, hIcon int, skin xcc.NotifyMsg_Skin_) int {
	return xc.XNotifyMsg_Popup(position, pTitle, pText, hIcon, skin)
}

// 通知消息_弹出扩展, 未实现, 预留接口.
//
// position: 位置, Position_Flag_.
//
// pTitle: 标题.
//
// pText: 内容.
//
// hIcon: 图标.
//
// skin: 外观类型, NotifyMsg_Skin_.
//
// bBtnClose: 是否启用关闭按钮.
//
// bAutoClose: 是否自动关闭.
//
// nWidth: 自定义宽度, -1(使用默认值).
//
// nHeight: 自定义高度, -1(使用默认值).
func (a *App) NotifyMsg_PopupEx(position xcc.Position_Flag_, pTitle, pText string, hIcon int, skin xcc.NotifyMsg_Skin_, bBtnClose, bAutoClose bool, nWidth, nHeight int) int {
	return xc.XNotifyMsg_PopupEx(position, pTitle, pText, hIcon, skin, bBtnClose, bAutoClose, nWidth, nHeight)
}

// 通知消息_置持续时间.
//
// hWindow: 通知消息所属窗口句柄, 如果未指定那么认为是桌面通知消息.
//
// duration: 持续时间.
func (a *App) NotifyMsg_SetDuration(hWindow, duration int) int {
	return xc.XNotifyMsg_SetDuration(hWindow, duration)
}

// 通知消息_置父边距 设置通知消息与父对象的四边间隔.
//
// hWindow: 通知消息所属窗口句柄, 如果未指定那么认为是桌面通知消息.
//
// left: 左侧间隔, 未实现, 预留功能.
//
// top: 顶部间隔.
//
// right: 右侧间隔.
//
// bottom: 底部间隔, 未实现, 预留功能.
func (a *App) NotifyMsg_SetParentMargin(hWindow, left, top, right, bottom int) int {
	return xc.XNotifyMsg_SetParentMargin(hWindow, left, top, right, bottom)
}

// 通知消息_置标题高度.
//
// hWindow: 通知消息所属窗口句柄, 如果未指定那么认为是桌面通知消息.
//
// nHeight: 高度.
func (a *App) NotifyMsg_SetCaptionHeight(hWindow, nHeight int) int {
	return xc.XNotifyMsg_SetCaptionHeight(hWindow, nHeight)
}

// 通知消息_置宽度, 设置默认宽度.
//
// hWindow: 通知消息所属窗口句柄, 如果未指定那么认为是桌面通知消息.
//
// nWidth: 宽度.
func (a *App) NotifyMsg_SetWidth(hWindow, nWidth int) int {
	return xc.XNotifyMsg_SetWidth(hWindow, nWidth)
}

// 通知消息_置间隔.
//
// hWindow: 通知消息所属窗口句柄, 如果未指定那么认为是桌面通知消息.
//
// nSpace: 间隔大小.
func (a *App) NotifyMsg_SetSpace(hWindow, nSpace int) int {
	return xc.XNotifyMsg_SetSpace(hWindow, nSpace)
}

// 通知消息_置边大小, 设置通知消息面板边大小.
//
// hWindow: 通知消息所属窗口句柄, 如果未指定那么认为是桌面通知消息.
//
// left: 左边.
//
// top: 顶边.
//
// right: 右边.
//
// bottom: 底边.
func (a *App) NotifyMsg_SetBorderSize(hWindow, left, top, right, bottom int) int {
	return xc.XNotifyMsg_SetBorderSize(hWindow, left, top, right, bottom)
}

// 炫彩_启用自动DPI. 当启用后, 创建窗口时自动检测DPI调整UI缩放, 处理DPI改变消息; 禁用后,当DPI改变,需要手动设置窗口DPI.
//
// bEnabel: 是否启用.
func (a *App) EnableAutoDPI(bEnabel bool) int {
	return xc.XC_EnableAutoDPI(bEnabel)
}

// 炫彩_置窗口图标. 全局窗口图标, 所有未设置图标的窗口, 都将使用此默认图标.
//
// hImage: 图标句柄.
func (a *App) SetWindowIcon(hImage int) int {
	return xc.XC_SetWindowIcon(hImage)
}

// 炫彩_启用DPI.
//
// 为go程序启用DPI的几种方式:
//  1. 使用程序清单文件.
//  2. 调用此函数.
//  3. 自己调用DPI函数.
//
// 参考[MSDN](https://learn.microsoft.com/zh-cn/windows/win32/hidpi/setting-the-default-dpi-awareness-for-a-process)
//
// bEnabel: 是否启用.
func (a *App) EnableDPI(bEnabel bool) bool {
	return xc.XC_EnableDPI(bEnabel)
}

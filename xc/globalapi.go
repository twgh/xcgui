package xc

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"

	"github.com/twgh/xcgui/xcc"
)

// XInitXCGUI 炫彩_初始化.
//
//	@Description 在调用本函数之前请先调用 xc.LoadXCGUI().
//	@param bD2D 是否启用D2D.
//	@return bool
func XInitXCGUI(bD2D bool) bool {
	r, _, _ := xInitXCGUI.Call(common.BoolPtr(bD2D))
	return r != 0
}

// XRunXCGUI 炫彩_运行, 运行消息循环, 当炫彩窗口数量为0时退出.
//
//	@return int
func XRunXCGUI() int {
	r, _, _ := xRunXCGUI.Call()
	return int(r)
}

// XExitXCGUI 炫彩_退出, 退出界面库释放资源.
//
//	@return int
func XExitXCGUI() int {
	r, _, _ := xExitXCGUI.Call()
	return int(r)
}

// XC_DebugToFileInfo 炫彩_输出调试信息到文件, 打印调试信息到文件xcgui_debug.txt.
//
//	@param pInfo 文本.
func XC_DebugToFileInfo(pInfo string) {
	xC_DebugToFileInfo.Call(XC_wtoa(pInfo))
}

// XC_SetActivateTopWindow 炫彩_激活窗口, 激活当前进程最上层窗口.
//
//	@return bool
func XC_SetActivateTopWindow() bool {
	r, _, _ := xC_SetActivateTopWindow.Call()
	return r != 0
}

// XC_GetDefaultFont 炫彩_取默认字体.
//
//	@return int 字体句柄.
func XC_GetDefaultFont() int {
	r, _, _ := xC_GetDefaultFont.Call()
	return int(r)
}

// XC_MessageBox 炫彩_消息框.
//
//	@param pTitle 标题.
//	@param pText 内容文本.
//	@param nFlags 标识: xcc.MessageBox_Flag_.
//	@param hWndParent 父窗口句柄(真实的窗口句柄).
//	@param XCStyle xcc.Window_Style_.
//	@return xcc.MessageBox_Flag_ , 返回: xcc.MessageBox_Flag_Ok: 点击确定按钮退出. xcc.MessageBox_Flag_Cancel: 点击取消按钮退出. xcc.MessageBox_Flag_Other: 其他方式退出.
func XC_MessageBox(pTitle, pText string, nFlags xcc.MessageBox_Flag_, hWndParent uintptr, XCStyle xcc.Window_Style_) xcc.MessageBox_Flag_ {
	r, _, _ := xC_MessageBox.Call(common.StrPtr(pTitle), common.StrPtr(pText), uintptr(nFlags), hWndParent, uintptr(XCStyle))
	return xcc.MessageBox_Flag_(r)
}

// XMsg_Create 消息框_创建, 此窗口是一个模态窗口, 弹出窗口请调用 xc.XModalWnd_DoModal().
//
//	@param pTitle 标题.
//	@param pText 内容文本.
//	@param nFlags 标识: xcc.MessageBox_Flag_.
//	@param hWndParent 父窗口句柄(真实的窗口句柄).
//	@param XCStyle xcc.Window_Style_.
//	@return int 返回消息框窗口句柄.
func XMsg_Create(pTitle, pText string, nFlags xcc.MessageBox_Flag_, hWndParent uintptr, XCStyle xcc.Window_Style_) int {
	r, _, _ := xMsg_Create.Call(common.StrPtr(pTitle), common.StrPtr(pText), uintptr(nFlags), hWndParent, uintptr(XCStyle))
	return int(r)
}

// XMsg_CreateEx 消息框_创建扩展, 此窗口是一个模态窗口, 弹出窗口请调用 xc.XModalWnd_DoModal().
//
//	@param dwExStyle 窗口扩展样式.
//	@param dwStyle 窗口样式.
//	@param lpClassName 窗口类名.
//	@param pTitle 标题.
//	@param pText 内容文本.
//	@param nFlags 标识: xcc.MessageBox_Flag_ .
//	@param hWndParent 父窗口句柄(真实的窗口句柄).
//	@param XCStyle xcc.Window_Style_ .
//	@return int 消息框窗口句柄.
func XMsg_CreateEx(dwExStyle int, dwStyle int, lpClassName, pTitle, pText string, nFlags xcc.MessageBox_Flag_, hWndParent uintptr, XCStyle xcc.Window_Style_) int {
	r, _, _ := xMsg_CreateEx.Call(uintptr(dwExStyle), uintptr(dwStyle), common.StrPtr(lpClassName), common.StrPtr(pTitle), common.StrPtr(pText), uintptr(nFlags), hWndParent, uintptr(XCStyle))
	return int(r)
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
func XC_SendMessage(hWindow int, msg uint32, wParam, lParam uint) uint {
	r, _, _ := xC_SendMessage.Call(uintptr(hWindow), uintptr(msg), uintptr(wParam), uintptr(lParam))
	return uint(r)
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
func XC_PostMessage(hWindow int, msg uint32, wParam int32, lParam int32) bool {
	r, _, _ := xC_PostMessage.Call(uintptr(hWindow), uintptr(msg), uintptr(wParam), uintptr(lParam))
	return r != 0
}

// XC_CallUiThread 炫彩_调用界面线程, 调用UI线程, 设置回调函数, 在回调函数里操作UI.
//
//	@Description: 回调函数尽量不要使用匿名函数, 使用匿名函数意味着你每次都在创建1个新的回调, 超过2000个时, 程序必将panic.
//	如果使用 xc.XC_CallUiThreadEx 和 xc.XC_CallUiThreader 则没有此限制.
//	@param pCall 回调函数.
//	@param data 传进回调函数的用户自定义数据.
//	@return int
func XC_CallUiThread(pCall func(data int) int, data int) int {
	r, _, _ := xC_CallUiThread.Call(syscall.NewCallback(pCall), uintptr(data))
	return int(r)
}

// 炫彩_判断元素, 判断是否为元素句柄.
//
// hEle: 元素句柄.
func XC_IsHELE(hEle int) bool {
	r, _, _ := xC_IsHELE.Call(uintptr(hEle))
	return r != 0
}

// 炫彩_判断窗口, 判断是否为窗口句柄.
//
// hWindow: 窗口句柄.
func XC_IsHWINDOW(hWindow int) bool {
	r, _, _ := xC_IsHWINDOW.Call(uintptr(hWindow))
	return r != 0
}

// 炫彩_判断形状对象, 判断是否为形状对象.
//
// hShape: 形状对象句柄.
func XC_IsShape(hShape int) bool {
	r, _, _ := xC_IsShape.Call(uintptr(hShape))
	return r != 0
}

// 炫彩_判断句柄包含类型, 判断句柄是否拥有该类型.
//
// hXCGUI: 炫彩句柄.
//
// nType: 句柄类型, XC_OBJECT_TYPE, 以XC_开头的常量.
func XC_IsHXCGUI(hXCGUI int, nType xcc.XC_OBJECT_TYPE) bool {
	r, _, _ := xC_IsHXCGUI.Call(uintptr(hXCGUI), uintptr(nType))
	return r != 0
}

// 炫彩_转换HWND到HWINDOW, 通过窗口HWND句柄获取HWINDOW句柄.
//
// hWnd: 窗口真实句柄HWND.
func XC_hWindowFromHWnd(hWnd uintptr) int {
	r, _, _ := xC_hWindowFromHWnd.Call(hWnd)
	return int(r)
}

// 炫彩_置属性, 设置对象属性.
//
// hXCGUI: 对象句柄.
//
// pName: 属性名.
//
// pValue: 属性值.
func XC_SetProperty(hXCGUI int, pName string, pValue string) bool {
	r, _, _ := xC_SetProperty.Call(uintptr(hXCGUI), common.StrPtr(pName), common.StrPtr(pValue))
	return r != 0
}

// 炫彩_取属性, 获取对象属性, 返回属性值.
//
// hXCGUI: 对象句柄.
//
// pName: 属性名.
func XC_GetProperty(hXCGUI int, pName string) string {
	r, _, _ := xC_GetProperty.Call(uintptr(hXCGUI), common.StrPtr(pName))
	return common.UintPtrToString(r)
}

// 炫彩_注册窗口类名, 如果是在DLL中使用, 那么DLL卸载时需要注销窗口类名, 否则DLL卸载后, 类名所指向的窗口过程地址失效.
//
// pClassName: 类名.
func XC_RegisterWindowClassName(pClassName string) bool {
	r, _, _ := xC_RegisterWindowClassName.Call(common.StrPtr(pClassName))
	return r != 0
}

// 炫彩_判断滚动视图扩展元素, 判断元素是否从滚动视图元素扩展的新元素, 包含滚动视图元素.
//
// hEle: 元素句柄.
func XC_IsSViewExtend(hEle int) bool {
	r, _, _ := xC_IsSViewExtend.Call(uintptr(hEle))
	return r != 0
}

// 炫彩_取对象类型, 获取句柄类型, 返回: XC_OBJECT_TYPE.
//
// hXCGUI: 炫彩对象句柄.
func XC_GetObjectType(hXCGUI int) xcc.XC_OBJECT_TYPE {
	r, _, _ := xC_GetObjectType.Call(uintptr(hXCGUI))
	return xcc.XC_OBJECT_TYPE(r)
}

// 炫彩_取对象从ID, 通过ID获取对象句柄, 不包括窗口对象.
//
// hWindow: 所属窗口句柄.
//
// nID: ID值.
func XC_GetObjectByID(hWindow int, nID int) int {
	r, _, _ := xC_GetObjectByID.Call(uintptr(hWindow), uintptr(nID))
	return int(r)
}

// 炫彩_取对象从ID名称, 通过ID名称获取对象句柄.
//
// hWindow: 所属窗口句柄.
//
// pName: ID名称.
func XC_GetObjectByIDName(hWindow int, pName string) int {
	r, _, _ := xC_GetObjectByIDName.Call(uintptr(hWindow), common.StrPtr(pName))
	return int(r)
}

// 炫彩_取对象从UID, 通过UID获取对象句柄, 不包括窗口对象.
//
// nUID: UID值.
func XC_GetObjectByUID(nUID int) int {
	r, _, _ := xC_GetObjectByUID.Call(uintptr(nUID))
	return int(r)
}

// 炫彩_取对象从UID名称, 通过UID名称获取对象句柄.
//
// pName: UID名称.
func XC_GetObjectByUIDName(pName string) int {
	r, _, _ := xC_GetObjectByUIDName.Call(common.StrPtr(pName))
	return int(r)
}

// 炫彩_取对象从名称, 通过name获取对象句柄.
//
// pName: name名称.
func XC_GetObjectByName(pName string) int {
	r, _, _ := xC_GetObjectByName.Call(common.StrPtr(pName))
	return int(r)
}

// 炫彩_置绘制频率, 设置UI的最小重绘频率.
//
// nMilliseconds: 重绘最小时间间隔, 单位毫秒.
func XC_SetPaintFrequency(nMilliseconds int) {
	xC_SetPaintFrequency.Call(uintptr(nMilliseconds))
}

// 炫彩_置文本渲染质量, 设置文本渲染质量.
//
// nType: 参见GDI+ TextRenderingHint 定义.
func XC_SetTextRenderingHint(nType int) {
	xC_SetTextRenderingHint.Call(uintptr(nType))
}

// 炫彩_启用GDI绘制文本, 将影响到以下函数: XDraw_TextOut, XDraw_TextOutEx, XDraw_TextOutA.
//
// bEnable: 是否启用.
func XC_EnableGdiDrawText(bEnable bool) {
	xC_EnableGdiDrawText.Call(common.BoolPtr(bEnable))
}

// 炫彩_判断矩形相交, 判断两个矩形是否相交及重叠.
//
// pRect1: 矩形1.
//
// pRect2: 矩形2.
func XC_RectInRect(pRect1 *RECT, pRect2 *RECT) bool {
	r, _, _ := xC_RectInRect.Call(uintptr(unsafe.Pointer(pRect1)), uintptr(unsafe.Pointer(pRect2)))
	return r != 0
}

// 炫彩_组合矩形, 组合两个矩形区域.
//
// pDest: 新的矩形区域.
//
// pSrc1: 源矩形1.
//
// pSrc2: 源矩形2.
func XC_CombineRect(pDest *RECT, pSrc1 *RECT, pSrc2 *RECT) {
	xC_CombineRect.Call(uintptr(unsafe.Pointer(pDest)), uintptr(unsafe.Pointer(pSrc1)), uintptr(unsafe.Pointer(pSrc2)))
}

// 炫彩_显示布局边界, 显示布局对象边界.
//
// bShow: 是否显示.
func XC_ShowLayoutFrame(bShow bool) {
	xC_ShowLayoutFrame.Call(common.BoolPtr(bShow))
}

// 炫彩_启用debug文件.
//
// bEnable: 是否启用.
func XC_EnableDebugFile(bEnable bool) {
	xC_EnableDebugFile.Call(common.BoolPtr(bEnable))
}

// 炫彩_启用资源监视器.
//
// bEnable: 是否启用.
func XC_EnableResMonitor(bEnable bool) {
	xC_EnableResMonitor.Call(common.BoolPtr(bEnable))
}

// 炫彩_置布局边界颜色.
//
// color: ABGR 颜色值.
func XC_SetLayoutFrameColor(color int) int {
	r, _, _ := xC_SetLayoutFrameColor.Call(uintptr(color))
	return int(r)
}

// 炫彩_启用错误弹窗, 启用错误弹出, 通过该接口可以设置遇到严重错误时不弹出消息提示框.
//
// bEnabel: 是否启用.
func XC_EnableErrorMessageBox(bEnabel bool) int {
	r, _, _ := xC_EnableErrorMessageBox.Call(common.BoolPtr(bEnabel))
	return int(r)
}

// 炫彩_启用自动退出程序, 启动或禁用自动退出程序, 当检测到所有用户创建的窗口都关闭时, 自动退出程序; 可调用 XC_PostQuitMessage() 手动退出程序.
//
// bEnabel: 是否启用.
func XC_EnableAutoExitApp(bEnabel bool) int {
	r, _, _ := xC_EnableAutoExitApp.Call(common.BoolPtr(bEnabel))
	return int(r)
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
func XC_GetTextSize(pString string, length int, hFontX int, pOutSize *SIZE) int {
	r, _, _ := xC_GetTextSize.Call(common.StrPtr(pString), uintptr(length), uintptr(hFontX), uintptr(unsafe.Pointer(pOutSize)))
	return int(r)
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
func XC_GetTextShowSize(pString string, length int, hFontX int, pOutSize *SIZE) int {
	r, _, _ := xC_GetTextShowSize.Call(common.StrPtr(pString), uintptr(length), uintptr(hFontX), uintptr(unsafe.Pointer(pOutSize)))
	return int(r)
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
func XC_GetTextShowSizeEx(pString string, length int, hFontX int, nTextAlign xcc.TextFormatFlag_, pOutSize *SIZE) int {
	r, _, _ := xC_GetTextShowSizeEx.Call(common.StrPtr(pString), uintptr(length), uintptr(hFontX), uintptr(nTextAlign), uintptr(unsafe.Pointer(pOutSize)))
	return int(r)
}

// XC_GetTextShowRect 炫彩_取文本显示矩形.
//
//	@param pString 字符串.
//	@param length 字符串长度.
//	@param hFontX 字体.
//	@param nTextAlign 文本对齐: xcc.TextFormatFlag_.
//	@param width 最大宽度.
//	@param pOutSize 接收返回大小.
//	@return int
func XC_GetTextShowRect(pString string, length int, hFontX int, nTextAlign xcc.TextFormatFlag_, width int, pOutSize *SIZE) int {
	r, _, _ := xC_GetTextShowRect.Call(common.StrPtr(pString), uintptr(length), uintptr(hFontX), uintptr(nTextAlign), uintptr(width), uintptr(unsafe.Pointer(pOutSize)))
	return int(r)
}

// 炫彩_置默认字体.
//
// hFontX: 炫彩字体句柄.
func XC_SetDefaultFont(hFontX int) int {
	r, _, _ := xC_SetDefaultFont.Call(uintptr(hFontX))
	return int(r)
}

// 炫彩_添加搜索路径, 添加文件搜索路径, 默认路径为exe目录和程序当前运行目录.
//
// pPath: 文件夹.
func XC_AddFileSearchPath(pPath string) int {
	r, _, _ := xC_AddFileSearchPath.Call(common.StrPtr(pPath))
	return int(r)
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
func XC_InitFont(pFont *LOGFONTW, pName string, size int, bBold bool, bItalic bool, bUnderline bool, bStrikeOut bool) int {
	r, _, _ := xC_InitFont.Call(uintptr(unsafe.Pointer(pFont)), common.StrPtr(pName), uintptr(size), common.BoolPtr(bBold), common.BoolPtr(bItalic), common.BoolPtr(bUnderline), common.BoolPtr(bStrikeOut))
	return int(r)
}

// 炫彩_分配内存, 在UI库中申请内存, 返回: 内存首地址.
//
// size: 大小, 字节为单位.
func XC_Malloc(size int) int {
	r, _, _ := xC_Malloc.Call(uintptr(size))
	return int(r)
}

// 炫彩_释放内存, 在UI库中释放内存.
//
// p: 内存首地址.
func XC_Free(p int) int {
	r, _, _ := xC_Free.Call(uintptr(p))
	return int(r)
}

// 炫彩_弹框, 弹出提示框.
//
// pTitle: 提示框标题.
//
// pText: 提示内容.
func XC_Alert(pTitle, pText string) int {
	r, _, _ := xC_Alert.Call(common.StrPtr(pTitle), common.StrPtr(pText))
	return int(r)
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
func XC_Sys_ShellExecute(hwnd uintptr, lpOperation string, lpFile string, lpParameters string, lpDirectory string, nShowCmd xcc.SW_) uintptr {
	r, _, _ := xC_Sys_ShellExecute.Call(hwnd, common.StrPtr(lpOperation), common.StrPtr(lpFile), common.StrPtr(lpParameters), common.StrPtr(lpDirectory), uintptr(nShowCmd))
	return r
}

// 炫彩_载入动态库, 系统API LoadLibrary, 返回动态库模块句柄.
//
// lpFileName: 文件名.
func XC_LoadLibrary(lpFileName string) uintptr {
	r, _, _ := xC_LoadLibrary.Call(common.StrPtr(lpFileName))
	return r
}

// 炫彩_取动态库中函数地址, 系统API GetProcAddress, 返回函数地址.
//
// hModule: 动态库模块句柄.
//
// lpProcName: 函数名.
func XC_GetProcAddress(hModule uintptr, lpProcName string) uintptr {
	r, _, _ := xC_GetProcAddress.Call(hModule, XC_wtoa(lpProcName))
	return r
}

// 炫彩_释放动态库, 系统API FreeLibrary.
//
// hModule: 动态库模块句柄.
func XC_FreeLibrary(hModule uintptr) bool {
	r, _, _ := xC_FreeLibrary.Call(hModule)
	return r != 0
}

// 炫彩_加载DLL, 返回DLL模块句柄. 加载指定DLL, 并且调用DLL中函数LoadDll(), DLL中导出函数格式: int WINAPI LoadDll().
//
// pDllFileName: DLL文件名.
func XC_LoadDll(pDllFileName string) uintptr {
	r, _, _ := xC_LoadDll.Call(common.StrPtr(pDllFileName))
	return r
}

// 炫彩_PostQuitMessage, 发送WM_QUIT消息退出消息循环.
//
// nExitCode: 退出码.
func XC_PostQuitMessage(nExitCode int) int {
	r, _, _ := xC_PostQuitMessage.Call(uintptr(nExitCode))
	return int(r)
}

// 炫彩_置D2D文本渲染模式.
//
// mode: 渲染模式, XC_DWRITE_RENDERING_MODE_ .
func XC_SetD2dTextRenderingMode(mode xcc.XC_DWRITE_RENDERING_MODE_) int {
	r, _, _ := xC_SetD2dTextRenderingMode.Call(uintptr(mode))
	return int(r)
}

// 炫彩_是否启用了D2D.
func XC_IsEnableD2D() bool {
	r, _, _ := xC_IsEnableD2D.Call()
	return r != 0
}

// 炫彩_W2A.
//
// pValue: 参数.
func XC_wtoa(pValue string) uintptr {
	r, _, _ := xC_wtoa.Call(common.StrPtr(pValue))
	return r
}

// 炫彩_A2W.
//
// pValue: 参数.
func XC_atow(pValue uintptr) string {
	r, _, _ := xC_atow.Call(pValue)
	return common.UintPtrToString(r)
}

// 炫彩_UTF8到文本W.
//
// pUtf8: 参数.
func XC_utf8tow(pUtf8 uintptr) string {
	r, _, _ := xC_utf8tow.Call(pUtf8)
	return common.UintPtrToString(r)
}

// 炫彩_UTF8到文本W扩展.
//
// pUtf8: utf8字符串指针.
//
// length: utf8字符串长度.
func XC_utf8towEx(pUtf8 uintptr, length int) string {
	r, _, _ := xC_utf8towEx.Call(pUtf8, uintptr(length))
	return common.UintPtrToString(r)
}

// 炫彩_UTF8到文本A.
//
// pUtf8: utf8字符串指针.
func XC_utf8toa(pUtf8 uintptr) uintptr {
	r, _, _ := xC_utf8toa.Call(pUtf8)
	return r
}

// 炫彩_文本A到UTF8.
//
// pValue: 参数.
func XC_atoutf8(pValue uintptr) uintptr {
	r, _, _ := xC_atoutf8.Call(pValue)
	return r
}

// 炫彩_文本W到UTF8.
//
// pValue: 字符串.
func XC_wtoutf8(pValue string) uintptr {
	r, _, _ := xC_wtoutf8.Call(common.StrPtr(pValue))
	return r
}

// 炫彩_文本W到UTF8扩展.
//
// pValue: 字符串.
//
// length: 字符串长度.
func XC_wtoutf8Ex(pValue string, length int) uintptr {
	r, _, _ := xC_wtoutf8Ex.Call(common.StrPtr(pValue), uintptr(length))
	return r
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
func XC_UnicodeToAnsi(pIn string, inLen int, pOut uintptr, outLen int) int {
	r, _, _ := xC_UnicodeToAnsi.Call(common.StrPtr(pIn), uintptr(inLen), pOut, uintptr(outLen))
	return int(r)
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
func XC_AnsiToUnicode(pIn uintptr, inLen int, pOut *string, outLen int) int {
	buf := make([]uint16, outLen)
	r, _, _ := xC_AnsiToUnicode.Call(pIn, uintptr(inLen), common.Uint16SliceDataPtr(&buf), uintptr(outLen))
	*pOut = syscall.UTF16ToString(buf[0:])
	return int(r)
}

// 炫彩_打印调试信息, 打印调试信息到文件xcgui_debug.txt.
//
// level: 级别.
//
// pInfo: 信息.
func XDebug_Print(level int, pInfo string) int {
	r, _, _ := xDebug_Print.Call(uintptr(level), XC_wtoa(pInfo))
	return int(r)
}

// 炫彩_显示边界.
//
// bShow: 是否显示.
func XC_ShowSvgFrame(bShow bool) int {
	r, _, _ := xC_ShowSvgFrame.Call(common.BoolPtr(bShow))
	return int(r)
}

// 炫彩_启用自动DPI. 当启用后, 创建窗口时自动检测DPI调整UI缩放, 处理DPI改变消息; 禁用后,当DPI改变,需要手动设置窗口DPI.
//
// bEnabel: 是否启用.
func XC_EnableAutoDPI(bEnabel bool) int {
	r, _, _ := xC_EnableAutoDPI.Call(common.BoolPtr(bEnabel))
	return int(r)
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
func XC_EnableDPI(bEnabel bool) bool {
	r, _, _ := xC_EnableDPI.Call(common.BoolPtr(bEnabel))
	return r != 0
}

// 炫彩_置窗口图标. 全局窗口图标, 所有未设置图标的窗口, 都将使用此默认图标.
//
// hImage: 图标句柄.
func XC_SetWindowIcon(hImage int) int {
	r, _, _ := xC_SetWindowIcon.Call(uintptr(hImage))
	return int(r)
}

/* // 炫彩_打印调试信息, 打印调试信息到文件xcgui_debug.txt.[无效]
//
// pString: 字符串.
func XDebug_OutputDebugStringW(pString string) int {
	r, _, _ := xDebug_OutputDebugStringW.Call(common.StrPtr(pString))
	return int(r)
}

// 炫彩_设置debug输出编码方式为utf8.[无效]
//
// bUTF8: 是否开启utf8编码.
func XDebug_Set_OutputDebugString_UTF8(bUTF8 bool) int {
	r, _, _ := xDebug_Set_OutputDebugString_UTF8.Call(common.BoolPtr(bUTF8))
	return int(r)
}

// 炫彩_整数到文本A.
//
// nValue: 参数.
func XC_itoa(nValue int) int {
	r, _, _ := xC_itoa.Call(uintptr(nValue))
	return int(r)
}

// 炫彩_整数到文本W.
//
// nValue: 参数.
func XC_itow(nValue int) int {
	r, _, _ := xC_itow.Call(uintptr(nValue))
	return int(r)
}

// 炫彩_浮点数到文本A.
//
// fValue: 参数.
func XC_ftoa(fValue int) int {
	r, _, _ := xC_ftoa.Call(uintptr(fValue))
	return int(r)
}

// 炫彩_浮点数到文本W.
//
// fValue: 参数.
func XC_ftow(fValue int) int {
	r, _, _ := xC_ftow.Call(uintptr(fValue))
	return int(r)
}
*/

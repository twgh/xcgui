package wapi

import (
	"github.com/twgh/xcgui/common"
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

var (
	// Library.
	shell32 = syscall.NewLazyDLL("shell32.dll")

	// Functions.
	dragQueryFileW       = shell32.NewProc("DragQueryFileW")
	dragFinish           = shell32.NewProc("DragFinish")
	dragQueryPoint       = shell32.NewProc("DragQueryPoint")
	shellExecuteW        = shell32.NewProc("ShellExecuteW")
	sHBrowseForFolderW   = shell32.NewProc("SHBrowseForFolderW")
	sHGetPathFromIDListW = shell32.NewProc("SHGetPathFromIDListW")
)

// DragQueryFileW 检索由成功的拖放操作产生的文件路径.
//
//	@Description 详见: https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-DragQueryFileW.
//	@param hDrop 句柄.
//	@param iFile 文件索引.
//	@param lpszFile 返回的文件路径.
//	@param cch 接收的文件路径的字符数, 通常为260.
//	@return int 返回文件路径的字符数.
func DragQueryFileW(hDrop uintptr, iFile uint32, lpszFile *string, cch uint32) int {
	buf := make([]uint16, cch)
	r, _, _ := dragQueryFileW.Call(hDrop, uintptr(iFile), common.Uint16SliceDataPtr(&buf), uintptr(cch))
	*lpszFile = syscall.UTF16ToString(buf[0:])
	return int(r)
}

// DragFinish 释放系统分配用于将文件名传输到应用程序的内存.
//
//	@Description 详见: https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-DragFinish.
//	@param hDrop 句柄.
func DragFinish(hDrop uintptr) {
	dragFinish.Call(hDrop)
}

// DragQueryPoint 检索在拖放文件时鼠标指针的位置.
//
//	@Description 详见: https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-DragQueryPoint.
//	@param hDrop 句柄.
//	@param ppt 接收鼠标指针的坐标.
//	@return bool 如果拖放发生在窗口的客户区, 返回true；否则返回false.
func DragQueryPoint(hDrop uintptr, ppt *xc.POINT) bool {
	r, _, _ := dragQueryPoint.Call(hDrop, uintptr(unsafe.Pointer(ppt)))
	return r != 0
}

// ShellExecuteW 对指定文件执行操作.
//
//	@Description 详见: https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shellexecutew.
//	@param hwnd 用于显示 UI 或错误消息的父窗口的句柄。如果操作与窗口无关，则此值可以为0.
//	@param lpOperation 填“open”则打开lpFlie文档.
//	@param lpFile 想用关联的程序打印或打开的一个程序名或文件名.
//	@param lpParameters 如果lpFile是一个可执行文件，则这个字串包含了传递给执行程序的参数.
//	@param lpDirectory 想使用的默认路径完整路径.
//	@param nShowCmd 定义了如何显示启动程序的常数值, xcc.SW_.
//	@return int 如果函数成功，则返回大于32的值。如果函数失败，则返回指示失败原因的错误值.
func ShellExecuteW(hwnd uintptr, lpOperation, lpFile, lpParameters, lpDirectory string, nShowCmd xcc.SW_) int {
	r, _, _ := shellExecuteW.Call(hwnd, common.StrPtr(lpOperation), common.StrPtr(lpFile), common.StrPtr(lpParameters), common.StrPtr(lpDirectory), uintptr(nShowCmd))
	return int(r)
}

// BrowseInfoW 包含用于显示对话框的信息。
//
//	@Description 详见: https://docs.microsoft.com/zh-cn/windows/win32/api/shlobj_core/ns-shlobj_core-browseinfow.
type BrowseInfoW struct {
	HwndOwner      uintptr // 父窗口句柄
	PidlRoot       uintptr // 指定开始浏览的根文件夹的位置。只有命名空间层次结构中的指定文件夹及其子文件夹出现在对话框中。该成员可以为NULL；在这种情况下，将使用默认位置.
	PszDisplayName uintptr // 指向缓冲区的指针，用于接收用户选择的文件夹的显示名称。此缓冲区的大小假定为 260 个字符.
	LpszTitle      uintptr // 指向显示在对话框中树视图控件上方的以空字符结尾的字符串的指针。使用 common.StrPtr()函数生成.
	UlFlags        BIF_    // 指定对话框选项的标志。可以为0，也可以是 wapi.BIF_ 的组合.
	Lpfn           uintptr // 指向应用程序定义函数的指针，当事件发生时对话框调用该函数.
	LParam         uintptr // 对话框传递给回调函数的应用程序定义的值（如果在lpfn中指定） .
	IImage         int32   // 接收与所选文件夹关联的图像索引，存储在系统图像列表中。
}

// BIF_ 是指定对话框选项的标志.
type BIF_ uint32

const (
	BIF_RETURNONLYFSDIRS BIF_ = 0x00000001 // 仅返回文件系统目录。如果用户选择不属于文件系统的文件夹，则“确定”按钮将变灰。注意对于“\\server”项目以及“\\server\share”和目录项目，“确定”按钮仍处于启用  状态。但是，如果用户选择“\\server”项，则将 SHBrowseForFolderW 返回的 PIDL 传递给 SHGetPathFromIDListW 将失败。

	BIF_DONTGOBELOWDOMAIN BIF_ = 0x00000002 // 不要在对话框的树视图控件中包含域级别以下的网络文件夹。

	BIF_STATUSTEXT BIF_ = 0x00000004 // 在对话框中包含一个状态区域。回调函数可以通过向对话框发送消息来设置状态文本。指定 BIF_NEWDIALOGSTYLE 时不支持此标志。

	BIF_RETURNFSANCESTORS BIF_ = 0x00000008 // 仅返回文件系统祖先。祖先是命名空间层次结构中根文件夹下的子文件夹。如果用户选择了不属于文件系统的根文件夹的祖先，则“确定”按钮将变灰。

	BIF_EDITBOX BIF_ = 0x00000010 // 在允许用户键入项目名称的浏览对话框中包含一个编辑控件。

	BIF_VALIDATE BIF_ = 0x00000020 // 如果用户在编辑框中输入了无效的名称，浏览对话框将调用应用程序的BrowseCallbackProc并发送BFFM_VALIDATEFAILED消息。如果未指定 BIF_EDITBOX，则忽略此标志。

	BIF_NEWDIALOGSTYLE BIF_ = 0x00000040 // 使用新的用户界面。设置此标志为用户提供了一个可以调整大小的更大对话框。该对话框具有多项新功能，包括：对话框内的拖放功能、重新排序、快捷菜单、新文件夹、删除和其他快捷菜单命令。注意  如果 COM 是通过CoInitializeEx初始化并设置了 COINIT_MULTITHREADED 标志，如果传递了 BIF_NEWDIALOGSTYLE，则 SHBrowseForFolderW 将失败。

	BIF_BROWSEINCLUDEURLS BIF_ = 0x00000080 // 浏览对话框可以显示 URL. BIF_USENEWUI 和 BIF_BROWSEINCLUDEFILES 标志也必须设置。如果未设置这三个标志中的任何一个，浏览器对话框将拒绝 URL。即使设置了这些标志，只有在包含所选项目的文件夹支持 URL 时，浏览对话框才会显示 URL。当调用文件夹的IShellFolder::GetAttributesOf方法来请求所选项目的属性时，文件夹必须设置SFGAO_FOLDER属性标志。否则，浏览对话框将不会显示 URL。

	BIF_USENEWUI = BIF_EDITBOX | BIF_NEWDIALOGSTYLE // 使用新的用户界面，包括一个编辑框。这个标志相当于 BIF_EDITBOX | BIF_NEWDIALOGSTYLE。

	BIF_UAHINT BIF_ = 0x00000100 // 与 BIF_NEWDIALOGSTYLE 结合使用时，会在对话框中添加使用提示来代替编辑框。BIF_EDITBOX 会覆盖此标志。

	BIF_NONEWFOLDERBUTTON BIF_ = 0x00000200 // 不要在浏览对话框中包含新建文件夹按钮。

	BIF_NOTRANSLATETARGETS BIF_ = 0x00000400 // 当所选项目是快捷方式时，返回快捷方式本身的 PIDL 而不是其目标。

	BIF_BROWSEFORCOMPUTER BIF_ = 0x00001000 // 只退回电脑。如果用户选择计算机以外的任何东西，则“确定”按钮将变灰。

	BIF_BROWSEFOPRINTER BIF_ = 0x00002000 // 只允许选择打印机。如果用户选择打印机以外的任何东西，则“确定”按钮将变灰。在 Windows XP 和更高版本的系统中，最佳做法是使用 Windows XP 样式的对话框，将对话框的根设置为Printers and Faxes文件夹 (CSIDL_PRINTERS)。

	BIF_BROWSEINCLUDEFILES BIF_ = 0x00004000 // 浏览对话框显示文件和文件夹。

	BIF_SHAREABLE BIF_ = 0x00008000 // 浏览对话框可以显示远程系统上的共享资源。这适用于希望在本地系统上公开远程共享的应用程序。BIF_NEWDIALOGSTYLE 标志也必须设置。

	BIF_BROWSEFILEJUNCTIONS BIF_ = 0x00010000 // Windows 7 及更高版本。允许浏览文件夹连接，例如库或具有 .zip 文件扩展名的压缩文件。
)

// SHBrowseForFolderW 显示一个对话框，使用户能够选择文件夹。
//
//	@Description 详见: https://docs.microsoft.com/zh-cn/windows/win32/api/shlobj_core/nf-shlobj_core-shbrowseforfolderw.
//	@param browseInfo 指向 wapi.BrowseInfoW 结构的指针，该结构包含用于显示对话框的信息。
//	@return uintptr 返回一个 PIDL，它指定所选文件夹相对于命名空间根的位置。如果用户在对话框中选择取消按钮，则返回值为NULL。返回的 PIDL 可能是文件夹快捷方式而不是文件夹。
func SHBrowseForFolderW(browseInfo *BrowseInfoW) uintptr {
	r, _, _ := sHBrowseForFolderW.Call(uintptr(unsafe.Pointer(browseInfo)))
	return r
}

// SHGetPathFromIDListW 将 SHBrowseForFolderW 的返回值转换为文件路径。
//
//	@Description 详见: https://docs.microsoft.com/zh-cn/windows/win32/api/shlobj_core/nf-shlobj_core-shgetpathfromidlistw.
//	@param pidl SHBrowseForFolderW 的返回值.
//	@param pszPath 返回的文件路径。
//	@return bool
func SHGetPathFromIDListW(pidl uintptr, pszPath *string) bool {
	buf := make([]uint16, 260)
	r, _, _ := sHGetPathFromIDListW.Call(pidl, common.Uint16SliceDataPtr(&buf))
	*pszPath = syscall.UTF16ToString(buf)
	return r != 0
}

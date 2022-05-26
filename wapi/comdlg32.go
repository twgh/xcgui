package wapi

import (
	"syscall"
	"unsafe"
)

var (
	// Library.
	comdlg32 = syscall.NewLazyDLL("comdlg32.dll")

	// Functions.
	getOpenFileNameW = comdlg32.NewProc("GetOpenFileNameW")
	getSaveFileNameW = comdlg32.NewProc("GetSaveFileNameW")
	chooseColorW     = comdlg32.NewProc("ChooseColorW")
)

// OpenFileNameW 包含 GetOpenFileNameW 和 GetSaveFileNameW 函数用于初始化“打开”或“另存为”对话框的信息。用户关闭对话框后，系统在此结构中返回有关用户选择的信息。
//	详见: https://docs.microsoft.com/zh-cn/windows/win32/api/commdlg/ns-commdlg-openfilenamea.
type OpenFileNameW struct {
	// 结构体大小.
	//	ofn := wapi.OpenFileNameW{...}
	//	ofn.LStructSize = uint32(unsafe.Sizeof(ofn))
	LStructSize uint32

	// 拥有对话框的窗口句柄。此成员可以是任何有效的窗口句柄，或者如果对话框没有所有者，它可以为NULL 。
	HwndOwner int

	// 如果在Flags成员中设置了OFN_ENABLETEMPLATEHANDLE标志，则hInstance是包含对话框模板的内存对象的句柄。如果设置了OFN_ENABLETEMPLATE标志，则hInstance是一个模块句柄，该模块包含一个由lpTemplateName成员命名的对话框模板。如果两个标志都没有设置，则忽略此成员。如果设置了OFN_EXPLORER标志，系统将使用指定的模板创建一个对话框，该对话框是默认资源管理器样式对话框的子对话框。如果未设置OFN_EXPLORER标志，则系统使用模板创建旧式对话框，替换默认对话框。
	HInstance int

	// 过滤器. 包含成对的以 null 结尾的过滤器字符串的缓冲区。缓冲区中的最后一个字符串必须以两个NULL字符终止。
	//
	//	每对中的第一个字符串是描述过滤器的显示字符串（例如:“文本文件”），第二个字符串指定过滤器模式（例如:“ .TXT”）。要为单个显示字符串指定多个过滤器模式，请使用分号分隔模式（例如:“ .TXT; .DOC; .BAK”）。模式字符串可以是有效文件名字符和星号 (*) 通配符的组合。不要在模式字符串中包含空格。
	//
	//	系统不会更改过滤器的顺序。它按照lpstrFilter中指定的顺序将它们显示在文件类型组合框中。
	//
	//	如果lpstrFilter为NULL，则对话框不显示任何过滤器。
	//
	//	例子:
	//	c := "\x00"
	//	lpstrFilter := strings.Join([]string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}, c) + c + c
	//	common.StringToUint16Ptr(lpstrFilter)
	LpstrFilter *uint16

	// 一个静态缓冲区，其中包含一对以空值结尾的过滤器字符串，用于保存用户选择的过滤器模式。第一个字符串是描述自定义过滤器的显示字符串，第二个字符串是用户选择的过滤器模式。当您的应用程序第一次创建对话框时，您指定第一个字符串，它可以是任何非空字符串。当用户选择一个文件时，对话框将当前过滤模式复制到第二个字符串。保留的过滤器模式可以是lpstrFilter缓冲区中指定的模式之一，也可以是用户键入的过滤器模式。下次创建对话框时，系统使用字符串初始化用户定义的文件过滤器。如果nFilterIndex成员为零，对话框使用自定义过滤器。
	//
	//	如果此成员为NULL，则对话框不保留用户定义的过滤器模式。
	//
	//	如果此成员不是NULL，则nMaxCustFilter成员的值必须指定lpstrCustomFilter缓冲区的大小（以字符为单位） 。
	LpstrCustomFilter *uint16

	// 由lpstrCustomFilter标识的缓冲区的大小（以字符为单位） 。此缓冲区应至少有 40 个字符长。如果lpstrCustomFilter为NULL或指向NULL字符串，则忽略此成员。
	NMaxCustFilter uint32

	// 当前选定过滤器的索引。lpstrFilter指向的缓冲区包含定义过滤器的字符串对。第一对字符串的索引值为 1，第二对字符串的索引值为 2，依此类推。索引为零表示由 lpstrCustomFilter指定的自定义过滤器。您可以在输入上指定一个索引，以指示对话框的初始过滤器描述和过滤器模式。当用户选择一个文件时，nFilterIndex返回当前显示的过滤器的索引。如果nFilterIndex为零且lpstrCustomFilter为NULL，则系统使用lpstrFilter中的第一个过滤器缓冲。如果所有三个成员都为零或NULL，则系统不使用任何过滤器，并且不会在对话框的文件列表控件中显示任何文件。
	NFilterIndex uint32

	// 用于初始化文件名编辑控件的文件名。如果不需要初始化，此缓冲区的第一个字符必须为NULL 。当 GetOpenFileNameW 或 GetSaveFileNameW 函数成功返回时，此缓冲区包含所选文件的驱动器指示符、路径、文件名和扩展名。
	//
	//	如果设置了 OFN_ALLOWMULTISELECT 标志并且用户选择了多个文件，则缓冲区包含当前目录，后跟所选文件的文件名。对于资源管理器样式的对话框，目录和文件名字符串是NULL分隔的，最后一个文件名后有一个额外的NULL字符。对于旧式对话框，字符串以空格分隔，函数使用短文件名作为带空格的文件名。您可以使用 FindFirstFile 函数在长文件名和短文件名之间进行转换。如果用户只选择一个文件，lpstrFile字符串在路径和文件名之间没有分隔符。
	//
	//	如果缓冲区太小，该函数返回FALSE并且 CommDlgExtendedError 函数返回 FNERR_BUFFERTOOSMALL 。在这种情况下，lpstrFile缓冲区的前两个字节包含所需的大小，以字节或字符为单位。
	//
	//	例子:
	//	lpstrFile := make([]uint16, 260)//初始大小如果是单选文件的话, 就填260. 多选文件的话, 可根据情况增大
	//	然后填写: &lpstrFile[0]
	LpstrFile *uint16

	// lpstrFile指向的缓冲区的大小（以字符为单位）。缓冲区必须足够大以存储路径和文件名字符串或字符串，包括终止NULL字符。如果缓冲区太小而无法包含文件信息， GetOpenFileNameW 和 GetSaveFileNameW 函数将返回FALSE 。缓冲区的长度应至少为 256 个字符。
	NMaxFile uint32

	// 所选文件的文件名和扩展名（不含路径信息）。该成员可以是NULL。
	//	例子:
	//	lpstrFileTitle := make([]uint16, 260)
	//	然后填写: &lpstrFileTitle[0]
	LpstrFileTitle *uint16

	// lpstrFileTitle指向的缓冲区的大小（以字符为单位）。如果lpstrFileTitle为NULL ，则忽略此成员。
	NMaxFileTitle uint32

	// 初始目录.
	//	例: common.StrPtr("D:").
	LpstrInitialDir uintptr

	// 要放置在对话框标题栏中的字符串。如果此成员为NULL，则系统使用默认标题（即"另存为"或"打开"）。
	//	例: common.StrPtr("打开文件")
	LpstrTitle uintptr

	// 标志: wapi.OFN_ , 可组合, 可为0.
	Flags OFN_

	// 从路径开头到lpstrFile指向的字符串中的文件名的从零开始的偏移量（以字符为单位）。对于 Unicode 版本，这是字符数。例如，如果lpstrFile指向以下字符串“c:\dir1\dir2\file.ext”，则该成员包含值 13 以指示“file.ext”字符串的偏移量。如果用户选择了多个文件，nFileOffset是第一个文件名的偏移量。
	NFileOffset uint16

	// 从路径开头到lpstrFile指向的字符串中的文件扩展名的从零开始的偏移量（以字符为单位）。对于 Unicode 版本，这是字符数。通常，文件扩展名是最后一次出现的点（“.”）字符之后的子字符串。例如txt是文件名readme.txt的扩展名，html是readme.txt.html的扩展名。因此，如果lpstrFile指向字符串“c:\dir1\dir2\readme.txt”，则该成员包含值 20。如果lpstrFile指向字符串“c:\dir1\dir2\readme.txt.html”，则此成员成员包含值 24。如果lpstrFile指向字符串“c:\dir1\dir2\readme.txt.html.”，该成员包含值 29。如果lpstrFile指向的字符串不包含任何“.” 字符如“c:\dir1\dir2\readme”，该成员包含零。
	NFileExtension uint16

	// 默认扩展名。如果用户未能键入扩展名， GetOpenFileNameW 和 GetSaveFileNameW 会将此扩展名附加到文件名中。此字符串可以是任意长度，但仅附加前三个字符。该字符串不应包含句点 (.)。如果此成员为NULL并且用户未能键入扩展名，则不会附加任何扩展名。
	//
	//	例子: common.StrPtr("txt")
	LpstrDefExt uintptr

	// 系统传递给由lpfnHook成员标识的钩子过程的应用程序定义的数据。当系统向挂钩过程发送 WM_INITDIALOG 消息时，该消息的lParam参数是一个指向对话框创建时指定的 OpenFileNameW 结构的指针。挂钩过程可以使用此指针来获取lCustData值。
	LCustData uintptr

	// 指向钩子过程的指针。除非Flags成员包含 OFN_ENABLEHOOK 标志，否则此成员将被忽略。
	//
	//	如果在Flags成员中没有设置 OFN_EXPLORER 标志， lpfnHook是一个指向 OFNHookProcOldStyle 挂钩过程的指针，该过程接收用于对话框的消息。挂钩过程返回FALSE以将消息传递给默认对话框过程，或返回TRUE以丢弃消息。
	//
	//	如果设置了 OFN_EXPLORER，lpfnHook是一个指向 OFNHookProc 挂钩过程的指针。挂钩过程接收从对话框发送的通知消息。挂钩过程还接收您通过指定子对话框模板定义的任何其他控件的消息。挂钩过程不接收用于默认对话框的标准控件的消息。
	LpfnHook uintptr

	// 由hInstance成员标识的模块中对话模板资源的名称。对于编号的对话框资源，这可以是 MAKEINTRESOURCE 宏返回的值。除非在Flags成员中设置了 OFN_ENABLETEMPLATE 标志，否则该成员将被忽略。如果设置了 OFN_EXPLORER 标志，系统将使用指定的模板创建一个对话框，该对话框是默认资源管理器样式对话框的子对话框。如果未设置 OFN_EXPLORER 标志，则系统使用模板创建旧式对话框，替换默认对话框。
	LpTemplateName uintptr
}

// OFN_ 是用于初始化对话框的位标志
type OFN_ uint32

const (
	OFN_ALLOWMULTISELECT OFN_ = 0x00000200 // 文件名列表框允许多选 。如果您还设置了 OFN_EXPLORER 标志，则对话框使用资源管理器样式的用户界面；否则，它使用旧式用户界面。如果用户选择了多个文件，lpstrFile缓冲区会返回当前目录的路径，后跟所选文件的文件名。nFileOffset成员是第一个文件名的偏移量，以字节或字符为单位，并且不使用nFileExtension成员。对于资源管理器样式的对话框，目录和文件名字符串是NULL分隔的，最后一个文件名后有一个额外的NULL字符。这种格式使 Explorer 样式的对话框能够返回包含空格的长文件名。对于旧式对话框，目录和文件名字符串用空格分隔，函数使用短文件名作为带空格的文件名。您可以使用FindFirstFile函数在长文件名和短文件名之间进行转换。如果为旧式对话框指定自定义模板，则文件名列表框的定义必须包含LBS_EXTENDEDSEL值。

	OFN_CREATEPROMPT OFN_ = 0x00002000 // 如果用户指定的文件不存在，则此标志会导致对话框提示用户授予创建文件的权限。如果用户选择创建文件，对话框关闭并且函数返回指定的名称；否则，对话框保持打开状态。如果将此标志与 OFN_ALLOWMULTISELECT 标志一起使用，则对话框允许用户仅指定一个不存在的文件。

	OFN_DONTADDDTORECENT OFN_ = 0x02000000 // 防止系统在包含用户最近使用的文档的文件系统目录中添加指向选定文件的链接。要检索此目录的位置，请使用 CSIDL_RECENT 标志 调用 SHGetSpecialFolderLocation 函数。

	OFN_ENABLEHOOK OFN_ = 0x00000020 // 启用在lpfnHook成员中指定的钩子函数。

	OFN_ENABLEINCLUDENOTIFY OFN_ = 0x00400000 // 当用户打开文件夹时， 使对话框将 CDN_INCLUDEITEM 通知消息发送到您的 OFNHookProc 挂钩过程。该对话框会为新打开的文件夹中的每个项目发送通知。这些消息使您能够控制对话框在文件夹的项目列表中显示的项目。

	OFN_ENABLESIZING OFN_ = 0x00800000 // 允许使用鼠标或键盘调整资源管理器样式对话框的大小。默认情况下，资源管理器样式的打开和另存为对话框允许调整对话框的大小，无论是否设置了此标志。仅当您提供挂钩过程或自定义模板时，才需要此标志。旧式对话框不允许调整大小。

	OFN_ENABLETEMPLATE OFN_ = 0x00000040 // lpTemplateName成员是指向模块中由 hInstance 成员标识的对话模板资源名称 的指针。如果设置了 OFN_EXPLORER 标志，系统将使用指定的模板创建一个对话框，该对话框是默认资源管理器样式对话框的子对话框。如果未设置 OFN_EXPLORER 标志，则系统使用模板创建旧式对话框，替换默认对话框。

	OFN_ENABLETEMPLATEHANDLE OFN_ = 0x00000080 // hInstance成员标识包含预加载对话框模板的数据块 。如果指定了此标志，系统将忽略lpTemplateName 。如果设置了 OFN_EXPLORER 标志，系统将使用指定的模板创建一个对话框，该对话框是默认资源管理器样式对话框的子对话框。如果未设置 OFN_EXPLORER 标志，则系统使用模板创建旧式对话框，替换默认对话框。

	OFN_EXPLORER OFN_ = 0x00080000 // 指示对“打开”或“另存为”对话框所做的任何自定义都使用资源管理器样式的自定义方法。有关详细信息，请参阅Explorer-Style Hook Procedures和Explorer-Style Custom Templates。默认情况下，无论是否设置了此标志，“打开”和“另存为”对话框都使用资源管理器样式的用户界面。仅当您提供挂钩过程或自定义模板或设置 OFN_ALLOWMULTISELECT 标志时，才需要此标志。如果您想要旧式用户界面，请省略 OFN_EXPLORER 标志并提供替换旧式模板或挂钩过程。如果您想要旧样式但不需要自定义模板或挂钩过程，只需提供一个始终返回FALSE的挂钩过程。

	OFN_EXTENSIONDIFFERENT OFN_ = 0x00000400 // 用户键入的文件扩展名与lpstrDefExt 指定的扩展名不同。如果lpstrDefExt为NULL ，则该函数不使用此标志。

	OFN_FILEMUSTEXIST OFN_ = 0x00001000 // 用户只能在文件名输入字段中键入现有文件的名称。如果指定了此标志并且用户输入了无效名称，则对话框过程会在消息框中显示警告。如果指定了此标志，则还使用 OFN_PATHMUSTEXIST 标志。此标志可在打开对话框中使用。它不能与另存为对话框一起使用。

	OFN_FORCESHOWHIDDEN OFN_ = 0x10000000 // 强制显示系统和隐藏文件，从而覆盖用户设置以显示或不显示隐藏文件。但是，未显示标记为系统和隐藏的文件。

	OFN_HIDEREADONLY OFN_ = 0x00000004 // 隐藏只读复选框。

	OFN_LONGNAMES OFN_ = 0x00200000 // 对于旧式对话框，此标志使对话框使用长文件名。如果未指定此标志，或者还设置了 OFN_ALLOWMULTISELECT 标志，则旧式对话框使用短文件名（8.3 格式）作为带空格的文件名。资源管理器样式的对话框忽略此标志并始终显示长文件名。

	OFN_NOCHANGEDIR OFN_ = 0x00000008 // 如果用户在搜索文件时更改了目录，则将当前目录恢复为其原始值。此标志对 GetOpenFileNameW 无效。

	OFN_NODEREFERENCELINKS OFN_ = 0x00100000 // 指示对话框返回所选快捷方式 (.LNK) 文件的路径和文件名。如果未指定此值，则对话框返回快捷方式引用的文件的路径和文件名。

	OFN_NOLONGNAMES OFN_ = 0x00040000 // 对于旧式对话框，此标志使对话框使用短文件名（8.3 格式）。资源管理器样式的对话框忽略此标志并始终显示长文件名。

	OFN_NONETWORKBUTTON OFN_ = 0x00020000 // 隐藏和禁用网络按钮。

	OFN_NOREADONLYRETURN OFN_ = 0x00008000 // 返回的文件没有选中只读复选框，并且不在写保护目录中。

	OFN_NOTESTFILECREATE OFN_ = 0x00010000 // 在关闭对话框之前不会创建文件。如果应用程序将文件保存在创建非修改网络共享上，则应指定此标志。当应用程序指定此标志时，库不检查写保护、磁盘已满、驱动器门打开或网络保护。使用此标志的应用程序必须小心执行文件操作，因为文件一旦关闭就无法重新打开。

	OFN_NOVALIDATE OFN_ = 0x00000100 // 常用对话框允许在返回的文件名中包含无效字符。通常，调用应用程序使用一个挂钩过程，该过程使用FILEOKSTRING消息检查文件名。如果编辑控件中的文本框为空或只包含空格，则更新文件和目录列表。如果编辑控件中的文本框包含其他任何内容，则 nFileOffset和nFileExtension将设置为通过解析文本生成的值。文本中没有添加默认扩展名，文本也没有复制到lpstrFileTitle指定的缓冲区。如果nFileOffset指定的值小于零，则文件名无效。否则，文件名有效，nFileExtension并且nFileOffset可以像未指定 OFN_NOVALIDATE 标志一样使用。

	OFN_OVERWRITEPROMPT OFN_ = 0x00000002 // 如果所选文件已存在，则使“另存为”对话框生成一个消息框。用户必须确认是否覆盖文件。

	OFN_PATHMUTEXIST OFN_ = 0x00000800 // 用户只能键入有效的路径和文件名。如果使用此标志并且用户在“文件名”输入字段中键入了无效的路径和文件名，则对话框功能会在消息框中显示警告。

	OFN_READONLY OFN_ = 0x00000001 // 导致在创建对话框时最初选中只读复选框。此标志指示对话框关闭时 只读复选框的状态。

	OFN_SHAREAWARE OFN_ = 0x00004000 // 指定如果对 OpenFile函数的调用由于网络共享冲突而失败，则忽略该错误并且对话框返回选定的文件名。如果未设置此标志，则当用户指定的文件名发生网络共享冲突时，对话框会通知您的挂钩过程。如果设置了 OFN_EXPLORER 标志，对话框将CDN_SHAREVIOLATION消息发送到挂钩过程。如果不设置OFN_EXPLORER，对话框会发送SHAREVISTRING注册消息到钩子过程。

	OFN_SHOWHELP OFN_ = 0x00000010 // 使对话框显示“帮助”按钮。hwndOwner成员必须指定窗口以接收当用户单击帮助按钮时对话框发送的HELPMSGSTRING注册消息。当用户单击“帮助”按钮 时，资源管理器样式的对话框会向您的挂钩过程发送CDN_HELP通知消息。
)

// GetOpenFileNameW 创建一个打开对话框，让用户指定要打开的文件或文件集的驱动器、目录和名称。
//	@Description 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/commdlg/nf-commdlg-getopenfilenamew.
//	@param unnamedParam1 指向包含用于初始化对话框的信息的 wapi.OpenFileNameW 结构的指针。当函数返回时，此结构包含有关用户文件选择的信息。
//	@return bool
//
func GetOpenFileNameW(unnamedParam1 *OpenFileNameW) bool {
	r, _, _ := getOpenFileNameW.Call(uintptr(unsafe.Pointer(unnamedParam1)))
	return r != 0
}

// GetSaveFileNameW 创建一个保存对话框，让用户指定要保存的文件的驱动器、目录和名称。
//	@Description 详情: https://docs.microsoft.com/zh-cn/windows/win32/api/commdlg/nf-commdlg-GetSaveFileNameW.
//	@param unnamedParam1 指向包含用于初始化对话框的信息的 wapi.OpenFileNameW 结构的指针。当函数返回时，此结构包含有关用户文件选择的信息。
//	@return bool
//
func GetSaveFileNameW(unnamedParam1 *OpenFileNameW) bool {
	r, _, _ := getSaveFileNameW.Call(uintptr(unsafe.Pointer(unnamedParam1)))
	return r != 0
}

//  ChooseColor 包含 ChooseColorW 函数用于初始化Color对话框的信息。用户关闭对话框后，系统在此结构中返回有关用户选择的信息。
//
type ChooseColor struct {
	// 结构的长度（以字节为单位）。
	//	cc := wapi.ChooseColor{...}
	//	cc.LStructSize = uint32(unsafe.Sizeof(cc))
	LStructSize uint32

	// 拥有对话框的窗口句柄。此成员可以是任何有效的窗口句柄，或者如果对话框没有所有者，它可以为NULL 。
	HwndOwner int

	// 如果在Flags成员中设置了 CC_ENABLETEMPLATEHANDLE 标志，则hInstance是包含对话框模板的内存对象的句柄。如果设置了 CC_ENABLETEMPLATE 标志，则hInstance是一个模块句柄，该模块包含一个由lpTemplateName成员命名的对话框模板。如果既未设置 CC_ENABLETEMPLATEHANDLE 也未设置 CC_ENABLETEMPLATE，则忽略此成员。
	HInstance int

	// 如果设置了 CC_RGBINIT 标志，则rgbResult指定创建对话框时最初选择的颜色。如果指定的颜色值不在可用颜色中，则系统选择最接近的可用纯色。如果rgbResult为零或未设置 CC_RGBINIT，则最初选择的颜色为黑色。如果用户单击OK按钮，则 rgbResult指定用户的颜色选择。要创建RGB颜色值，请使用: xc.RGB().
	RgbResult uint32

	// 指向包含对话框中自定义颜色框的红、绿、蓝 (RGB) 值的 16 个值的数组的指针。如果用户修改了这些颜色，系统将使用新的 RGB 值更新数组。要在调用 ChooseColorW 函数之间保留新的自定义颜色，您应该为数组分配静态内存。要创建RGB颜色值，请使用: xc.RGB().
	//	例子:
	//	var lpCustColors [16]uint32
	//	然后填 &lpCustColors[0]
	LpCustColors *uint32

	// 一组可用于初始化颜色对话框的位标志。当对话框返回时，它会设置这些标志来指示用户的输入。该成员可以是 wapi.CC_ 的组合。
	Flags CC_

	// 系统传递给由lpfnHook成员标识的钩子过程的应用程序定义的数据。当系统向挂钩过程发送 WM_INITDIALOG 消息时，该消息的lParam参数是一个指向对话框创建时指定的 ChooseColor 结构的指针。挂钩过程可以使用此指针来获取lCustData值。
	LCustData uintptr

	// 指向可以处理用于对话框的消息的CCHookProc挂钩过程的指针。除非在Flags成员中设置了 CC_ENABLEHOOK 标志，否则该成员将被忽略。
	LpfnHook uintptr

	// hInstance成员标识的模块中对话框模板资源的名称。此模板替代了标准对话框模板。对于编号的对话框资源，lpTemplateName可以是 MAKEINTRESOURCE宏返回的值。除非在Flags成员中设置了 CC_ENABLETEMPLATE 标志，否则此成员将被忽略。
	LpTemplateName uintptr
}

// CC_ 是可用于初始化颜色对话框的位标志。
type CC_ uint32

const (
	CC_ANYCOLOR CC_ = 0x00000100 // 使对话框显示一组基本颜色中的所有可用颜色。

	CC_ENABLEHOOK CC_ = 0x00000010 // 启用在此结构的lpfnHook成员中指定的挂钩过程。此标志仅用于初始化对话框。

	CC_ENABLETEMPLATE CC_ = 0x00000020 // hInstance和 lpTemplateName成员指定一个对话框模板来代替默认模板。此标志仅用于初始化对话框。

	CC_ENABLETEMPLATEHANDLE CC_ = 0x00000040 // hInstance成员标识包含预加载对话框模板的数据块 。如果指定了此标志，系统将忽略lpTemplateName成员。此标志仅用于初始化对话框。

	CC_FULLOPEN CC_ = 0x00000002 // 使对话框显示允许用户创建自定义颜色的附加控件。如果未设置此标志，用户必须单击定义自定义颜色按钮以显示自定义颜色控件。

	CC_PREVENTFULLOPEN CC_ = 0x00000004 // 禁用定义自定义颜色按钮。

	CC_RGBINIT CC_ = 0x00000001 // 使对话框使用rgbResult成员中指定的颜色作为初始颜色选择。

	CC_SHOWHELP CC_ = 0x00000008 // 使对话框显示“帮助”按钮。hwndOwner成员必须指定窗口以接收当用户单击帮助按钮 时对话框发送的 HELPMSGSTRING 注册消息。

	CC_SOLIDCOLOR CC_ = 0x00000080 // 使对话框仅显示基本颜色集中的纯色。
)

// ChooseColorW 创建一个颜色对话框，使用户能够选择一种颜色。
//	@Description 详情: https://docs.microsoft.com/zh-cn/previous-versions/windows/desktop/legacy/ms646912(v=vs.85).
//	@param lpcc 指向 wapi.ChooseColor 结构的指针，该结构包含用于初始化对话框的信息。当函数返回时，此结构包含有关用户颜色选择的信息。
//	@return bool
//
func ChooseColorW(lpcc *ChooseColor) bool {
	r, _, _ := chooseColorW.Call(uintptr(unsafe.Pointer(lpcc)))
	return r != 0
}

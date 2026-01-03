package window

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// ModalWindow 模态窗口.
type ModalWindow struct {
	windowBase
}

// 模态窗口_创建, 创建模态窗口; 当模态窗口关闭时, 会自动销毁模态窗口资源句柄, 失败返回 nil.
//
// nWidth: 宽度.
//
// nHeight: 高度.
//
// title: 窗口标题内容.
//
// hWndParent: 父窗口句柄.
//
// XCStyle: 炫彩窗口样式: xcc.Window_Style_.
func NewModalWindow(nWidth, nHeight int32, title string, hWndParent uintptr, XCStyle xcc.Window_Style_) *ModalWindow {
	return NewModalWindowByHandle(xc.XModalWnd_Create(nWidth, nHeight, title, hWndParent, XCStyle))
}

// 模态窗口_创建扩展, 创建模态窗口, 增强功能, 失败返回 nil.
//
// dwExStyle: 窗口扩展样式.
//
// dwStyle: 窗口样式.
//
// lpClassName: 窗口类名.
//
// x: 窗口左上角x坐标.
//
// y: 窗口左上角y坐标.
//
// cx: 窗口宽度.
//
// cy: 窗口高度.
//
// title: 窗口名.
//
// hWndParent: 父窗口.
//
// XCStyle: GUI库窗口样式: xcc.Window_Style_.
func NewModalWindowEx(dwExStyle, dwStyle uint32, lpClassName string, x, y, cx, cy int32, title string, hWndParent uintptr, XCStyle xcc.Window_Style_) *ModalWindow {
	return NewModalWindowByHandle(xc.XModalWnd_CreateEx(dwExStyle, dwStyle, title, x, y, cx, cy, lpClassName, hWndParent, XCStyle))
}

// NewModalWindowByLayout 从布局文件创建对象, 失败返回 nil.
//
// pFileName: 布局文件名.
//
// hParent: 父对象句柄.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func NewModalWindowByLayout(pFileName string, hParent int, hAttachWnd uintptr) *ModalWindow {
	return NewModalWindowByHandle(xc.XC_LoadLayout(pFileName, hParent, hAttachWnd))
}

// NewModalWindowByLayoutZip 从压缩包中的布局文件创建对象, 失败返回 nil.
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
func NewModalWindowByLayoutZip(pZipFileName string, pFileName string, pPassword string, hParent int, hAttachWnd uintptr) *ModalWindow {
	return NewModalWindowByHandle(xc.XC_LoadLayoutZip(pZipFileName, pFileName, pPassword, hParent, hAttachWnd))
}

// NewModalWindowByLayoutZipMem 从内存压缩包中的布局文件创建对象, 失败返回 nil.
//
// data: 布局文件数据.
//
// pFileName: 布局文件名.
//
// pPassword: zip密码.
//
// hParent: 父对象句柄.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func NewModalWindowByLayoutZipMem(data []byte, pFileName string, pPassword string, hParent int, hAttachWnd uintptr) *ModalWindow {
	return NewModalWindowByHandle(xc.XC_LoadLayoutZipMem(data, pFileName, pPassword, hParent, hAttachWnd))
}

// NewModalWindowByLayoutStringW 从布局文件字符串W创建对象, 失败返回 nil.
//
// pStringXML: 字符串.
//
// hParent: 父对象.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func NewModalWindowByLayoutStringW(pStringXML string, hParent int, hAttachWnd uintptr) *ModalWindow {
	return NewModalWindowByHandle(xc.XC_LoadLayoutFromStringW(pStringXML, hParent, hAttachWnd))
}

// NewModalWindowByLayoutEx 从布局文件创建对象, 失败返回 nil.
//
// pFileName: 布局文件名.
//
// pPrefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func NewModalWindowByLayoutEx(pFileName, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *ModalWindow {
	return NewModalWindowByHandle(xc.XC_LoadLayoutEx(pFileName, pPrefixName, hParent, hParentWnd, hAttachWnd))
}

// NewModalWindowByLayoutZipResEx 从RC资源zip压缩包中的布局文件创建对象, 失败返回 nil.
//
// id: RC资源ID.
//
// pFileName: 布局文件名.
//
// pPassword: zip密码.
//
// pPrefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
//
// hModule: 模块句柄, 可填0.
func NewModalWindowByLayoutZipResEx(id int32, pFileName, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd, hModule uintptr) *ModalWindow {
	return NewModalWindowByHandle(xc.XC_LoadLayoutZipResEx(id, pFileName, pPassword, pPrefixName, hParent, hParentWnd, hAttachWnd, hModule))
}

// NewModalWindowByLayoutZipEx 从压缩包中的布局文件创建对象, 失败返回 nil.
//
// pZipFileName: zip文件名.
//
// pFileName: 布局文件名.
//
// pPassword: zip密码.
//
// pPrefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func NewModalWindowByLayoutZipEx(pZipFileName string, pFileName string, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *ModalWindow {
	return NewModalWindowByHandle(xc.XC_LoadLayoutZipEx(pZipFileName, pFileName, pPassword, pPrefixName, hParent, hParentWnd, hAttachWnd))
}

// NewModalWindowByLayoutZipMemEx 从内存压缩包中的布局文件创建对象, 失败返回 nil.
//
// data: 布局文件数据.
//
// pFileName: 布局文件名.
//
// pPassword: zip密码.
//
// pPrefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func NewModalWindowByLayoutZipMemEx(data []byte, pFileName string, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *ModalWindow {
	return NewModalWindowByHandle(xc.XC_LoadLayoutZipMemEx(data, pFileName, pPassword, pPrefixName, hParent, hParentWnd, hAttachWnd))
}

// NewModalWindowByLayoutStringWEx 从布局文件字符串W创建对象, 失败返回 nil.
//
// pStringXML: 字符串.
//
// pPrefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func NewModalWindowByLayoutStringWEx(pStringXML, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *ModalWindow {
	return NewModalWindowByHandle(xc.XC_LoadLayoutFromStringWEx(pStringXML, pPrefixName, hParent, hParentWnd, hAttachWnd))
}

// 模态窗口_附加窗口, 返回窗口对象, 失败返回 nil.
//
// hWnd: 要附加的外部窗口句柄.
//
// XCStyle: 炫彩窗口样式: xcc.Window_Style_.
func ModalWnd_Attach(hWnd uintptr, XCStyle xcc.Window_Style_) *Window {
	return NewByHandle(xc.XModalWnd_Attach(hWnd, XCStyle))
}

// 从句柄创建对象, 失败返回 nil.
func NewModalWindowByHandle(handle int) *ModalWindow {
	if handle <= 0 {
		return nil
	}
	p := &ModalWindow{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回 nil.
func NewModalWindowByName(name string) *ModalWindow {
	return NewModalWindowByHandle(xc.XC_GetObjectByName(name))
}

// 从UID创建对象, 失败返回 nil.
func NewModalWindowByUID(nUID int32) *ModalWindow {
	return NewModalWindowByHandle(xc.XC_GetObjectByUID(nUID))
}

// 从UID名称创建对象, 失败返回 nil.
func NewModalWindowByUIDName(name string) *ModalWindow {
	return NewModalWindowByHandle(xc.XC_GetObjectByUIDName(name))
}

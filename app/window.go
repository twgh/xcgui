package app

import (
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

// NewWindow 窗口_创建.
//
// x: 窗口左上角x坐标.
//
// y: 窗口左上角y坐标.
//
// cx: 窗口宽度.
//
// cy: 窗口高度.
//
// pTitle: 窗口标题.
//
// hWndParent: 父窗口.
//
// XCStyle: 窗口样式: xcc.Window_Style_.
func NewWindow(x, y, cx, cy int32, pTitle string, hWndParent uintptr, XCStyle xcc.Window_Style_) *window.Window {
	return window.New(x, y, cx, cy, pTitle, hWndParent, XCStyle)
}

// NewWindowEx 窗口_创建扩展.
//
// dwExStyle: 窗口扩展样式. xcc.WS_EX_ .
//
// dwStyle: 窗口样式. xcc.WS_ .
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
// pTitle: 窗口名.
//
// hWndParent: 父窗口.
//
// XCStyle: 窗口样式, xcc.Window_Style_.
func NewWindowEx(dwExStyle xcc.WS_EX_, dwStyle xcc.WS_, lpClassName string, x, y, cx, cy int32, pTitle string, hWndParent uintptr, XCStyle xcc.Window_Style_) *window.Window {
	return window.NewEx(dwExStyle, dwStyle, lpClassName, x, y, cx, cy, pTitle, hWndParent, XCStyle)
}

// NewWindowByAttach 窗口_附加窗口, 返回窗口对象.
//
// hWnd: 要附加的外部窗口句柄.
//
// XCStyle: 窗口样式: xcc.Window_Style_.
func NewWindowByAttach(hWnd uintptr, XCStyle xcc.Window_Style_) *window.Window {
	return window.Attach(hWnd, XCStyle)
}

// NewWindowByHandle 从句柄创建窗口对象.
//
// hWindow: 窗口资源句柄.
func NewWindowByHandle(hWindow int) *window.Window {
	return window.NewByHandle(hWindow)
}

// NewWindowByLayout 从布局文件创建窗口对象, 失败返回nil.
//
// pFileName: 布局文件名.
//
// hParent: 父对象句柄.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func NewWindowByLayout(pFileName string, hParent int, hAttachWnd uintptr) *window.Window {
	return window.NewByLayout(pFileName, hParent, hAttachWnd)
}

// NewWindowByLayoutZip 从压缩包中的布局文件创建窗口对象, 失败返回nil.
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
func NewWindowByLayoutZip(pZipFileName string, pFileName string, pPassword string, hParent int, hAttachWnd uintptr) *window.Window {
	return window.NewByLayoutZip(pZipFileName, pFileName, pPassword, hParent, hAttachWnd)
}

// NewWindowByLayoutZipMem 从内存压缩包中的布局文件创建窗口对象, 失败返回nil.
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
func NewWindowByLayoutZipMem(data []byte, pFileName string, pPassword string, hParent int, hAttachWnd uintptr) *window.Window {
	return window.NewByLayoutZipMem(data, pFileName, pPassword, hParent, hAttachWnd)
}

// NewWindowByLayoutStringW 从布局文件字符串W创建窗口对象, 失败返回nil.
//
// pStringXML: 字符串.
//
// hParent: 父对象.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func NewWindowByLayoutStringW(pStringXML string, hParent int, hAttachWnd uintptr) *window.Window {
	return window.NewByLayoutStringW(pStringXML, hParent, hAttachWnd)
}

// NewWindowByLayoutEx 从布局文件创建窗口对象, 失败返回nil.
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
func NewWindowByLayoutEx(pFileName, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *window.Window {
	return window.NewByLayoutEx(pFileName, pPrefixName, hParent, hParentWnd, hAttachWnd)
}

// NewWindowByLayoutZipEx 从压缩包中的布局文件创建窗口对象, 失败返回nil.
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
func NewWindowByLayoutZipEx(pZipFileName string, pFileName string, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *window.Window {
	return window.NewByLayoutZipEx(pZipFileName, pFileName, pPassword, pPrefixName, hParent, hParentWnd, hAttachWnd)
}

// NewWindowByLayoutZipResEx 从RC资源zip压缩包中的布局文件创建窗口对象, 失败返回nil.
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
func NewWindowByLayoutZipResEx(id int32, pFileName, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd, hModule uintptr) *window.Window {
	return window.NewByLayoutZipResEx(id, pFileName, pPassword, pPrefixName, hParent, hParentWnd, hAttachWnd, hModule)
}

// NewWindowByLayoutZipMemEx 从内存压缩包中的布局文件创建窗口对象, 失败返回nil.
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
func NewWindowByLayoutZipMemEx(data []byte, pFileName string, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *window.Window {
	return window.NewByLayoutZipMemEx(data, pFileName, pPassword, pPrefixName, hParent, hParentWnd, hAttachWnd)
}

// NewWindowByLayoutStringWEx 从布局文件字符串W创建窗口对象, 失败返回nil.
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
func NewWindowByLayoutStringWEx(pStringXML, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *window.Window {
	return window.NewByLayoutStringWEx(pStringXML, pPrefixName, hParent, hParentWnd, hAttachWnd)
}

// NewWindowByName 从name创建对象, 失败返回nil.
//
// name: name名称.
func NewWindowByName(name string) *window.Window {
	return window.NewByName(name)
}

// NewWindowByUID 从UID创建窗口对象, 失败返回nil.
//
// nUID: UID值.
func NewWindowByUID(nUID int32) *window.Window {
	return window.NewByUID(nUID)
}

// NewWindowByUIDName 从UID名称创建对象, 失败返回nil.
//
// name: name名称.
func NewWindowByUIDName(name string) *window.Window {
	return window.NewByUIDName(name)
}

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
// title: 窗口标题.
//
// hWndParent: 父窗口.
//
// XCStyle: 窗口样式: xcc.Window_Style_.
func NewWindow(x, y, cx, cy int32, title string, hWndParent uintptr, XCStyle xcc.Window_Style_) *window.Window {
	return window.New(x, y, cx, cy, title, hWndParent, XCStyle)
}

// NewWindowEx 窗口_创建扩展.
//
// dwExStyle: 窗口扩展样式. xcc.WS_EX_ .
//
// dwStyle: 窗口样式. xcc.WS_ .
//
// className: 窗口类名.
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
// XCStyle: 窗口样式, xcc.Window_Style_.
func NewWindowEx(dwExStyle xcc.WS_EX_, dwStyle xcc.WS_, className string, x, y, cx, cy int32, title string, hWndParent uintptr, XCStyle xcc.Window_Style_) *window.Window {
	return window.NewEx(dwExStyle, dwStyle, className, x, y, cx, cy, title, hWndParent, XCStyle)
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
// fileName: 布局文件名.
//
// hParent: 父对象句柄.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func NewWindowByLayout(fileName string, hParent int, hAttachWnd uintptr) *window.Window {
	return window.NewByLayout(fileName, hParent, hAttachWnd)
}

// NewWindowByLayoutZip 从压缩包中的布局文件创建窗口对象, 失败返回nil.
//
// zipFileName: zip文件名.
//
// fileName: 布局文件名.
//
// password: zip密码.
//
// hParent: 父对象句柄.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func NewWindowByLayoutZip(zipFileName string, fileName string, password string, hParent int, hAttachWnd uintptr) *window.Window {
	return window.NewByLayoutZip(zipFileName, fileName, password, hParent, hAttachWnd)
}

// NewWindowByLayoutZipMem 从内存压缩包中的布局文件创建窗口对象, 失败返回nil.
//
// data: 布局文件数据.
//
// fileName: 布局文件名.
//
// password: zip密码.
//
// hParent: 父对象句柄.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func NewWindowByLayoutZipMem(data []byte, fileName string, password string, hParent int, hAttachWnd uintptr) *window.Window {
	return window.NewByLayoutZipMem(data, fileName, password, hParent, hAttachWnd)
}

// NewWindowByLayoutStringW 从布局文件字符串W创建窗口对象, 失败返回nil.
//
// xmlStr: 字符串.
//
// hParent: 父对象.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func NewWindowByLayoutStringW(xmlStr string, hParent int, hAttachWnd uintptr) *window.Window {
	return window.NewByLayoutStringW(xmlStr, hParent, hAttachWnd)
}

// NewWindowByLayoutEx 从布局文件创建窗口对象, 失败返回nil.
//
// fileName: 布局文件名.
//
// prefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func NewWindowByLayoutEx(fileName, prefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *window.Window {
	return window.NewByLayoutEx(fileName, prefixName, hParent, hParentWnd, hAttachWnd)
}

// NewWindowByLayoutZipEx 从压缩包中的布局文件创建窗口对象, 失败返回nil.
//
// zipFileName: zip文件名.
//
// fileName: 布局文件名.
//
// password: zip密码.
//
// prefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func NewWindowByLayoutZipEx(zipFileName string, fileName string, password, prefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *window.Window {
	return window.NewByLayoutZipEx(zipFileName, fileName, password, prefixName, hParent, hParentWnd, hAttachWnd)
}

// NewWindowByLayoutZipResEx 从RC资源zip压缩包中的布局文件创建窗口对象, 失败返回nil.
//
// id: RC资源ID.
//
// fileName: 布局文件名.
//
// password: zip密码.
//
// prefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
//
// hModule: 模块句柄, 可填0.
func NewWindowByLayoutZipResEx(id int32, fileName, password, prefixName string, hParent int, hParentWnd, hAttachWnd, hModule uintptr) *window.Window {
	return window.NewByLayoutZipResEx(id, fileName, password, prefixName, hParent, hParentWnd, hAttachWnd, hModule)
}

// NewWindowByLayoutZipMemEx 从内存压缩包中的布局文件创建窗口对象, 失败返回nil.
//
// data: 布局文件数据.
//
// fileName: 布局文件名.
//
// password: zip密码.
//
// prefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func NewWindowByLayoutZipMemEx(data []byte, fileName string, password, prefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *window.Window {
	return window.NewByLayoutZipMemEx(data, fileName, password, prefixName, hParent, hParentWnd, hAttachWnd)
}

// NewWindowByLayoutStringWEx 从布局文件字符串W创建窗口对象, 失败返回nil.
//
// xmlStr: 字符串.
//
// prefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func NewWindowByLayoutStringWEx(xmlStr, prefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *window.Window {
	return window.NewByLayoutStringWEx(xmlStr, prefixName, hParent, hParentWnd, hAttachWnd)
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

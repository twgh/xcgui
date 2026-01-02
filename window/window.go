package window

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// Window 普通窗口.
type Window struct {
	windowBase
}

// New 窗口_创建, 失败返回 nil.
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
// hWndParent: 父窗口(真实窗口句柄).
//
// XCStyle: 窗口样式: xcc.Window_Style_.
func New(x, y, cx, cy int32, pTitle string, hWndParent uintptr, XCStyle xcc.Window_Style_) *Window {
	return NewByHandle(xc.XWnd_Create(x, y, cx, cy, pTitle, hWndParent, XCStyle))
}

// NewEx 窗口_创建扩展, 失败返回 nil.
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
// hWndParent: 父窗口(真实窗口句柄).
//
// XCStyle: 窗口样式, xcc.Window_Style_.
func NewEx(dwExStyle xcc.WS_EX_, dwStyle xcc.WS_, lpClassName string, x, y, cx, cy int32, pTitle string, hWndParent uintptr, XCStyle xcc.Window_Style_) *Window {
	return NewByHandle(xc.XWnd_CreateEx(dwExStyle, dwStyle, lpClassName, x, y, cx, cy, pTitle, hWndParent, XCStyle))
}

// Attach 窗口_附加窗口, 返回窗口对象, 失败返回 nil.
//
// hWnd: 要附加的外部窗口句柄.
//
// XCStyle: 窗口样式: xcc.Window_Style_.
func Attach(hWnd uintptr, XCStyle xcc.Window_Style_) *Window {
	return NewByHandle(xc.XWnd_Attach(hWnd, XCStyle))
}

// NewByHandle 从句柄创建窗口对象, 失败返回 nil.
//
// hWindow: 窗口资源句柄.
func NewByHandle(hWindow int) *Window {
	if hWindow <= 0 {
		return nil
	}
	p := &Window{}
	p.SetHandle(hWindow)
	return p
}

// NewByLayout 从布局文件创建窗口对象, 失败返回 nil.
//
// pFileName: 布局文件名.
//
// hParent: 父对象句柄.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func NewByLayout(pFileName string, hParent int, hAttachWnd uintptr) *Window {
	return NewByHandle(xc.XC_LoadLayout(pFileName, hParent, hAttachWnd))
}

// NewByLayoutZip 从压缩包中的布局文件创建窗口对象, 失败返回 nil.
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
func NewByLayoutZip(pZipFileName string, pFileName string, pPassword string, hParent int, hAttachWnd uintptr) *Window {
	return NewByHandle(xc.XC_LoadLayoutZip(pZipFileName, pFileName, pPassword, hParent, hAttachWnd))
}

// NewByLayoutZipMem 从内存压缩包中的布局文件创建窗口对象, 失败返回 nil.
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
func NewByLayoutZipMem(data []byte, pFileName string, pPassword string, hParent int, hAttachWnd uintptr) *Window {
	return NewByHandle(xc.XC_LoadLayoutZipMem(data, pFileName, pPassword, hParent, hAttachWnd))
}

// NewByLayoutStringW 从布局文件字符串W创建窗口对象, 失败返回 nil.
//
// pStringXML: 字符串.
//
// hParent: 父对象句柄.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func NewByLayoutStringW(pStringXML string, hParent int, hAttachWnd uintptr) *Window {
	return NewByHandle(xc.XC_LoadLayoutFromStringW(pStringXML, hParent, hAttachWnd))
}

// NewByLayoutEx 从布局文件创建窗口对象, 失败返回 nil.
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
func NewByLayoutEx(pFileName, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *Window {
	return NewByHandle(xc.XC_LoadLayoutEx(pFileName, pPrefixName, hParent, hParentWnd, hAttachWnd))
}

// NewByLayoutZipEx 从压缩包中的布局文件创建窗口对象, 失败返回 nil.
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
func NewByLayoutZipEx(pZipFileName string, pFileName string, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *Window {
	return NewByHandle(xc.XC_LoadLayoutZipEx(pZipFileName, pFileName, pPassword, pPrefixName, hParent, hParentWnd, hAttachWnd))
}

// NewByLayoutZipResEx 从RC资源zip压缩包中的布局文件创建窗口对象, 失败返回 nil.
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
func NewByLayoutZipResEx(id int32, pFileName, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd, hModule uintptr) *Window {
	return NewByHandle(xc.XC_LoadLayoutZipResEx(id, pFileName, pPassword, pPrefixName, hParent, hParentWnd, hAttachWnd, hModule))
}

// NewByLayoutZipMemEx 从内存压缩包中的布局文件创建窗口对象, 失败返回 nil.
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
func NewByLayoutZipMemEx(data []byte, pFileName string, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *Window {
	return NewByHandle(xc.XC_LoadLayoutZipMemEx(data, pFileName, pPassword, pPrefixName, hParent, hParentWnd, hAttachWnd))
}

// NewByLayoutStringWEx 从布局文件字符串W创建窗口对象, 失败返回 nil.
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
func NewByLayoutStringWEx(pStringXML, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *Window {
	return NewByHandle(xc.XC_LoadLayoutFromStringWEx(pStringXML, pPrefixName, hParent, hParentWnd, hAttachWnd))
}

// NewByName 从name创建对象, 失败返回 nil.
//
// name: name名称.
func NewByName(name string) *Window {
	return NewByHandle(xc.XC_GetObjectByName(name))
}

// NewByUID 从UID创建窗口对象, 失败返回 nil.
//
// nUID: UID值.
func NewByUID(nUID int32) *Window {
	return NewByHandle(xc.XC_GetObjectByUID(nUID))
}

// NewByUIDName 从UID名称创建对象, 失败返回 nil.
//
// name: name名称.
func NewByUIDName(name string) *Window {
	return NewByHandle(xc.XC_GetObjectByUIDName(name))
}

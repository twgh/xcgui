package window

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// Window 普通窗口.
type Window struct {
	windowBase
}

// New 窗口_创建.
//
//	@param x 窗口左上角x坐标.
//	@param y 窗口左上角y坐标.
//	@param cx 窗口宽度.
//	@param cy 窗口高度.
//	@param pTitle 窗口标题.
//	@param hWndParent 父窗口.
//	@param XCStyle 窗口样式: xcc.Window_Style_.
//	@return *Window
func New(x, y, cx, cy int32, pTitle string, hWndParent uintptr, XCStyle xcc.Window_Style_) *Window {
	p := &Window{}
	p.SetHandle(xc.XWnd_Create(x, y, cx, cy, pTitle, hWndParent, XCStyle))
	return p
}

// NewEx 窗口_创建扩展.
//
//	@param dwExStyle 窗口扩展样式.
//	@param dwStyle 窗口样式.
//	@param lpClassName 窗口类名.
//	@param x 窗口左上角x坐标.
//	@param y 窗口左上角y坐标.
//	@param cx 窗口宽度.
//	@param cy 窗口高度.
//	@param pTitle 窗口名.
//	@param hWndParent 父窗口.
//	@param XCStyle 窗口样式, xcc.Window_Style_.
//	@return *Window
func NewEx(dwExStyle int, dwStyle int, lpClassName string, x int, y int, cx int, cy int, pTitle string, hWndParent uintptr, XCStyle xcc.Window_Style_) *Window {
	p := &Window{}
	p.SetHandle(xc.XWnd_CreateEx(dwExStyle, dwStyle, lpClassName, x, y, cx, cy, pTitle, hWndParent, XCStyle))
	return p
}

// Attach 窗口_附加窗口, 返回窗口对象.
//
//	@param hWnd 要附加的外部窗口句柄.
//	@param XCStyle 窗口样式: xcc.Window_Style_.
//	@return *Window
func Attach(hWnd uintptr, XCStyle xcc.Window_Style_) *Window {
	p := &Window{}
	p.SetHandle(xc.XWnd_Attach(hWnd, XCStyle))
	return p
}

// NewByHandle 从句柄创建窗口对象.
//
//	@param hWindow 窗口资源句柄.
//	@return *Window
func NewByHandle(hWindow int) *Window {
	p := &Window{}
	p.SetHandle(hWindow)
	return p
}

// NewByLayout 从布局文件创建窗口对象, 失败返回nil.
//
//	@param pFileName 布局文件名.
//	@param hParent 父对象句柄.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *Window
func NewByLayout(pFileName string, hParent int, hAttachWnd uintptr) *Window {
	handle := xc.XC_LoadLayout(pFileName, hParent, hAttachWnd)
	if handle > 0 {
		p := &Window{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewByLayoutZip 从压缩包中的布局文件创建窗口对象, 失败返回nil.
//
//	@param pZipFileName zip文件名.
//	@param pFileName 布局文件名.
//	@param pPassword zip密码.
//	@param hParent 父对象句柄.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *Window
func NewByLayoutZip(pZipFileName string, pFileName string, pPassword string, hParent int, hAttachWnd uintptr) *Window {
	handle := xc.XC_LoadLayoutZip(pZipFileName, pFileName, pPassword, hParent, hAttachWnd)
	if handle > 0 {
		p := &Window{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewByLayoutZipMem 从内存压缩包中的布局文件创建窗口对象, 失败返回nil.
//
//	@param data 布局文件数据.
//	@param pFileName 布局文件名.
//	@param pPassword zip密码.
//	@param hParent 父对象句柄.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *Window
func NewByLayoutZipMem(data []byte, pFileName string, pPassword string, hParent int, hAttachWnd uintptr) *Window {
	handle := xc.XC_LoadLayoutZipMem(data, pFileName, pPassword, hParent, hAttachWnd)
	if handle > 0 {
		p := &Window{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewByLayoutStringW 从布局文件字符串W创建窗口对象, 失败返回nil.
//
//	@param pStringXML 字符串.
//	@param hParent 父对象.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *Window
func NewByLayoutStringW(pStringXML string, hParent int, hAttachWnd uintptr) *Window {
	handle := xc.XC_LoadLayoutFromStringW(pStringXML, hParent, hAttachWnd)
	if handle > 0 {
		p := &Window{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewByLayoutEx 从布局文件创建窗口对象, 失败返回nil.
//
//	@param pFileName 布局文件名.
//	@param pPrefixName 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//	@param hParent 父对象句柄.
//	@param hParentWnd 父窗口句柄HWND, 提供给第三方窗口使用.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *Window
func NewByLayoutEx(pFileName, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *Window {
	handle := xc.XC_LoadLayoutEx(pFileName, pPrefixName, hParent, hParentWnd, hAttachWnd)
	if handle > 0 {
		p := &Window{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewByLayoutZipEx 从压缩包中的布局文件创建窗口对象, 失败返回nil.
//
//	@param pZipFileName zip文件名.
//	@param pFileName 布局文件名.
//	@param pPassword zip密码.
//	@param pPrefixName 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//	@param hParent 父对象句柄.
//	@param hParentWnd 父窗口句柄HWND, 提供给第三方窗口使用.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *Window
func NewByLayoutZipEx(pZipFileName string, pFileName string, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *Window {
	handle := xc.XC_LoadLayoutZipEx(pZipFileName, pFileName, pPassword, pPrefixName, hParent, hParentWnd, hAttachWnd)
	if handle > 0 {
		p := &Window{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewByLayoutZipResEx 从RC资源zip压缩包中的布局文件创建窗口对象, 失败返回nil.
//
//	@param id RC资源ID.
//	@param pFileName 布局文件名.
//	@param pPassword zip密码.
//	@param pPrefixName 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//	@param hParent 父对象句柄.
//	@param hParentWnd 父窗口句柄HWND, 提供给第三方窗口使用.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@param hModule 模块句柄, 可填0.
//	@return *Window
func NewByLayoutZipResEx(id int32, pFileName, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd, hModule uintptr) *Window {
	handle := xc.XC_LoadLayoutZipResEx(id, pFileName, pPassword, pPrefixName, hParent, hParentWnd, hAttachWnd, hModule)
	if handle > 0 {
		p := &Window{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewByLayoutZipMemEx 从内存压缩包中的布局文件创建窗口对象, 失败返回nil.
//
//	@param data 布局文件数据.
//	@param pFileName 布局文件名.
//	@param pPassword zip密码.
//	@param pPrefixName 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//	@param hParent 父对象句柄.
//	@param hParentWnd 父窗口句柄HWND, 提供给第三方窗口使用.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *Window
func NewByLayoutZipMemEx(data []byte, pFileName string, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *Window {
	handle := xc.XC_LoadLayoutZipMemEx(data, pFileName, pPassword, pPrefixName, hParent, hParentWnd, hAttachWnd)
	if handle > 0 {
		p := &Window{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewByLayoutStringWEx 从布局文件字符串W创建窗口对象, 失败返回nil.
//
//	@param pStringXML 字符串.
//	@param pPrefixName 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//	@param hParent 父对象.
//	@param hParentWnd 父窗口句柄HWND, 提供给第三方窗口使用.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *Window
func NewByLayoutStringWEx(pStringXML, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *Window {
	handle := xc.XC_LoadLayoutFromStringWEx(pStringXML, pPrefixName, hParent, hParentWnd, hAttachWnd)
	if handle > 0 {
		p := &Window{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewByName 从name创建对象, 失败返回nil.
//
//	@param name
//	@return *Window
func NewByName(name string) *Window {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &Window{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewByUID 从UID创建窗口对象, 失败返回nil.
//
//	@param nUID
//	@return *Window
func NewByUID(nUID int) *Window {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &Window{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewByUIDName 从UID名称创建对象, 失败返回nil.
//
//	@param name
//	@return *Window
func NewByUIDName(name string) *Window {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &Window{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

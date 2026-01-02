package widget

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// LayoutEle 布局元素.
type LayoutEle struct {
	Element
}

// NewLayoutEle 布局_创建, 创建布局元素, 失败返回 nil.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父为窗口句柄或元素句柄.
func NewLayoutEle(x, y, cx, cy int32, hParent int) *LayoutEle {
	return NewLayoutEleByHandle(xc.XLayout_Create(x, y, cx, cy, hParent))
}

// NewLayoutEleEx 布局_创建扩展, 创建布局元素, 失败返回 nil.
//
// hParent: 父为窗口句柄或元素句柄.
func NewLayoutEleEx(hParent int) *LayoutEle {
	return NewLayoutEleByHandle(xc.XLayout_CreateEx(hParent))
}

// NewLayoutEleByHandle 从句柄创建对象, 失败返回 nil.
//
// handle: 布局元素句柄.
func NewLayoutEleByHandle(handle int) *LayoutEle {
	if handle <= 0 {
		return nil
	}
	p := &LayoutEle{}
	p.SetHandle(handle)
	return p
}

// NewLayoutEleByLayout 从布局文件创建对象, 失败返回 nil.
//
// pFileName: 布局文件名.
//
// hParent: 父对象句柄.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func NewLayoutEleByLayout(pFileName string, hParent int, hAttachWnd uintptr) *LayoutEle {
	return NewLayoutEleByHandle(xc.XC_LoadLayout(pFileName, hParent, hAttachWnd))
}

// NewLayoutEleByLayoutZip 从压缩包中的布局文件创建对象, 失败返回 nil.
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
func NewLayoutEleByLayoutZip(pZipFileName string, pFileName string, pPassword string, hParent int, hAttachWnd uintptr) *LayoutEle {
	return NewLayoutEleByHandle(xc.XC_LoadLayoutZip(pZipFileName, pFileName, pPassword, hParent, hAttachWnd))
}

// NewLayoutEleByLayoutZipResEx 从RC资源zip压缩包中的布局文件创建对象, 失败返回 nil.
//
// id: RC资源ID.
//
// pFileName: zip文件名.
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
func NewLayoutEleByLayoutZipResEx(id int32, pFileName, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd, hModule uintptr) *LayoutEle {
	return NewLayoutEleByHandle(xc.XC_LoadLayoutZipResEx(id, pFileName, pPassword, pPrefixName, hParent, hParentWnd, hAttachWnd, hModule))
}

// NewLayoutEleByLayoutZipMem 从内存压缩包中的布局文件创建对象, 失败返回 nil.
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
func NewLayoutEleByLayoutZipMem(data []byte, pFileName string, pPassword string, hParent int, hAttachWnd uintptr) *LayoutEle {
	return NewLayoutEleByHandle(xc.XC_LoadLayoutZipMem(data, pFileName, pPassword, hParent, hAttachWnd))
}

// NewLayoutEleByStringW 从布局文件字符串W创建对象, 失败返回 nil.
//
// pStringXML: 字符串.
//
// hParent: 父对象.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func NewLayoutEleByStringW(pStringXML string, hParent int, hAttachWnd uintptr) *LayoutEle {
	return NewLayoutEleByHandle(xc.XC_LoadLayoutFromStringW(pStringXML, hParent, hAttachWnd))
}

// NewLayoutEleByLayoutEx 从布局文件创建对象, 失败返回 nil.
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
func NewLayoutEleByLayoutEx(pFileName, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *LayoutEle {
	return NewLayoutEleByHandle(xc.XC_LoadLayoutEx(pFileName, pPrefixName, hParent, hParentWnd, hAttachWnd))
}

// NewLayoutEleByLayoutZipEx 从压缩包中的布局文件创建对象, 失败返回 nil.
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
func NewLayoutEleByLayoutZipEx(pZipFileName string, pFileName string, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *LayoutEle {
	return NewLayoutEleByHandle(xc.XC_LoadLayoutZipEx(pZipFileName, pFileName, pPassword, pPrefixName, hParent, hParentWnd, hAttachWnd))
}

// NewLayoutEleByLayoutZipMemEx 从内存压缩包中的布局文件创建对象, 失败返回 nil.
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
func NewLayoutEleByLayoutZipMemEx(data []byte, pFileName string, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *LayoutEle {
	return NewLayoutEleByHandle(xc.XC_LoadLayoutZipMemEx(data, pFileName, pPassword, pPrefixName, hParent, hParentWnd, hAttachWnd))
}

// NewLayoutEleByStringWEx 从布局文件字符串W创建对象, 失败返回 nil.
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
func NewLayoutEleByStringWEx(pStringXML, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *LayoutEle {
	return NewLayoutEleByHandle(xc.XC_LoadLayoutFromStringWEx(pStringXML, pPrefixName, hParent, hParentWnd, hAttachWnd))
}

// NewLayoutEleByName 从name创建对象, 失败返回 nil.
//
// name: name名称.
func NewLayoutEleByName(name string) *LayoutEle {
	return NewLayoutEleByHandle(xc.XC_GetObjectByName(name))
}

// NewLayoutEleByID 从ID创建对象, 失败返回 nil.
//
// hWindow: 父窗口句柄.
//
// nID: ID值.
func NewLayoutEleByID(hWindow int, nID int32) *LayoutEle {
	return NewLayoutEleByHandle(xc.XC_GetObjectByID(hWindow, nID))
}

// NewLayoutEleByIDName 从ID名称创建对象, 失败返回 nil.
//
// hWindow: 父窗口句柄.
//
// name: name名称.
func NewLayoutEleByIDName(hWindow int, name string) *LayoutEle {
	return NewLayoutEleByHandle(xc.XC_GetObjectByIDName(hWindow, name))
}

// NewLayoutEleByUID 从UID创建对象, 失败返回 nil.
//
// nUID: ID值.
func NewLayoutEleByUID(nUID int32) *LayoutEle {
	return NewLayoutEleByHandle(xc.XC_GetObjectByUID(nUID))
}

// NewLayoutEleByUIDName 从UID名称创建对象, 失败返回 nil.
//
// name: name名称.
func NewLayoutEleByUIDName(name string) *LayoutEle {
	return NewLayoutEleByHandle(xc.XC_GetObjectByUIDName(name))
}

// IsEnableLayout 布局_判断启用, 是否已经启用布局功能.
func (l *LayoutEle) IsEnableLayout() bool {
	return xc.XLayout_IsEnableLayout(l.Handle)
}

// EnableLayout 布局_启用, 启用布局功能.
//
// bEnable: 是否启用.
func (l *LayoutEle) EnableLayout(bEnable bool) *LayoutEle {
	xc.XLayout_EnableLayout(l.Handle, bEnable)
	return l
}

// ShowLayoutFrame 布局_显示布局边界, 显示布局边界.
//
// bEnable: 是否显示.
func (l *LayoutEle) ShowLayoutFrame(bEnable bool) *LayoutEle {
	xc.XLayout_ShowLayoutFrame(l.Handle, bEnable)
	return l
}

// GetWidthIn 布局_取内宽度, 获取宽度, 不包含内边距大小.
func (l *LayoutEle) GetWidthIn() int32 {
	return xc.XLayout_GetWidthIn(l.Handle)
}

// GetHeightIn 布局_取内高度, 获取高度, 不包含内边距大小.
func (l *LayoutEle) GetHeightIn() int32 {
	return xc.XLayout_GetHeightIn(l.Handle)
}

/*
LayoutBox-布局盒子
*/

// EnableHorizon 布局盒子_启用水平.
//
// bEnable: 是否启用.
func (l *LayoutEle) EnableHorizon(bEnable bool) *LayoutEle {
	xc.XLayoutBox_EnableHorizon(l.Handle, bEnable)
	return l
}

// EnableAutoWrap 布局盒子_启用自动换行.
//
// bEnable: 是否启用.
func (l *LayoutEle) EnableAutoWrap(bEnable bool) *LayoutEle {
	xc.XLayoutBox_EnableAutoWrap(l.Handle, bEnable)
	return l
}

// EnableOverflowHide 布局盒子_启用溢出隐藏.
//
// bEnable: 是否启用.
func (l *LayoutEle) EnableOverflowHide(bEnable bool) *LayoutEle {
	xc.XLayoutBox_EnableOverflowHide(l.Handle, bEnable)
	return l
}

// SetAlignH 布局盒子_置水平对齐.
//
// nAlign: 对齐方式: xcc.Layout_Align_.
func (l *LayoutEle) SetAlignH(nAlign xcc.Layout_Align_) *LayoutEle {
	xc.XLayoutBox_SetAlignH(l.Handle, nAlign)
	return l
}

// SetAlignV 布局盒子_置垂直对齐.
//
// nAlign: 对齐方式: xcc.Layout_Align_.
func (l *LayoutEle) SetAlignV(nAlign xcc.Layout_Align_) *LayoutEle {
	xc.XLayoutBox_SetAlignV(l.Handle, nAlign)
	return l
}

// SetAlignBaseline 布局盒子_置对齐基线.
//
// nAlign: 对齐方式: xcc.Layout_Align_Axis_.
func (l *LayoutEle) SetAlignBaseline(nAlign xcc.Layout_Align_Axis_) *LayoutEle {
	xc.XLayoutBox_SetAlignBaseline(l.Handle, nAlign)
	return l
}

// SetSpace 布局盒子_置间距.
//
// nSpace: 项间距大小.
func (l *LayoutEle) SetSpace(nSpace int32) *LayoutEle {
	xc.XLayoutBox_SetSpace(l.Handle, nSpace)
	return l
}

// SetSpaceRow 布局盒子_置行距.
//
// nSpace: 行间距大小.
func (l *LayoutEle) SetSpaceRow(nSpace int32) *LayoutEle {
	xc.XLayoutBox_SetSpaceRow(l.Handle, nSpace)
	return l
}

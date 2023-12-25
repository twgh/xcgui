package widget

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// LayoutEle 布局元素.
type LayoutEle struct {
	Element
}

// NewLayoutEle 布局_创建, 创建布局元素.
//
//	@param x 元素x坐标.
//	@param y 元素y坐标.
//	@param cx 宽度.
//	@param cy 高度.
//	@param hParent 父为窗口句柄或元素句柄.
//	@return *LayoutEle
func NewLayoutEle(x int, y int, cx int, cy int, hParent int) *LayoutEle {
	p := &LayoutEle{}
	p.SetHandle(xc.XLayout_Create(x, y, cx, cy, hParent))
	return p
}

// NewLayoutEleEx 布局_创建扩展, 创建布局元素.
//
//	@param hParent 父为窗口句柄或元素句柄.
//	@return *LayoutEle
func NewLayoutEleEx(hParent int) *LayoutEle {
	p := &LayoutEle{}
	p.SetHandle(xc.XLayout_CreateEx(hParent))
	return p
}

// NewLayoutEleByHandle 从句柄创建对象.
//
//	@param handle
//	@return *LayoutEle
func NewLayoutEleByHandle(handle int) *LayoutEle {
	p := &LayoutEle{}
	p.SetHandle(handle)
	return p
}

// NewLayoutEleByLayout 从布局文件创建对象, 失败返回nil.
//
//	@param pFileName 布局文件名.
//	@param hParent 父对象句柄.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *LayoutEle
func NewLayoutEleByLayout(pFileName string, hParent int, hAttachWnd uintptr) *LayoutEle {
	handle := xc.XC_LoadLayout(pFileName, hParent, hAttachWnd)
	if handle > 0 {
		p := &LayoutEle{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewLayoutEleByLayoutZip 从压缩包中的布局文件创建对象, 失败返回nil.
//
//	@param pZipFileName zip文件名.
//	@param pFileName 布局文件名.
//	@param pPassword zip密码.
//	@param hParent 父对象句柄.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *LayoutEle
func NewLayoutEleByLayoutZip(pZipFileName string, pFileName string, pPassword string, hParent int, hAttachWnd uintptr) *LayoutEle {
	handle := xc.XC_LoadLayoutZip(pZipFileName, pFileName, pPassword, hParent, hAttachWnd)
	if handle > 0 {
		p := &LayoutEle{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewLayoutEleByLayoutZipResEx 从RC资源zip压缩包中的布局文件创建对象, 失败返回nil.
//
//	@param id RC资源ID.
//	@param pFileName zip文件名.
//	@param pPassword zip密码.
//	@param pPrefixName 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//	@param hParent 父对象句柄.
//	@param hParentWnd 父窗口句柄HWND, 提供给第三方窗口使用.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@param hModule 模块句柄, 可填0.
//	@return *LayoutEle
func NewLayoutEleByLayoutZipResEx(id int32, pFileName, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd, hModule uintptr) *LayoutEle {
	handle := xc.XC_LoadLayoutZipResEx(id, pFileName, pPassword, pPrefixName, hParent, hParentWnd, hAttachWnd, hModule)
	if handle > 0 {
		p := &LayoutEle{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewLayoutEleByLayoutZipMem 从内存压缩包中的布局文件创建对象, 失败返回nil.
//
//	@param data 布局文件数据.
//	@param pFileName 布局文件名.
//	@param pPassword zip密码.
//	@param hParent 父对象句柄.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *LayoutEle
func NewLayoutEleByLayoutZipMem(data []byte, pFileName string, pPassword string, hParent int, hAttachWnd uintptr) *LayoutEle {
	handle := xc.XC_LoadLayoutZipMem(data, pFileName, pPassword, hParent, hAttachWnd)
	if handle > 0 {
		p := &LayoutEle{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewLayoutEleByStringW 从布局文件字符串W创建对象, 失败返回nil.
//
//	@param pStringXML 字符串.
//	@param hParent 父对象.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *LayoutEle
func NewLayoutEleByStringW(pStringXML string, hParent int, hAttachWnd uintptr) *LayoutEle {
	handle := xc.XC_LoadLayoutFromStringW(pStringXML, hParent, hAttachWnd)
	if handle > 0 {
		p := &LayoutEle{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewLayoutEleByLayoutEx 从布局文件创建对象, 失败返回nil.
//
//	@param pFileName 布局文件名.
//	@param pPrefixName 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//	@param hParent 父对象句柄.
//	@param hParentWnd 父窗口句柄HWND, 提供给第三方窗口使用.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *LayoutEle
func NewLayoutEleByLayoutEx(pFileName, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *LayoutEle {
	handle := xc.XC_LoadLayoutEx(pFileName, pPrefixName, hParent, hParentWnd, hAttachWnd)
	if handle > 0 {
		p := &LayoutEle{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewLayoutEleByLayoutZipEx 从压缩包中的布局文件创建对象, 失败返回nil.
//
//	@param pZipFileName zip文件名.
//	@param pFileName 布局文件名.
//	@param pPassword zip密码.
//	@param pPrefixName 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//	@param hParent 父对象句柄.
//	@param hParentWnd 父窗口句柄HWND, 提供给第三方窗口使用.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *LayoutEle
func NewLayoutEleByLayoutZipEx(pZipFileName string, pFileName string, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *LayoutEle {
	handle := xc.XC_LoadLayoutZipEx(pZipFileName, pFileName, pPassword, pPrefixName, hParent, hParentWnd, hAttachWnd)
	if handle > 0 {
		p := &LayoutEle{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewLayoutEleByLayoutZipMemEx 从内存压缩包中的布局文件创建对象, 失败返回nil.
//
//	@param data 布局文件数据.
//	@param pFileName 布局文件名.
//	@param pPassword zip密码.
//	@param pPrefixName 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//	@param hParent 父对象句柄.
//	@param hParentWnd 父窗口句柄HWND, 提供给第三方窗口使用.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *LayoutEle
func NewLayoutEleByLayoutZipMemEx(data []byte, pFileName string, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *LayoutEle {
	handle := xc.XC_LoadLayoutZipMemEx(data, pFileName, pPassword, pPrefixName, hParent, hParentWnd, hAttachWnd)
	if handle > 0 {
		p := &LayoutEle{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewLayoutEleByStringWEx 从布局文件字符串W创建对象, 失败返回nil.
//
//	@param pStringXML 字符串.
//	@param pPrefixName 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//	@param hParent 父对象.
//	@param hParentWnd 父窗口句柄HWND, 提供给第三方窗口使用.
//	@param hAttachWnd 附加窗口句柄, 附加到指定的窗口, 可填0.
//	@return *LayoutEle
func NewLayoutEleByStringWEx(pStringXML, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) *LayoutEle {
	handle := xc.XC_LoadLayoutFromStringWEx(pStringXML, pPrefixName, hParent, hParentWnd, hAttachWnd)
	if handle > 0 {
		p := &LayoutEle{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewLayoutEleByName 从name创建对象, 失败返回nil.
//
//	@param name
//	@return *LayoutEle
func NewLayoutEleByName(name string) *LayoutEle {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &LayoutEle{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewLayoutEleByID 从ID创建对象, 失败返回nil.
//
//	@param hWindow 父窗口句柄
//	@param nID
//	@return *LayoutEle
func NewLayoutEleByID(hWindow, nID int) *LayoutEle {
	handle := xc.XC_GetObjectByID(hWindow, nID)
	if handle > 0 {
		p := &LayoutEle{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewLayoutEleByIDName 从ID名称创建对象, 失败返回nil.
//
//	@param hWindow 父窗口句柄
//	@param name
//	@return *LayoutEle
func NewLayoutEleByIDName(hWindow int, name string) *LayoutEle {
	handle := xc.XC_GetObjectByIDName(hWindow, name)
	if handle > 0 {
		p := &LayoutEle{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewLayoutEleByUID 从UID创建对象, 失败返回nil.
//
//	@param nUID
//	@return *LayoutEle
func NewLayoutEleByUID(nUID int) *LayoutEle {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &LayoutEle{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewLayoutEleByUIDName 从UID名称创建对象, 失败返回nil.
//
//	@param name
//	@return *LayoutEle
func NewLayoutEleByUIDName(name string) *LayoutEle {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &LayoutEle{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// IsEnableLayout 布局_判断启用, 是否已经启用布局功能.
//
//	@return bool
func (l *LayoutEle) IsEnableLayout() bool {
	return xc.XLayout_IsEnableLayout(l.Handle)
}

// EnableLayout 布局_启用, 启用布局功能.
//
//	@param bEnable 是否启用.
//	@return int
func (l *LayoutEle) EnableLayout(bEnable bool) int {
	return xc.XLayout_EnableLayout(l.Handle, bEnable)
}

// ShowLayoutFrame 布局_显示布局边界, 显示布局边界.
//
//	@param bEnable 是否显示.
//	@return int
func (l *LayoutEle) ShowLayoutFrame(bEnable bool) int {
	return xc.XLayout_ShowLayoutFrame(l.Handle, bEnable)
}

// GetWidthIn 布局_取内宽度, 获取宽度,不包含内边距大小.
//
//	@return int
func (l *LayoutEle) GetWidthIn() int {
	return xc.XLayout_GetWidthIn(l.Handle)
}

// GetHeightIn 布局_取内高度, 获取高度,不包含内边距大小.
//
//	@return int
func (l *LayoutEle) GetHeightIn() int {
	return xc.XLayout_GetHeightIn(l.Handle)
}

/*
LayoutBox-布局盒子
*/

// EnableHorizon 布局盒子_启用水平.
//
//	@param bEnable 是否启用.
//	@return int
func (l *LayoutEle) EnableHorizon(bEnable bool) int {
	return xc.XLayoutBox_EnableHorizon(l.Handle, bEnable)
}

// EnableAutoWrap 布局盒子_启用自动换行.
//
//	@param bEnable 是否启用.
//	@return int
func (l *LayoutEle) EnableAutoWrap(bEnable bool) int {
	return xc.XLayoutBox_EnableAutoWrap(l.Handle, bEnable)
}

// EnableOverflowHide 布局盒子_启用溢出隐藏.
//
//	@param bEnable 是否启用.
//	@return int
func (l *LayoutEle) EnableOverflowHide(bEnable bool) int {
	return xc.XLayoutBox_EnableOverflowHide(l.Handle, bEnable)
}

// SetAlignH 布局盒子_置水平对齐.
//
//	@param nAlign 对齐方式: xcc.Layout_Align_.
//	@return int
func (l *LayoutEle) SetAlignH(nAlign xcc.Layout_Align_) int {
	return xc.XLayoutBox_SetAlignH(l.Handle, nAlign)
}

// SetAlignV 布局盒子_置垂直对齐.
//
//	@param nAlign 对齐方式: xcc.Layout_Align_.
//	@return int
func (l *LayoutEle) SetAlignV(nAlign xcc.Layout_Align_) int {
	return xc.XLayoutBox_SetAlignV(l.Handle, nAlign)
}

// SetAlignBaseline 布局盒子_置对齐基线.
//
//	@param nAlign 对齐方式: xcc.Layout_Align_Axis_.
//	@return int
func (l *LayoutEle) SetAlignBaseline(nAlign xcc.Layout_Align_Axis_) int {
	return xc.XLayoutBox_SetAlignBaseline(l.Handle, nAlign)
}

// SetSpace 布局盒子_置间距.
//
//	@param nSpace 项间距大小.
//	@return int
func (l *LayoutEle) SetSpace(nSpace int) int {
	return xc.XLayoutBox_SetSpace(l.Handle, nSpace)
}

// SetSpaceRow 布局盒子_置行距.
//
//	@param nSpace 行间距大小.
//	@return int
func (l *LayoutEle) SetSpaceRow(nSpace int) int {
	return xc.XLayoutBox_SetSpaceRow(l.Handle, nSpace)
}

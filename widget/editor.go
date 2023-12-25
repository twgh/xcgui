package widget

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// 代码编辑框.
type Editor struct {
	Edit
}

// 代码编辑框_创建.
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
func NewEditor(x int, y int, cx int, cy int, hParent int) *Editor {
	p := &Editor{}
	p.SetHandle(xc.XEditor_Create(x, y, cx, cy, hParent))
	return p
}

// 从句柄创建对象.
func NewEditorByHandle(handle int) *Editor {
	p := &Editor{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func NewEditorByName(name string) *Editor {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &Editor{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func NewEditorByUID(nUID int) *Editor {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &Editor{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func NewEditorByUIDName(name string) *Editor {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &Editor{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 代码编辑框_启用空格选择自动匹配项.
//
// bEnable: 是否启用.
func (e *Editor) EnableAutoMatchSpaseSelect(bEnable bool) int {
	return xc.XEidtor_EnableAutoMatchSpaseSelect(e.Handle, bEnable)
}

// 代码编辑框_判断断点.
//
// iRow: 行索引.
func (e *Editor) IsBreakpoint(iRow int) bool {
	return xc.XEditor_IsBreakpoint(e.Handle, iRow)
}

// 代码编辑框_置断点.
//
// iRow: 行索引.
//
// bActivate: 是否激活.
func (e *Editor) SetBreakpoint(iRow int, bActivate bool) bool {
	return xc.XEditor_SetBreakpoint(e.Handle, iRow, bActivate)
}

// 代码编辑框_移除断点.
//
// iRow: 行索引.
func (e *Editor) RemoveBreakpoint(iRow int) bool {
	return xc.XEditor_RemoveBreakpoint(e.Handle, iRow)
}

// 代码编辑框_清空断点.
func (e *Editor) ClearBreakpoint() int {
	return xc.XEditor_ClearBreakpoint(e.Handle)
}

// 代码编辑框_置当前运行.
//
// iRow: 行索引.
func (e *Editor) SetRunRow(iRow int) bool {
	return xc.XEditor_SetRunRow(e.Handle, iRow)
}

// 代码编辑框_取颜色信息.
//
// pInfo: 颜色信息结构体指针.
func (e *Editor) GetColor(pInfo *xc.Editor_Color_) int {
	return xc.XEditor_GetColor(e.Handle, pInfo)
}

// 代码编辑框_置颜色.
//
// pInfo: 颜色信息结构体指针.
func (e *Editor) SetColor(pInfo *xc.Editor_Color_) int {
	return xc.XEditor_SetColor(e.Handle, pInfo)
}

// 代码编辑框_置常量样式.
//
// iStyle: 样式.
func (e *Editor) SetStyleKeyword(iStyle int) int {
	return xc.XEditor_SetStyleKeyword(e.Handle, iStyle)
}

// 代码编辑框_置函数样式.
//
// iStyle: 样式.
func (e *Editor) SetStyleFunction(iStyle int) int {
	return xc.XEditor_SetStyleFunction(e.Handle, iStyle)
}

// 代码编辑框_置变量样式.
//
// iStyle: 样式.
func (e *Editor) SetStyleVar(iStyle int) int {
	return xc.XEditor_SetStyleVar(e.Handle, iStyle)
}

// 代码编辑框_置数据类型样式.
//
// iStyle: 样式.
func (e *Editor) SetStyleDataType(iStyle int) int {
	return xc.XEditor_SetStyleDataType(e.Handle, iStyle)
}

// 代码编辑框_置类样式.
//
// iStyle: 样式.
func (e *Editor) SetStyleClass(iStyle int) int {
	return xc.XEditor_SetStyleClass(e.Handle, iStyle)
}

// 代码编辑框_置宏样式.
//
// iStyle: 样式.
func (e *Editor) SetStyleMacro(iStyle int) int {
	return xc.XEditor_SetStyleMacro(e.Handle, iStyle)
}

// 代码编辑框_置字符串样式.
//
// iStyle: 样式.
func (e *Editor) SetStyleString(iStyle int) int {
	return xc.XEditor_SetStyleString(e.Handle, iStyle)
}

// 代码编辑框_置注释样式.
//
// iStyle: 样式.
func (e *Editor) SetStyleComment(iStyle int) int {
	return xc.XEditor_SetStyleComment(e.Handle, iStyle)
}

// 代码编辑框_置数字样式.
//
// iStyle: 样式.
func (e *Editor) SetStyleNumber(iStyle int) int {
	return xc.XEditor_SetStyleNumber(e.Handle, iStyle)
}

// 代码编辑框_取断点数量.
func (e *Editor) GetBreakpointCount() int {
	return xc.XEditor_GetBreakpointCount(e.Handle)
}

// 代码编辑框_取全部断点, 返回实际获取断点数量.
//
// aPoints: 接收断点数组.
//
// nCount: 数组大小.
func (e *Editor) GetBreakpoints(aPoints *[]int32, nCount int) int {
	return xc.XEditor_GetBreakpoints(e.Handle, aPoints, nCount)
}

// 代码编辑框_设置当前行, 跳过收缩行.
//
// iRow: 行索引.
func (e *Editor) SetCurRow(iRow int) int {
	return xc.XEditor_SetCurRow(e.Handle, iRow)
}

// 代码编辑框_获取深度.
//
// iRow: 行索引.
func (e *Editor) GetDepth(iRow int) int {
	return xc.XEditor_GetDepth(e.Handle, iRow)
}

// 代码编辑框_转换到展开行, 跳过收缩行.
//
// iRow: 行索引.
func (e *Editor) ToExpandRow(iRow int) int {
	return xc.XEditor_ToExpandRow(e.Handle, iRow)
}

// 代码编辑框_展开扩展, 完全展开指定行, 例如:行包含在折叠内容中, 将其展开.
//
// iRow: 行索引.
func (e *Editor) ExpandEx(iRow int) int {
	return xc.XEditor_ExpandEx(e.Handle, iRow)
}

// 代码编辑框_展开全部.
//
// bExpand: 是否展开.
func (e *Editor) ExpandAll(bExpand bool) int {
	return xc.XEditor_ExpandAll(e.Handle, bExpand)
}

// 代码编辑框_展开指定行.
//
// iRow: 行索引.
//
// bExpand: 是否展开.
func (e *Editor) Expand(iRow int, bExpand bool) int {
	return xc.XEditor_Expand(e.Handle, iRow, bExpand)
}

// 代码编辑框_添加关键字.
//
// pKey: 字符串.
//
// iStyle: 样式.
func (e *Editor) AddKeyword(pKey string, iStyle int) int {
	return xc.XEditor_AddKeyword(e.Handle, pKey, iStyle)
}

// 代码编辑框_添加自动匹配字符串.
//
// pKey: 字符串.
func (e *Editor) AddConst(pKey string) int {
	return xc.XEditor_AddConst(e.Handle, pKey)
}

// 代码编辑框_添加自动匹配函数.
//
// pKey: 字符串.
func (e *Editor) AddFunction(pKey string) int {
	return xc.XEditor_AddFunction(e.Handle, pKey)
}

// 代码编辑框_添加排除定义变量关键字, 排除定义变量的关键字, 用于排除定义变量, 因为定义变量禁用自动匹配; 此关键字不加入自动匹配,仅用于排除定义变量.
//
// pKeyword: 字符串.
func (e *Editor) AddExcludeDefVarKeyword(pKeyword string) int {
	return xc.XEditor_AddExcludeDefVarKeyword(e.Handle, pKeyword)
}

// 代码编辑框_获取折叠状态.
func (e *Editor) GetExpandState() string {
	return xc.XEditor_GetExpandState(e.Handle)
}

// 代码编辑框_设置折叠状态.
//
// pString: .
func (e *Editor) SetExpandState(pString string) int {
	return xc.XEditor_SetExpandState(e.Handle, pString)
}

// 代码编辑框_获取缩进.
//
// iRow: 行.
func (e *Editor) GetIndentation(iRow int) int {
	return xc.XEditor_GetIndentation(e.Handle, iRow)
}

// 代码编辑框_是否为空行.
//
// iRow: 行.
func (e *Editor) IsEmptyRow(iRow int) int {
	return xc.XEidtor_IsEmptyRow(e.Handle, iRow)
}

// 代码编辑框_置自动匹配结果显示模式.
//
// mode: 0:中英文, 1:英文, 3:中文.
func (e *Editor) SetAutoMatchMode(mode int) int {
	return xc.XEditor_SetAutoMatchMode(e.Handle, mode)
}

/*
以下都是事件
*/

type XE_EDITOR_MODIFY_ROWS func(iRow int32, nRows int32, pbHandled *bool) int                 // 多行内容改变事件 例如:区块注释操作, 区块缩进操作, 代码格式化. iRow: 开始行. nRows: 改变行数量.
type XE_EDITOR_MODIFY_ROWS1 func(hEle int, iRow int32, nRows int32, pbHandled *bool) int      // 多行内容改变事件 例如:区块注释操作, 区块缩进操作, 代码格式化. iRow: 开始行. nRows: 改变行数量.
type XE_EDITOR_SETBREAKPOINT func(iRow int32, bCheck bool, pbHandled *bool) int               // 代码编辑框_设置断点.
type XE_EDITOR_SETBREAKPOINT1 func(hEle int, iRow int32, bCheck bool, pbHandled *bool) int    // 代码编辑框_设置断点.
type XE_EDITOR_REMOVEBREAKPOINT func(iRow int32, pbHandled *bool) int                         // 代码编辑框_移除断点.
type XE_EDITOR_REMOVEBREAKPOINT1 func(hEle int, iRow int32, pbHandled *bool) int              // 代码编辑框_移除断点.
type XE_EDITOR_AUTOMATCH_SELECT func(iRow int32, nRows int32, pbHandled *bool) int            // 代码编辑框_自动匹配选择.
type XE_EDITOR_AUTOMATCH_SELECT1 func(hEle int, iRow int32, nRows int32, pbHandled *bool) int // 代码编辑框_自动匹配选择.

// 多行内容改变事件 例如:区块注释操作, 区块缩进操作, 代码格式化.
func (e *Editor) Event_EDITOR_MODIFY_ROWS(pFun XE_EDITOR_MODIFY_ROWS) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_EDITOR_MODIFY_ROWS, pFun)
}

// 多行内容改变事件 例如:区块注释操作, 区块缩进操作, 代码格式化.
func (e *Editor) Event_EDITOR_MODIFY_ROWS1(pFun XE_EDITOR_MODIFY_ROWS1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_EDITOR_MODIFY_ROWS, pFun)
}

// 代码编辑框_设置断点.
func (e *Editor) Event_EDITOR_SETBREAKPOINT(pFun XE_EDITOR_SETBREAKPOINT) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_EDITOR_SETBREAKPOINT, pFun)
}

// 代码编辑框_设置断点.
func (e *Editor) Event_EDITOR_SETBREAKPOINT1(pFun XE_EDITOR_SETBREAKPOINT1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_EDITOR_SETBREAKPOINT, pFun)
}

// 代码编辑框_移除断点.
func (e *Editor) Event_EDITOR_REMOVEBREAKPOINT(pFun XE_EDITOR_REMOVEBREAKPOINT) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_EDITOR_REMOVEBREAKPOINT, pFun)
}

// 代码编辑框_移除断点.
func (e *Editor) Event_EDITOR_REMOVEBREAKPOINT1(pFun XE_EDITOR_REMOVEBREAKPOINT1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_EDITOR_REMOVEBREAKPOINT, pFun)
}

// 代码编辑框_自动匹配选择.
func (e *Editor) Event_EDITOR_AUTOMATCH_SELECT(pFun XE_EDITOR_AUTOMATCH_SELECT) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_EDITOR_AUTOMATCH_SELECT, pFun)
}

// 代码编辑框_自动匹配选择.
func (e *Editor) Event_EDITOR_AUTOMATCH_SELECT1(pFun XE_EDITOR_AUTOMATCH_SELECT1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_EDITOR_AUTOMATCH_SELECT, pFun)
}

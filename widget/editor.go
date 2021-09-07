package widget

import (
	"github.com/twgh/xcgui/xc"
)

type Editor struct {
	Element

	HEle int
}

// 代码编辑框_创建, 返回元素句柄
// x: 元素x坐标.
// y: 元素y坐标.
// cx: 宽度.
// cy: 高度.
// hParent: 父为窗口句柄或元素句柄.
func NewEditor(x int, y int, cx int, cy int, hParent int) *Editor {
	p := &Editor{
		HEle: xc.XEditor_Create(x, y, cx, cy, hParent),
	}
	p.HEle_ = p.HEle
	return p
}

// 代码编辑框_启用空格选择自动匹配项
// bEnable: 是否启用
func (e *Editor) EnableAutoMatchSpaseSelect(bEnable bool) int {
	return xc.XEidtor_EnableAutoMatchSpaseSelect(e.HEle, bEnable)
}

// 代码编辑框_判断断点
// iRow: 行索引
func (e *Editor) IsBreakpoint(iRow int) bool {
	return xc.XEditor_IsBreakpoint(e.HEle, iRow)
}

// 代码编辑框_置断点
// iRow: 行索引
// bActivate: 是否激活
func (e *Editor) SetBreakpoint(iRow int, bActivate bool) bool {
	return xc.XEditor_SetBreakpoint(e.HEle, iRow, bActivate)
}

// 代码编辑框_移除断点
// iRow: 行索引
func (e *Editor) RemoveBreakpoint(iRow int) bool {
	return xc.XEditor_RemoveBreakpoint(e.HEle, iRow)
}

// 代码编辑框_清空断点
func (e *Editor) ClearBreakpoint() int {
	return xc.XEditor_ClearBreakpoint(e.HEle)
}

// 代码编辑框_置当前运行
// iRow: 行索引
func (e *Editor) SetRunRow(iRow int) bool {
	return xc.XEditor_SetRunRow(e.HEle, iRow)
}

// 代码编辑框_取颜色信息
// pInfo: 颜色信息结构体指针
func (e *Editor) GetColor(pInfo *xc.Editor_Color_) int {
	return xc.XEditor_GetColor(e.HEle, pInfo)
}

// 代码编辑框_置颜色
// pInfo: 颜色信息结构体指针
func (e *Editor) SetColor(pInfo *xc.Editor_Color_) int {
	return xc.XEditor_SetColor(e.HEle, pInfo)
}

// 代码编辑框_置常量样式
// iStyle: 样式
func (e *Editor) SetStyleKeyword(iStyle int) int {
	return xc.XEditor_SetStyleKeyword(e.HEle, iStyle)
}

// 代码编辑框_置函数样式
// iStyle: 样式
func (e *Editor) SetStyleFunction(iStyle int) int {
	return xc.XEditor_SetStyleFunction(e.HEle, iStyle)
}

// 代码编辑框_置变量样式
// iStyle: 样式
func (e *Editor) SetStyleVar(iStyle int) int {
	return xc.XEditor_SetStyleVar(e.HEle, iStyle)
}

// 代码编辑框_置数据类型样式
// iStyle: 样式
func (e *Editor) SetStyleDataType(iStyle int) int {
	return xc.XEditor_SetStyleDataType(e.HEle, iStyle)
}

// 代码编辑框_置类样式
// iStyle: 样式
func (e *Editor) SetStyleClass(iStyle int) int {
	return xc.XEditor_SetStyleClass(e.HEle, iStyle)
}

// 代码编辑框_置宏样式
// iStyle: 样式
func (e *Editor) SetStyleMacro(iStyle int) int {
	return xc.XEditor_SetStyleMacro(e.HEle, iStyle)
}

// 代码编辑框_置字符串样式
// iStyle: 样式
func (e *Editor) SetStyleString(iStyle int) int {
	return xc.XEditor_SetStyleString(e.HEle, iStyle)
}

// 代码编辑框_置注释样式
// iStyle: 样式
func (e *Editor) SetStyleComment(iStyle int) int {
	return xc.XEditor_SetStyleComment(e.HEle, iStyle)
}

// 代码编辑框_置数字样式
// iStyle: 样式
func (e *Editor) SetStyleNumber(iStyle int) int {
	return xc.XEditor_SetStyleNumber(e.HEle, iStyle)
}

// 代码编辑框_取断点数量
func (e *Editor) GetBreakpointCount() int {
	return xc.XEditor_GetBreakpointCount(e.HEle)
}

// 代码编辑框_取全部断点, 返回实际获取断点数量
// aPoints: 接收断点数组
// nCount: 数组大小
func (e *Editor) GetBreakpoints(aPoints int, nCount int) int {
	return xc.XEditor_GetBreakpoints(e.HEle, aPoints, nCount)
}

// 代码编辑框_设置当前行, 跳过收缩行
// iRow: 行索引
func (e *Editor) SetCurRow(iRow int) int {
	return xc.XEditor_SetCurRow(e.HEle, iRow)
}

// 代码编辑框_获取深度
// iRow: 行索引
func (e *Editor) GetDepth(iRow int) int {
	return xc.XEditor_GetDepth(e.HEle, iRow)
}

// 代码编辑框_转换到展开行, 跳过收缩行
// iRow: 行索引
func (e *Editor) ToExpandRow(iRow int) int {
	return xc.XEditor_ToExpandRow(e.HEle, iRow)
}

// 代码编辑框_展开扩展, 完全展开指定行, 例如:行包含在折叠内容中, 将其展开
// iRow: 行索引
func (e *Editor) ExpandEx(iRow int) int {
	return xc.XEditor_ExpandEx(e.HEle, iRow)
}

// 代码编辑框_展开全部
// bExpand: 是否展开
func (e *Editor) ExpandAll(bExpand bool) int {
	return xc.XEditor_ExpandAll(e.HEle, bExpand)
}

// 代码编辑框_展开指定行
// iRow: 行索引
// bExpand: 是否展开
func (e *Editor) Expand(iRow int, bExpand bool) int {
	return xc.XEditor_Expand(e.HEle, iRow, bExpand)
}

// 代码编辑框_添加关键字
// pKey: 字符串
// iStyle: 样式
func (e *Editor) AddKeyword(pKey string, iStyle int) int {
	return xc.XEditor_AddKeyword(e.HEle, pKey, iStyle)
}

// 代码编辑框_添加自动匹配字符串
// pKey: 字符串
func (e *Editor) AddConst(pKey string) int {
	return xc.XEditor_AddConst(e.HEle, pKey)
}

// 代码编辑框_添加自动匹配函数
// pKey: 字符串
func (e *Editor) AddFunction(pKey string) int {
	return xc.XEditor_AddFunction(e.HEle, pKey)
}

// 代码编辑框_添加排除定义变量关键字, 排除定义变量的关键字, 用于排除定义变量, 因为定义变量禁用自动匹配; 此关键字不加入自动匹配,仅用于排除定义变量
// pKeyword: 字符串
func (e *Editor) AddExcludeDefVarKeyword(pKeyword string) int {
	return xc.XEditor_AddExcludeDefVarKeyword(e.HEle, pKeyword)
}

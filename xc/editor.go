package xc

import "unsafe"

// 代码编辑框_创建, 返回元素句柄.
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
func XEditor_Create(x int, y int, cx int, cy int, hParent int) int {
	r, _, _ := xEditor_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent))
	return int(r)
}

// 代码编辑框_启用空格选择自动匹配项.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XEidtor_EnableAutoMatchSpaseSelect(hEle int, bEnable bool) int {
	r, _, _ := xEidtor_EnableAutoMatchSpaseSelect.Call(uintptr(hEle), boolPtr(bEnable))
	return int(r)
}

// 代码编辑框_判断断点.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
func XEditor_IsBreakpoint(hEle int, iRow int) bool {
	r, _, _ := xEditor_IsBreakpoint.Call(uintptr(hEle), uintptr(iRow))
	return int(r) != 0
}

// 代码编辑框_置断点.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
//
// bActivate: 是否激活.
func XEditor_SetBreakpoint(hEle int, iRow int, bActivate bool) bool {
	r, _, _ := xEditor_SetBreakpoint.Call(uintptr(hEle), uintptr(iRow), boolPtr(bActivate))
	return int(r) != 0
}

// 代码编辑框_移除断点.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
func XEditor_RemoveBreakpoint(hEle int, iRow int) bool {
	r, _, _ := xEditor_RemoveBreakpoint.Call(uintptr(hEle), uintptr(iRow))
	return int(r) != 0
}

// 代码编辑框_清空断点.
//
// hEle: 元素句柄.
func XEditor_ClearBreakpoint(hEle int) int {
	r, _, _ := xEditor_ClearBreakpoint.Call(uintptr(hEle))
	return int(r)
}

// 代码编辑框_置当前运行.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
func XEditor_SetRunRow(hEle int, iRow int) bool {
	r, _, _ := xEditor_SetRunRow.Call(uintptr(hEle), uintptr(iRow))
	return int(r) != 0
}

// 代码编辑框_取颜色信息.
//
// hEle: 元素句柄.
//
// pInfo: 颜色信息结构体指针.
func XEditor_GetColor(hEle int, pInfo *Editor_Color_) int {
	r, _, _ := xEditor_GetColor.Call(uintptr(hEle), uintptr(unsafe.Pointer(pInfo)))
	return int(r)
}

// 代码编辑框_置颜色.
//
// hEle: 元素句柄.
//
// pInfo: 颜色信息结构体指针.
func XEditor_SetColor(hEle int, pInfo *Editor_Color_) int {
	r, _, _ := xEditor_SetColor.Call(uintptr(hEle), uintptr(unsafe.Pointer(pInfo)))
	return int(r)
}

// 代码编辑框_置常量样式.
//
// hEle: 元素句柄.
//
// iStyle: 样式.
func XEditor_SetStyleKeyword(hEle int, iStyle int) int {
	r, _, _ := xEditor_SetStyleKeyword.Call(uintptr(hEle), uintptr(iStyle))
	return int(r)
}

// 代码编辑框_置函数样式.
//
// hEle: 元素句柄.
//
// iStyle: 样式.
func XEditor_SetStyleFunction(hEle int, iStyle int) int {
	r, _, _ := xEditor_SetStyleFunction.Call(uintptr(hEle), uintptr(iStyle))
	return int(r)
}

// 代码编辑框_置变量样式.
//
// hEle: 元素句柄.
//
// iStyle: 样式.
func XEditor_SetStyleVar(hEle int, iStyle int) int {
	r, _, _ := xEditor_SetStyleVar.Call(uintptr(hEle), uintptr(iStyle))
	return int(r)
}

// 代码编辑框_置数据类型样式.
//
// hEle: 元素句柄.
//
// iStyle: 样式.
func XEditor_SetStyleDataType(hEle int, iStyle int) int {
	r, _, _ := xEditor_SetStyleDataType.Call(uintptr(hEle), uintptr(iStyle))
	return int(r)
}

// 代码编辑框_置类样式.
//
// hEle: 元素句柄.
//
// iStyle: 样式.
func XEditor_SetStyleClass(hEle int, iStyle int) int {
	r, _, _ := xEditor_SetStyleClass.Call(uintptr(hEle), uintptr(iStyle))
	return int(r)
}

// 代码编辑框_置宏样式.
//
// hEle: 元素句柄.
//
// iStyle: 样式.
func XEditor_SetStyleMacro(hEle int, iStyle int) int {
	r, _, _ := xEditor_SetStyleMacro.Call(uintptr(hEle), uintptr(iStyle))
	return int(r)
}

// 代码编辑框_置字符串样式.
//
// hEle: 元素句柄.
//
// iStyle: 样式.
func XEditor_SetStyleString(hEle int, iStyle int) int {
	r, _, _ := xEditor_SetStyleString.Call(uintptr(hEle), uintptr(iStyle))
	return int(r)
}

// 代码编辑框_置注释样式.
//
// hEle: 元素句柄.
//
// iStyle: 样式.
func XEditor_SetStyleComment(hEle int, iStyle int) int {
	r, _, _ := xEditor_SetStyleComment.Call(uintptr(hEle), uintptr(iStyle))
	return int(r)
}

// 代码编辑框_置数字样式.
//
// hEle: 元素句柄.
//
// iStyle: 样式.
func XEditor_SetStyleNumber(hEle int, iStyle int) int {
	r, _, _ := xEditor_SetStyleNumber.Call(uintptr(hEle), uintptr(iStyle))
	return int(r)
}

// 代码编辑框_取断点数量.
//
// hEle: 元素句柄.
func XEditor_GetBreakpointCount(hEle int) int {
	r, _, _ := xEditor_GetBreakpointCount.Call(uintptr(hEle))
	return int(r)
}

// 代码编辑框_取全部断点, 返回实际获取断点数量.
//
// hEle: 元素句柄.
//
// aPoints: 接收断点数组.
//
// nCount: 数组大小.
func XEditor_GetBreakpoints(hEle int, aPoints int, nCount int) int {
	r, _, _ := xEditor_GetBreakpoints.Call(uintptr(hEle), uintptr(aPoints), uintptr(nCount))
	return int(r)
}

// 代码编辑框_设置当前行, 跳过收缩行.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
func XEditor_SetCurRow(hEle int, iRow int) int {
	r, _, _ := xEditor_SetCurRow.Call(uintptr(hEle), uintptr(iRow))
	return int(r)
}

// 代码编辑框_获取深度.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
func XEditor_GetDepth(hEle int, iRow int) int {
	r, _, _ := xEditor_GetDepth.Call(uintptr(hEle), uintptr(iRow))
	return int(r)
}

// 代码编辑框_转换到展开行, 跳过收缩行.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
func XEditor_ToExpandRow(hEle int, iRow int) int {
	r, _, _ := xEditor_ToExpandRow.Call(uintptr(hEle), uintptr(iRow))
	return int(r)
}

// 代码编辑框_展开扩展, 完全展开指定行, 例如:行包含在折叠内容中, 将其展开.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
func XEditor_ExpandEx(hEle int, iRow int) int {
	r, _, _ := xEditor_ExpandEx.Call(uintptr(hEle), uintptr(iRow))
	return int(r)
}

// 代码编辑框_展开全部.
//
// hEle: 元素句柄.
//
// bExpand: 是否展开.
func XEditor_ExpandAll(hEle int, bExpand bool) int {
	r, _, _ := xEditor_ExpandAll.Call(uintptr(hEle), boolPtr(bExpand))
	return int(r)
}

// 代码编辑框_展开指定行.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
//
// bExpand: 是否展开.
func XEditor_Expand(hEle int, iRow int, bExpand bool) int {
	r, _, _ := xEditor_Expand.Call(uintptr(hEle), uintptr(iRow), boolPtr(bExpand))
	return int(r)
}

// 代码编辑框_添加关键字.
//
// hEle: 元素句柄.
//
// pKey: 字符串.
//
// iStyle: 样式.
func XEditor_AddKeyword(hEle int, pKey string, iStyle int) int {
	r, _, _ := xEditor_AddKeyword.Call(uintptr(hEle), strPtr(pKey), uintptr(iStyle))
	return int(r)
}

// 代码编辑框_添加自动匹配字符串.
//
// hEle: 元素句柄.
//
// pKey: 字符串.
func XEditor_AddConst(hEle int, pKey string) int {
	r, _, _ := xEditor_AddConst.Call(uintptr(hEle), strPtr(pKey))
	return int(r)
}

// 代码编辑框_添加自动匹配函数.
//
// hEle: 元素句柄.
//
// pKey: 字符串.
func XEditor_AddFunction(hEle int, pKey string) int {
	r, _, _ := xEditor_AddFunction.Call(uintptr(hEle), strPtr(pKey))
	return int(r)
}

// 代码编辑框_添加排除定义变量关键字, 排除定义变量的关键字, 用于排除定义变量, 因为定义变量禁用自动匹配; 此关键字不加入自动匹配,仅用于排除定义变量.
//
// hEle: 元素句柄.
//
// pKeyword: 字符串.
func XEditor_AddExcludeDefVarKeyword(hEle int, pKeyword string) int {
	r, _, _ := xEditor_AddExcludeDefVarKeyword.Call(uintptr(hEle), strPtr(pKeyword))
	return int(r)
}

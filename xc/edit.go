package xc

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"

	"github.com/twgh/xcgui/xcc"
)

// 编辑框_创建, 返回元素句柄.
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
func XEdit_Create(x int, y int, cx int, cy int, hParent int) int {
	r, _, _ := xEdit_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent))
	return int(r)
}

// 编辑框_创建扩展, 返回元素句柄.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// nType: 类型, Edit_Type_.
//
// hParent: 父为窗口句柄或元素句柄.
func XEdit_CreateEx(x int, y int, cx int, cy int, nType xcc.Edit_Type_, hParent int) int {
	r, _, _ := xEdit_CreateEx.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(nType), uintptr(hParent))
	return int(r)
}

// 编辑框_启用自动换行.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XEdit_EnableAutoWrap(hEle int, bEnable bool) int {
	r, _, _ := xEdit_EnableAutoWrap.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 编辑框_启用只读.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XEdit_EnableReadOnly(hEle int, bEnable bool) int {
	r, _, _ := xEdit_EnableReadOnly.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 编辑框_启用多行.
//
// hEle:.
//
// bEnable:.
func XEdit_EnableMultiLine(hEle int, bEnable bool) int {
	r, _, _ := xEdit_EnableMultiLine.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 编辑框_启用密码, 启用密码模式(只支持默认类型编辑框).
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XEdit_EnablePassword(hEle int, bEnable bool) int {
	r, _, _ := xEdit_EnablePassword.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 编辑框_启用自动选择, 当获得焦点时,自动选择所有内容.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XEdit_EnableAutoSelAll(hEle int, bEnable bool) int {
	r, _, _ := xEdit_EnableAutoSelAll.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 编辑框_启用自动取消选择, 当失去焦点时自动取消选择.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func XEdit_EnableAutoCancelSel(hEle int, bEnable bool) int {
	r, _, _ := xEdit_EnableAutoCancelSel.Call(uintptr(hEle), common.BoolPtr(bEnable))
	return int(r)
}

// 编辑框_是否只读.
//
// hEle: 元素句柄.
func XEdit_IsReadOnly(hEle int) bool {
	r, _, _ := xEdit_IsReadOnly.Call(uintptr(hEle))
	return r != 0
}

// 编辑框_是否多行.
//
// hEle: 元素句柄.
func XEdit_IsMultiLine(hEle int) bool {
	r, _, _ := xEdit_IsMultiLine.Call(uintptr(hEle))
	return r != 0
}

// 编辑框_是否密码.
//
// hEle: 元素句柄.
func XEdit_IsPassword(hEle int) bool {
	r, _, _ := xEdit_IsPassword.Call(uintptr(hEle))
	return r != 0
}

// 编辑框_是否自动换行.
//
// hEle: 元素句柄.
func XEdit_IsAutoWrap(hEle int) bool {
	r, _, _ := xEdit_IsAutoWrap.Call(uintptr(hEle))
	return r != 0
}

// 编辑框_判断为空.
//
// hEle: 元素句柄.
func XEdit_IsEmpty(hEle int) bool {
	r, _, _ := xEdit_IsEmpty.Call(uintptr(hEle))
	return r != 0
}

// 编辑框_是否在选择区域.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
func XEdit_IsInSelect(hEle int, iRow int, iCol int) bool {
	r, _, _ := xEdit_IsInSelect.Call(uintptr(hEle), uintptr(iRow), uintptr(iCol))
	return r != 0
}

// 编辑框_取总行数.
//
// hEle: 元素句柄.
func XEdit_GetRowCount(hEle int) int {
	r, _, _ := xEdit_GetRowCount.Call(uintptr(hEle))
	return int(r)
}

// 编辑框_取数据, 包含文本或非文本内容.
//
// hEle: 元素句柄.
func XEdit_GetData(hEle int) Edit_Data_Copy_ {
	r, _, _ := xEdit_GetData.Call(uintptr(hEle))
	return *(*Edit_Data_Copy_)(unsafe.Pointer(&r))
}

// 编辑框_添加数据.
//
// hEle: 元素句柄.
//
// pData: 数据结构.
//
// styleTable: 样式表.
//
// nStyleCount: 样式数量.
func XEdit_AddData(hEle int, pData *Edit_Data_Copy_, styleTable []uint16, nStyleCount int) int {
	r, _, _ := xEdit_AddData.Call(uintptr(hEle), uintptr(unsafe.Pointer(pData)), uintptr(unsafe.Pointer(&styleTable[0])), uintptr(nStyleCount))
	return int(r)
}

// 编辑框_释放数据.
//
// pData: 数据结构.
func XEdit_FreeData(pData *Edit_Data_Copy_) int {
	r, _, _ := xEdit_FreeData.Call(uintptr(unsafe.Pointer(pData)))
	return int(r)
}

// 编辑框_置默认文本, 当内容为空时, 显示默认文本.
//
// hEle: 元素句柄.
//
// pString: 文本内容.
func XEdit_SetDefaultText(hEle int, pString string) int {
	r, _, _ := xEdit_SetDefaultText.Call(uintptr(hEle), common.StrPtr(pString))
	return int(r)
}

// 编辑框_置默认文本颜色.
//
// hEle: 元素句柄.
//
// color: ABGR 颜色值.
func XEdit_SetDefaultTextColor(hEle int, color int) int {
	r, _, _ := xEdit_SetDefaultTextColor.Call(uintptr(hEle), uintptr(color))
	return int(r)
}

// 编辑框_置密码字符.
//
// hEle: 元素句柄.
//
// ch: 字符.
func XEdit_SetPasswordCharacter(hEle int, ch int) int {
	r, _, _ := xEdit_SetPasswordCharacter.Call(uintptr(hEle), uintptr(ch))
	return int(r)
}

// 编辑框_置文本对齐, 单行模式下有效.
//
// hEle: 元素句柄.
//
// align: 对齐方式, Edit_TextAlign_Flag_.
func XEdit_SetTextAlign(hEle int, align xcc.Edit_TextAlign_Flag_) int {
	r, _, _ := xEdit_SetTextAlign.Call(uintptr(hEle), uintptr(align))
	return int(r)
}

// 编辑框_置TAB空格.
//
// hEle: 元素句柄.
//
// nSpace: 空格数量.
func XEdit_SetTabSpace(hEle int, nSpace int) int {
	r, _, _ := xEdit_SetTabSpace.Call(uintptr(hEle), uintptr(nSpace))
	return int(r)
}

// 编辑框_置文本.
//
// hEle: 元素句柄.
//
// pString: 字符串.
func XEdit_SetText(hEle int, pString string) int {
	r, _, _ := xEdit_SetText.Call(uintptr(hEle), common.StrPtr(pString))
	return int(r)
}

// 编辑框_置文本整数.
//
// hEle: 元素句柄.
//
// nValue: 整数值.
func XEdit_SetTextInt(hEle int, nValue int) int {
	r, _, _ := xEdit_SetTextInt.Call(uintptr(hEle), uintptr(nValue))
	return int(r)
}

// 编辑框_取文本, 不包含非文本内容, 返回实际接收文本长度.
//
// hEle: 元素句柄.
//
// pOut: 接收文本内存指针.
//
// nOutlen: 内存大小. 例: xc.XEdit_GetLength()+1 .
func XEdit_GetText(hEle int, pOut *string, nOutlen int) int {
	buf := make([]uint16, nOutlen)
	r, _, _ := xEdit_GetText.Call(uintptr(hEle), common.Uint16SliceDataPtr(&buf), uintptr(nOutlen))
	*pOut = syscall.UTF16ToString(buf[0:])
	return int(r)
}

// 编辑框_取文本行.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
//
// pOut: 接收文本内存指针.
//
// nOutlen: 接收文本内存块长度. 例: xc.XEdit_GetLengthRow()+1 .
func XEdit_GetTextRow(hEle int, iRow int, pOut *string, nOutlen int) int {
	buf := make([]uint16, nOutlen)
	r, _, _ := xEdit_GetTextRow.Call(uintptr(hEle), uintptr(iRow), common.Uint16SliceDataPtr(&buf), uintptr(nOutlen))
	*pOut = syscall.UTF16ToString(buf[0:])
	return int(r)
}

// 编辑框_取内容长度, 包含非文本内容.
//
// hEle: 元素句柄.
func XEdit_GetLength(hEle int) int {
	r, _, _ := xEdit_GetLength.Call(uintptr(hEle))
	return int(r)
}

// 编辑框_取内容长度行, 包含非文本内容.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
func XEdit_GetLengthRow(hEle int, iRow int) int {
	r, _, _ := xEdit_GetLengthRow.Call(uintptr(hEle), uintptr(iRow))
	return int(r)
}

// 编辑框_取字符, 返回指定位置字符.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
func XEdit_GetAt(hEle int, iRow int, iCol int) int {
	r, _, _ := xEdit_GetAt.Call(uintptr(hEle), uintptr(iRow), uintptr(iCol))
	return int(r)
}

// 编辑框_插入文本.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// pString: 字符串.
func XEdit_InsertText(hEle int, iRow int, iCol int, pString string) int {
	r, _, _ := xEdit_InsertText.Call(uintptr(hEle), uintptr(iRow), uintptr(iCol), common.StrPtr(pString))
	return int(r)
}

// XEdit_AddTextUser 编辑框_插入文本模拟用户操作, 自动刷新UI, 支持撤销/恢复.
//
//	@param hEle 元素句柄.
//	@param pString 字符串.
//	@return int
func XEdit_AddTextUser(hEle int, pString string) int {
	r, _, _ := xEdit_AddTextUser.Call(uintptr(hEle), common.StrPtr(pString))
	return int(r)
}

// 编辑框_添加文本.
//
// hEle: 元素句柄.
//
// pString: 字符串.
func XEdit_AddText(hEle int, pString string) int {
	r, _, _ := xEdit_AddText.Call(uintptr(hEle), common.StrPtr(pString))
	return int(r)
}

// 编辑框_添加文本扩展.
//
// hEle: 元素句柄.
//
// pString: 字符串.
//
// iStyle: 样式索引.
func XEdit_AddTextEx(hEle int, pString string, iStyle int) int {
	r, _, _ := xEdit_AddTextEx.Call(uintptr(hEle), common.StrPtr(pString), uintptr(iStyle))
	return int(r)
}

// 编辑框_添加对象, 例如: 字体, 图片, UI对象, 返回样式索引.
//
// hEle: 元素句柄.
//
// hObj: 对象句柄.
func XEdit_AddObject(hEle int, hObj int) int {
	r, _, _ := xEdit_AddObject.Call(uintptr(hEle), uintptr(hObj))
	return int(r)
}

// 编辑框_添加对象从样式, 当样式为图片时有效.
//
// hEle: 元素句柄.
//
// iStyle: 样式索引.
func XEdit_AddByStyle(hEle int, iStyle int) int {
	r, _, _ := xEdit_AddByStyle.Call(uintptr(hEle), uintptr(iStyle))
	return int(r)
}

// 编辑框_添加样式, 返回样式索引.
//
// hEle: 元素句柄.
//
// hFont_image_Obj: 字体.
//
// color: 颜色.
//
// bColor: 是否使用颜色.
func XEdit_AddStyle(hEle int, hFont_image_Obj int, color int, bColor bool) int {
	r, _, _ := xEdit_AddStyle.Call(uintptr(hEle), uintptr(hFont_image_Obj), uintptr(color), common.BoolPtr(bColor))
	return int(r)
}

// 编辑框_添加样式扩展, 返回样式索引.
//
// hEle: 元素句柄.
//
// fontName: 字体名称.
//
// fontSize: 字体大小.
//
// fontStyle: 字体样式, FontStyle_.
//
// color: 颜色.
//
// bColor: 是否使用颜色.
func XEdit_AddStyleEx(hEle int, fontName string, fontSize int, fontStyle xcc.FontStyle_, color int, bColor bool) int {
	r, _, _ := xEdit_AddStyleEx.Call(uintptr(hEle), common.StrPtr(fontName), uintptr(fontSize), uintptr(fontStyle), uintptr(color), common.BoolPtr(bColor))
	return int(r)
}

// 编辑框_取样式信息.
//
// hEle: 元素句柄.
//
// iStyle: 样式索引.
//
// info: 返回样式信息.
func XEdit_GetStyleInfo(hEle int, iStyle int, info *Edit_Style_Info_) bool {
	r, _, _ := xEdit_GetStyleInfo.Call(uintptr(hEle), uintptr(iStyle), uintptr(unsafe.Pointer(info)))
	return r != 0
}

// 编辑框_置当前样式.
//
// hEle: 元素句柄.
//
// iStyle: 样式索引.
func XEdit_SetCurStyle(hEle int, iStyle int) int {
	r, _, _ := xEdit_SetCurStyle.Call(uintptr(hEle), uintptr(iStyle))
	return int(r)
}

// 编辑框_置插入符颜色.
//
// hEle: 元素句柄.
//
// color: 颜色.
func XEdit_SetCaretColor(hEle int, color int) int {
	r, _, _ := xEdit_SetCaretColor.Call(uintptr(hEle), uintptr(color))
	return int(r)
}

// 编辑框_置插入符宽度.
//
// hEle: 元素句柄.
//
// nWidth: 宽度.
func XEdit_SetCaretWidth(hEle int, nWidth int) int {
	r, _, _ := xEdit_SetCaretWidth.Call(uintptr(hEle), uintptr(nWidth))
	return int(r)
}

// 编辑框_置选择背景颜色.
//
// hEle: 元素句柄.
//
// color: ABGR 颜色.
func XEdit_SetSelectBkColor(hEle int, color int) int {
	r, _, _ := xEdit_SetSelectBkColor.Call(uintptr(hEle), uintptr(color))
	return int(r)
}

// 编辑框_置默认行高.
//
// hEle: 元素句柄.
//
// nHeight: 行高.
func XEdit_SetRowHeight(hEle int, nHeight int) int {
	r, _, _ := xEdit_SetRowHeight.Call(uintptr(hEle), uintptr(nHeight))
	return int(r)
}

// 编辑框_置指定行高度, 类型为 Edit_Type_Richedit 支持指定不同行高.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
//
// nHeight: 高度.
func XEdit_SetRowHeightEx(hEle int, iRow int, nHeight int) int {
	r, _, _ := xEdit_SetRowHeightEx.Call(uintptr(hEle), uintptr(iRow), uintptr(nHeight))
	return int(r)
}

// XEdit_SetCurPos 编辑框_置当前位置.
//
//	@param hEle: 元素句柄.
//	@param iRow: 行索引.
//	@return int
func XEdit_SetCurPos(hEle int, iRow int) int {
	r, _, _ := xEdit_SetCurPos.Call(uintptr(hEle), uintptr(iRow))
	return int(r)
}

// 编辑框_取当前位置点, 返回范围位置点.
//
// hEle: 元素句柄.
func XEdit_GetCurPos(hEle int) int {
	r, _, _ := xEdit_GetCurPos.Call(uintptr(hEle))
	return int(r)
}

// 编辑框_取当前行, 返回行索引.
//
// hEle: 元素句柄.
func XEdit_GetCurRow(hEle int) int {
	r, _, _ := xEdit_GetCurRow.Call(uintptr(hEle))
	return int(r)
}

// 编辑框_取当前列, 返回列索引.
//
// hEle: 元素句柄.
func XEdit_GetCurCol(hEle int) int {
	r, _, _ := xEdit_GetCurCol.Call(uintptr(hEle))
	return int(r)
}

// 编辑框_取坐标点.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// pOut: 接收返回坐标点.
func XEdit_GetPoint(hEle int, iRow int, iCol int, pOut *POINT) int {
	r, _, _ := xEdit_GetPoint.Call(uintptr(hEle), uintptr(iRow), uintptr(iCol), uintptr(unsafe.Pointer(pOut)))
	return int(r)
}

// 编辑框_自动滚动, 视图自动滚动到当前插入符位置.
//
// hEle: 元素句柄.
func XEdit_AutoScroll(hEle int) bool {
	r, _, _ := xEdit_AutoScroll.Call(uintptr(hEle))
	return r != 0
}

// 编辑框_自动滚动扩展, 视图自动滚动到指定位置.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
func XEdit_AutoScrollEx(hEle int, iRow int, iCol int) bool {
	r, _, _ := xEdit_AutoScrollEx.Call(uintptr(hEle), uintptr(iRow), uintptr(iCol))
	return r != 0
}

// XEdit_PosToRowCol 编辑框_转换位置, 转换位置点到行列.
//
//	@param hEle 元素句柄.
//	@param iPos 位置点.
//	@param pInfo 行列.
//	@return int
func XEdit_PosToRowCol(hEle int, iPos int, pInfo *Position_) int {
	r, _, _ := xEdit_PosToRowCol.Call(uintptr(hEle), uintptr(iPos), uintptr(unsafe.Pointer(pInfo)))
	return int(r)
}

// 编辑框_选择全部.
//
// hEle: 元素句柄.
func XEdit_SelectAll(hEle int) bool {
	r, _, _ := xEdit_SelectAll.Call(uintptr(hEle))
	return r != 0
}

// 编辑框_取消选择.
//
// hEle: 元素句柄.
func XEdit_CancelSelect(hEle int) bool {
	r, _, _ := xEdit_CancelSelect.Call(uintptr(hEle))
	return r != 0
}

// 编辑框_删除选择内容.
//
// hEle: 元素句柄.
func XEdit_DeleteSelect(hEle int) bool {
	r, _, _ := xEdit_DeleteSelect.Call(uintptr(hEle))
	return r != 0
}

// 编辑框_置选择.
//
// hEle: 元素句柄.
//
// iStartRow: 起始行索引.
//
// iStartCol: 起始行列索引.
//
// iEndRow: 结束行索引.
//
// iEndCol: 结束行列索引.
func XEdit_SetSelect(hEle int, iStartRow int, iStartCol int, iEndRow int, iEndCol int) bool {
	r, _, _ := xEdit_SetSelect.Call(uintptr(hEle), uintptr(iStartRow), uintptr(iStartCol), uintptr(iEndRow), uintptr(iEndCol))
	return r != 0
}

// 编辑框_取选择文本, 不包括非文本内容, 返回接收文本内容实际长度.
//
// hEle: 元素句柄.
//
// pOut: 接收返回文本内容.
//
// nOutLen: 接收内存大小. xc.XEdit_GetSelectTextLength()+1 .
func XEdit_GetSelectText(hEle int, pOut *string, nOutLen int) int {
	buf := make([]uint16, nOutLen)
	r, _, _ := xEdit_GetSelectText.Call(uintptr(hEle), common.Uint16SliceDataPtr(&buf), uintptr(nOutLen))
	*pOut = syscall.UTF16ToString(buf[0:])
	return int(r)
}

// 编辑框_取选择内容范围.
//
// hEle: 元素句柄.
//
// pBegin: 起始位置.
//
// pEnd: 结束位置.
func XEdit_GetSelectRange(hEle int, pBegin *Position_, pEnd *Position_) bool {
	r, _, _ := xEdit_GetSelectRange.Call(uintptr(hEle), uintptr(unsafe.Pointer(pBegin)), uintptr(unsafe.Pointer(pEnd)))
	return r != 0
}

// 编辑框_取可视行范围.
//
// hEle: 元素句柄.
//
// piStart: 起始行索引.
//
// piEnd: 结束行索引.
func XEdit_GetVisibleRowRange(hEle int, piStart *int32, piEnd *int32) int {
	r, _, _ := xEdit_GetVisibleRowRange.Call(uintptr(hEle), uintptr(unsafe.Pointer(piStart)), uintptr(unsafe.Pointer(piEnd)))
	return int(r)
}

// 编辑框_删除, 删除指定范围内容.
//
// hEle: 元素句柄.
//
// iStartRow: 起始行索引.
//
// iStartCol: 起始行列索引.
//
// iEndRow: 结束行索引.
//
// iEndCol: 结束行列索引.
func XEdit_Delete(hEle int, iStartRow int, iStartCol int, iEndRow int, iEndCol int) bool {
	r, _, _ := xEdit_Delete.Call(uintptr(hEle), uintptr(iStartRow), uintptr(iStartCol), uintptr(iEndRow), uintptr(iEndCol))
	return r != 0
}

// 编辑框_删除行.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
func XEdit_DeleteRow(hEle int, iRow int) bool {
	r, _, _ := xEdit_DeleteRow.Call(uintptr(hEle), uintptr(iRow))
	return r != 0
}

// 编辑框_剪贴板剪切.
//
// hEle: 元素句柄.
func XEdit_ClipboardCut(hEle int) bool {
	r, _, _ := xEdit_ClipboardCut.Call(uintptr(hEle))
	return r != 0
}

// 编辑框_剪贴板复制.
//
// hEle: 元素句柄.
func XEdit_ClipboardCopy(hEle int) bool {
	r, _, _ := xEdit_ClipboardCopy.Call(uintptr(hEle))
	return r != 0
}

// 编辑框_剪贴板粘贴.
//
// hEle: 元素句柄.
func XEdit_ClipboardPaste(hEle int) bool {
	r, _, _ := xEdit_ClipboardPaste.Call(uintptr(hEle))
	return r != 0
}

// 编辑框_撤销.
//
// hEle: 元素句柄.
func XEdit_Undo(hEle int) bool {
	r, _, _ := xEdit_Undo.Call(uintptr(hEle))
	return r != 0
}

// 编辑框_恢复.
//
// hEle: 元素句柄.
func XEdit_Redo(hEle int) bool {
	r, _, _ := xEdit_Redo.Call(uintptr(hEle))
	return r != 0
}

// 编辑框_添加气泡开始, 当前行开始.
//
// hEle: 元素句柄.
//
// hImageAvatar: 头像.
//
// hImageBubble: 气泡背景.
//
// nFlag: 标志, Chat_Flag_.
func XEdit_AddChatBegin(hEle int, hImageAvatar int, hImageBubble int, nFlag xcc.Chat_Flag_) int {
	r, _, _ := xEdit_AddChatBegin.Call(uintptr(hEle), uintptr(hImageAvatar), uintptr(hImageBubble), uintptr(nFlag))
	return int(r)
}

// 编辑框_添加气泡结束, 当前行结束.
//
// hEle: 元素句柄.
func XEdit_AddChatEnd(hEle int) int {
	r, _, _ := xEdit_AddChatEnd.Call(uintptr(hEle))
	return int(r)
}

// 编辑框_置气泡缩进, 设置聊天气泡内容缩进.
//
// hEle: 元素句柄.
//
// nIndentation: 缩进值.
func XEdit_SetChatIndentation(hEle int, nIndentation int) int {
	r, _, _ := xEdit_SetChatIndentation.Call(uintptr(hEle), uintptr(nIndentation))
	return int(r)
}

// 编辑框_置行间隔, 设置行间隔大小, 多行模式有效.
//
// hEle: 元素句柄.
//
// nSpace: 行间隔大小.
func XEdit_SetRowSpace(hEle int, nSpace int) int {
	r, _, _ := xEdit_SetRowSpace.Call(uintptr(hEle), uintptr(nSpace))
	return int(r)
}

// 编辑框_置当前位置扩展.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
func XEdit_SetCurPosEx(hEle int, iRow, iCol int32) int {
	r, _, _ := xEdit_SetCurPosEx.Call(uintptr(hEle), uintptr(iRow), uintptr(iCol))
	return int(r)
}

// 编辑框_取当前位置扩展.
//
// hEle: 元素句柄.
//
// iRow: 返回行索引.
//
// iCol: 返回列索引.
func XEdit_GetCurPosEx(hEle int, iRow, iCol *int32) int {
	r, _, _ := xEdit_GetCurPosEx.Call(uintptr(hEle), uintptr(unsafe.Pointer(iRow)), uintptr(unsafe.Pointer(iCol)))
	return int(r)
}

// 编辑框_移动到末尾.
//
// hEle: 元素句柄.
func XEdit_MoveEnd(hEle int) int {
	r, _, _ := xEdit_MoveEnd.Call(uintptr(hEle))
	return int(r)
}

// 编辑框_行列到位置, 返回位置点.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
func XEdit_RowColToPos(hEle int, iRow int, iCol int) int {
	r, _, _ := xEdit_RowColToPos.Call(uintptr(hEle), uintptr(iRow), uintptr(iCol))
	return int(r)
}

// 编辑框_置后备字体, 置中文字体. 如果已设置, 当遇到中文字符时使用后备字体, 解决不支持中文的字体的问题
//
// hEle: 元素句柄.
//
// hFont: 字体.
func XEdit_SetBackFont(hEle int, hFont int) int {
	r, _, _ := xEdit_SetBackFont.Call(uintptr(hEle), uintptr(hFont))
	return int(r)
}

// 编辑框_释放样式.
//
// hEle: 元素句柄.
//
// iStyle: 样式.
func XEdit_ReleaseStyle(hEle int, iStyle int) bool {
	r, _, _ := xEdit_ReleaseStyle.Call(uintptr(hEle), uintptr(iStyle))
	return r != 0
}

// 编辑框_修改样式.
//
// hEle: 元素句柄.
//
// iStyle: 样式索引.
//
// hFont: 字体句柄.
//
// color: ABGR 颜色.
//
// bColor: 是否使用颜色.
func XEdit_ModifyStyle(hEle int, iStyle int, hFont int, color int, bColor bool) bool {
	r, _, _ := xEdit_ModifyStyle.Call(uintptr(hEle), uintptr(iStyle), uintptr(hFont), uintptr(color), common.BoolPtr(bColor))
	return r != 0
}

// 编辑框_置空格大小.
//
// hEle: 元素句柄.
//
// size: 空格大小.
func XEdit_SetSpaceSize(hEle int, size int) int {
	r, _, _ := xEdit_SetSpaceSize.Call(uintptr(hEle), uintptr(size))
	return int(r)
}

// 编辑框_置字符间距.
//
// hEle: 元素句柄.
//
// size: 英文字符间距大小.
//
// sizeZh: 中文字符间距大小.
func XEdit_SetCharSpaceSize(hEle int, size int, sizeZh int) int {
	r, _, _ := xEdit_SetCharSpaceSize.Call(uintptr(hEle), uintptr(size), uintptr(sizeZh))
	return int(r)
}

// 编辑框_取选择文本长度, 不包括非文本内容, 返回文本内容长度.
//
// hEle: 元素句柄.
func XEdit_GetSelectTextLength(hEle int) int {
	r, _, _ := xEdit_GetSelectTextLength.Call(uintptr(hEle))
	return int(r)
}

// 编辑框_置选择文本样式.
//
// hEle: 元素句柄.
//
// iStyle: 样式索引.
func XEdit_SetSelectTextStyle(hEle int, iStyle int) int {
	r, _, _ := xEdit_SetSelectTextStyle.Call(uintptr(hEle), uintptr(iStyle))
	return int(r)
}

// 编辑框_取文本_临时, 不包含非文本内容. 返回临时文本, 临时缓存区大小: xcc.Text_Buffer_Size .
//
// hEle: 元素句柄.
func XEdit_GetText_Temp(hEle int) string {
	r, _, _ := xEdit_GetText_Temp.Call(uintptr(hEle))
	return common.UintPtrToString(r)
}

// 编辑框_取文本行_临时, 获取指定行文本内容. 返回临时文本, 临时缓存区大小: xcc.Text_Buffer_Size .
//
// hEle: 元素句柄.
//
// iRow: 行索引.
func XEdit_GetTextRow_Temp(hEle int, iRow int) string {
	r, _, _ := xEdit_GetTextRow_Temp.Call(uintptr(hEle), uintptr(iRow))
	return common.UintPtrToString(r)
}

// 编辑框_取选择文本, 不包含非文本内容. 返回临时文本, 临时缓存区大小: xcc.Text_Buffer_Size .
//
// hEle: 元素句柄.
func XEdit_GetSelectText_Temp(hEle int) string {
	r, _, _ := xEdit_GetSelectText_Temp.Call(uintptr(hEle))
	return common.UintPtrToString(r)
}

// 编辑框_插入气泡开始, 当前行开始.
//
// hEle: 元素句柄.
//
// hImageAvatar: 头像图片句柄.
//
// hImageBubble: 气泡背景图片句柄.
//
// nFlag: 聊天气泡对齐方式: xcc.Chat_Flag_ .
func XEdit_InsertChatBegin(hEle int, hImageAvatar int, hImageBubble int, nFlag xcc.Chat_Flag_) int {
	r, _, _ := xEdit_InsertChatBegin.Call(uintptr(hEle), uintptr(hImageAvatar), uintptr(hImageBubble), uintptr(nFlag))
	return int(r)
}

// 编辑框_取指定行气泡标识. 返回行标识: xcc.Chat_Flag_
//
// hEle: 元素句柄.
//
// iRow: 行索引.
func XEdit_GetChatFlags(hEle int, iRow int) xcc.Chat_Flag_ {
	r, _, _ := xEdit_GetChatFlags.Call(uintptr(hEle), uintptr(iRow))
	return xcc.Chat_Flag_(r)
}

// 编辑框_插入文本扩展.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// pString: 字符串.
//
// iStyle: 样式.
func XEdit_InsertTextEx(hEle int, iRow int, iCol int, pString string, iStyle int) int {
	r, _, _ := xEdit_InsertTextEx.Call(uintptr(hEle), uintptr(iRow), uintptr(iCol), common.StrPtr(pString), uintptr(iStyle))
	return int(r)
}

// 编辑框_插入对象.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// hObj: 对象句柄.
func XEdit_InsertObject(hEle int, iRow int, iCol int, hObj int) int {
	r, _, _ := xEdit_InsertObject.Call(uintptr(hEle), uintptr(iRow), uintptr(iCol), uintptr(hObj))
	return int(r)
}

// 编辑框_置气泡最大宽度. 当值为0时代表不限制宽度.
//
// hEle: 元素句柄.
//
// nWidth: 最大宽度.
func XEdit_SetChatMaxWidth(hEle int, nWidth int32) {
	xEdit_SetChatMaxWidth.Call(uintptr(hEle), uintptr(nWidth))
}

package widget

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// 编辑框(常规, 富文本, 聊天气泡).
type Edit struct {
	ScrollView
}

// 编辑框_创建.
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
func NewEdit(x int, y int, cx int, cy int, hParent int) *Edit {
	p := &Edit{}
	p.SetHandle(xc.XEdit_Create(x, y, cx, cy, hParent))
	return p
}

// 编辑框_创建扩展.
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
func NewEditEx(x int, y int, cx int, cy int, nType int, hParent int) *Edit {
	p := &Edit{}
	p.Handle = xc.XEdit_CreateEx(x, y, cx, cy, nType, hParent)
	return p
}

// 从句柄创建对象.
func NewEditByHandle(handle int) *Edit {
	p := &Edit{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func NewEditByName(name string) *Edit {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &Edit{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func NewEditByUID(nUID int) *Edit {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &Edit{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func NewEditByUIDName(name string) *Edit {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &Edit{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 编辑框_启用自动换行.
//
// bEnable: 是否启用.
func (e *Edit) EnableAutoWrap(bEnable bool) int {
	return xc.XEdit_EnableAutoWrap(e.Handle, bEnable)
}

// 编辑框_启用只读.
//
// bEnable: 是否启用.
func (e *Edit) EnableReadOnly(bEnable bool) int {
	return xc.XEdit_EnableReadOnly(e.Handle, bEnable)
}

// 编辑框_启用多行.
//
// bEnable:.
func (e *Edit) EnableMultiLine(bEnable bool) int {
	return xc.XEdit_EnableMultiLine(e.Handle, bEnable)
}

// 编辑框_启用密码, 启用密码模式(只支持默认类型编辑框).
//
// bEnable: 是否启用.
func (e *Edit) EnablePassword(bEnable bool) int {
	return xc.XEdit_EnablePassword(e.Handle, bEnable)
}

// 编辑框_启用自动选择, 当获得焦点时,自动选择所有内容.
//
// bEnable: 是否启用.
func (e *Edit) EnableAutoSelAll(bEnable bool) int {
	return xc.XEdit_EnableAutoSelAll(e.Handle, bEnable)
}

// 编辑框_启用自动取消选择, 当失去焦点时自动取消选择.
//
// bEnable: 是否启用.
func (e *Edit) EnableAutoCancelSel(bEnable bool) int {
	return xc.XEdit_EnableAutoCancelSel(e.Handle, bEnable)
}

// 编辑框_是否只读.
func (e *Edit) IsReadOnly() bool {
	return xc.XEdit_IsReadOnly(e.Handle)
}

// 编辑框_是否多行.
func (e *Edit) IsMultiLine() bool {
	return xc.XEdit_IsMultiLine(e.Handle)
}

// 编辑框_是否密码.
func (e *Edit) IsPassword() bool {
	return xc.XEdit_IsPassword(e.Handle)
}

// 编辑框_是否自动换行.
func (e *Edit) IsAutoWrap() bool {
	return xc.XEdit_IsAutoWrap(e.Handle)
}

// 编辑框_判断为空.
func (e *Edit) IsEmpty() bool {
	return xc.XEdit_IsEmpty(e.Handle)
}

// 编辑框_是否在选择区域.
//
// iRow: 行索引.
//
// iCol: 列索引.
func (e *Edit) IsInSelect(iRow int, iCol int) bool {
	return xc.XEdit_IsInSelect(e.Handle, iRow, iCol)
}

// 编辑框_取总行数.
func (e *Edit) GetRowCount() int {
	return xc.XEdit_GetRowCount(e.Handle)
}

// 编辑框_取数据, 包含文本或非文本内容.
func (e *Edit) GetData() xc.Edit_Data_Copy_ {
	return xc.XEdit_GetData(e.Handle)
}

// 编辑框_添加数据.
//
// pData: 数据结构.
//
// styleTable: 样式表.
//
// nStyleCount: 样式数量.
func (e *Edit) AddData(pData *xc.Edit_Data_Copy_, styleTable []uint16, nStyleCount int) int {
	return xc.XEdit_AddData(e.Handle, pData, styleTable, nStyleCount)
}

// 编辑框_释放数据.
func (e *Edit) FreeData(pData *xc.Edit_Data_Copy_) int {
	return xc.XEdit_FreeData(pData)
}

// 编辑框_置默认文本, 当内容为空时, 显示默认文本.
//
// pString: 文本内容.
func (e *Edit) SetDefaultText(pString string) int {
	return xc.XEdit_SetDefaultText(e.Handle, pString)
}

// 编辑框_置默认文本颜色.
//
// color: ABGR颜色值.
func (e *Edit) SetDefaultTextColor(color int) int {
	return xc.XEdit_SetDefaultTextColor(e.Handle, color)
}

// 编辑框_置密码字符.
//
// ch: 字符.
func (e *Edit) SetPasswordCharacter(ch int) int {
	return xc.XEdit_SetPasswordCharacter(e.Handle, ch)
}

// 编辑框_置文本对齐, 单行模式下有效.
//
// align: 对齐方式, Edit_TextAlign_Flag_.
func (e *Edit) SetTextAlign(align int) int {
	return xc.XEdit_SetTextAlign(e.Handle, align)
}

// 编辑框_置TAB空格.
//
// nSpace: 空格数量.
func (e *Edit) SetTabSpace(nSpace int) int {
	return xc.XEdit_SetTabSpace(e.Handle, nSpace)
}

// 编辑框_置文本.
//
// pString: 字符串.
func (e *Edit) SetText(pString string) int {
	return xc.XEdit_SetText(e.Handle, pString)
}

// 编辑框_置文本整数.
//
// nValue: 整数值.
func (e *Edit) SetTextInt(nValue int) int {
	return xc.XEdit_SetTextInt(e.Handle, nValue)
}

// 编辑框_取文本, 不包含非文本内容, 返回实际接收文本长度.
//
// pOut: 接收文本内存指针.
//
// nOutlen: 内存大小.
func (e *Edit) GetText(pOut *string, nOutlen int) int {
	return xc.XEdit_GetText(e.Handle, pOut, nOutlen)
}

// 编辑框_取文本行.
//
// iRow: 行索引.
//
// pOut: 接收文本内存指针.
//
// nOutlen: 接收文本内存块长度.
func (e *Edit) GetTextRow(iRow int, pOut *string, nOutlen int) int {
	return xc.XEdit_GetTextRow(e.Handle, iRow, pOut, nOutlen)
}

// 编辑框_取内容长度, 包含非文本内容.
func (e *Edit) GetLength() int {
	return xc.XEdit_GetLength(e.Handle)
}

// 编辑框_取内容长度行, 包含非文本内容.
//
// iRow: 行索引.
func (e *Edit) GetLengthRow(iRow int) int {
	return xc.XEdit_GetLengthRow(e.Handle, iRow)
}

// 编辑框_取字符, 返回指定位置字符.
//
// iRow: 行索引.
//
// iCol: 列索引.
func (e *Edit) GetAt(iRow int, iCol int) int {
	return xc.XEdit_GetAt(e.Handle, iRow, iCol)
}

// 编辑框_插入文本.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// pString: 字符串.
func (e *Edit) InsertText(iRow int, iCol int, pString string) int {
	return xc.XEdit_InsertText(e.Handle, iRow, iCol, pString)
}

// 编辑框_插入文本模拟用户操作, 自动刷新UI, 支持撤销/恢复.
//
// pString: 字符串.
func (e *Edit) InsertTextUser(pString string) int {
	return xc.XEdit_InsertTextUser(e.Handle, pString)
}

// 编辑框_添加文本.
//
// pString: 字符串.
func (e *Edit) AddText(pString string) int {
	return xc.XEdit_AddText(e.Handle, pString)
}

// 编辑框_添加文本扩展.
//
// pString: 字符串.
//
// iStyle: 样式索引.
func (e *Edit) AddTextEx(pString string, iStyle int) int {
	return xc.XEdit_AddTextEx(e.Handle, pString, iStyle)
}

// 编辑框_添加对象, 例如: 字体, 图片, UI对象, 返回样式索引.
//
// hObj: 对象句柄.
func (e *Edit) AddObject(hObj int) int {
	return xc.XEdit_AddObject(e.Handle, hObj)
}

// 编辑框_添加对象从样式, 当样式为图片时有效.
//
// iStyle: 样式索引.
func (e *Edit) AddByStyle(iStyle int) int {
	return xc.XEdit_AddByStyle(e.Handle, iStyle)
}

// 编辑框_添加样式, 返回样式索引.
//
// hFont_image_Obj: 字体.
//
// color: 颜色.
//
// bColor: 是否使用颜色.
func (e *Edit) AddStyle(hFont_image_Obj int, color int, bColor bool) int {
	return xc.XEdit_AddStyle(e.Handle, hFont_image_Obj, color, bColor)
}

// 编辑框_添加样式扩展, 返回样式索引.
//
// fontName: 字体名称.
//
// fontSize: 字体大小.
//
// fontStyle: 字体样式.
//
// color: 颜色.
//
// bColor: 是否使用颜色.
func (e *Edit) AddStyleEx(fontName string, fontSize int, fontStyle int, color int, bColor bool) int {
	return xc.XEdit_AddStyleEx(e.Handle, fontName, fontSize, fontStyle, color, bColor)
}

// 编辑框_取样式信息.
//
// iStyle: 样式索引.
//
// info: 返回样式信息.
func (e *Edit) GetStyleInfo(iStyle int, info *xc.Edit_Style_Info_) bool {
	return xc.XEdit_GetStyleInfo(e.Handle, iStyle, info)
}

// 编辑框_置当前样式.
//
// iStyle: 样式索引.
func (e *Edit) SetCurStyle(iStyle int) int {
	return xc.XEdit_SetCurStyle(e.Handle, iStyle)
}

// 编辑框_置插入符颜色.
//
// color: 颜色.
func (e *Edit) SetCaretColor(color int) int {
	return xc.XEdit_SetCaretColor(e.Handle, color)
}

// 编辑框_置插入符宽度.
//
// nWidth: 宽度.
func (e *Edit) SetCaretWidth(nWidth int) int {
	return xc.XEdit_SetCaretWidth(e.Handle, nWidth)
}

// 编辑框_置选择背景颜色.
//
// color: ABGR颜色.
func (e *Edit) SetSelectBkColor(color int) int {
	return xc.XEdit_SetSelectBkColor(e.Handle, color)
}

// 编辑框_置默认行高.
//
// nHeight: 行高.
func (e *Edit) SetRowHeight(nHeight int) int {
	return xc.XEdit_SetRowHeight(e.Handle, nHeight)
}

// 编辑框_置指定行高度, 类型为 Edit_Type_Richedit 支持指定不同行高.
//
// iRow: 行索引.
//
// nHeight: 高度.
func (e *Edit) SetRowHeightEx(iRow int, nHeight int) int {
	return xc.XEdit_SetRowHeightEx(e.Handle, iRow, nHeight)
}

// 编辑框_置当前位置.
//
// iRow: 行索引.
//
// iCol: 列索引.
func (e *Edit) SetCurPos(iRow int, iCol int) int {
	return xc.XEdit_SetCurPos(e.Handle, iRow, iCol)
}

// 编辑框_取当前位置点, 返回范围位置点.
func (e *Edit) GetCurPos() int {
	return xc.XEdit_GetCurPos(e.Handle)
}

// 编辑框_取当前行, 返回行索引.
func (e *Edit) GetCurRow() int {
	return xc.XEdit_GetCurRow(e.Handle)
}

// 编辑框_取当前列, 返回列索引.
func (e *Edit) GetCurCol() int {
	return xc.XEdit_GetCurCol(e.Handle)
}

// 编辑框_取坐标点.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// pOut: 接收返回坐标点.
func (e *Edit) GetPoint(iRow int, iCol int, pOut *xc.POINT) int {
	return xc.XEdit_GetPoint(e.Handle, iRow, iCol, pOut)
}

// 编辑框_自动滚动, 视图自动滚动到当前插入符位置.
func (e *Edit) AutoScroll() bool {
	return xc.XEdit_AutoScroll(e.Handle)
}

// 编辑框_自动滚动扩展, 视图自动滚动到指定位置.
//
// iRow: 行索引.
//
// iCol: 列索引.
func (e *Edit) AutoScrollEx(iRow int, iCol int) bool {
	return xc.XEdit_AutoScrollEx(e.Handle, iRow, iCol)
}

// 编辑框_转换位置, 转换位置点到行列.
//
// iPos: 位置点.
//
// pInfo: 行列.
func (e *Edit) PositionToInfo(iPos int, pInfo *xc.Position_) int {
	return xc.XEdit_PositionToInfo(e.Handle, iPos, pInfo)
}

// 编辑框_选择全部.
func (e *Edit) SelectAll() bool {
	return xc.XEdit_SelectAll(e.Handle)
}

// 编辑框_取消选择.
func (e *Edit) CancelSelect() bool {
	return xc.XEdit_CancelSelect(e.Handle)
}

// 编辑框_删除选择内容.
func (e *Edit) DeleteSelect() bool {
	return xc.XEdit_DeleteSelect(e.Handle)
}

// 编辑框_置选择.
//
// iStartRow: 起始行索引.
//
// iStartCol: 起始行列索引.
//
// iEndRow: 结束行索引.
//
// iEndCol: 结束行列索引.
func (e *Edit) SetSelect(iStartRow int, iStartCol int, iEndRow int, iEndCol int) bool {
	return xc.XEdit_SetSelect(e.Handle, iStartRow, iStartCol, iEndRow, iEndCol)
}

// 编辑框_取选择文本, 不包括非文本内容, 返回接收文本内容实际长度.
//
// pOut: 接收返回文本内容.
//
// nOutLen: 接收内存大小.
func (e *Edit) GetSelectText(pOut *string, nOutLen int) int {
	return xc.XEdit_GetSelectText(e.Handle, pOut, nOutLen)
}

// 编辑框_取选择内容范围.
//
// pBegin: 起始位置.
//
// pEnd: 结束位置.
func (e *Edit) GetSelectRange(pBegin *xc.Position_, pEnd *xc.Position_) bool {
	return xc.XEdit_GetSelectRange(e.Handle, pBegin, pEnd)
}

// 编辑框_取可视行范围.
//
// piStart: 起始行索引.
//
// piEnd: 结束行索引.
func (e *Edit) GetVisibleRowRange(piStart *int, piEnd *int) int {
	return xc.XEdit_GetVisibleRowRange(e.Handle, piStart, piEnd)
}

// 编辑框_删除, 删除指定范围内容.
//
// iStartRow: 起始行索引.
//
// iStartCol: 起始行列索引.
//
// iEndRow: 结束行索引.
//
// iEndCol: 结束行列索引.
func (e *Edit) Delete(iStartRow int, iStartCol int, iEndRow int, iEndCol int) bool {
	return xc.XEdit_Delete(e.Handle, iStartRow, iStartCol, iEndRow, iEndCol)
}

// 编辑框_删除行.
//
// iRow: 行索引.
func (e *Edit) DeleteRow(iRow int) bool {
	return xc.XEdit_DeleteRow(e.Handle, iRow)
}

// 编辑框_剪贴板剪切.
func (e *Edit) ClipboardCut() bool {
	return xc.XEdit_ClipboardCut(e.Handle)
}

// 编辑框_剪贴板复制.
func (e *Edit) ClipboardCopy() bool {
	return xc.XEdit_ClipboardCopy(e.Handle)
}

// 编辑框_剪贴板粘贴.
func (e *Edit) ClipboardPaste() bool {
	return xc.XEdit_ClipboardPaste(e.Handle)
}

// 编辑框_撤销.
func (e *Edit) Undo() bool {
	return xc.XEdit_Undo(e.Handle)
}

// 编辑框_恢复.
func (e *Edit) Redo() bool {
	return xc.XEdit_Redo(e.Handle)
}

// 编辑框_添加气泡开始, 当前行开始.
//
// hImageAvatar: 头像.
//
// hImageBubble: 气泡背景.
//
// nFlag: 标志, Chat_Flag_.
func (e *Edit) AddChatBegin(hImageAvatar int, hImageBubble int, nFlag int) int {
	return xc.XEdit_AddChatBegin(e.Handle, hImageAvatar, hImageBubble, nFlag)
}

// 编辑框_添加气泡结束, 当前行结束.
func (e *Edit) AddChatEnd() int {
	return xc.XEdit_AddChatEnd(e.Handle)
}

// 编辑框_置气泡缩进, 设置聊天气泡内容缩进.
//
// nIndentation: 缩进值.
func (e *Edit) SetChatIndentation(nIndentation int) int {
	return xc.XEdit_SetChatIndentation(e.Handle, nIndentation)
}

/*
以下都是事件
*/

type XE_EDIT_SET func(pbHandled *bool) int                                               // 编辑框_置文本.
type XE_EDIT_SET1 func(hEle int, pbHandled *bool) int                                    // 编辑框_置文本.
type XE_EDIT_DRAWROW func(hDraw int, iRow int, pbHandled *bool) int                      // 和XE_EDIT_CHANGED的对换.
type XE_EDIT_DRAWROW1 func(hEle int, hDraw int, iRow int, pbHandled *bool) int           // 和XE_EDIT_CHANGED的对换.
type XE_EDIT_CHANGED func(pbHandled *bool) int                                           // 编辑框_内容被改变.
type XE_EDIT_CHANGED1 func(hEle int, pbHandled *bool) int                                // 编辑框_内容被改变.
type XE_EDIT_POS_CHANGED func(iPos int, pbHandled *bool) int                             // 编辑框_光标位置_被改变.
type XE_EDIT_POS_CHANGED1 func(hEle int, iPos int, pbHandled *bool) int                  // 编辑框_光标位置_被改变.
type XE_EDIT_STYLE_CHANGED func(iStyle int, pbHandled *bool) int                         // 编辑框_样式_被改变.
type XE_EDIT_STYLE_CHANGED1 func(hEle int, iStyle int, pbHandled *bool) int              // 编辑框_样式_被改变.
type XE_EDIT_ENTER_GET_TABALIGN func(pbHandled *bool) int                                // 编辑框_回车_获取标签?.
type XE_EDIT_ENTER_GET_TABALIGN1 func(hEle int, pbHandled *bool) int                     // 编辑框_回车_获取标签?.
type XE_EDIT_ROW_CHANGED func(iRow int, nChangeRows int, pbHandled *bool) int            // 编辑框_行_被改变.
type XE_EDIT_ROW_CHANGED1 func(hEle int, iRow int, nChangeRows int, pbHandled *bool) int // 编辑框_行_被改变.

// 编辑框_置文本.
func (e *Edit) Event_EDIT_SET(pFun XE_EDIT_SET) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_EDIT_SET, pFun)
}

// 编辑框_置文本.
func (e *Edit) Event_EDIT_SET1(pFun XE_EDIT_SET1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_EDIT_SET, pFun)
}

// 和XE_EDIT_CHANGED的对换.
func (e *Edit) Event_EDIT_DRAWROW(pFun XE_EDIT_DRAWROW) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_EDIT_DRAWROW, pFun)
}

// 和XE_EDIT_CHANGED的对换.
func (e *Edit) Event_EDIT_DRAWROW1(pFun XE_EDIT_DRAWROW1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_EDIT_DRAWROW, pFun)
}

// 编辑框_内容被改变.
func (e *Edit) Event_EDIT_CHANGED(pFun XE_EDIT_CHANGED) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_EDIT_CHANGED, pFun)
}

// 编辑框_内容被改变.
func (e *Edit) Event_EDIT_CHANGED1(pFun XE_EDIT_CHANGED1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_EDIT_CHANGED, pFun)
}

// 编辑框_光标位置_被改变.
func (e *Edit) Event_EDIT_POS_CHANGED(pFun XE_EDIT_POS_CHANGED) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_EDIT_POS_CHANGED, pFun)
}

// 编辑框_光标位置_被改变.
func (e *Edit) Event_EDIT_POS_CHANGED1(pFun XE_EDIT_POS_CHANGED1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_EDIT_POS_CHANGED, pFun)
}

// 编辑框_样式_被改变.
func (e *Edit) Event_EDIT_STYLE_CHANGED(pFun XE_EDIT_STYLE_CHANGED) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_EDIT_STYLE_CHANGED, pFun)
}

// 编辑框_样式_被改变.
func (e *Edit) Event_EDIT_STYLE_CHANGED1(pFun XE_EDIT_STYLE_CHANGED1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_EDIT_STYLE_CHANGED, pFun)
}

// 编辑框_回车_获取标签?.
func (e *Edit) Event_EDIT_ENTER_GET_TABALIGN(pFun XE_EDIT_ENTER_GET_TABALIGN) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_EDIT_ENTER_GET_TABALIGN, pFun)
}

// 编辑框_回车_获取标签?.
func (e *Edit) Event_EDIT_ENTER_GET_TABALIGN1(pFun XE_EDIT_ENTER_GET_TABALIGN1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_EDIT_ENTER_GET_TABALIGN, pFun)
}

// 编辑框_行_被改变.
func (e *Edit) Event_EDIT_ROW_CHANGED(pFun XE_EDIT_ROW_CHANGED) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_EDIT_ROW_CHANGED, pFun)
}

// 编辑框_行_被改变.
func (e *Edit) Event_EDIT_ROW_CHANGED1(pFun XE_EDIT_ROW_CHANGED1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_EDIT_ROW_CHANGED, pFun)
}

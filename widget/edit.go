package widget

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// 编辑框(常规, 富文本, 聊天气泡).
type Edit struct {
	ScrollView
}

// 编辑框_创建, 失败返回 nil.
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
func NewEdit(x, y, cx, cy int32, hParent int) *Edit {
	return NewEditByHandle(xc.XEdit_Create(x, y, cx, cy, hParent))
}

// 编辑框_创建扩展, 失败返回 nil.
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
func NewEditEx(x, y, cx, cy int32, nType xcc.Edit_Type_, hParent int) *Edit {
	return NewEditByHandle(xc.XEdit_CreateEx(x, y, cx, cy, nType, hParent))
}

// 从句柄创建对象, 失败返回 nil.
func NewEditByHandle(handle int) *Edit {
	if handle <= 0 {
		return nil
	}
	p := &Edit{}
	p.SetHandle(handle)
	return p
}

// 从 name 创建对象, 失败返回 nil.
func NewEditByName(name string) *Edit {
	return NewEditByHandle(xc.XC_GetObjectByName(name))
}

// 从 UID 创建对象, 失败返回 nil.
func NewEditByUID(nUID int32) *Edit {
	return NewEditByHandle(xc.XC_GetObjectByUID(nUID))
}

// 从 UID 名称创建对象, 失败返回 nil.
func NewEditByUIDName(name string) *Edit {
	return NewEditByHandle(xc.XC_GetObjectByUIDName(name))
}

// 编辑框_启用自动换行.
//
// bEnable: 是否启用.
func (e *Edit) EnableAutoWrap(bEnable bool) *Edit {
	xc.XEdit_EnableAutoWrap(e.Handle, bEnable)
	return e
}

// 编辑框_启用只读.
//
// bEnable: 是否启用.
func (e *Edit) EnableReadOnly(bEnable bool) *Edit {
	xc.XEdit_EnableReadOnly(e.Handle, bEnable)
	return e
}

// 编辑框_启用多行.
//
// bEnable:.
func (e *Edit) EnableMultiLine(bEnable bool) *Edit {
	xc.XEdit_EnableMultiLine(e.Handle, bEnable)
	return e
}

// 编辑框_启用密码, 启用密码模式(只支持默认类型编辑框).
//
// bEnable: 是否启用.
func (e *Edit) EnablePassword(bEnable bool) *Edit {
	xc.XEdit_EnablePassword(e.Handle, bEnable)
	return e
}

// 编辑框_启用自动选择, 当获得焦点时,自动选择所有内容.
//
// bEnable: 是否启用.
func (e *Edit) EnableAutoSelAll(bEnable bool) *Edit {
	xc.XEdit_EnableAutoSelAll(e.Handle, bEnable)
	return e
}

// 编辑框_启用自动取消选择, 当失去焦点时自动取消选择.
//
// bEnable: 是否启用.
func (e *Edit) EnableAutoCancelSel(bEnable bool) *Edit {
	xc.XEdit_EnableAutoCancelSel(e.Handle, bEnable)
	return e
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
func (e *Edit) IsInSelect(iRow, iCol int32) bool {
	return xc.XEdit_IsInSelect(e.Handle, iRow, iCol)
}

// 编辑框_取总行数.
func (e *Edit) GetRowCount() int32 {
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
func (e *Edit) AddData(pData *xc.Edit_Data_Copy_, styleTable []uint16, nStyleCount int32) *Edit {
	xc.XEdit_AddData(e.Handle, pData, styleTable, nStyleCount)
	return e
}

// 编辑框_释放数据.
func (e *Edit) FreeData(pData *xc.Edit_Data_Copy_) *Edit {
	xc.XEdit_FreeData(pData)
	return e
}

// 编辑框_置默认文本, 当内容为空时, 显示默认文本.
//
// pString: 文本内容.
func (e *Edit) SetDefaultText(pString string) *Edit {
	xc.XEdit_SetDefaultText(e.Handle, pString)
	return e
}

// 编辑框_置默认文本颜色.
//
// color: xc.RGBA 颜色值.
func (e *Edit) SetDefaultTextColor(color uint32) *Edit {
	xc.XEdit_SetDefaultTextColor(e.Handle, color)
	return e
}

// 编辑框_置密码字符.
//
// ch: 字符. 使用单引号包裹的字符, 如: '#'.
func (e *Edit) SetPasswordCharacter(ch uint16) *Edit {
	xc.XEdit_SetPasswordCharacter(e.Handle, ch)
	return e
}

// 编辑框_置文本对齐, 单行模式下有效.
//
// align: 对齐方式, Edit_TextAlign_Flag_.
func (e *Edit) SetTextAlign(align xcc.Edit_TextAlign_Flag_) *Edit {
	xc.XEdit_SetTextAlign(e.Handle, align)
	return e
}

// 编辑框_置TAB空格.
//
// nSpace: 空格数量.
func (e *Edit) SetTabSpace(nSpace int32) *Edit {
	xc.XEdit_SetTabSpace(e.Handle, nSpace)
	return e
}

// 编辑框_置文本.
//
// pString: 字符串.
func (e *Edit) SetText(pString string) *Edit {
	xc.XEdit_SetText(e.Handle, pString)
	return e
}

// 编辑框_置文本整数.
//
// nValue: 整数值.
func (e *Edit) SetTextInt(nValue int32) *Edit {
	xc.XEdit_SetTextInt(e.Handle, nValue)
	return e
}

// 编辑框_取文本, 不包含非文本内容, 返回实际接收文本长度.
//
// pOut: 接收文本内存指针.
//
// nOutlen: 内存大小. 例: GetLength()+1 .
func (e *Edit) GetText(pOut *string, nOutlen int32) int32 {
	return xc.XEdit_GetText(e.Handle, pOut, nOutlen)
}

// 编辑框_取文本Ex, 不包含非文本内容, 返回文本.
func (e *Edit) Text() string {
	return e.GetTextEx()
}

// 编辑框_取文本Ex, 不包含非文本内容, 返回文本.
func (e *Edit) GetTextEx() string {
	var s string
	xc.XEdit_GetText(e.Handle, &s, xc.XEdit_GetLength(e.Handle)+1)
	return s
}

// 编辑框_取选择文本Ex, 不包括非文本内容, 返回文本.
func (e *Edit) GetSelectTextEx() string {
	var s string
	xc.XEdit_GetSelectText(e.Handle, &s, xc.XEdit_GetSelectTextLength(e.Handle)+1)
	return s
}

// 编辑框_取文本行.
//
// iRow: 行索引.
//
// pOut: 接收文本内存指针.
//
// nOutlen: 接收文本内存块长度. 例: GetLengthRow()+1 .
func (e *Edit) GetTextRow(iRow int32, pOut *string, nOutlen int32) int32 {
	return xc.XEdit_GetTextRow(e.Handle, iRow, pOut, nOutlen)
}

// 编辑框_取文本行Ex, 返回文本.
//
// iRow: 行索引.
func (e *Edit) GetTextRowEx(iRow int32) string {
	var s string
	xc.XEdit_GetTextRow(e.Handle, iRow, &s, xc.XEdit_GetLengthRow(e.Handle, iRow)+1)
	return s
}

// 编辑框_取内容长度, 包含非文本内容.
func (e *Edit) GetLength() int32 {
	return xc.XEdit_GetLength(e.Handle)
}

// 编辑框_取内容长度行, 包含非文本内容.
//
// iRow: 行索引.
func (e *Edit) GetLengthRow(iRow int32) int32 {
	return xc.XEdit_GetLengthRow(e.Handle, iRow)
}

// 编辑框_取字符, 返回指定位置字符.
//
// iRow: 行索引.
//
// iCol: 列索引.
func (e *Edit) GetAt(iRow, iCol int32) rune {
	return xc.XEdit_GetAt(e.Handle, iRow, iCol)
}

// 编辑框_插入文本.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// pString: 字符串.
func (e *Edit) InsertText(iRow, iCol int32, pString string) *Edit {
	xc.XEdit_InsertText(e.Handle, iRow, iCol, pString)
	return e
}

// AddTextUser 编辑框_插入文本模拟用户操作, 自动刷新UI, 支持撤销/恢复.
//
// pString: 字符串.
func (e *Edit) AddTextUser(pString string) *Edit {
	xc.XEdit_AddTextUser(e.Handle, pString)
	return e
}

// 编辑框_添加文本.
//
// pString: 字符串.
func (e *Edit) AddText(pString string) *Edit {
	xc.XEdit_AddText(e.Handle, pString)
	return e
}

// 编辑框_添加文本扩展.
//
// pString: 字符串.
//
// iStyle: 样式索引.
func (e *Edit) AddTextEx(pString string, iStyle int32) *Edit {
	xc.XEdit_AddTextEx(e.Handle, pString, iStyle)
	return e
}

// 编辑框_添加对象, 例如: 字体, 图片, UI对象, 返回样式索引.
//
// hObj: 对象句柄.
func (e *Edit) AddObject(hObj int) int32 {
	return xc.XEdit_AddObject(e.Handle, hObj)
}

// 编辑框_添加对象从样式, 当样式为图片时有效.
//
// iStyle: 样式索引.
func (e *Edit) AddByStyle(iStyle int32) *Edit {
	xc.XEdit_AddByStyle(e.Handle, iStyle)
	return e
}

// 编辑框_添加样式, 返回样式索引.
//
// hFont_image_Obj: 字体.
//
// color: xc.RGBA 颜色.
//
// bColor: 是否使用颜色.
func (e *Edit) AddStyle(hFont_image_Obj int, color uint32, bColor bool) int32 {
	return xc.XEdit_AddStyle(e.Handle, hFont_image_Obj, color, bColor)
}

// 编辑框_添加样式扩展, 返回样式索引.
//
// fontName: 字体名称.
//
// fontSize: 字体大小.
//
// fontStyle: 字体样式, FontStyle_.
//
// color: xc.RGBA 颜色.
//
// bColor: 是否使用颜色.
func (e *Edit) AddStyleEx(fontName string, fontSize int32, fontStyle xcc.FontStyle_, color uint32, bColor bool) int32 {
	return xc.XEdit_AddStyleEx(e.Handle, fontName, fontSize, fontStyle, color, bColor)
}

// 编辑框_取样式信息.
//
// iStyle: 样式索引.
//
// info: 返回样式信息.
func (e *Edit) GetStyleInfo(iStyle int32, info *xc.Edit_Style_Info_) bool {
	return xc.XEdit_GetStyleInfo(e.Handle, iStyle, info)
}

// 编辑框_置当前样式.
//
// iStyle: 样式索引.
func (e *Edit) SetCurStyle(iStyle int32) *Edit {
	xc.XEdit_SetCurStyle(e.Handle, iStyle)
	return e
}

// 编辑框_置插入符颜色.
//
// color: xc.RGBA 颜色.
func (e *Edit) SetCaretColor(color uint32) *Edit {
	xc.XEdit_SetCaretColor(e.Handle, color)
	return e
}

// 编辑框_置插入符宽度.
//
// nWidth: 宽度.
func (e *Edit) SetCaretWidth(nWidth int32) *Edit {
	xc.XEdit_SetCaretWidth(e.Handle, nWidth)
	return e
}

// 编辑框_置选择背景颜色.
//
// color: xc.RGBA 颜色.
func (e *Edit) SetSelectBkColor(color uint32) *Edit {
	xc.XEdit_SetSelectBkColor(e.Handle, color)
	return e
}

// 编辑框_置默认行高.
//
// nHeight: 行高.
func (e *Edit) SetRowHeight(nHeight int32) *Edit {
	xc.XEdit_SetRowHeight(e.Handle, nHeight)
	return e
}

// 编辑框_置指定行高度, 类型为 Edit_Type_Richedit 支持指定不同行高.
//
// iRow: 行索引.
//
// nHeight: 高度.
func (e *Edit) SetRowHeightEx(iRow, nHeight int32) *Edit {
	xc.XEdit_SetRowHeightEx(e.Handle, iRow, nHeight)
	return e
}

// SetCurPos 编辑框_置当前位置.
//
// iRow: 行索引.
func (e *Edit) SetCurPos(iRow int32) bool {
	return xc.XEdit_SetCurPos(e.Handle, iRow)
}

// 编辑框_取当前位置点, 返回范围位置点.
func (e *Edit) GetCurPos() int32 {
	return xc.XEdit_GetCurPos(e.Handle)
}

// 编辑框_取当前行, 返回行索引.
func (e *Edit) GetCurRow() int32 {
	return xc.XEdit_GetCurRow(e.Handle)
}

// 编辑框_取当前列, 返回列索引.
func (e *Edit) GetCurCol() int32 {
	return xc.XEdit_GetCurCol(e.Handle)
}

// 编辑框_取坐标点.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// pOut: 接收返回坐标点.
func (e *Edit) GetPoint(iRow, iCol int32, pOut *xc.POINT) *Edit {
	xc.XEdit_GetPoint(e.Handle, iRow, iCol, pOut)
	return e
}

// 编辑框_取坐标点Ex.
//
// iRow: 行索引.
//
// iCol: 列索引.
func (e *Edit) GetPointEx(iRow, iCol int32) xc.POINT {
	var pt xc.POINT
	xc.XEdit_GetPoint(e.Handle, iRow, iCol, &pt)
	return pt
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
func (e *Edit) AutoScrollEx(iRow, iCol int32) bool {
	return xc.XEdit_AutoScrollEx(e.Handle, iRow, iCol)
}

// PosToRowCol 编辑框_转换位置, 转换位置点到行列.
//
// iPos: 位置点.
//
// pInfo: 行列.
func (e *Edit) PosToRowCol(iPos int32, pInfo *xc.Position_) *Edit {
	xc.XEdit_PosToRowCol(e.Handle, iPos, pInfo)
	return e
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
func (e *Edit) SetSelect(iStartRow, iStartCol, iEndRow, iEndCol int32) bool {
	return xc.XEdit_SetSelect(e.Handle, iStartRow, iStartCol, iEndRow, iEndCol)
}

// 编辑框_取选择文本, 不包括非文本内容, 返回接收文本内容实际长度.
//
// pOut: 接收返回文本内容.
//
// nOutLen: 接收内存大小. GetSelectTextLength()+1 .
func (e *Edit) GetSelectText(pOut *string, nOutLen int32) int32 {
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
func (e *Edit) GetVisibleRowRange(piStart *int32, piEnd *int32) *Edit {
	xc.XEdit_GetVisibleRowRange(e.Handle, piStart, piEnd)
	return e
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
func (e *Edit) Delete(iStartRow, iStartCol, iEndRow, iEndCol int32) bool {
	return xc.XEdit_Delete(e.Handle, iStartRow, iStartCol, iEndRow, iEndCol)
}

// 编辑框_删除行.
//
// iRow: 行索引.
func (e *Edit) DeleteRow(iRow int32) bool {
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
func (e *Edit) AddChatBegin(hImageAvatar int, hImageBubble int, nFlag xcc.Chat_Flag_) *Edit {
	xc.XEdit_AddChatBegin(e.Handle, hImageAvatar, hImageBubble, nFlag)
	return e
}

// 编辑框_添加气泡结束, 当前行结束.
func (e *Edit) AddChatEnd() *Edit {
	xc.XEdit_AddChatEnd(e.Handle)
	return e
}

// 编辑框_置气泡缩进, 设置聊天气泡内容缩进.
//
// nIndentation: 缩进值.
func (e *Edit) SetChatIndentation(nIndentation int32) *Edit {
	xc.XEdit_SetChatIndentation(e.Handle, nIndentation)
	return e
}

// 编辑框_置行间隔, 设置行间隔大小, 多行模式有效.
//
// nSpace: 行间隔大小.
func (e *Edit) SetRowSpace(nSpace int32) *Edit {
	xc.XEdit_SetRowSpace(e.Handle, nSpace)
	return e
}

// 编辑框_置当前位置扩展.
//
// iRow: 行索引.
//
// iCol: 列索引.
func (e *Edit) SetCurPosEx(iRow, iCol int32) *Edit {
	xc.XEdit_SetCurPosEx(e.Handle, iRow, iCol)
	return e
}

// 编辑框_取当前位置扩展.
//
// iRow: 返回行索引.
//
// iCol: 返回列索引.
func (e *Edit) GetCurPosEx(iRow, iCol *int32) *Edit {
	xc.XEdit_GetCurPosEx(e.Handle, iRow, iCol)
	return e
}

// 编辑框_移动到末尾.
func (e *Edit) MoveEnd() *Edit {
	xc.XEdit_MoveEnd(e.Handle)
	return e
}

// 编辑框_行列到位置, 返回位置点.
//
// iRow: 行索引.
//
// iCol: 列索引.
func (e *Edit) RowColToPos(iRow, iCol int32) int32 {
	return xc.XEdit_RowColToPos(e.Handle, iRow, iCol)
}

// 编辑框_置后备字体, 置中文字体. 如果已设置, 当遇到中文字符时使用后备字体, 解决不支持中文的字体的问题
//
// hFont: 字体.
func (e *Edit) SetBackFont(hFont int) *Edit {
	xc.XEdit_SetBackFont(e.Handle, hFont)
	return e
}

// 编辑框_释放样式.
//
// iStyle: 样式.
func (e *Edit) ReleaseStyle(iStyle int32) bool {
	return xc.XEdit_ReleaseStyle(e.Handle, iStyle)
}

// 编辑框_修改样式.
//
// iStyle: 样式索引.
//
// hFont: 字体句柄.
//
// color: xc.RGBA 颜色.
//
// bColor: 是否使用颜色.
func (e *Edit) ModifyStyle(iStyle int32, hFont int, color uint32, bColor bool) bool {
	return xc.XEdit_ModifyStyle(e.Handle, iStyle, hFont, color, bColor)
}

// 编辑框_置空格大小.
//
// size: 空格大小.
func (e *Edit) SetSpaceSize(size int32) *Edit {
	xc.XEdit_SetSpaceSize(e.Handle, size)
	return e
}

// 编辑框_置字符间距.
//
// size: 英文字符间距大小.
//
// sizeZh: 中文字符间距大小.
func (e *Edit) SetCharSpaceSize(size, sizeZh int32) *Edit {
	xc.XEdit_SetCharSpaceSize(e.Handle, size, sizeZh)
	return e
}

// 编辑框_取选择文本长度, 不包括非文本内容, 返回文本内容长度.
func (e *Edit) GetSelectTextLength() int32 {
	return xc.XEdit_GetSelectTextLength(e.Handle)
}

// 编辑框_置选择文本样式.
//
// iStyle: 样式索引.
func (e *Edit) SetSelectTextStyle(iStyle int32) *Edit {
	xc.XEdit_SetSelectTextStyle(e.Handle, iStyle)
	return e
}

// 编辑框_取文本_临时, 不包含非文本内容. 返回临时文本, 临时缓存区大小: xcc.Text_Buffer_Size .
func (e *Edit) GetText_Temp() string {
	return xc.XEdit_GetText_Temp(e.Handle)
}

// 编辑框_取文本行_临时, 获取指定行文本内容. 返回临时文本, 临时缓存区大小: xcc.Text_Buffer_Size .
//
// iRow: 行索引.
func (e *Edit) GetTextRow_Temp(iRow int32) string {
	return xc.XEdit_GetTextRow_Temp(e.Handle, iRow)
}

// 编辑框_取选择文本, 不包含非文本内容. 返回临时文本, 临时缓存区大小: xcc.Text_Buffer_Size .
func (e *Edit) GetSelectText_Temp() string {
	return xc.XEdit_GetSelectText_Temp(e.Handle)
}

// 编辑框_插入气泡开始, 当前行开始.
//
// hImageAvatar: 头像图片句柄.
//
// hImageBubble: 气泡背景图片句柄.
//
// nFlag: 聊天气泡对齐方式: xcc.Chat_Flag_ .
func (e *Edit) InsertChatBegin(hImageAvatar int, hImageBubble int, nFlag xcc.Chat_Flag_) *Edit {
	xc.XEdit_InsertChatBegin(e.Handle, hImageAvatar, hImageBubble, nFlag)
	return e
}

// 编辑框_取指定行气泡标识. 返回行标识: xcc.Chat_Flag_
//
// iRow: 行索引.
func (e *Edit) GetChatFlags(iRow int32) xcc.Chat_Flag_ {
	return xc.XEdit_GetChatFlags(e.Handle, iRow)
}

// 编辑框_插入文本扩展.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// pString: 字符串.
//
// iStyle: 样式.
func (e *Edit) InsertTextEx(iRow, iCol int32, pString string, iStyle int32) *Edit {
	xc.XEdit_InsertTextEx(e.Handle, iRow, iCol, pString, iStyle)
	return e
}

// 编辑框_插入对象.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// hObj: 对象句柄.
func (e *Edit) InsertObject(iRow, iCol int32, hObj int) *Edit {
	xc.XEdit_InsertObject(e.Handle, iRow, iCol, hObj)
	return e
}

// 编辑框_置气泡最大宽度. 当值为0时代表不限制宽度.
//
// nWidth: 最大宽度.
func (e *Edit) SetChatMaxWidth(nWidth int32) *Edit {
	xc.XEdit_SetChatMaxWidth(e.Handle, nWidth)
	return e
}

// 编辑框_取总行数扩展. 包含自动换行数量, 返回总行数.
func (e *Edit) GetRowCountEx() int32 {
	return xc.XEdit_GetRowCountEx(e.Handle)
}

// 编辑框_剪贴板复制. 复制全部内容.
func (e *Edit) ClipboardCopyAll() bool {
	return xc.XEdit_ClipboardCopyAll(e.Handle)
}

// ------------------------- AddEvent ------------------------- //

// AddEvent_Edit_Set 添加编辑框设置事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Edit) AddEvent_Edit_Set(pFun xc.XE_EDIT_SET1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_EDIT_SET, onXE_EDIT_SET, pFun, allowAddingMultiple...)
}

// onXE_EDIT_SET 编辑框设置事件.
func onXE_EDIT_SET(hEle int, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_EDIT_SET)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_EDIT_SET1); ok {
			ret = cb(hEle, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_Edit_DrawRow 添加编辑框绘制行事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Edit) AddEvent_Edit_DrawRow(pFun xc.XE_EDIT_DRAWROW1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_EDIT_DRAWROW, onXE_EDIT_DRAWROW, pFun, allowAddingMultiple...)
}

// onXE_EDIT_DRAWROW 编辑框绘制行事件.
func onXE_EDIT_DRAWROW(hEle int, hDraw int, iRow int32, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_EDIT_DRAWROW)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_EDIT_DRAWROW1); ok {
			ret = cb(hEle, hDraw, iRow, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_Edit_Changed 添加编辑框内容被改变事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Edit) AddEvent_Edit_Changed(pFun xc.XE_EDIT_CHANGED1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_EDIT_CHANGED, onXE_EDIT_CHANGED, pFun, allowAddingMultiple...)
}

// onXE_EDIT_CHANGED 编辑框内容被改变事件.
func onXE_EDIT_CHANGED(hEle int, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_EDIT_CHANGED)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_EDIT_CHANGED1); ok {
			ret = cb(hEle, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_Edit_Pos_Changed 添加编辑框光标位置被改变事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Edit) AddEvent_Edit_Pos_Changed(pFun xc.XE_EDIT_POS_CHANGED1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_EDIT_POS_CHANGED, onXE_EDIT_POS_CHANGED, pFun, allowAddingMultiple...)
}

// onXE_EDIT_POS_CHANGED 编辑框光标位置被改变事件.
func onXE_EDIT_POS_CHANGED(hEle int, iPos int32, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_EDIT_POS_CHANGED)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_EDIT_POS_CHANGED1); ok {
			ret = cb(hEle, iPos, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_Edit_Style_Changed 添加编辑框样式被改变事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Edit) AddEvent_Edit_Style_Changed(pFun xc.XE_EDIT_STYLE_CHANGED1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_EDIT_STYLE_CHANGED, onXE_EDIT_STYLE_CHANGED, pFun, allowAddingMultiple...)
}

// onXE_EDIT_STYLE_CHANGED 编辑框样式被改变事件.
func onXE_EDIT_STYLE_CHANGED(hEle int, iStyle int32, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_EDIT_STYLE_CHANGED)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_EDIT_STYLE_CHANGED1); ok {
			ret = cb(hEle, iStyle, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_Edit_Enter_Get_TabAlign 添加编辑框回车TAB对齐事件, 返回需要TAB数量.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Edit) AddEvent_Edit_Enter_Get_TabAlign(pFun xc.XE_EDIT_ENTER_GET_TABALIGN1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_EDIT_ENTER_GET_TABALIGN, onXE_EDIT_ENTER_GET_TABALIGN, pFun, allowAddingMultiple...)
}

// onXE_EDIT_ENTER_GET_TABALIGN 编辑框回车TAB对齐事件.
func onXE_EDIT_ENTER_GET_TABALIGN(hEle int, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_EDIT_ENTER_GET_TABALIGN)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_EDIT_ENTER_GET_TABALIGN1); ok {
			ret = cb(hEle, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_Edit_Row_Changed 添加编辑框行被改变事件.
//   - iRow: 更改行开始位置索引, if(nChangeRows>0) iEnd= iRow + nChangeRows
//   - nChangeRows: 改变行数, 正数添加行, 负数删除行
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Edit) AddEvent_Edit_Row_Changed(pFun xc.XE_EDIT_ROW_CHANGED1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_EDIT_ROW_CHANGED, onXE_EDIT_ROW_CHANGED, pFun, allowAddingMultiple...)
}

// onXE_EDIT_ROW_CHANGED 编辑框行被改变事件.
func onXE_EDIT_ROW_CHANGED(hEle int, iRow int32, nChangeRows int32, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_EDIT_ROW_CHANGED)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_EDIT_ROW_CHANGED1); ok {
			ret = cb(hEle, iRow, nChangeRows, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_Edit_SwapRow 添加编辑框交换行事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Edit) AddEvent_Edit_SwapRow(pFun xc.XE_EDIT_SWAPROW1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_EDIT_SWAPROW, onXE_EDIT_SWAPROW, pFun, allowAddingMultiple...)
}

// onXE_EDIT_SWAPROW 编辑框交换行事件.
func onXE_EDIT_SWAPROW(hEle int, iRow int32, bArrowUp int32, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_EDIT_SWAPROW)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_EDIT_SWAPROW1); ok {
			ret = cb(hEle, iRow, bArrowUp, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_Edit_Color_Change 添加编辑框颜色被改变事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (e *Edit) AddEvent_Edit_Color_Change(pFun xc.XE_EDIT_COLOR_CHANGE1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(e.Handle, xcc.XE_EDIT_COLOR_CHANGE, onXE_EDIT_COLOR_CHANGE, pFun, allowAddingMultiple...)
}

// onXE_EDIT_COLOR_CHANGE 编辑框颜色被改变事件.
func onXE_EDIT_COLOR_CHANGE(hEle int, color uint32, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_EDIT_COLOR_CHANGE)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_EDIT_COLOR_CHANGE1); ok {
			ret = cb(hEle, color, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// ------------------------- 事件 ------------------------- //

// 编辑框_颜色被改变.
func (e *Edit) Event_EDIT_COLOR_CHANGE(pFun xc.XE_EDIT_COLOR_CHANGE) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_EDIT_COLOR_CHANGE, pFun)
}

// 编辑框_颜色被改变.
func (e *Edit) Event_EDIT_COLOR_CHANGE1(pFun xc.XE_EDIT_COLOR_CHANGE1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_EDIT_COLOR_CHANGE, pFun)
}

// 元素事件_交换行.
func (e *Edit) Event_EDIT_SWAPROW(pFun xc.XE_EDIT_SWAPROW) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_EDIT_SWAPROW, pFun)
}

// 元素事件_交换行.
func (e *Edit) Event_EDIT_SWAPROW1(pFun xc.XE_EDIT_SWAPROW1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_EDIT_SWAPROW, pFun)
}

// 元素事件_编辑框设置.
func (e *Edit) Event_EDIT_SET(pFun xc.XE_EDIT_SET) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_EDIT_SET, pFun)
}

// 元素事件_编辑框设置.
func (e *Edit) Event_EDIT_SET1(pFun xc.XE_EDIT_SET1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_EDIT_SET, pFun)
}

// 暂未使用.
func (e *Edit) Event_EDIT_DRAWROW(pFun xc.XE_EDIT_DRAWROW) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_EDIT_DRAWROW, pFun)
}

// 暂未使用.
func (e *Edit) Event_EDIT_DRAWROW1(pFun xc.XE_EDIT_DRAWROW1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_EDIT_DRAWROW, pFun)
}

// 编辑框_内容被改变.
func (e *Edit) Event_EDIT_CHANGED(pFun xc.XE_EDIT_CHANGED) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_EDIT_CHANGED, pFun)
}

// 编辑框_内容被改变.
func (e *Edit) Event_EDIT_CHANGED1(pFun xc.XE_EDIT_CHANGED1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_EDIT_CHANGED, pFun)
}

// 编辑框_光标位置_被改变.
func (e *Edit) Event_EDIT_POS_CHANGED(pFun xc.XE_EDIT_POS_CHANGED) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_EDIT_POS_CHANGED, pFun)
}

// 编辑框_光标位置_被改变.
func (e *Edit) Event_EDIT_POS_CHANGED1(pFun xc.XE_EDIT_POS_CHANGED1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_EDIT_POS_CHANGED, pFun)
}

// 编辑框_样式_被改变.
func (e *Edit) Event_EDIT_STYLE_CHANGED(pFun xc.XE_EDIT_STYLE_CHANGED) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_EDIT_STYLE_CHANGED, pFun)
}

// 编辑框_样式_被改变.
func (e *Edit) Event_EDIT_STYLE_CHANGED1(pFun xc.XE_EDIT_STYLE_CHANGED1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_EDIT_STYLE_CHANGED, pFun)
}

// 回车TAB对齐,返回需要TAB数量.
func (e *Edit) Event_EDIT_ENTER_GET_TABALIGN(pFun xc.XE_EDIT_ENTER_GET_TABALIGN) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_EDIT_ENTER_GET_TABALIGN, pFun)
}

// 回车TAB对齐,返回需要TAB数量.
func (e *Edit) Event_EDIT_ENTER_GET_TABALIGN1(pFun xc.XE_EDIT_ENTER_GET_TABALIGN1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_EDIT_ENTER_GET_TABALIGN, pFun)
}

// 编辑框_行_被改变.
//
// iRow: 更改行开始位置索引,  if(nChangeRows>0) iEnd= iRow + nChangeRows
//
// nChangeRows: 改变行数, 正数添加行, 负数删除行
func (e *Edit) Event_EDIT_ROW_CHANGED(pFun xc.XE_EDIT_ROW_CHANGED) bool {
	return xc.XEle_RegEventC(e.Handle, xcc.XE_EDIT_ROW_CHANGED, pFun)
}

// 编辑框_行_被改变.
//
// iRow: 更改行开始位置索引,  if(nChangeRows>0) iEnd= iRow + nChangeRows
//
// nChangeRows: 改变行数, 正数添加行, 负数删除行
func (e *Edit) Event_EDIT_ROW_CHANGED1(pFun xc.XE_EDIT_ROW_CHANGED1) bool {
	return xc.XEle_RegEventC1(e.Handle, xcc.XE_EDIT_ROW_CHANGED, pFun)
}

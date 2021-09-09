package widget

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// 按钮
type Button struct {
	Element
}

// 按钮_创建
// x: 按钮x坐标
// y: 按钮y坐标
// cx: 宽度
// cy: 高度
// pName: 标题
// hParent: 父为窗口句柄或元素句柄.
func NewButton(x int, y int, cx int, cy int, pName string, hParent int) *Button {
	p := &Button{}
	p.SetHandle(xc.XBtn_Create(x, y, cx, cy, pName, hParent))
	return p
}

// 按钮_判断选中, 是否选中状态
func (b *Button) IsCheck() bool {
	return xc.XBtn_IsCheck(b.Handle)
}

// 按钮_置选中, 设置选中状态
// bCheck: 是否选中.
func (b *Button) SetCheck(bCheck bool) bool {
	return xc.XBtn_SetCheck(b.Handle, bCheck)
}

// 按钮_置样式
// nStyle: 样式, Button_Style_
func (b *Button) SetStyle(nStyle int) int {
	return xc.XBtn_SetStyle(b.Handle, nStyle)
}

// 按钮_置状态
// nState: 按钮状态, Common_State3_
func (b *Button) SetState(nState int) int {
	return xc.XBtn_SetState(b.Handle, nState)
}

// 按钮_取状态, 返回: Common_State3_
func (b *Button) GetState() int {
	return xc.XBtn_GetState(b.Handle)
}

// 按钮_取状态扩展, 返回: Button_State_
func (b *Button) GetStateEx() int {
	return xc.XBtn_GetStateEx(b.Handle)
}

// 按钮_取样式, 返回: Button_Style_
func (b *Button) GetStyle() int {
	return xc.XBtn_GetStyle(b.Handle)
}

// 按钮_置类型, 设置按钮功能类型.
// nType: 按钮类型, Button_Type_
func (b *Button) SetType(nType int) int {
	return xc.XBtn_SetType(b.Handle, nType)
}

// 按钮_置类型扩展, 设置按钮类型并自动修改样式和文本对齐方式.
// nType: 按钮类型, Button_Type_
func (b *Button) SetTypeEx(nType int) int {
	return xc.XBtn_SetTypeEx(b.Handle, nType)
}

// 按钮_取类型, 获取按钮功能类型, 返回: Button_Type_
func (b *Button) GetType() int {
	return xc.XBtn_GetType(b.Handle)
}

// 按钮_置组ID
// nID: 组ID.
func (b *Button) SetGroupID(nID int) int {
	return xc.XBtn_SetGroupID(b.Handle, nID)
}

// 按钮_取组ID
func (b *Button) GetGroupID() int {
	return xc.XBtn_GetGroupID(b.Handle)
}

// 按钮_置绑定元素
// hBindEle: 将要绑定的元素.
func (b *Button) SetBindEle(hBindEle int) int {
	return xc.XBtn_SetBindEle(b.Handle, hBindEle)
}

// 按钮_取绑定元素, 返回: 绑定的元素句柄
func (b *Button) GetBindEle() int {
	return xc.XBtn_GetBindEle(b.Handle)
}

// 按钮_置文本对齐
// nFlags: 对齐方式, TextFormatFlag_ , TextAlignFlag_ , TextTrimming_
func (b *Button) SetTextAlign(nFlags int) int {
	return xc.XBtn_SetTextAlign(b.Handle, nFlags)
}

// 按钮_取文本对齐方式, 返回: TextFormatFlag_ , TextAlignFlag_ , TextTrimming_
func (b *Button) GetTextAlign() int {
	return xc.XBtn_GetTextAlign(b.Handle)
}

// 按钮_置图标对齐
// align: 对齐方式, Button_Icon_Align_
func (b *Button) SetIconAlign(align int) int {
	return xc.XBtn_SetIconAlign(b.Handle, align)
}

// 按钮_置偏移, 设置按钮文本坐标偏移量
// x: 偏移量.
// y: 偏移量.
func (b *Button) SetOffset(x int, y int) int {
	return xc.XBtn_SetOffset(b.Handle, x, y)
}

// 按钮_置图标偏移, 设置按钮图标坐标偏移量
// x: 偏移量.
// y: 偏移量.
func (b *Button) SetOffsetIcon(x int, y int) int {
	return xc.XBtn_SetOffsetIcon(b.Handle, x, y)
}

// 按钮_置图标间隔, 设置图标与文本间隔大小
// size: 间隔大小.
func (b *Button) SetIconSpace(size int) int {
	return xc.XBtn_SetIconSpace(b.Handle, size)
}

// 按钮_置文本
// pName: 文本内容.
func (b *Button) SetText(pName string) int {
	return xc.XBtn_SetText(b.Handle, pName)
}

// 按钮_取文本
func (b *Button) GetText() string {
	return xc.XBtn_GetText(b.Handle)
}

// 按钮_置图标
// hImage: 图片句柄.
func (b *Button) SetIcon(hImage int) int {
	return xc.XBtn_SetIcon(b.Handle, hImage)
}

// 按钮_置禁用图标
// hImage: 图片句柄.
func (b *Button) SetIconDisable(hImage int) int {
	return xc.XBtn_SetIconDisable(b.Handle, hImage)
}

// 按钮_取图标, 返回图标句柄
// nType: 图标类型, 0:默认图标, 1:禁用状态图标.
func (b *Button) GetIcon(nType int) int {
	return xc.XBtn_GetIcon(b.Handle, nType)
}

// 按钮_添加动画帧
// hImage: 图片句柄
// uElapse: 图片帧延时, 单位毫秒.
func (b *Button) AddAnimationFrame(hImage int, uElapse int) int {
	return xc.XBtn_AddAnimationFrame(b.Handle, hImage, uElapse)
}

// 按钮_启用动画, 开始或关闭图片动画的播放
// bEnable: 开始播放动画TRUE, 关闭播放动画FALSE.
// bLoopPlay: 是否循环播放.
func (b *Button) EnableAnimation(bEnable bool, bLoopPlay bool) int {
	return xc.XBtn_EnableAnimation(b.Handle, bEnable, bLoopPlay)
}

// 按钮_添加背景边框, 添加背景内容边框.
// nState: 按钮状态, Button_State_
// color: RGB颜色.
// alpha: 透明度.
// width: 线宽.
func (b *Button) AddBkBorder(nState int, color int, alpha uint8, width int) int {
	return xc.XBtn_AddBkBorder(b.Handle, nState, color, alpha, width)
}

// 按钮_添加填充背景, 添加背景内容填充
// nState: 按钮状态, Button_State_
// color: RGB颜色.
// alpha: 透明度.
func (b *Button) AddBkFill(nState int, color int, alpha uint8) int {
	return xc.XBtn_AddBkFill(b.Handle, nState, color, alpha)
}

// 按钮_添加背景图片
// nState: 按钮状态, Button_State_
// hImage: 图片句柄.
func (b *Button) AddBkImage(nState int, hImage int) int {
	return xc.XBtn_AddBkImage(b.Handle, nState, hImage)
}

// 按钮_取背景对象数量, 获取背景内容数量, 成功返回背景内容数量, 否则返回: XC_ID_ERROR
func (b *Button) GetBkInfoCount() int {
	return xc.XBtn_GetBkInfoCount(b.Handle)
}

// 按钮_清空背景对象, 清空背景内容; 如果背景没有内容,将使用系统默认内容,以便保证背景正确
func (b *Button) ClearBkInfo() int {
	return xc.XBtn_ClearBkInfo(b.Handle)
}

/*
下面都是事件
*/

// 事件_按钮被单击
// pFun: 事件函数指针.
func (b *Button) Event_BnClick(pFun func(pbHandled *bool) int) bool {
	return xc.XEle_RegEventC(b.Handle, xcc.XE_BNCLICK, pFun)
}

// 事件_按钮被单击1
// pFun: 事件函数指针.
func (b *Button) Event_BnClick1(pFun func(hEle int, pbHandled *bool) int) bool {
	return xc.XEle_RegEventC1(b.Handle, xcc.XE_BNCLICK, pFun)
}

package widget

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// Button 按钮.
type Button struct {
	Element
}

// 按钮_创建.
//
// x: 按钮x坐标.
//
// y: 按钮y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// pName: 标题.
//
// hParent: 父为窗口句柄或元素句柄.
func NewButton(x, y, cx, cy int32, pName string, hParent int) *Button {
	p := &Button{}
	p.SetHandle(xc.XBtn_Create(x, y, cx, cy, pName, hParent))
	return p
}

// 从句柄创建对象.
func NewButtonByHandle(handle int) *Button {
	p := &Button{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func NewButtonByName(name string) *Button {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &Button{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func NewButtonByUID(nUID int32) *Button {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &Button{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func NewButtonByUIDName(name string) *Button {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &Button{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 按钮_判断选中, 是否选中状态.
func (b *Button) IsCheck() bool {
	return xc.XBtn_IsCheck(b.Handle)
}

// 按钮_置选中, 设置选中状态.
//
// bCheck: 是否选中.
func (b *Button) SetCheck(bCheck bool) bool {
	return xc.XBtn_SetCheck(b.Handle, bCheck)
}

// SetState 按钮_置状态.
//
// nState: 按钮状态: xcc.Common_State3_.
func (b *Button) SetState(nState xcc.Common_State3_) *Button {
	xc.XBtn_SetState(b.Handle, nState)
	return b
}

// GetState 按钮_取状态.
func (b *Button) GetState() xcc.Common_State3_ {
	return xc.XBtn_GetState(b.Handle)
}

// GetStateEx 按钮_取状态扩展.
func (b *Button) GetStateEx() xcc.Button_State_ {
	return xc.XBtn_GetStateEx(b.Handle)
}

// 按钮_置类型扩展, 设置按钮类型并自动修改样式和文本对齐方式.
//
// nType: 按钮类型, Button_Type_ , element_type_ , xc_ex_error.
func (b *Button) SetTypeEx(nType xcc.XC_OBJECT_TYPE_EX) *Button {
	xc.XBtn_SetTypeEx(b.Handle, nType)
	return b
}

// 按钮_置组ID.
//
// nID: 组ID.
func (b *Button) SetGroupID(nID int32) *Button {
	xc.XBtn_SetGroupID(b.Handle, nID)
	return b
}

// 按钮_取组ID.
func (b *Button) GetGroupID() int32 {
	return xc.XBtn_GetGroupID(b.Handle)
}

// 按钮_置绑定元素.
//
// hBindEle: 将要绑定的元素.
func (b *Button) SetBindEle(hBindEle int) *Button {
	xc.XBtn_SetBindEle(b.Handle, hBindEle)
	return b
}

// 按钮_取绑定元素, 返回: 绑定的元素句柄.
func (b *Button) GetBindEle() int {
	return xc.XBtn_GetBindEle(b.Handle)
}

// 按钮_置文本对齐.
//
// nFlags: 对齐方式, TextFormatFlag_ , TextAlignFlag_ , TextTrimming_.
func (b *Button) SetTextAlign(nFlags xcc.TextFormatFlag_) *Button {
	xc.XBtn_SetTextAlign(b.Handle, nFlags)
	return b
}

// 按钮_取文本对齐方式, 返回: TextFormatFlag_ , TextAlignFlag_ , TextTrimming_.
func (b *Button) GetTextAlign() xcc.TextFormatFlag_ {
	return xc.XBtn_GetTextAlign(b.Handle)
}

// 按钮_置图标对齐.
//
// align: 对齐方式, Button_Icon_Align_.
func (b *Button) SetIconAlign(align xcc.Button_Icon_Align_) *Button {
	xc.XBtn_SetIconAlign(b.Handle, align)
	return b
}

// 按钮_置偏移, 设置按钮文本坐标偏移量.
//
// x: 偏移量.
//
// y: 偏移量.
func (b *Button) SetOffset(x, y int32) *Button {
	xc.XBtn_SetOffset(b.Handle, x, y)
	return b
}

// 按钮_置图标偏移, 设置按钮图标坐标偏移量.
//
// x: 偏移量.
//
// y: 偏移量.
func (b *Button) SetOffsetIcon(x, y int32) *Button {
	xc.XBtn_SetOffsetIcon(b.Handle, x, y)
	return b
}

// 按钮_置图标间隔, 设置图标与文本间隔大小.
//
// size: 间隔大小.
func (b *Button) SetIconSpace(size int32) *Button {
	xc.XBtn_SetIconSpace(b.Handle, size)
	return b
}

// 按钮_置文本.
//
// pName: 文本内容.
func (b *Button) SetText(pName string) *Button {
	xc.XBtn_SetText(b.Handle, pName)
	return b
}

// 按钮_取文本.
func (b *Button) GetText() string {
	return xc.XBtn_GetText(b.Handle)
}

// 按钮_置图标.
//
// hImage: 图片句柄.
func (b *Button) SetIcon(hImage int) *Button {
	xc.XBtn_SetIcon(b.Handle, hImage)
	return b
}

// 按钮_置禁用图标.
//
// hImage: 图片句柄.
func (b *Button) SetIconDisable(hImage int) *Button {
	xc.XBtn_SetIconDisable(b.Handle, hImage)
	return b
}

// 按钮_取图标, 返回图标句柄.
//
// nType: 图标类型, 0:默认图标, 1:禁用状态图标.
func (b *Button) GetIcon(nType int32) int {
	return xc.XBtn_GetIcon(b.Handle, nType)
}

// 按钮_添加动画帧.
//
// hImage: 图片句柄.
//
// uElapse: 图片帧延时, 单位毫秒.
func (b *Button) AddAnimationFrame(hImage int, uElapse uint32) *Button {
	xc.XBtn_AddAnimationFrame(b.Handle, hImage, uElapse)
	return b
}

// 按钮_启用动画, 开始或关闭图片动画的播放.
//
// bEnable: 开始播放动画TRUE, 关闭播放动画FALSE.
//
// bLoopPlay: 是否循环播放.
func (b *Button) EnableAnimation(bEnable bool, bLoopPlay bool) *Button {
	xc.XBtn_EnableAnimation(b.Handle, bEnable, bLoopPlay)
	return b
}

// 按钮_启用热键前缀. 对当前文本内容处理, 将&符号后面的一个字符加上下划线.
//
// bEnable: 是否启用.
func (b *Button) EnableHotkeyPrefix(bEnable bool) *Button {
	xc.XBtn_EnableHotkeyPrefix(b.Handle, bEnable)
	return b
}

// 按钮_清除动画.
func (b *Button) ClearAnimation() *Button {
	xc.XBtn_ClearAnimation(b.Handle)
	return b
}

// ------------------------- AddEvent ------------------------- //

// AddEvent_BnClick 添加按钮点击事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (b *Button) AddEvent_BnClick(pFun xc.XE_BNCLICK1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(b.Handle, xcc.XE_BNCLICK, onXE_BNCLICK, pFun, allowAddingMultiple...)
}

// onXE_BNCLICK 按钮点击事件.
func onXE_BNCLICK(hEle int, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_BNCLICK)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_BNCLICK1); ok {
			ret = cb(hEle, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// AddEvent_Button_Check 添加按钮选中事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (b *Button) AddEvent_Button_Check(pFun xc.XE_BUTTON_CHECK1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(b.Handle, xcc.XE_BUTTON_CHECK, onXE_BUTTON_CHECK, pFun, allowAddingMultiple...)
}

// onXE_BUTTON_CHECK 按钮选中事件.
func onXE_BUTTON_CHECK(hEle int, bCheck bool, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_BUTTON_CHECK)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		if cb, ok := cbs[i].(xc.XE_BUTTON_CHECK1); ok {
			ret = cb(hEle, bCheck, pbHandled)
		}
		if *pbHandled {
			break
		}
	}
	return ret
}

// ------------------------- 事件 ------------------------- //

// 事件_按钮被单击.
func (b *Button) Event_BnClick(pFun xc.XE_BNCLICK) bool {
	return xc.XEle_RegEventC(b.Handle, xcc.XE_BNCLICK, pFun)
}

// 事件_按钮被单击1.
func (b *Button) Event_BnClick1(pFun xc.XE_BNCLICK1) bool {
	return xc.XEle_RegEventC1(b.Handle, xcc.XE_BNCLICK, pFun)
}

// 按钮选中事件.
func (b *Button) Event_BUTTON_CHECK(pFun xc.XE_BUTTON_CHECK) bool {
	return xc.XEle_RegEventC(b.Handle, xcc.XE_BUTTON_CHECK, pFun)
}

// 按钮选中事件.
func (b *Button) Event_BUTTON_CHECK1(pFun xc.XE_BUTTON_CHECK1) bool {
	return xc.XEle_RegEventC1(b.Handle, xcc.XE_BUTTON_CHECK, pFun)
}

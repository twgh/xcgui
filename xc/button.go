package xc

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
func XBtn_Create(x int, y int, cx int, cy int, pName string, hParent int) int {
	r, _, _ := xBtn_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), strPtr(pName), uintptr(hParent))
	return int(r)
}

// 按钮_判断选中, 是否选中状态.
//
// hEle: 元素句柄.
func XBtn_IsCheck(hEle int) bool {
	r, _, _ := xBtn_IsCheck.Call(uintptr(hEle))
	return int(r) != 0
}

// 按钮_置选中, 设置选中状态.
//
// hEle: 元素句柄.
//
// bCheck: 是否选中.
func XBtn_SetCheck(hEle int, bCheck bool) bool {
	r, _, _ := xBtn_SetCheck.Call(uintptr(hEle), boolPtr(bCheck))
	return int(r) != 0
}

// 按钮_置状态.
//
// hEle: 元素句柄.
//
// nState: 按钮状态, Common_State3_.
func XBtn_SetState(hEle int, nState int) int {
	r, _, _ := xBtn_SetState.Call(uintptr(hEle), uintptr(nState))
	return int(r)
}

// 按钮_取状态, 返回: Common_State3_.
//
// hEle: 元素句柄.
func XBtn_GetState(hEle int) int {
	r, _, _ := xBtn_GetState.Call(uintptr(hEle))
	return int(r)
}

// 按钮_取状态扩展, 返回: Button_State_.
//
// hEle: 元素句柄.
func XBtn_GetStateEx(hEle int) int {
	r, _, _ := xBtn_GetStateEx.Call(uintptr(hEle))
	return int(r)
}

// 按钮_置类型扩展, 设置按钮类型并自动修改样式和文本对齐方式.
//
// hEle: 元素句柄.
//
// nType: 按钮类型, Button_Type_.
func XBtn_SetTypeEx(hEle int, nType int) int {
	r, _, _ := xBtn_SetTypeEx.Call(uintptr(hEle), uintptr(nType))
	return int(r)
}

// 按钮_置组ID.
//
// hEle: 元素句柄.
//
// nID: 组ID.
func XBtn_SetGroupID(hEle int, nID int) int {
	r, _, _ := xBtn_SetGroupID.Call(uintptr(hEle), uintptr(nID))
	return int(r)
}

// 按钮_取组ID.
//
// hEle: 元素句柄.
func XBtn_GetGroupID(hEle int) int {
	r, _, _ := xBtn_GetGroupID.Call(uintptr(hEle))
	return int(r)
}

// 按钮_置绑定元素.
//
// hEle: 元素句柄.
//
// hBindEle: 将要绑定的元素.
func XBtn_SetBindEle(hEle int, hBindEle int) int {
	r, _, _ := xBtn_SetBindEle.Call(uintptr(hEle), uintptr(hBindEle))
	return int(r)
}

// 按钮_取绑定元素, 返回: 绑定的元素句柄.
//
// hEle: 元素句柄.
func XBtn_GetBindEle(hEle int) int {
	r, _, _ := xBtn_GetBindEle.Call(uintptr(hEle))
	return int(r)
}

// 按钮_置文本对齐.
//
// hEle: 元素句柄.
//
// nFlags: 对齐方式, TextFormatFlag_ , TextAlignFlag_ , TextTrimming_.
func XBtn_SetTextAlign(hEle int, nFlags int) int {
	r, _, _ := xBtn_SetTextAlign.Call(uintptr(hEle), uintptr(nFlags))
	return int(r)
}

// 按钮_取文本对齐方式, 返回: TextFormatFlag_ , TextAlignFlag_ , TextTrimming_.
//
// hEle: 元素句柄.
func XBtn_GetTextAlign(hEle int) int {
	r, _, _ := xBtn_GetTextAlign.Call(uintptr(hEle))
	return int(r)
}

// 按钮_置图标对齐.
//
// hEle: 元素句柄.
//
// align: 对齐方式, Button_Icon_Align_.
func XBtn_SetIconAlign(hEle int, align int) int {
	r, _, _ := xBtn_SetIconAlign.Call(uintptr(hEle), uintptr(align))
	return int(r)
}

// 按钮_置偏移, 设置按钮文本坐标偏移量.
//
// hEle: 元素句柄.
//
// x: 偏移量.
//
// y: 偏移量.
func XBtn_SetOffset(hEle int, x int, y int) int {
	r, _, _ := xBtn_SetOffset.Call(uintptr(hEle), uintptr(x), uintptr(y))
	return int(r)
}

// 按钮_置图标偏移, 设置按钮图标坐标偏移量.
//
// hEle: 元素句柄.
//
// x: 偏移量.
//
// y: 偏移量.
func XBtn_SetOffsetIcon(hEle int, x int, y int) int {
	r, _, _ := xBtn_SetOffsetIcon.Call(uintptr(hEle), uintptr(x), uintptr(y))
	return int(r)
}

// 按钮_置图标间隔, 设置图标与文本间隔大小.
//
// hEle: 元素句柄.
//
// size: 间隔大小.
func XBtn_SetIconSpace(hEle int, size int) int {
	r, _, _ := xBtn_SetIconSpace.Call(uintptr(hEle), uintptr(size))
	return int(r)
}

// 按钮_置文本.
//
// hEle: 元素句柄.
//
// pName: 文本内容.
func XBtn_SetText(hEle int, pName string) int {
	r, _, _ := xBtn_SetText.Call(uintptr(hEle), strPtr(pName))
	return int(r)
}

// 按钮_取文本.
//
// hEle: 元素句柄.
func XBtn_GetText(hEle int) string {
	r, _, _ := xBtn_GetText.Call(uintptr(hEle))
	return UintPtrToString(r)
}

// 按钮_置图标.
//
// hEle: 元素句柄.
//
// hImage: 图片句柄.
func XBtn_SetIcon(hEle int, hImage int) int {
	r, _, _ := xBtn_SetIcon.Call(uintptr(hEle), uintptr(hImage))
	return int(r)
}

// 按钮_置禁用图标.
//
// hEle: 元素句柄.
//
// hImage: 图片句柄.
func XBtn_SetIconDisable(hEle int, hImage int) int {
	r, _, _ := xBtn_SetIconDisable.Call(uintptr(hEle), uintptr(hImage))
	return int(r)
}

// 按钮_取图标, 返回图标句柄.
//
// hEle: 元素句柄.
//
// nType: 图标类型, 0:默认图标, 1:禁用状态图标.
func XBtn_GetIcon(hEle int, nType int) int {
	r, _, _ := xBtn_GetIcon.Call(uintptr(hEle), uintptr(nType))
	return int(r)
}

// 按钮_添加动画帧.
//
// hEle: 元素句柄.
//
// hImage: 图片句柄.
//
// uElapse: 图片帧延时, 单位毫秒.
func XBtn_AddAnimationFrame(hEle int, hImage int, uElapse int) int {
	r, _, _ := xBtn_AddAnimationFrame.Call(uintptr(hEle), uintptr(hImage), uintptr(uElapse))
	return int(r)
}

// 按钮_启用动画, 开始或关闭图片动画的播放.
//
// hEle: 元素句柄.
//
// bEnable: 开始播放动画TRUE, 关闭播放动画FALSE.
//
// bLoopPlay: 是否循环播放.
func XBtn_EnableAnimation(hEle int, bEnable bool, bLoopPlay bool) int {
	r, _, _ := xBtn_EnableAnimation.Call(uintptr(hEle), boolPtr(bEnable), boolPtr(bLoopPlay))
	return int(r)
}

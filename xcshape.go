package xcgui

import "unsafe"

// 形状_创建, 获取所在元素句柄
// hShape: 形状对象句柄.
func XShape_GetParentEle(hShape int) int {
	r, _, _ := xShape_GetParentEle.Call(uintptr(hShape))
	return int(r)
}

// 形状_取窗口句柄, 获取窗口句柄
// hShape: 形状对象句柄.
func XShape_GetHWINDOW(hShape int) int {
	r, _, _ := xShape_GetHWINDOW.Call(uintptr(hShape))
	return int(r)
}

// 形状_取父对象, 获取父对象, 父可能是元素或窗口, 返回父对象
// hShape: 形状对象句柄.
func XShape_GetParent(hShape int) int {
	r, _, _ := xShape_GetParent.Call(uintptr(hShape))
	return int(r)
}

// 形状_移除, 从父UI元素或窗口,和父布局对象中移除
// hShape: 形状对象句柄.
func XShape_RemoveShape(hShape int) int {
	r, _, _ := xShape_RemoveShape.Call(uintptr(hShape))
	return int(r)
}

// 形状_置ID
// hShape: 形状对象句柄.
// nID: ID值.
func XShape_SetID(hShape int, nID int) int {
	r, _, _ := xShape_SetID.Call(uintptr(hShape), uintptr(nID))
	return int(r)
}

// 形状_取ID
// hShape: 形状对象句柄.
func XShape_GetID(hShape int) int {
	r, _, _ := xShape_GetID.Call(uintptr(hShape))
	return int(r)
}

// 形状_置UID
// hShape: 形状对象句柄.
// nUID: UID值.
func XShape_SetUID(hShape int, nUID int) int {
	r, _, _ := xShape_SetUID.Call(uintptr(hShape), uintptr(nUID))
	return int(r)
}

// 形状_取UID
// hShape: 形状对象句柄.
func XShape_GetUID(hShape int) int {
	r, _, _ := xShape_GetUID.Call(uintptr(hShape))
	return int(r)
}

// 形状_置名称, 设置形状对象name
// hShape: 形状对象句柄.
// pName: name值
func XShape_SetName(hShape int, pName string) int {
	r, _, _ := xShape_SetName.Call(uintptr(hShape), strPtr(pName))
	return int(r)
}

// 形状_取名称, 获取形状对象name
// hShape: 形状对象句柄.
func XShape_GetName(hShape int) string {
	r, _, _ := xShape_GetName.Call(uintptr(hShape))
	return uintPtrToString(r)
}

// 形状_取Z序, 获取形状对象Z序, 成功返回索引值, 否则返回 XC_ID_ERROR
// hShape: 形状对象句柄.
func XShape_GetZOrder(hShape int) int {
	r, _, _ := xShape_GetZOrder.Call(uintptr(hShape))
	return int(r)
}

// 形状_重绘, 重绘形状对象
// hShape: 形状对象句柄.
func XShape_Redraw(hShape int) int {
	r, _, _ := xShape_Redraw.Call(uintptr(hShape))
	return int(r)
}

// 形状_取宽度, 获取内容宽度
// hShape: 形状对象句柄.
func XShape_GetWidth(hShape int) int {
	r, _, _ := xShape_GetWidth.Call(uintptr(hShape))
	return int(r)
}

// 形状_取高度, 获取内容高度
// hShape: 形状对象句柄.
func XShape_GetHeight(hShape int) int {
	r, _, _ := xShape_GetHeight.Call(uintptr(hShape))
	return int(r)
}

// 形状_移动位置
// hShape: 形状对象句柄
// x: x坐标
// y: y坐标
func XShape_Move(hShape int, x int, y int) int {
	r, _, _ := xShape_Move.Call(uintptr(hShape), uintptr(x), uintptr(y))
	return int(r)
}

// 形状_取坐标
// hShape: 形状对象句柄.
// pRect: 接收返回坐标.
func XShape_GetRect(hShape int, pRect *RECT) int {
	r, _, _ := xShape_GetRect.Call(uintptr(hShape), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 形状_置坐标
// hShape: 形状对象句柄.
// pRect: 坐标.
func XShape_SetRect(hShape int, pRect *RECT) int {
	r, _, _ := xShape_SetRect.Call(uintptr(hShape), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 形状_置逻辑坐标, 设置元素坐标, 逻辑坐标, 包含滚动视图偏移
// hShape:
// pRect: 坐标.
// bRedraw: 是否重绘.
func XShape_SetRectLogic(hShape int, pRect *RECT, bRedraw bool) bool {
	r, _, _ := xShape_SetRectLogic.Call(uintptr(hShape), uintptr(unsafe.Pointer(pRect)), boolPtr(bRedraw))
	return int(r) != 0
}

// 形状_取逻辑坐标, 获取元素坐标, 逻辑坐标, 包含滚动视图偏移
// hShape: 形状对象句柄.
// pRect: 坐标.
func XShape_GetRectLogic(hShape int, pRect *RECT) int {
	r, _, _ := xShape_GetRectLogic.Call(uintptr(hShape), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 形状_取基于窗口客户区坐标, 基于窗口客户区坐标
// hShape: 形状对象句柄.
// pRect: 坐标.
func XShape_GetWndClientRect(hShape int, pRect *RECT) int {
	r, _, _ := xShape_GetWndClientRect.Call(uintptr(hShape), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 形状_取内容大小 ,仅计算有效内容, 填充父, 权重依赖父级所以无法计算
// hShape: 形状对象句柄.
// pSize: 接收返回内容大小值.
func XShape_GetContentSize(hShape int, pSize *SIZE) int {
	r, _, _ := xShape_GetContentSize.Call(uintptr(hShape), uintptr(unsafe.Pointer(pSize)))
	return int(r)
}

// 形状_显示布局边界, 是否显示布局边界.
// hShape: 形状对象句柄.
// bShow: 是否显示.
func XShape_ShowLayout(hShape int, bShow bool) int {
	r, _, _ := xShape_ShowLayout.Call(uintptr(hShape), boolPtr(bShow))
	return int(r)
}

// 形状_调整布局
// hShape: 形状对象句柄.
func XShape_AdjustLayout(hShape int) int {
	r, _, _ := xShape_AdjustLayout.Call(uintptr(hShape))
	return int(r)
}

// 形状_销毁, 销毁形状对象
// hShape: 形状对象句柄.
func XShape_Destroy(hShape int) int {
	r, _, _ := xShape_Destroy.Call(uintptr(hShape))
	return int(r)
}

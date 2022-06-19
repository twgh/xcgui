package xc

import (
	"github.com/twgh/xcgui/common"
	"unsafe"

	"github.com/twgh/xcgui/xcc"
)

// 背景_创建, 创建背景管理器, 返回背景管理器句柄.
func XBkM_Create() int {
	r, _, _ := xBkM_Create.Call()
	return int(r)
}

// 背景_销毁.
//
// hBkInfoM: 背景管理器句柄.
func XBkM_Destroy(hBkInfoM int) int {
	r, _, _ := xBkM_Destroy.Call(uintptr(hBkInfoM))
	return int(r)
}

// !废弃函数, 保留为了兼容旧版; 请使用 XBkM_SetInfo().
//
// 背景_置内容, 设置背景内容, 返回设置的背景对象数量.
//
// hBkInfoM: 背景管理器句柄.
//
// pText: 背景内容字符串.
func XBkM_SetBkInfo(hBkInfoM int, pText string) int {
	r, _, _ := xBkM_SetBkInfo.Call(uintptr(hBkInfoM), common.StrPtr(pText))
	return int(r)
}

// 背景_添加内容, 添加背景内容, 返回添加的背景对象数量.
//
// hBkInfoM: 背景管理器句柄.
//
// pText: 背景内容字符串.
func XBkM_AddInfo(hBkInfoM int, pText string) int {
	r, _, _ := xBkM_AddInfo.Call(uintptr(hBkInfoM), common.StrPtr(pText))
	return int(r)
}

// 背景_添加边框, 添加背景内容边框.
//
// hBkInfoM: 背景管理器句柄.
//
// nState: 组合状态.
//
// color: ABGR 颜色.
//
// width: 线宽.
//
// id: 背景对象ID, 可忽略(填0).
func XBkM_AddBorder(hBkInfoM int, nState xcc.CombinedState, color, width, id int) int {
	r, _, _ := xBkM_AddBorder.Call(uintptr(hBkInfoM), uintptr(nState), uintptr(color), uintptr(width), uintptr(id))
	return int(r)
}

// 背景_添加填充, 添加背景内容填充.
//
// hBkInfoM: 背景管理器句柄.
//
// nState: 组合状态.
//
// color: ABGR 颜色.
//
// id: 背景对象ID, 可忽略(填0).
func XBkM_AddFill(hBkInfoM int, nState xcc.CombinedState, color, id int) int {
	r, _, _ := xBkM_AddFill.Call(uintptr(hBkInfoM), uintptr(nState), uintptr(color), uintptr(id))
	return int(r)
}

// 背景_添加图片, 添加背景内容图片.
//
// hBkInfoM: 背景管理器句柄.
//
// nState: 组合状态.
//
// hImage: 图片句柄.
//
// id: 背景对象ID, 可忽略(填0).
func XBkM_AddImage(hBkInfoM int, nState xcc.CombinedState, hImage, id int) int {
	r, _, _ := xBkM_AddImage.Call(uintptr(hBkInfoM), uintptr(nState), uintptr(hImage), uintptr(id))
	return int(r)
}

// 背景_取数量, 获取背景内容数量.
//
// hBkInfoM: 背景管理器句柄.
func XBkM_GetCount(hBkInfoM int) int {
	r, _, _ := xBkM_GetCount.Call(uintptr(hBkInfoM))
	return int(r)
}

// 背景_清空, 清空背景内容.
//
// hBkInfoM: 背景管理器句柄.
func XBkM_Clear(hBkInfoM int) int {
	r, _, _ := xBkM_Clear.Call(uintptr(hBkInfoM))
	return int(r)
}

// 背景_绘制, 绘制背景内容.
//
// hBkInfoM: 背景管理器句柄.
//
// nState: 组合状态.
//
// hDraw: 图形绘制句柄.
//
// pRect: 区域坐标.
func XBkM_Draw(hBkInfoM int, nState xcc.CombinedState, hDraw int, pRect *RECT) bool {
	r, _, _ := xBkM_Draw.Call(uintptr(hBkInfoM), uintptr(nState), uintptr(hDraw), uintptr(unsafe.Pointer(pRect)))
	return r != 0
}

// 背景_绘制扩展, 绘制背景内容, 设置条件.
//
// hBkInfoM: 背景管理器句柄.
//
// nState: 组合状态.
//
// hDraw: 图形绘制句柄.
//
// pRect: 区域坐标.
//
// nStateEx: 当(nState)中包含(nStateEx)中的一个或多个状态时有效.
//
// 注解: 例如用来绘制列表项时, nState中包含项的状态(nStateEx)才会绘制, 避免列表项与元素背景叠加.
func XBkM_DrawEx(hBkInfoM int, nState xcc.CombinedState, hDraw int, pRect *RECT, nStateEx xcc.CombinedState) bool {
	r, _, _ := xBkM_DrawEx.Call(uintptr(hBkInfoM), uintptr(nState), uintptr(hDraw), uintptr(unsafe.Pointer(pRect)), uintptr(nStateEx))
	return r != 0
}

// 背景_启用自动销毁, 是否自动销毁.
//
// hBkInfoM: 背景管理器句柄.
//
// bEnable: 是否启用.
func XBkM_EnableAutoDestroy(hBkInfoM int, bEnable bool) int {
	r, _, _ := xBkM_EnableAutoDestroy.Call(uintptr(hBkInfoM), common.BoolPtr(bEnable))
	return int(r)
}

// 背景_增加引用计数.
//
// hBkInfoM: 背景管理器句柄.
func XBkM_AddRef(hBkInfoM int) int {
	r, _, _ := xBkM_AddRef.Call(uintptr(hBkInfoM))
	return int(r)
}

// 背景_释放引用计数.
//
// hBkInfoM: 背景管理器句柄.
func XBkM_Release(hBkInfoM int) int {
	r, _, _ := xBkM_Release.Call(uintptr(hBkInfoM))
	return int(r)
}

// 背景_取引用计数.
//
// hBkInfoM: 背景管理器句柄.
func XBkM_GetRefCount(hBkInfoM int) int {
	r, _, _ := xBkM_GetRefCount.Call(uintptr(hBkInfoM))
	return int(r)
}

// 背景_取引用计数, 设置背景内容, 返回设置的背景对象数量.
//
// hBkInfoM: 背景管理器句柄.
//
// pText: 背景内容字符串.
func XBkM_SetInfo(hBkInfoM int, pText string) int {
	r, _, _ := xBkM_SetInfo.Call(uintptr(hBkInfoM), common.StrPtr(pText))
	return int(r)
}

// 背景_取指定状态文本颜色.
//
// hBkInfoM: 背景管理器句柄.
//
// nState: 组合状态.
//
// color: 接收返回的ABGR 颜色.
func XBkM_GetStateTextColor(hBkInfoM int, nState xcc.CombinedState, color *int) bool {
	r, _, _ := xBkM_GetStateTextColor.Call(uintptr(hBkInfoM), uintptr(nState), uintptr(unsafe.Pointer(color)))
	return r != 0
}

// 背景_取背景对象, 返回BkObj对象句柄.
//
// hBkInfoM: 背景管理器句柄.
//
// id: 背景对象ID.
func XBkM_GetObject(hBkInfoM int, id int) int {
	r, _, _ := xBkM_GetObject.Call(uintptr(hBkInfoM), uintptr(id))
	return int(r)
}

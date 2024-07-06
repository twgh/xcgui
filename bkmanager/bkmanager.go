package bkmanager

import (
	"github.com/twgh/xcgui/objectbase"
	"github.com/twgh/xcgui/res"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// BkManager 背景管理器.
type BkManager struct {
	objectbase.ObjectBase
}

// New 背景_创建, 创建背景管理器.
func New() *BkManager {
	p := &BkManager{}
	p.SetHandle(xc.XBkM_Create())
	return p
}

// NewByHandle 从句柄创建背景管理器对象.
func NewByHandle(handle int) *BkManager {
	p := &BkManager{}
	p.SetHandle(handle)
	return p
}

// NewByName 从name创建背景管理器对象, 失败返回nil.
func NewByName(name string) *BkManager {
	handle := res.GetBkM(name)
	if handle > 0 {
		p := &BkManager{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 背景_销毁.
func (b *BkManager) Destroy() *BkManager {
	xc.XBkM_Destroy(b.Handle)
	return b
}

// 背景_添加内容, 添加背景内容, 返回添加的背景对象数量.
//
//	pText: 背景内容字符串.
func (b *BkManager) AddInfo(pText string) int32 {
	return xc.XBkM_AddInfo(b.Handle, pText)
}

// 背景_添加边框, 添加背景内容边框.
//
//	nState: 组合状态.
//
//	color: ARGB 颜色.
//
//	width: 线宽.
//
//	id: 背景对象ID, 可忽略(填0).
func (b *BkManager) AddBorder(nState xcc.CombinedState, color int, width, id int32) *BkManager {
	xc.XBkM_AddBorder(b.Handle, nState, color, width, id)
	return b
}

// 背景_添加填充, 添加背景内容填充.
//
//	nState: 组合状态.
//
//	color: ARGB 颜色.
//
//	id: 背景对象ID, 可忽略(填0).
func (b *BkManager) AddFill(nState xcc.CombinedState, color int, id int32) *BkManager {
	xc.XBkM_AddFill(b.Handle, nState, color, id)
	return b
}

// 背景_添加图片, 添加背景内容图片.
//
//	nState: 组合状态.
//
//	hImage: 图片句柄.
//
//	id: 背景对象ID, 可忽略(填0).
func (b *BkManager) AddImage(nState xcc.CombinedState, hImage int, id int32) *BkManager {
	xc.XBkM_AddImage(b.Handle, nState, hImage, id)
	return b
}

// 背景_取数量, 获取背景内容数量.
func (b *BkManager) GetCount() int32 {
	return xc.XBkM_GetCount(b.Handle)
}

// 背景_清空, 清空背景内容.
func (b *BkManager) Clear() *BkManager {
	xc.XBkM_Clear(b.Handle)
	return b
}

// 背景_绘制, 绘制背景内容.
//
//	nState: 组合状态.
//
//	hDraw: 图形绘制句柄.
//
//	pRect: 区域坐标.
func (b *BkManager) Draw(nState xcc.CombinedState, hDraw int, pRect *xc.RECT) bool {
	return xc.XBkM_Draw(b.Handle, nState, hDraw, pRect)
}

// 背景_绘制扩展, 绘制背景内容, 设置条件.
//
//	nState: 组合状态.
//
//	hDraw: 图形绘制句柄.
//
//	pRect: 区域坐标.
//
//	nStateEx: 当(nState)中包含(nStateEx)中的一个或多个状态时有效.
//
// 注解: 例如用来绘制列表项时, nState中包含项的状态(nStateEx)才会绘制, 避免列表项与元素背景叠加.
func (b *BkManager) DrawEx(nState xcc.CombinedState, hDraw int, pRect *xc.RECT, nStateEx xcc.CombinedState) bool {
	return xc.XBkM_DrawEx(b.Handle, nState, hDraw, pRect, nStateEx)
}

// 背景_启用自动销毁, 是否自动销毁.
//
//	bEnable: 是否启用.
func (b *BkManager) EnableAutoDestroy(bEnable bool) *BkManager {
	xc.XBkM_EnableAutoDestroy(b.Handle, bEnable)
	return b
}

// 背景_增加引用计数.
func (b *BkManager) AddRef() *BkManager {
	xc.XBkM_AddRef(b.Handle)
	return b
}

// 背景_释放引用计数.
func (b *BkManager) Release() *BkManager {
	xc.XBkM_Release(b.Handle)
	return b
}

// 背景_取引用计数.
func (b *BkManager) GetRefCount() int32 {
	return xc.XBkM_GetRefCount(b.Handle)
}

// 背景_取引用计数, 设置背景内容, 返回设置的背景对象数量.
//
//	pText: 背景内容字符串.
func (b *BkManager) SetInfo(pText string) int32 {
	return xc.XBkM_SetInfo(b.Handle, pText)
}

// 背景_取指定状态文本颜色.
//
//	nState: 组合状态.
//
//	color: 接收返回的ARGB 颜色.
func (b *BkManager) GetStateTextColor(nState xcc.CombinedState, color *int) bool {
	return xc.XBkM_GetStateTextColor(b.Handle, nState, color)
}

// 背景_取背景对象, 返回BkObj对象句柄.
//
//	id: 背景对象ID.
func (b *BkManager) GetObject(id int32) int {
	return xc.XBkM_GetObject(b.Handle, id)
}

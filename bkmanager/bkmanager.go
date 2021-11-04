// 背景管理器.
package bkmanager

import (
	"github.com/twgh/xcgui/objectbase"
	"github.com/twgh/xcgui/res"
	"github.com/twgh/xcgui/xc"
)

// 背景管理器.
type BkManager struct {
	objectbase.ObjectBase
}

// 背景_创建, 创建背景管理器.
func NewBkManager() *BkManager {
	p := &BkManager{}
	p.SetHandle(xc.XBkM_Create())
	return p
}

// 从句柄创建对象.
func NewBkManagerByHandle(handle int) *BkManager {
	p := &BkManager{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func NewBkManagerByName(name string) *BkManager {
	handle := res.GetBkM(name)
	if handle > 0 {
		p := &BkManager{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 背景_销毁.
func (b *BkManager) Destroy() int {
	return xc.XBkM_Destroy(b.Handle)
}

// 背景_置内容, 设置背景内容, 返回设置的背景对象数量.
//
// pText: 背景内容字符串.
func (b *BkManager) SetBkInfo(pText string) int {
	return xc.XBkM_SetBkInfo(b.Handle, pText)
}

// 背景_添加内容, 添加背景内容, 返回添加的背景对象数量.
//
// pText: 背景内容字符串.
func (b *BkManager) AddInfo(pText string) int {
	return xc.XBkM_AddInfo(b.Handle, pText)
}

// 背景_添加边框, 添加背景内容边框.
//
// nState: 组合状态.
//
// color: ABGR颜色.
//
// width: 线宽.
func (b *BkManager) AddBorder(nState int, color int, width int) int {
	return xc.XBkM_AddBorder(b.Handle, nState, color, width)
}

// 背景_添加填充, 添加背景内容填充.
//
// nState: 组合状态, Window_State_Flag_.
//
// color: ABGR颜色.
func (b *BkManager) AddFill(nState int, color int) int {
	return xc.XBkM_AddFill(b.Handle, nState, color)
}

// 背景_添加图片, 添加背景内容图片.
//
// nState: 组合状态.
//
// hImage: 图片句柄.
func (b *BkManager) AddImage(nState int, hImage int) int {
	return xc.XBkM_AddImage(b.Handle, nState, hImage)
}

// 背景_取数量, 获取背景内容数量.
func (b *BkManager) GetCount() int {
	return xc.XBkM_GetCount(b.Handle)
}

// 背景_清空, 清空背景内容.
func (b *BkManager) Clear() int {
	return xc.XBkM_Clear(b.Handle)
}

// 背景_绘制, 绘制背景内容.
//
// nState: 组合状态.
//
// hDraw: 图形绘制句柄.
//
// pRect: 区域坐标.
func (b *BkManager) Draw(nState int, hDraw int, pRect *xc.RECT) bool {
	return xc.XBkM_Draw(b.Handle, nState, hDraw, pRect)
}

// 背景_绘制扩展, 绘制背景内容, 设置条件.
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
func (b *BkManager) DrawEx(nState int, hDraw int, pRect *xc.RECT, nStateEx int) bool {
	return xc.XBkM_DrawEx(b.Handle, nState, hDraw, pRect, nStateEx)
}

// 背景_启用自动销毁, 是否自动销毁.
//
// bEnable: 是否启用.
func (b *BkManager) EnableAutoDestroy(bEnable bool) int {
	return xc.XBkM_EnableAutoDestroy(b.Handle, bEnable)
}

// 背景_增加引用计数.
func (b *BkManager) AddRef() int {
	return xc.XBkM_AddRef(b.Handle)
}

// 背景_释放引用计数.
func (b *BkManager) Release() int {
	return xc.XBkM_Release(b.Handle)
}

// 背景_取引用计数.
func (b *BkManager) GetRefCount() int {
	return xc.XBkM_GetRefCount(b.Handle)
}

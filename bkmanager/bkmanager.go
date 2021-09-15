// 背景管理器
package bkmanager

import (
	"github.com/twgh/xcgui/xc"
)

// 背景管理器
type BkManager struct {
	HBKM int
}

// 背景_创建, 创建背景管理器
func NewBkManager() *BkManager {
	p := &BkManager{
		HBKM: xc.XBkM_Create(),
	}
	return p
}

// 给本类的HBKM赋值
func (b *BkManager) SetHBKM(hBkm int) {
	b.HBKM = hBkm
}

// 背景_销毁
func (b *BkManager) Destroy() int {
	return xc.XBkM_Destroy(b.HBKM)
}

// 背景_置内容, 设置背景内容, 返回设置的背景对象数量
// pText: 背景内容字符串.
func (b *BkManager) SetBkInfo(pText string) int {
	return xc.XBkM_SetBkInfo(b.HBKM, pText)
}

// 背景_添加内容, 添加背景内容, 返回添加的背景对象数量
// pText: 背景内容字符串.
func (b *BkManager) AddInfo(pText string) int {
	return xc.XBkM_AddInfo(b.HBKM, pText)
}

// 背景_添加边框, 添加背景内容边框
// nState: 组合状态
// color: RGB颜色.
// alpha: 透明度.
// width: 线宽.
func (b *BkManager) AddBorder(nState int, color int, alpha uint8, width int) int {
	return xc.XBkM_AddBorder(b.HBKM, nState, color, alpha, width)
}

// 背景_添加填充, 添加背景内容填充
// nState: 组合状态, Window_State_Flag_
// color: RGB颜色.
// alpha: 透明度.
func (b *BkManager) AddFill(nState int, color int, alpha uint8) int {
	return xc.XBkM_AddFill(b.HBKM, nState, color, alpha)
}

// 背景_添加图片, 添加背景内容图片
// nState: 组合状态
// hImage: 图片句柄.
func (b *BkManager) AddImage(nState int, hImage int) int {
	return xc.XBkM_AddImage(b.HBKM, nState, hImage)
}

// 背景_取数量, 获取背景内容数量
func (b *BkManager) GetCount() int {
	return xc.XBkM_GetCount(b.HBKM)
}

// 背景_清空, 清空背景内容
func (b *BkManager) Clear() int {
	return xc.XBkM_Clear(b.HBKM)
}

// 背景_绘制, 绘制背景内容
// nState: 组合状态
// hDraw: 图形绘制句柄.
// pRect: 区域坐标.
func (b *BkManager) Draw(nState int, hDraw int, pRect *xc.RECT) bool {
	return xc.XBkM_Draw(b.HBKM, nState, hDraw, pRect)
}

// 背景_绘制扩展, 绘制背景内容, 设置条件
// nState: 组合状态.
// hDraw: 图形绘制句柄.
// pRect: 区域坐标.
// nStateEx: 当(nState)中包含(nStateEx)中的一个或多个状态时有效.
// 注解: 例如用来绘制列表项时, nState中包含项的状态(nStateEx)才会绘制, 避免列表项与元素背景叠加
func (b *BkManager) DrawEx(nState int, hDraw int, pRect *xc.RECT, nStateEx int) bool {
	return xc.XBkM_DrawEx(b.HBKM, nState, hDraw, pRect, nStateEx)
}

// 背景_启用自动销毁, 是否自动销毁
// bEnable: 是否启用
func (b *BkManager) EnableAutoDestroy(bEnable bool) int {
	return xc.XBkM_EnableAutoDestroy(b.HBKM, bEnable)
}

// 背景_增加引用计数
func (b *BkManager) AddRef() int {
	return xc.XBkM_AddRef(b.HBKM)
}

// 背景_释放引用计数
func (b *BkManager) Release() int {
	return xc.XBkM_Release(b.HBKM)
}

// 背景_取引用计数
func (b *BkManager) GetRefCount() int {
	return xc.XBkM_GetRefCount(b.HBKM)
}

// SVG矢量图形.
package svg

import (
	"github.com/twgh/xcgui/objectbase"
	"github.com/twgh/xcgui/xc"
)

// SVG矢量图形.
type Svg struct {
	objectbase.ObjectBase
}

// SVG_加载从文件, 返回Svg对象.
//
// pFileName: 文件名.
func NewSvg_LoadFile(pFileName string) *Svg {
	p := &Svg{}
	p.SetHandle(xc.XSvg_LoadFile(pFileName))
	return p
}

// SVG_加载从字符串, 返回Svg对象.
//
// pString: 字符串指针.
//
// nLength: 字符串长度.
func NewSvg_LoadString(pString string, nLength int) *Svg {
	p := &Svg{}
	p.SetHandle(xc.XSvg_LoadString(pString, nLength))
	return p
}

// SVG_加载从ZIP, 返回Svg对象.
//
// pZipFileName: zip文件名.
//
// pFileName: svg文件名.
//
// pPassword: zip密码.
func NewSvg_LoadZip(pZipFileName, pFileName, pPassword string) *Svg {
	p := &Svg{}
	p.SetHandle(xc.XSvg_LoadZip(pZipFileName, pFileName, pPassword))
	return p
}

// SVG_加载从资源, 返回Svg对象.
//
// id: 资源ID.
//
// pType: 资源类型.在rc资源文件中.
//
// hModule: 从指定模块加载.
func NewSvg_LoadRes(id int, pType string, hModule int) *Svg {
	p := &Svg{}
	p.SetHandle(xc.XSvg_LoadRes(id, pType, hModule))
	return p
}

// SVG_置大小.
//
// nWidth: 宽度.
//
// nHeight: 高度.
func (s *Svg) SetSize(nWidth, nHeight int) int {
	return xc.XSvg_SetSize(s.Handle, nWidth, nHeight)
}

// SVG_取大小.
//
// pWidth: 接收返回宽度.
//
// pHeight: 接收返回高度.
func (s *Svg) GetSize(pWidth, pHeight *int) int {
	return xc.XSvg_GetSize(s.Handle, pWidth, pHeight)
}

// SVG_取宽度.
func (s *Svg) GetWidth() int {
	return xc.XSvg_GetWidth(s.Handle)
}

// SVG_取高度.
func (s *Svg) GetHeight() int {
	return xc.XSvg_GetHeight(s.Handle)
}

// SVG_置偏移.
//
// x: x轴偏移.
//
// y: y轴偏移.
func (s *Svg) SetOffset(x int, y int) int {
	return xc.XSvg_SetOffset(s.Handle, x, y)
}

// SVG_取偏移.
//
// pX: x轴偏移.
//
// pY: y轴偏移.
func (s *Svg) GetOffset(pX, pY *int) int {
	return xc.XSvg_GetOffset(s.Handle, pX, pY)
}

// SVG_取视图框.
//
// pViewBox: 接收返回视图框.
func (s *Svg) GetViewBox(pViewBox *xc.RECT) int {
	return xc.XSvg_GetViewBox(s.Handle, pViewBox)
}

// SVG_启用自动销毁.
//
// bEnable: 是否自动销毁.
func (s *Svg) EnableAutoDestroy(bEnable bool) int {
	return xc.XSvg_EnableAutoDestroy(s.Handle, bEnable)
}

// SVG_增加引用计数.
func (s *Svg) AddRef() int {
	return xc.XSvg_AddRef(s.Handle)
}

// SVG_释放引用计数.
func (s *Svg) Release() int {
	return xc.XSvg_Release(s.Handle)
}

// SVG_取引用计数.
func (s *Svg) GetRefCount() int {
	return xc.XSvg_GetRefCount(s.Handle)
}

// SVG_销毁.
func (s *Svg) Destroy() int {
	return xc.XSvg_Destroy(s.Handle)
}

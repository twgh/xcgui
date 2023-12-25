package imagex

import (
	"github.com/twgh/xcgui/objectbase"
	"github.com/twgh/xcgui/res"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// Image 图片操作.
type Image struct {
	objectbase.ObjectBase
}

// 图片_加载从图片源.
//
// hImageSrc: 图片源句柄.
func NewBySrc(hImageSrc int) *Image {
	p := &Image{}
	p.SetHandle(xc.XImage_LoadSrc(hImageSrc))
	return p
}

// 图片_加载从文件.
//
// pFileName: 图片文件.
func NewByFile(pFileName string) *Image {
	p := &Image{}
	p.SetHandle(xc.XImage_LoadFile(pFileName))
	return p
}

// 图片_加载从文件自适应, 加载图片从文件, 自适应图片.
//
// pFileName: 图片文件.
//
// leftSize: 坐标.
//
// topSize: 坐标.
//
// rightSize: 坐标.
//
// bottomSize: 坐标.
func NewByFileAdaptive(pFileName string, leftSize, topSize, rightSize, bottomSize int32) *Image {
	p := &Image{}
	p.SetHandle(xc.XImage_LoadFileAdaptive(pFileName, leftSize, topSize, rightSize, bottomSize))
	return p
}

// 图片_加载从文件指定区域, 加载图片, 指定区位置及大小.
//
// pFileName: 图片文件.
//
// x: 坐标.
//
// y: 坐标.
//
// cx: 宽度.
//
// cy: 高度.
func NewByFileRect(pFileName string, x, y, cx, cy int32) *Image {
	p := &Image{}
	p.SetHandle(xc.XImage_LoadFileRect(pFileName, x, y, cx, cy))
	return p
}

// 图片_加载从资源自适应, 加载图片从资源, 自适应图片.
//
// id: 资源ID.
//
// pType: 资源类型.
//
// leftSize: 坐标.
//
// topSize: 坐标.
//
// rightSize: 坐标.
//
// bottomSize: 坐标.
//
// hModule:	从指定模块加载, 例如:DLL, EXE; 如果为空, 从当前EXE加载.
func NewByResAdaptive(id int32, pType string, leftSize, topSize, rightSize, bottomSize int32, hModule uintptr) *Image {
	p := &Image{}
	p.SetHandle(xc.XImage_LoadResAdaptive(id, pType, leftSize, topSize, rightSize, bottomSize, hModule))
	return p
}

// 图片_加载从资源.
//
// id: 资源ID.
//
// pType: 资源类型.
//
// bStretch: 是否拉伸图片.
//
// hModule:	从指定模块加载, 例如:DLL, EXE; 如果为空, 从当前EXE加载.
func NewByRes(id int32, pType string, bStretch bool, hModule uintptr) *Image {
	p := &Image{}
	p.SetHandle(xc.XImage_LoadRes(id, pType, bStretch, hModule))
	return p
}

// 图片_加载从ZIP, 加载图片从ZIP压缩包.
//
// pZipFileName: ZIP压缩包文件名.
//
// pFileName: 图片文件名.
//
// pPassword: ZIP压缩包密码.
func NewByZip(pZipFileName string, pFileName string, pPassword string) *Image {
	p := &Image{}
	p.SetHandle(xc.XImage_LoadZip(pZipFileName, pFileName, pPassword))
	return p
}

// 图片_加载从资源ZIP.
//
// id: RC资源ID.
//
// pFileName: 图片文件名.
//
// pPassword: ZIP压缩包密码.
//
// hModule: 模块句柄, 可填0.
func NewByZipRes(id int32, pFileName string, pPassword string, hModule uintptr) *Image {
	p := &Image{}
	p.SetHandle(xc.XImage_LoadZipRes(id, pFileName, pPassword, hModule))
	return p
}

// 图片_加载从ZIP自适应, 加载图片从ZIP压缩包, 自适应图片.
//
// pZipFileName: ZIP压缩包文件名.
//
// pFileName: 图片文件名.
//
// pPassword: ZIP压缩包密码.
//
// x1: 坐标.
//
// x2: 坐标.
//
// y1: 坐标.
//
// y2: 坐标.
func NewByZipAdaptive(pZipFileName string, pFileName string, pPassword string, x1, x2, y1, y2 int32) *Image {
	p := &Image{}
	p.SetHandle(xc.XImage_LoadZipAdaptive(pZipFileName, pFileName, pPassword, x1, x2, y1, y2))
	return p
}

// 图片_加载从ZIP指定区域, 加载ZIP图片, 指定区位置及大小.
//
// pZipFileName: ZIP文件.
//
// pFileName: 图片名称.
//
// pPassword: 密码.
//
// x: 坐标.
//
// y: 坐标.
//
// cx: 宽度.
//
// cy: 高度.
func NewByZipRect(pZipFileName string, pFileName string, pPassword string, x, y, cx, cy int32) *Image {
	p := &Image{}
	p.SetHandle(xc.XImage_LoadZipRect(pZipFileName, pFileName, pPassword, x, y, cx, cy))
	return p
}

// 图片_加载从内存ZIP.
//
// data: 图片数据.
//
// pFileName: 图片名称.
//
// pPassword: zip压缩包密码.
func NewByZipMem(data []byte, pFileName string, pPassword string) *Image {
	p := &Image{}
	p.SetHandle(xc.XImage_LoadZipMem(data, pFileName, pPassword))
	return p
}

// 图片_加载从内存, 加载流图片.
//
// pBuffer: 图片数据.
func NewByMem(pBuffer []byte) *Image {
	p := &Image{}
	p.SetHandle(xc.XImage_LoadMemory(pBuffer))
	return p
}

// 图片_加载从内存指定区域, 加载流图片, 指定区位置及大小.
//
// pBuffer: 图片数据.
//
// x: 坐标.
//
// y: 坐标.
//
// cx: 宽度.
//
// cy: 高度.
func NewByMemRect(pBuffer []byte, x, y, cx, cy int32) *Image {
	p := &Image{}
	p.SetHandle(xc.XImage_LoadMemoryRect(pBuffer, x, y, cx, cy))
	return p
}

// 图片_加载从内存自适应, 加载流图片压缩包, 自适应图片(九宫格).
//
// pBuffer: 图片数据.
//
// leftSize: 坐标.
//
// topSize: 坐标.
//
// rightSize: 坐标.
//
// bottomSize: 坐标.
func NewByMemAdaptive(pBuffer []byte, leftSize, topSize, rightSize, bottomSize int32) *Image {
	p := &Image{}
	p.SetHandle(xc.XImage_LoadMemoryAdaptive(pBuffer, leftSize, topSize, rightSize, bottomSize))
	return p
}

// 图片_加载从Image, 加载图片从GDI+的Image对象.
//
// pImage: GDI图片对象指针Image*.
func NewByImage(pImage uintptr) *Image {
	p := &Image{}
	p.SetHandle(xc.XImage_LoadFromImage(pImage))
	return p
}

// 图片_加载文件图标, 加载文件图标, 从一个EXE文件或DLL文件或图标文件; 例如:*.exe文件的图标.
//
// pFileName: 文件名.
func NewByExtractIcon(pFileName string) *Image {
	p := &Image{}
	p.SetHandle(xc.XImage_LoadFromExtractIcon(pFileName))
	return p
}

// 图片_加载从HICON, 创建一个炫彩图片句柄, 从一个现有的图标句柄HICON.
//
// hIcon: 图标句柄.
func NewByHICON(hIcon uintptr) *Image {
	p := &Image{}
	p.SetHandle(xc.XImage_LoadFromHICON(hIcon))
	return p
}

// 图片_加载从HBITMAP, 创建一个炫彩图片句柄, 从一个现有的位图句柄HBITMAP.
//
// hBitmap: 位图句柄.
func NewByHBITMAP(hBitmap uintptr) *Image {
	p := &Image{}
	p.SetHandle(xc.XImage_LoadFromHBITMAP(hBitmap))
	return p
}

// 图片_加载从SVG.
//
// hSvg: SVG句柄.
func NewBySvg(hSvg int) *Image {
	p := &Image{}
	p.SetHandle(xc.XImage_LoadSvg(hSvg))
	return p
}

// 图片_加载从SVG文件.
//
// pFileName: 文件名.
func NewBySvgFile(pFileName string) *Image {
	p := &Image{}
	p.SetHandle(xc.XImage_LoadSvgFile(pFileName))
	return p
}

// 图片_加载从SVG字符串.
//
// pString: 字符串.
func NewBySvgString(pString string) *Image {
	p := &Image{}
	p.SetHandle(xc.XImage_LoadSvgString(pString))
	return p
}

// 图片_加载从SVG字符串W.
//
// pString: 字符串.
func NewBySvgStringW(pString string) *Image {
	p := &Image{}
	p.SetHandle(xc.XImage_LoadSvgStringW(pString))
	return p
}

// 图片_加载从SVG字符串UTF8, 更推荐使用 imagex.NewBySvgStringW.
//
// pString: 字符串.
func NewBySvgStringUtf8(pString string) *Image {
	p := &Image{}
	p.SetHandle(xc.XImage_LoadSvgStringUtf8(pString))
	return p
}

// 从句柄创建对象.
func NewByHandle(handle int) *Image {
	p := &Image{}
	p.SetHandle(handle)
	return p
}

// 根据资源文件中的name创建对象, 失败返回nil.
//
// pName: 资源名称.
func NewByName(name string) *Image {
	handle := res.GetImage(name)
	if handle > 0 {
		p := &Image{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从指定的资源文件中, 根据name创建对象, 失败返回nil.
//
// pFileName: 资源文件名.
//
// pName: 资源名称.
func NewByNameEx(fileName, name string) *Image {
	handle := res.GetImageEx(fileName, name)
	if handle > 0 {
		p := &Image{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 图片_取SVG, 返回SVG句柄.
//
// hImage: 图片句柄.
func (i *Image) GetSvg() int {
	return xc.XImage_GetSvg(i.Handle)
}

// 图片_判断缩放, 是否为拉伸图片句柄.
func (i *Image) IsStretch() bool {
	return xc.XImage_IsStretch(i.Handle)
}

// 图片_判断自适应, 是否为自适应图片句柄.
func (i *Image) IsAdaptive() bool {
	return xc.XImage_IsAdaptive(i.Handle)
}

// 图片_判断平铺, 是否为平铺图片.
func (i *Image) IsTile() bool {
	return xc.XImage_IsTile(i.Handle)
}

// 图片_置绘制类型, 设置图片绘制类型.
//
// nType: 图片绘制类型, Image_Draw_Type_.
func (i *Image) SetDrawType(nType xcc.Image_Draw_Type_) bool {
	return xc.XImage_SetDrawType(i.Handle, nType)
}

// 图片_置绘制类型自适应, 设置图片自适应(九宫格).
//
// leftSize: 坐标.
//
// topSize: 坐标.
//
// rightSize: 坐标.
//
// bottomSize: 坐标.
func (i *Image) SetDrawTypeAdaptive(leftSize, topSize, rightSize, bottomSize int32) bool {
	return xc.XImage_SetDrawTypeAdaptive(i.Handle, leftSize, topSize, rightSize, bottomSize)
}

// 图片_置透明色, 指定图片透明颜色.
//
// color: ABGR 颜色.
func (i *Image) SetTranColor(color int) {
	xc.XImage_SetTranColor(i.Handle, color)
}

// 图片_置透明色扩展, 指定图片透明颜色及透明度.
//
// color: ABGR 颜色.
//
// tranColor: 透明色的透明度.
func (i *Image) SetTranColorEx(color int, tranColor byte) {
	xc.XImage_SetTranColorEx(i.Handle, color, tranColor)
}

// 图片_置旋转角度, 设置旋转角度, 返回先前角度.
//
// fAngle: 选择角度.
func (i *Image) SetRotateAngle(fAngle float32) float32 {
	return xc.XImage_SetRotateAngle(i.Handle, fAngle)
}

// 图片_置等分.
//
// nCount: 等分数量.
//
// iIndex: 索引.
func (i *Image) SetSplitEqual(nCount, iIndex int32) {
	xc.XImage_SetSplitEqual(i.Handle, nCount, iIndex)
}

// 图片_启用透明色, 启用或关闭图片透明色.
//
// bEnable: 启用TRUE.
func (i *Image) EnableTranColor(bEnable bool) {
	xc.XImage_EnableTranColor(i.Handle, bEnable)
}

// 图片_启用自动销毁, 启用或关闭自动销毁, 当与UI元素关联时有效.
//
// bEnable: 启用自动销毁TRUE.
func (i *Image) EnableAutoDestroy(bEnable bool) {
	xc.XImage_EnableAutoDestroy(i.Handle, bEnable)
}

// 图片_启用居中, 启用或关闭图片居中显示，默认属性图片有效.
//
// bCenter: 是否居中显示.
func (i *Image) EnableCenter(bCenter bool) {
	xc.XImage_EnableCenter(i.Handle, bCenter)
}

// 图片_判断居中, 判断图片是否居中显示.
func (i *Image) IsCenter() bool {
	return xc.XImage_IsCenter(i.Handle)
}

// 图片_取绘制类型, 获取图片绘制类型, 返回: xcc.Image_Draw_Type_.
func (i *Image) GetDrawType() xcc.Image_Draw_Type_ {
	return xc.XImage_GetDrawType(i.Handle)
}

// 图片_取宽度.
func (i *Image) GetWidth() int32 {
	return xc.XImage_GetWidth(i.Handle)
}

// 图片_取高度.
func (i *Image) GetHeight() int32 {
	return xc.XImage_GetHeight(i.Handle)
}

// 图片_取图片源.
func (i *Image) GetImageSrc() int {
	return xc.XImage_GetImageSrc(i.Handle)
}

// 图片_增加引用计数.
func (i *Image) AddRef() {
	xc.XImage_AddRef(i.Handle)
}

// 图片_释放引用计数, 释放引用计数, 当引用计数为0时, 自动销毁.
func (i *Image) Release() {
	xc.XImage_Release(i.Handle)
}

// 图片_取引用计数.
func (i *Image) GetRefCount() int32 {
	return xc.XImage_GetRefCount(i.Handle)
}

// 图片_销毁, 强制销毁图片, 谨慎使用, 建议使用 XImage_Release() 释放.
func (i *Image) Destroy() {
	xc.XImage_Destroy(i.Handle)
}

// 图片_置缩放大小, 启用缩放属性后有效, 值大于0有效.
//
// width: 宽度.
//
// height: 高度.
func (i *Image) SetScaleSize(width, height int32) {
	xc.XImage_SetScaleSize(i.Handle, width, height)
}

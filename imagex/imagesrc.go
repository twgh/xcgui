package imagex

import (
	"image"

	"github.com/twgh/xcgui/objectbase"
	"github.com/twgh/xcgui/xc"
)

// ImageSrc 图片源.
//   - 提供图片加载, 图片共享, 无渲染属性, 仅作为内存中的资源共享, 已保证每个图片的唯一性, 避免重复浪费资源.
//   - 主要支持: 加载图片文件格式: bmp,jpg,png,gif,ico. 从文件加载, 从程序资源加载, 从压缩包中加载, 自适应图片, 平铺, 透明色支持.
type ImageSrc struct {
	objectbase.ObjectBase
}

// NewImageSrcByFile 图片源_加载从文件, 失败返回 nil.
//
// fileName: 图片文件.
func NewImageSrcByFile(fileName string) *ImageSrc {
	return NewImageSrcByHandle(xc.XImgSrc_LoadFile(fileName))
}

// NewImageSrcByFileRect 图片源_加载从文件指定区域, 加载图片, 指定区位置及大小, 失败返回 nil.
//
// fileName: 图片文件.
//
// x: 坐标.
//
// y: 坐标.
//
// cx: 宽度.
//
// cy: 高度.
func NewImageSrcByFileRect(fileName string, x, y, cx, cy int32) *ImageSrc {
	return NewImageSrcByHandle(xc.XImgSrc_LoadFileRect(fileName, x, y, cx, cy))
}

// NewImageSrcByRes 图片源_加载从资源, 失败返回 nil.
//
// id: 资源ID.
//
// Type: 资源类型, 在rc资源文件中,资源的类型, 例如:xcgui.rc, 用记事本打开可以看见资源类型; 例如:BITMAP, PNG; 参见MSDN.
//
// hModule: 从指定模块加载, 例如:DLL, EXE; 如果为空, 从当前EXE加载.
func NewImageSrcByRes(id int32, Type string, hModule ...uintptr) *ImageSrc {
	return NewImageSrcByHandle(xc.XImgSrc_LoadRes(id, Type, hModule...))
}

// NewImageSrcByZip 图片源_加载从ZIP, 加载图片从ZIP压缩包, 失败返回 nil.
//
// zipFileName: ZIP压缩包文件名.
//
// fileName: 图片文件名.
//
// password: ZIP压缩包密码, 不填默认为空.
func NewImageSrcByZip(zipFileName, fileName string, password ...string) *ImageSrc {
	return NewImageSrcByHandle(xc.XImgSrc_LoadZip(zipFileName, fileName, password...))
}

// NewImageSrcByZipRect 图片源_加载从ZIP指定区域, 加载ZIP图片, 指定区位置及大小, 失败返回 nil.
//
// zipFileName: ZIP文件.
//
// fileName: 图片名称.
//
// password: 密码.
//
// x: 坐标.
//
// y: 坐标.
//
// cx: 宽度.
//
// cy: 高度.
func NewImageSrcByZipRect(zipFileName, fileName, password string, x, y, cx, cy int32) *ImageSrc {
	return NewImageSrcByHandle(xc.XImgSrc_LoadZipRect(zipFileName, fileName, password, x, y, cx, cy))
}

// NewImageSrcByZipMem 图片源_加载从内存ZIP, 失败返回 nil.
//
// data: 图片数据.
//
// fileName: 图片名称.
//
// password: zip压缩包密码, 不填默认为空.
func NewImageSrcByZipMem(data []byte, fileName string, password ...string) *ImageSrc {
	return NewImageSrcByHandle(xc.XImgSrc_LoadZipMem(data, fileName, password...))
}

// NewImageSrcByMem 图片源_加载从内存, 加载流图片, 失败返回 nil.
//
// buffer: 图片数据.
func NewImageSrcByMem(buffer []byte) *ImageSrc {
	return NewImageSrcByHandle(xc.XImgSrc_LoadMemory(buffer))
}

// NewImageSrcByMemRect 图片源_加载从内存指定区域, 加载流图片, 指定区位置及大小, 失败返回 nil.
//
// buffer: 图片数据.
//
// x: 坐标.
//
// y: 坐标.
//
// cx: 宽度.
//
// cy: 高度.
func NewImageSrcByMemRect(buffer []byte, x, y, cx, cy int32) *ImageSrc {
	return NewImageSrcByHandle(xc.XImgSrc_LoadMemoryRect(buffer, x, y, cx, cy))
}

// NewImageSrcByImage 图片源_加载从Image, 加载图片从GDI+的Image对象, 失败返回 nil.
//
// pImage: GDI图片对象指针 Bitmap*.
func NewImageSrcByImage(pImage uintptr) *ImageSrc {
	return NewImageSrcByHandle(xc.XImgSrc_LoadFromImage(pImage))
}

// NewImageSrcByExtractIcon 图片源_加载文件图标, 加载文件图标, 从一个EXE文件或DLL文件或图标文件; 例如:*.exe文件的图标, 失败返回 nil.
//
// fileName: 文件名.
func NewImageSrcByExtractIcon(fileName string) *ImageSrc {
	return NewImageSrcByHandle(xc.XImgSrc_LoadFromExtractIcon(fileName))
}

// NewImageSrcByHICON 图片源_加载从HICON, 创建一个炫彩图片句柄, 从一个现有的图标句柄HICON, 失败返回 nil.
//
// hIcon: 图标句柄.
func NewImageSrcByHICON(hIcon uintptr) *ImageSrc {
	return NewImageSrcByHandle(xc.XImgSrc_LoadFromHICON(hIcon))
}

// NewImageSrcByHBITMAP 图片源_加载从HBITMAP, 创建一个炫彩图片句柄, 从一个现有的位图句柄HBITMAP, 失败返回 nil.
//
// hBitmap: 位图句柄.
func NewImageSrcByHBITMAP(hBitmap uintptr) *ImageSrc {
	return NewImageSrcByHandle(xc.XImgSrc_LoadFromHBITMAP(hBitmap))
}

// NewImageSrcByData 图片源_加载从数据, 从内存加载原始像素数据(BGRA32位格式), 1个像素占4个字节,分别是(B,G,R,A), 失败返回 nil.
//
// data: 数据(BGRA32位格式).
//
// width: 图片宽度.
//
// height: 图片高度.
func NewImageSrcByData(data []byte, width, height int32) *ImageSrc {
	return NewImageSrcByHandle(xc.XImgSrc_LoadFromData(data, width, height))
}

// NewImageSrcByDataRGBA 图片源_加载从数据RGBA, 从内存加载原始像素数据(RGBA32位格式), 1个像素占4个字节,分别是(R,G,B,A), 失败返回 nil.
//
// data: 数据(RGBA32位格式).
//
// width: 图片宽度.
//
// height: 图片高度.
func NewImageSrcByDataRGBA(data []byte, width, height int32) *ImageSrc {
	return NewImageSrcByHandle(xc.XImgSrc_LoadFromDataRGBA(data, width, height))
}

// NewImageSrcByGoImage 图片源_加载从GoImage, 失败返回 nil.
//
// img: *image.RGBA.
func NewImageSrcByGoImage(img *image.RGBA) *ImageSrc {
	return NewImageSrcByHandle(xc.XImgSrc_LoadFromGoImage(img))
}

// NewImageSrcByHandle 从句柄创建对象, 失败返回 nil.
//
// handle: 图片句柄.
func NewImageSrcByHandle(handle int) *ImageSrc {
	if handle < 1 {
		return nil
	}
	p := &ImageSrc{}
	p.SetHandle(handle)
	return p
}

// EnableAutoDestroy 图片源_启用自动销毁, 启用或关闭自动销毁, 当与UI元素关联时有效.
//
// bEnable: 是否启用, 不填默认为 true.
func (i *ImageSrc) EnableAutoDestroy(bEnable ...bool) *ImageSrc {
	xc.XImgSrc_EnableAutoDestroy(i.Handle, bEnable...)
	return i
}

// GetWidth 图片源_取宽度.
func (i *ImageSrc) GetWidth() int32 {
	return xc.XImgSrc_GetWidth(i.Handle)
}

// GetHeight 图片源_取高度.
func (i *ImageSrc) GetHeight() int32 {
	return xc.XImgSrc_GetHeight(i.Handle)
}

// GetFile 图片源_取文件名.
func (i *ImageSrc) GetFile() string {
	return xc.XImgSrc_GetFile(i.Handle)
}

// AddRef 图片源_增加引用计数.
func (i *ImageSrc) AddRef() *ImageSrc {
	xc.XImgSrc_AddRef(i.Handle)
	return i
}

// Release 图片源_释放引用计数, 当引用计数为0时, 自动销毁.
func (i *ImageSrc) Release() *ImageSrc {
	xc.XImgSrc_Release(i.Handle)
	return i
}

// GetRefCount 图片源_取引用计数.
func (i *ImageSrc) GetRefCount() int32 {
	return xc.XImgSrc_GetRefCount(i.Handle)
}

// Destroy 图片源_销毁, 强制销毁图片, 谨慎使用, 建议使用 Release() 释放.
func (i *ImageSrc) Destroy() *ImageSrc {
	xc.XImgSrc_Destroy(i.Handle)
	return i
}

// ModifyData 图片源_修改数据, 修改图片内存数据(BGRA32位格式), 1个像素占4个字节,分别是(B,G,R,A).
//
// data: 数据(BGRA32位格式).
//
// width: 图片宽度.
//
// height: 图片高度.
func (i *ImageSrc) ModifyData(data []byte, width, height int32) bool {
	return xc.XImgSrc_ModifyData(i.Handle, data, width, height)
}

// ModifyDataRGBA 图片源_修改数据RGBA, 修改图片内存数据(RGBA32位格式), 1个像素占4个字节,分别是(R,G,B,A).
//
// data: 数据(RGBA32位格式).
//
// width: 图片宽度.
//
// height: 图片高度.
func (i *ImageSrc) ModifyDataRGBA(data []byte, width, height int32) bool {
	return xc.XImgSrc_ModifyDataRGBA(i.Handle, data, width, height)
}

// ModifyDataGoImage 图片源_修改数据从GoImage, 修改图片内存数据.
//
// img: *image.RGBA.
func (i *ImageSrc) ModifyDataGoImage(img *image.RGBA) bool {
	return xc.XImgSrc_ModifyDataGoImage(i.Handle, img)
}

// GetWicBitMap 图片源_取WicBitmap. 获取WIC位图指针, 返回 IWICBitmapSource*.
func (i *ImageSrc) GetWicBitMap() uintptr {
	return xc.XImgSrc_GetWicBitMap(i.Handle)
}

// GetGdiplusBitmap 图片源_取GdiplusBitmap. 获取GDI+位图指针, 返回 Gdiplus::Bitmap*.
func (i *ImageSrc) GetGdiplusBitmap() uintptr {
	return xc.XImgSrc_GetGdiplusBitmap(i.Handle)
}

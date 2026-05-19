package xc

import (
	"image"

	"github.com/twgh/xcgui/common"
)

// 图片源_加载从文件, 返回图片句柄.
//
// fileName: 图片文件.
func XImgSrc_LoadFile(fileName string) int {
	r, _, _ := xImgSrc_LoadFile.Call(common.StrPtr(fileName))
	return int(r)
}

// 图片源_加载从文件指定区域, 加载图片, 指定区位置及大小, 返回图片句柄.
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
func XImgSrc_LoadFileRect(fileName string, x, y, cx, cy int32) int {
	r, _, _ := xImgSrc_LoadFileRect.Call(common.StrPtr(fileName), uintptr(x), uintptr(y), uintptr(cx), uintptr(cy))
	return int(r)
}

// 图片源_加载从资源, 返回图片句柄.
//
// id: 资源ID.
//
// Type: 资源类型, 在rc资源文件中,资源的类型, 例如:xcgui.rc, 用记事本打开可以看见资源类型; 例如:BITMAP, PNG; 参见MSDN.
//
// hModule: 从指定模块加载, 例如:DLL, EXE; 如果为空, 从当前EXE加载.
func XImgSrc_LoadRes(id int32, Type string, hModule ...uintptr) int {
	hMod := uintptr(0)
	if len(hModule) > 0 {
		hMod = hModule[0]
	}
	r, _, _ := xImgSrc_LoadRes.Call(uintptr(id), common.StrPtr(Type), hMod)
	return int(r)
}

// 图片源_加载从ZIP, 加载图片从ZIP压缩包, 返回图片句柄.
//
// zipFileName: ZIP压缩包文件名.
//
// fileName: 图片文件名.
//
// password: ZIP压缩包密码, 不填默认为空.
func XImgSrc_LoadZip(zipFileName, fileName string, password ...string) int {
	pwd := ""
	if len(password) > 0 {
		pwd = password[0]
	}
	r, _, _ := xImgSrc_LoadZip.Call(common.StrPtr(zipFileName), common.StrPtr(fileName), common.StrPtr(pwd))
	return int(r)
}

// 图片源_加载从ZIP指定区域, 加载ZIP图片, 指定区位置及大小, 返回图片句柄.
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
func XImgSrc_LoadZipRect(zipFileName, fileName, password string, x, y, cx, cy int32) int {
	r, _, _ := xImgSrc_LoadZipRect.Call(common.StrPtr(zipFileName), common.StrPtr(fileName), common.StrPtr(password), uintptr(x), uintptr(y), uintptr(cx), uintptr(cy))
	return int(r)
}

// 图片源_加载从内存ZIP, 返回图片句柄.
//
// data: 图片数据.
//
// fileName: 图片名称.
//
// password: zip压缩包密码, 不填默认为空.
func XImgSrc_LoadZipMem(data []byte, fileName string, password ...string) int {
	pwd := ""
	if len(password) > 0 {
		pwd = password[0]
	}
	r, _, _ := xImgSrc_LoadZipMem.Call(common.ByteSliceDataPtr(&data), uintptr(len(data)), common.StrPtr(fileName), common.StrPtr(pwd))
	return int(r)
}

// 图片源_加载从内存, 加载流图片, 返回图片句柄.
//
// buffer: 图片数据.
func XImgSrc_LoadMemory(buffer []byte) int {
	r, _, _ := xImgSrc_LoadMemory.Call(common.ByteSliceDataPtr(&buffer), uintptr(len(buffer)))
	return int(r)
}

// 图片源_加载从内存指定区域, 加载流图片, 指定区位置及大小, 返回图片句柄.
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
func XImgSrc_LoadMemoryRect(buffer []byte, x, y, cx, cy int32) int {
	r, _, _ := xImgSrc_LoadMemoryRect.Call(common.ByteSliceDataPtr(&buffer), uintptr(len(buffer)), uintptr(x), uintptr(y), uintptr(cx), uintptr(cy))
	return int(r)
}

// 图片源_加载从Image, 加载图片从GDI+的Image对象, 返回图片句柄.
//
// pImage: GDI图片对象指针 Bitmap*.
func XImgSrc_LoadFromImage(pImage uintptr) int {
	r, _, _ := xImgSrc_LoadFromImage.Call(pImage)
	return int(r)
}

// 图片源_加载文件图标, 加载文件图标, 从一个EXE文件或DLL文件或图标文件; 例如:*.exe文件的图标, 返回图片句柄.
//
// fileName: 文件名.
func XImgSrc_LoadFromExtractIcon(fileName string) int {
	r, _, _ := xImgSrc_LoadFromExtractIcon.Call(common.StrPtr(fileName))
	return int(r)
}

// 图片源_加载从HICON, 创建一个炫彩图片句柄, 从一个现有的图标句柄HICON, 返回图片句柄.
//
// hIcon: 图标句柄.
func XImgSrc_LoadFromHICON(hIcon uintptr) int {
	r, _, _ := xImgSrc_LoadFromHICON.Call(hIcon)
	return int(r)
}

// 图片源_加载从HBITMAP, 创建一个炫彩图片句柄, 从一个现有的位图句柄HBITMAP, 返回图片句柄.
//
// hBitmap: 位图句柄.
func XImgSrc_LoadFromHBITMAP(hBitmap uintptr) int {
	r, _, _ := xImgSrc_LoadFromHBITMAP.Call(hBitmap)
	return int(r)
}

// 图片源_启用自动销毁, 启用或关闭自动销毁, 当与UI元素关联时有效.
//
// hImage: 图片句柄.
//
// bEnable: 是否启用, 不填默认为 true.
func XImgSrc_EnableAutoDestroy(hImage int, bEnable ...bool) {
	enable := true
	if len(bEnable) > 0 {
		enable = bEnable[0]
	}
	xImgSrc_EnableAutoDestroy.Call(uintptr(hImage), common.BoolPtr(enable))
}

// 图片源_取宽度.
//
// hImage: 图片句柄.
func XImgSrc_GetWidth(hImage int) int32 {
	r, _, _ := xImgSrc_GetWidth.Call(uintptr(hImage))
	return int32(r)
}

// 图片源_取高度.
//
// hImage: 图片句柄.
func XImgSrc_GetHeight(hImage int) int32 {
	r, _, _ := xImgSrc_GetHeight.Call(uintptr(hImage))
	return int32(r)
}

// 图片源_取文件名.
//
// hImage: 图片句柄.
func XImgSrc_GetFile(hImage int) string {
	r, _, _ := xImgSrc_GetFile.Call(uintptr(hImage))
	return common.UintPtrToString(r)
}

// 图片源_增加引用计数.
//
// hImage: 图片句柄.
func XImgSrc_AddRef(hImage int) {
	xImgSrc_AddRef.Call(uintptr(hImage))
}

// 图片源_释放引用计数, 当引用计数为0时, 自动销毁.
//
// hImage: 图片句柄.
func XImgSrc_Release(hImage int) {
	xImgSrc_Release.Call(uintptr(hImage))
}

// 图片源_取引用计数.
//
// hImage: 图片句柄.
func XImgSrc_GetRefCount(hImage int) int32 {
	r, _, _ := xImgSrc_GetRefCount.Call(uintptr(hImage))
	return int32(r)
}

// 图片源_销毁, 强制销毁图片, 谨慎使用, 建议使用 XImgSrc_Release() 释放.
//
// hImage: 图片句柄.
func XImgSrc_Destroy(hImage int) {
	xImgSrc_Destroy.Call(uintptr(hImage))
}

// 图片源_加载从数据, 从内存加载原始像素数据(BGRA32位格式), 1个像素占4个字节,分别是(B,G,R,A), 返回炫彩图片句柄.
//
// data: 数据(BGRA32位格式).
//
// width: 图片宽度.
//
// height: 图片高度.
func XImgSrc_LoadFromData(data []byte, width, height int32) int {
	if len(data) < int(width*height*4) {
		return 0
	}
	r, _, _ := xImgSrc_LoadFromData.Call(common.ByteSliceDataPtr(&data), uintptr(width), uintptr(height))
	return int(r)
}

// 图片源_加载从数据RGBA, 从内存加载原始像素数据(RGBA32位格式), 1个像素占4个字节,分别是(R,G,B,A), 返回炫彩图片句柄.
//
// data: 数据(RGBA32位格式).
//
// width: 图片宽度.
//
// height: 图片高度.
func XImgSrc_LoadFromDataRGBA(data []byte, width, height int32) int {
	if len(data) < int(width*height*4) {
		return 0
	}
	d := common.RGBABytesToBGRA(data)
	r, _, _ := xImgSrc_LoadFromData.Call(common.ByteSliceDataPtr(&d), uintptr(width), uintptr(height))
	return int(r)
}

// 图片源_加载从GoImage, 返回炫彩图片句柄.
//
// img: *image.RGBA.
func XImgSrc_LoadFromGoImage(img *image.RGBA) int {
	d := common.RGBABytesToBGRA(img.Pix)
	r, _, _ := xImgSrc_LoadFromData.Call(common.ByteSliceDataPtr(&d), uintptr(img.Rect.Dx()), uintptr(img.Rect.Dy()))
	return int(r)
}

// 图片源_修改数据, 修改图片内存数据(BGRA32位格式), 1个像素占4个字节,分别是(B,G,R,A).
//
// hImage: 图片句柄.
//
// data: 数据(BGRA32位格式).
//
// width: 图片宽度.
//
// height: 图片高度.
func XImgSrc_ModifyData(hImage int, data []byte, width, height int32) bool {
	if len(data) < int(width*height*4) {
		return false
	}
	r, _, _ := xImgSrc_ModifyData.Call(uintptr(hImage), common.ByteSliceDataPtr(&data), uintptr(width), uintptr(height))
	return r != 0
}

// 图片源_修改数据RGBA, 修改图片内存数据(RGBA32位格式), 1个像素占4个字节,分别是(R,G,B,A).
//
// hImage: 图片句柄.
//
// data: 数据(RGBA32位格式).
//
// width: 图片宽度.
//
// height: 图片高度.
func XImgSrc_ModifyDataRGBA(hImage int, data []byte, width, height int32) bool {
	if len(data) < int(width*height*4) {
		return false
	}
	d := common.RGBABytesToBGRA(data)
	r, _, _ := xImgSrc_ModifyData.Call(uintptr(hImage), common.ByteSliceDataPtr(&d), uintptr(width), uintptr(height))
	return r != 0
}

// 图片源_修改数据从GoImage, 修改图片内存数据.
//
// hImage: 图片句柄.
//
// img: *image.RGBA.
func XImgSrc_ModifyDataGoImage(hImage int, img *image.RGBA) bool {
	d := common.RGBABytesToBGRA(img.Pix)
	r, _, _ := xImgSrc_ModifyData.Call(uintptr(hImage), common.ByteSliceDataPtr(&d), uintptr(img.Rect.Dx()), uintptr(img.Rect.Dy()))
	return r != 0
}

// 图片源_取WicBitmap. 获取WIC位图指针, 返回 IWICBitmapSource*.
//
// hImage: 图片句柄.
func XImgSrc_GetWicBitMap(hImage int) uintptr {
	r, _, _ := xImgSrc_GetWicBitMap.Call(uintptr(hImage))
	return r
}

// 图片源_取GdiplusBitmap. 获取GDI+位图指针, 返回 Gdiplus::Bitmap*.
//
// hImage: 图片句柄.
func XImgSrc_GetGdiplusBitmap(hImage int) uintptr {
	r, _, _ := xImgSrc_GetGdiplusBitmap.Call(uintptr(hImage))
	return r
}

package xc

// 图片_加载从图片源.
//
// hImageSrc: 图片源句柄.
func XImage_LoadSrc(hImageSrc int) int {
	r, _, _ := xImage_LoadSrc.Call(uintptr(hImageSrc))
	return int(r)
}

// 图片_加载从文件.
//
// pFileName: 图片文件.
func XImage_LoadFile(pFileName string) int {
	r, _, _ := xImage_LoadFile.Call(strPtr(pFileName))
	return int(r)
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
func XImage_LoadFileAdaptive(pFileName string, leftSize int, topSize int, rightSize int, bottomSize int) int {
	r, _, _ := xImage_LoadFileAdaptive.Call(strPtr(pFileName), uintptr(leftSize), uintptr(topSize), uintptr(rightSize), uintptr(bottomSize))
	return int(r)
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
func XImage_LoadFileRect(pFileName string, x int, y int, cx int, cy int) int {
	r, _, _ := xImage_LoadFileRect.Call(strPtr(pFileName), uintptr(x), uintptr(y), uintptr(cx), uintptr(cy))
	return int(r)
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
func XImage_LoadResAdaptive(id int, pType string, leftSize int, topSize int, rightSize int, bottomSize, hModule int) int {
	r, _, _ := xImage_LoadResAdaptive.Call(uintptr(id), strPtr(pType), uintptr(leftSize), uintptr(topSize), uintptr(rightSize), uintptr(bottomSize), uintptr(hModule))
	return int(r)
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
func XImage_LoadRes(id int, pType string, bStretch bool, hModule int) int {
	r, _, _ := xImage_LoadRes.Call(uintptr(id), strPtr(pType), boolPtr(bStretch), uintptr(hModule))
	return int(r)
}

// 图片_加载从ZIP, 加载图片从ZIP压缩包.
//
// pZipFileName: ZIP压缩包文件名.
//
// pFileName: 图片文件名.
//
// pPassword: ZIP压缩包密码.
func XImage_LoadZip(pZipFileName string, pFileName string, pPassword string) int {
	r, _, _ := xImage_LoadZip.Call(strPtr(pZipFileName), strPtr(pFileName), strPtr(pPassword))
	return int(r)
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
func XImage_LoadZipAdaptive(pZipFileName string, pFileName string, pPassword string, x1 int, x2 int, y1 int, y2 int) int {
	r, _, _ := xImage_LoadZipAdaptive.Call(strPtr(pZipFileName), strPtr(pFileName), strPtr(pPassword), uintptr(x1), uintptr(x2), uintptr(y1), uintptr(y2))
	return int(r)
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
func XImage_LoadZipRect(pZipFileName string, pFileName string, pPassword string, x int, y int, cx int, cy int) int {
	r, _, _ := xImage_LoadZipRect.Call(strPtr(pZipFileName), strPtr(pFileName), strPtr(pPassword), uintptr(x), uintptr(y), uintptr(cx), uintptr(cy))
	return int(r)
}

// 图片_加载从内存ZIP.
//
// data: 内存块指针.
//
// length: 内存块大小.
//
// pFileName: 图片名称.
//
// pPassword: zip压缩包密码.
func XImage_LoadZipMem(data int, length int, pFileName string, pPassword string) int {
	r, _, _ := xImage_LoadZipMem.Call(uintptr(data), uintptr(length), strPtr(pFileName), strPtr(pPassword))
	return int(r)
}

// 图片_加载从内存, 加载流图片, 指定区位置及大小.
//
// pBuffer: 图片缓冲区.
//
// nSize: 图片缓冲区大小.
func XImage_LoadMemory(pBuffer int, nSize int) int {
	r, _, _ := xImage_LoadMemory.Call(uintptr(pBuffer), uintptr(nSize))
	return int(r)
}

// 图片_加载从内存指定区域, 加载流图片, 指定区位置及大小.
//
// pBuffer: 图片缓冲区.
//
// nSize: 图片缓冲区大小.
//
// x: 坐标.
//
// y: 坐标.
//
// cx: 宽度.
//
// cy: 高度.
func XImage_LoadMemoryRect(pBuffer int, nSize int, x int, y int, cx int, cy int) int {
	r, _, _ := xImage_LoadMemoryRect.Call(uintptr(pBuffer), uintptr(nSize), uintptr(x), uintptr(y), uintptr(cx), uintptr(cy))
	return int(r)
}

// 图片_加载从内存自适应, 加载流图片压缩包, 自适应图片(九宫格).
//
// pBuffer: 图片缓冲区.
//
// nSize: 图片缓冲区大小.
//
// leftSize: 坐标.
//
// topSize: 坐标.
//
// rightSize: 坐标.
//
// bottomSize: 坐标.
func XImage_LoadMemoryAdaptive(pBuffer int, nSize int, leftSize int, topSize int, rightSize int, bottomSize int) int {
	r, _, _ := xImage_LoadMemoryAdaptive.Call(uintptr(pBuffer), uintptr(nSize), uintptr(leftSize), uintptr(topSize), uintptr(rightSize), uintptr(bottomSize))
	return int(r)
}

// 图片_加载从Image, 加载图片从GDI+的Image对象.
//
// pImage: GDI图片对象指针Image*.
func XImage_LoadFromImage(pImage int) int {
	r, _, _ := xImage_LoadFromImage.Call(uintptr(pImage))
	return int(r)
}

// 图片_加载文件图标, 加载文件图标, 从一个EXE文件或DLL文件或图标文件; 例如:*.exe文件的图标.
//
// pFileName: 文件名.
func XImage_LoadFromExtractIcon(pFileName string) int {
	r, _, _ := xImage_LoadFromExtractIcon.Call(strPtr(pFileName))
	return int(r)
}

// 图片_加载从HICON, 创建一个炫彩图片句柄, 从一个现有的图标句柄HICON.
//
// hIcon: 图标句柄.
func XImage_LoadFromHICON(hIcon int) int {
	r, _, _ := xImage_LoadFromHICON.Call(uintptr(hIcon))
	return int(r)
}

// 图片_加载从HBITMAP, 创建一个炫彩图片句柄, 从一个现有的位图句柄HBITMAP.
//
// hBitmap: 位图句柄.
func XImage_LoadFromHBITMAP(hBitmap int) int {
	r, _, _ := xImage_LoadFromHBITMAP.Call(uintptr(hBitmap))
	return int(r)
}

// 图片_判断缩放, 是否为拉伸图片句柄.
//
// hImage: 图片句柄.
func XImage_IsStretch(hImage int) bool {
	r, _, _ := xImage_IsStretch.Call(uintptr(hImage))
	return int(r) != 0
}

// 图片_判断自适应, 是否为自适应图片句柄.
//
// hImage: 图片句柄.
func XImage_IsAdaptive(hImage int) bool {
	r, _, _ := xImage_IsAdaptive.Call(uintptr(hImage))
	return int(r) != 0
}

// 图片_判断平铺, 是否为平铺图片.
//
// hImage: 图片句柄.
func XImage_IsTile(hImage int) bool {
	r, _, _ := xImage_IsTile.Call(uintptr(hImage))
	return int(r) != 0
}

// 图片_置绘制类型, 设置图片绘制类型.
//
// hImage: 图片句柄.
//
// nType: 图片绘制类型, Image_Draw_Type_.
func XImage_SetDrawType(hImage int, nType int) bool {
	r, _, _ := xImage_SetDrawType.Call(uintptr(hImage), uintptr(nType))
	return int(r) != 0
}

// 图片_置绘制类型自适应, 设置图片自适应(九宫格).
//
// hImage: 图片句柄.
//
// leftSize: 坐标.
//
// topSize: 坐标.
//
// rightSize: 坐标.
//
// bottomSize: 坐标.
func XImage_SetDrawTypeAdaptive(hImage int, leftSize int, topSize int, rightSize int, bottomSize int) bool {
	r, _, _ := xImage_SetDrawTypeAdaptive.Call(uintptr(hImage), uintptr(leftSize), uintptr(topSize), uintptr(rightSize), uintptr(bottomSize))
	return int(r) != 0
}

// 图片_置透明色, 指定图片透明颜色.
//
// hImage: 图片句柄.
//
// color: ABGR颜色.
func XImage_SetTranColor(hImage int, color int) int {
	r, _, _ := xImage_SetTranColor.Call(uintptr(hImage), uintptr(color))
	return int(r)
}

// 图片_置透明色扩展, 指定图片透明颜色及透明度.
//
// hImage: 图片句柄.
//
// color: ABGR颜色.
//
// tranColor: 透明色的透明度.
func XImage_SetTranColorEx(hImage int, color int, tranColor uint8) int {
	r, _, _ := xImage_SetTranColorEx.Call(uintptr(hImage), uintptr(color), uintptr(tranColor))
	return int(r)
}

// 图片_置旋转角度, 设置旋转角度, 返回先前角度.
//
// hImage: 图片句柄.
//
// fAngle: 选择角度.
func XImage_SetRotateAngle(hImage int, fAngle float32) int {
	r, _, _ := xImage_SetRotateAngle.Call(uintptr(hImage), float32Ptr(fAngle))
	return int(r)
}

// 图片_置等分.
//
// hImage: 图片句柄.
//
// nCount: 等分数量.
//
// iIndex: 索引.
func XImage_SetSplitEqual(hImage int, nCount int, iIndex int) int {
	r, _, _ := xImage_SetSplitEqual.Call(uintptr(hImage), uintptr(nCount), uintptr(iIndex))
	return int(r)
}

// 图片_启用透明色, 启用或关闭图片透明色.
//
// hImage: 图片句柄.
//
// bEnable: 启用TRUE.
func XImage_EnableTranColor(hImage int, bEnable bool) int {
	r, _, _ := xImage_EnableTranColor.Call(uintptr(hImage), boolPtr(bEnable))
	return int(r)
}

// 图片_启用自动销毁, 启用或关闭自动销毁, 当与UI元素关联时有效.
//
// hImage: 图片句柄.
//
// bEnable: 启用自动销毁TRUE.
func XImage_EnableAutoDestroy(hImage int, bEnable bool) int {
	r, _, _ := xImage_EnableAutoDestroy.Call(uintptr(hImage), boolPtr(bEnable))
	return int(r)
}

// 图片_启用居中, 启用或关闭图片居中显示，默认属性图片有效.
//
// hImage: 图片句柄.
//
// bCenter: 是否居中显示.
func XImage_EnableCenter(hImage int, bCenter bool) int {
	r, _, _ := xImage_EnableCenter.Call(uintptr(hImage), boolPtr(bCenter))
	return int(r)
}

// 图片_判断居中, 判断图片是否居中显示.
//
// hImage: 图片句柄.
func XImage_IsCenter(hImage int) bool {
	r, _, _ := xImage_IsCenter.Call(uintptr(hImage))
	return int(r) != 0
}

// 图片_取绘制类型, 获取图片绘制类型, 返回: Image_Draw_Type_.
//
// hImage: 图片句柄.
func XImage_GetDrawType(hImage int) int {
	r, _, _ := xImage_GetDrawType.Call(uintptr(hImage))
	return int(r)
}

// 图片_取宽度.
//
// hImage: 图片句柄.
func XImage_GetWidth(hImage int) int {
	r, _, _ := xImage_GetWidth.Call(uintptr(hImage))
	return int(r)
}

// 图片_取高度.
//
// hImage: 图片句柄.
func XImage_GetHeight(hImage int) int {
	r, _, _ := xImage_GetHeight.Call(uintptr(hImage))
	return int(r)
}

// 图片_取图片源.
//
// hImage: 图片句柄.
func XImage_GetImageSrc(hImage int) int {
	r, _, _ := xImage_GetImageSrc.Call(uintptr(hImage))
	return int(r)
}

// 图片_增加引用计数.
//
// hImage: 图片句柄.
func XImage_AddRef(hImage int) int {
	r, _, _ := xImage_AddRef.Call(uintptr(hImage))
	return int(r)
}

// 图片_释放引用计数, 释放引用计数, 当引用计数为0时, 自动销毁.
//
// hImage: 图片句柄.
func XImage_Release(hImage int) int {
	r, _, _ := xImage_Release.Call(uintptr(hImage))
	return int(r)
}

// 图片_取引用计数.
//
// hImage: 图片句柄.
func XImage_GetRefCount(hImage int) int {
	r, _, _ := xImage_GetRefCount.Call(uintptr(hImage))
	return int(r)
}

// 图片_销毁, 强制销毁图片, 谨慎使用, 建议使用 XImage_Release() 释放.
//
// hImage: 图片句柄.
func XImage_Destroy(hImage int) int {
	r, _, _ := xImage_Destroy.Call(uintptr(hImage))
	return int(r)
}

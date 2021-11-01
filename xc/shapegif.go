package xc

// 形状GIF_创建, 创建形状对象GIF, 成功返回形状对象GIF句柄, 否则返回NULL.
//
// x: X坐标.
//
// y: Y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父对象句柄.
func XShapeGif_Create(x int, y int, cx int, cy int, hParent int) int {
	r, _, _ := xShapeGif_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent))
	return int(r)
}

// 形状GIF_置图片, 设置GIF图片.
//
// hShape: 形状对象句柄.
//
// hImage: 图片句柄.
func XShapeGif_SetImage(hShape int, hImage int) int {
	r, _, _ := xShapeGif_SetImage.Call(uintptr(hShape), uintptr(hImage))
	return int(r)
}

// 形状GIF_取图片, 获取图片句柄.
//
// hShape: 形状对象句柄.
func XShapeGif_GetImage(hShape int) int {
	r, _, _ := xShapeGif_GetImage.Call(uintptr(hShape))
	return int(r)
}

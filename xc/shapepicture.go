package xc

// XShapePic_Create 形状图片_创建, 创建形状对象-图片, 成功返回图片对象句柄, 否则返回NULL.
//
// x: x坐标.
//
// y: y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父对象句柄.
func XShapePic_Create(x int, y int, cx int, cy int, hParent int) int {
	r, _, _ := xShapePic_Create.Call(uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(hParent))
	return int(r)
}

// XShapePic_SetImage 形状图片_置图片, 设置图片.
//
// hShape: 形状对象句柄.
//
// hImage: 图片句柄.
func XShapePic_SetImage(hShape int, hImage int) int {
	r, _, _ := xShapePic_SetImage.Call(uintptr(hShape), uintptr(hImage))
	return int(r)
}

// XShapePic_GetImage 形状图片_取图片, 获取图片句柄.
//
// hShape: 形状对象句柄.
func XShapePic_GetImage(hShape int) int {
	r, _, _ := xShapePic_GetImage.Call(uintptr(hShape))
	return int(r)
}

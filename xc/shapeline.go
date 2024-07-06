package xc

// 形状线_创建.
//
//	x1: 坐标.
//
//	y1: 坐标.
//
//	x2: 坐标.
//
//	y2: 坐标.
//
//	hParent: 父对象句柄.
func XShapeLine_Create(x1, y1, x2, y2 int32, hParent int) int {
	r, _, _ := xShapeLine_Create.Call(uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2), uintptr(hParent))
	return int(r)
}

// 形状线_置位置.
//
//	hShape: 形状对象句柄.
//
//	x1: 坐标.
//
//	y1: 坐标.
//
//	x2: 坐标.
//
//	y2: 坐标.
func XShapeLine_SetPosition(hShape int, x1, y1, x2, y2 int32) {
	xShapeLine_SetPosition.Call(uintptr(hShape), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2))
}

// 形状线_置颜色, 设置直线颜色.
//
//	hShape: 形状对象句柄.
//
//	color: ARGB 颜色值.
func XShapeLine_SetColor(hShape int, color int) {
	xShapeLine_SetColor.Call(uintptr(hShape), uintptr(color))
}

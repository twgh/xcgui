package xc

// XC_GetD2dFactory 炫彩_取D2D工厂, 开启D2D有效, 返回 *ID2D1Factory.
//
// 拿到的是裸接口指针, 可直接调用 Direct2D 的方法(纯 Go, 不需要 cgo), 例如建路径几何;
// 常用封装见 github.com/twgh/xcgui/d2d 包. 该对象由 xcgui 持有, 借用即可, 不要 Release.
func XC_GetD2dFactory() uintptr {
	r, _, _ := xC_GetD2dFactory.Call()
	return r
}

// XC_GetDWriteFactory 炫彩_取DWrite工厂, 开启D2D有效, 返回 *IDWriteFactory.
func XC_GetDWriteFactory() uintptr {
	r, _, _ := xC_GetDWriteFactory.Call()
	return r
}

// XC_GetWicFactory 炫彩_取WIC工厂, 开启D2D有效, 返回 *IWICImagingFactory.
func XC_GetWicFactory() uintptr {
	r, _, _ := xC_GetWicFactory.Call()
	return r
}

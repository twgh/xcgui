package xc

// 炫彩_取D2D工厂, 开启D2D有效, 返回 ID2D1Factory* .
func XC_GetD2dFactory() int {
	r, _, _ := xC_GetD2dFactory.Call()
	return int(r)
}

// 炫彩_取DWrite工厂, 开启D2D有效, 返回 IDWriteFactory* .
func XC_GetDWriteFactory() int {
	r, _, _ := xC_GetDWriteFactory.Call()
	return int(r)
}

// 炫彩_取WIC工厂, 开启D2D有效, 返回 IWICImagingFactory* .
func XC_GetWicFactory() int {
	r, _, _ := xC_GetWicFactory.Call()
	return int(r)
}

package xc

// XC_GetD2dFactory 炫彩_取D2D工厂, 开启D2D有效, 返回 *ID2D1Factory.
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

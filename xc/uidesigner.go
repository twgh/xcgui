package xc

// 炫彩_加载布局文件, 返回窗口句柄或布局句柄或元素句柄.
//
// pFileName: 布局文件名.
//
// hParent: 父对象句柄.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func XC_LoadLayout(pFileName string, hParent, hAttachWnd int) int {
	r, _, _ := xC_LoadLayout.Call(StrPtr(pFileName), uintptr(hParent), uintptr(hAttachWnd))
	return int(r)
}

// 炫彩_加载布局文件ZIP, 加载布局文件从zip压缩包中, 返回窗口句柄或布局句柄或元素句柄.
//
// pZipFileName: zip文件名.
//
// pFileName: 布局文件名.
//
// pPassword: zip密码.
//
// hParent: 父对象句柄.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func XC_LoadLayoutZip(pZipFileName string, pFileName string, pPassword string, hParent, hAttachWnd int) int {
	r, _, _ := xC_LoadLayoutZip.Call(StrPtr(pZipFileName), StrPtr(pFileName), StrPtr(pPassword), uintptr(hParent), uintptr(hAttachWnd))
	return int(r)
}

// 炫彩_加载布局文件内存ZIP, 加载布局文件从zip压缩包中, 返回窗口句柄或布局句柄或元素句柄.
//
// data: 布局文件数据.
//
// pFileName: 布局文件名.
//
// pPassword: zip密码.
//
// hParent: 父对象句柄.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func XC_LoadLayoutZipMem(data []byte, pFileName string, pPassword string, hParent, hAttachWnd int) int {
	r, _, _ := xC_LoadLayoutZipMem.Call(ByteSliceDataPtr(&data), uintptr(len(data)), StrPtr(pFileName), StrPtr(pPassword), uintptr(hParent), uintptr(hAttachWnd))
	return int(r)
}

/* // 炫彩_加载布局文件从字符串, 加载布局文件从内存字符串, 返回窗口句柄或布局句柄或元素句柄.
//
// pStringXML: 字符串指针.
//
// hParent: 父对象.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func XC_LoadLayoutFromString(pStringXML string, hParent, hAttachWnd int) int {
	r, _, _ := xC_LoadLayoutFromString.Call(XC_wtoa(pStringXML), uintptr(hParent), uintptr(hAttachWnd))
	return int(r)
} */

// 炫彩_加载布局文件从字符串UTF8, 加载布局文件从内存字符串, 返回窗口句柄或布局句柄或元素句柄.
//
// pStringXML: 字符串指针.
//
// hParent: 父对象.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func XC_LoadLayoutFromStringUtf8(pStringXML string, hParent, hAttachWnd int) int {
	r, _, _ := xC_LoadLayoutFromStringUtf8.Call(XC_wtoutf8(pStringXML), uintptr(hParent), uintptr(hAttachWnd))
	return int(r)
}

// 炫彩_加载样式文件.
//
// pFileName: 样式文件名称.
func XC_LoadStyle(pFileName string) bool {
	r, _, _ := xC_LoadStyle.Call(StrPtr(pFileName))
	return int(r) != 0
}

// 炫彩_加载样式文件ZIP.
//
// pZipFile: ZIP文件名.
//
// pFileName: 文件名.
//
// pPassword: 密码.
func XC_LoadStyleZip(pZipFile string, pFileName string, pPassword string) bool {
	r, _, _ := xC_LoadStyleZip.Call(StrPtr(pZipFile), StrPtr(pFileName), StrPtr(pPassword))
	return int(r) != 0
}

// 炫彩_加载样式文件从内存ZIP.
//
// data: 样式文件数据.
//
// pFileName: 文件名.
//
// pPassword: 密码.
func XC_LoadStyleZipMem(data []byte, pFileName string, pPassword string) bool {
	r, _, _ := xC_LoadStyleZipMem.Call(ByteSliceDataPtr(&data), uintptr(len(data)), StrPtr(pFileName), StrPtr(pPassword))
	return int(r) != 0
}

// 炫彩_加载资源文件.
//
// pFileName: 资源文件名.
func XC_LoadResource(pFileName string) bool {
	r, _, _ := xC_LoadResource.Call(StrPtr(pFileName))
	return int(r) != 0
}

// 炫彩_加载资源文件ZIP.
//
// pZipFileName: zip文件名.
//
// pFileName: 资源文件名.
//
// pPassword: zip压缩包密码.
func XC_LoadResourceZip(pZipFileName string, pFileName string, pPassword string) bool {
	r, _, _ := xC_LoadResourceZip.Call(StrPtr(pZipFileName), StrPtr(pFileName), StrPtr(pPassword))
	return int(r) != 0
}

// 炫彩_加载资源文件内存ZIP.
//
// data: 资源文件数据.
//
// pFileName: 资源文件名.
//
// pPassword: zip压缩包密码.
func XC_LoadResourceZipMem(data []byte, pFileName string, pPassword string) bool {
	r, _, _ := xC_LoadResourceZipMem.Call(ByteSliceDataPtr(&data), uintptr(len(data)), StrPtr(pFileName), StrPtr(pPassword))
	return int(r) != 0
}

// 炫彩_加载资源文件从字符串UTF8.
//
// pStringXML: 字符串指针.
//
// pFileName: 资源文件名.
func XC_LoadResourceFromStringUtf8(pStringXML string, pFileName string) bool {
	r, _, _ := xC_LoadResourceFromStringUtf8.Call(XC_wtoutf8(pStringXML), StrPtr(pFileName))
	return int(r) != 0
}

// 炫彩_加载样式文件从字符串W.
//
// pFileName: 样式文件名.
//
// pString: 字符串.
func XC_LoadStyleFromStringW(pFileName string, pString string) bool {
	r, _, _ := xC_LoadStyleFromStringW.Call(StrPtr(pFileName), StrPtr(pString))
	return int(r) != 0
}

/*
// 炫彩_加载资源文件从字符串.
//
// pStringXML: 字符串指针.
//
// pFileName: 资源文件名.
func XC_LoadResourceFromString(pStringXML string, pFileName string) bool {
	r, _, _ := xC_LoadResourceFromString.Call(strPtr(pStringXML), strPtr(pFileName))
	return int(r) != 0
}

// 炫彩_加载样式文件从字符串.
//
// pFileName: 样式文件名, 用于打印错误文件和定位关联资源文件位置.
//
// pString: 字符串.
func XC_LoadStyleFromString(pFileName string, pString string) bool {
	r, _, _ := xC_LoadStyleFromString.Call(strPtr(pFileName), strPtr(pString))
	return int(r) != 0
} */

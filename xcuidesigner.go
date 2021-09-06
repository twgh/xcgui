package xcgui

// 炫彩_加载布局文件, 返回窗口句柄或布局句柄或元素句柄
// pFileName: 布局文件名.
// hParent: 父对象句柄
func XC_LoadLayout(pFileName string, hParent int) int {
	r, _, _ := xC_LoadLayout.Call(strPtr(pFileName), uintptr(hParent))
	return int(r)
}

// 炫彩_加载布局文件ZIP, 加载布局文件从zip压缩包中, 返回窗口句柄或布局句柄或元素句柄
// pZipFileName: zip文件名.
// pFileName: 布局文件名.
// pPassword: zip密码.
// hParent: 父对象句柄
func XC_LoadLayoutZip(pZipFileName string, pFileName string, pPassword string, hParent int) int {
	r, _, _ := xC_LoadLayoutZip.Call(strPtr(pZipFileName), strPtr(pFileName), strPtr(pPassword), uintptr(hParent))
	return int(r)
}

// 炫彩_加载布局文件内存ZIP, 加载布局文件从zip压缩包中, 返回窗口句柄或布局句柄或元素句柄
// data: 内存块指针
// length: 内存块大小
// pFileName: 布局文件名.
// pPassword: zip密码.
// hParent: 父对象句柄
func XC_LoadLayoutZipMem(data int, length int, pFileName string, pPassword string, hParent int) int {
	r, _, _ := xC_LoadLayoutZipMem.Call(uintptr(data), uintptr(length), strPtr(pFileName), strPtr(pPassword), uintptr(hParent))
	return int(r)
}

// 炫彩_加载布局文件从字符串, 加载布局文件从内存字符串, 返回窗口句柄或布局句柄或元素句柄
// pStringXML: 字符串指针.
// hParent: 父对象
func XC_LoadLayoutFromString(pStringXML string, hParent int) int {
	r, _, _ := xC_LoadLayoutFromString.Call(strPtr(pStringXML), uintptr(hParent))
	return int(r)
}

// 炫彩_加载布局文件从字符串UTF8, 加载布局文件从内存字符串, 返回窗口句柄或布局句柄或元素句柄
// pStringXML: 字符串指针.
// hParent: 父对象
func XC_LoadLayoutFromStringUtf8(pStringXML string, hParent int) int {
	r, _, _ := xC_LoadLayoutFromStringUtf8.Call(strPtr(pStringXML), uintptr(hParent))
	return int(r)
}

// 炫彩_加载样式文件
// pFileName: 样式文件名称.
func XC_LoadStyle(pFileName string) bool {
	r, _, _ := xC_LoadStyle.Call(strPtr(pFileName))
	return int(r) != 0
}

// 炫彩_加载样式文件ZIP
// pZipFile: ZIP文件名
// pFileName: 文件名
// pPassword: 密码
func XC_LoadStyleZip(pZipFile string, pFileName string, pPassword string) bool {
	r, _, _ := xC_LoadStyleZip.Call(strPtr(pZipFile), strPtr(pFileName), strPtr(pPassword))
	return int(r) != 0
}

// 炫彩_加载样式文件从内存ZIP
// data: 内存块指针
// length: 内存块大小
// pFileName: 文件名
// pPassword: 密码
func XC_LoadStyleZipMem(data int, length int, pFileName string, pPassword string) bool {
	r, _, _ := xC_LoadStyleZipMem.Call(uintptr(data), uintptr(length), strPtr(pFileName), strPtr(pPassword))
	return int(r) != 0
}

// 炫彩_加载资源文件
// pFileName: 资源文件名.
func XC_LoadResource(pFileName string) bool {
	r, _, _ := xC_LoadResource.Call(strPtr(pFileName))
	return int(r) != 0
}

// 炫彩_加载资源文件ZIP
// pZipFileName: zip文件名.
// pFileName: 资源文件名.
// pPassword: zip压缩包密码.
func XC_LoadResourceZip(pZipFileName string, pFileName string, pPassword string) bool {
	r, _, _ := xC_LoadResourceZip.Call(strPtr(pZipFileName), strPtr(pFileName), strPtr(pPassword))
	return int(r) != 0
}

// 炫彩_加载资源文件内存ZIP
// data: 内存块指针
// length: 内存块大小
// pFileName: 资源文件名
// pPassword: zip压缩包密码
func XC_LoadResourceZipMem(data int, length int, pFileName string, pPassword string) bool {
	r, _, _ := xC_LoadResourceZipMem.Call(uintptr(data), uintptr(length), strPtr(pFileName), strPtr(pPassword))
	return int(r) != 0
}

// 炫彩_加载资源文件从字符串
// pStringXML: 字符串指针.
// pFileName: 资源文件名
func XC_LoadResourceFromString(pStringXML string, pFileName string) bool {
	r, _, _ := xC_LoadResourceFromString.Call(strPtr(pStringXML), strPtr(pFileName))
	return int(r) != 0
}

// 炫彩_加载资源文件从字符串UTF8
// pStringXML: 字符串指针.
// pFileName: 资源文件名
func XC_LoadResourceFromStringUtf8(pStringXML string, pFileName string) bool {
	r, _, _ := xC_LoadResourceFromStringUtf8.Call(strPtr(pStringXML), strPtr(pFileName))
	return int(r) != 0
}

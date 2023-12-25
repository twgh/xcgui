package xc

import "github.com/twgh/xcgui/common"

// 炫彩_加载布局文件, 返回窗口句柄或布局句柄或元素句柄.
//
// pFileName: 布局文件名.
//
// hParent: 父对象句柄, 窗口句柄或UI元素句柄.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func XC_LoadLayout(pFileName string, hParent int, hAttachWnd uintptr) int {
	r, _, _ := xC_LoadLayout.Call(common.StrPtr(pFileName), uintptr(hParent), hAttachWnd)
	return int(r)
}

// 炫彩_加载布局文件Ex, 返回窗口句柄或布局句柄或元素句柄.
//
// pFileName: 布局文件名.
//
// pPrefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄, 窗口句柄或UI元素句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func XC_LoadLayoutEx(pFileName, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) int {
	r, _, _ := xC_LoadLayoutEx.Call(common.StrPtr(pFileName), common.StrPtr(pPrefixName), uintptr(hParent), hParentWnd, hAttachWnd)
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
// hParent: 父对象句柄, 窗口句柄或UI元素句柄.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func XC_LoadLayoutZip(pZipFileName string, pFileName string, pPassword string, hParent int, hAttachWnd uintptr) int {
	r, _, _ := xC_LoadLayoutZip.Call(common.StrPtr(pZipFileName), common.StrPtr(pFileName), common.StrPtr(pPassword), uintptr(hParent), hAttachWnd)
	return int(r)
}

// 炫彩_加载布局文件ZIPEx, 加载布局文件从zip压缩包中, 返回窗口句柄或布局句柄或元素句柄.
//
// pZipFileName: zip文件名.
//
// pFileName: 布局文件名.
//
// pPassword: zip密码.
//
// pPrefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄, 窗口句柄或UI元素句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func XC_LoadLayoutZipEx(pZipFileName string, pFileName string, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) int {
	r, _, _ := xC_LoadLayoutZipEx.Call(common.StrPtr(pZipFileName), common.StrPtr(pFileName), common.StrPtr(pPassword), common.StrPtr(pPrefixName), uintptr(hParent), hParentWnd, hAttachWnd)
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
// hParent: 父对象句柄, 窗口句柄或UI元素句柄.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func XC_LoadLayoutZipMem(data []byte, pFileName string, pPassword string, hParent int, hAttachWnd uintptr) int {
	r, _, _ := xC_LoadLayoutZipMem.Call(common.ByteSliceDataPtr(&data), uintptr(len(data)), common.StrPtr(pFileName), common.StrPtr(pPassword), uintptr(hParent), hAttachWnd)
	return int(r)
}

// 炫彩_加载布局文件内存ZIPEx, 加载布局文件从zip压缩包中, 返回窗口句柄或布局句柄或元素句柄.
//
// data: 布局文件数据.
//
// pFileName: 布局文件名.
//
// pPassword: zip密码.
//
// pPrefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄, 窗口句柄或UI元素句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func XC_LoadLayoutZipMemEx(data []byte, pFileName string, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) int {
	r, _, _ := xC_LoadLayoutZipMemEx.Call(common.ByteSliceDataPtr(&data), uintptr(len(data)), common.StrPtr(pFileName), common.StrPtr(pPassword), common.StrPtr(pPrefixName), uintptr(hParent), hParentWnd, hAttachWnd)
	return int(r)
}

/* // 炫彩_加载布局文件从字符串, 加载布局文件从内存字符串, 返回窗口句柄或布局句柄或元素句柄.
//
// pStringXML: 字符串.
//
// hParent: 父对象.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func XC_LoadLayoutFromString(pStringXML string, hParent int, hAttachWnd uintptr) int {
	r, _, _ := xC_LoadLayoutFromString.Call(XC_wtoa(pStringXML), uintptr(hParent), hAttachWnd)
	return int(r)
} */

// 炫彩_加载布局文件从字符串W, 加载布局文件从内存字符串, 返回窗口句柄或布局句柄或元素句柄.
//
// pStringXML: 字符串.
//
// hParent: 父对象句柄, 窗口句柄或UI元素句柄.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func XC_LoadLayoutFromStringW(pStringXML string, hParent int, hAttachWnd uintptr) int {
	r, _, _ := xC_LoadLayoutFromStringUtf8.Call(XC_wtoutf8(pStringXML), uintptr(hParent), hAttachWnd)
	return int(r)
}

// 炫彩_加载布局文件从字符串WEx, 加载布局文件从内存字符串, 返回窗口句柄或布局句柄或元素句柄.
//
// pStringXML: 字符串.
//
// pPrefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄, 窗口句柄或UI元素句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func XC_LoadLayoutFromStringWEx(pStringXML, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) int {
	r, _, _ := xC_LoadLayoutFromStringUtf8Ex.Call(XC_wtoutf8(pStringXML), common.StrPtr(pPrefixName), uintptr(hParent), hParentWnd, hAttachWnd)
	return int(r)
}

// 炫彩_加载样式文件.
//
// pFileName: 样式文件名称.
func XC_LoadStyle(pFileName string) bool {
	r, _, _ := xC_LoadStyle.Call(common.StrPtr(pFileName))
	return r != 0
}

// 炫彩_加载样式文件ZIP.
//
// pZipFile: ZIP文件名.
//
// pFileName: 文件名.
//
// pPassword: 密码.
func XC_LoadStyleZip(pZipFile string, pFileName string, pPassword string) bool {
	r, _, _ := xC_LoadStyleZip.Call(common.StrPtr(pZipFile), common.StrPtr(pFileName), common.StrPtr(pPassword))
	return r != 0
}

// 炫彩_加载样式文件从内存ZIP.
//
// data: 样式文件数据.
//
// pFileName: 文件名.
//
// pPassword: 密码.
func XC_LoadStyleZipMem(data []byte, pFileName string, pPassword string) bool {
	r, _, _ := xC_LoadStyleZipMem.Call(common.ByteSliceDataPtr(&data), uintptr(len(data)), common.StrPtr(pFileName), common.StrPtr(pPassword))
	return r != 0
}

// 炫彩_加载资源文件.
//
// pFileName: 资源文件名.
func XC_LoadResource(pFileName string) bool {
	r, _, _ := xC_LoadResource.Call(common.StrPtr(pFileName))
	return r != 0
}

// 炫彩_加载资源文件ZIP.
//
// pZipFileName: zip文件名.
//
// pFileName: 资源文件名.
//
// pPassword: zip压缩包密码.
func XC_LoadResourceZip(pZipFileName string, pFileName string, pPassword string) bool {
	r, _, _ := xC_LoadResourceZip.Call(common.StrPtr(pZipFileName), common.StrPtr(pFileName), common.StrPtr(pPassword))
	return r != 0
}

// 炫彩_加载资源文件内存ZIP.
//
// data: 资源文件数据.
//
// pFileName: 资源文件名.
//
// pPassword: zip压缩包密码.
func XC_LoadResourceZipMem(data []byte, pFileName string, pPassword string) bool {
	r, _, _ := xC_LoadResourceZipMem.Call(common.ByteSliceDataPtr(&data), uintptr(len(data)), common.StrPtr(pFileName), common.StrPtr(pPassword))
	return r != 0
}

// 炫彩_加载资源文件从字符串W.
//
// pStringXML: 字符串.
//
// pFileName: 资源文件名.
func XC_LoadResourceFromStringW(pStringXML string, pFileName string) bool {
	r, _, _ := xC_LoadResourceFromStringUtf8.Call(XC_wtoutf8(pStringXML), common.StrPtr(pFileName))
	return r != 0
}

// 炫彩_加载样式文件从字符串W.
//
// pFileName: 样式文件名.
//
// pString: 字符串.
func XC_LoadStyleFromStringW(pFileName string, pString string) bool {
	r, _, _ := xC_LoadStyleFromStringW.Call(common.StrPtr(pFileName), common.StrPtr(pString))
	return r != 0
}

// 炫彩_加载布局文件资源ZIP扩展. 加载布局文件从RC资源zip压缩包中, 返回窗口句柄或元素句柄.
//
// id: RC资源ID.
//
// pFileName: 布局文件名.
//
// pPassword: zip密码.
//
// pPrefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄, 窗口句柄或UI元素句柄, 可填0.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用, 可填0.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
//
// hModule: 模块句柄, 可填0.
func XC_LoadLayoutZipResEx(id int32, pFileName string, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd, hModule uintptr) int {
	r, _, _ := xC_LoadLayoutZipResEx.Call(uintptr(id), common.StrPtr(pFileName), common.StrPtr(pPassword), common.StrPtr(pPrefixName), uintptr(hParent), hParentWnd, hAttachWnd, hModule)
	return int(r)
}

// 炫彩_加载资源文件资源ZIP. 加载资源文件从RC资源zip压缩包中.
//
// id: RC资源ID.
//
// pFileName: 资源文件名.
//
// pPassword: zip压缩包密码.
//
// hModule: 模块句柄, 可填0.
func XC_LoadResourceZipRes(id int, pFileName string, pPassword string, hModule uintptr) bool {
	r, _, _ := xC_LoadResourceZipRes.Call(uintptr(id), common.StrPtr(pFileName), common.StrPtr(pPassword), hModule)
	return r != 0
}

// 炫彩_加载样式文件从资源ZIP. 从RC资源中的ZIP包中, 加载样式文件.
//
// id: RC资源ID.
//
// pFileName: 文件名.
//
// pPassword: 密码.
//
// hModule: 模块句柄, 可填0.
func XC_LoadStyleZipRes(id int, pFileName string, pPassword string, hModule uintptr) bool {
	r, _, _ := xC_LoadStyleZipRes.Call(uintptr(id), common.StrPtr(pFileName), common.StrPtr(pPassword), hModule)
	return r != 0
}

/*
// 炫彩_加载资源文件从字符串.
//
// pStringXML: 字符串.
//
// pFileName: 资源文件名.
func XC_LoadResourceFromString(pStringXML string, pFileName string) bool {
	r, _, _ := xC_LoadResourceFromString.Call(strPtr(pStringXML), strPtr(pFileName))
	return r!=0
}

// 炫彩_加载样式文件从字符串.
//
// pFileName: 样式文件名, 用于打印错误文件和定位关联资源文件位置.
//
// pString: 字符串.
func XC_LoadStyleFromString(pFileName string, pString string) bool {
	r, _, _ := xC_LoadStyleFromString.Call(strPtr(pFileName), strPtr(pString))
	return r!=0
} */

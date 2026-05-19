package xc

import "github.com/twgh/xcgui/common"

// 炫彩_加载布局文件, 返回窗口句柄或布局句柄或元素句柄.
//
// fileName: 布局文件名.
//
// hParent: 父对象句柄, 窗口句柄或UI元素句柄.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func XC_LoadLayout(fileName string, hParent int, hAttachWnd uintptr) int {
	r, _, _ := xC_LoadLayout.Call(common.StrPtr(fileName), uintptr(hParent), hAttachWnd)
	return int(r)
}

// 炫彩_加载布局文件Ex, 返回窗口句柄或布局句柄或元素句柄.
//
// fileName: 布局文件名.
//
// pPrefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄, 窗口句柄或UI元素句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func XC_LoadLayoutEx(fileName, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) int {
	r, _, _ := xC_LoadLayoutEx.Call(common.StrPtr(fileName), common.StrPtr(pPrefixName), uintptr(hParent), hParentWnd, hAttachWnd)
	return int(r)
}

// 炫彩_加载布局文件ZIP, 加载布局文件从zip压缩包中, 返回窗口句柄或布局句柄或元素句柄.
//
// zipFileName: zip文件名.
//
// fileName: 布局文件名.
//
// password: zip密码.
//
// hParent: 父对象句柄, 窗口句柄或UI元素句柄.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func XC_LoadLayoutZip(zipFileName string, fileName string, password string, hParent int, hAttachWnd uintptr) int {
	r, _, _ := xC_LoadLayoutZip.Call(common.StrPtr(zipFileName), common.StrPtr(fileName), common.StrPtr(password), uintptr(hParent), hAttachWnd)
	return int(r)
}

// 炫彩_加载布局文件ZIPEx, 加载布局文件从zip压缩包中, 返回窗口句柄或布局句柄或元素句柄.
//
// zipFileName: zip文件名.
//
// fileName: 布局文件名.
//
// password: zip密码.
//
// pPrefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄, 窗口句柄或UI元素句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func XC_LoadLayoutZipEx(zipFileName string, fileName string, password, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) int {
	r, _, _ := xC_LoadLayoutZipEx.Call(common.StrPtr(zipFileName), common.StrPtr(fileName), common.StrPtr(password), common.StrPtr(pPrefixName), uintptr(hParent), hParentWnd, hAttachWnd)
	return int(r)
}

// 炫彩_加载布局文件内存ZIP, 加载布局文件从zip压缩包中, 返回窗口句柄或布局句柄或元素句柄.
//
// data: 布局文件数据.
//
// fileName: 布局文件名.
//
// password: zip密码.
//
// hParent: 父对象句柄, 窗口句柄或UI元素句柄.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func XC_LoadLayoutZipMem(data []byte, fileName string, password string, hParent int, hAttachWnd uintptr) int {
	r, _, _ := xC_LoadLayoutZipMem.Call(common.ByteSliceDataPtr(&data), uintptr(len(data)), common.StrPtr(fileName), common.StrPtr(password), uintptr(hParent), hAttachWnd)
	return int(r)
}

// 炫彩_加载布局文件内存ZIPEx, 加载布局文件从zip压缩包中, 返回窗口句柄或布局句柄或元素句柄.
//
// data: 布局文件数据.
//
// fileName: 布局文件名.
//
// password: zip密码.
//
// pPrefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄, 窗口句柄或UI元素句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func XC_LoadLayoutZipMemEx(data []byte, fileName string, password, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) int {
	r, _, _ := xC_LoadLayoutZipMemEx.Call(common.ByteSliceDataPtr(&data), uintptr(len(data)), common.StrPtr(fileName), common.StrPtr(password), common.StrPtr(pPrefixName), uintptr(hParent), hParentWnd, hAttachWnd)
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
// fileName: 样式文件名称.
func XC_LoadStyle(fileName string) bool {
	r, _, _ := xC_LoadStyle.Call(common.StrPtr(fileName))
	return r != 0
}

// 炫彩_加载样式文件ZIP.
//
// zipFile: ZIP文件名.
//
// fileName: 文件名.
//
// password: 密码, 不填默认为空.
func XC_LoadStyleZip(zipFile string, fileName string, password ...string) bool {
	pwd := ""
	if len(password) > 0 {
		pwd = password[0]
	}
	r, _, _ := xC_LoadStyleZip.Call(common.StrPtr(zipFile), common.StrPtr(fileName), common.StrPtr(pwd))
	return r != 0
}

// 炫彩_加载样式文件从内存ZIP.
//
// data: 样式文件数据.
//
// fileName: 文件名.
//
// password: 密码, 不填默认为空.
func XC_LoadStyleZipMem(data []byte, fileName string, password ...string) bool {
	pwd := ""
	if len(password) > 0 {
		pwd = password[0]
	}
	r, _, _ := xC_LoadStyleZipMem.Call(common.ByteSliceDataPtr(&data), uintptr(len(data)), common.StrPtr(fileName), common.StrPtr(pwd))
	return r != 0
}

// 炫彩_加载资源文件.
//
// fileName: 资源文件名.
func XC_LoadResource(fileName string) bool {
	r, _, _ := xC_LoadResource.Call(common.StrPtr(fileName))
	return r != 0
}

// 炫彩_加载资源文件ZIP.
//
// zipFileName: zip文件名.
//
// fileName: 资源文件名.
//
// password: zip压缩包密码, 不填默认为空.
func XC_LoadResourceZip(zipFileName string, fileName string, password ...string) bool {
	pwd := ""
	if len(password) > 0 {
		pwd = password[0]
	}
	r, _, _ := xC_LoadResourceZip.Call(common.StrPtr(zipFileName), common.StrPtr(fileName), common.StrPtr(pwd))
	return r != 0
}

// 炫彩_加载资源文件内存ZIP.
//
// data: 资源文件数据.
//
// fileName: 资源文件名.
//
// password: zip压缩包密码, 不填默认为空.
func XC_LoadResourceZipMem(data []byte, fileName string, password ...string) bool {
	pwd := ""
	if len(password) > 0 {
		pwd = password[0]
	}
	r, _, _ := xC_LoadResourceZipMem.Call(common.ByteSliceDataPtr(&data), uintptr(len(data)), common.StrPtr(fileName), common.StrPtr(pwd))
	return r != 0
}

// 炫彩_加载资源文件从字符串W.
//
// pStringXML: 字符串.
//
// fileName: 资源文件名.
func XC_LoadResourceFromStringW(pStringXML string, fileName string) bool {
	r, _, _ := xC_LoadResourceFromStringUtf8.Call(XC_wtoutf8(pStringXML), common.StrPtr(fileName))
	return r != 0
}

// 炫彩_加载样式文件从字符串W.
//
// fileName: 样式文件名.
//
// str: 字符串.
func XC_LoadStyleFromStringW(fileName string, str string) bool {
	r, _, _ := xC_LoadStyleFromStringW.Call(common.StrPtr(fileName), common.StrPtr(str))
	return r != 0
}

// 炫彩_加载布局文件资源ZIP扩展. 加载布局文件从RC资源zip压缩包中, 返回窗口句柄或元素句柄.
//
// id: RC资源ID.
//
// fileName: 布局文件名.
//
// password: zip密码.
//
// pPrefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄, 窗口句柄或UI元素句柄, 可填0.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用, 可填0.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
//
// hModule: 模块句柄, 不填默认为0.
func XC_LoadLayoutZipResEx(id int32, fileName string, password, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr, hModule ...uintptr) int {
	module := uintptr(0)
	if len(hModule) > 0 {
		module = hModule[0]
	}
	r, _, _ := xC_LoadLayoutZipResEx.Call(uintptr(id), common.StrPtr(fileName), common.StrPtr(password), common.StrPtr(pPrefixName), uintptr(hParent), hParentWnd, hAttachWnd, module)
	return int(r)
}

// 炫彩_加载资源文件资源ZIP. 加载资源文件从RC资源zip压缩包中.
//
// id: RC资源ID.
//
// fileName: 资源文件名.
//
// password: zip压缩包密码.
//
// hModule: 模块句柄, 不填默认为0.
func XC_LoadResourceZipRes(id int32, fileName string, password string, hModule ...uintptr) bool {
	module := uintptr(0)
	if len(hModule) > 0 {
		module = hModule[0]
	}
	r, _, _ := xC_LoadResourceZipRes.Call(uintptr(id), common.StrPtr(fileName), common.StrPtr(password), module)
	return r != 0
}

// 炫彩_加载样式文件从资源ZIP. 从RC资源中的ZIP包中, 加载样式文件.
//
// id: RC资源ID.
//
// fileName: 文件名.
//
// password: 密码.
//
// hModule: 模块句柄, 不填默认为0.
func XC_LoadStyleZipRes(id int32, fileName string, password string, hModule ...uintptr) bool {
	module := uintptr(0)
	if len(hModule) > 0 {
		module = hModule[0]
	}
	r, _, _ := xC_LoadStyleZipRes.Call(uintptr(id), common.StrPtr(fileName), common.StrPtr(password), module)
	return r != 0
}

/*
// 炫彩_加载资源文件从字符串.
//
// pStringXML: 字符串.
//
// fileName: 资源文件名.
func XC_LoadResourceFromString(pStringXML string, fileName string) bool {
	r, _, _ := xC_LoadResourceFromString.Call(strPtr(pStringXML), strPtr(fileName))
	return r!=0
}

// 炫彩_加载样式文件从字符串.
//
// fileName: 样式文件名, 用于打印错误文件和定位关联资源文件位置.
//
// str: 字符串.
func XC_LoadStyleFromString(fileName string, str string) bool {
	r, _, _ := xC_LoadStyleFromString.Call(strPtr(fileName), strPtr(str))
	return r!=0
} */

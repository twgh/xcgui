package app

import (
	"github.com/twgh/xcgui/svg"
)

// NewSvgByFile SVG_加载从文件, 返回 Svg 对象, 失败返回 nil.
//
// fileName: 文件名.
func NewSvgByFile(fileName string) *svg.Svg {
	return svg.NewByFile(fileName)
}

// NewSvgByString SVG_加载从字符串, 返回 Svg 对象, 失败返回 nil.
//
// str: 字符串.
func NewSvgByString(str string) *svg.Svg {
	return svg.NewByString(str)
}

// NewSvgByZip SVG_加载从ZIP, 返回 Svg 对象, 失败返回 nil.
//
// zipFileName: zip 文件名.
//
// fileName: svg 文件名.
//
// password: zip 密码.
func NewSvgByZip(zipFileName, fileName, password string) *svg.Svg {
	return svg.NewByZip(zipFileName, fileName, password)
}

// NewSvgByZipRes SVG_加载从资源ZIP, 返回 Svg 对象, 失败返回 nil.
//
// id: 资源 ID.
//
// fileName: svg 文件名.
//
// password: zip 密码.
//
// hModule: 模块句柄, 可填0.
func NewSvgByZipRes(id int32, fileName, password string, hModule uintptr) *svg.Svg {
	return svg.NewByZipRes(id, fileName, password, hModule)
}

// NewSvgByZipMem SVG_加载从内存ZIP, 返回 Svg 对象, 失败返回 nil.
//
// data: zip 数据.
//
// fileName: svg 文件名.
//
// password: zip 密码.
func NewSvgByZipMem(data []byte, fileName, password string) *svg.Svg {
	return svg.NewByZipMem(data, fileName, password)
}

// NewSvgByRes SVG_加载从资源, 返回 Svg 对象, 失败返回 nil.
//
// id: 资源 ID.
//
// Type: 资源类型. 在 rc 资源文件中.
//
// hModule: 从指定模块加载.
func NewSvgByRes(id int32, Type string, hModule uintptr) *svg.Svg {
	return svg.NewByRes(id, Type, hModule)
}

// NewSvgByHandle 从句柄创建 Svg 对象, 失败返回 nil.
//
// hSvg: SVG 句柄.
func NewSvgByHandle(hSvg int) *svg.Svg {
	return svg.NewByHandle(hSvg)
}

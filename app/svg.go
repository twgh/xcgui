package app

import (
	"github.com/twgh/xcgui/svg"
)

// NewSvgByFile SVG_加载从文件, 返回 Svg 对象, 失败返回 nil.
//
// pFileName: 文件名.
func NewSvgByFile(pFileName string) *svg.Svg {
	return svg.NewByFile(pFileName)
}

// NewSvgByString SVG_加载从字符串, 返回 Svg 对象, 失败返回 nil.
//
// pString: 字符串.
func NewSvgByString(pString string) *svg.Svg {
	return svg.NewByString(pString)
}

// NewSvgByZip SVG_加载从ZIP, 返回 Svg 对象, 失败返回 nil.
//
// pZipFileName: zip 文件名.
//
// pFileName: svg 文件名.
//
// pPassword: zip 密码.
func NewSvgByZip(pZipFileName, pFileName, pPassword string) *svg.Svg {
	return svg.NewByZip(pZipFileName, pFileName, pPassword)
}

// NewSvgByZipRes SVG_加载从资源ZIP, 返回 Svg 对象, 失败返回 nil.
//
// id: 资源 ID.
//
// pFileName: svg 文件名.
//
// pPassword: zip 密码.
//
// hModule: 模块句柄, 可填0.
func NewSvgByZipRes(id int32, pFileName, pPassword string, hModule uintptr) *svg.Svg {
	return svg.NewByZipRes(id, pFileName, pPassword, hModule)
}

// NewSvgByZipMem SVG_加载从内存ZIP, 返回 Svg 对象, 失败返回 nil.
//
// data: zip 数据.
//
// pFileName: svg 文件名.
//
// pPassword: zip 密码.
func NewSvgByZipMem(data []byte, pFileName, pPassword string) *svg.Svg {
	return svg.NewByZipMem(data, pFileName, pPassword)
}

// NewSvgByRes SVG_加载从资源, 返回 Svg 对象, 失败返回 nil.
//
// id: 资源 ID.
//
// pType: 资源类型. 在 rc 资源文件中.
//
// hModule: 从指定模块加载.
func NewSvgByRes(id int32, pType string, hModule uintptr) *svg.Svg {
	return svg.NewByRes(id, pType, hModule)
}

// NewSvgByHandle 从句柄创建 Svg 对象, 失败返回 nil.
//
// hSvg: SVG 句柄.
func NewSvgByHandle(hSvg int) *svg.Svg {
	return svg.NewByHandle(hSvg)
}

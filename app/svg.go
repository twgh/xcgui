package app

import (
	"github.com/twgh/xcgui/svg"
)

// NewSvgByFile SVG_加载从文件, 返回Svg对象.
//
// pFileName: 文件名.
func NewSvgByFile(pFileName string) *svg.Svg {
	return svg.NewByFile(pFileName)
}

// NewSvgByString SVG_加载从字符串, 返回Svg对象.
//
// pString: 字符串.
func NewSvgByString(pString string) *svg.Svg {
	return svg.NewByString(pString)
}

// NewSvgByStringW SVG_加载从字符串W.
//
// pString: 字符串.
func NewSvgByStringW(pString string) *svg.Svg {
	return svg.NewByStringW(pString)
}

// NewSvgByZip SVG_加载从ZIP, 返回Svg对象.
//
// pZipFileName: zip文件名.
//
// pFileName: svg文件名.
//
// pPassword: zip密码.
func NewSvgByZip(pZipFileName, pFileName, pPassword string) *svg.Svg {
	return svg.NewByZip(pZipFileName, pFileName, pPassword)
}

// NewSvgByZipRes SVG_加载从资源ZIP, 返回SVG对象.
//
// id: 资源ID.
//
// pFileName: svg文件名.
//
// pPassword: zip密码.
//
// hModule: 模块句柄, 可填0.
func NewSvgByZipRes(id int32, pFileName, pPassword string, hModule uintptr) *svg.Svg {
	return svg.NewByZipRes(id, pFileName, pPassword, hModule)
}

// NewSvgByZipMem SVG_加载从内存ZIP, 返回Svg对象.
//
// data: zip数据.
//
// pFileName: svg文件名.
//
// pPassword: zip密码.
func NewSvgByZipMem(data []byte, pFileName, pPassword string) *svg.Svg {
	return svg.NewByZipMem(data, pFileName, pPassword)
}

// NewSvgByRes SVG_加载从资源, 返回Svg对象.
//
// id: 资源ID.
//
// pType: 资源类型.在rc资源文件中.
//
// hModule: 从指定模块加载.
func NewSvgByRes(id int32, pType string, hModule uintptr) *svg.Svg {
	return svg.NewByRes(id, pType, hModule)
}

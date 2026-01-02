package app

import (
	"github.com/twgh/xcgui/tmpl"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// LoadItemTemplate 模板_加载从文件, 项模板文件载入.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pFileName: 文件名.
func LoadItemTemplate(nType xcc.ListItemTemp_Type_, pFileName string) int {
	return xc.XTemp_Load(nType, pFileName)
}

// LoadItemTemplateZip 模板_加载从ZIP, 加载项模板从zip压缩包中.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pZipFile: zip文件.
//
// pFileName: 文件名.
//
// pPassword: zip密码.
func LoadItemTemplateZip(nType xcc.ListItemTemp_Type_, pZipFile string, pFileName string, pPassword string) int {
	return xc.XTemp_LoadZip(nType, pZipFile, pFileName, pPassword)
}

// LoadItemTemplateZipRes 模板_加载从ZIP, 加载项模板从zip压缩包中.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// id: RC资源ID.
//
// pFileName: 模板文件名.
//
// pPassword: zip密码.
//
// hModule: 模块句柄, 可填0.
func LoadItemTemplateZipRes(nType xcc.ListItemTemp_Type_, id int32, pFileName string, pPassword string, hModule uintptr) int {
	return xc.XTemp_LoadZipRes(nType, id, pFileName, pPassword, hModule)
}

// LoadItemTemplateZipMem 模板_加载从内存ZIP, 加载项模板从内存zip压缩包中.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// data: 模板文件数据.
//
// pFileName: 文件名.
//
// pPassword: zip密码.
func LoadItemTemplateZipMem(nType xcc.ListItemTemp_Type_, data []byte, pFileName string, pPassword string) int {
	return xc.XTemp_LoadZipMem(nType, data, pFileName, pPassword)
}

// LoadItemTemplateFromMem 项模板_加载从内存. 加载项模板文件从内存.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// data: 模板文件数据.
func LoadItemTemplateFromMem(nType xcc.ListItemTemp_Type_, data []byte) int {
	return xc.XTemp_LoadFromMem(nType, data)
}

// LoadItemTemplateFromString 模板_加载从字符串, 加载项模板文件从内存字符串.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pStringXML: 字符串.
func LoadItemTemplateFromString(nType xcc.ListItemTemp_Type_, pStringXML string) int {
	return xc.XTemp_LoadFromString(nType, pStringXML)
}

// LoadItemTemplateEx 模板_加载从文件扩展, 加载项模板从文件.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pFileName: 文件名.
//
// pOutTemp1: 返回模板句柄1, 项模板.
//
// pOutTemp2: 返回模板句柄2, 列表头模板或列表视组模板.
func LoadItemTemplateEx(nType xcc.ListItemTemp_Type_, pFileName string, pOutTemp1 *int, pOutTemp2 *int) bool {
	return xc.XTemp_LoadEx(nType, pFileName, pOutTemp1, pOutTemp2)
}

// 模板_加载从文件扩展, 加载项模板从文件, 返回模板对象, 失败返回 nil.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pFileName: 文件名.
//
// 返回值:
//   - 项模板对象.
//   - 列表头模板对象或列表视组模板对象.
func LoadItemTemplateObjEx(nType xcc.ListItemTemp_Type_, pFileName string) (*tmpl.ListItemTemplate, *tmpl.ListItemTemplate) {
	return tmpl.LoadObjEx(nType, pFileName)
}

// LoadItemTemplateZipEx 模板_加载从ZIP扩展, 加载项模板从zip压缩包中.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pZipFile: zip文件.
//
// pFileName: 文件名.
//
// pPassword: zip密码.
//
// pOutTemp1: 返回模板句柄1, 项模板.
//
// pOutTemp2: 返回模板句柄2, 列表头模板或列表视组模板.
func LoadItemTemplateZipEx(nType xcc.ListItemTemp_Type_, pZipFile string, pFileName string, pPassword string, pOutTemp1 *int, pOutTemp2 *int) bool {
	return xc.XTemp_LoadZipEx(nType, pZipFile, pFileName, pPassword, pOutTemp1, pOutTemp2)
}

// 模板_加载从ZIP扩展, 加载项模板从zip压缩包中, 返回模板对象, 失败返回 nil.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pZipFile: zip文件.
//
// pFileName: 文件名.
//
// pPassword: zip密码.
//
// 返回值:
//   - 项模板对象.
//   - 列表头模板对象或列表视组模板对象.
func LoadItemTemplateObjZipEx(nType xcc.ListItemTemp_Type_, pZipFile string, pFileName string, pPassword string) (*tmpl.ListItemTemplate, *tmpl.ListItemTemplate) {
	return tmpl.LoadObjZipEx(nType, pZipFile, pFileName, pPassword)
}

// LoadItemTemplateZipMemEx 模板_加载从内存ZIP扩展, 加载项模板从内存zip压缩包中.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// data: 模板文件数据.
//
// pFileName: 文件名.
//
// pPassword: zip密码.
//
// pOutTemp1: 返回模板句柄1, 项模板.
//
// pOutTemp2: 返回模板句柄2, 列表头模板或列表视组模板.
func LoadItemTemplateZipMemEx(nType xcc.ListItemTemp_Type_, data []byte, pFileName string, pPassword string, pOutTemp1 *int, pOutTemp2 *int) bool {
	return xc.XTemp_LoadZipMemEx(nType, data, pFileName, pPassword, pOutTemp1, pOutTemp2)
}

// 模板_加载从内存ZIP扩展, 加载项模板从内存zip压缩包中, 返回模板对象, 失败返回 nil.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// data: 内存块数据.
//
// pFileName: 文件名.
//
// pPassword: zip密码.
//
// 返回值:
//   - 项模板对象.
//   - 列表头模板对象或列表视组模板对象.
func LoadItemTemplateObjZipMemEx(nType xcc.ListItemTemp_Type_, data []byte, pFileName string, pPassword string) (*tmpl.ListItemTemplate, *tmpl.ListItemTemplate) {
	return tmpl.LoadObjZipMemEx(nType, data, pFileName, pPassword)
}

// LoadItemTemplateFromStringEx 模板_加载从字符串扩展, 加载项模板文件从内存字符串.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pStringXML: 字符串内容.
//
// pOutTemp1: 返回模板句柄1, 项模板.
//
// pOutTemp2: 返回模板句柄2, 列表头模板或列表视组模板.
func LoadItemTemplateFromStringEx(nType xcc.ListItemTemp_Type_, pStringXML string, pOutTemp1 *int, pOutTemp2 *int) bool {
	return xc.XTemp_LoadFromStringEx(nType, pStringXML, pOutTemp1, pOutTemp2)
}

// 模板_加载从字符串扩展, 加载项模板文件从内存字符串, 返回模板对象, 失败返回 nil.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pStringXML: 字符串内容.
//
// 返回值:
//   - 项模板对象.
//   - 列表头模板对象或列表视组模板对象.
func LoadItemTemplateObjFromStringEx(nType xcc.ListItemTemp_Type_, pStringXML string) (*tmpl.ListItemTemplate, *tmpl.ListItemTemplate) {
	return tmpl.LoadObjFromStringEx(nType, pStringXML)
}

// LoadItemTemplateFromMemEx 项模板_加载从内存扩展. 加载项模板文件从内存.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// data: 模板文件数据.
//
// pOutTemp1: 返回模板句柄1, 项模板.
//
// pOutTemp2: 返回模板句柄2, 列表头模板或列表视组模板.
func LoadItemTemplateFromMemEx(nType xcc.ListItemTemp_Type_, data []byte, pOutTemp1 *int, pOutTemp2 *int) bool {
	return xc.XTemp_LoadFromMemEx(nType, data, pOutTemp1, pOutTemp2)
}

// 项模板_加载从内存扩展. 加载项模板文件从内存, 返回模板对象, 失败返回 nil.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// data: 内存块数据.
//
// 返回值:
//   - 项模板对象.
//   - 列表头模板对象或列表视组模板对象.
func LoadItemTemplateObjFromMemEx(nType xcc.ListItemTemp_Type_, data []byte) (*tmpl.ListItemTemplate, *tmpl.ListItemTemplate) {
	return tmpl.LoadObjFromMemEx(nType, data)
}

// LoadItemTemplateZipResEx 项模板_加载从资源ZIP扩展. 加载项模板文件从RC资源ZIP.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// id: RC资源ID.
//
// pFileName: 模板文件名.
//
// pPassword: zip密码.
//
// pOutTemp1: 返回模板句柄1, 项模板.
//
// pOutTemp2: 返回模板句柄2, 列表头模板或列表视组模板.
//
// hModule: 模块句柄, 可填0.
func LoadItemTemplateZipResEx(nType xcc.ListItemTemp_Type_, id int32, pFileName string, pPassword string, pOutTemp1 *int, pOutTemp2 *int, hModule uintptr) bool {
	return xc.XTemp_LoadZipResEx(nType, id, pFileName, pPassword, pOutTemp1, pOutTemp2, hModule)
}

// 项模板_加载从资源ZIP扩展. 加载项模板文件从RC资源ZIP, 返回模板对象, 失败返回 nil.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// id: RC资源ID.
//
// pFileName: 模板文件名.
//
// pPassword: zip密码.
//
// hModule: 模块句柄, 可填0.
func LoadItemTemplateObjZipResEx(nType xcc.ListItemTemp_Type_, id int32, pFileName string, pPassword string, hModule uintptr) (*tmpl.ListItemTemplate, *tmpl.ListItemTemplate) {
	return tmpl.LoadObjZipResEx(nType, id, pFileName, pPassword, hModule)
}

// NewItemTemplate 模板_创建, 创建项模板对象, 失败返回 nil.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
func NewItemTemplate(nType xcc.ListItemTemp_Type_) *tmpl.ListItemTemplate {
	return tmpl.New(nType)
}

// NewItemTemplateByHandle 从句柄创建对象, 失败返回 nil.
func NewItemTemplateByHandle(handle int) *tmpl.ListItemTemplate {
	return tmpl.NewByHandle(handle)
}

// NewItemTemplateByXML 模板_加载从文件, 项模板文件载入, 失败返回 nil.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pFileName: 文件名.
func NewItemTemplateByXML(nType xcc.ListItemTemp_Type_, pFileName string) *tmpl.ListItemTemplate {
	return tmpl.NewByXML(nType, pFileName)
}

// NewItemTemplateByZip 模板_加载从ZIP, 加载项模板从zip压缩包中, 失败返回 nil.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pZipFile: zip文件.
//
// pFileName: 文件名.
//
// pPassword: zip密码.
func NewItemTemplateByZip(nType xcc.ListItemTemp_Type_, pZipFile string, pFileName string, pPassword string) *tmpl.ListItemTemplate {
	return tmpl.NewByZip(nType, pZipFile, pFileName, pPassword)
}

// NewItemTemplateByZipRes 模板_加载从ZIP, 加载项模板从zip压缩包中, 失败返回 nil.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// id: RC资源ID.
//
// pFileName: 模板文件名.
//
// pPassword: zip密码.
//
// hModule: 模块句柄, 可填0.
func NewItemTemplateByZipRes(nType xcc.ListItemTemp_Type_, id int32, pFileName string, pPassword string, hModule uintptr) *tmpl.ListItemTemplate {
	return tmpl.NewByZipRes(nType, id, pFileName, pPassword, hModule)
}

// NewItemTemplateByZipMem 模板_加载从内存ZIP, 加载项模板从内存zip压缩包中, 失败返回 nil.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// data: 模板文件数据.
//
// pFileName: 文件名.
//
// pPassword: zip密码.
func NewItemTemplateByZipMem(nType xcc.ListItemTemp_Type_, data []byte, pFileName string, pPassword string) *tmpl.ListItemTemplate {
	return tmpl.NewByZipMem(nType, data, pFileName, pPassword)
}

// NewItemTemplateByXmlMem 项模板_加载从内存. 加载项模板文件从内存, 失败返回 nil.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// data: 模板文件数据.
func NewItemTemplateByXmlMem(nType xcc.ListItemTemp_Type_, data []byte) *tmpl.ListItemTemplate {
	return tmpl.NewByXmlMem(nType, data)
}

// NewItemTemplateByString 模板_加载从字符串, 加载项模板文件从内存字符串, 失败返回 nil.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pStringXML: 字符串.
func NewItemTemplateByString(nType xcc.ListItemTemp_Type_, pStringXML string) *tmpl.ListItemTemplate {
	return tmpl.NewByString(nType, pStringXML)
}

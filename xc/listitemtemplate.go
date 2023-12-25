package xc

import (
	"unsafe"

	"github.com/twgh/xcgui/common"

	"github.com/twgh/xcgui/xcc"
)

// 模板_加载从文件, 列表项模板文件载入, 返回模板句柄.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pFileName: 文件名.
func XTemp_Load(nType xcc.ListItemTemp_Type_, pFileName string) int {
	r, _, _ := xTemp_Load.Call(uintptr(nType), common.StrPtr(pFileName))
	return int(r)
}

// 模板_加载从ZIP, 加载列表项模板从zip压缩包中, 返回模板句柄.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pZipFile: zip文件.
//
// pFileName: 文件名.
//
// pPassword: zip密码.
func XTemp_LoadZip(nType xcc.ListItemTemp_Type_, pZipFile string, pFileName string, pPassword string) int {
	r, _, _ := xTemp_LoadZip.Call(uintptr(nType), common.StrPtr(pZipFile), common.StrPtr(pFileName), common.StrPtr(pPassword))
	return int(r)
}

// 模板_加载从内存ZIP, 加载列表项模板从内存zip压缩包中, 返回模板句柄.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// data: 模板文件数据.
//
// pFileName: 文件名.
//
// pPassword: zip密码.
func XTemp_LoadZipMem(nType xcc.ListItemTemp_Type_, data []byte, pFileName string, pPassword string) int {
	r, _, _ := xTemp_LoadZipMem.Call(uintptr(nType), common.ByteSliceDataPtr(&data), uintptr(len(data)), common.StrPtr(pFileName), common.StrPtr(pPassword))
	return int(r)
}

// 模板_加载从文件扩展, 加载列表项模板从文件.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pFileName: 文件名.
//
// pOutTemp1: 返回模板句柄1, 项模板.
//
// pOutTemp2: 返回模板句柄2, 列表头模板或列表视组模板.
func XTemp_LoadEx(nType xcc.ListItemTemp_Type_, pFileName string, pOutTemp1 *int, pOutTemp2 *int) bool {
	r, _, _ := xTemp_LoadEx.Call(uintptr(nType), common.StrPtr(pFileName), uintptr(unsafe.Pointer(&pOutTemp1)), uintptr(unsafe.Pointer(&pOutTemp2)))
	return r != 0
}

// 项模板_加载从内存. 加载列表项模板文件从内存, 返回模板句柄.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// data: 模板文件数据.
func XTemp_LoadFromMem(nType xcc.ListItemTemp_Type_, data []byte) int {
	r, _, _ := xTemp_LoadFromMem.Call(uintptr(nType), common.ByteSliceDataPtr(&data), uintptr(len(data)))
	return int(r)
}

// 项模板_加载从内存扩展. 加载列表项模板文件从内存.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// data: 模板文件数据.
//
// pOutTemp1: 返回模板句柄1, 项模板.
//
// pOutTemp2: 返回模板句柄2, 列表头模板或列表视组模板.
func XTemp_LoadFromMemEx(nType xcc.ListItemTemp_Type_, data []byte, pOutTemp1 *int, pOutTemp2 *int) bool {
	r, _, _ := xTemp_LoadFromMemEx.Call(uintptr(nType), common.ByteSliceDataPtr(&data), uintptr(len(data)), uintptr(unsafe.Pointer(&pOutTemp1)), uintptr(unsafe.Pointer(&pOutTemp2)))
	return r != 0
}

// 项模板_加载从资源ZIP. 加载列表项模板文件从RC资源ZIP, 返回模板句柄.
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
func XTemp_LoadZipRes(nType xcc.ListItemTemp_Type_, id int32, pFileName string, pPassword string, hModule uintptr) int {
	r, _, _ := xTemp_LoadZipRes.Call(uintptr(nType), uintptr(id), common.StrPtr(pFileName), common.StrPtr(pPassword), hModule)
	return int(r)
}

// 项模板_加载从资源ZIP扩展. 加载列表项模板文件从RC资源ZIP, 返回模板句柄.
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
func XTemp_LoadZipResEx(nType xcc.ListItemTemp_Type_, id int32, pFileName string, pPassword string, pOutTemp1 *int, pOutTemp2 *int, hModule uintptr) int {
	r, _, _ := xTemp_LoadZipResEx.Call(uintptr(nType), uintptr(id), common.StrPtr(pFileName), common.StrPtr(pPassword), uintptr(unsafe.Pointer(&pOutTemp1)), uintptr(unsafe.Pointer(&pOutTemp2)), hModule)
	return int(r)
}

// 模板_加载从ZIP扩展, 加载列表项模板从zip压缩包中.
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
func XTemp_LoadZipEx(nType xcc.ListItemTemp_Type_, pZipFile string, pFileName string, pPassword string, pOutTemp1 *int, pOutTemp2 *int) bool {
	r, _, _ := xTemp_LoadZipEx.Call(uintptr(nType), common.StrPtr(pZipFile), common.StrPtr(pFileName), common.StrPtr(pPassword), uintptr(unsafe.Pointer(&pOutTemp1)), uintptr(unsafe.Pointer(&pOutTemp2)))
	return r != 0
}

// 模板_加载从内存ZIP扩展, 加载列表项模板从内存zip压缩包中.
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
func XTemp_LoadZipMemEx(nType xcc.ListItemTemp_Type_, data []byte, pFileName string, pPassword string, pOutTemp1 *int, pOutTemp2 *int) bool {
	r, _, _ := xTemp_LoadZipMemEx.Call(uintptr(nType), common.ByteSliceDataPtr(&data), uintptr(len(data)), common.StrPtr(pFileName), common.StrPtr(pPassword), uintptr(unsafe.Pointer(&pOutTemp1)), uintptr(unsafe.Pointer(&pOutTemp2)))
	return r != 0
}

// 模板_加载从字符串, 加载列表项模板文件从内存字符串.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pStringXML: 字符串.
func XTemp_LoadFromString(nType xcc.ListItemTemp_Type_, pStringXML string) int {
	r, _, _ := xTemp_LoadFromString.Call(uintptr(nType), XC_wtoa(pStringXML))
	return int(r)
}

// 模板_加载从字符串扩展, 加载列表项模板文件从内存字符串.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pStringXML: 字符串内容.
//
// pOutTemp1: 返回模板句柄1, 项模板.
//
// pOutTemp2: 返回模板句柄2, 列表头模板或列表视组模板.
func XTemp_LoadFromStringEx(nType xcc.ListItemTemp_Type_, pStringXML string, pOutTemp1 *int, pOutTemp2 *int) bool {
	r, _, _ := xTemp_LoadFromStringEx.Call(uintptr(nType), XC_wtoa(pStringXML), uintptr(unsafe.Pointer(&pOutTemp1)), uintptr(unsafe.Pointer(&pOutTemp2)))
	return r != 0
}

// 模板_取类型, 获取列表项模板类型, 返回: xcc.ListItemTemp_Type_.
//
// hTemp: 列表项模板句柄.
func XTemp_GetType(hTemp int) xcc.ListItemTemp_Type_ {
	r, _, _ := xTemp_GetType.Call(uintptr(hTemp))
	return xcc.ListItemTemp_Type_(r)
}

// 模板_销毁, 项模板销毁.
//
// hTemp: 项模板句柄.
func XTemp_Destroy(hTemp int) bool {
	r, _, _ := xTemp_Destroy.Call(uintptr(hTemp))
	return r != 0
}

// 模板_创建, 创建项模板, 返回模板句柄.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
func XTemp_Create(nType xcc.ListItemTemp_Type_) int {
	r, _, _ := xTemp_Create.Call(uintptr(nType))
	return int(r)
}

// 模板_添加根节点.
//
// hTemp: 项模板句柄.
//
// pNode: 节点指针.
func XTemp_AddNodeRoot(hTemp int, pNode int) bool {
	r, _, _ := xTemp_AddNodeRoot.Call(uintptr(hTemp), uintptr(pNode))
	return r != 0
}

// 模板_添加子节点.
//
// pParentNode: 父节点指针.
//
// pNode: 节点指针.
func XTemp_AddNode(pParentNode int, pNode int) bool {
	r, _, _ := xTemp_AddNode.Call(uintptr(pParentNode), uintptr(pNode))
	return r != 0
}

// 模板_创建节点.
//
// nType: 对象类型: xcc.XC_OBJECT_TYPE.
func XTemp_CreateNode(nType xcc.XC_OBJECT_TYPE) int {
	r, _, _ := xTemp_CreateNode.Call(uintptr(nType))
	return int(r)
}

// 模板_置节点属性.
//
// pNode: 节点指针.
//
// pName: 属性名.
//
// pAttr: 属性值.
func XTemp_SetNodeAttribute(pNode int, pName string, pAttr string) bool {
	r, _, _ := xTemp_SetNodeAttribute.Call(uintptr(pNode), common.StrPtr(pName), common.StrPtr(pAttr))
	return r != 0
}

// 模板_置节点属性扩展.
//
// pNode: 节点指针.
//
// itemID: 模板项ID.
//
// pName: 属性名.
//
// pAttr: 属性值.
func XTemp_SetNodeAttributeEx(pNode int, itemID int32, pName string, pAttr string) bool {
	r, _, _ := xTemp_SetNodeAttributeEx.Call(uintptr(pNode), uintptr(itemID), common.StrPtr(pName), common.StrPtr(pAttr))
	return r != 0
}

// 模板_取列表中的节点.
//
// hTemp: 模板句柄.
//
// index: 节点位置索引.
func XTemp_List_GetNode(hTemp int, index int32) int {
	r, _, _ := xTemp_List_GetNode.Call(uintptr(hTemp), uintptr(index))
	return int(r)
}

// 模板_取节点, 获取节点, 根据itemID. 返回itemID对应的节点指针.
//
// pNode: 节点指针.
//
// itemID: ID.
func XTemp_GetNode(pNode int, itemID int32) int {
	r, _, _ := xTemp_GetNode.Call(uintptr(pNode), uintptr(itemID))
	return int(r)
}

// 模板_克隆节点, 克隆一个节点, 返回克隆的节点.
//
// pNode: 节点指针.
func XTemp_CloneNode(pNode int) int {
	r, _, _ := xTemp_CloneNode.Call(uintptr(pNode))
	return int(r)
}

// 项模板_克隆, 复制一份新的项模板, 返回模板句柄.
//
// hTemp: 列表项模板句柄.
func XTemp_Clone(hTemp int) int {
	r, _, _ := xTemp_Clone.Call(uintptr(hTemp))
	return int(r)
}

// 项模板_列表_插入节点.
//
// hTemp: 列表项模板句柄.
//
// index: 插入位置索引.
//
// pNode: 节点指针.
func XTemp_List_InsertNode(hTemp int, index int32, pNode int) bool {
	r, _, _ := xTemp_List_InsertNode.Call(uintptr(hTemp), uintptr(index), uintptr(pNode))
	return r != 0
}

// 项模板_列表_删除节点.
//
// hTemp: 列表项模板句柄.
//
// index: 位置索引.
func XTemp_List_DeleteNode(hTemp int, index int32) bool {
	r, _, _ := xTemp_List_DeleteNode.Call(uintptr(hTemp), uintptr(index))
	return r != 0
}

// 项模板_列表_取数量, 取子节点数量, 只当前层子节点.
//
// hTemp: 列表项模板句柄.
func XTemp_List_GetCount(hTemp int) int32 {
	r, _, _ := xTemp_List_GetCount.Call(uintptr(hTemp))
	return int32(r)
}

// 项模板_列表_移动列, 将指定列移动到目标位置.
//
// hTemp: 列表项模板句柄.
//
// iColSrc: 源列索引.
//
// iColDest: 目标列索引.
func XTemp_List_MoveColumn(hTemp int, iColSrc, iColDest int32) bool {
	r, _, _ := xTemp_List_MoveColumn.Call(uintptr(hTemp), uintptr(iColSrc), uintptr(iColDest))
	return r != 0
}

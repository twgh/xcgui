package tmpl

import (
	"github.com/twgh/xcgui/objectbase"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// ListItemTemplate 项模板.
type ListItemTemplate struct {
	objectbase.ObjectBase
}

// 模板_创建, 创建项模板, 失败返回 nil.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
func New(nType xcc.ListItemTemp_Type_) *ListItemTemplate {
	return NewByHandle(xc.XTemp_Create(nType))
}

// NewByHandle 从句柄创建对象, 失败返回 nil.
func NewByHandle(handle int) *ListItemTemplate {
	if handle <= 0 {
		return nil
	}
	p := &ListItemTemplate{}
	p.SetHandle(handle)
	return p
}

// 模板_加载从文件, 项模板文件载入, 失败返回 nil.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pFileName: 文件名.
func NewByXML(nType xcc.ListItemTemp_Type_, pFileName string) *ListItemTemplate {
	return NewByHandle(xc.XTemp_Load(nType, pFileName))
}

// 模板_加载从ZIP, 加载项模板从 zip 压缩包中, 失败返回 nil.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pZipFile: zip文件.
//
// pFileName: 文件名.
//
// pPassword: zip密码.
func NewByZip(nType xcc.ListItemTemp_Type_, pZipFile string, pFileName string, pPassword string) *ListItemTemplate {
	return NewByHandle(xc.XTemp_LoadZip(nType, pZipFile, pFileName, pPassword))
}

// 模板_加载从ZIP, 加载项模板从zip压缩包中, 失败返回 nil.
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
func NewByZipRes(nType xcc.ListItemTemp_Type_, id int32, pFileName string, pPassword string, hModule uintptr) *ListItemTemplate {
	return NewByHandle(xc.XTemp_LoadZipRes(nType, id, pFileName, pPassword, hModule))
}

// 模板_加载从内存ZIP, 加载项模板从内存 zip 压缩包中, 失败返回 nil.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// data: 内存块数据.
//
// pFileName: 文件名.
//
// pPassword: zip密码.
func NewByZipMem(nType xcc.ListItemTemp_Type_, data []byte, pFileName string, pPassword string) *ListItemTemplate {
	return NewByHandle(xc.XTemp_LoadZipMem(nType, data, pFileName, pPassword))
}

// 项模板_加载从内存. 加载项模板文件从内存, 失败返回 nil.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// data: 内存块数据.
func NewByXmlMem(nType xcc.ListItemTemp_Type_, data []byte) *ListItemTemplate {
	return NewByHandle(xc.XTemp_LoadFromMem(nType, data))
}

// 模板_加载从字符串, 加载项模板文件从内存字符串, 失败返回 nil.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pStringXML: 字符串.
func NewByString(nType xcc.ListItemTemp_Type_, pStringXML string) *ListItemTemplate {
	return NewByHandle(xc.XTemp_LoadFromString(nType, pStringXML))
}

// 模板_克隆, 复制一份新的项模板, 返回模板对象, 失败返回 nil.
func (l *ListItemTemplate) Clone() *ListItemTemplate {
	return NewByHandle(xc.XTemp_Clone(l.Handle))
}

// 模板_取类型, 获取项模板类型, 返回: xcc.ListItemTemp_Type_.
func (l *ListItemTemplate) GetType() xcc.ListItemTemp_Type_ {
	return xc.XTemp_GetType(l.Handle)
}

// 模板_销毁, 项模板销毁.
func (l *ListItemTemplate) Destroy() bool {
	return xc.XTemp_Destroy(l.Handle)
}

// 模板_添加根节点.
//
// pNode: 节点指针.
func (l *ListItemTemplate) AddNodeRoot(pNode int) bool {
	return xc.XTemp_AddNodeRoot(l.Handle, pNode)
}

// 模板_取列表中的节点.
//
// index: 节点位置索引.
func (l *ListItemTemplate) List_GetNode(index int32) int {
	return xc.XTemp_List_GetNode(l.Handle, index)
}

// 模板_取列表中的节点对象, 失败返回 nil.
//
// index: 节点位置索引.
func (l *ListItemTemplate) List_GetNodeObj(index int32) *Node {
	return NewNodeByHandle(xc.XTemp_List_GetNode(l.Handle, index))
}

// 项模板_列表_插入节点.
//
// index: 插入位置索引.
//
// pNode: 节点指针.
func (l *ListItemTemplate) List_InsertNode(index int32, pNode int) bool {
	return xc.XTemp_List_InsertNode(l.Handle, index, pNode)
}

// 项模板_列表_删除节点.
//
// index: 位置索引.
func (l *ListItemTemplate) List_DeleteNode(index int32) bool {
	return xc.XTemp_List_DeleteNode(l.Handle, index)
}

// 项模板_列表_取数量, 取子节点数量, 只当前层子节点.
func (l *ListItemTemplate) List_GetCount() int32 {
	return xc.XTemp_List_GetCount(l.Handle)
}

// 模板_获取内置模板, 返回模板句柄.
func (l *ListItemTemplate) Get(nType xcc.ListItemTemp_Type_) int {
	return xc.XTemp_Get(nType)
}

// 项模板_列表_移动列, 将指定列移动到目标位置.
//
// iColSrc: 源列索引.
//
// iColDest: 目标列索引.
func (l *ListItemTemplate) List_MoveColumn(iColSrc, iColDest int32) bool {
	return xc.XTemp_List_MoveColumn(l.Handle, iColSrc, iColDest)
}

// 模板_加载从文件扩展, 加载项模板从文件.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pFileName: 文件名.
//
// pOutTemp1: 返回模板句柄1, 项模板.
//
// pOutTemp2: 返回模板句柄2, 列表头模板或列表视组模板.
func LoadEx(nType xcc.ListItemTemp_Type_, pFileName string, pOutTemp1 *int, pOutTemp2 *int) bool {
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
func LoadObjEx(nType xcc.ListItemTemp_Type_, pFileName string) (*ListItemTemplate, *ListItemTemplate) {
	var pOutTemp1, pOutTemp2 int
	if !xc.XTemp_LoadEx(nType, pFileName, &pOutTemp1, &pOutTemp2) {
		return nil, nil
	}
	return NewByHandle(pOutTemp1), NewByHandle(pOutTemp2)
}

// 模板_加载从ZIP扩展, 加载项模板从zip压缩包中.
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
func LoadZipEx(nType xcc.ListItemTemp_Type_, pZipFile string, pFileName string, pPassword string, pOutTemp1 *int, pOutTemp2 *int) bool {
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
func LoadObjZipEx(nType xcc.ListItemTemp_Type_, pZipFile string, pFileName string, pPassword string) (*ListItemTemplate, *ListItemTemplate) {
	var pOutTemp1, pOutTemp2 int
	if !xc.XTemp_LoadZipEx(nType, pZipFile, pFileName, pPassword, &pOutTemp1, &pOutTemp2) {
		return nil, nil
	}
	return NewByHandle(pOutTemp1), NewByHandle(pOutTemp2)
}

// 模板_加载从内存ZIP扩展, 加载项模板从内存zip压缩包中.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// data: 内存块数据.
//
// pFileName: 文件名.
//
// pPassword: zip密码.
//
// pOutTemp1: 返回模板句柄1, 项模板.
//
// pOutTemp2: 返回模板句柄2, 列表头模板或列表视组模板.
func LoadZipMemEx(nType xcc.ListItemTemp_Type_, data []byte, pFileName string, pPassword string, pOutTemp1 *int, pOutTemp2 *int) bool {
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
func LoadObjZipMemEx(nType xcc.ListItemTemp_Type_, data []byte, pFileName string, pPassword string) (*ListItemTemplate, *ListItemTemplate) {
	var pOutTemp1, pOutTemp2 int
	if !xc.XTemp_LoadZipMemEx(nType, data, pFileName, pPassword, &pOutTemp1, &pOutTemp2) {
		return nil, nil
	}
	return NewByHandle(pOutTemp1), NewByHandle(pOutTemp2)
}

// 模板_加载从字符串扩展, 加载项模板文件从内存字符串.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// pStringXML: 字符串内容.
//
// pOutTemp1: 返回模板句柄1, 项模板.
//
// pOutTemp2: 返回模板句柄2, 列表头模板或列表视组模板.
func LoadFromStringEx(nType xcc.ListItemTemp_Type_, pStringXML string, pOutTemp1 *int, pOutTemp2 *int) bool {
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
func LoadObjFromStringEx(nType xcc.ListItemTemp_Type_, pStringXML string) (*ListItemTemplate, *ListItemTemplate) {
	var pOutTemp1, pOutTemp2 int
	if !xc.XTemp_LoadFromStringEx(nType, pStringXML, &pOutTemp1, &pOutTemp2) {
		return nil, nil
	}
	return NewByHandle(pOutTemp1), NewByHandle(pOutTemp2)
}

// 项模板_加载从内存扩展. 加载项模板文件从内存.
//
// nType: 模板类型, xcc.ListItemTemp_Type_.
//
// data: 内存块数据.
//
// pOutTemp1: 返回模板句柄1, 项模板.
//
// pOutTemp2: 返回模板句柄2, 列表头模板或列表视组模板.
func LoadFromMemEx(nType xcc.ListItemTemp_Type_, data []byte, pOutTemp1 *int, pOutTemp2 *int) bool {
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
func LoadObjFromMemEx(nType xcc.ListItemTemp_Type_, data []byte) (*ListItemTemplate, *ListItemTemplate) {
	var pOutTemp1, pOutTemp2 int
	if !xc.XTemp_LoadFromMemEx(nType, data, &pOutTemp1, &pOutTemp2) {
		return nil, nil
	}
	return NewByHandle(pOutTemp1), NewByHandle(pOutTemp2)
}

// 项模板_加载从资源ZIP扩展. 加载项模板文件从RC资源ZIP.
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
func LoadZipResEx(nType xcc.ListItemTemp_Type_, id int32, pFileName string, pPassword string, pOutTemp1 *int, pOutTemp2 *int, hModule uintptr) bool {
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
func LoadObjZipResEx(nType xcc.ListItemTemp_Type_, id int32, pFileName string, pPassword string, hModule uintptr) (*ListItemTemplate, *ListItemTemplate) {
	var pOutTemp1, pOutTemp2 int
	if !xc.XTemp_LoadZipResEx(nType, id, pFileName, pPassword, &pOutTemp1, &pOutTemp2, hModule) {
		return nil, nil
	}
	return NewByHandle(pOutTemp1), NewByHandle(pOutTemp2)
}

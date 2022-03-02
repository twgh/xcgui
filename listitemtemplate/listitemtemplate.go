// 列表项模板.
package listitemtemplate

import (
	"github.com/twgh/xcgui/objectbase"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// 列表项模板.
type ListItemTemplate struct {
	objectbase.ObjectBase
}

// 模板_加载从文件, 列表项模板文件载入.
//
// nType: 模板类型, ListItemTemp_Type_.
//
// pFileName: 文件名.
func NewListItemTemplate_Load(nType xcc.ListItemTemp_Type_, pFileName string) *ListItemTemplate {
	p := &ListItemTemplate{}
	p.SetHandle(xc.XTemp_Load(nType, pFileName))
	return p
}

// 模板_加载从ZIP, 加载列表项模板从zip压缩包中.
//
// nType: 模板类型.
//
// pZipFile: zip文件.
//
// pFileName: 文件名.
//
// pPassword: zip密码.
func NewListItemTemplate_LoadZip(nType int, pZipFile string, pFileName string, pPassword string) *ListItemTemplate {
	p := &ListItemTemplate{}
	p.SetHandle(xc.XTemp_LoadZip(nType, pZipFile, pFileName, pPassword))
	return p
}

// 模板_加载从内存ZIP, 加载列表项模板从内存zip压缩包中.
//
// nType: 模板类型, ListItemTemp_Type_.
//
// data: 模板文件数据.
//
// pFileName: 文件名.
//
// pPassword: zip密码.
func NewListItemTemplate_LoadZipMem(nType xcc.ListItemTemp_Type_, data []byte, pFileName string, pPassword string) *ListItemTemplate {
	p := &ListItemTemplate{}
	p.SetHandle(xc.XTemp_LoadZipMem(nType, data, pFileName, pPassword))
	return p
}

// 模板_加载从字符串, 加载列表项模板文件从内存字符串.
//
// nType: 模板类型, ListItemTemp_Type_.
//
// pStringXML: 字符串指针.
func NewListItemTemplate_LoadFromString(nType xcc.ListItemTemp_Type_, pStringXML string) *ListItemTemplate {
	p := &ListItemTemplate{}
	p.SetHandle(xc.XTemp_LoadFromString(nType, pStringXML))
	return p
}

// 模板_创建, 创建项模板.
//
// nType: 模板类型, ListItemTemp_Type_.
func NewListItemTemplate(nType xcc.ListItemTemp_Type_) *ListItemTemplate {
	p := &ListItemTemplate{}
	p.SetHandle(xc.XTemp_Create(nType))
	return p
}

// 从句柄创建对象.
func NewListItemTemplateByHandle(handle int) *ListItemTemplate {
	p := &ListItemTemplate{}
	p.SetHandle(handle)
	return p
}

// 模板_取类型, 获取列表项模板类型, 返回: ListItemTemp_Type_.
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
func (l *ListItemTemplate) List_GetNode(index int) int {
	return xc.XTemp_List_GetNode(l.Handle, index)
}

// 模板_加载从文件扩展, 加载列表项模板从文件.
//
// nType: 模板类型, ListItemTemp_Type_.
//
// pFileName: 文件名.
//
// pOutTemp1: 返回模板句柄1, 项模板.
//
// pOutTemp2: 返回模板句柄2, 列表头模板或列表视组模板.
func LoadEx(nType xcc.ListItemTemp_Type_, pFileName string, pOutTemp1 *int, pOutTemp2 *int) bool {
	return xc.XTemp_LoadEx(nType, pFileName, pOutTemp1, pOutTemp2)
}

// 模板_加载从ZIP扩展, 加载列表项模板从zip压缩包中.
//
// nType: 模板类型, ListItemTemp_Type_.
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

// 模板_加载从内存ZIP扩展, 加载列表项模板从内存zip压缩包中.
//
// nType: 模板类型, ListItemTemp_Type_.
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
func LoadZipMemEx(nType xcc.ListItemTemp_Type_, data []byte, pFileName string, pPassword string, pOutTemp1 *int, pOutTemp2 *int) bool {
	return xc.XTemp_LoadZipMemEx(nType, data, pFileName, pPassword, pOutTemp1, pOutTemp2)
}

// 模板_加载从字符串扩展, 加载列表项模板文件从内存字符串.
//
// nType: 模板类型, ListItemTemp_Type_.
//
// pStringXML: 字符串内容.
//
// pOutTemp1: 返回模板句柄1, 项模板.
//
// pOutTemp2: 返回模板句柄2, 列表头模板或列表视组模板.
func LoadFromStringEx(nType xcc.ListItemTemp_Type_, pStringXML string, pOutTemp1 *int, pOutTemp2 *int) bool {
	return xc.XTemp_LoadFromStringEx(nType, pStringXML, pOutTemp1, pOutTemp2)
}

// 节点.
type Node struct {
	PNode int // 节点指针.
}

// 模板_创建节点.
//
// nType: 对象类型: XC_.
func NewNode(nType xcc.XC_OBJECT_TYPE) *Node {
	p := &Node{
		PNode: xc.XTemp_CreateNode(nType),
	}
	return p
}

// 给本类的PNode赋值.
func (n *Node) SetPNode(pNode int) {
	n.PNode = pNode
}

// 模板_取节点, 获取节点, 根据itemID. 返回节点对象.
//
// pNode: 节点指针.
//
// itemID: ID.
func (n *Node) GetNode(itemID int) *Node {
	p := &Node{
		PNode: xc.XTemp_GetNode(n.PNode, itemID),
	}
	return p
}

// 模板_克隆节点, 获取列表项模板类型, 返回节点对象.
func (n *Node) CloneNode() *Node {
	p := &Node{
		PNode: xc.XTemp_CloneNode(n.PNode),
	}
	return p
}

// 模板_添加子节点.
//
// pNode: 节点指针.
func (n *Node) AddNode(pNode int) bool {
	return xc.XTemp_AddNode(n.PNode, pNode)
}

// 模板_置节点属性.
//
// pName: 属性名.
//
// pAttr: 属性值.
func (n *Node) SetNodeAttribute(pName string, pAttr string) bool {
	return xc.XTemp_SetNodeAttribute(n.PNode, pName, pAttr)
}

// 模板_置节点属性扩展.
//
// itemID: 模板项ID.
//
// pName: 属性名.
//
// pAttr: 属性值.
func (n *Node) SetNodeAttributeEx(itemID int, pName string, pAttr string) bool {
	return xc.XTemp_SetNodeAttributeEx(n.PNode, itemID, pName, pAttr)
}

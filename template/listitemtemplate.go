// 模板
package template

import (
	"github.com/twgh/xcgui/xc"
)

// 列表项模板
type ListItemTemplate struct {
	HTEMP int //模板句柄
}

// 模板_加载从文件, 列表项模板文件载入
// nType: 模板类型, ListItemTemp_Type_
// pFileName: 文件名.
func New_Load(nType int, pFileName string) *ListItemTemplate {
	p := &ListItemTemplate{
		HTEMP: xc.XTemp_Load(nType, pFileName),
	}
	return p
}

// 模板_加载从ZIP, 加载列表项模板从zip压缩包中
// nType: 模板类型
// pZipFile: zip文件
// pFileName: 文件名
// pPassword: zip密码
func New_LoadZip(nType int, pZipFile string, pFileName string, pPassword string) *ListItemTemplate {
	p := &ListItemTemplate{
		HTEMP: xc.XTemp_LoadZip(nType, pZipFile, pFileName, pPassword),
	}
	return p
}

// 模板_加载从内存ZIP, 加载列表项模板从内存zip压缩包中
// nType: 模板类型, ListItemTemp_Type_
// data: 内存块指针
// length: 内存块大小
// pFileName: 文件名
// pPassword: zip密码
func New_LoadZipMem(nType int, data int, length int, pFileName string, pPassword string) *ListItemTemplate {
	p := &ListItemTemplate{
		HTEMP: xc.XTemp_LoadZipMem(nType, data, length, pFileName, pPassword),
	}
	return p
}

// 模板_加载从字符串, 加载列表项模板文件从内存字符串
// nType: 模板类型, ListItemTemp_Type_
// pStringXML: 字符串指针.
func New_LoadFromString(nType int, pStringXML int) *ListItemTemplate {
	p := &ListItemTemplate{
		HTEMP: xc.XTemp_LoadFromString(nType, pStringXML),
	}
	return p
}

// 模板_创建, 创建项模板
// nType: 模板类型, ListItemTemp_Type_
func New(nType int) *ListItemTemplate {
	p := &ListItemTemplate{
		HTEMP: xc.XTemp_Create(nType),
	}
	return p
}

// 给本类的HTEMP赋值
func (l *ListItemTemplate) SetHTemp(hTemp int) {
	l.HTEMP = hTemp
}

// 模板_取类型, 获取列表项模板类型, 返回: ListItemTemp_Type_
func (l *ListItemTemplate) GetType() int {
	return xc.XTemp_GetType(l.HTEMP)
}

// 模板_销毁, 项模板销毁
func (l *ListItemTemplate) Destroy() bool {
	return xc.XTemp_Destroy(l.HTEMP)
}

// 模板_添加根节点
// pNode: 节点指针.
func (l *ListItemTemplate) AddNodeRoot(pNode int) bool {
	return xc.XTemp_AddNodeRoot(l.HTEMP, pNode)
}

// 模板_取列表中的节点
// index: 节点位置索引
func (l *ListItemTemplate) List_GetNode(index int) int {
	return xc.XTemp_List_GetNode(l.HTEMP, index)
}

// 模板_加载从文件扩展, 加载列表项模板从文件
// nType: 模板类型, ListItemTemp_Type_
// pFileName: 文件名
// pOutTemp1: 返回模板句柄1
// pOutTemp2: 返回模板句柄2
func LoadEx(nType int, pFileName string, pOutTemp1 int, pOutTemp2 int) bool {
	return xc.XTemp_LoadEx(nType, pFileName, pOutTemp1, pOutTemp2)
}

// 模板_加载从ZIP扩展, 加载列表项模板从zip压缩包中
// nType: 模板类型, ListItemTemp_Type_
// pZipFile: zip文件
// pFileName: 文件名
// pPassword: zip密码
// pOutTemp1: 返回模板句柄1, 项模板
// pOutTemp2: 返回模板句柄2, 列表头模板或列表视组模板
func LoadZipEx(nType int, pZipFile string, pFileName string, pPassword string, pOutTemp1 int, pOutTemp2 int) bool {
	return xc.XTemp_LoadZipEx(nType, pZipFile, pFileName, pPassword, pOutTemp1, pOutTemp2)
}

// 模板_加载从内存ZIP扩展, 加载列表项模板从内存zip压缩包中
// nType: 模板类型, ListItemTemp_Type_
// data: 内存块指针
// length: 内存块大小, 字节为单位
// pFileName: 文件名
// pPassword: zip密码
// pOutTemp1: 返回模板句柄1, 项模板
// pOutTemp2: 返回模板句柄2, 列表头模板或列表视组模板
func LoadZipMemEx(nType int, data int, length int, pFileName string, pPassword string, pOutTemp1 int, pOutTemp2 int) bool {
	return xc.XTemp_LoadZipMemEx(nType, data, length, pFileName, pPassword, pOutTemp1, pOutTemp2)
}

// 模板_加载从字符串扩展, 加载列表项模板文件从内存字符串
// nType: 模板类型, ListItemTemp_Type_
// pStringXML: 字符串内容
// pOutTemp1: 返回模板句柄1
// pOutTemp2: 返回模板句柄2
func LoadFromStringEx(nType int, pStringXML int, pOutTemp1 int, pOutTemp2 int) bool {
	return xc.XTemp_LoadFromStringEx(nType, pStringXML, pOutTemp1, pOutTemp2)
}

// 节点
type Node struct {
	PNode int // 节点指针
}

// 模板_创建节点
// nType: 对象类型: XC_
func NewNode(nType int) *Node {
	p := &Node{
		PNode: xc.XTemp_CreateNode(nType),
	}
	return p
}

// 模板_取节点, 获取节点, 根据itemID
// pNode: 节点指针
// itemID: ID.
func (n *Node) GetNode(itemID int) int {
	return xc.XTemp_GetNode(n.PNode, itemID)
}

// 模板_克隆节点, 获取列表项模板类型, 返回节点对象
func (n *Node) CloneNode() *Node {
	p := &Node{
		PNode: xc.XTemp_CloneNode(n.PNode),
	}
	return p
}

// 模板_添加子节点
// pNode: 节点指针.
func (n *Node) AddNode(pNode int) bool {
	return xc.XTemp_AddNode(n.PNode, pNode)
}

// 模板_置节点属性
// pName: 属性名.
// pAttr: 属性值.
func (n *Node) SetNodeAttribute(pName string, pAttr string) bool {
	return xc.XTemp_SetNodeAttribute(n.PNode, pName, pAttr)
}

// 模板_置节点属性扩展
// itemID: 模板项ID
// pName: 属性名
// pAttr: 属性值
func (n *Node) SetNodeAttributeEx(itemID int, pName string, pAttr string) bool {
	return xc.XTemp_SetNodeAttributeEx(n.PNode, itemID, pName, pAttr)
}

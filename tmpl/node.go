package tmpl

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// 节点.
type Node struct {
	Handle int
}

// 模板_创建节点.
//
// nType: 对象类型: xcc.XC_OBJECT_TYPE : xcc.XC_.
func NewNode(nType xcc.XC_OBJECT_TYPE) *Node {
	p := &Node{
		Handle: xc.XTemp_CreateNode(nType),
	}
	return p
}

// 给本类的Handle赋值.
func (n *Node) SetHandle(handle int) {
	n.Handle = handle
}

// 模板_取节点, 获取节点, 根据itemID. 返回节点对象.
//
// pNode: 节点指针.
//
// itemID: ID.
func (n *Node) GetNode(itemID int32) *Node {
	p := &Node{
		Handle: xc.XTemp_GetNode(n.Handle, itemID),
	}
	return p
}

// 模板_克隆节点, 获取列表项模板类型, 返回节点对象.
func (n *Node) CloneNode() *Node {
	p := &Node{
		Handle: xc.XTemp_CloneNode(n.Handle),
	}
	return p
}

// 模板_添加子节点.
//
// pNode: 节点指针.
func (n *Node) AddNode(pNode int) bool {
	return xc.XTemp_AddNode(n.Handle, pNode)
}

// 模板_置节点属性.
//
// pName: 属性名.
//
// pAttr: 属性值.
func (n *Node) SetNodeAttribute(pName string, pAttr string) bool {
	return xc.XTemp_SetNodeAttribute(n.Handle, pName, pAttr)
}

// 模板_置节点属性扩展.
//
// itemID: 模板项ID.
//
// pName: 属性名.
//
// pAttr: 属性值.
func (n *Node) SetNodeAttributeEx(itemID int32, pName string, pAttr string) bool {
	return xc.XTemp_SetNodeAttributeEx(n.Handle, itemID, pName, pAttr)
}

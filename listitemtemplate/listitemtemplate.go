package listitemtemplate

import (
	"github.com/twgh/xcgui/xc"
)

type ListItemTemplate struct {
	HImage int
}

// 模板_加载从文件, 列表项模板文件载入, 返回模板句柄
// nType: 模板类型, ListItemTemp_Type_
// pFileName: 文件名.
func NewTemp_Load(nType int, pFileName string) *ListItemTemplate {
	p := &ListItemTemplate{
		HImage: xc.XTemp_Load(nType, pFileName),
	}
	return p
}

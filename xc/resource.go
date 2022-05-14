package xc

import (
	"github.com/twgh/xcgui/common"
	"syscall"
)

// 资源_启用延迟加载, 启用延迟加载; 图片文件, 列表项模板文件.
//
// bEnable: 是否启用.
func XRes_EnableDelayLoad(bEnable bool) int {
	r, _, _ := xRes_EnableDelayLoad.Call(common.BoolPtr(bEnable))
	return int(r)
}

// 资源_置文件加载回调, 设置文件加载回调函数.
//
// pFun: 回调函数.
func XRes_SetLoadFileCallback(pFun interface{}) int {
	r, _, _ := xRes_SetLoadFileCallback.Call(syscall.NewCallback(pFun))
	return int(r)
}

// 资源_取ID值, 获取资源ID整型值.
//
// pName: 资源ID名称.
func XRes_GetIDValue(pName string) int {
	r, _, _ := xRes_GetIDValue.Call(common.StrPtr(pName))
	return int(r)
}

// 资源_取图片, 查找资源图片.
//
// pName: 资源名称.
func XRes_GetImage(pName string) int {
	r, _, _ := xRes_GetImage.Call(common.StrPtr(pName))
	return int(r)
}

// 资源_取图片扩展, 从指定的资源文件中查找图片.
//
// pFileName: 资源文件名.
//
// pName: 资源名称.
func XRes_GetImageEx(pFileName string, pName string) int {
	r, _, _ := xRes_GetImageEx.Call(common.StrPtr(pFileName), common.StrPtr(pName))
	return int(r)
}

// 资源_取颜色, 从资源中查找颜色.
//
// pName: 资源名称.
func XRes_GetColor(pName string) int {
	r, _, _ := xRes_GetColor.Call(common.StrPtr(pName))
	return int(r)
}

// 资源_取字体, 从资源中查找字体.
//
// pName: 资源名称.
func XRes_GetFont(pName string) int {
	r, _, _ := xRes_GetFont.Call(common.StrPtr(pName))
	return int(r)
}

// 资源_取背景管理器, 从资源中查找背景.
//
// pName: 资源名称.
func XRes_GetBkM(pName string) int {
	r, _, _ := xRes_GetBkM.Call(common.StrPtr(pName))
	return int(r)
}

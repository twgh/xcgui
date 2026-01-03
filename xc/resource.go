package xc

import (
	"syscall"

	"github.com/twgh/xcgui/common"
)

// 资源_启用延迟加载, 启用延迟加载; 图片文件, 列表项模板文件.
//
// bEnable: 是否启用.
func XRes_EnableDelayLoad(bEnable bool) {
	xRes_EnableDelayLoad.Call(common.BoolPtr(bEnable))
}

// 资源_置文件加载回调, 设置文件加载回调函数.
//
// pFun: 回调函数.
func XRes_SetLoadFileCallback(pFun FunLoadFile) {
	xRes_SetLoadFileCallback.Call(syscall.NewCallback(pFun))
}

// 资源_取ID值, 获取资源ID整型值.
//
// name: 资源ID名称.
func XRes_GetIDValue(name string) int32 {
	r, _, _ := xRes_GetIDValue.Call(common.StrPtr(name))
	return int32(r)
}

// 资源_取图片, 查找资源图片.
//
// name: 资源名称.
func XRes_GetImage(name string) int {
	r, _, _ := xRes_GetImage.Call(common.StrPtr(name))
	return int(r)
}

// 资源_取图片扩展, 从指定的资源文件中查找图片.
//
// pFileName: 资源文件名.
//
// name: 资源名称.
func XRes_GetImageEx(pFileName string, name string) int {
	r, _, _ := xRes_GetImageEx.Call(common.StrPtr(pFileName), common.StrPtr(name))
	return int(r)
}

// 资源_取颜色, 从资源中查找颜色.
//
// name: 资源名称.
func XRes_GetColor(name string) uint32 {
	r, _, _ := xRes_GetColor.Call(common.StrPtr(name))
	return uint32(r)
}

// 资源_取字体, 从资源中查找字体.
//
// name: 资源名称.
func XRes_GetFont(name string) int {
	r, _, _ := xRes_GetFont.Call(common.StrPtr(name))
	return int(r)
}

// 资源_取背景管理器, 从资源中查找背景.
//
// name: 资源名称.
func XRes_GetBkM(name string) int {
	r, _, _ := xRes_GetBkM.Call(common.StrPtr(name))
	return int(r)
}

// FunLoadFile 图片资源文件加载回调.
type FunLoadFile func(pFileName uintptr) bool

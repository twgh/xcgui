package res

import "github.com/twgh/xcgui/xc"

// 资源_启用延迟加载, 启用延迟加载; 图片文件, 列表项模板文件.
//
// bEnable: 是否启用.
func EnableDelayLoad(bEnable bool) {
	xc.XRes_EnableDelayLoad(bEnable)
}

// 资源_置文件加载回调, 设置文件加载回调函数.
//
// pFun: 回调函数.
func SetLoadFileCallback(pFun xc.FunLoadFile) {
	xc.XRes_SetLoadFileCallback(pFun)
}

// 资源_取ID值, 获取资源ID整型值.
//
// name: 资源ID名称.
func GetIDValue(name string) int32 {
	return xc.XRes_GetIDValue(name)
}

// 资源_取图片, 查找资源图片.
//
// name: 资源名称.
func GetImage(name string) int {
	return xc.XRes_GetImage(name)
}

// 资源_取图片扩展, 从指定的资源文件中查找图片.
//
// pFileName: 资源文件名.
//
// name: 资源名称.
func GetImageEx(pFileName string, name string) int {
	return xc.XRes_GetImageEx(pFileName, name)
}

// 资源_取颜色, 从资源中查找颜色.
//
// name: 资源名称.
func GetColor(name string) uint32 {
	return xc.XRes_GetColor(name)
}

// 资源_取字体, 从资源中查找字体.
//
// name: 资源名称.
func GetFont(name string) int {
	return xc.XRes_GetFont(name)
}

// 资源_取背景管理器, 从资源中查找背景.
//
// name: 资源名称.
func GetBkM(name string) int {
	return xc.XRes_GetBkM(name)
}

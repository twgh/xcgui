package res

import "github.com/twgh/xcgui/xc"

// 资源_启用延迟加载, 启用延迟加载; 图片文件, 列表项模板文件.
//
// bEnable: 是否启用.
func EnableDelayLoad(bEnable bool) int {
	return xc.XRes_EnableDelayLoad(bEnable)
}

// 资源_置文件加载回调, 设置文件加载回调函数.
//
// pFun: 回调函数.
func SetLoadFileCallback(pFun interface{}) int {
	return xc.XRes_SetLoadFileCallback(pFun)
}

// 资源_取ID值, 获取资源ID整型值.
//
// pName: 资源ID名称.
func GetIDValue(pName string) int {
	return xc.XRes_GetIDValue(pName)
}

// 资源_取图片, 查找资源图片.
//
// pName: 资源名称.
func GetImage(pName string) int {
	return xc.XRes_GetImage(pName)
}

// 资源_取图片扩展, 从指定的资源文件中查找图片.
//
// pFileName: 资源文件名.
//
// pName: 资源名称.
func GetImageEx(pFileName string, pName string) int {
	return xc.XRes_GetImageEx(pFileName, pName)
}

// 资源_取颜色, 从资源中查找颜色.
//
// pName: 资源名称.
func GetColor(pName string) int {
	return xc.XRes_GetColor(pName)
}

// 资源_取字体, 从资源中查找字体.
//
// pName: 资源名称.
func GetFont(pName string) int {
	return xc.XRes_GetFont(pName)
}

// 资源_取背景管理器, 从资源中查找背景.
//
// pName: 资源名称.
func GetBkM(pName string) int {
	return xc.XRes_GetBkM(pName)
}

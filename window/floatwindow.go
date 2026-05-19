package window

import (
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/xc"
)

// FloatWindow 浮动窗口.
type FloatWindow struct {
	windowBase
}

// NewFloatWindowByHandle 从句柄创建浮动窗口对象, 失败返回 nil.
//
// hWindow: 窗口句柄.
func NewFloatWindowByHandle(hWindow int) *FloatWindow {
	if hWindow <= 0 {
		return nil
	}
	p := &FloatWindow{}
	p.SetHandle(hWindow)
	return p
}

// NewFloatWindowByLayout 从布局文件创建对象, 失败返回 nil.
//
// fileName: 布局文件名.
//
// hParent: 父对象句柄.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 不填默认为0.
func NewFloatWindowByLayout(fileName string, hParent int, hAttachWnd ...uintptr) *FloatWindow {
	return NewFloatWindowByHandle(xc.XC_LoadLayout(fileName, hParent, hAttachWnd...))
}

// NewFloatWindowByLayoutZip 从压缩包中的布局文件创建对象, 失败返回 nil.
//
// zipFileName: zip文件名.
//
// fileName: 布局文件名.
//
// password: zip密码.
//
// hParent: 父对象句柄.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 不填默认为0.
func NewFloatWindowByLayoutZip(zipFileName string, fileName string, password string, hParent int, hAttachWnd ...uintptr) *FloatWindow {
	return NewFloatWindowByHandle(xc.XC_LoadLayoutZip(zipFileName, fileName, password, hParent, hAttachWnd...))
}

// NewFloatWindowByLayoutZipMem 从内存压缩包中的布局文件创建对象, 失败返回 nil.
//
// data: 布局文件数据.
//
// fileName: 布局文件名.
//
// password: zip密码.
//
// hParent: 父对象句柄.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 不填默认为0.
func NewFloatWindowByLayoutZipMem(data []byte, fileName string, password string, hParent int, hAttachWnd ...uintptr) *FloatWindow {
	return NewFloatWindowByHandle(xc.XC_LoadLayoutZipMem(data, fileName, password, hParent, hAttachWnd...))
}

// NewFloatWindowByLayoutString 从布局文件字符串创建对象, 失败返回 nil.
//
// xmlStr: 字符串.
//
// hParent: 父对象.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 不填默认为0.
func NewFloatWindowByLayoutString(xmlStr string, hParent int, hAttachWnd ...uintptr) *FloatWindow {
	return NewFloatWindowByHandle(xc.XC_LoadLayoutFromStringW(xmlStr, hParent, hAttachWnd...))
}

// NewFloatWindowByLayoutEx 从布局文件创建对象, 失败返回 nil.
//
// fileName: 布局文件名.
//
// prefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 不填默认为0.
func NewFloatWindowByLayoutEx(fileName, prefixName string, hParent int, hParentWnd uintptr, hAttachWnd ...uintptr) *FloatWindow {
	return NewFloatWindowByHandle(xc.XC_LoadLayoutEx(fileName, prefixName, hParent, hParentWnd, hAttachWnd...))
}

// NewFloatWindowByLayoutZipResEx 从RC资源zip压缩包中的布局文件创建对象, 失败返回 nil.
//
// id: RC资源ID.
//
// fileName: 布局文件名.
//
// password: zip密码.
//
// prefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
//
// hModule: 模块句柄, 不填默认为0.
func NewFloatWindowByLayoutZipResEx(id int32, fileName, password, prefixName string, hParent int, hParentWnd, hAttachWnd uintptr, hModule ...uintptr) *FloatWindow {
	return NewFloatWindowByHandle(xc.XC_LoadLayoutZipResEx(id, fileName, password, prefixName, hParent, hParentWnd, hAttachWnd, hModule...))
}

// NewFloatWindowByLayoutZipEx 从压缩包中的布局文件创建对象, 失败返回 nil.
//
// zipFileName: zip文件名.
//
// fileName: 布局文件名.
//
// password: zip密码.
//
// prefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 不填默认为0.
func NewFloatWindowByLayoutZipEx(zipFileName string, fileName string, password, prefixName string, hParent int, hParentWnd uintptr, hAttachWnd ...uintptr) *FloatWindow {
	return NewFloatWindowByHandle(xc.XC_LoadLayoutZipEx(zipFileName, fileName, password, prefixName, hParent, hParentWnd, hAttachWnd...))
}

// NewFloatWindowByLayoutZipMemEx 从内存压缩包中的布局文件创建对象, 失败返回 nil.
//
// data: 布局文件数据.
//
// fileName: 布局文件名.
//
// password: zip密码.
//
// prefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 不填默认为0.
func NewFloatWindowByLayoutZipMemEx(data []byte, fileName string, password, prefixName string, hParent int, hParentWnd uintptr, hAttachWnd ...uintptr) *FloatWindow {
	return NewFloatWindowByHandle(xc.XC_LoadLayoutZipMemEx(data, fileName, password, prefixName, hParent, hParentWnd, hAttachWnd...))
}

// NewFloatWindowByLayoutStringEx 从布局文件字符串创建对象, 失败返回 nil.
//
// xmlStr: 字符串.
//
// prefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 不填默认为0.
func NewFloatWindowByLayoutStringEx(xmlStr, prefixName string, hParent int, hParentWnd uintptr, hAttachWnd ...uintptr) *FloatWindow {
	return NewFloatWindowByHandle(xc.XC_LoadLayoutFromStringWEx(xmlStr, prefixName, hParent, hParentWnd, hAttachWnd...))
}

// 从name创建对象, 失败返回 nil.
func NewFloatWindowByName(name string) *FloatWindow {
	return NewFloatWindowByHandle(xc.XC_GetObjectByName(name))
}

// 从UID创建对象, 失败返回 nil.
func NewFloatWindowByUID(nUID int32) *FloatWindow {
	return NewFloatWindowByHandle(xc.XC_GetObjectByUID(nUID))
}

// 从UID名称创建对象, 失败返回 nil.
func NewFloatWindowByUIDName(name string) *FloatWindow {
	return NewFloatWindowByHandle(xc.XC_GetObjectByUIDName(name))
}

// EnableCaptionContent 浮动窗口_启用标题栏内容, 启用后,将创建标题(形状文本)和关闭按钮.
//
// bEnable: 是否启用, 不填默认为 true.
func (fw *FloatWindow) EnableCaptionContent(bEnable ...bool) bool {
	return xc.XFloatWnd_EnableCaptionContent(fw.Handle, bEnable...)
}

// GetCaptionShapeText 浮动窗口_获取标题形状文本, 启用标题栏内容后有效, 返回形状文本句柄.
func (fw *FloatWindow) GetCaptionShapeText() int {
	return xc.XFloatWnd_GetCaptionShapeText(fw.Handle)
}

// GetCaptionShapeTextObj 浮动窗口_获取标题形状文本, 启用标题栏内容后有效, 返回形状文本对象.
func (fw *FloatWindow) GetCaptionShapeTextObj() *widget.ShapeText {
	return widget.NewShapeTextByHandle(xc.XFloatWnd_GetCaptionShapeText(fw.Handle))
}

// GetCaptionButtonClose 浮动窗口_获取标题关闭按钮, 启用标题栏内容后有效, 返回关闭按钮句柄.
func (fw *FloatWindow) GetCaptionButtonClose() int {
	return xc.XFloatWnd_GetCaptionButtonClose(fw.Handle)
}

// GetCaptionButtonCloseObj 浮动窗口_获取标题关闭按钮, 启用标题栏内容后有效, 返回关闭按钮对象.
func (fw *FloatWindow) GetCaptionButtonCloseObj() *widget.Button {
	return widget.NewButtonByHandle(xc.XFloatWnd_GetCaptionButtonClose(fw.Handle))
}

// SetTitle 浮动窗口_设置标题, 启用标题栏内容后有效.
//
// pTitle: 标题文本.
func (fw *FloatWindow) SetTitle(pTitle string) *FloatWindow {
	xc.XFloatWnd_SetTitle(fw.Handle, pTitle)
	return fw
}

// GetTitle 浮动窗口_获取标题, 启用标题栏内容后有效, 返回标题文本.
func (fw *FloatWindow) GetTitle() string {
	return xc.XFloatWnd_GetTitle(fw.Handle)
}

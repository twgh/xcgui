package window

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// ModalWindow 模态窗口.
type ModalWindow struct {
	windowBase
}

// 模态窗口_创建, 创建模态窗口; 当模态窗口关闭时, 会自动销毁模态窗口资源句柄, 失败返回 nil.
//
// nWidth: 宽度.
//
// nHeight: 高度.
//
// title: 窗口标题内容.
//
// hWndParent: 父窗口句柄.
//
// XCStyle: 炫彩窗口样式: xcc.Window_Style_.
func NewModalWindow(nWidth, nHeight int32, title string, hWndParent uintptr, XCStyle xcc.Window_Style_) *ModalWindow {
	return NewModalWindowByHandle(xc.XModalWnd_Create(nWidth, nHeight, title, hWndParent, XCStyle))
}

// 模态窗口_创建扩展, 创建模态窗口, 增强功能, 失败返回 nil.
//
// dwExStyle: 窗口扩展样式.
//
// dwStyle: 窗口样式.
//
// lpClassName: 窗口类名.
//
// x: 窗口左上角x坐标.
//
// y: 窗口左上角y坐标.
//
// cx: 窗口宽度.
//
// cy: 窗口高度.
//
// title: 窗口名.
//
// hWndParent: 父窗口.
//
// XCStyle: GUI库窗口样式: xcc.Window_Style_.
func NewModalWindowEx(dwExStyle, dwStyle uint32, lpClassName string, x, y, cx, cy int32, title string, hWndParent uintptr, XCStyle xcc.Window_Style_) *ModalWindow {
	return NewModalWindowByHandle(xc.XModalWnd_CreateEx(dwExStyle, dwStyle, title, x, y, cx, cy, lpClassName, hWndParent, XCStyle))
}

// NewModalWindowByLayout 从布局文件创建对象, 失败返回 nil.
//
// fileName: 布局文件名.
//
// hParent: 父对象句柄.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 不填默认为0.
func NewModalWindowByLayout(fileName string, hParent int, hAttachWnd ...uintptr) *ModalWindow {
	return NewModalWindowByHandle(xc.XC_LoadLayout(fileName, hParent, hAttachWnd...))
}

// NewModalWindowByLayoutZip 从压缩包中的布局文件创建对象, 失败返回 nil.
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
func NewModalWindowByLayoutZip(zipFileName string, fileName string, password string, hParent int, hAttachWnd ...uintptr) *ModalWindow {
	return NewModalWindowByHandle(xc.XC_LoadLayoutZip(zipFileName, fileName, password, hParent, hAttachWnd...))
}

// NewModalWindowByLayoutZipMem 从内存压缩包中的布局文件创建对象, 失败返回 nil.
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
func NewModalWindowByLayoutZipMem(data []byte, fileName string, password string, hParent int, hAttachWnd ...uintptr) *ModalWindow {
	return NewModalWindowByHandle(xc.XC_LoadLayoutZipMem(data, fileName, password, hParent, hAttachWnd...))
}

// NewModalWindowByLayoutStringW 从布局文件字符串W创建对象, 失败返回 nil.
//
// pStringXML: 字符串.
//
// hParent: 父对象.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 不填默认为0.
func NewModalWindowByLayoutStringW(pStringXML string, hParent int, hAttachWnd ...uintptr) *ModalWindow {
	return NewModalWindowByHandle(xc.XC_LoadLayoutFromStringW(pStringXML, hParent, hAttachWnd...))
}

// NewModalWindowByLayoutEx 从布局文件创建对象, 失败返回 nil.
//
// fileName: 布局文件名.
//
// pPrefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 不填默认为0.
func NewModalWindowByLayoutEx(fileName, pPrefixName string, hParent int, hParentWnd uintptr, hAttachWnd ...uintptr) *ModalWindow {
	return NewModalWindowByHandle(xc.XC_LoadLayoutEx(fileName, pPrefixName, hParent, hParentWnd, hAttachWnd...))
}

// NewModalWindowByLayoutZipResEx 从RC资源zip压缩包中的布局文件创建对象, 失败返回 nil.
//
// id: RC资源ID.
//
// fileName: 布局文件名.
//
// password: zip密码.
//
// pPrefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
//
// hModule: 模块句柄, 不填默认为0.
func NewModalWindowByLayoutZipResEx(id int32, fileName, password, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr, hModule ...uintptr) *ModalWindow {
	return NewModalWindowByHandle(xc.XC_LoadLayoutZipResEx(id, fileName, password, pPrefixName, hParent, hParentWnd, hAttachWnd, hModule...))
}

// NewModalWindowByLayoutZipEx 从压缩包中的布局文件创建对象, 失败返回 nil.
//
// zipFileName: zip文件名.
//
// fileName: 布局文件名.
//
// password: zip密码.
//
// pPrefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 不填默认为0.
func NewModalWindowByLayoutZipEx(zipFileName string, fileName string, password, pPrefixName string, hParent int, hParentWnd uintptr, hAttachWnd ...uintptr) *ModalWindow {
	return NewModalWindowByHandle(xc.XC_LoadLayoutZipEx(zipFileName, fileName, password, pPrefixName, hParent, hParentWnd, hAttachWnd...))
}

// NewModalWindowByLayoutZipMemEx 从内存压缩包中的布局文件创建对象, 失败返回 nil.
//
// data: 布局文件数据.
//
// fileName: 布局文件名.
//
// password: zip密码.
//
// pPrefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 不填默认为0.
func NewModalWindowByLayoutZipMemEx(data []byte, fileName string, password, pPrefixName string, hParent int, hParentWnd uintptr, hAttachWnd ...uintptr) *ModalWindow {
	return NewModalWindowByHandle(xc.XC_LoadLayoutZipMemEx(data, fileName, password, pPrefixName, hParent, hParentWnd, hAttachWnd...))
}

// NewModalWindowByLayoutStringWEx 从布局文件字符串W创建对象, 失败返回 nil.
//
// pStringXML: 字符串.
//
// pPrefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 不填默认为0.
func NewModalWindowByLayoutStringWEx(pStringXML, pPrefixName string, hParent int, hParentWnd uintptr, hAttachWnd ...uintptr) *ModalWindow {
	return NewModalWindowByHandle(xc.XC_LoadLayoutFromStringWEx(pStringXML, pPrefixName, hParent, hParentWnd, hAttachWnd...))
}

// 模态窗口_附加窗口, 返回窗口对象, 失败返回 nil.
//
// hWnd: 要附加的外部窗口句柄.
//
// XCStyle: 炫彩窗口样式: xcc.Window_Style_.
func ModalWnd_Attach(hWnd uintptr, XCStyle xcc.Window_Style_) *Window {
	return NewByHandle(xc.XModalWnd_Attach(hWnd, XCStyle))
}

// 从句柄创建对象, 失败返回 nil.
func NewModalWindowByHandle(handle int) *ModalWindow {
	if handle <= 0 {
		return nil
	}
	p := &ModalWindow{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回 nil.
func NewModalWindowByName(name string) *ModalWindow {
	return NewModalWindowByHandle(xc.XC_GetObjectByName(name))
}

// 从UID创建对象, 失败返回 nil.
func NewModalWindowByUID(nUID int32) *ModalWindow {
	return NewModalWindowByHandle(xc.XC_GetObjectByUID(nUID))
}

// 从UID名称创建对象, 失败返回 nil.
func NewModalWindowByUIDName(name string) *ModalWindow {
	return NewModalWindowByHandle(xc.XC_GetObjectByUIDName(name))
}

// 模态窗口_启用自动关闭, 是否自动关闭窗口, 当窗口失去焦点时.
//
// bEnable: 是否启用, 不填默认为 true.
func (w *ModalWindow) EnableAutoClose(bEnable ...bool) *ModalWindow {
	xc.XModalWnd_EnableAutoClose(w.Handle, bEnable...)
	return w
}

// 模态窗口_启用ESC关闭, 当用户按ESC键时自动关闭模态窗口.
//
// bEnable: 是否启用, 不填默认为 true.
func (w *ModalWindow) EnableEscClose(bEnable ...bool) *ModalWindow {
	xc.XModalWnd_EnableEscClose(w.Handle, bEnable...)
	return w
}

// 模态窗口_启动, 启动显示模态窗口, 当窗口关闭时返回:
//   - xcc.MessageBox_Flag_Ok: 点击确定按钮退出.
//   - xcc.MessageBox_Flag_Cancel: 点击取消按钮退出.
//   - xcc.MessageBox_Flag_Other: 其他方式退出.
func (w *ModalWindow) DoModal() xcc.MessageBox_Flag_ {
	return xc.XModalWnd_DoModal(w.Handle)
}

// 模态窗口_结束, 结束模态窗口.
//
// nResult: 用作 xc.XModalWnd_DoModal 的返回值:
//   - xcc.MessageBox_Flag_Ok: 点击确定按钮退出.
//   - xcc.MessageBox_Flag_Cancel: 点击取消按钮退出.
//   - xcc.MessageBox_Flag_Other: 其他方式退出.
func (w *ModalWindow) EndModal(nResult xcc.MessageBox_Flag_) *ModalWindow {
	xc.XModalWnd_EndModal(w.Handle, nResult)
	return w
}

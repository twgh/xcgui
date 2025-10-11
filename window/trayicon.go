package window

import (
	"math/rand"
	"time"

	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// TrayIcon 封装且简化了对托盘图标的操作. 更多操作可使用 xc.XTrayIcon_ 系列函数.
type TrayIcon struct {
	HIcon   uintptr // 图标句柄. 仅用于记录.
	HWindow int     // 炫彩窗口句柄. 仅用于记录.
	Tips    string  // 提示文本. 仅用于记录.
	Id      int32   // 托盘图标唯一标识符, 会传到托盘事件的第一个参数里. 仅用于记录.
	Result  bool    // 由于采用了链式调用的写法, 所以在这里判断函数是否执行成功. 或者在链式调用后使用 IsOk 函数来判断.
}

// 用于记录已有的托盘图标标识符.
var ids []int32

// NewTrayIcon 创建托盘图标对象. 之后需调用 Show 函数来显示到托盘.
//
// hWindow: 炫彩窗口句柄.
//
// hIcon: 图标句柄. 可使用 wutil.HIcon 函数生成.
//
// tips: 提示文本, 长度不能超过 127 个字符.
func NewTrayIcon(hWindow int, hIcon uintptr, tips string) *TrayIcon {
	xc.XTrayIcon_SetIcon(hIcon)
	xc.XTrayIcon_SetTips(tips)

	rand.Seed(time.Now().UnixNano())
	var id int32
	for {
		id = rand.Int31()
		if id != 0 && sliceIndexOf(ids, id) == -1 { // 不等于0且不是已有的
			ids = append(ids, id)
			break
		}
	}
	return &TrayIcon{HWindow: hWindow, HIcon: hIcon, Result: true, Id: id}
}

// IsOk 用于判断上一个函数是否执行成功, 因为采用了链式调用的写法.
func (t *TrayIcon) IsOk() bool {
	return t.Result
}

// Show 托盘图标_显示. 从系统托盘显示或隐藏图标.
//
// isShow: 是否显示. 内部原理如下:
//   - true: 执行 xc.XTrayIcon_Add().
//   - false: 执行 xc.XTrayIcon_Del().
func (t *TrayIcon) Show(isShow bool) *TrayIcon {
	if isShow {
		t.Result = xc.XTrayIcon_Add(t.HWindow, t.Id)
	} else {
		t.Result = xc.XTrayIcon_Del()
	}
	return t
}

// Del 托盘图标_删除. 从系统托盘删除图标. 等同于 Show(false).
func (t *TrayIcon) Del() *TrayIcon {
	t.Result = xc.XTrayIcon_Del()
	return t
}

// Modify 托盘图标_修改. 应用对于托盘图标的修改.
func (t *TrayIcon) Modify() *TrayIcon {
	t.Result = xc.XTrayIcon_Modify()
	return t
}

// SetFocus 托盘图标_置焦点. 将焦点设置回系统托盘.
func (t *TrayIcon) SetFocus() *TrayIcon {
	t.Result = xc.XTrayIcon_SetFocus()
	return t
}

// SetIcon 托盘图标_置图标. 之后需调用 Modify 函数以应用修改.
//
// hIcon: 图标句柄. 可使用 wutil.HIcon 函数生成.
func (t *TrayIcon) SetIcon(hIcon uintptr) *TrayIcon {
	xc.XTrayIcon_SetIcon(hIcon)
	t.HIcon = hIcon
	t.Result = true
	return t
}

// SetTips 托盘图标_置提示文本. 之后需调用 Modify 函数以应用修改.
//
// pTips: 提示文本内容, 长度不能超过 127 个字符.
func (t *TrayIcon) SetTips(pTips string) *TrayIcon {
	xc.XTrayIcon_SetTips(pTips)
	t.Tips = pTips
	t.Result = true
	return t
}

// SetPopupBalloon 托盘图标_置弹出气泡. 设置弹出气泡信息. 之后需调用 Modify 函数以应用修改.
//
// pTitle: 弹出气泡标题.
//
// pText: 弹出气泡内容.
//
// hBalloonIcon: 气泡图标. 可填0.
//
// flags: 标识, 可设置默认图标类型, 禁用声音等: xcc.TrayIcon_Flag_
func (t *TrayIcon) SetPopupBalloon(pTitle, pText string, hBalloonIcon uintptr, flags xcc.TrayIcon_Flag_) *TrayIcon {
	xc.XTrayIcon_SetPopupBalloon(pTitle, pText, hBalloonIcon, flags)
	t.Result = true
	return t
}

// Reset 托盘图标_重置. 重置数据, 当未添加到系统托盘状态才可调用.
//   - 注意: 并不是清除本对象中存储的信息, 也并没有释放图标句柄.
func (t *TrayIcon) Reset() *TrayIcon {
	xc.XTrayIcon_Reset()
	t.Result = true
	return t
}

// AddEvent_TrayIcon 添加托盘图标事件.
//
// hWindow: 炫彩窗口句柄.
//
// fun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (t *TrayIcon) AddEvent_TrayIcon(hWindow int, fun xc.XWM_TRAYICON, allowAddingMultiple ...bool) int {
	return xc.WndEventHandler.AddCallBack(hWindow, xcc.XWM_TRAYICON, onXWM_TRAYICON, fun, allowAddingMultiple...)
}

// 托盘图标事件.
//
// hWindow: 炫彩窗口句柄.
func (t *TrayIcon) Event_TRAYICON(hWindow int, fun xc.XWM_TRAYICON) bool {
	return xc.XWnd_RegEventC(hWindow, xcc.XWM_TRAYICON, fun)
}

// 获取元素在slice中的索引, 找不到则返回-1.
func sliceIndexOf(slice []int32, element int32) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == element {
			return i
		}
	}
	return -1
}

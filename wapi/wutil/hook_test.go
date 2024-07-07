package wutil_test

import (
	"fmt"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/tf"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/wapi/wutil"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
	"testing"
)

func TestNewHookKeyboard(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		widget.NewShapeText(40, 40, 300, 30, "在任何窗口操作键盘都能够监听到", w.Handle)
		widget.NewEdit(40, 80, 300, 30, w.Handle).SetFocus()

		kbHook := wutil.NewHookKeyboard(func(nCode int32, wParam xcc.WM_, lParam *wapi.KBDLLHOOKSTRUCT) uintptr {
			if nCode < 0 { // nCode小于0时不应继续处理
				return wutil.CallNextHookEx_Keyboard(nCode, wParam, lParam)
			}

			if wParam == xcc.WM_KEYDOWN { // 键盘按下
				fmt.Printf("按键按下: 虚拟键码=%d, 扫描码=%d\n", lParam.VkCode, lParam.ScanCode)

				if lParam.VkCode == xcc.VK_A {
					fmt.Println("拦截了A键, 是不会输入文本框的")
					return 1 // 返回1可拦截, 这时按下A键是不会输入文本框的
				}
			} else if wParam == xcc.WM_KEYUP { // 键盘弹起
				fmt.Printf("按键弹起: 虚拟键码=%d, 扫描码=%d\n", lParam.VkCode, lParam.ScanCode)
			}
			return wutil.CallNextHookEx_Keyboard(nCode, wParam, lParam)
		})

		w.Event_CLOSE(func(pbHandled *bool) int {
			kbHook.Unhook()
			return 0
		})
	})
}

func TestNewHookMouse(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		widget.NewShapeText(40, 40, 300, 30, "在任何窗口操作鼠标都能够监听到", w.Handle)
		checkBtn := widget.NewButton(40, 80, 300, 30, "拦截鼠标右键按下消息", w.Handle).SetTypeEx(xcc.Button_Type_Check)
		checkBtn.EnableBkTransparent(true)

		// 注册事件_窗口鼠标右键按下, 用来检测是否真的拦截了鼠标右键按下消息
		w.Event_RBUTTONDOWN(func(nFlags uint, pPt *xc.POINT, pbHandled *bool) int {
			fmt.Println("响应了炫彩窗口鼠标右键被按下消息, 证明没有被拦截:", nFlags, pPt)
			return 0
		})

		msHook := wutil.NewHookMouse(func(nCode int32, wParam xcc.WM_, lParam *wapi.MSLLHOOKSTRUCT) uintptr {
			if nCode < 0 { // nCode小于0时不应继续处理
				return wutil.CallNextHookEx_Mouse(nCode, wParam, lParam)
			}

			switch wParam {
			case xcc.WM_LBUTTONDOWN: // 鼠标左键按下
				fmt.Println("鼠标左键按下, 坐标:", lParam.PT)
			case xcc.WM_RBUTTONDOWN: // 鼠标右键按下
				if checkBtn.GetStateEx() == xcc.Button_State_Check {
					fmt.Println("拦截了鼠标右键按下, 是不会真实响应鼠标右键消息的, 你在任务栏上右键已经没用了, 有些程序窗口拦截不了自行研究, 坐标:", lParam.PT)
					return 1 // 返回1可拦截, 这时按下鼠标右键是不会有响应的, 部分软件窗口拦截不了有多方面原因比如该程序做了特殊处理, 自行研究
				}
			case xcc.WM_MBUTTONDOWN: // 鼠标中键按下
				fmt.Println("鼠标中键按下, 坐标:", lParam.PT)
			case xcc.WM_XBUTTONDOWN: // 鼠标侧键按下
				value := wutil.GetHigh16Bits(lParam.MouseData)
				if value == 1 {
					fmt.Println("鼠标侧键1按下, 坐标:", lParam.PT)
				} else if value == 2 {
					fmt.Println("鼠标侧键2按下, 坐标:", lParam.PT)
				}
			case xcc.WM_MOUSEWHEEL: // 鼠标滚轮滚动
				value := wutil.GetHigh16Bits(lParam.MouseData)
				if lParam.MouseData > 0 {
					fmt.Printf("鼠标滚轮向上滚动, 滚轮增量: %d, 坐标:%v, lParam.MouseData: %v\n", value, lParam.PT, lParam.MouseData)
				} else if lParam.MouseData < 0 {
					fmt.Printf("鼠标滚轮向下滚动, 滚轮增量: %d, 坐标:%v, lParam.MouseData: %v\n", value, lParam.PT, lParam.MouseData)
				}
			}
			return wutil.CallNextHookEx_Mouse(nCode, wParam, lParam)
		})

		w.Event_CLOSE(func(pbHandled *bool) int {
			msHook.Unhook()
			return 0
		})
	})
}

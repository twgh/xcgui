package edge

import (
	"github.com/twgh/xcgui/wapi"
	"unsafe"
)

// WebView2_Objs 中声明了 WebView2_2 - WebView2_27,
// 默认值为 nil, 使用前需赋值.
//   - 可调用 InitWebView2_Objs 进行统一赋值.
//   - 在 WebView 关闭时会自动调用 ReleaseWebView2_Objs, 自己额外创建的 WebView2_Objs 得自己手动释放.
type WebView2_Objs struct {
	WebView2_2  *ICoreWebView2_2  // 默认值为 nil, 使用前需赋值, 会在 webview 关闭时自动释放
	WebView2_3  *ICoreWebView2_3  // 默认值为 nil, 使用前需赋值, 会在 webview 关闭时自动释放
	WebView2_4  *ICoreWebView2_4  // 默认值为 nil, 使用前需赋值, 会在 webview 关闭时自动释放
	WebView2_5  *ICoreWebView2_5  // 默认值为 nil, 使用前需赋值, 会在 webview 关闭时自动释放
	WebView2_6  *ICoreWebView2_6  // 默认值为 nil, 使用前需赋值, 会在 webview 关闭时自动释放
	WebView2_7  *ICoreWebView2_7  // 默认值为 nil, 使用前需赋值, 会在 webview 关闭时自动释放
	WebView2_8  *ICoreWebView2_8  // 默认值为 nil, 使用前需赋值, 会在 webview 关闭时自动释放
	WebView2_9  *ICoreWebView2_9  // 默认值为 nil, 使用前需赋值, 会在 webview 关闭时自动释放
	WebView2_10 *ICoreWebView2_10 // 默认值为 nil, 使用前需赋值, 会在 webview 关闭时自动释放
	WebView2_11 *ICoreWebView2_11 // 默认值为 nil, 使用前需赋值, 会在 webview 关闭时自动释放
	WebView2_12 *ICoreWebView2_12 // 默认值为 nil, 使用前需赋值, 会在 webview 关闭时自动释放
	WebView2_13 *ICoreWebView2_13 // 默认值为 nil, 使用前需赋值, 会在 webview 关闭时自动释放
	WebView2_14 *ICoreWebView2_14 // 默认值为 nil, 使用前需赋值, 会在 webview 关闭时自动释放
	WebView2_15 *ICoreWebView2_15 // 默认值为 nil, 使用前需赋值, 会在 webview 关闭时自动释放
	WebView2_16 *ICoreWebView2_16 // 默认值为 nil, 使用前需赋值, 会在 webview 关闭时自动释放
	WebView2_17 *ICoreWebView2_17 // 默认值为 nil, 使用前需赋值, 会在 webview 关闭时自动释放
	WebView2_18 *ICoreWebView2_18 // 默认值为 nil, 使用前需赋值, 会在 webview 关闭时自动释放
	WebView2_19 *ICoreWebView2_19 // 默认值为 nil, 使用前需赋值, 会在 webview 关闭时自动释放
	WebView2_20 *ICoreWebView2_20 // 默认值为 nil, 使用前需赋值, 会在 webview 关闭时自动释放
	WebView2_21 *ICoreWebView2_21 // 默认值为 nil, 使用前需赋值, 会在 webview 关闭时自动释放
	WebView2_22 *ICoreWebView2_22 // 默认值为 nil, 使用前需赋值, 会在 webview 关闭时自动释放
	WebView2_23 *ICoreWebView2_23 // 默认值为 nil, 使用前需赋值, 会在 webview 关闭时自动释放
	WebView2_24 *ICoreWebView2_24 // 默认值为 nil, 使用前需赋值, 会在 webview 关闭时自动释放
	WebView2_25 *ICoreWebView2_25 // 默认值为 nil, 使用前需赋值, 会在 webview 关闭时自动释放
	WebView2_26 *ICoreWebView2_26 // 默认值为 nil, 使用前需赋值, 会在 webview 关闭时自动释放
	WebView2_27 *ICoreWebView2_27 // 默认值为 nil, 使用前需赋值, 会在 webview 关闭时自动释放
}

// InitWebView2_Objs 初始化所有的 WebView2_ 系列对象.
//   - 从 2 到 27 按顺序获取, 出错时返回出错序号.
//   - 比如获取 WebView2_2 时失败返回 2, 获取 WebView2_3 时失败返回 3, 以此类推.
//   - 如果返回0, 代表成功获取到了 WebView2_2 到 WebView2_27.
//   - 这些对象, 只要运行时版本支持, 基本是不会获取出错的.
func (w *WebView2_Objs) InitWebView2_Objs(CoreWebView *ICoreWebView2) int {
	var err error
	if w.WebView2_2 == nil {
		w.WebView2_2, err = CoreWebView.GetICoreWebView2_2()
		if err != nil {
			return 2
		}
	}
	if w.WebView2_3 == nil {
		w.WebView2_3, err = CoreWebView.GetICoreWebView2_3()
		if err != nil {
			return 3
		}
	}
	if w.WebView2_4 == nil {
		w.WebView2_4, err = CoreWebView.GetICoreWebView2_4()
		if err != nil {
			return 4
		}
	}
	if w.WebView2_5 == nil {
		w.WebView2_5, err = CoreWebView.GetICoreWebView2_5()
		if err != nil {
			return 5
		}
	}
	if w.WebView2_6 == nil {
		w.WebView2_6, err = CoreWebView.GetICoreWebView2_6()
		if err != nil {
			return 6
		}
	}
	if w.WebView2_7 == nil {
		w.WebView2_7, err = CoreWebView.GetICoreWebView2_7()
		if err != nil {
			return 7
		}
	}
	if w.WebView2_8 == nil {
		w.WebView2_8, err = CoreWebView.GetICoreWebView2_8()
		if err != nil {
			return 8
		}
	}
	if w.WebView2_9 == nil {
		w.WebView2_9, err = CoreWebView.GetICoreWebView2_9()
		if err != nil {
			return 9
		}
	}
	if w.WebView2_10 == nil {
		w.WebView2_10, err = CoreWebView.GetICoreWebView2_10()
		if err != nil {
			return 10
		}
	}
	if w.WebView2_11 == nil {
		w.WebView2_11, err = CoreWebView.GetICoreWebView2_11()
		if err != nil {
			return 11
		}
	}
	if w.WebView2_12 == nil {
		w.WebView2_12, err = CoreWebView.GetICoreWebView2_12()
		if err != nil {
			return 12
		}
	}
	if w.WebView2_13 == nil {
		w.WebView2_13, err = CoreWebView.GetICoreWebView2_13()
		if err != nil {
			return 13
		}
	}
	if w.WebView2_14 == nil {
		w.WebView2_14, err = CoreWebView.GetICoreWebView2_14()
		if err != nil {
			return 14
		}
	}
	if w.WebView2_15 == nil {
		w.WebView2_15, err = CoreWebView.GetICoreWebView2_15()
		if err != nil {
			return 15
		}
	}
	if w.WebView2_16 == nil {
		w.WebView2_16, err = CoreWebView.GetICoreWebView2_16()
		if err != nil {
			return 16
		}
	}
	if w.WebView2_17 == nil {
		w.WebView2_17, err = CoreWebView.GetICoreWebView2_17()
		if err != nil {
			return 17
		}
	}
	if w.WebView2_18 == nil {
		w.WebView2_18, err = CoreWebView.GetICoreWebView2_18()
		if err != nil {
			return 18
		}
	}
	if w.WebView2_19 == nil {
		w.WebView2_19, err = CoreWebView.GetICoreWebView2_19()
		if err != nil {
			return 19
		}
	}
	if w.WebView2_20 == nil {
		w.WebView2_20, err = CoreWebView.GetICoreWebView2_20()
		if err != nil {
			return 20
		}
	}
	if w.WebView2_21 == nil {
		w.WebView2_21, err = CoreWebView.GetICoreWebView2_21()
		if err != nil {
			return 21
		}
	}
	if w.WebView2_22 == nil {
		w.WebView2_22, err = CoreWebView.GetICoreWebView2_22()
		if err != nil {
			return 22
		}
	}
	if w.WebView2_23 == nil {
		w.WebView2_23, err = CoreWebView.GetICoreWebView2_23()
		if err != nil {
			return 23
		}
	}
	if w.WebView2_24 == nil {
		w.WebView2_24, err = CoreWebView.GetICoreWebView2_24()
		if err != nil {
			return 24
		}
	}
	if w.WebView2_25 == nil {
		w.WebView2_25, err = CoreWebView.GetICoreWebView2_25()
		if err != nil {
			return 25
		}
	}
	if w.WebView2_26 == nil {
		w.WebView2_26, err = CoreWebView.GetICoreWebView2_26()
		if err != nil {
			return 26
		}
	}
	if w.WebView2_27 == nil {
		w.WebView2_27, err = CoreWebView.GetICoreWebView2_27()
		if err != nil {
			return 27
		}
	}
	return 0
}

// ReleaseWebView2_Objs 释放所有的 WebView2_ 系列对象.
func (w *WebView2_Objs) ReleaseWebView2_Objs() {
	if w.WebView2_2 != nil {
		w.WebView2_2.Release()
		w.WebView2_2 = nil
	}
	if w.WebView2_3 != nil {
		w.WebView2_3.Release()
		w.WebView2_3 = nil
	}
	if w.WebView2_4 != nil {
		w.WebView2_4.Release()
		w.WebView2_4 = nil
	}
	if w.WebView2_5 != nil {
		w.WebView2_5.Release()
		w.WebView2_5 = nil
	}
	if w.WebView2_6 != nil {
		w.WebView2_6.Release()
		w.WebView2_6 = nil
	}
	if w.WebView2_7 != nil {
		w.WebView2_7.Release()
		w.WebView2_7 = nil
	}
	if w.WebView2_8 != nil {
		w.WebView2_8.Release()
		w.WebView2_8 = nil
	}
	if w.WebView2_9 != nil {
		w.WebView2_9.Release()
		w.WebView2_9 = nil
	}
	if w.WebView2_10 != nil {
		w.WebView2_10.Release()
		w.WebView2_10 = nil
	}
	if w.WebView2_11 != nil {
		w.WebView2_11.Release()
		w.WebView2_11 = nil
	}
	if w.WebView2_12 != nil {
		w.WebView2_12.Release()
		w.WebView2_12 = nil
	}
	if w.WebView2_13 != nil {
		w.WebView2_13.Release()
		w.WebView2_13 = nil
	}
	if w.WebView2_14 != nil {
		w.WebView2_14.Release()
		w.WebView2_14 = nil
	}
	if w.WebView2_15 != nil {
		w.WebView2_15.Release()
		w.WebView2_15 = nil
	}
	if w.WebView2_16 != nil {
		w.WebView2_16.Release()
		w.WebView2_16 = nil
	}
	if w.WebView2_17 != nil {
		w.WebView2_17.Release()
		w.WebView2_17 = nil
	}
	if w.WebView2_18 != nil {
		w.WebView2_18.Release()
		w.WebView2_18 = nil
	}
	if w.WebView2_19 != nil {
		w.WebView2_19.Release()
		w.WebView2_19 = nil
	}
	if w.WebView2_20 != nil {
		w.WebView2_20.Release()
		w.WebView2_20 = nil
	}
	if w.WebView2_21 != nil {
		w.WebView2_21.Release()
		w.WebView2_21 = nil
	}
	if w.WebView2_22 != nil {
		w.WebView2_22.Release()
		w.WebView2_22 = nil
	}
	if w.WebView2_23 != nil {
		w.WebView2_23.Release()
		w.WebView2_23 = nil
	}
	if w.WebView2_24 != nil {
		w.WebView2_24.Release()
		w.WebView2_24 = nil
	}
	if w.WebView2_25 != nil {
		w.WebView2_25.Release()
		w.WebView2_25 = nil
	}
	if w.WebView2_26 != nil {
		w.WebView2_26.Release()
		w.WebView2_26 = nil
	}
	if w.WebView2_27 != nil {
		w.WebView2_27.Release()
		w.WebView2_27 = nil
	}
}

// GetICoreWebView2_2 获取 ICoreWebView2_2。
func (i *ICoreWebView2) GetICoreWebView2_2() (*ICoreWebView2_2, error) {
	var result *ICoreWebView2_2
	iidICoreWebView2_2 := wapi.NewGUID(IID_ICoreWebView2_2)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2_2),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetICoreWebView2_2 获取 ICoreWebView2_2。出错时会触发全局错误回调。
func (i *ICoreWebView2) MustGetICoreWebView2_2() *ICoreWebView2_2 {
	result, err := i.GetICoreWebView2_2()
	ReportErrorAtuo(err)
	return result
}

// GetICoreWebView2_3 获取 ICoreWebView2_3。
func (i *ICoreWebView2) GetICoreWebView2_3() (*ICoreWebView2_3, error) {
	var result *ICoreWebView2_3
	iidICoreWebView2_3 := wapi.NewGUID(IID_ICoreWebView2_3)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2_3),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetICoreWebView2_3 获取 ICoreWebView2_3。出错时会触发全局错误回调。
func (i *ICoreWebView2) MustGetICoreWebView2_3() *ICoreWebView2_3 {
	result, err := i.GetICoreWebView2_3()
	ReportErrorAtuo(err)
	return result
}

// GetICoreWebView2_4 获取 ICoreWebView2_4。
func (i *ICoreWebView2) GetICoreWebView2_4() (*ICoreWebView2_4, error) {
	var result *ICoreWebView2_4
	iidICoreWebView2_4 := wapi.NewGUID(IID_ICoreWebView2_4)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2_4),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetICoreWebView2_4 获取 ICoreWebView2_4。出错时会触发全局错误回调。
func (i *ICoreWebView2) MustGetICoreWebView2_4() *ICoreWebView2_4 {
	result, err := i.GetICoreWebView2_4()
	ReportErrorAtuo(err)
	return result
}

// GetICoreWebView2_5 获取 ICoreWebView2_5。
func (i *ICoreWebView2) GetICoreWebView2_5() (*ICoreWebView2_5, error) {
	var result *ICoreWebView2_5
	iidICoreWebView2_5 := wapi.NewGUID(IID_ICoreWebView2_5)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2_5),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetICoreWebView2_5 获取 ICoreWebView2_5。出错时会触发全局错误回调。
func (i *ICoreWebView2) MustGetICoreWebView2_5() *ICoreWebView2_5 {
	result, err := i.GetICoreWebView2_5()
	ReportErrorAtuo(err)
	return result
}

// GetICoreWebView2_6 获取 ICoreWebView2_6。
func (i *ICoreWebView2) GetICoreWebView2_6() (*ICoreWebView2_6, error) {
	var result *ICoreWebView2_6
	iidICoreWebView2_6 := wapi.NewGUID(IID_ICoreWebView2_6)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2_6),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetICoreWebView2_6 获取 ICoreWebView2_6。出错时会触发全局错误回调。
func (i *ICoreWebView2) MustGetICoreWebView2_6() *ICoreWebView2_6 {
	result, err := i.GetICoreWebView2_6()
	ReportErrorAtuo(err)
	return result
}

// GetICoreWebView2_7 获取 ICoreWebView2_7。
func (i *ICoreWebView2) GetICoreWebView2_7() (*ICoreWebView2_7, error) {
	var result *ICoreWebView2_7
	iidICoreWebView2_7 := wapi.NewGUID(IID_ICoreWebView2_7)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2_7),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetICoreWebView2_7 获取 ICoreWebView2_7。出错时会触发全局错误回调。
func (i *ICoreWebView2) MustGetICoreWebView2_7() *ICoreWebView2_7 {
	result, err := i.GetICoreWebView2_7()
	ReportErrorAtuo(err)
	return result
}

// GetICoreWebView2_8 获取 ICoreWebView2_8。
func (i *ICoreWebView2) GetICoreWebView2_8() (*ICoreWebView2_8, error) {
	var result *ICoreWebView2_8
	iidICoreWebView2_8 := wapi.NewGUID(IID_ICoreWebView2_8)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2_8),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetICoreWebView2_8 获取 ICoreWebView2_8。出错时会触发全局错误回调。
func (i *ICoreWebView2) MustGetICoreWebView2_8() *ICoreWebView2_8 {
	result, err := i.GetICoreWebView2_8()
	ReportErrorAtuo(err)
	return result
}

// GetICoreWebView2_9 获取 ICoreWebView2_9。
func (i *ICoreWebView2) GetICoreWebView2_9() (*ICoreWebView2_9, error) {
	var result *ICoreWebView2_9
	iidICoreWebView2_9 := wapi.NewGUID(IID_ICoreWebView2_9)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2_9),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetICoreWebView2_9 获取 ICoreWebView2_9。出错时会触发全局错误回调。
func (i *ICoreWebView2) MustGetICoreWebView2_9() *ICoreWebView2_9 {
	result, err := i.GetICoreWebView2_9()
	ReportErrorAtuo(err)
	return result
}

// GetICoreWebView2_10 获取 ICoreWebView2_10。
func (i *ICoreWebView2) GetICoreWebView2_10() (*ICoreWebView2_10, error) {
	var result *ICoreWebView2_10
	iidICoreWebView2_10 := wapi.NewGUID(IID_ICoreWebView2_10)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2_10),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetICoreWebView2_10 获取 ICoreWebView2_10。出错时会触发全局错误回调。
func (i *ICoreWebView2) MustGetICoreWebView2_10() *ICoreWebView2_10 {
	result, err := i.GetICoreWebView2_10()
	ReportErrorAtuo(err)
	return result
}

// GetICoreWebView2_11 获取 ICoreWebView2_11。
func (i *ICoreWebView2) GetICoreWebView2_11() (*ICoreWebView2_11, error) {
	var result *ICoreWebView2_11
	iidICoreWebView2_11 := wapi.NewGUID(IID_ICoreWebView2_11)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2_11),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetICoreWebView2_11 获取 ICoreWebView2_11。出错时会触发全局错误回调。
func (i *ICoreWebView2) MustGetICoreWebView2_11() *ICoreWebView2_11 {
	result, err := i.GetICoreWebView2_11()
	ReportErrorAtuo(err)
	return result
}

// GetICoreWebView2_12 获取 ICoreWebView2_12。
func (i *ICoreWebView2) GetICoreWebView2_12() (*ICoreWebView2_12, error) {
	var result *ICoreWebView2_12
	iidICoreWebView2_12 := wapi.NewGUID(IID_ICoreWebView2_12)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2_12),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetICoreWebView2_12 获取 ICoreWebView2_12。出错时会触发全局错误回调。
func (i *ICoreWebView2) MustGetICoreWebView2_12() *ICoreWebView2_12 {
	result, err := i.GetICoreWebView2_12()
	ReportErrorAtuo(err)
	return result
}

// GetICoreWebView2_13 获取 ICoreWebView2_13。
func (i *ICoreWebView2) GetICoreWebView2_13() (*ICoreWebView2_13, error) {
	var result *ICoreWebView2_13
	iidICoreWebView2_13 := wapi.NewGUID(IID_ICoreWebView2_13)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2_13),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetICoreWebView2_13 获取 ICoreWebView2_13。出错时会触发全局错误回调。
func (i *ICoreWebView2) MustGetICoreWebView2_13() *ICoreWebView2_13 {
	result, err := i.GetICoreWebView2_13()
	ReportErrorAtuo(err)
	return result
}

// GetICoreWebView2_14 获取 ICoreWebView2_14。
func (i *ICoreWebView2) GetICoreWebView2_14() (*ICoreWebView2_14, error) {
	var result *ICoreWebView2_14
	iidICoreWebView2_14 := wapi.NewGUID(IID_ICoreWebView2_14)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2_14),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetICoreWebView2_14 获取 ICoreWebView2_14。出错时会触发全局错误回调。
func (i *ICoreWebView2) MustGetICoreWebView2_14() *ICoreWebView2_14 {
	result, err := i.GetICoreWebView2_14()
	ReportErrorAtuo(err)
	return result
}

// GetICoreWebView2_15 获取 ICoreWebView2_15。
func (i *ICoreWebView2) GetICoreWebView2_15() (*ICoreWebView2_15, error) {
	var result *ICoreWebView2_15
	iidICoreWebView2_15 := wapi.NewGUID(IID_ICoreWebView2_15)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2_15),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetICoreWebView2_15 获取 ICoreWebView2_15。出错时会触发全局错误回调。
func (i *ICoreWebView2) MustGetICoreWebView2_15() *ICoreWebView2_15 {
	result, err := i.GetICoreWebView2_15()
	ReportErrorAtuo(err)
	return result
}

// GetICoreWebView2_16 获取 ICoreWebView2_16。
func (i *ICoreWebView2) GetICoreWebView2_16() (*ICoreWebView2_16, error) {
	var result *ICoreWebView2_16
	iidICoreWebView2_16 := wapi.NewGUID(IID_ICoreWebView2_16)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2_16),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetICoreWebView2_16 获取 ICoreWebView2_16。出错时会触发全局错误回调。
func (i *ICoreWebView2) MustGetICoreWebView2_16() *ICoreWebView2_16 {
	result, err := i.GetICoreWebView2_16()
	ReportErrorAtuo(err)
	return result
}

// GetICoreWebView2_17 获取 ICoreWebView2_17。
func (i *ICoreWebView2) GetICoreWebView2_17() (*ICoreWebView2_17, error) {
	var result *ICoreWebView2_17
	iidICoreWebView2_17 := wapi.NewGUID(IID_ICoreWebView2_17)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2_17),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetICoreWebView2_17 获取 ICoreWebView2_17。出错时会触发全局错误回调。
func (i *ICoreWebView2) MustGetICoreWebView2_17() *ICoreWebView2_17 {
	result, err := i.GetICoreWebView2_17()
	ReportErrorAtuo(err)
	return result
}

// GetICoreWebView2_18 获取 ICoreWebView2_18。
func (i *ICoreWebView2) GetICoreWebView2_18() (*ICoreWebView2_18, error) {
	var result *ICoreWebView2_18
	iidICoreWebView2_18 := wapi.NewGUID(IID_ICoreWebView2_18)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2_18),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetICoreWebView2_18 获取 ICoreWebView2_18。出错时会触发全局错误回调。
func (i *ICoreWebView2) MustGetICoreWebView2_18() *ICoreWebView2_18 {
	result, err := i.GetICoreWebView2_18()
	ReportErrorAtuo(err)
	return result
}

// GetICoreWebView2_19 获取 ICoreWebView2_19。
func (i *ICoreWebView2) GetICoreWebView2_19() (*ICoreWebView2_19, error) {
	var result *ICoreWebView2_19
	iidICoreWebView2_19 := wapi.NewGUID(IID_ICoreWebView2_19)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2_19),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetICoreWebView2_19 获取 ICoreWebView2_19。出错时会触发全局错误回调。
func (i *ICoreWebView2) MustGetICoreWebView2_19() *ICoreWebView2_19 {
	result, err := i.GetICoreWebView2_19()
	ReportErrorAtuo(err)
	return result
}

// GetICoreWebView2_20 获取 ICoreWebView2_20。
func (i *ICoreWebView2) GetICoreWebView2_20() (*ICoreWebView2_20, error) {
	var result *ICoreWebView2_20
	iidICoreWebView2_20 := wapi.NewGUID(IID_ICoreWebView2_20)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2_20),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetICoreWebView2_20 获取 ICoreWebView2_20。出错时会触发全局错误回调。
func (i *ICoreWebView2) MustGetICoreWebView2_20() *ICoreWebView2_20 {
	result, err := i.GetICoreWebView2_20()
	ReportErrorAtuo(err)
	return result
}

// GetICoreWebView2_21 获取 ICoreWebView2_21。
func (i *ICoreWebView2) GetICoreWebView2_21() (*ICoreWebView2_21, error) {
	var result *ICoreWebView2_21
	iidICoreWebView2_21 := wapi.NewGUID(IID_ICoreWebView2_21)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2_21),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetICoreWebView2_21 获取 ICoreWebView2_21。出错时会触发全局错误回调。
func (i *ICoreWebView2) MustGetICoreWebView2_21() *ICoreWebView2_21 {
	result, err := i.GetICoreWebView2_21()
	ReportErrorAtuo(err)
	return result
}

// GetICoreWebView2_22 获取 ICoreWebView2_22。
func (i *ICoreWebView2) GetICoreWebView2_22() (*ICoreWebView2_22, error) {
	var result *ICoreWebView2_22
	iidICoreWebView2_22 := wapi.NewGUID(IID_ICoreWebView2_22)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2_22),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetICoreWebView2_22 获取 ICoreWebView2_22。出错时会触发全局错误回调。
func (i *ICoreWebView2) MustGetICoreWebView2_22() *ICoreWebView2_22 {
	result, err := i.GetICoreWebView2_22()
	ReportErrorAtuo(err)
	return result
}

// GetICoreWebView2_23 获取 ICoreWebView2_23。
func (i *ICoreWebView2) GetICoreWebView2_23() (*ICoreWebView2_23, error) {
	var result *ICoreWebView2_23
	iidICoreWebView2_23 := wapi.NewGUID(IID_ICoreWebView2_23)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2_23),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetICoreWebView2_23 获取 ICoreWebView2_23。出错时会触发全局错误回调。
func (i *ICoreWebView2) MustGetICoreWebView2_23() *ICoreWebView2_23 {
	result, err := i.GetICoreWebView2_23()
	ReportErrorAtuo(err)
	return result
}

// GetICoreWebView2_24 获取 ICoreWebView2_24。
func (i *ICoreWebView2) GetICoreWebView2_24() (*ICoreWebView2_24, error) {
	var result *ICoreWebView2_24
	iidICoreWebView2_24 := wapi.NewGUID(IID_ICoreWebView2_24)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2_24),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetICoreWebView2_24 获取 ICoreWebView2_24。出错时会触发全局错误回调。
func (i *ICoreWebView2) MustGetICoreWebView2_24() *ICoreWebView2_24 {
	result, err := i.GetICoreWebView2_24()
	ReportErrorAtuo(err)
	return result
}

// GetICoreWebView2_25 获取 ICoreWebView2_25。
func (i *ICoreWebView2) GetICoreWebView2_25() (*ICoreWebView2_25, error) {
	var result *ICoreWebView2_25
	iidICoreWebView2_25 := wapi.NewGUID(IID_ICoreWebView2_25)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2_25),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetICoreWebView2_25 获取 ICoreWebView2_25。出错时会触发全局错误回调。
func (i *ICoreWebView2) MustGetICoreWebView2_25() *ICoreWebView2_25 {
	result, err := i.GetICoreWebView2_25()
	ReportErrorAtuo(err)
	return result
}

// GetICoreWebView2_26 获取 ICoreWebView2_26。
func (i *ICoreWebView2) GetICoreWebView2_26() (*ICoreWebView2_26, error) {
	var result *ICoreWebView2_26
	iidICoreWebView2_26 := wapi.NewGUID(IID_ICoreWebView2_26)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2_26),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetICoreWebView2_26 获取 ICoreWebView2_26。出错时会触发全局错误回调。
func (i *ICoreWebView2) MustGetICoreWebView2_26() *ICoreWebView2_26 {
	result, err := i.GetICoreWebView2_26()
	ReportErrorAtuo(err)
	return result
}

// GetICoreWebView2_27 获取 ICoreWebView2_27。
func (i *ICoreWebView2) GetICoreWebView2_27() (*ICoreWebView2_27, error) {
	var result *ICoreWebView2_27
	iidICoreWebView2_27 := wapi.NewGUID(IID_ICoreWebView2_27)
	err := i.QueryInterface(
		unsafe.Pointer(iidICoreWebView2_27),
		unsafe.Pointer(&result))
	return result, err
}

// MustGetICoreWebView2_27 获取 ICoreWebView2_27。出错时会触发全局错误回调。
func (i *ICoreWebView2) MustGetICoreWebView2_27() *ICoreWebView2_27 {
	result, err := i.GetICoreWebView2_27()
	ReportErrorAtuo(err)
	return result
}

package edge

// Webview2_Objs 中声明了 WebView2_2 - WebView2_27,
// 默认值为 nil, 使用前需赋值, 或调用一次 Webview2.InitAllWebView2_Objs() 进行统一赋值, 在 webview 关闭时会自动释放所有的 WebView2_ 系列对象.
type Webview2_Objs struct {
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

// InitAllWebView2_Objs 初始化所有的 WebView2_ 系列对象, 给本类 Webview2_Objs 中的成员赋值.
//   - 从 2 到 27 按顺序获取, 出错时返回出错序号.
//   - 比如获取 WebView2_2 时失败返回 2, 获取 WebView2_3 时失败返回 3, 以此类推.
//   - 如果返回0, 代表成功获取到了 WebView2_2 到 WebView2_27.
//   - 这些对象, 只要运行时版本支持, 基本是不会获取出错的.
func (e *Webview2) InitAllWebView2_Objs() int {
	var err error
	if e.WebView2_2 == nil {
		e.WebView2_2, err = e.CoreWebView.GetICoreWebView2_2()
		if err != nil {
			return 2
		}
	}
	if e.WebView2_3 == nil {
		e.WebView2_3, err = e.CoreWebView.GetICoreWebView2_3()
		if err != nil {
			return 3
		}
	}
	if e.WebView2_4 == nil {
		e.WebView2_4, err = e.CoreWebView.GetICoreWebView2_4()
		if err != nil {
			return 4
		}
	}
	if e.WebView2_5 == nil {
		e.WebView2_5, err = e.CoreWebView.GetICoreWebView2_5()
		if err != nil {
			return 5
		}
	}
	if e.WebView2_6 == nil {
		e.WebView2_6, err = e.CoreWebView.GetICoreWebView2_6()
		if err != nil {
			return 6
		}
	}
	if e.WebView2_7 == nil {
		e.WebView2_7, err = e.CoreWebView.GetICoreWebView2_7()
		if err != nil {
			return 7
		}
	}
	if e.WebView2_8 == nil {
		e.WebView2_8, err = e.CoreWebView.GetICoreWebView2_8()
		if err != nil {
			return 8
		}
	}
	if e.WebView2_9 == nil {
		e.WebView2_9, err = e.CoreWebView.GetICoreWebView2_9()
		if err != nil {
			return 9
		}
	}
	if e.WebView2_10 == nil {
		e.WebView2_10, err = e.CoreWebView.GetICoreWebView2_10()
		if err != nil {
			return 10
		}
	}
	if e.WebView2_11 == nil {
		e.WebView2_11, err = e.CoreWebView.GetICoreWebView2_11()
		if err != nil {
			return 11
		}
	}
	if e.WebView2_12 == nil {
		e.WebView2_12, err = e.CoreWebView.GetICoreWebView2_12()
		if err != nil {
			return 12
		}
	}
	if e.WebView2_13 == nil {
		e.WebView2_13, err = e.CoreWebView.GetICoreWebView2_13()
		if err != nil {
			return 13
		}
	}
	if e.WebView2_14 == nil {
		e.WebView2_14, err = e.CoreWebView.GetICoreWebView2_14()
		if err != nil {
			return 14
		}
	}
	if e.WebView2_15 == nil {
		e.WebView2_15, err = e.CoreWebView.GetICoreWebView2_15()
		if err != nil {
			return 15
		}
	}
	if e.WebView2_16 == nil {
		e.WebView2_16, err = e.CoreWebView.GetICoreWebView2_16()
		if err != nil {
			return 16
		}
	}
	if e.WebView2_17 == nil {
		e.WebView2_17, err = e.CoreWebView.GetICoreWebView2_17()
		if err != nil {
			return 17
		}
	}
	if e.WebView2_18 == nil {
		e.WebView2_18, err = e.CoreWebView.GetICoreWebView2_18()
		if err != nil {
			return 18
		}
	}
	if e.WebView2_19 == nil {
		e.WebView2_19, err = e.CoreWebView.GetICoreWebView2_19()
		if err != nil {
			return 19
		}
	}
	if e.WebView2_20 == nil {
		e.WebView2_20, err = e.CoreWebView.GetICoreWebView2_20()
		if err != nil {
			return 20
		}
	}
	if e.WebView2_21 == nil {
		e.WebView2_21, err = e.CoreWebView.GetICoreWebView2_21()
		if err != nil {
			return 21
		}
	}
	if e.WebView2_22 == nil {
		e.WebView2_22, err = e.CoreWebView.GetICoreWebView2_22()
		if err != nil {
			return 22
		}
	}
	if e.WebView2_23 == nil {
		e.WebView2_23, err = e.CoreWebView.GetICoreWebView2_23()
		if err != nil {
			return 23
		}
	}
	if e.WebView2_24 == nil {
		e.WebView2_24, err = e.CoreWebView.GetICoreWebView2_24()
		if err != nil {
			return 24
		}
	}
	if e.WebView2_25 == nil {
		e.WebView2_25, err = e.CoreWebView.GetICoreWebView2_25()
		if err != nil {
			return 25
		}
	}
	if e.WebView2_26 == nil {
		e.WebView2_26, err = e.CoreWebView.GetICoreWebView2_26()
		if err != nil {
			return 26
		}
	}
	if e.WebView2_27 == nil {
		e.WebView2_27, err = e.CoreWebView.GetICoreWebView2_27()
		if err != nil {
			return 27
		}
	}
	return 0
}

// 释放所有的 WebView2_ 系列对象. 会在使用 Close 时自动调用.
func (e *Webview2) releaseAllWebView2_Objs() {
	if e.WebView2_2 != nil {
		e.WebView2_2.Release()
		e.WebView2_2 = nil
	}
	if e.WebView2_3 != nil {
		e.WebView2_3.Release()
		e.WebView2_3 = nil
	}
	if e.WebView2_4 != nil {
		e.WebView2_4.Release()
		e.WebView2_4 = nil
	}
	if e.WebView2_5 != nil {
		e.WebView2_5.Release()
		e.WebView2_5 = nil
	}
	if e.WebView2_6 != nil {
		e.WebView2_6.Release()
		e.WebView2_6 = nil
	}
	if e.WebView2_7 != nil {
		e.WebView2_7.Release()
		e.WebView2_7 = nil
	}
	if e.WebView2_8 != nil {
		e.WebView2_8.Release()
		e.WebView2_8 = nil
	}
	if e.WebView2_9 != nil {
		e.WebView2_9.Release()
		e.WebView2_9 = nil
	}
	if e.WebView2_10 != nil {
		e.WebView2_10.Release()
		e.WebView2_10 = nil
	}
	if e.WebView2_11 != nil {
		e.WebView2_11.Release()
		e.WebView2_11 = nil
	}
	if e.WebView2_12 != nil {
		e.WebView2_12.Release()
		e.WebView2_12 = nil
	}
	if e.WebView2_13 != nil {
		e.WebView2_13.Release()
		e.WebView2_13 = nil
	}
	if e.WebView2_14 != nil {
		e.WebView2_14.Release()
		e.WebView2_14 = nil
	}
	if e.WebView2_15 != nil {
		e.WebView2_15.Release()
		e.WebView2_15 = nil
	}
	if e.WebView2_16 != nil {
		e.WebView2_16.Release()
		e.WebView2_16 = nil
	}
	if e.WebView2_17 != nil {
		e.WebView2_17.Release()
		e.WebView2_17 = nil
	}
	if e.WebView2_18 != nil {
		e.WebView2_18.Release()
		e.WebView2_18 = nil
	}
	if e.WebView2_19 != nil {
		e.WebView2_19.Release()
		e.WebView2_19 = nil
	}
	if e.WebView2_20 != nil {
		e.WebView2_20.Release()
		e.WebView2_20 = nil
	}
	if e.WebView2_21 != nil {
		e.WebView2_21.Release()
		e.WebView2_21 = nil
	}
	if e.WebView2_22 != nil {
		e.WebView2_22.Release()
		e.WebView2_22 = nil
	}
	if e.WebView2_23 != nil {
		e.WebView2_23.Release()
		e.WebView2_23 = nil
	}
	if e.WebView2_24 != nil {
		e.WebView2_24.Release()
		e.WebView2_24 = nil
	}
	if e.WebView2_25 != nil {
		e.WebView2_25.Release()
		e.WebView2_25 = nil
	}
	if e.WebView2_26 != nil {
		e.WebView2_26.Release()
		e.WebView2_26 = nil
	}
	if e.WebView2_27 != nil {
		e.WebView2_27.Release()
		e.WebView2_27 = nil
	}

	if e.CoreWebView != nil {
		e.CoreWebView.Release()
		e.CoreWebView = nil
	}
}

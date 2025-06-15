package edge

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

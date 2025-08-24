package edge

// webview 选项函数类型
type WebViewOption func(*WebViewOptions)

// WebViewOptions 选项.
type WebViewOptions struct {
	// ProfileName 配置文件名称, 可不填。
	//   - 它的最大长度为64个字符（不包括空字符终止符）。它不区分 ASCII 大小写。
	//   - 注意：文本不能以句号“.”或空格“ ”结尾。
	//   - 此外，虽然允许使用大写字母，但它们会被当作小写字母处理，因为配置文件名称将映射到磁盘上真实的配置文件目录路径，而 Windows 文件系统在处理路径名称时不区分大小写。
	ProfileName string
	// webview 宿主窗口标题
	Title string
	// webview 宿主窗口类名
	ClassName string

	// 资源文件中的图标资源 id, 给原生窗口设置图标, 如果为 0, 则使用默认图标.
	IconId uint

	WebViewSize

	// 填充父, 如果为 true, 则 webview 会填充父, WebViewSize 里的固定坐标和尺寸会失效.
	FillParent bool

	// Debug 是否可打开开发者工具.
	Debug bool
	// AppDrag 是否启用非客户区域支持。WebView2 运行时版本大于 123.0.2420.47 时有效。
	//   - 当此属性为 true 时，将启用所有非客户端区域功能：将启用可拖动区域，它们是网页上用 CSS 属性 app-region:drag/no-drag 标记的区域。
	//   - 设置为拖动时，这些区域将被视为窗口的标题栏，支持拖动整个 webview 及其宿主应用程序窗口；
	//   - 系统菜单在右键单击时显示，双击将触发最大化/恢复窗口大小。
	AppDrag bool

	// AutoFocus 将在窗口获得焦点时尝试保持 WebView 的焦点。
	AutoFocus bool
	// PrivateMode 是否启用隐私模式。
	PrivateMode bool

	DefaultEnabledWebViewOption
}

// WebViewSize 是 webview 的固定位置与大小.
type WebViewSize struct {
	Left   int32
	Top    int32
	Width  int32
	Height int32
}

// DefaultEnabledWebViewOption 里的 webview 选项都是默认会开启的.
type DefaultEnabledWebViewOption struct {
	// 默认的上下文菜单
	DefaultContextMenus bool
	// 状态栏
	StatusBar bool
	// 缩放控件, 控制是否可以缩放
	ZoomControl bool
	// 浏览器快捷键, 是浏览器里面默认的一些快捷键, 组合键
	BrowserAcceleratorKeys bool
}

// WebView 默认选项
func defaultWebViewOptions() *WebViewOptions {
	return &WebViewOptions{
		DefaultEnabledWebViewOption: DefaultEnabledWebViewOption{
			DefaultContextMenus:    true,
			StatusBar:              true,
			ZoomControl:            true,
			BrowserAcceleratorKeys: true,
		},
	}
}

// WithProfileName 设置配置文件名称.
//   - 它的最大长度为64个字符（不包括空字符终止符）。它不区分 ASCII 大小写。
//   - 注意：文本不能以句号"."或空格" "结尾。
//   - 此外，虽然允许使用大写字母，但它们会被当作小写字母处理，因为配置文件名称将映射到磁盘上真实的配置文件目录路径，而 Windows 文件系统在处理路径名称时不区分大小写。
func WithProfileName(name string) WebViewOption {
	return func(o *WebViewOptions) {
		o.ProfileName = name
	}
}

// WithTitle 设置 webview 宿主窗口标题
func WithTitle(title string) WebViewOption {
	return func(o *WebViewOptions) {
		o.Title = title
	}
}

// WithClassName 设置 webview 宿主窗口类名
func WithClassName(className string) WebViewOption {
	return func(o *WebViewOptions) {
		o.ClassName = className
	}
}

// WithIconId 设置资源文件中的图标资源id, 给原生窗口设置图标, 如果为0, 则使用默认图标.
func WithIconId(id uint) WebViewOption {
	return func(o *WebViewOptions) {
		o.IconId = id
	}
}

// WithWebViewSize 设置 webview 固定位置与大小.
func WithWebViewSize(left, top, width, height int32) WebViewOption {
	return func(o *WebViewOptions) {
		o.WebViewSize = WebViewSize{
			Left:   left,
			Top:    top,
			Width:  width,
			Height: height,
		}
	}
}

// WithFillParent 设置是否填充父, 如果为 true, 则 webview 会填充父,  WebViewSize 里的固定坐标和尺寸会失效.
func WithFillParent(enable bool) WebViewOption {
	return func(o *WebViewOptions) {
		o.FillParent = enable
	}
}

// WithDebug 设置是否可打开开发者工具.
func WithDebug(enable bool) WebViewOption {
	return func(o *WebViewOptions) {
		o.Debug = enable
	}
}

// WithAppDrag 设置是否启用非客户区域支持。 WebView2 运行时版本大于 123.0.2420.47 时有效。
//   - 当此属性为 true 时，将启用所有非客户端区域功能：将启用可拖动区域，它们是网页上用 CSS 属性 app-region:drag/no-drag 标记的区域。
//   - 设置为拖动时，这些区域将被视为窗口的标题栏，支持拖动整个 webview 及其宿主应用程序窗口；
//   - 系统菜单在右键单击时显示，双击将触发最大化/恢复窗口大小。
func WithAppDrag(enable bool) WebViewOption {
	return func(o *WebViewOptions) {
		o.AppDrag = enable
	}
}

// WithAutoFocus 设置是否在窗口获得焦点时尝试保持 WebView 的焦点.
func WithAutoFocus(enable bool) WebViewOption {
	return func(o *WebViewOptions) {
		o.AutoFocus = enable
	}
}

// WithPrivateMode 设置是否启用隐私模式.
func WithPrivateMode(enable bool) WebViewOption {
	return func(o *WebViewOptions) {
		o.PrivateMode = enable
	}
}

// WithDefaultContextMenus 设置是否启用默认的上下文菜单.
func WithDefaultContextMenus(enable bool) WebViewOption {
	return func(o *WebViewOptions) {
		o.DefaultContextMenus = enable
	}
}

// WithStatusBar 设置是否启用状态栏.
func WithStatusBar(enable bool) WebViewOption {
	return func(o *WebViewOptions) {
		o.StatusBar = enable
	}
}

// WithZoomControl 设置是否启用缩放控件, 控制是否可以缩放.
func WithZoomControl(enable bool) WebViewOption {
	return func(o *WebViewOptions) {
		o.ZoomControl = enable
	}
}

// WithBrowserAcceleratorKeys 设置是否启用浏览器快捷键, 是浏览器里面默认的一些快捷键, 组合键.
func WithBrowserAcceleratorKeys(enable bool) WebViewOption {
	return func(o *WebViewOptions) {
		o.BrowserAcceleratorKeys = enable
	}
}

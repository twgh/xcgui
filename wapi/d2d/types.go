package d2d

// 本文件是 Direct2D 里"按值/按指针传递的简单数据类型"的 Go 形式, 命名与 Windows SDK 保持一致.
//
// https://learn.microsoft.com/zh-cn/windows/win32/direct2d/direct2d-reference

// D2D1_COLOR_F 颜色(分量 0~1, 顺序 RGBA), 按指针传递.
//
// https://learn.microsoft.com/zh-cn/windows/win32/direct2d/d2d1-color-f
type D2D1_COLOR_F struct {
	R, G, B, A float32
}

// NewColorF 从 0~255 的 RGBA 分量构造颜色.
func NewColorF(r, g, b, a uint8) D2D1_COLOR_F {
	const k = 1.0 / 255.0
	return D2D1_COLOR_F{R: float32(r) * k, G: float32(g) * k, B: float32(b) * k, A: float32(a) * k}
}

// D2D1_POINT_2F 二维点, 按值传递(调用时需打包成一个 64 位整数, 见 pointArg).
//
// https://learn.microsoft.com/zh-cn/windows/win32/direct2d/d2d1-point-2f
type D2D1_POINT_2F struct {
	X, Y float32
}

// D2D1_RECT_F 矩形, 按指针传递.
//
// https://learn.microsoft.com/zh-cn/windows/win32/direct2d/d2d1-rect-f
type D2D1_RECT_F struct {
	Left, Top, Right, Bottom float32
}

// D2D1_ELLIPSE 椭圆: 中心点 + 两个半径, 按指针传递(DrawEllipse / FillEllipse / CreateEllipseGeometry).
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/ns-d2d1-d2d1_ellipse
type D2D1_ELLIPSE struct {
	Point   D2D1_POINT_2F
	RadiusX float32
	RadiusY float32
}

// NewEllipse 从中心与半径构造椭圆.
func NewEllipse(cx, cy, rx, ry float32) D2D1_ELLIPSE {
	return D2D1_ELLIPSE{Point: D2D1_POINT_2F{X: cx, Y: cy}, RadiusX: rx, RadiusY: ry}
}

// D2D1_ROUNDED_RECT 圆角矩形, 按指针传递(CreateRoundedRectangleGeometry).
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/ns-d2d1-d2d1_rounded_rect
type D2D1_ROUNDED_RECT struct {
	Rect    D2D1_RECT_F
	RadiusX float32
	RadiusY float32
}

// D2D1_MATRIX_3X2_F 是 3x2 仿射变换矩阵, 按指针传递:
//
//	| M11 M12 |
//	| M21 M22 |
//	| DX  DY  |
//
// https://learn.microsoft.com/zh-cn/windows/win32/direct2d/d2d1-matrix-3x2-f
type D2D1_MATRIX_3X2_F struct {
	M11, M12 float32
	M21, M22 float32
	DX, DY   float32
}

// D2D1_CAP_STYLE 线帽样式.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/ne-d2d1-d2d1_cap_style
type D2D1_CAP_STYLE int32

const (
	D2D1_CAP_STYLE_FLAT D2D1_CAP_STYLE = iota
	D2D1_CAP_STYLE_SQUARE
	D2D1_CAP_STYLE_ROUND
	D2D1_CAP_STYLE_TRIANGLE
)

// D2D1_DASH_CAP_STYLE 虚线端点样式(取值同 D2D1_CAP_STYLE).
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/ne-d2d1-d2d1_dash_cap_style
type D2D1_DASH_CAP_STYLE int32

const (
	D2D1_DASH_CAP_STYLE_FLAT D2D1_DASH_CAP_STYLE = iota
	D2D1_DASH_CAP_STYLE_SQUARE
	D2D1_DASH_CAP_STYLE_ROUND
	D2D1_DASH_CAP_STYLE_TRIANGLE
)

// D2D1_LINE_JOIN 连接样式.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/ne-d2d1-d2d1_line_join
type D2D1_LINE_JOIN int32

const (
	D2D1_LINE_JOIN_MITER D2D1_LINE_JOIN = iota
	D2D1_LINE_JOIN_BEVEL
	D2D1_LINE_JOIN_ROUND
	D2D1_LINE_JOIN_MITER_OR_BEVEL
)

// D2D1_DASH_STYLE 虚线样式.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/ne-d2d1-d2d1_dash_style
type D2D1_DASH_STYLE int32

const (
	D2D1_DASH_STYLE_SOLID D2D1_DASH_STYLE = iota
	D2D1_DASH_STYLE_DASH
	D2D1_DASH_STYLE_DOT
	D2D1_DASH_STYLE_DASH_DOT
	D2D1_DASH_STYLE_DASH_DOT_DOT
	D2D1_DASH_STYLE_CUSTOM
)

// D2D1_STROKE_STYLE_PROPERTIES 描边样式属性(按指针传给 ID2D1Factory::CreateStrokeStyle).
//
// 字段顺序即 C 结构体布局(4 个枚举 + 1 个 float + 1 个枚举 = 24 字节), 不要改.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/ns-d2d1-d2d1_stroke_style_properties
type D2D1_STROKE_STYLE_PROPERTIES struct {
	StartCap   D2D1_CAP_STYLE
	EndCap     D2D1_CAP_STYLE
	DashCap    D2D1_DASH_CAP_STYLE
	LineJoin   D2D1_LINE_JOIN
	MiterLimit float32
	DashStyle  D2D1_DASH_STYLE
}

// RoundStrokeStyleProperties 返回"圆角连接 + 圆头 + 实线"的描边属性 —— 即 SVG 里
// `stroke-linecap="round" stroke-linejoin="round"` 的语义(lucide 图标就是这个).
func RoundStrokeStyleProperties() D2D1_STROKE_STYLE_PROPERTIES {
	return D2D1_STROKE_STYLE_PROPERTIES{
		StartCap:  D2D1_CAP_STYLE_ROUND,
		EndCap:    D2D1_CAP_STYLE_ROUND,
		DashCap:   D2D1_DASH_CAP_STYLE_ROUND,
		LineJoin:  D2D1_LINE_JOIN_ROUND,
		DashStyle: D2D1_DASH_STYLE_SOLID,
	}
}

// IdentityMatrix 返回单位矩阵.
func IdentityMatrix() D2D1_MATRIX_3X2_F {
	return D2D1_MATRIX_3X2_F{M11: 1, M22: 1}
}

// TranslationMatrix 返回平移矩阵(把绘制内容整体位移 dx/dy).
func TranslationMatrix(dx, dy float32) D2D1_MATRIX_3X2_F {
	return D2D1_MATRIX_3X2_F{M11: 1, M22: 1, DX: dx, DY: dy}
}

// DXGI_FORMAT 是 DXGI 的像素格式枚举(这里只列出用得到的值).
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/dxgiformat/ne-dxgiformat-dxgi_format
type DXGI_FORMAT uint32

const (
	// DXGI_FORMAT_B8G8R8A8_UNORM 32bpp 蓝绿红 + alpha, 每通道 8 位; 就是 GDI 的 32bpp DIB 布局.
	DXGI_FORMAT_B8G8R8A8_UNORM DXGI_FORMAT = 87
	// DXGI_FORMAT_B8G8R8A8_UNORM_SRGB 同上但走 sRGB 色彩空间(GDI 的 DIB 用它会被 D2D 拒绝).
	DXGI_FORMAT_B8G8R8A8_UNORM_SRGB DXGI_FORMAT = 91
)

// D2D1_ALPHA_MODE 像素里 alpha 通道的语义.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/dcommon/ne-dcommon-d2d1_alpha_mode
type D2D1_ALPHA_MODE uint32

const (
	// D2D1_ALPHA_MODE_UNKNOWN alpha 语义未知(由渲染目标自行决定).
	D2D1_ALPHA_MODE_UNKNOWN D2D1_ALPHA_MODE = 0
	// D2D1_ALPHA_MODE_PREMULTIPLIED 预乘: 颜色分量已经乘过 alpha(c×a/255). D2D 的 DC 渲染目标
	// 只支持这一种带 alpha 的语义, 取出的像素因此要**反预乘**一趟.
	D2D1_ALPHA_MODE_PREMULTIPLIED D2D1_ALPHA_MODE = 1
	// D2D1_ALPHA_MODE_STRAIGHT 直通(非预乘)alpha —— DC 渲染目标不支持.
	D2D1_ALPHA_MODE_STRAIGHT D2D1_ALPHA_MODE = 2
	// D2D1_ALPHA_MODE_IGNORE 忽略 alpha(当作不透明).
	D2D1_ALPHA_MODE_IGNORE D2D1_ALPHA_MODE = 3
)

// D2D1_PIXEL_FORMAT 像素格式: DXGI 格式 + alpha 语义. 按值嵌在渲染目标属性里(不是指针).
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/dcommon/ns-dcommon-d2d1_pixel_format
type D2D1_PIXEL_FORMAT struct {
	Format    DXGI_FORMAT
	AlphaMode D2D1_ALPHA_MODE
}

// D2D1_RENDER_TARGET_TYPE 渲染目标的实现类型: 硬件 / 软件 / 默认.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/ne-d2d1-d2d1_render_target_type
type D2D1_RENDER_TARGET_TYPE uint32

const (
	// D2D1_RENDER_TARGET_TYPE_DEFAULT 默认(能用硬件就用硬件).
	D2D1_RENDER_TARGET_TYPE_DEFAULT D2D1_RENDER_TARGET_TYPE = 0
	// D2D1_RENDER_TARGET_TYPE_SOFTWARE 强制软件渲染.
	D2D1_RENDER_TARGET_TYPE_SOFTWARE D2D1_RENDER_TARGET_TYPE = 1
	// D2D1_RENDER_TARGET_TYPE_HARDWARE 强制硬件渲染.
	D2D1_RENDER_TARGET_TYPE_HARDWARE D2D1_RENDER_TARGET_TYPE = 2
)

// D2D1_RENDER_TARGET_USAGE 渲染目标的用途(远程/兼容 GDI 的开关).
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/ne-d2d1-d2d1_render_target_usage
type D2D1_RENDER_TARGET_USAGE uint32

const (
	// D2D1_RENDER_TARGET_USAGE_NONE 无特殊用途.
	D2D1_RENDER_TARGET_USAGE_NONE D2D1_RENDER_TARGET_USAGE = 0x00000000
	// D2D1_RENDER_TARGET_USAGE_FORCE_BITMAP_REMOTING 允许远程会话里回落到位图重映射.
	D2D1_RENDER_TARGET_USAGE_FORCE_BITMAP_REMOTING D2D1_RENDER_TARGET_USAGE = 0x00000001
	// D2D1_RENDER_TARGET_USAGE_GDI_COMPATIBLE 允许 GDI 与 D2D 混画(会降低性能).
	D2D1_RENDER_TARGET_USAGE_GDI_COMPATIBLE D2D1_RENDER_TARGET_USAGE = 0x00000002
)

// D2D1_FEATURE_LEVEL 渲染目标要求的最低功能级别.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/ne-d2d1-d2d1_feature_level
type D2D1_FEATURE_LEVEL uint32

const (
	// D2D1_FEATURE_LEVEL_DEFAULT 默认(D2D 自行决定).
	D2D1_FEATURE_LEVEL_DEFAULT D2D1_FEATURE_LEVEL = 0
	// D2D1_FEATURE_LEVEL_9 Direct3D 9 级功能.
	D2D1_FEATURE_LEVEL_9 D2D1_FEATURE_LEVEL = 0x9100
	// D2D1_FEATURE_LEVEL_10 Direct3D 10 级功能.
	D2D1_FEATURE_LEVEL_10 D2D1_FEATURE_LEVEL = 0xa000
)

// D2D1_FACTORY_TYPE 工厂的线程模型.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/ne-d2d1-d2d1_factory_type
type D2D1_FACTORY_TYPE uint32

const (
	// D2D1_FACTORY_TYPE_SINGLE_THREADED 单线程工厂: 建资源不带锁, 更快, 但**调用方**要自己保证
	// 同一时刻只有一个线程在用(d2d1.dll 的 D2D1CreateFactory 默认推荐值).
	D2D1_FACTORY_TYPE_SINGLE_THREADED D2D1_FACTORY_TYPE = 0
	// D2D1_FACTORY_TYPE_MULTI_THREADED 多线程工厂: 内部加锁, 但会拉低单线程下的性能.
	D2D1_FACTORY_TYPE_MULTI_THREADED D2D1_FACTORY_TYPE = 1
)

// D2D1_RENDER_TARGET_PROPERTIES 渲染目标的属性(按指针传给 CreateDCRenderTarget / IsSupported).
//
// 字段顺序即 C 结构体布局(type 4 字节 + pixelFormat 8 字节 + dpi 8 字节 + usage 4 字节 +
// minLevel 4 字节 = 28 字节), 不要改.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/ns-d2d1-d2d1_render_target_properties
type D2D1_RENDER_TARGET_PROPERTIES struct {
	Type        D2D1_RENDER_TARGET_TYPE
	PixelFormat D2D1_PIXEL_FORMAT
	DpiX        float32
	DpiY        float32
	Usage       D2D1_RENDER_TARGET_USAGE
	MinLevel    D2D1_FEATURE_LEVEL
}

// NewDCRenderTargetProperties 返回"32bpp BGRA + 预乘 alpha + 96 DPI"的属性 —— 离屏渲染成
// 一张可读出像素的图片(见 RenderToRGBA)用的就是这套: 96 DPI 让 1 DIP == 1 像素, 画布坐标即
// 目标像素坐标, 不用再做 DPI 换算.
func NewDCRenderTargetProperties() D2D1_RENDER_TARGET_PROPERTIES {
	return D2D1_RENDER_TARGET_PROPERTIES{
		Type: D2D1_RENDER_TARGET_TYPE_DEFAULT,
		PixelFormat: D2D1_PIXEL_FORMAT{
			Format:    DXGI_FORMAT_B8G8R8A8_UNORM,
			AlphaMode: D2D1_ALPHA_MODE_PREMULTIPLIED,
		},
		DpiX:     96,
		DpiY:     96,
		Usage:    D2D1_RENDER_TARGET_USAGE_NONE,
		MinLevel: D2D1_FEATURE_LEVEL_DEFAULT,
	}
}

// D2D1_FILL_MODE 填充模式.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/ne-d2d1-d2d1_fill_mode
type D2D1_FILL_MODE uint32

const (
	// D2D1_FILL_MODE_ALTERNATE 交替填充(奇偶环绕): 重叠区被挖空, 是 D2D 的默认值.
	D2D1_FILL_MODE_ALTERNATE D2D1_FILL_MODE = 0
	// D2D1_FILL_MODE_WINDING 环绕填充(非零环绕): 重叠区取并集 —— 互相重叠的多边形
	// (例如把描边展开成的一条条轮廓)必须用它.
	D2D1_FILL_MODE_WINDING D2D1_FILL_MODE = 1
)

// D2D1_FIGURE_BEGIN 子路径的起笔方式.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/ne-d2d1-d2d1_figure_begin
type D2D1_FIGURE_BEGIN uint32

const (
	// D2D1_FIGURE_BEGIN_FILLED 子路径围出的区域参与填充.
	D2D1_FIGURE_BEGIN_FILLED D2D1_FIGURE_BEGIN = 0
	// D2D1_FIGURE_BEGIN_HOLLOW 空心(不填充), 只用于描边类几何.
	D2D1_FIGURE_BEGIN_HOLLOW D2D1_FIGURE_BEGIN = 1
)

// D2D1_FIGURE_END 子路径的收笔方式.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/ne-d2d1-d2d1_figure_end
type D2D1_FIGURE_END uint32

const (
	// D2D1_FIGURE_END_OPEN 不闭合.
	D2D1_FIGURE_END_OPEN D2D1_FIGURE_END = 0
	// D2D1_FIGURE_END_CLOSED 自动闭合(首尾相连).
	D2D1_FIGURE_END_CLOSED D2D1_FIGURE_END = 1
)

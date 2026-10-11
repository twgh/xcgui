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

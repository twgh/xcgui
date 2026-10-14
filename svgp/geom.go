package svgp

import (
	"errors"

	"github.com/twgh/xcgui/wapi/d2d"
)

// Geometries 一组可直接绘制的几何: 描边子路径与填充子路径分开(图标通常只有描边一侧).
//
// 它们是**设备无关**资源(归属 d2d 工厂), 建好后可以反复绘制; 工厂换了(D2D 设备重建)要整体重建 ——
// 建议配合 d2d.GeometryCache 按 key 缓存, 不要每帧新建.
type Geometries struct {
	Stroke *d2d.ID2D1PathGeometry // 描边折线(保留各自的开口/闭合)
	Fill   *d2d.ID2D1PathGeometry // 填充形状(闭合, 绕向已统一)
}

// NewGeometries 把子路径建成 d2d 几何.
//
// scale 会把点坐标统一缩放(例: 24 网格的图标画到 size 像素时传 size/24); 传 1 表示点已是目标坐标.
// 两条几何都可以是 nil(该侧没有子路径), 不算错误 —— 用 Empty 判断; 全都建不出来才返回 error.
func NewGeometries(f *d2d.ID2D1Factory, subs []SubPath, scale float32) (Geometries, error) {
	var g Geometries
	if f == nil {
		return g, errors.New("svgp: d2d 工厂为空(未开启 D2D?)")
	}
	if scale <= 0 {
		scale = 1
	}
	var strokeFigs, fillFigs []d2d.GeometryFigure
	for _, sub := range subs {
		if len(sub.Points) < 2 {
			continue
		}
		pts := make([]float32, 0, len(sub.Points)*2)
		for _, p := range sub.Points {
			pts = append(pts, p.X*scale, p.Y*scale)
		}
		if sub.Filled {
			// 填充形状: 闭合 + 统一绕向(WINDING 下反向重叠会把区域挖空)
			fillFigs = append(fillFigs, d2d.GeometryFigure{Points: sameOrientation(pts), Closed: true})
			continue
		}
		strokeFigs = append(strokeFigs, d2d.GeometryFigure{Points: pts, Closed: sub.Closed})
	}
	if len(strokeFigs) > 0 {
		if pg, err := f.NewGeometryFigures(strokeFigs); err == nil {
			g.Stroke = pg
		}
	}
	if len(fillFigs) > 0 {
		if pg, err := f.NewGeometryFigures(fillFigs); err == nil {
			g.Fill = pg
		}
	}
	if g.Empty() {
		return g, errors.New("svgp: 没有可用的子路径(每条至少 2 个点)")
	}
	return g, nil
}

// ColorGroup 一组**同色**几何: 颜色 + 该颜色下按填充/描边分好的几何.
//
// 它是多色 SVG(iconfont 的彩色图标: 每条 path 自带 fill="#xxx")的中间形态 —— 绘制就是逐组
// 走 Geometries.Draw. 单色图标(全部 currentColor)只会得到 1 组.
type ColorGroup struct {
	// Color 是该组的颜色: 子路径自带的色(HasColor=true), 或调用方给的 override.
	Color d2d.D2D1_COLOR_F
	// Geom 是该颜色下的几何(描边 + 填充两侧, 见 Geometries).
	Geom Geometries
}

// Release 释放该组的几何(nil 安全).
func (g ColorGroup) Release() { g.Geom.Release() }

// NewColorGroups 按"颜色 + 填充/描边"把子路径分组建几何(for 多色图标).
//
// override 用于 **HasColor=false** 的子路径(currentColor / 没写颜色 / 解析不了的色),
// 也就是"整体单色"图标的颜色来源; HasColor=true 的子路径用自己解析出来的色.
// 其余参数语义与 NewGeometries 一致(scale 是统一缩放; 组的先后顺序 = 颜色首次出现的顺序).
//
// 某组建不出几何(点太少等)会被跳过; **全部**组都建不出来才返回 error.
func NewColorGroups(f *d2d.ID2D1Factory, subs []SubPath, scale float32, override d2d.D2D1_COLOR_F) ([]ColorGroup, error) {
	if f == nil {
		return nil, errors.New("svgp: d2d 工厂为空(未开启 D2D?)")
	}
	type groupKey struct {
		color d2d.D2D1_COLOR_F
		fill  bool
	}
	keys := make([]groupKey, 0, 4)
	buckets := make(map[groupKey][]SubPath)
	for _, sub := range subs {
		c := override
		if sub.HasColor {
			c = sub.Color
		}
		k := groupKey{color: c, fill: sub.Filled}
		if _, ok := buckets[k]; !ok {
			keys = append(keys, k)
		}
		buckets[k] = append(buckets[k], sub)
	}
	groups := make([]ColorGroup, 0, len(keys))
	for _, k := range keys {
		g, err := NewGeometries(f, buckets[k], scale)
		if err != nil {
			continue // 这一组建不出来就跳过它(整体都建不出来时下面统一报错)
		}
		groups = append(groups, ColorGroup{Color: k.color, Geom: g})
	}
	if len(groups) == 0 {
		return nil, errors.New("svgp: 没有可用的子路径(每条至少 2 个点)")
	}
	return groups, nil
}

// ReleaseColorGroups 释放全部分组的几何(便利函数, nil 安全).
func ReleaseColorGroups(groups []ColorGroup) {
	for _, g := range groups {
		g.Release()
	}
}

// Empty 两条几何都没有.
func (g Geometries) Empty() bool { return g.Stroke == nil && g.Fill == nil }

// Release 释放两条几何(nil 安全).
//
// 与 d2d 的 Release 一样是值接收者: 释放之后请不要再使用这个值(需要重复释放就自己置零).
func (g Geometries) Release() {
	if g.Stroke != nil {
		g.Stroke.Release()
	}
	if g.Fill != nil {
		g.Fill.Release()
	}
}

// Draw 一步绘制: 描边子路径走 DrawGeometry(线宽 + 可选样式), 填充子路径走 FillGeometry.
//
// color 直接给 d2d 颜色; strokeWidth 是对应网格单位的线宽(SVG 的 24 网格就是 stroke-width 的值).
// style 传 nil = 默认(平头 + 斜接); SVG 的圆角语义用 d2d.RoundStrokeStyleProperties() 建.
func (g Geometries) Draw(rt *d2d.ID2D1RenderTarget, color d2d.D2D1_COLOR_F, strokeWidth float32, style *d2d.ID2D1StrokeStyle) error {
	if rt == nil {
		return errors.New("svgp: 渲染目标为空")
	}
	if g.Stroke != nil {
		if err := rt.DrawGeometryColor(g.Stroke.Base(), color, strokeWidth, style); err != nil {
			return err
		}
	}
	if g.Fill != nil {
		if err := rt.FillGeometryColor(g.Fill.Base(), color); err != nil {
			return err
		}
	}
	return nil
}

// sameOrientation 把多边形统一成"正向"(鞋带公式面积为正), 需要时反转点序.
//
// 为什么必须: WINDING 填充模式下只有缠绕组数非 0 的区域才填充 —— 反向重叠的多边形会让 winding 归零,
// 重叠区被挖空. 描边子路径交给 D2D 自己处理, 这条只对填充形状重要.
func sameOrientation(pts []float32) []float32 {
	n := len(pts) / 2
	if n < 3 {
		return pts
	}
	var sum float64
	for i := 0; i < n-1; i++ {
		sum += float64(pts[i*2])*float64(pts[i*2+3]) - float64(pts[i*2+2])*float64(pts[i*2+1])
	}
	sum += float64(pts[(n-1)*2])*float64(pts[1]) - float64(pts[0])*float64(pts[(n-1)*2+1])
	if sum >= 0 {
		return pts
	}
	rev := make([]float32, len(pts))
	for i := 0; i < n; i++ {
		rev[i*2] = pts[(n-1-i)*2]
		rev[i*2+1] = pts[(n-1-i)*2+1]
	}
	return rev
}

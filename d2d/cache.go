package d2d

// GeometryCache 按 key 缓存几何.
//
// 为什么值得有个现成的: 几何是**设备无关**资源(归属工厂), 建一次可以跨帧、跨窗口(不同渲染目标)
// 复用; 但它随工厂一起失效 —— 工厂变了(D2D 设备重建)必须整体作废重建; 还有"建失败也要记一笔"
// 这条容易漏(否则每帧都会重试建一条坏几何). 这几件事写成工具就没有各写一遍的余地了.
//
// 用法:
//
//	var cache d2d.GeometryCache[myKey]         // 零值可用; key 必须可比较
//	geom := cache.Get(factory, myKey{...}, func() (*d2d.ID2D1Geometry, error) {
//		pg, err := factory.NewGeometryFigures(figs)
//		if pg == nil {
//			return nil, err
//		}
//		return pg.Base(), nil                  // 路径几何转基类视图; 矩形/椭圆几何自带 Base()
//	})
//	if geom != nil {
//		rt.FillGeometry(geom, brush)
//	}
//
// 边界:
//   - **不是并发安全的** —— 几何挂在工厂上, 建议只在 UI 线程用(绘制回调就是 UI 线程);
//   - key 由调用方定, 把"会改变几何形状的参数"都编进去(例如"路径 + 尺寸"); **位置/线宽不该进 key** ——
//     它们由绘制时的变换矩阵与线宽参数施加, 编进去只会让缓存在窗口 resize 时整体失效;
//   - 缓存里的几何统一由本 cache Release; 工厂变更时自动作废(旧几何全部 Release 掉).
type GeometryCache[K comparable] struct {
	factory *ID2D1Factory
	geoms   map[K]*ID2D1Geometry
}

// Get 取几何: 命中(含"上次建失败"的空项)直接返回; 未命中则调 build 建一次.
//
// build 返回错误时记下空项并返回 nil —— 不会每帧重试(坏 key 只付一次代价).
func (c *GeometryCache[K]) Get(factory *ID2D1Factory, key K, build func() (*ID2D1Geometry, error)) *ID2D1Geometry {
	if factory == nil {
		return nil
	}
	if factory != c.factory { // 工厂变了 → 旧几何全部作废
		c.Clear()
		c.factory = factory
	}
	if c.geoms == nil {
		c.geoms = make(map[K]*ID2D1Geometry)
	}
	g, ok := c.geoms[key]
	if !ok {
		if build != nil {
			g, _ = build()
		}
		c.geoms[key] = g // 失败(nil)也记下, 免得每帧重试
	}
	return g
}

// Clear 释放全部缓存几何并清空(工厂变更时会自动调用).
func (c *GeometryCache[K]) Clear() {
	for _, g := range c.geoms {
		g.Release()
	}
	c.geoms = nil
}

// Len 缓存项数(含"建失败"的空项).
func (c *GeometryCache[K]) Len() int { return len(c.geoms) }

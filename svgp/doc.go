// Package svgp 把 SVG 的 path 数据与 SVG 文档解析成"折线子路径", 需要时再交给 d2d 画出来.
//
// 它是**纯 Go** 的 path 子集实现 —— 不经过 xcgui.dll 的 svg 支持(那套是 DLL 里的 C++ 实现,
// 对官方图标站导出的紧凑写法解析不了、圆弧走样、fill="none" 会被糊成填充块; 本包与之并存,
// 名字也刻意避开 svg). 因为不依赖 DLL, 它可以独立升级.
//
// # 能力边界(先用它之前要知道)
//
// 支持:
//   - 全部 path 命令: M L H V Z C S Q T A(含相对形式, 含 SVG 允许的紧凑写法 "a.25.25 0 00.407.195");
//   - <svg> / <g> / <path> 三个元素, 含 fill / stroke / stroke-width 的**继承**;
//   - 颜色: #rrggbb / #rgb(以及 #rgba / #rrggbbaa)形式的 fill 或 stroke 会保留在子路径上
//     (见 SubPath.Color / HasColor), 于是**多色图标**(iconfont 的彩色图标: 每条 path 自带
//     fill="#xxx")也能出对色 —— 渲染时用 NewColorGroups 按颜色分组; 其余写法(currentColor /
//     命名色 / rgb() / url(#渐变))算"没颜色", 由调用方给的颜色兜底;
//   - viewBox 提取(缩放交给 ScaleTo), 多个 <path> 合并.
//
// 不支持(碰到就忽略, 不报错):
//   - <circle> / <rect> / <line> / <polyline> / <polygon> 等形状元素;
//   - transform 属性、CSS 样式表(class 里的 fill/stroke)、渐变 / 裁剪 / 蒙版 / 文字;
//   - 精确曲线(贝塞尔与圆弧统一离散成折线: 三次贝塞尔 8 段、圆弧每段圆心角 <= 90°).
//
// 换句话说: 它服务的是**图标级别的矢量图形**(lucide / heroicons / feather / iconfont 那类),
// 不是通用 SVG 渲染引擎. 名字取 "svgp" = SVG path.
//
// # 用法
//
// 只要折线(自己画、自己算):
//
//	subs, err := svgp.ParsePath("M12 3 L21 21 L3 21 Z")
//	subs, err = svgp.ParseDocSubPaths(svgText) // 整段 SVG(含 fill/stroke 判定)
//
// 直接画到炫彩的 Direct2D 渲染目标上(在绘制回调里):
//
//	doc, err := svgp.ParseDoc(svgText)                       // 或 svgp.ParsePath(...) 自己组
//	subs := svgp.ScaleTo(doc.Paths, doc.ViewBox, 24)         // 归一到 24 网格(可按需)
//	geoms, err := svgp.NewGeometries(factory, subs, 1)       // 建几何(设备无关, 可缓存)
//	defer geoms.Release()
//	err = geoms.Draw(rt, color, 2, style)                    // 描边走 DrawGeometry, 填充走 FillGeometry
//
// 同一段代码也支持多色图标 —— 把 NewGeometries 换成 NewColorGroups, 逐组绘制:
//
//	groups, err := svgp.NewColorGroups(factory, subs, 1, color) // 子路径按"颜色 + 填充/描边"分组
//	defer svgp.ReleaseColorGroups(groups)
//	for _, g := range groups {
//		g.Geom.Draw(rt, g.Color, 2, style)
//	}
//
// 把 SVG 渲染成**一张图片**(不与窗口/绘制回调绑定, 即"离屏"):
//
//	doc, err := svgp.ParseDoc(svgText)
//	subs := svgp.ScaleTo(doc.Paths, doc.ViewBox, 32)                 // 目标像素网格
//	img, err := svgp.RenderToRGBA(subs, 32, color, 2*32.0/24, nil)   // 线宽按 viewBox 换算
//	// img 是**非预乘** RGBA, 可直接喂 xc.XImage_LoadFromGoImage 拿图片句柄
//
// RenderToRGBA 自建 D2D 工厂(不依赖炫彩初始化), 所以 GDI 渲染模式(app.New(false))、
// 甚至"还没建窗口"时都能用; 内部按尺寸复用离屏目标, 32px 图标单次约 0.5ms 量级.
//
// # 与 d2d 的关系
//
// 本包只做"数据 → 折线 → 几何", 绘制靠同仓的 github.com/twgh/xcgui/wapi/d2d(直通系统 Direct2D).
// 几何是设备无关资源(归属工厂), 建议配合 d2d.GeometryCache 按 key 缓存, 不要每帧重建.
package svgp

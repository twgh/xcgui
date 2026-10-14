package svgp

import (
	"image"
	"math"
	"testing"

	"github.com/twgh/xcgui/wapi/d2d"
	"github.com/twgh/xcgui/wapi/gdi"
	"github.com/twgh/xcgui/wapi/types"
)

// ---------------- 颜色解析(纯逻辑, 不碰 D2D) ----------------

func TestParseColor(t *testing.T) {
	cases := []struct {
		in         string
		ok         bool
		r, g, b, a uint8
	}{
		{"#f00", true, 255, 0, 0, 255},       // #RGB 每位重复
		{"#ff0000", true, 255, 0, 0, 255},    // #RRGGBB
		{"#00ff00", true, 0, 255, 0, 255},    //
		{"#0000FF", true, 0, 0, 255, 255},    // 大小写不敏感
		{"#f008", true, 255, 0, 0, 0x88},     // #RGBA
		{"#ff000080", true, 255, 0, 0, 0x80}, // #RRGGBBAA
		{" #123456 ", true, 0x12, 0x34, 0x56, 255},
		{"currentColor", false, 0, 0, 0, 0}, // 由调用方决定
		{"none", false, 0, 0, 0, 0},         //
		{"red", false, 0, 0, 0, 0},          // 命名色不支持
		{"rgb(255,0,0)", false, 0, 0, 0, 0}, // rgb() 不支持
		{"url(#grad)", false, 0, 0, 0, 0},   // 渐变不支持
		{"#gg0000", false, 0, 0, 0, 0},      // 非法十六进制
		{"#12345", false, 0, 0, 0, 0},       // 5 位: 不支持的长度
		{"", false, 0, 0, 0, 0},             //
	}
	for _, c := range cases {
		got, ok := parseColor(c.in)
		if ok != c.ok {
			t.Errorf("parseColor(%q) ok = %v, 想要 %v", c.in, ok, c.ok)
			continue
		}
		if !ok {
			continue
		}
		r, g, b, a := to255(got.R), to255(got.G), to255(got.B), to255(got.A)
		if r != c.r || g != c.g || b != c.b || a != c.a {
			t.Errorf("parseColor(%q) = rgba(%d,%d,%d,%d), 想要 rgba(%d,%d,%d,%d)",
				c.in, r, g, b, a, c.r, c.g, c.b, c.a)
		}
	}
}

func to255(v float32) uint8 { return uint8(math.Round(float64(v) * 255)) }

// ---------------- ParseDoc 的颜色归属 ----------------

func TestParseDocColors(t *testing.T) {
	const svg = `<svg viewBox="0 0 24 24">
  <path d="M2 2 L22 2" stroke="currentColor"/>
  <path d="M2 4 L22 4" fill="#ff0000"/>
  <path d="M2 6 L22 6" fill="none" stroke="#0f0"/>
  <g fill="#0000ff"><path d="M2 8 L22 8"/></g>
</svg>`
	doc, err := ParseDoc(svg)
	if err != nil {
		t.Fatal(err)
	}
	if len(doc.Paths) != 4 {
		t.Fatalf("子路径数 = %d, 想要 4", len(doc.Paths))
	}
	// currentColor: 没颜色, 用调用方的
	if doc.Paths[0].HasColor {
		t.Error("stroke=currentColor 不该有自己的颜色")
	}
	// 显式 fill: 填充 + 自带红
	p1 := doc.Paths[1]
	if !p1.Filled || !p1.HasColor || to255(p1.Color.R) != 255 || to255(p1.Color.G) != 0 {
		t.Errorf("fill=#ff0000 的子路径 = {Filled:%v HasColor:%v R:%d}", p1.Filled, p1.HasColor, to255(p1.Color.R))
	}
	// fill=none + stroke=#0f0: 描边 + 自带绿
	p2 := doc.Paths[2]
	if p2.Filled || !p2.HasColor || to255(p2.Color.G) != 255 {
		t.Errorf("fill=none stroke=#0f0 的子路径 = {Filled:%v HasColor:%v G:%d}", p2.Filled, p2.HasColor, to255(p2.Color.G))
	}
	// <g fill=#0000ff> 继承: 填充 + 自带蓝
	p3 := doc.Paths[3]
	if !p3.Filled || !p3.HasColor || to255(p3.Color.B) != 255 {
		t.Errorf("<g fill=#0000ff> 的子路径 = {Filled:%v HasColor:%v B:%d}", p3.Filled, p3.HasColor, to255(p3.Color.B))
	}
}

// 根 <svg> 上的 fill / stroke 也参与继承(iconfont 的彩色图标常把色写在根上).
func TestParseDocRootInherit(t *testing.T) {
	doc, err := ParseDoc(`<svg viewBox="0 0 24 24" fill="none" stroke="#cc0000"><path d="M2 6h20"/></svg>`)
	if err != nil {
		t.Fatal(err)
	}
	p := doc.Paths[0]
	if p.Filled {
		t.Error("根 fill=none → 子路径应是描边")
	}
	if !p.HasColor || to255(p.Color.R) != 0xcc {
		t.Errorf("根 stroke=#cc0000 没继承下来: HasColor=%v Color=%+v", p.HasColor, p.Color)
	}
}

// 缩放只改坐标, 颜色/填充/闭合这些语义字段必须原样带过去.
func TestScaleToKeepsColor(t *testing.T) {
	doc, err := ParseDoc(`<svg viewBox="0 0 12 12"><path d="M0 0 L12 0" stroke="#ff0000"/></svg>`)
	if err != nil {
		t.Fatal(err)
	}
	subs := ScaleTo(doc.Paths, doc.ViewBox, 24)
	if !subs[0].HasColor || to255(subs[0].Color.R) != 255 {
		t.Errorf("缩放后颜色丢了: HasColor=%v Color=%+v", subs[0].HasColor, subs[0].Color)
	}
	if p := subs[0].Points[1]; p.X != 24 {
		t.Errorf("缩放后末点 X = %v, 想要 24", p.X)
	}
	// ParsePath 单独一条数据时无从判定颜色
	raw, err := ParsePath("M0 0 L1 1")
	if err != nil {
		t.Fatal(err)
	}
	if raw[0].HasColor {
		t.Error("ParsePath 不该猜颜色")
	}
}

// ---------------- 反预乘(纯逻辑) ----------------

func TestUnpremul(t *testing.T) {
	cases := []struct{ c, a, want byte }{
		{200, 255, 200}, // a=255 原样
		{128, 128, 255}, // 预乘满值还原
		{64, 128, 128},  // 半值还原
		{0, 1, 0},       // 预乘为 0 就还原不出东西
		{1, 1, 255},     // a=1 下任意非零都是满值
		{200, 100, 255}, // 预乘舍入过头时钳到 255
	}
	for _, c := range cases {
		if got := unpremul(c.c, c.a); got != c.want {
			t.Errorf("unpremul(%d, %d) = %d, 想要 %d", c.c, c.a, got, c.want)
		}
	}
}

// ---------------- 真机渲染(缺 D2D 时跳过) ----------------

// requireD2D 探测环境里能不能自建 D2D 工厂(不能就跳过真机测试, 不阻塞 CI).
func requireD2D(t *testing.T) {
	t.Helper()
	f, err := d2d.D2D1CreateFactory(d2d.D2D1_FACTORY_TYPE_SINGLE_THREADED)
	if err != nil {
		t.Skipf("环境没有可用的 D2D: %v", err)
	}
	f.Release()
}

// fillRectSub 造一条"填充矩形"的子路径(像素网格).
func fillRectSub(x1, y1, x2, y2 float32) SubPath {
	return SubPath{
		Points: []Point{{x1, y1}, {x2, y1}, {x2, y2}, {x1, y2}},
		Filled: true,
		Closed: true,
	}
}

func TestRenderToRGBABasic(t *testing.T) {
	requireD2D(t)
	const size = 16
	// 左上 8×8 填不透明红(边界与像素边界齐平, 没有 AA 过渡像素).
	img, err := RenderToRGBA([]SubPath{fillRectSub(0, 0, 8, 8)}, size, d2d.NewColorF(255, 0, 0, 255), 0, nil)
	if err != nil {
		t.Fatal(err)
	}
	if got := img.Bounds().Dx(); got != size {
		t.Fatalf("图宽 = %d, 想要 %d", got, size)
	}
	checkPx := func(x, y int, r, g, b, a uint8, what string) {
		t.Helper()
		i := img.PixOffset(x, y)
		if img.Pix[i] != r || img.Pix[i+1] != g || img.Pix[i+2] != b || img.Pix[i+3] != a {
			t.Errorf("%s(%d,%d) = %v, 想要 [%d %d %d %d]", what, x, y, img.Pix[i:i+4], r, g, b, a)
		}
	}
	checkPx(0, 0, 255, 0, 0, 255, "方块内")
	checkPx(7, 7, 255, 0, 0, 255, "方块内")
	checkPx(8, 8, 0, 0, 0, 0, "方块外") // 反预乘后透明像素必须干干净净
	checkPx(15, 15, 0, 0, 0, 0, "方块外")
}

// 半透明填充走一遍"预乘 → 反预乘": alpha 保留, 颜色还原(R 回到 255).
func TestRenderToRGBAAlphaRoundTrip(t *testing.T) {
	requireD2D(t)
	img, err := RenderToRGBA([]SubPath{fillRectSub(0, 0, 16, 16)}, 16, d2d.NewColorF(255, 0, 0, 128), 0, nil)
	if err != nil {
		t.Fatal(err)
	}
	i := img.PixOffset(8, 8)
	r, g, b, a := img.Pix[i], img.Pix[i+1], img.Pix[i+2], img.Pix[i+3]
	if a < 126 || a > 130 {
		t.Errorf("alpha = %d, 想要 ~128", a)
	}
	if r < 253 {
		t.Errorf("反预乘后 R = %d, 想要 ~255(预乘丢过精度, 误差应只在 ±2 内)", r)
	}
	if g > 2 || b > 2 {
		t.Errorf("G/B = %d/%d, 想要 ~0", g, b)
	}
}

// 多色: 自带颜色的子路径用自己那份色, 没颜色的子路径用 override.
func TestRenderToRGBAMultiColor(t *testing.T) {
	requireD2D(t)
	subs := []SubPath{
		fillRectSub(0, 0, 8, 8), // 红块(override)
		func() SubPath { // 绿块(自带色)
			s := fillRectSub(8, 8, 16, 16)
			s.Color, s.HasColor = d2d.NewColorF(0, 255, 0, 255), true
			return s
		}(),
	}
	img, err := RenderToRGBA(subs, 16, d2d.NewColorF(255, 0, 0, 255), 0, nil)
	if err != nil {
		t.Fatal(err)
	}
	at := func(x, y int) []byte { i := img.PixOffset(x, y); return img.Pix[i : i+4] }
	if got := at(2, 2); got[0] != 255 || got[1] != 0 || got[3] != 255 {
		t.Errorf("override 块 = %v, 想要红", got)
	}
	if got := at(12, 12); got[1] != 255 || got[0] != 0 || got[3] != 255 {
		t.Errorf("自带色块 = %v, 想要绿", got)
	}
}

// 描边走的是 DrawGeometry(带线宽): 一条水平线在上下各占线宽的一半.
func TestRenderToRGBAStroke(t *testing.T) {
	requireD2D(t)
	const size = 16
	// y=8 处一条水平线, 线宽 4 → 覆盖 y∈[6,10] 的像素.
	sub := SubPath{Points: []Point{{0, 8}, {16, 8}}}
	img, err := RenderToRGBA([]SubPath{sub}, size, d2d.NewColorF(0, 0, 255, 255), 4, nil)
	if err != nil {
		t.Fatal(err)
	}
	i := img.PixOffset(8, 8)
	if img.Pix[i+2] != 255 || img.Pix[i+3] != 255 {
		t.Errorf("线中心像素 = %v, 想要不透明蓝", img.Pix[i:i+4])
	}
	// 远离线的地方要干净
	j := img.PixOffset(8, 0)
	if img.Pix[j+3] != 0 {
		t.Errorf("线外像素 = %v, 想要全透明", img.Pix[j:j+4])
	}
}

// 失败路径: 参数不合法要报错而不是崩.
func TestRenderToRGBAErrors(t *testing.T) {
	if _, err := RenderToRGBA([]SubPath{fillRectSub(0, 0, 4, 4)}, 0, d2d.D2D1_COLOR_F{}, 1, nil); err == nil {
		t.Error("size=0 应报错")
	}
	if _, err := RenderToRGBA(nil, 8, d2d.D2D1_COLOR_F{}, 1, nil); err == nil {
		t.Error("没有子路径应报错")
	}
	// 点太少的子路径: 建不出几何
	if _, err := RenderToRGBA([]SubPath{{Points: []Point{{0, 0}}}}, 8, d2d.D2D1_COLOR_F{}, 1, nil); err == nil {
		t.Error("没有可用子路径应报错")
	}
}

// ---------------- 与 xui 直通绘制路径对照 ----------------

// lucide star(紧凑写法 + 圆弧): 对照用的真实图标.
const testStarSVG = `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M11.525 2.295a.53.53 0 0 1 .95 0l2.31 4.679a2.123 2.123 0 0 0 1.595 1.16l5.166.756a.53.53 0 0 1 .294.904l-3.736 3.638a2.123 2.123 0 0 0-.611 1.878l.882 5.14a.53.53 0 0 1-.771.56l-4.618-2.428a2.122 2.122 0 0 0-1.973 0L6.396 21.01a.53.53 0 0 1-.77-.56l.881-5.139a2.122 2.122 0 0 0-.611-1.879L2.16 9.795a.53.53 0 0 1 .294-.906l5.165-.755a2.122 2.122 0 0 0 1.597-1.16z"/></svg>`

// xuiPathRender 复刻 xui 的直通绘制路径: 子路径归一到 24 网格, 建几何时整体乘 size/24,
// 线宽也按 size/24 换算(与 xui/components 的 d2dStroke 同样的两步).
func xuiPathRender(t *testing.T, doc *Doc, size int32, color d2d.D2D1_COLOR_F, strokeWidth float32) *image.RGBA {
	t.Helper()
	f, err := d2d.D2D1CreateFactory(d2d.D2D1_FACTORY_TYPE_SINGLE_THREADED)
	if err != nil {
		t.Skipf("环境没有可用的 D2D: %v", err)
	}
	defer f.Release()

	subs24 := ScaleTo(doc.Paths, doc.ViewBox, 24) // 任意网格 → 24(图标约定)
	geoms, err := NewGeometries(f, subs24, float32(size)/24)
	if err != nil {
		t.Fatal(err)
	}
	defer geoms.Release()
	props := d2d.RoundStrokeStyleProperties()
	style, err := f.CreateStrokeStyle(&props)
	if err != nil {
		t.Fatal(err)
	}
	defer style.Release()

	dib, err := gdi.NewDIB(size, size)
	if err != nil {
		t.Fatal(err)
	}
	defer dib.Close()
	rtProps := d2d.NewDCRenderTargetProperties()
	rt, err := f.CreateDCRenderTarget(&rtProps)
	if err != nil {
		t.Fatal(err)
	}
	defer rt.Release()
	if err := rt.BindDC(dib.HDC(), &types.RECT{Left: 0, Top: 0, Right: size, Bottom: size}); err != nil {
		t.Fatal(err)
	}
	rt.BeginDraw()
	rt.Clear(nil)
	if err := geoms.Draw(&rt.ID2D1RenderTarget, color, strokeWidth*float32(size)/24, style); err != nil {
		t.Fatal(err)
	}
	if err := rt.EndDraw(); err != nil {
		t.Fatal(err)
	}
	img, err := dibToRGBA(dib, size)
	if err != nil {
		t.Fatal(err)
	}
	return img
}

// 两条几何构造路径的**同图对照**: 本方案(ScaleTo 到像素网格 + Scale=1) vs xui 直通
// (24 网格 + 建几何时乘 size/24). 同一 DC 渲染目标上应逐像素一致, 只允许浮点结合律级的
// 边缘 AA 微差(float64 与 float32 两条乘法链).
func TestRenderPathMatchesXuiPath(t *testing.T) {
	requireD2D(t)
	doc, err := ParseDoc(testStarSVG)
	if err != nil {
		t.Fatal(err)
	}
	const size = 48
	color := d2d.NewColorF(60, 60, 60, 255)
	sw := float32(2 * size / 24)

	a, err := RenderToRGBA(ScaleTo(doc.Paths, doc.ViewBox, size), size, color, sw, nil)
	if err != nil {
		t.Fatal(err)
	}
	b := xuiPathRender(t, doc, size, color, 2)

	var maxDiff, over2 int
	for i := 0; i < len(a.Pix); i++ {
		d := int(a.Pix[i]) - int(b.Pix[i])
		if d < 0 {
			d = -d
		}
		if d > maxDiff {
			maxDiff = d
		}
		if d > 2 {
			over2++
		}
	}
	if maxDiff > 8 || over2 > len(a.Pix)/100 {
		t.Errorf("两条路径的渲染结果差太多: 最大通道差 = %d, 差>2 的通道 = %d/%d", maxDiff, over2, len(a.Pix))
	}
	t.Logf("最大通道差 = %d, 差>2 的通道 = %d/%d", maxDiff, over2, len(a.Pix))
}

// 反复生成/释放不泄漏 D2D 资源(几何/渲染目标都是即建即放; 这里的意义是"跑得完不炸").
func TestRenderToRGBARepeat(t *testing.T) {
	if testing.Short() {
		t.Skip("短模式跳过")
	}
	requireD2D(t)
	subs := []SubPath{fillRectSub(0, 0, 16, 16)}
	for i := 0; i < 200; i++ {
		img, err := RenderToRGBA(subs, 16, d2d.NewColorF(255, 0, 0, 255), 0, nil)
		if err != nil {
			t.Fatalf("第 %d 次失败: %v", i, err)
		}
		if img == nil {
			t.Fatalf("第 %d 次拿到空图", i)
		}
	}
}

package svgp

import (
	"errors"
	"strings"
	"testing"
)

func TestParsePathBasic(t *testing.T) {
	subs, err := ParsePath("M2 2 L22 2 L22 22 Z")
	if err != nil {
		t.Fatal(err)
	}
	if len(subs) != 1 {
		t.Fatalf("子路径数 = %d, 想要 1", len(subs))
	}
	s := subs[0]
	if len(s.Points) != 3 {
		t.Errorf("点数 = %d, 想要 3(Z 不补重复点)", len(s.Points))
	}
	if !s.Closed {
		t.Error("Z 收尾的子路径应为 Closed")
	}
	if s.Filled {
		t.Error("ParsePath 不判定填充(一律 false)")
	}
	if p := s.Points[0]; p.X != 2 || p.Y != 2 {
		t.Errorf("首点 = (%v,%v), 想要 (2,2)", p.X, p.Y)
	}
}

// H/V 与相对命令: "m4 4 h16 v16 z" 应等价于绝对写法.
func TestParsePathRelAndHV(t *testing.T) {
	subs, err := ParsePath("m4 4 h16 v16 z")
	if err != nil {
		t.Fatal(err)
	}
	if len(subs) != 1 || len(subs[0].Points) != 3 {
		t.Fatalf("解析结果不对: %+v", subs)
	}
	want := []Point{{4, 4}, {20, 4}, {20, 20}}
	for i, w := range want {
		if subs[0].Points[i] != w {
			t.Errorf("点 %d = %+v, 想要 %+v", i, subs[0].Points[i], w)
		}
	}
	if !subs[0].Closed {
		t.Error("z 也应算闭合")
	}
}

// M 后多组坐标按 L 处理(隐式重复), 且一个 M 内不会再开子路径.
func TestParsePathImplicitLine(t *testing.T) {
	subs, err := ParsePath("M0 0 10 0 10 10")
	if err != nil {
		t.Fatal(err)
	}
	if len(subs) != 1 || len(subs[0].Points) != 3 {
		t.Fatalf("隐式 L 没生效: %+v", subs)
	}
}

// 图标站导出的紧凑写法: 省略分隔符 + 单字符标志位(a.25.25 0 00.407.195).
func TestParsePathCompactArc(t *testing.T) {
	const d = "M11.48 3.499a.562.562 0 011.04 0l2.125 5.111a.563.563 0 00.475.345"
	subs, err := ParsePath(d)
	if err != nil {
		t.Fatalf("紧凑写法应能解析: %v", err)
	}
	if len(subs) != 1 {
		t.Fatalf("子路径数 = %d, 想要 1", len(subs))
	}
	if n := len(subs[0].Points); n < 10 {
		t.Errorf("点数 = %d, 弧线没展开?", n)
	}
}

// Z 之后直接跟坐标: 按 M(绝对)处理(规范行为), 开新子路径.
func TestParsePathAfterClose(t *testing.T) {
	subs, err := ParsePath("M0 0 L10 0 Z M20 20 L30 20")
	if err != nil {
		t.Fatal(err)
	}
	if len(subs) != 2 {
		t.Fatalf("子路径数 = %d, 想要 2", len(subs))
	}
	if !subs[0].Closed || subs[1].Closed {
		t.Errorf("闭合标志不对: %v %v", subs[0].Closed, subs[1].Closed)
	}
	if p := subs[1].Points[0]; p.X != 20 || p.Y != 20 {
		t.Errorf("第二条起点 = (%v,%v), 想要 (20,20)", p.X, p.Y)
	}
}

func TestParsePathErrors(t *testing.T) {
	if _, err := ParsePath("X0 0"); err == nil {
		t.Error("未知命令应报错")
	}
	if _, err := ParsePath("M0"); err == nil {
		t.Error("参数不足应报错")
	}
	if _, err := ParsePath("0 0 L1 1"); err == nil {
		t.Error("不是命令开头应报错")
	}
	// 报错时已解析的部分一并返回
	subs, err := ParsePath("M0 0 L10 0 L20")
	if err == nil {
		t.Fatal("应报错")
	}
	if len(subs) == 0 {
		t.Error("报错时应返回已解析的部分")
	}
	var pe *PathError
	if !strings.Contains(err.Error(), "svgp:") || !errors.As(err, &pe) {
		t.Errorf("错误类型不对: %#v", err)
	}
}

func TestParseDocViewBoxAndFill(t *testing.T) {
	const svg = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
  <path d="M2 2 L22 2"/>
  <g fill="currentColor" stroke="none">
    <path d="M6 6 L10 6 L10 10 Z"/>
  </g>
</svg>`
	doc, err := ParseDoc(svg)
	if err != nil {
		t.Fatal(err)
	}
	if doc.ViewBox != [4]float64{0, 0, 24, 24} {
		t.Errorf("ViewBox = %v", doc.ViewBox)
	}
	if doc.StrokeWidth != 1.5 {
		t.Errorf("StrokeWidth = %v, 想要 1.5", doc.StrokeWidth)
	}
	if len(doc.Paths) != 2 {
		t.Fatalf("子路径数 = %d, 想要 2", len(doc.Paths))
	}
	if doc.Paths[0].Filled {
		t.Error("根节点 fill=none 的那条应是描边")
	}
	if !doc.Paths[1].Filled {
		t.Error("<g fill=currentColor stroke=none> 里的应是填充")
	}
}

// 多 path 与形状元素混排: 形状元素被忽略, path 照收.
func TestParseDocIgnoresShapes(t *testing.T) {
	const svg = `<svg viewBox="0 0 16 16">
  <circle cx="8" cy="8" r="6"/>
  <path d="M2 2 L14 2"/>
  <rect x="0" y="0" width="4" height="4"/>
</svg>`
	doc, err := ParseDoc(svg)
	if err != nil {
		t.Fatal(err)
	}
	if len(doc.Paths) != 1 {
		t.Fatalf("子路径数 = %d, 想要 1(形状元素应被忽略)", len(doc.Paths))
	}
}

func TestScaleTo(t *testing.T) {
	doc, err := ParseDoc(`<svg viewBox="0 0 16 16"><path d="M2 2 L14 2"/></svg>`)
	if err != nil {
		t.Fatal(err)
	}
	subs := ScaleTo(doc.Paths, doc.ViewBox, 24) // 16 → 24
	if p := subs[0].Points[0]; p.X != 3 || p.Y != 3 {
		t.Errorf("首点 = (%v,%v), 想要 (3,3)(源 2 × 1.5)", p.X, p.Y)
	}
	if f := ScaleFactor(doc.ViewBox, 24); f != 1.5 {
		t.Errorf("ScaleFactor = %v, 想要 1.5", f)
	}
	// min-x/min-y 要平掉
	doc2, _ := ParseDoc(`<svg viewBox="10 20 24 24"><path d="M10 20 L12 20"/></svg>`)
	subs2 := ScaleTo(doc2.Paths, doc2.ViewBox, 24)
	if p := subs2[0].Points[0]; p.X != 0 || p.Y != 0 {
		t.Errorf("viewBox 偏移没平掉: (%v,%v)", p.X, p.Y)
	}
	// viewBox 无效 → 原样返回
	same := ScaleTo(doc.Paths, [4]float64{}, 24)
	if &same[0] != &doc.Paths[0] {
		t.Error("viewBox 无效时应原样返回(不复制)")
	}
}

func TestParseDocErrors(t *testing.T) {
	if _, err := ParseDoc(`<svg viewBox="0 0 24 24"></svg>`); err == nil {
		t.Error("没有 path 应报错")
	}
	if _, err := ParseDoc("不是 svg"); err == nil {
		t.Error("不是 SVG 应报错")
	}
}

func TestNewGeometriesSafety(t *testing.T) {
	// 没有工厂(未开启 D2D)时要安全报错, 不 panic
	if _, err := NewGeometries(nil, nil, 1); err == nil {
		t.Error("空工厂应报错")
	}
}

// 扩展语法: `!` 标记填充子路径; 一条数据里可以混着描边与填充.
func TestParsePathFillMark(t *testing.T) {
	subs, err := ParsePath("!M2 2 L6 2 L6 6 Z M10 10 L12 12")
	if err != nil {
		t.Fatal(err)
	}
	if len(subs) != 2 {
		t.Fatalf("子路径数 = %d, 想要 2", len(subs))
	}
	if !subs[0].Filled || subs[1].Filled {
		t.Errorf("填充标记不对: %v %v", subs[0].Filled, subs[1].Filled)
	}
	if !subs[0].Closed {
		t.Error("第一条以 Z 收尾应为 Closed")
	}
}

// 相对指令 / 隐式重复 / H V / 粘连负号 / 多子路径 / 参数不足与未知命令要如实报错.
func TestParsePathSyntaxTable(t *testing.T) {
	cases := []struct {
		name   string
		d      string
		subs   int
		filled int
		lastX  float32
		lastY  float32
		ok     bool
	}{
		{"相对折线", "m2 2 4 0 0 4", 1, 0, 6, 6, true},     // m 之后的隐式组按 l 处理
		{"粘连负号", "M10 10l-2-2", 1, 0, 8, 8, true},      // "l-2-2" 是两个数
		{"hv", "M2 2h4v4", 1, 0, 6, 6, true},           // H/V 保持当前另一轴
		{"多子路径", "M2 2L3 3M5 5L6 6", 2, 0, 6, 6, true}, // M 起新子路径
		{"闭合", "M2 2L6 2L6 6Z", 1, 0, 6, 6, true},      // Z 回到子路径起点
		{"填充标记", "!M2 2A2 2 0 1 1 6 2Z M10 10L12 12", 2, 1, 12, 12, true},
		{"不完整", "M2 2L", 0, 0, 0, 0, false}, // 参数不足必须如实报告
		{"未知指令", "M2 2Q", 0, 0, 0, 0, false},
	}
	for _, c := range cases {
		subs, err := ParsePath(c.d)
		if (err == nil) != c.ok {
			t.Errorf("%s: err = %v, 想要 ok=%v", c.name, err, c.ok)
		}
		if len(subs) != c.subs {
			t.Errorf("%s: 子路径 %d 条, 想要 %d", c.name, len(subs), c.subs)
			continue
		}
		nFill := 0
		for _, s := range subs {
			if s.Filled {
				nFill++
			}
		}
		if nFill != c.filled {
			t.Errorf("%s: 填充子路径 %d 条, 想要 %d", c.name, nFill, c.filled)
		}
		if c.subs > 0 {
			last := subs[len(subs)-1].Points
			if p := last[len(last)-1]; p.X != c.lastX || p.Y != c.lastY {
				t.Errorf("%s: 末点 = (%v,%v), 想要 (%v,%v)", c.name, p.X, p.Y, c.lastX, c.lastY)
			}
		}
	}
}

func TestSameOrientation(t *testing.T) {
	// 逆时针(面积为正)原样; 顺时针反转
	ccw := []float32{0, 0, 10, 0, 0, 10}
	if got := sameOrientation(ccw); &got[0] != &ccw[0] {
		t.Error("已正向的多边形不该被复制/反转")
	}
	cw := []float32{0, 0, 0, 10, 10, 0}
	got := sameOrientation(cw)
	// 反转后首点应是原末点
	if got[0] != cw[4] || got[1] != cw[5] {
		t.Errorf("反转后首点 = (%v,%v), 想要 (%v,%v)", got[0], got[1], cw[4], cw[5])
	}
}

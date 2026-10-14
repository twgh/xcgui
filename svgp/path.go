package svgp

import (
	"fmt"
	"math"
	"strconv"
	"unicode"

	"github.com/twgh/xcgui/wapi/d2d"
)

// Point 折线上的一个点.
type Point struct {
	X, Y float32
}

// SubPath 一条子路径: 折线点 + 语义.
type SubPath struct {
	// Points 折线点(源坐标系; 缩放见 ScaleTo, 至少 2 个点).
	Points []Point
	// Filled 为 true 表示这是**填充形状**(绘制走 FillGeometry), false 表示**描边折线**(走 DrawGeometry).
	//
	// ParsePath 解析单独一条 path 数据时无从判定(一律 false) —— 由调用方按 fill/stroke 属性决定;
	// ParseDoc 会按元素(含 <g> 继承)的 fill/stroke 自动填好.
	Filled bool
	// Closed 表示这条子路径以 Z 收尾(首尾相连). 渲染时把它映射成几何的"闭合"标志
	// (d2d.GeometryFigure.Closed), **不要**自己补一个重复点.
	Closed bool
	// Color 是这条子路径**自己**的颜色(#rrggbb / #rgb 形式的 fill, 或 fill="none" 时的
	// stroke); 只在 HasColor 为 true 时有意义.
	Color d2d.D2D1_COLOR_F
	// HasColor 为 false 表示子路径没带可用颜色(currentColor / 没写 / 命名色 / rgb() /
	// url(#渐变) 都算), 绘制时用调用方给的颜色(即 SVG 里 currentColor 的取值).
	//
	// ParsePath 解析单独一条 path 数据时无从判定(一律 false); ParseDoc 按元素(含 <g> 继承)填.
	HasColor bool
}

// PathError path 数据不合法.
type PathError struct {
	Pos int    // 出错处的字节偏移
	Cmd byte   // 出错时的命令(0 = 还没读到命令)
	Why string // 原因
}

func (e *PathError) Error() string {
	if e.Cmd == 0 {
		return fmt.Sprintf("svgp: path 数据在偏移 %d 处不合法: %s", e.Pos, e.Why)
	}
	return fmt.Sprintf("svgp: path 数据在偏移 %d 处不合法(命令 %q): %s", e.Pos, string(e.Cmd), e.Why)
}

// ParsePath 把 SVG 的 path 数据(d 属性)解析成若干条折线子路径.
//
// 支持全部命令: M L H V Z C S Q T A(含相对形式)、隐式重复(M 后多组坐标按 L 处理)、
// 以及 SVG 允许的紧凑写法(省略逗号/空格, 如 "a.25.25 0 00.407.195"). 曲线统一离散成折线:
// 三次/二次贝塞尔 8 段, 圆弧每段圆心角不超过 90°.
//
// **扩展语法**: 子路径前加 `!` 表示它是填充形状(SubPath.Filled = true) —— path 数据本身不携带
// fill/stroke 信息, 而"一条数据里既有描边又有填充"很常见(图标里的色点), `!` 让这种数据能放进
// 一个字符串. 纯 SVG 数据不含它, 一律按描边.
//
// 出错时返回已成功解析的部分 + error —— 调用方可以降级使用这部分, 也可以当失败处理.
func ParsePath(d string) ([]SubPath, error) {
	sc := &scanner{s: d}
	var (
		subs      []SubPath
		cur       []Point
		curClosed bool
		curFilled bool
		x, y      float32 // 当前点
		sx, sy    float32 // 当前子路径起点(Z 回到这里)
		cmd       byte
		lastC     Point // 上一个三次贝塞尔控制点(S 的反射源)
		lastQ     Point // 上一个二次贝塞尔控制点(T 的反射源)
		hasC      bool
		hasQ      bool
	)
	push := func(px, py float32) {
		// 重合的相邻点(路径数据里很常见)直接丢: 零长度段在描边时可能出毛刺
		if n := len(cur); n > 0 && cur[n-1].X == px && cur[n-1].Y == py {
			return
		}
		cur = append(cur, Point{px, py})
	}
	flush := func() {
		if cur == nil { // 没有正在写的子路径: 不要动填充标记(它属于接下来的那条)
			return
		}
		if len(cur) >= 2 { // 单点围不出形状、也描不出线
			subs = append(subs, SubPath{Points: cur, Filled: curFilled, Closed: curClosed})
		}
		cur = nil
		curClosed = false
		curFilled = false
	}
	// 三次贝塞尔离散: 固定 8 段(与图标资源的离散精度一致, 24 网格下误差远小于像素)
	bezier := func(x1, y1, x2, y2, x3, y3, x4, y4 float32) {
		const n = 8
		for k := 1; k <= n; k++ {
			t := float32(k) / n
			u := 1 - t
			push(
				u*u*u*x1+3*u*u*t*x2+3*u*t*t*x3+t*t*t*x4,
				u*u*u*y1+3*u*u*t*y2+3*u*t*t*y3+t*t*t*y4,
			)
		}
	}

	for sc.pos < len(sc.s) {
		sc.skipSep()
		if sc.pos >= len(sc.s) {
			break // 尾部空白: 正常结束
		}
		if sc.s[sc.pos] == '!' { // 扩展语法: 接下来的子路径按填充画
			flush()
			curFilled = true
			sc.pos++
			continue
		}
		if c, ok := sc.cmd(); ok {
			cmd = c
		} else if cmd == 0 {
			flush()
			return subs, &PathError{Pos: sc.pos, Why: "数据不是以命令字母开头"}
		} else if cmd == 'M' {
			cmd = 'L' // 隐式重复: M 后的多组坐标按 L 处理
		} else if cmd == 'm' {
			cmd = 'l'
		}
		rel := unicode.IsLower(rune(cmd))
		ok := true

		switch unicode.ToUpper(rune(cmd)) {
		case 'M':
			var ax, ay float32
			if ax, ok = sc.num(); ok {
				ay, ok = sc.num()
			}
			if !ok {
				break
			}
			flush()
			if rel {
				x, y = x+ax, y+ay
			} else {
				x, y = ax, ay
			}
			sx, sy = x, y
			push(x, y)
			hasC, hasQ = false, false
		case 'L':
			var ax, ay float32
			if ax, ok = sc.num(); ok {
				ay, ok = sc.num()
			}
			if !ok {
				break
			}
			if rel {
				x, y = x+ax, y+ay
			} else {
				x, y = ax, ay
			}
			push(x, y)
			hasC, hasQ = false, false
		case 'H':
			ax, o := sc.num()
			if !o {
				ok = false
				break
			}
			if rel {
				x += ax
			} else {
				x = ax
			}
			push(x, y)
		case 'V':
			ay, o := sc.num()
			if !o {
				ok = false
				break
			}
			if rel {
				y += ay
			} else {
				y = ay
			}
			push(x, y)
		case 'Z':
			curClosed = true
			flush()
			x, y = sx, sy
			cmd = 'M' // 规范: Z 后若直接跟坐标, 按 M(绝对)处理
		case 'C', 'S', 'Q', 'T':
			ok = parseCurve(cmd, rel, sc, &x, &y, &lastC, &lastQ, &hasC, &hasQ, bezier)
		case 'A':
			// 圆弧: 端点参数化转多段三次贝塞尔(每段圆心角 <= 90°)。
			// 参数依次 rx ry rot largeArc sweep x y —— 中间两个标志位是**单字符**,
			// 必须逐字符取("a.25.25 0 00.407.195" 这样的连写只能这么读)。
			var v [7]float32
			for k := 0; k < 7; k++ {
				var (
					f float32
					o bool
				)
				if k == 3 || k == 4 {
					f, o = sc.flag()
				} else {
					f, o = sc.num()
				}
				if !o {
					ok = false
					break
				}
				v[k] = f
			}
			if ok {
				ex, ey := v[5], v[6]
				if rel {
					ex, ey = x+ex, y+ey
				}
				arcToBeziers(x, y, v[0], v[1], v[2], v[3] != 0, v[4] != 0, ex, ey, push, bezier)
				x, y = ex, ey
			}
			hasC, hasQ = false, false
		default:
			flush()
			return subs, &PathError{Pos: sc.pos, Cmd: cmd, Why: "不认识的命令"}
		}
		if !ok {
			flush()
			return subs, &PathError{Pos: sc.pos, Cmd: cmd, Why: "参数不足或数字格式不对"}
		}
	}
	flush()
	return subs, nil
}

// ---------------- 扫描器 ----------------

type scanner struct {
	s   string
	pos int
}

func (sc *scanner) skipSep() {
	for sc.pos < len(sc.s) {
		switch sc.s[sc.pos] {
		case ' ', ',', '\n', '\r', '\t':
			sc.pos++
		default:
			return
		}
	}
}

// cmd 取下一个命令字母(没有则返回 false, 表示沿用上一个命令).
func (sc *scanner) cmd() (byte, bool) {
	sc.skipSep()
	if sc.pos >= len(sc.s) {
		return 0, false
	}
	c := sc.s[sc.pos]
	if unicode.IsLetter(rune(c)) {
		sc.pos++
		return c, true
	}
	return 0, false
}

// num 取下一个数字(支持 "-4.3-4.3" 这类连写).
//
// 遇到第二个小数点就停: SVG 允许 ".25.25" 表示两个数(.25 与 .25), 贪婪吃到第二个 '.'
// 会让解析失败并把后面整段图形丢掉(图标站的 path 里很常见, 症状是"只画出一半")。
func (sc *scanner) num() (float32, bool) {
	sc.skipSep()
	st := sc.pos
	if sc.pos < len(sc.s) && (sc.s[sc.pos] == '-' || sc.s[sc.pos] == '+') {
		sc.pos++
	}
	dot := false
	for sc.pos < len(sc.s) {
		c := sc.s[sc.pos]
		if c >= '0' && c <= '9' {
			sc.pos++
			continue
		}
		if c == '.' && !dot {
			dot = true
			sc.pos++
			continue
		}
		break
	}
	if sc.pos == st {
		return 0, false
	}
	v, err := strconv.ParseFloat(sc.s[st:sc.pos], 32)
	if err != nil {
		return 0, false
	}
	return float32(v), true
}

// flag 取一个 arc 标志位: SVG 规定 large-arc/sweep 是**单字符** 0/1,
// 所以 "00.407" 必须按 0 / 0 / .407 读, 不能按数字切(见 num 的注释)。
func (sc *scanner) flag() (float32, bool) {
	sc.skipSep()
	if sc.pos >= len(sc.s) || (sc.s[sc.pos] != '0' && sc.s[sc.pos] != '1') {
		return 0, false
	}
	c := sc.s[sc.pos]
	sc.pos++
	return float32(c - '0'), true
}

// ---------------- 曲线与圆弧 ----------------

// parseCurve 处理三次/二次贝塞尔相关命令(C/S/Q/T), 返回是否解析成功。
func parseCurve(cmd byte, rel bool, sc *scanner,
	x, y *float32, lastC, lastQ *Point, hasC, hasQ *bool,
	bezier func(float32, float32, float32, float32, float32, float32, float32, float32)) bool {

	// 参数个数: C=6, S/Q=4, T=2
	upper := unicode.ToUpper(rune(cmd))
	n := 2
	switch upper {
	case 'C':
		n = 6
	case 'S', 'Q':
		n = 4
	}
	var v [6]float32
	for k := 0; k < n; k++ {
		f, ok := sc.num()
		if !ok {
			return false
		}
		v[k] = f
	}
	ex, ey := v[n-2], v[n-1]
	if rel {
		ex, ey = *x+ex, *y+ey
	}

	switch upper {
	case 'C':
		c1x, c1y, c2x, c2y := v[0], v[1], v[2], v[3]
		if rel {
			c1x, c1y, c2x, c2y = *x+c1x, *y+c1y, *x+c2x, *y+c2y
		}
		bezier(*x, *y, c1x, c1y, c2x, c2y, ex, ey)
		*lastC, *hasC, *hasQ = Point{c2x, c2y}, true, false
	case 'S':
		c2x, c2y := v[0], v[1]
		c1x, c1y := *x, *y
		if rel {
			c2x, c2y = *x+c2x, *y+c2y
		}
		if *hasC { // 反射上一个控制点
			c1x, c1y = 2**x-lastC.X, 2**y-lastC.Y
		}
		bezier(*x, *y, c1x, c1y, c2x, c2y, ex, ey)
		*lastC, *hasC, *hasQ = Point{c2x, c2y}, true, false
	case 'Q', 'T':
		var c1x, c1y float32
		if upper == 'Q' {
			c1x, c1y = v[0], v[1]
			if rel {
				c1x, c1y = *x+c1x, *y+c1y
			}
		} else {
			c1x, c1y = *x, *y
			if *hasQ { // 反射上一个控制点
				c1x, c1y = 2**x-lastQ.X, 2**y-lastQ.Y
			}
		}
		// 二次转三次(控制点按 2/3 位置内插)
		bezier(*x, *y,
			*x+2*(c1x-*x)/3, *y+2*(c1y-*y)/3,
			ex+2*(c1x-ex)/3, ey+2*(c1y-ey)/3,
			ex, ey)
		*lastQ, *hasQ, *hasC = Point{c1x, c1y}, true, false
	}
	*x, *y = ex, ey
	return true
}

// arcToBeziers 把 SVG 圆弧指令(A, 端点参数化)转换为若干段三次贝塞尔折线点.
//
// 端点/圆心换算与角度拆分见 W3C SVG 1.1 附录 F.6.5/F.6.6; 每段圆心角不超过 90°,
// 用 4/3·tan(θ/4) 的切向量长度做三次近似(与整圆图元肉眼无差).
func arcToBeziers(x1, y1, rx, ry, phiDeg float32, largeArc, sweep bool, x2, y2 float32,
	push func(float32, float32),
	bezier func(x1, y1, x2, y2, x3, y3, x4, y4 float32)) {

	if x1 == x2 && y1 == y2 {
		return
	}
	rx, ry = abs32(rx), abs32(ry)
	if rx == 0 || ry == 0 { // 半径为 0 退化为直线
		push(x2, y2)
		return
	}
	phi := float64(phiDeg) * math.Pi / 180
	cosP, sinP := float32(math.Cos(phi)), float32(math.Sin(phi))

	// 端点换到以椭圆中点为原点、沿椭圆主轴的坐标系
	dx, dy := (x1-x2)/2, (y1-y2)/2
	x1p := cosP*dx + sinP*dy
	y1p := -sinP*dx + cosP*dy

	// 半径过小则按端点距离等比放大
	if lam := x1p*x1p/(rx*rx) + y1p*y1p/(ry*ry); lam > 1 {
		s := float32(math.Sqrt(float64(lam)))
		rx, ry = rx*s, ry*s
	}

	// 圆心(换算回原坐标系)。两个候选圆心对称于弦, 符号按 SVG 规范 F.6.5:
	// large-arc 与 sweep 不相同时取 +√q, 相同时取 -√q(取错会绕远弧, 图形整个乱掉)。
	sign := float32(-1)
	if largeArc != sweep {
		sign = 1
	}
	nm := rx*rx*ry*ry - rx*rx*y1p*y1p - ry*ry*x1p*x1p
	den := rx*rx*y1p*y1p + ry*ry*x1p*x1p
	// 半径按端点距离放大后 nm 理论上为 0, 浮点误差下可能是极小负数, 必须钳到 0
	var co float32
	if q := nm / den; q > 0 {
		co = sign * float32(math.Sqrt(float64(q)))
	}
	cxp := co * rx * y1p / ry
	cyp := -co * ry * x1p / rx
	cx := cosP*cxp - sinP*cyp + (x1+x2)/2
	cy := sinP*cxp + cosP*cyp + (y1+y2)/2

	ang := func(ux, uy, vx, vy float32) float32 {
		c := (ux*vx + uy*vy) / float32(math.Sqrt(float64((ux*ux+uy*uy)*(vx*vx+vy*vy))))
		if c > 1 {
			c = 1
		} else if c < -1 {
			c = -1
		}
		a := float32(math.Acos(float64(c)))
		if ux*vy-uy*vx < 0 {
			a = -a
		}
		return a
	}
	ux, uy := (x1p-cxp)/rx, (y1p-cyp)/ry
	vx, vy := (-x1p-cxp)/rx, (-y1p-cyp)/ry
	th1 := ang(1, 0, ux, uy)
	dth := ang(ux, uy, vx, vy)
	if !sweep && dth > 0 {
		dth -= 2 * math.Pi
	} else if sweep && dth < 0 {
		dth += 2 * math.Pi
	}

	// 拆成不超过 90° 的段逐段转贝塞尔; 点函数 E(a) 是椭圆上的点, D(a) 是切向导数
	n := int(math.Ceil(math.Abs(float64(dth))/(math.Pi/2))) + 1
	seg := dth / float32(n)
	k := float32(4.0 / 3.0 * math.Tan(float64(seg)/4))
	e := func(a float32) (float32, float32) {
		ca, sa := float32(math.Cos(float64(a))), float32(math.Sin(float64(a)))
		return cx + rx*cosP*ca - ry*sinP*sa, cy + rx*sinP*ca + ry*cosP*sa
	}
	d := func(a float32) (float32, float32) {
		sa, ca := float32(math.Sin(float64(a))), float32(math.Cos(float64(a)))
		return -rx*cosP*sa - ry*sinP*ca, -rx*sinP*sa + ry*cosP*ca
	}
	for i := 0; i < n; i++ {
		a1, a2 := th1+seg*float32(i), th1+seg*float32(i+1)
		p1x, p1y := e(a1)
		p2x, p2y := e(a2)
		d1x, d1y := d(a1)
		d2x, d2y := d(a2)
		bezier(p1x, p1y, p1x+k*d1x, p1y+k*d1y, p2x-k*d2x, p2y-k*d2y, p2x, p2y)
	}
}

func abs32(v float32) float32 {
	if v < 0 {
		return -v
	}
	return v
}

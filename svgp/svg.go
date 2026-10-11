package svgp

import (
	"encoding/xml"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

// Doc 从一整段 SVG 文本里收集到的东西(只认 svg / g / path 三个元素).
type Doc struct {
	// Paths 全部 <path> 解析成的折线子路径; Filled 已按元素(含 <g> 继承)的 fill/stroke 判定好.
	Paths []SubPath
	// ViewBox 是 min-x, min-y, width, height; 没写 viewBox 时退化成 width/height, 都没有则全 0
	// (此时 ScaleTo 会原样返回). 缩放不在这里做 —— 用 ScaleTo 显式换算.
	ViewBox [4]float64
	// StrokeWidth 是元素上的 stroke-width(**源网格单位**); 没写则 0.
	// 要换算到目标网格(比如 24)就乘 ScaleFactor(doc.ViewBox, 24).
	StrokeWidth float64
}

// ParseDoc 解析一整段 SVG 文本(从图标站复制来的 `<svg viewBox="…"><path d="…"/></svg>` 那种).
//
//   - 识别 <path> 的 d / fill / stroke / stroke-width, <g> 上的同名属性会被子元素继承;
//   - 单条 path 解析失败只跳它自己, 不牵连整张图; 全都没有可用 path 时返回 error;
//   - <circle>/<rect>/<line> 等形状元素、transform 属性、CSS 样式表一律**忽略**(不报错).
func ParseDoc(svg string) (*Doc, error) {
	dec := xml.NewDecoder(strings.NewReader(svg))
	dec.Strict = false // SVG 里常有未声明的实体/宽松写法

	// <g> 的属性继承栈
	type inherited struct{ fill, stroke, sw string }
	stack := []inherited{{}}
	cur := func() inherited { return stack[len(stack)-1] }
	get := func(a []xml.Attr, name string) string {
		for _, at := range a {
			if strings.EqualFold(at.Name.Local, name) {
				return strings.TrimSpace(at.Value)
			}
		}
		return ""
	}

	doc := &Doc{}
	hasBox, hasSize := false, false

	for {
		tok, e := dec.Token()
		if e == io.EOF {
			break
		}
		if e != nil {
			return nil, fmt.Errorf("svgp: SVG 解析失败: %w", e)
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch strings.ToLower(t.Name.Local) {
			case "svg":
				if vb := get(t.Attr, "viewBox"); vb != "" {
					if f := parseNums(vb); len(f) == 4 && f[2] > 0 && f[3] > 0 {
						doc.ViewBox = [4]float64{f[0], f[1], f[2], f[3]}
						hasBox = true
					}
				}
				if !hasBox && !hasSize {
					w := parseLen(get(t.Attr, "width"))
					h := parseLen(get(t.Attr, "height"))
					if w > 0 && h > 0 {
						doc.ViewBox = [4]float64{0, 0, w, h}
						hasSize = true
					}
				}
				if v := parseLen(get(t.Attr, "stroke-width")); v > 0 {
					doc.StrokeWidth = v
					stack[len(stack)-1].sw = get(t.Attr, "stroke-width")
				}
			case "g":
				n := cur()
				if v := get(t.Attr, "fill"); v != "" {
					n.fill = v
				}
				if v := get(t.Attr, "stroke"); v != "" {
					n.stroke = v
				}
				if v := get(t.Attr, "stroke-width"); v != "" {
					n.sw = v
				}
				stack = append(stack, n)
			case "path":
				d := get(t.Attr, "d")
				if d == "" {
					continue
				}
				n := cur()
				fill, stroke, sws := n.fill, n.stroke, n.sw
				if v := get(t.Attr, "fill"); v != "" {
					fill = v
				}
				if v := get(t.Attr, "stroke"); v != "" {
					stroke = v
				}
				if v := get(t.Attr, "stroke-width"); v != "" {
					sws = v
				}
				if doc.StrokeWidth <= 0 {
					if v := parseLen(sws); v > 0 {
						doc.StrokeWidth = v
					}
				}
				// 判定"填充还是描边": 没有 stroke 且 fill 非 none → 填充; 其余(尤其
				// fill="none" stroke="currentColor")→ 描边 —— 图标集里后者是绝对主流.
				filled := fill != "" && !strings.EqualFold(fill, "none") &&
					(stroke == "" || strings.EqualFold(stroke, "none"))
				subs, err := ParsePath(d)
				if err != nil { // 单条 path 解析不动就跳过它
					continue
				}
				for _, s := range subs {
					s.Filled = filled
					doc.Paths = append(doc.Paths, s)
				}
			}
		case xml.EndElement:
			if strings.EqualFold(t.Name.Local, "g") && len(stack) > 1 {
				stack = stack[:len(stack)-1]
			}
		}
	}
	if len(doc.Paths) == 0 {
		return nil, fmt.Errorf("svgp: SVG 里没有可用的 <path d=…>")
	}
	return doc, nil
}

// ParseDocSubPaths 只要折线时的便捷入口(等价于 ParseDoc 之后取 Doc.Paths).
func ParseDocSubPaths(svg string) ([]SubPath, error) {
	doc, err := ParseDoc(svg)
	if err != nil {
		return nil, err
	}
	return doc.Paths, nil
}

// ScaleFactor 返回"viewBox → target(较长边)"的缩放系数; viewBox 无效或 target <= 0 时返回 1.
//
// 用法: 线宽也要一起换算 —— 24 网格下的线宽 = doc.StrokeWidth * ScaleFactor(doc.ViewBox, 24).
func ScaleFactor(viewBox [4]float64, target float64) float64 {
	if target <= 0 || viewBox[2] <= 0 || viewBox[3] <= 0 {
		return 1
	}
	return target / math.Max(viewBox[2], viewBox[3])
}

// ScaleTo 把子路径从 viewBox 坐标换算到"较长边为 target、左上角对齐原点"的坐标系(保持比例).
//
// viewBox 的 min-x/min-y 会一并平掉; viewBox 无效或 target <= 0 时**原样返回**(不复制).
func ScaleTo(subs []SubPath, viewBox [4]float64, target float32) []SubPath {
	if target <= 0 || viewBox[2] <= 0 || viewBox[3] <= 0 {
		return subs
	}
	s := ScaleFactor(viewBox, float64(target))
	out := make([]SubPath, len(subs))
	for i, sub := range subs {
		pts := make([]Point, len(sub.Points))
		for j, p := range sub.Points {
			pts[j] = Point{
				X: float32((float64(p.X) - viewBox[0]) * s),
				Y: float32((float64(p.Y) - viewBox[1]) * s),
			}
		}
		out[i] = SubPath{Points: pts, Filled: sub.Filled, Closed: sub.Closed}
	}
	return out
}

// parseNums 从 "0 0 24 24"(也容忍逗号)里取一串数字.
func parseNums(s string) []float64 {
	s = strings.ReplaceAll(s, ",", " ")
	fields := strings.Fields(s)
	out := make([]float64, 0, len(fields))
	for _, f := range fields {
		if v, err := strconv.ParseFloat(f, 64); err == nil {
			out = append(out, v)
		}
	}
	return out
}

// parseLen 解析长度值("24" / "24px" / "1.5"), 拿不到就返回 0.
func parseLen(s string) float64 {
	s = strings.TrimSpace(s)
	s = strings.TrimSuffix(s, "px")
	if v, err := strconv.ParseFloat(s, 64); err == nil {
		return v
	}
	return 0
}

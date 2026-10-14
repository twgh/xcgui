package svgp

import (
	"errors"
	"fmt"
	"image"
	"sync"

	"github.com/twgh/xcgui/wapi/d2d"
	"github.com/twgh/xcgui/wapi/gdi"
	"github.com/twgh/xcgui/wapi/types"
)

// ==================== 离屏渲染: 子路径 → RGBA 位图 ====================
//
// 这条路径的用途: 把 SVG 画成一张**图片**(而不是直接画到窗口), 之后交给任何吃图片句柄的 API
// (xc.XImage_LoadFromGoImage / XImage_LoadFromHBITMAP …).
//
// 它自建 D2D 工厂(见 d2d.D2D1CreateFactory), 与炫彩的渲染管线完全无关 —— 因此:
//   - GDI 渲染模式(app.New(false))下同样可用;
//   - "还没建窗口"时也能用;
//   - 不会因炫彩重建 D2D 设备而失效.
//
// 工厂是 SINGLE_THREADED 的(建资源不带锁, 更快), 所以整条渲染流程用一把互斥锁串行化.
// 图标生成是低频操作(调用方建一次、缓存复用), 串行化的代价无关紧要.

var (
	offMu          sync.Mutex            // 保护整条离屏渲染流程 + 下面这些单例
	offFactory     *d2d.ID2D1Factory     // 离屏渲染的共享工厂(进程生命周期, 不释放)
	offFactoryErr  error                 // 工厂创建失败的原因(记住, 不反复重试)
	offStrokeStyle *d2d.ID2D1StrokeStyle // SVG 语义的圆头圆角样式(建不出来时为 nil)
	offTargets     []*offscreenTarget    // 按尺寸复用的"位图 + DC 渲染目标"(见 acquireTarget)
)

// offscreenTarget 是一套可复用的离屏绘制目标: DIB(像素仓库) + 绑好该 DIB 的 DC 渲染目标.
//
// 为什么要复用: 现建一套(CreateDIBSection + CreateDCRenderTarget)约 1ms, 是 32px 图标整条
// 渲染路径的主要成本; 复用之后每次只剩"绘制 + 拷像素"(约 1/3). 像素每次都清(见 RenderToRGBA
// 的 Clear), 所以复用不会串上上一次的内容.
type offscreenTarget struct {
	size int32
	dib  *gdi.DIB
	rt   *d2d.ID2D1DCRenderTarget
}

// maxOffscreenTargets 是复用目标的个数上限(图标尺寸种类很少; 超了按 LRU 丢最久未用的).
const maxOffscreenTargets = 4

// acquireTarget 取(或新建)该尺寸的离屏目标(调用方必须持有 offMu).
func acquireTarget(f *d2d.ID2D1Factory, size int32) (*offscreenTarget, error) {
	for i, t := range offTargets {
		if t.size == size {
			// 命中: 挪到末尾(LRU 的"最近使用"端)
			copy(offTargets[i:], offTargets[i+1:])
			offTargets[len(offTargets)-1] = t
			return t, nil
		}
	}
	dib, err := gdi.NewDIB(size, size)
	if err != nil {
		return nil, err
	}
	props := d2d.NewDCRenderTargetProperties()
	rt, err := f.CreateDCRenderTarget(&props)
	if err != nil {
		dib.Close()
		return nil, err
	}
	if err := rt.BindDC(dib.HDC(), &types.RECT{Left: 0, Top: 0, Right: size, Bottom: size}); err != nil {
		rt.Release()
		dib.Close()
		return nil, err
	}
	t := &offscreenTarget{size: size, dib: dib, rt: rt}
	if len(offTargets) >= maxOffscreenTargets { // 满了: 丢最久未用的那套
		offTargets[0].release()
		offTargets = offTargets[1:]
	}
	offTargets = append(offTargets, t)
	return t, nil
}

// release 释放这套目标.
func (t *offscreenTarget) release() {
	if t == nil {
		return
	}
	t.rt.Release() // 先放渲染目标(它引用着 DC), 再删位图
	t.dib.Close()
}

// offscreenD2D 取离屏渲染用的共享工厂(调用方必须持有 offMu).
func offscreenD2D() (*d2d.ID2D1Factory, error) {
	if offFactory != nil || offFactoryErr != nil {
		return offFactory, offFactoryErr
	}
	f, err := d2d.D2D1CreateFactory(d2d.D2D1_FACTORY_TYPE_SINGLE_THREADED)
	if err != nil {
		offFactoryErr = err
		return nil, err
	}
	// SVG 的圆角语义(stroke-linecap/linejoin="round")是图标集的标准; 建失败也不致命
	// (绘制时传 nil = 默认平头斜接, 只是角不圆).
	props := d2d.RoundStrokeStyleProperties()
	if s, err := f.CreateStrokeStyle(&props); err == nil {
		offStrokeStyle = s
	}
	offFactory = f
	return f, nil
}

// RenderToRGBA 把子路径离屏渲染成一张 size×size 的 RGBA 位图.
//
// 参数:
//   - subs: 已经**换算到目标像素网格**的子路径(用 ScaleTo(subs, viewBox, size) 得到);
//     坐标单位即像素, 原点在左上角;
//   - size: 输出位图的边长(物理像素, 正方形);
//   - override: HasColor=false 的子路径(currentColor / 没写颜色)用的颜色;
//   - strokeWidth: 线宽, 单位与 subs 的坐标一致(即像素); 由调用方换算好再传
//     (SVG 的 24 网格线宽 → 像素要乘 ScaleFactor), <=0 表示不描边(填充形状不受影响);
//   - style: 描边样式(圆头 / 圆角 / 虚线); **传 nil 用 SVG 惯例的圆头圆角**.
//
// 返回的图像是**非预乘**的 RGBA(D2D 只输出预乘 BGRA, 这里已反预乘), 与 image.RGBA 的约定
// 一致 —— 可以直接喂 xc.XImage_LoadFromGoImage.
//
// 环境里没有 D2D(d2d1.dll 不可用)或任何一步失败都会返回 error, 调用方应回落(而不是硬用).
func RenderToRGBA(subs []SubPath, size int32, override d2d.D2D1_COLOR_F, strokeWidth float32, style *d2d.ID2D1StrokeStyle) (*image.RGBA, error) {
	if size <= 0 {
		return nil, fmt.Errorf("svgp: 图片边长必须为正, 得到 %d", size)
	}
	if len(subs) == 0 {
		return nil, errors.New("svgp: 没有可渲染的子路径")
	}

	offMu.Lock()
	defer offMu.Unlock()

	f, err := offscreenD2D()
	if err != nil {
		return nil, err
	}
	if style == nil {
		style = offStrokeStyle // 可能仍为 nil(样式没建出来), 绘制层会退成默认样式
	}

	groups, err := NewColorGroups(f, subs, 1, override)
	if err != nil {
		return nil, err
	}
	defer ReleaseColorGroups(groups)

	// 取该尺寸的复用目标(DIB + 绑好的 DC 渲染目标); 像素每次都会清, 不会串上一次的内容.
	target, err := acquireTarget(f, size)
	if err != nil {
		return nil, err
	}
	rt, dib := target.rt, target.dib

	// 绘制层吃的是 ID2D1RenderTarget 基类视图(值嵌入, 取嵌入字段的地址即可, 布局与 COM 一致).
	base := &rt.ID2D1RenderTarget
	rt.BeginDraw()
	rt.Clear(nil) // DIB 内容可能是上一次的/不保证是零: 先清成透明黑
	for _, g := range groups {
		if err := g.Geom.Draw(base, g.Color, strokeWidth, style); err != nil {
			rt.EndDraw() // 已进入绘制态, 失败也要收尾再走
			return nil, err
		}
	}
	if err := rt.EndDraw(); err != nil {
		return nil, err
	}

	return dibToRGBA(dib, size)
}

// dibToRGBA 把 DIB 里的**预乘 BGRA** 像素转成**非预乘 RGBA** 位图.
//
// D2D 的 DC 渲染目标只输出预乘 alpha(颜色分量已经乘过 a/255), 而 Go 的 image.RGBA 与
// 炫彩的图片 API 都按非预乘理解, 所以要反预乘一趟: c = round(c×255/a), 钳到 255.
// a=0 的像素整像素置 0.
func dibToRGBA(dib *gdi.DIB, size int32) (*image.RGBA, error) {
	src := dib.Pixels()
	if len(src) < int(size)*int(size)*4 {
		return nil, errors.New("svgp: DIB 像素不完整")
	}
	img := image.NewRGBA(image.Rect(0, 0, int(size), int(size)))
	for i := 0; i < len(img.Pix); i += 4 {
		b, g, r, a := src[i], src[i+1], src[i+2], src[i+3]
		if a == 0 {
			continue // 预乘下 a=0 必是全 0; 目标里保持 (0,0,0,0)
		}
		if a != 255 {
			r, g, b = unpremul(r, a), unpremul(g, a), unpremul(b, a)
		}
		img.Pix[i], img.Pix[i+1], img.Pix[i+2], img.Pix[i+3] = r, g, b, a
	}
	return img, nil
}

// unpremul 把一个预乘过的分量还原(c = round(c×255/a), 钳到 255 —— 预乘那一步本身做过舍入,
// 反算结果可能略微超出范围).
func unpremul(c, a byte) byte {
	v := (int(c)*255 + int(a)/2) / int(a)
	if v > 255 {
		return 255
	}
	return byte(v)
}

package widget

import (
	"errors"
	"io"
	"sync"
	"time"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// GifPlayer Gif 播放器.
type GifPlayer struct {
	HImages         []int           // Gif 帧图片句柄
	Delays          []time.Duration // Gif 帧延迟
	ImmediateRedraw bool            // 控制在播放 Gif 帧时是否立即重绘, 默认为 false.
}

// NewGifPlayer 创建 Gif 播放器.
//   - 会加载 Gif 的所有帧图片到 GifPlayer.HImages.
//   - 图片不再使用时, 可调用 GifPlayer.ReleaseImages 释放所有帧图片句柄.
func NewGifPlayer(gifReader io.Reader) (*GifPlayer, error) {
	// 读取 GIF 数据
	frames, err := common.ExtractGifFrames(gifReader)
	if err != nil {
		return nil, errors.New("读取 GIF 数据时出错: " + err.Error())
	}
	frameCount := len(frames)

	gp := &GifPlayer{}
	gp.HImages = make([]int, 0, frameCount)
	gp.Delays = make([]time.Duration, 0, frameCount)
	// 将 GIF 帧加载为图片句柄
	for i := 0; i < frameCount; i++ {
		hImage := xc.XImage_LoadMemory(frames[i].ImageData)
		if hImage > 0 {
			gp.HImages = append(gp.HImages, hImage)
			gp.Delays = append(gp.Delays, frames[i].Delay)
		}
	}
	return gp, nil
}

// Play 播放 Gif, 返回 GifPlayerHandler.
//   - 如果你想改变 Gif 的宽高, 可改变 hEle 的宽高然后启用 fullEle, 或者操作 GifPlayer.HImages.
//   - 内部注册了元素绘制事件.
//
// hEle: 元素句柄. 会给元素添加绘制事件, 可在不需要时调用 GifPlayerHandler.Destroy.
//
// fullEle: 是否填满元素. true: 自适应填满元素. false: 保持原大小.
//
// loopCount: 控制 Gif 在显示期间重新启动的次数。
//   - < 1: 永远循环。
//   - > 0: 循环 loopCount 次。
//
// onFrame: 帧事件, 此事件中的代码是在 UI 线程执行的. 可不填.
func (p *GifPlayer) Play(hEle int, fullEle bool, loopCount int, onFrame ...func(h *GifPlayerHandler, frame int)) *GifPlayerHandler {
	gph := NewGifPlayerHandler()
	gph.hEle = hEle

	// 添加元素绘制事件.
	ele := NewElementByHandle(hEle)
	ele.AddEvent_Paint(func(hEle int, hDraw int, pbHandled *bool) int {
		*pbHandled = true
		if fullEle {
			var rc xc.RECT
			xc.XEle_GetRect(hEle, &rc)
			rc.Left = 0
			rc.Top = 0
			xc.XDraw_ImageAdaptive(hDraw, p.HImages[gph.GetCurrentFrame()], &rc, false)
		} else {
			xc.XDraw_Image(hDraw, p.HImages[gph.GetCurrentFrame()], 0, 0)
		}
		return 0
	}, true)

	// 帧事件
	if len(onFrame) > 0 {
		gph.OnFrame = onFrame[0]
	}
	// 最大帧索引
	gph.maxFrame = len(p.HImages) - 1

	gph.stopped = false
	gph.paused = false
	go func() {
		for {
			if gph.isReturn() {
				return
			}

			// 更新帧
			xc.XC_CallUT(func() {
				if gph.IsStopped() {
					return
				}
				if gph.OnFrame != nil {
					gph.OnFrame(gph, gph.GetCurrentFrame())
				}
				if xc.XC_IsHELE(hEle) {
					xc.XEle_Redraw(hEle, p.ImmediateRedraw)
				}
			})

			// 帧延迟
			time.Sleep(p.Delays[gph.GetCurrentFrame()])

			if gph.isReturn() {
				return
			}

			// 更新当前帧索引
			newFrame := gph.GetCurrentFrame() + 1

			if newFrame > gph.GetMaxFrame() {
				// 处理循环次数, loopCount > 0: 循环指定次数
				if loopCount > 0 {
					if loopCount--; loopCount == 0 {
						gph.mu.Lock()
						gph.stopped = true // 标记为已停止
						gph.mu.Unlock()
						break
					}
				}
				gph.SetCurrentFrame(0)
			} else {
				// 设置当前帧索引
				gph.SetCurrentFrame(newFrame)
			}
		}
	}()
	return gph
}

// ReleaseImages 释放已加载的 Gif 所有帧图片句柄.
func (p *GifPlayer) ReleaseImages() {
	for i := 0; i < len(p.HImages); i++ {
		xc.XImage_Release(p.HImages[i])
	}
}

// GifPlayerHandler 用于操作 Gif 播放器.
type GifPlayerHandler struct {
	// 帧事件, 此事件是在 UI 线程执行的
	OnFrame func(h *GifPlayerHandler, frame int)

	rwx sync.RWMutex

	mu   sync.Mutex
	cond *sync.Cond

	hEle     int // 元素句柄
	curFrame int // 当前帧索引
	maxFrame int // 最大帧索引

	paused  bool
	stopped bool
}

func NewGifPlayerHandler() *GifPlayerHandler {
	gph := &GifPlayerHandler{}
	gph.cond = sync.NewCond(&gph.mu)
	return gph
}

func (g *GifPlayerHandler) isReturn() bool {
	// 状态检查
	g.mu.Lock()
	// 优先检查停止状态
	if g.stopped {
		g.mu.Unlock()
		return true
	}
	// 严格循环检查暂停条件
	for g.paused && !g.stopped {
		g.cond.Wait()
	}
	// 再次确认是否停止
	if g.stopped {
		g.mu.Unlock()
		return true
	}
	g.mu.Unlock()
	return false
}

// SetCurrentFrame 设置当前帧索引.
func (g *GifPlayerHandler) SetCurrentFrame(frame int) {
	if frame < 0 {
		frame = 0
	}
	g.rwx.Lock()
	g.curFrame = frame
	g.rwx.Unlock()
}

// GetCurrentFrame 获取当前帧索引.
func (g *GifPlayerHandler) GetCurrentFrame() int {
	g.rwx.RLock()
	defer g.rwx.RUnlock()
	return g.curFrame
}

// SetMaxFrame 设置最大帧索引.
func (g *GifPlayerHandler) SetMaxFrame(frame int) {
	if frame < 0 {
		return
	}
	g.rwx.Lock()
	g.maxFrame = frame
	g.rwx.Unlock()
}

// GetMaxFrame 获取最大帧索引.
func (g *GifPlayerHandler) GetMaxFrame() int {
	g.rwx.RLock()
	defer g.rwx.RUnlock()
	return g.maxFrame
}

// Destroy 会停止播放且移除添加给元素的绘制事件.
func (g *GifPlayerHandler) Destroy() {
	g.Stop()
	ele := NewElementByHandle(g.hEle)
	xc.Auto(func() {
		ele.RemoveEvent(xcc.XE_PAINT)
		ele.Redraw(true)
	})
}

// Stop 停止播放.
func (g *GifPlayerHandler) Stop() {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.stopped = true
	g.paused = false
	g.cond.Broadcast()
}

// Pause 暂停播放.
func (g *GifPlayerHandler) Pause() {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.paused = true
}

// Resume 恢复播放.
func (g *GifPlayerHandler) Resume() {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.paused = false
	g.cond.Signal()
}

// IsPaused 是否已暂停.
func (g *GifPlayerHandler) IsPaused() bool {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.paused
}

// IsStopped 是否已停止.
func (g *GifPlayerHandler) IsStopped() bool {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.stopped
}

package widget

import (
	"errors"
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
	"io"
	"sync"
	"time"
)

// GifPlayer Gif 播放器.
type GifPlayer struct {
	HImages []int           // Gif 帧图片句柄
	Delays  []time.Duration // Gif 帧延迟
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

	gp := &GifPlayer{}
	// 将 GIF 帧加载为图片句柄
	for i := 0; i < len(frames); i++ {
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
//
// hEle: 元素句柄. 会给元素添加绘制事件, 可在不需要时调用 GifPlayerHandler.Destroy.
//
// fullEle: 是否填满元素. true: 自适应填满元素. false: 保持原大小.
//
// loopCount: 控制 Gif 在显示期间重新启动的次数。默认为0.
//   - < 1: 永远循环。
//   - > 0: 循环 loopCount 次。
func (p *GifPlayer) Play(hEle int, fullEle bool, loopCount ...int) *GifPlayerHandler {
	gph := NewGifPlayerHandler()
	gph.hEle = hEle

	// 添加元素绘制事件.
	ele := NewElementByHandle(hEle)
	ele.AddEvent_Paint(func(hEle int, hDraw int, pbHandled *bool) int {
		// 在自绘事件函数中, 用户手动调用绘制元素, 以便控制绘制顺序
		xc.XEle_DrawEle(hEle, hDraw)

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
	})

	// Gif 循环次数
	LoopCount := 0
	if len(loopCount) > 0 {
		LoopCount = loopCount[0]
	}
	// 最大帧索引
	gph.maxFrame = len(p.HImages)

	gph.stopped = false
	gph.paused = false
	go func() {
		for {
			// 状态检查
			gph.mu.Lock()
			// 优先检查停止状态
			if gph.stopped {
				gph.mu.Unlock()
				return
			}
			// 严格循环检查暂停条件
			for gph.paused && !gph.stopped {
				gph.cond.Wait()
			}
			// 再次确认是否停止
			if gph.stopped {
				gph.mu.Unlock()
				return
			}
			gph.mu.Unlock()

			// 帧延迟
			time.Sleep(p.Delays[gph.GetCurrentFrame()])

			// 设置当前帧索引
			gph.SetCurrentFrame((gph.GetCurrentFrame() + 1) % gph.GetMaxFrame())

			// 更新到下一帧
			xc.XC_CallUT(func() {
				if xc.XC_IsHELE(hEle) {
					xc.XEle_Redraw(hEle, false)
				}
			})

			// 处理循环次数
			if LoopCount > 0 && gph.GetCurrentFrame() == gph.GetMaxFrame()-1 { // >0: 循环指定次数
				if LoopCount--; LoopCount == 0 {
					gph.mu.Lock()
					gph.stopped = true // 标记为已停止
					gph.mu.Unlock()
					break
				}
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

// SetCurrentFrame 设置当前帧索引.
func (g *GifPlayerHandler) SetCurrentFrame(frame int) {
	if frame < 0 || frame >= g.maxFrame {
		return
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
	ele.RemoveEvent(xcc.XE_PAINT)
	ele.Redraw(false)
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

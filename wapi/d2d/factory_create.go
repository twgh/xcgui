package d2d

import (
	"errors"
	"fmt"
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/wapi/com"
)

// 本文件封装 d2d1.dll 的 flat API `D2D1CreateFactory` —— 用户自建 D2D 工厂的唯一入口.
//
// 为什么不借 xcgui 的工厂(见 NewFactoryByHandle): 那个工厂跟着炫彩的 D2D 渲染走, GDI 渲染
// 模式(app.New(false))下压根没有, 也没法在"还没建窗口"时使用. 自建的工厂与炫彩完全无关,
// 因此在任何模式下都能用 —— 离屏渲染(把 SVG 画成图片句柄)依赖这一点.

var (
	// D2D1 是 d2d1.dll. 用 LazyDLL: 只有真正调用本文件的方法时才会去加载它(与借工厂的路径无关).
	D2D1 = syscall.NewLazyDLL("d2d1.dll")

	procD2D1CreateFactory = D2D1.NewProc("D2D1CreateFactory")
)

const IID_ID2D1Factory = "06152247-6f50-465a-9245-118bfd3b6007"

// D2D1CreateFactory 创建一个**独立于 xcgui** 的 D2D 工厂(对应 d2d1.dll 的 flat API, 一进一出).
//
// 拿到的工厂由调用方持有: 长期复用没问题(设备无关资源都挂在它上面), 不用时必须 Release ——
// 通常是"进程级单例, 不释放". 线程模型自选: 单线程工厂(SINGLE_THREADED, 更快)要求调用方
// 自己串行化建资源, 多线程工厂自己加锁.
//
// 它是 `app.New(true/false)` 之外的一条路: 不需要炫彩初始化、不需要窗口, 也不会因为炫彩
// 重建 D2D 设备而失效.
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/nf-d2d1-d2d1createfactory
func D2D1CreateFactory(factoryType D2D1_FACTORY_TYPE) (*ID2D1Factory, error) {
	// Find 先查一遍 dll/导出表: LazyProc.Call 找不到时会直接 panic, 而调用方要的是"能失败".
	if err := procD2D1CreateFactory.Find(); err != nil {
		return nil, fmt.Errorf("d2d: 加载 d2d1.dll 的 D2D1CreateFactory 失败: %w", err)
	}
	var factory *ID2D1Factory
	r, _, _ := procD2D1CreateFactory.Call(
		uintptr(factoryType),
		uintptr(com.NewGUIDPointer(IID_ID2D1Factory)),
		0, // D2D1_FACTORY_OPTIONS 传 NULL = 默认(不开调试层)
		uintptr(unsafe.Pointer(&factory)),
	)
	if r != 0 {
		return nil, errno(r)
	}
	if factory == nil {
		return nil, errors.New("d2d: D2D1CreateFactory 成功但没有返回工厂")
	}
	return factory, nil
}

// Package d2d 为"开启 D2D 渲染"时炫彩交出来的 Direct2D 对象提供封装.
//
// xc.XC_GetD2dFactory / xc.XDraw_GetD2dRenderTarget / xc.XDraw_GetD2dBitmap 只返回裸接口指针
// (uintptr), Go 侧既没有类型化接口, 也没有建路径几何 / 填充几何 / 建画刷的能力. 本包补上这一层:
// 借用宿主持有的 ID2D1Factory / ID2D1RenderTarget, 创建并填充 ID2D1PathGeometry.
//
// 纯 Go 实现, **不需要 cgo 也不需要汇编** —— Direct2D 的平滑(抗锯齿)对直通调用同样生效.
//
// 封装方式与 xcgui 既有的 edge 包(WebView2)一致: 每个 COM 接口一个文件, 虚表结构体按 Windows SDK
// 声明顺序**完整列出**各方法(索引由字段顺序决定, 不写魔数), 调用一律走 ComProc.Call(见 comproc.go
// 的 uintptrescapes 说明).
//
// # 继承关系与复用
//
// 接口结构体按 COM 的继承链**值嵌入**排布:
//
//	IUnknown{ Vtbl }  →  ID2D1Resource  →  { ID2D1RenderTarget, ID2D1Geometry → ID2D1PathGeometry,
//	                                        ID2D1Brush → ID2D1SolidColorBrush }
//	IUnknown{ Vtbl }  →  ID2D1SimplifiedGeometrySink  →  ID2D1GeometrySink
//	IUnknown{ Vtbl }  →  ID2D1Factory
//
// 好处是方法真的被继承: AddRef/Release/QueryInterface 只在 IUnknown 上实现一次, GetFactory 只在
// ID2D1Resource 上实现一次; 各接口自己的方法用未导出的 `vtbl()` 把**同一个对象指针**重新解释成自己
// 的虚表类型(同一个对象、同一个虚表指针, 只是槽位更多), 例如
// `(*ID2D1RenderTargetVtbl)(unsafe.Pointer(i)).FillGeometry`.
//
// 代价是这些结构体**不能加字段**(加了就不再等于一个虚表指针, 布局与真实 COM 对象错位) —— d2d.go
// 里有编译期断言守着这一条. 需要基类指针时直接取嵌入字段地址, 例如 `&brush.ID2D1Brush`.
//
// nil 语义: 构造函数在句柄为 0(未开启 D2D)时返回 nil, **调用方必须先判空**; nil 上只有 Release
// 与各接口自有方法是安全的(值嵌入下"被提升的基类方法"在 nil 外层指针上会先解引用嵌入字段而 panic,
// 这是 Go 的语义, 基类里判空也救不了).
//
// # ABI 三条硬规则(违反任意一条: 轻则整帧被毁/窗口隐形, 重则进程崩溃)
//
//  1. 形参里出现 float32/float64 的方法在 **amd64** 上不能用 syscall 直接调 —— 浮点参数走 XMM
//     寄存器, 而 syscall.SyscallN 只能设置通用寄存器(位打包也不行, 实测 SetDpi(float,float)
//     直接崩). 要调这类方法请用 wapi/com 的汇编转发器(CallWithFloat32/CallWithFloat64/
//     CallWithFloat32RetFloat) —— 注意浮点槽号可能不是 XMM0, 见该包注释的"汇编转发器"一节.
//     **386** 下则按位拆成 32 位槽压栈即可. 本包自己只用"形参全是整数/指针"的方法:
//     CreatePathGeometry / Open(只有出参)、CreateSolidColorBrush(颜色按**指针**传)、
//     FillGeometry(全是指针/整数)、SimplifiedGeometrySink 的 SetFillMode/BeginFigure/AddLines/
//     EndFigure. 虚表里其余方法照列(为了索引正确), 但标注了不能这样调用的原因 —— 例如
//     ID2D1RenderTarget::DrawGeometry 的线宽是 float 形参(想要带描边的几何, 只能在 CPU 侧把描边
//     展开成填充多边形再 FillGeometry), ID2D1Brush::SetTransform 的矩阵含 float32.
//  2. 按值传递的 8 字节结构体要打包成一个 64 位整数: D2D1_POINT_2F 用 pointArg 打包.
//  3. 其余参数一律用指针/整数传是安全的.
//
// # 坐标: 直通调用拿不到"当前 UI 对象"的变换与裁剪
//
// 炫彩是在**刷新它自己的绘制命令**(XDraw_*)时才施加元素偏移与裁剪的; 直通 D2D 调用发生在绘制
// 回调里、命令队列之前, 因此拿到的是**父级(窗口)坐标系**. 直接按客户区坐标绘制会跑到窗口左上角
// (典型症状: 幽灵态按钮的图标"消失", 其实画到标题栏底下去了).
//
// 正解: 先用 xc.XDraw_ConvRect(hDraw, *xc.RECT) 把"当前 UI 对象的客户区坐标"换算成"画布对齐
// 坐标"(自动含元素偏移 + DPI 缩放), 换算结果就是**物理像素**, 不要再乘 DPI.
//
// # 生命周期: 谁持有谁释放
//
//   - xcgui 持有的对象(ID2D1Factory / ID2D1RenderTarget)用 NewFactoryByHandle /
//     NewRenderTargetByHandle **借用**: 调用方不要 Release(设备重建时对象由 xcgui 销毁),
//     并且只在绘制回调内使用.
//   - 本包创建的对象(ID2D1PathGeometry / ID2D1SolidColorBrush)用完必须 Release.
//     几何是**设备无关**资源(归属 ID2D1Factory), 可以长期缓存复用, 但工厂换了(设备重建)要整体
//     释放重建; 画刷是**设备相关**资源(只能由创建它的渲染目标使用), 建议即建即放.
//
// # 示例(在 AddEvent_Paint 的绘制回调内)
//
//	hDraw := paint.Draw.Handle
//	rt := d2d.NewRenderTargetByHandle(xc.XDraw_GetD2dRenderTarget(hDraw))
//	if rt == nil {
//		// GDI 渲染模式(XC_IsEnableD2D() == false), 走回落路径.
//		return
//	}
//	rc := *clientRect
//	xc.XDraw_ConvRect(hDraw, &rc) // 客户区坐标 → 画布对齐坐标(物理像素)
//
//	geometry, err := d2d.NewFactoryByHandle(xc.XC_GetD2dFactory()).NewGeometry([][]float32{
//		{float32(rc.Left), float32(rc.Top), float32(rc.Right), float32(rc.Top),
//			float32(rc.Right), float32(rc.Bottom)}, // 一条闭合折线(至少 3 个点)
//	})
//	if err != nil {
//		return
//	}
//	defer geometry.Release() // 长期缓存时不要 defer, 见"生命周期"
//
//	rt.FillGeometryColor(geometry, d2d.NewColorF(0, 120, 215, 255))
//
// 相关文档:
//   - ID2D1RenderTarget: https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/nn-d2d1-id2d1rendertarget
//   - ID2D1Factory: https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/nn-d2d1-id2d1factory
//   - ID2D1PathGeometry: https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/nn-d2d1-id2d1pathgeometry
//   - ID2D1SimplifiedGeometrySink: https://learn.microsoft.com/zh-cn/windows/win32/api/d2d1/nn-d2d1-id2d1simplifiedgeometrysink
package d2d

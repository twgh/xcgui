package xc

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/wapi"
)

// getAnimaItemCallbackPtr 获取动画项回调函数指针
func getAnimaItemCallbackPtr() uintptr {
	animaItemCallbackOnce.Do(func() {
		animaItemCallbackPtr = createFloatCallback(syscall.NewCallback(animaItemCallbackShell))
	})
	return animaItemCallbackPtr
}

// 汇编代码（机器码）.
// 功能：将 XMM1 (float) 移动到 RDX，然后跳转到 Go 函数.
// 逻辑：
//
//	MOVD RDX, XMM1  (把 XMM1 里的浮点数位模式移到 RDX)
//	MOV R10, [目标地址]
//	JMP R10
func createFloatCallback(targetFn uintptr) uintptr {
	// 机器码模板
	// 48 0F 7E CA  -> MOVD RDX, XMM1
	// 49 BA [8字节地址] -> MOV R10, imm64
	// 41 FF E2     -> JMP R10

	code := []byte{
		// MOVD RDX, XMM1 (前缀)
		0x66, 0x48, 0x0F, 0x7E, 0xCA,
		// MOV R10, imm64 (将目标函数地址放入 R10)
		0x49, 0xBA,
		// 8字节地址
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		// JMP R10
		0x41, 0xFF, 0xE2,
	}

	// 填入 Go 函数地址（小端序）
	for i := 0; i < 8; i++ {
		code[7+i] = byte(targetFn >> (i * 8))
	}

	// 分配可执行内存
	mem := wapi.VirtualAlloc(0, uintptr(len(code)), wapi.MEM_COMMIT|wapi.MEM_RESERVE, wapi.PAGE_EXECUTE_READWRITE)
	if mem == 0 {
		panic("VirtualAlloc failed")
	}

	// 将机器码写入内存
	// 1. 将 mem (uintptr) 转换为 unsafe.Pointer
	// 2. 再转换为 *byte 指针
	// 3. 使用 unsafe.Slice 创建一个长度为 len(code) 的字节切片
	// 4. 将机器码拷贝到内存
	executableMemory := unsafe.Slice((*byte)(unsafe.Pointer(mem)), len(code))
	copy(executableMemory, code)
	return mem
}

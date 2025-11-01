package wapi

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
)

var (
	// Winmm 是 Windows 系统中负责音频播放、录音、MIDI 音乐控制等多媒体功能的核心组件。
	Winmm = syscall.NewLazyDLL("winmm.dll")

	// Functions.
	procPlaySoundW         = Winmm.NewProc("PlaySoundW")
	procMciSendStringW     = Winmm.NewProc("mciSendStringW")
	procMciGetErrorStringW = Winmm.NewProc("mciGetErrorStringW")
)

// PlaySound 播放由给定文件名、资源或系统事件指定的声音。
//   - 主要是用来播放 wav 格式的音频文件, 不要用来播放 mp3。
//   - 系统事件可能与注册表或 WIN.INI 文件中的声音相关联。
//   - 会在以下目录中搜索声音文件：当前目录；Windows 目录；Windows 系统目录；PATH 环境变量中列出的目录；以及网络中映射的目录列表。如果函数找不到指定的声音且未指定 SND_NODEFAULT 标志，则会使用默认的系统事件声音代替。如果函数既找不到系统默认条目也找不到默认声音，则不发出声音并返回 FALSE。
//
// pszSound: 指定要播放的声音的字符串。包括空终止符在内的最大长度为 256 个字符。
//   - 如果此参数为空，则停止当前正在播放的任何波形声音。
//   - 要停止非波形声音，请在 fdwSound 参数中指定 SND_PURGE。
//   - fdwSound 中的三个标志（ SND_ALIAS, SND_FILENAME, SND_RESOURCE ）决定了名称是被解释为系统事件的别名、文件名还是资源标识符。如果未指定这些标志中的任何一个，则会在注册表或 WIN.INI 文件中搜索与指定声音名称的关联。如果找到关联，则播放声音事件。如果在注册表中未找到关联，则该名称将被解释为文件名。
//
// hmod: 包含要加载资源的可执行文件的句柄。除非在 fdwSound 中指定了 SND_RESOURCE, 否则此参数必须为0。
//
// fdwSound: 播放声音的标志组合, wapi.SND_ .
//
// https://learn.microsoft.com/zh-cn/previous-versions//dd743680(v=vs.85)
func PlaySound(pszSound string, hmod uintptr, fdwSound SND_) bool {
	ret, _, _ := procPlaySoundW.Call(
		common.StrPtr(pszSound),
		hmod,
		uintptr(fdwSound),
	)
	return ret != 0
}

// SND_ 播放声音的标志位.
type SND_ uint32

const (
	SND_SYNC        SND_ = 0x0000     // 同步播放声音，函数在声音播放完成后才返回。这是默认行为
	SND_ASYNC       SND_ = 0x0001     // 异步播放声音，函数在开始播放后立即返回。要终止异步播放的声音，可以使用 NULL 调用 PlaySound
	SND_NODEFAULT   SND_ = 0x0002     // 找不到声音时不播放默认声音，保持静默
	SND_MEMORY      SND_ = 0x0004     // 指针指向内存中的声音数据
	SND_LOOP        SND_ = 0x0008     // 循环播放声音直到再次调用 PlaySound 且 pszSound 参数为 NULL。设置此标志时必须同时设置 SND_ASYNC 标志
	SND_NOSTOP      SND_ = 0x0010     // 指定的声音事件会让步于同一进程中已经在播放的另一个声音事件。如果因为生成声音所需的资源正忙于播放其他声音而导致声音无法播放，函数会立即返回 FALSE 而不播放请求的声音
	SND_NOWAIT      SND_ = 0x00002000 // 不支持（注意：早期文档错误地标记为此标志受支持，但函数会忽略此标志）
	SND_ALIAS       SND_ = 0x00010000 // pszSound 参数是注册表或 WIN.INI 文件中的系统事件别名。不能与 SND_FILENAME 或 SND_RESOURCE 一起使用
	SND_ALIAS_ID    SND_ = 0x00110000 // pszSound 参数是系统事件别名的预定义标识符
	SND_FILENAME    SND_ = 0x00020000 // pszSound 参数是一个文件名。如果找不到文件，则播放默认声音，除非设置了 SND_NODEFAULT 标志
	SND_RESOURCE    SND_ = 0x00040004 // pszSound 参数是一个资源标识符；hmod 必须标识包含该资源的实例
	SND_PURGE       SND_ = 0x0040     // 要停止非波形声音，请在 fdwSound 参数中指定。
	SND_APPLICATION SND_ = 0x0080     // pszSound 参数是注册表中的应用程序特定别名。您可以将此标志与 SND_ALIAS 或 SND_ALIAS_ID 标志结合使用，以指定应用程序定义的声音别名
	SND_SENTRY      SND_ = 0x00080000 // 触发 SoundSentry 事件（需要 Windows Vista 或更高版本）。 SoundSentry 是一项辅助功能，当播放声音时会导致计算机显示视觉提示
	SND_RING        SND_ = 0x00100000 // 将此视为来自通信应用程序的"铃声" - 不降低音量
	SND_SYSTEM      SND_ = 0x00200000 // 分配给系统通知声音的音频会话（需要 Windows Vista 或更高版本）。系统音量控制程序(SndVol)会显示一个控制系统通知声音的音量滑块
)

// 预定义的系统声音别名

const (
	SND_ALIAS_SYSTEMASTERISK    SND_ = ('*' << 8) | 'S'
	SND_ALIAS_SYSTEMQUESTION    SND_ = ('?' << 8) | 'S'
	SND_ALIAS_SYSTEMHAND        SND_ = ('H' << 8) | 'S'
	SND_ALIAS_SYSTEMEXIT        SND_ = ('E' << 8) | 'S'
	SND_ALIAS_SYSTEMSTART       SND_ = ('S' << 8) | 'S'
	SND_ALIAS_SYSTEMWELCOME     SND_ = ('W' << 8) | 'S'
	SND_ALIAS_SYSTEMEXCLAMATION SND_ = ('!' << 8) | 'S'
	SND_ALIAS_SYSTEMDEFAULT     SND_ = ('D' << 8) | 'S'
)

// MciSendString 向 MCI 设备发送命令字符串。命令发送的目标设备在该命令字符串中指定。
//
// command: 指定 MCI 命令字符串。有关列表，请参见多媒体命令字符串: https://learn.microsoft.com/zh-cn/windows/win32/multimedia/multimedia-command-strings?redirectedfrom=MSDN。
//
// returnString: 指向接收返回信息的缓冲区的指针。如果不需要返回信息，则此参数可以为 nil。
//
// cchReturn: 缓冲区的大小（以字符为单位）。
//
// hwndCallback: 如果在命令字符串中指定了"notify"标志，则为回调窗口的句柄。
//
// 返回值:
//   - 成功时返回0，否则返回错误。返回值的低字包含错误返回值。
//   - 如果错误是设备特定的，则返回值的高字是驱动程序标识符；否则高字为零。
//   - 有关可能的错误值列表，请参见 MCIERR 返回值: https://learn.microsoft.com/zh-cn/windows/win32/multimedia/mcierr-return-values?redirectedfrom=MSDN
//   - 要获取返回值的文本描述，请将返回值传递给 MciGetErrorString 函数。
//
// https://learn.microsoft.com/zh-cn/previous-versions//dd757161(v=vs.85)
func MciSendString(command string, returnString *string, cchReturn uint32, hwndCallback uintptr) uint32 {
	var retStrPtr *uint16
	var retStr []uint16

	// 如果提供了返回字符串缓冲区
	if returnString != nil && cchReturn > 0 {
		retStr = make([]uint16, cchReturn)
		retStrPtr = &retStr[0]
	}

	ret, _, _ := procMciSendStringW.Call(
		common.StrPtr(command),
		uintptr(unsafe.Pointer(retStrPtr)),
		uintptr(cchReturn),
		hwndCallback,
	)

	// 如果提供了返回字符串缓冲区，将其转换为Go字符串
	if returnString != nil && ret == 0 {
		*returnString = syscall.UTF16ToString(retStr)
	}
	return uint32(ret)
}

// MciGetErrorString 获取 MCI 错误代码的文本描述.
//   - 如果失败或未知错误代码，则返回空字符串.
//
// fdwError: MciSendCommand 或 MciSendString 函数返回的错误代码。
//
// https://learn.microsoft.com/zh-cn/previous-versions//dd757158(v=vs.85)
func MciGetErrorString(fdwError uint32) string {
	cchErrorText := uint32(128)
	retStr := make([]uint16, cchErrorText)
	retStrPtr := &retStr[0]

	ret, _, _ := procMciGetErrorStringW.Call(
		uintptr(fdwError),
		uintptr(unsafe.Pointer(retStrPtr)),
		uintptr(cchErrorText),
	)

	if ret != 0 {
		return syscall.UTF16ToString(retStr)
	}
	return ""
}

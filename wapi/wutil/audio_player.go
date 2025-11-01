package wutil

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
)

// AudioPlayer 音频播放器.
type AudioPlayer struct {
	alias string // 音频别名
}

// NewAudioPlayer 创建新的音频播放器
func NewAudioPlayer() *AudioPlayer {
	return &AudioPlayer{}
}

// Open 打开音频文件.
//
// fileName: 文件路径.
func (ap *AudioPlayer) Open(fileName string) error {
	if ap.alias != "" {
		return nil
	}

	alias := "audio_" + generateID()
	cmd := fmt.Sprintf(`open "%s" alias %s`, fileName, alias)

	err := _MciSendString(cmd)
	if err != nil {
		return err
	}

	ap.alias = alias
	return nil
}

// Play 播放音频文件.
//   - 最后必须调用 Close() 以释放系统资源.
//
// sync: 是否同步播放, 为空默认为 false, 异步播放.
func (ap *AudioPlayer) Play(sync ...bool) error {
	if ap.alias == "" {
		return fmt.Errorf("must open the audio first")
	}

	// 先停止当前播放
	_ = ap.Stop()

	// 确定播放方式
	isSync := false
	if len(sync) > 0 {
		isSync = sync[0]
	}

	wait := ""
	if isSync {
		wait = " wait"
	}

	// 开始播放
	cmd := fmt.Sprintf("play %s%s", ap.alias, wait)
	return _MciSendString(cmd)
}

// Pause 暂停播放.
//   - 暂停后可以调用 Resume() 恢复播放.
func (ap *AudioPlayer) Pause() error {
	if ap.alias == "" {
		return nil
	}

	cmd := fmt.Sprintf("pause %s", ap.alias)
	return _MciSendString(cmd)
}

// Resume 恢复播放.
func (ap *AudioPlayer) Resume() error {
	if ap.alias == "" {
		return fmt.Errorf("must open the audio first")
	}

	cmd := fmt.Sprintf("resume %s", ap.alias)
	return _MciSendString(cmd)
}

// Stop 停止播放.
func (ap *AudioPlayer) Stop() error {
	if ap.alias == "" {
		return nil
	}

	cmd := fmt.Sprintf("stop %s", ap.alias)
	return _MciSendString(cmd)
}

// Close 关闭音频设备.
//   - 关闭前最好先 Stop().
func (ap *AudioPlayer) Close() error {
	if ap.alias == "" {
		return nil
	}

	cmd := fmt.Sprintf("close %s", ap.alias)
	err := _MciSendString(cmd)
	if err != nil {
		return err
	}

	ap.alias = ""
	return nil
}

// SetVolume 设置音量 (0-100).
//
// volume: 音量范围, 0-100.
func (ap *AudioPlayer) SetVolume(volume int) error {
	if ap.alias == "" {
		return fmt.Errorf("must open the audio first")
	}

	if volume < 0 || volume > 100 {
		return fmt.Errorf("the volume must be between 0 and 100")
	}

	// 先尝试 0-100 范围
	cmd := fmt.Sprintf("setaudio %s volume to %d", ap.alias, volume)
	err := _MciSendString(cmd)
	if err == nil {
		return nil
	}

	// 如果失败，尝试 0-1000 范围
	volume *= 10
	if volume > 1000 {
		volume = 1000
	}

	cmd = fmt.Sprintf("setaudio %s volume to %d", ap.alias, volume)
	return _MciSendString(cmd)
}

// GetVolume 获取音量 (0-100).
func (ap *AudioPlayer) GetVolume() (int, error) {
	if ap.alias == "" {
		return 0, fmt.Errorf("must open the audio first")
	}

	// 查询音量
	volume, err := ap.getStatus("volume")
	if err != nil {
		return 0, err
	}

	// MCI 音量范围可能是 0-1000，转换为 0-100
	if volume > 100 {
		volume = volume / 10
	}
	return volume, nil
}

// Seek 跳转到指定位置.
//
// positionMs: 跳转位置 (毫秒).
func (ap *AudioPlayer) Seek(positionMs int) error {
	// 获取音频总长度
	length, err := ap.GetLength()
	if err != nil {
		return fmt.Errorf("failed to get the length of the audio: %v", err)
	}

	// 确保位置在有效范围内
	if positionMs < 0 {
		positionMs = 0
	}
	if positionMs > length {
		positionMs = length
	}

	// 保存当前状态
	originalMode, _ := ap.GetPlaybackMode()

	// 执行跳转
	cmd := fmt.Sprintf("seek %s to %d", ap.alias, positionMs)
	if err := _MciSendString(cmd); err != nil {
		return err
	}

	// 根据原始状态恢复播放
	switch originalMode {
	case "playing":
		// 跳转后自动继续播放
		if err := ap.Play(); err != nil {
			return fmt.Errorf("failed to resume playback after jumping: %v", err)
		}
	case "paused":
		// 保持在暂停状态，但位置已改变
		// 不需要额外操作
	default:
		// 停止或其他状态，保持现状
	}

	return nil
}

// SeekToStart 跳转到开头.
func (ap *AudioPlayer) SeekToStart() error {
	return ap.Seek(0)
}

// SeekToEnd 跳转到结尾.
func (ap *AudioPlayer) SeekToEnd() error {
	length, err := ap.GetLength()
	if err != nil {
		return err
	}
	return ap.Seek(length)
}

// SeekRelative 相对当前位置跳转.
//
// offsetMs: 相对位置 (毫秒), 正数向前，负数向后.
func (ap *AudioPlayer) SeekRelative(offsetMs int) error {
	// 获取当前位置
	currentPos, err := ap.GetPosition()
	if err != nil {
		return fmt.Errorf("failed to get current location: %v", err)
	}

	// 计算新位置
	newPos := currentPos + offsetMs

	// 确保位置有效
	if newPos < 0 {
		newPos = 0
	}

	// 获取总长度
	length, err := ap.GetLength()
	if err == nil && newPos > length {
		newPos = length
	}

	// 执行跳转
	return ap.Seek(newPos)
}

// SeekPercent 按百分比跳转.
//
// percent: 百分比, 0.0 - 1.0.
func (ap *AudioPlayer) SeekPercent(percent float64) error {
	if percent < 0 {
		percent = 0
	}
	if percent > 1 {
		percent = 1
	}

	// 获取总长度
	length, err := ap.GetLength()
	if err != nil {
		return fmt.Errorf("failed to get the length of the audio: %v", err)
	}

	// 计算位置
	position := int(float64(length) * percent)

	// 执行跳转
	return ap.Seek(position)
}

// GetPosition 获取当前播放位置 (毫秒).
func (ap *AudioPlayer) GetPosition() (int, error) {
	if ap.alias == "" {
		return 0, fmt.Errorf("must open the audio first")
	}
	return ap.getStatus("position")
}

// GetLength 获取音频长度 (毫秒).
func (ap *AudioPlayer) GetLength() (int, error) {
	if ap.alias == "" {
		return 0, fmt.Errorf("must open the audio first")
	}
	return ap.getStatus("length")
}

// GetPlaybackMode 获取播放模式.
//
// 返回值: 播放模式字符串, 对应 Mode 开头的常量, 如: ModePlaying.
func (ap *AudioPlayer) GetPlaybackMode() (string, error) {
	if ap.alias == "" {
		return "", fmt.Errorf("must open the audio first")
	}

	cmd := fmt.Sprintf("status %s mode", ap.alias)

	var resultStr string
	ret := wapi.MciSendString(cmd, &resultStr, 256, 0)
	if ret != 0 {
		errText := wapi.MciGetErrorString(uint32(common.GetLowWord(ret)))
		return "unknown", fmt.Errorf("failed to get playback mode: %s", errText)
	}

	return strings.TrimSpace(resultStr), nil
}

// 常见的 MCI 播放模式

const (
	ModePlaying = "playing" // 正在播放
	ModePaused  = "paused"  // 已暂停
	ModeStopped = "stopped" // 已停止
	ModeOpen    = "open"    // 设备已打开但未播放
	ModeSeeking = "seeking" // 正在寻址
	ModeReady   = "ready"   // 设备就绪
	ModeUnknown = "unknown" // 获取播放模式失败
	ModeEmpty   = ""        // 设备未打开, 就是没有调用 Open()
)

// IsPlaying 是否正在播放.
func (ap *AudioPlayer) IsPlaying() bool {
	mode, err := ap.GetPlaybackMode()
	if err != nil {
		return false
	}
	return mode == ModePlaying
}

// IsPaused 是否已暂停.
func (ap *AudioPlayer) IsPaused() bool {
	mode, err := ap.GetPlaybackMode()
	if err != nil {
		return false
	}
	return mode == ModePaused
}

// IsStopped 是否已停止.
func (ap *AudioPlayer) IsStopped() bool {
	mode, err := ap.GetPlaybackMode()
	if err != nil {
		return true // 如果获取失败，认为已停止
	}
	return mode == ModeStopped || mode == ModeOpen || mode == ModeReady
}

// CanPlay 是否可以开始播放.
func (ap *AudioPlayer) CanPlay() bool {
	mode, err := ap.GetPlaybackMode()
	if err != nil {
		return false
	}

	// 在停止、暂停、就绪状态下可以播放
	return mode == ModeStopped || mode == ModePaused || mode == ModeOpen || mode == ModeReady
}

// WaitForPlaybackComplete 等待播放完成.
//
// timeout: 超时时间.
func (ap *AudioPlayer) WaitForPlaybackComplete(timeout time.Duration) error {
	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {
		mode, err := ap.GetPlaybackMode()
		if err != nil {
			return err
		}

		if mode == ModeStopped || mode == ModeOpen {
			return nil // 播放完成
		}

		if !ap.IsPlaying() && !ap.IsPaused() {
			return fmt.Errorf("playback terminated abnormally")
		}

		time.Sleep(100 * time.Millisecond) // 每 100 ms 检查一次
	}

	return fmt.Errorf("waiting for playback to complete timeout")
}

// getStatus 通用的状态查询方法.
func (ap *AudioPlayer) getStatus(statusType string) (int, error) {
	cmd := fmt.Sprintf("status %s %s", ap.alias, statusType)

	var resultStr string
	ret := wapi.MciSendString(cmd, &resultStr, 256, 0)
	if ret != 0 {
		errText := wapi.MciGetErrorString(uint32(common.GetLowWord(ret)))
		return 0, fmt.Errorf("query status failed, type: %s, reason: %s", statusType, errText)
	}

	var value int
	_, err := fmt.Sscanf(resultStr, "%d", &value)
	if err != nil {
		return 0, fmt.Errorf("parse status value failed: %v", err)
	}

	return value, nil
}

// 生成随机 ID
func generateID() string {
	bytes := make([]byte, 4)
	_, _ = rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

// 调用 MCI 命令
func _MciSendString(cmd string) error {
	ret := wapi.MciSendString(cmd, nil, 0, 0)
	if ret != 0 {
		errText := wapi.MciGetErrorString(uint32(common.GetLowWord(ret)))

		// 取命令名, 从 cmd 开头截取到第一个空格
		name := cmd[:strings.Index(cmd, " ")]
		return errors.New(name + " failed: " + errText)
	}
	return nil
}

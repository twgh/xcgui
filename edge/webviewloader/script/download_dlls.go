// 下载最新正式版本的 webviewloader.dll 和头文件
package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// 主函数
func main() {
	// 获取最新版本
	version, err := getLatestVersion()
	if err != nil {
		fmt.Printf("错误: %v\n", err)
		return
	}
	fmt.Println("最新版本: " + version)

	// 创建临时目录
	tmpDir, err := os.MkdirTemp("", "webview2_")
	if err != nil {
		fmt.Printf("创建临时目录失败: %v\n", err)
		return
	}
	defer os.RemoveAll(tmpDir)

	// 下载 NuGet 包
	url := fmt.Sprintf("https://api.nuget.org/v3-flatcontainer/microsoft.web.webview2/%s/microsoft.web.webview2.%s.nupkg",
		version, version)
	pkgPath := filepath.Join(tmpDir, "webview2.nupkg")

	fmt.Println("开始下载 NuGet 包")
	if err := downloadFile(url, pkgPath); err != nil {
		fmt.Printf("下载失败: %v\n", err)
		return
	}

	// 解压并提取文件
	fmt.Println("开始解压文件")
	if err := extractDlls(pkgPath); err != nil {
		fmt.Printf("提取失败: %v\n", err)
		return
	}

	fmt.Println("操作成功完成")
}

// 获取最新版本号
func getLatestVersion() (string, error) {
	resp, err := http.Get("https://api.nuget.org/v3-flatcontainer/microsoft.web.webview2/index.json")
	if err != nil {
		return "", fmt.Errorf("获取版本失败: %v", err)
	}
	defer resp.Body.Close()

	var data struct {
		Versions []string `json:"versions"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", fmt.Errorf("解析版本失败: %v", err)
	}

	var stableVersions []string
	for _, v := range data.Versions {
		// 排除包含预发布标识的版本
		if !strings.Contains(v, "-") {
			stableVersions = append(stableVersions, v)
		}
	}

	if len(stableVersions) == 0 {
		return "", fmt.Errorf("未找到稳定版本")
	}

	// 按语义化版本排序
	sort.Slice(stableVersions, func(i, j int) bool {
		v1 := strings.Split(stableVersions[i], ".")
		v2 := strings.Split(stableVersions[j], ".")
		return compareVersion(v1, v2) > 0
	})

	return stableVersions[0], nil
}

// 比较版本号（返回 1 表示 v1 > v2）
func compareVersion(v1, v2 []string) int {
	maxLen := len(v1)
	if len(v2) > maxLen {
		maxLen = len(v2)
	}

	for i := 0; i < maxLen; i++ {
		var n1, n2 int
		if i < len(v1) {
			n1 = parseInt(v1[i])
		}
		if i < len(v2) {
			n2 = parseInt(v2[i])
		}
		if n1 != n2 {
			return n1 - n2
		}
	}
	return 0
}

func parseInt(s string) int {
	var n int
	fmt.Sscanf(s, "%d", &n)
	return n
}

// 下载文件
func downloadFile(url, dest string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP 状态码 %d", resp.StatusCode)
	}

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

// 解压并提取 DLL 和头文件
func extractDlls(pkgPath string) error {
	// 创建目标目录
	baseDir, _ := os.Getwd()
	sdkDir := filepath.Join(baseDir, "edge/webviewloader/sdk")
	dirs := map[string]string{
		"win-x86": filepath.Join(sdkDir, "x86"),
		"win-x64": filepath.Join(sdkDir, "x64"),
	}

	// 创建目录
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}

	// 打开 ZIP 文件
	r, err := zip.OpenReader(pkgPath)
	if err != nil {
		return err
	}
	defer r.Close()

	headerFiles := []string{
		"build/native/include/WebView2EnvironmentOptions.h",
		"build/native/include/WebView2.h",
	}

	// 遍历文件
	for _, f := range r.File {
		if strings.HasPrefix(f.Name, "runtimes/") {
			// 匹配目标文件路径
			for arch, destDir := range dirs {
				targetPath := fmt.Sprintf("runtimes/%s/native/WebView2Loader.dll", arch)
				if f.Name == targetPath {
					if err := extractFile(f, destDir); err != nil {
						return err
					}
				}
			}
		} else if strings.HasPrefix(f.Name, "build/native/include/") {
			//	提取头文件到 sdkDir 目录
			for _, headerFile := range headerFiles {
				if f.Name == headerFile {
					if err := extractFile(f, sdkDir); err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

// 提取单个文件
func extractFile(f *zip.File, destDir string) error {
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	destPath := filepath.Join(destDir, filepath.Base(f.Name))
	out, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, rc)
	return err
}

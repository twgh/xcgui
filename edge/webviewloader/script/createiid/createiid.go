// 读取 WebView2.h 文件生成 IID 常量 Go 代码文件
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	inputFile  = "edge/webviewloader/sdk/WebView2.h"
	outputFile = "edge/webview2_iids.go"
)

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	defer file.Close()

	// 匹配IID声明行
	iidRegex := regexp.MustCompile(`EXTERN_C __declspec\(selectany\) const IID (IID_\w+)`)
	// 匹配完整的GUID格式（包含所有11个组件）
	guidRegex := regexp.MustCompile(`\{0x([0-9a-fA-F]{8}),0x([0-9a-fA-F]{4}),0x([0-9a-fA-F]{4}),\{0x([0-9a-fA-F]{2}),0x([0-9a-fA-F]{2}),0x([0-9a-fA-F]{2}),0x([0-9a-fA-F]{2}),0x([0-9a-fA-F]{2}),0x([0-9a-fA-F]{2}),0x([0-9a-fA-F]{2}),0x([0-9a-fA-F]{2})\}\}`)

	constants := make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if iidMatch := iidRegex.FindStringSubmatch(line); iidMatch != nil {
			iidName := iidMatch[1]

			if guidMatch := guidRegex.FindStringSubmatch(line); guidMatch != nil {
				// 确保有足够的匹配组
				if len(guidMatch) >= 12 {
					// 构建完整的GUID字符串，包含所有11个组件
					guid := fmt.Sprintf("%s-%s-%s-%s%s-%s%s%s%s%s%s",
						strings.ToLower(guidMatch[1]),  // 8位
						strings.ToLower(guidMatch[2]),  // 4位
						strings.ToLower(guidMatch[3]),  // 4位
						strings.ToLower(guidMatch[4]),  // 2位
						strings.ToLower(guidMatch[5]),  // 2位
						strings.ToLower(guidMatch[6]),  // 2位
						strings.ToLower(guidMatch[7]),  // 2位
						strings.ToLower(guidMatch[8]),  // 2位
						strings.ToLower(guidMatch[9]),  // 2位
						strings.ToLower(guidMatch[10]), // 2位
						strings.ToLower(guidMatch[11])) // 2位

					constants[iidName] = guid
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件时出错:", err)
		return
	}

	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("创建输出文件时出错:", err)
		return
	}
	defer outFile.Close()

	outFile.WriteString("package edge\n\n")
	outFile.WriteString("const (\n")
	for name, value := range constants {
		outFile.WriteString(fmt.Sprintf("\t%s = \"%s\"\n", name, value))
	}
	outFile.WriteString(")\n")

	fmt.Printf("成功生成 %s, IID 数量: %s\n", outputFile, len(constants))
}

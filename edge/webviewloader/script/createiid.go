// 读取 WebView2.h 文件生成 IID 常量 Go 代码文件
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	// 读取 WebView2.h 文件
	file, err := os.Open("edge/webviewloader/sdk/WebView2.h")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 正则表达式匹配IID声明行
	iidRegex := regexp.MustCompile(`EXTERN_C __declspec\(selectany\) const IID (IID_\w+)`)
	guidRegex := regexp.MustCompile(`\{0x([0-9a-fA-F]{8}),0x([0-9a-fA-F]{4}),0x([0-9a-fA-F]{4}),\{0x([0-9a-fA-F]{2}),0x([0-9a-fA-F]{2}),0x([0-9a-fA-F]{2}),0x([0-9a-fA-F]{2}),0x([0-9a-fA-F]{2}),0x([0-9a-fA-F]{2}),0x([0-9a-fA-F]{2}),0x([0-9a-fA-F]{2})\}\}`)

	// 存储提取的IID常量
	constants := make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// 检查是否包含IID声明
		if iidMatch := iidRegex.FindStringSubmatch(line); iidMatch != nil {
			iidName := iidMatch[1]

			// 提取GUID部分
			if guidMatch := guidRegex.FindStringSubmatch(line); guidMatch != nil {
				// 构建GUID字符串
				guid := fmt.Sprintf("%s-%s-%s-%s%s-%s%s%s%s",
					strings.ToLower(guidMatch[1]),
					strings.ToLower(guidMatch[2]),
					strings.ToLower(guidMatch[3]),
					strings.ToLower(guidMatch[4]),
					strings.ToLower(guidMatch[5]),
					strings.ToLower(guidMatch[6]),
					strings.ToLower(guidMatch[7]),
					strings.ToLower(guidMatch[8]),
					strings.ToLower(guidMatch[9]),
				)

				constants[iidName] = guid
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 生成 Go 常量文件
	outFile, err := os.Create("edge/webview2_iids.go")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outFile.Close()

	outFile.WriteString("package edge\n\n")
	outFile.WriteString("const (\n")
	for name, value := range constants {
		outFile.WriteString(fmt.Sprintf("\t%s = \"%s\"\n", name, value))
	}
	outFile.WriteString(")\n")

	fmt.Println("Successfully generated webview2_iids.go, count:", len(constants))
}

package main

import (
	"bufio"
	"fmt"
	"mirrorsTestTools/main/task"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 调用cfginput.ReadFile()函数读取配置文件，并将其转换为切片
	mirrorNames, mirrorURLs := task.ReadFile("urls.json")
	// fmt.Println(mirrorNames, mirrorURLs)

	fmt.Println("欢迎使用镜像测速工具")
	fmt.Println("1. 批量自定义镜像站URL. 2. 批量选择库中镜像站. 3. Exit")

	var choice int
	fmt.Scanln(&choice)

	switch choice {

	case 1:
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("请输入要测试的镜像站的 URL: (Default:Tsinghua)\n")

		// 读取用户批量输入的镜像站URL，并append到mirrorURLs切片中
		for {
			input, _ := reader.ReadString('\n')
			// 移除末尾的换行符
			input = strings.TrimSuffix(input, "\n")
			if input == "" {
				break
			}
			mirrorURLs = append(mirrorURLs, input)
		}

		// 执行测试
		task.Test(mirrorNames, mirrorURLs)

	case 2:
		fmt.Println("请选择要测试的镜像站(中间用空格隔开):")
		for i, name := range mirrorNames {
			fmt.Printf("%d. %s\n", i+1, name)
		}

		// 读取用户选择的镜像站序号
		var index string
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取输入错误:", err)
			return
		}

		// 移除换行符
		index = strings.TrimSuffix(line, "\n")

		// 以空格分割输入字符串
		data := strings.Fields(index) //data: [1 3 4]

		for _, d := range data {
			// 将字符串转换为整数
			i, err := strconv.Atoi(d)
			if err != nil {
				fmt.Println("无效的输入，请重新输入")
			}
			mirrorURLs = append(mirrorURLs, mirrorURLs[i-1])
		}

		// 执行测试
		task.Test(mirrorNames, mirrorURLs)

	case 3:
		return
	}

}

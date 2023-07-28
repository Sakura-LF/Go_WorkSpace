package main

import (
	"fmt"
)

const (
	Black  = 0
	Red    = 1
	Green  = 2
	Yellow = 3
	Blue   = 4
	Purple = 5
	Cyon   = 6
	Gray   = 7
)

func PrintColor(color int, text string, isBackGround bool) {
	if isBackGround {
		fmt.Printf("\033[3%dm%s\033[0m", color, text)
	} else {
		fmt.Printf("\033[4%dm%s\033[0m", color, text)
	}
}

func main() {
	PrintColor(Red, "红色", true)
	//前景色
	fmt.Println("\033[30m你好,这是黑色\033[0m")
	fmt.Println("\033[31m你好,这是红色\033[0m")
	fmt.Println("\033[32m你好,这是绿色\033[0m")
	fmt.Println("\033[33m你好,这是黄色\033[0m")
	fmt.Println("\033[34m你好,这是盗色\033[0m")
	fmt.Println("\033[35m你好,这是紫色\033[0m")
	fmt.Println("\033[36m你好,这是青色\033[0m")
	fmt.Println("\033[37m你好,这是灰色\033[0m")

	//背景色
	fmt.Println("\033[40m你好,这是黑色\033[0m")
	fmt.Println("\033[41m你好,这是红色\033[0m")
	fmt.Println("\033[42m你好,这是绿色\033[0m")
	fmt.Println("\033[43m你好,这是黄色\033[0m")
	fmt.Println("\033[44m你好,这是蓝色\033[0m")
	fmt.Println("\033[45m你好,这是紫色\033[0m")
	fmt.Println("\033[46m你好,这是青色\033[0m")
	fmt.Println("\033[47m你好,这是灰色\033[0m")
}

func Test() {

}

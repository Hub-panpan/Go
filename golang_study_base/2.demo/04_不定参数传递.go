package main

import "fmt"

func main() {
	Test(1, 2, 3, 4, 5)
}

func Test(args ...int) {
	// 全部元素传递给MyFunc 参数就变成了args...
	MyFunc(args...)
	fmt.Println("分隔线------ ")
	MyFunc01(args[:2]...) // args[0] ~ args[2] (不包括args[2]) 传递过去
	fmt.Println("分隔线------ ")
	MyFunc01(args[2:]...) // 从args[2]开始（包括本身），把后面所有元素传递过去
}

func MyFunc(temp ...int) {//函数接受的参数也是 不定参数
	for _, data := range temp {
		fmt.Println("data = ", data)
	}
}

func MyFunc01(temp ...int) {
	for _, data := range temp {
		fmt.Println("data = ", data)
	}
}

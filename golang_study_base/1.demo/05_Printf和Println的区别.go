package main

import "fmt"

func main() {
	a := 1
	// 自动换行 这是一段输出 一段一段处理
	fmt.Println("a = ", a)

	//格式化输出，把a的内容放在%d的位置 不能够自动换行 因此需要 自己 加一个换行符
	fmt.Printf("a = %d\n", a)

	// := 自动推导类型  声明变量以及赋值
	b := 20
	c := 30
	//这一种格式不容易控制格式
	fmt.Println("a = ", a, "b = ", b, "c = ", c)
	//如果有多个变量 的话 下面这个写起来比较方便
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)

}

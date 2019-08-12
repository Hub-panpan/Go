package main

import "fmt"

func main() {
	a, b := 10, 20
	swap(a, b) // 变量本身的传递，值传递
	fmt.Printf("main: a = %d, b = %d\n", a, b)
}

func swap(a, b int) (int, int) {
	a, b = b, a
	fmt.Printf("swap: a = %d, b = %d\n", a, b)
	return a, b
}
//分析结果 发现虽然swap函数当中的值改变了 但是main函数当中的值并没有改变

/**

很好理解 创建了副本 改变的是副本 

*/
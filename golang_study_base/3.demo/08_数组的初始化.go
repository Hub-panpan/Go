package main

import "fmt"

func main() {
	// 声明的同时赋值，叫初始化
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("a = ", a)

	//简便写法 非常的便捷 自动推导类型
	b := [5]int{11, 22, 33, 44, 55}
	fmt.Println("b = ", b)

	// 部分初始化 最后面两个元素的值会被初始化为0
	c := [5]int{1, 3, 5}
	fmt.Println("c = ", c)

	// 指定的元素初始化 下标为2与4的进行初始化操作
	d := [5]int{2: 222, 4: 6666}
	fmt.Println("d = ", d)
}

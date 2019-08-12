package main

import "fmt"

func main() {

	//在go语言中  字符串也是一种类型来的

	var str1 string // 声明变量
	str1 = "abc" //赋值
	fmt.Println("str1 = ", str1)

	// 自动推导类型
	str2 := "mike"
	//%T指的是类型
	fmt.Printf("str2 类型是 %T\n", str2)

	// 内建函数，len() 可以测字符串的长度，有多少个字符
	fmt.Println("len(str2) = ", len(str2))
}

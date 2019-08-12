package main

import "fmt"

func main() {
	//自动推到类型
	s := 1
	if s == 1 {//左括号 要和条件在同一行  记住是两个等于号
		fmt.Println("s等于1")
	}

	//支持一个初始化操作 只支持一个 ；分号隔离判断条件

	if a:=10; a==10{

		fmt.Println("a=10")
	}
}

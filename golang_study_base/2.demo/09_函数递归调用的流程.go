package main

import "fmt"

func main() {
	fmt.Println("main start")
	funca(5)
	fmt.Println("main end")
}

//换种方式进行演示 函数的调用
func funca(a int) {
	fmt.Println("a = ", a)
	//出口条件
	if a < 1 {
		return //函数终止 当函数a<1 将a的值返回
	}
	//函数调用自身 
	funca(a - 1)
}

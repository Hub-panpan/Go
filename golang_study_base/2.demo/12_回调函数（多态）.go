package main

import "fmt"

func main() {
	res := Calc(1, 1, Add) // 回调函数，函数有一个参数是函数类型，这个函数就是回调函数
	res2 := Calc(1, 1, Mul)
	fmt.Println("res = ", res)
	fmt.Println("res2 = ", res2)
}

type FuncType func(int, int) int

// 回调函数，函数有一个参数是函数类型，这个函数就是回调函数
// 计算器，实现四则运算
// 多态，多种形态，调用同一个接口，可以实现不同的表现，比如加减乘除

// 先定义方法，后实现功能
func Calc(a int, b int, fTest FuncType) (result int) {
	result = fTest(a, b)  //就是传入进来的都是计算 表现出来的加减乘除功能 你传什么我就运行啥功能 现有想法再实现功能
	return
}

func Add(a int, b int) int {
	return a - b
}

func Mul(a int, b int) int {
	return a * b
}


//相当与简化了操作
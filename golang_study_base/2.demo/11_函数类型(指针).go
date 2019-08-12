package main

import "fmt"

func main() {
	var result int
	result = Add(1, 2)
	fmt.Println("result = ", result) // 传统调用方式

	var fTest FuncType     // 声明一个函数类型的变量，变量名叫 fTest
	fTest = Add            // 是变量就可以赋值
	result = fTest(11, 22) // 等价于 Add(11, 22)
	fmt.Println("resultF = ", result)

	fTest = Minus
	result = fTest(22, 11) // 等价于 Minus(22, 11)
	fmt.Println("result = ", result)
}

// 函数也是一种数据类型，通过 type 给一个函数类型起名
// FuncType 它是一个函数类型
//下面的意思是 只的是 这种函数类型有两个int类型的参数 一个int类型的返回值
type FuncType func(int, int) int // 没有函数名字，没有 {}  ，类似指针

func Add(a int, b int) int {
	return a + b
}

func Minus(a int, b int) int {
	return a - b
}


//留个问题 为啥不用传统的那种方式呢 这就涉及到多态了 所谓的回调函数 
//函数有一个参数是函数类型  这个函数就是回调函数
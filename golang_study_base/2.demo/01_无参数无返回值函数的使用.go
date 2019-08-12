package main

import "fmt"

//main函数 就是主函数
func main() {//大括号 就是函数体 圆括号中放参数
	//函数的调用： 函数名（）不调用不会执行 因为main函数是程序的入口
	MyFunc()
}

// 无参无返回值函数的定义
func MyFunc() {
	a := 0111
	fmt.Println("a = ", a)
}


//函数放在前面 放在后面不影响程序的执行
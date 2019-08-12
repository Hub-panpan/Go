package main

import "fmt"

func main() {
	a := 10 //整型 10
	str := "hello"

	// 匿名函数，没有函数名字，函数定义，还没有调用
	// 以前是函数func xxx()现在把名字去了
	f1 := func() {
		//这个匿名函数可以捕获它外部的变量 现在这个方法在主函数里调用才可以
		fmt.Println("a = ", a)
		fmt.Println("str = ", str)
	}
	f1()

	// 给函数类型起别名
	type FuncType func() // 函数没有参数，没有返回值
	var f2 FuncType  //声明变量
	f2 = f1         //赋值
	f2()   //调用



	// 定义匿名函数同时调用
	func() {
		fmt.Printf("a = %d, str = %s\n", a, str)
	}() // 后面的 （） 代表调用次匿名函数

	// 带参数的匿名函数
	f3 := func(i int, j int) {
		fmt.Printf("i = %d, j = %d\n", i, j)
	}
	f3(11, 22)

	// 定义匿名函数，同时调用
	func(i int, j int) {
		fmt.Printf("i = %d, j = %d\n", i, j)
	}(1, 2)

	// 匿名函数，有参有返回值
	x, y := func(i int, j int) (max, min int) {
		if i > j {
			max = i
			min = j
		} else {
			max = j
			min = i
		}
		return
	}(123, 321)

	fmt.Printf("x = %d, y = %d\n", x, y)

	// 匿名函数，有参有返回值
	f4 := func(i int, j int) (max, min int) {
		if i > j {
			max = i
			min = j
		} else {
			max = j
			min = i
		}
		return
	}
	max, min := f4(111, 222)
	fmt.Printf("max = %d, min = %d\n", max, min)
}

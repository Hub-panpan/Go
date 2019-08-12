package main

import "fmt"

func main() {
	//多调用几次 结果都是1  函数被调用时
	// fmt.Println(test01())
	// fmt.Println(test01())
	// fmt.Println(test01())
	// fmt.Println(test01())

	//使用一个变量直接接 接的是返回的匿名函数
	f := test02()
	//f来调用闭包函数 闭包的使用的变量 不关心是否超出作用域 也就是继续存在着 不会被释放
	//即 闭包中的变量不遵循变量的生命周期
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

}

func test01() int {
	// 函数被调用时，i才分配空间，才初始化为 0
	var i int // 没有初始化，默认值为 0
	i++
	return i * i // 函数调用完毕，i 的值才释放
}

func test02() func() int {  //注意 func() int 返回的是匿名函数
	var x int
	return func() int {
		x++
		return x * x
	}
}

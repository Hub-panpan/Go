package main

import "fmt"

func main() {
	//不同类型变量的声明（定义）
	// var a int
	// var b float32

	// var (
	// 	a int     = 5
	// 	b float64 = 1.2
	// )

	// 这里推荐更加简便的一种方法  自动推导类型
	var (
		//注意哦 这里没有加变量类型 实际上是可以自动推到类型的 
		a = 3
		b = 3.14
	)

	a, b = 10, 3.14
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	// const i int = 10
	// const j float64 = 3.14

	// const (
	// 	i int     = 1
	// 	j float64 = 1.1
	// )

	// 常量也是一样的  自动推导类型
	const (
		i = 3
		j = 3.14
	)

	fmt.Println("i = ", i)
	fmt.Println("j = ", j)

}

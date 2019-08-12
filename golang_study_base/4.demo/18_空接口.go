package main

import "fmt"
//空接口 实际上是一个万能类型 可以保存任意类型的值
func main() {
	var i interface{}
	i = 1
	fmt.Println("i = ", i)

	i = "abv"
	fmt.Println("i = ", i)
}

// func xxx(args ...interface()){

// }

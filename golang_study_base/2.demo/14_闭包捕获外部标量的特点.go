package main

import "fmt"

func main() {
	a := 1
	str := "hello"
	func() {
		// 闭包以引用的方式捕获外部变量  也就是直接该表值  而不是副本
		a = 777
		str = "go"
		fmt.Printf("内部：a = %d, str = %s\n", a, str)
	}()//代表直接调用

	fmt.Printf("外部：a = %d, str = %s\n", a, str)
}

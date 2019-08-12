package main

import "fmt"

//类型别名 非常很重要 在以后的结构体 方法中都会用的到
func main() {
	// 给int64 起个乳名 叫做很大的一个整形的数
	type bigint int64
	var a bigint // 等价于 var a int64
	fmt.Printf("a type is %T\n", a)

	//如果是起多个别名的话 和引入包是一样的 
	type (
		long int64
		char byte
	)

	var b long = 11
	var ch char = 'a'
	fmt.Printf("b = %d, ch = %c\n", b, ch)
}

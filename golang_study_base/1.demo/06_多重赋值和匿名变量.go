package main

import "fmt"
//go语言可以返回多个返回值  如果返回一个的话 就不用加括号 完美的匹配abc
func test() (a, b, c int) {
	return 1, 2, 3
}

func main() {
	//go语言特殊的一点可以多重赋值 
	/*
		传统的写法 很多条
	*/
	i, j := 10, 20
	// 只有 Pritf 可以使用格式化输出， 格式化指带有 %d 的
	fmt.Printf("i = %d, j = %d\n", i, j)
	//传统的变量的交换比较麻烦 C/java
	i, j = j, i
	fmt.Printf("i = %d, j = %d\n", i, j)

	// _ 匿名变量，丢弃数据不处理，_ 匿名变量配合函数返回值使用，才有优势
	var f, d, e int
	f, d, e = test()
	fmt.Printf("c = %d, d = %d, e = %d\n", f, d, e)
	
	//但是假设我f不想用 直接丢了就可以了啦
	_, d, e = test()
	fmt.Printf("d = %d, e = %d\n", d, e)
}

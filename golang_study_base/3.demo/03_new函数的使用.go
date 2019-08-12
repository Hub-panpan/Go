package main

import "fmt"

func main() {


	var p *int
	p = new(int) // 指向一块空内存 给了一块整型的操作空间
	fmt.Println("p = ", p)

	*p = 666
	fmt.Println("*p = ", *p)

	q := new(int) // 自动推导类型
	*q = 777
	fmt.Println("*q = ", *q)
	fmt.Println("q = ", q)
}

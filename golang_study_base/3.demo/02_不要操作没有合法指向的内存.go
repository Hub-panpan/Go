package main

import "fmt"

func main() {
	var p *int
	fmt.Printf("p = %v\n", p)
	//没有赋值  默认值为nil

	// *p = 666  // error 不是合法的指向 无效的内存 一定要有合法的指向才能够赋值
	// fmt.Printf("p = %v\n", p)

	a := 111
	p = &a // p 指向 a
	*p = 666
	fmt.Printf("a = %v\n", a)
}

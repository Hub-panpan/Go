package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	Modify(&a) //传值的时候 使用传递的地址
	fmt.Println("a = ", a)
}

// p指向数组a,它是指向数组，它是数组指针
// *p 代表指针所指向的内存地址，就是实参a
func Modify(p *[5]int) { //取地址则就是值 *[5]int 这是数组指针类型

	(*p)[0] = 666
	//内部进行改变
	fmt.Println("a = ", *p)
}

package main

import "fmt"
//结构体
type Student struct {
	id   int
	name string
	age  int
	sex  byte
	addr string
}

func main() {
	var p1 *Student = &Student{1, "mike", 11, 'm', "address"}
	fmt.Println("*p1 = ", *p1)
	//直接这样打印也是可以的  但是感觉怪怪的
	fmt.Println("p1 = ", p1)
	p2 := &Student{name: "jack", age: 22}
	fmt.Printf("*p2 type is %T\n", *p2)
	fmt.Println("*p2 = ", *p2)
	fmt.Println("p2 = ", p2)
}

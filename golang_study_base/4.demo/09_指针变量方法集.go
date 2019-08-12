package main

import "fmt"

type Person struct {
	id   int
	name string
	age  int
}

func main() {
	// var p Person = &Person{1, "go", 22}
	// 结构体变量是一个指针变量，它能够调用的方法的集合，简称方法集
	p := &Person{1, "go", 22}
	// 内部做了方法自动转换 实际上它是非常的灵活
	//先把指针p 转成*p就是值了再调用
	p.setInfoPointer()
	(*p).setInfoPointer()
	//调用的时候 不受方法的约束 编辑器总是查找全部的方法 并且进行自动转换receiver实参
	p.setInfoValue()
	(*p).setInfoValue()
}

func (tmp Person) setInfoValue() {
	fmt.Println("setInfoValue")
}

func (tmp *Person) setInfoPointer() {
	fmt.Println("setInfoPointer")
}

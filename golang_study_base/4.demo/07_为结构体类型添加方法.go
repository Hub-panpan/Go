package main

import "fmt"


//接受者不能是指针 接口   因为是抽象的不是具体的实例  没办法绑定方法
type Person struct {
	id   int
	name string
	age  int
}

func main() {
	//定义同时初始化
	p := Person{1, "mike", 11} //这是我的属性值
	//因为已经给结构体绑定了方法 因此可以通过结构体对象调用方法
	p.PrintInfo()

	var p2 Person
	//你要是想改内部的内容的值 那么就得传地址
	(&p2).SetInfo(2, "jane", 22) // 取地址传参  (&p2)就是指针类型 放的是地址
	p2.PrintInfo()
}

// 带有接收者的函数叫方法  为结构体添加一个方法
func (tmp Person) PrintInfo() {
	fmt.Println("tmp = ", tmp)
}

// 通过一个函数，给成员赋值  就是为结构体添加一个设置的方法
func (p *Person) SetInfo(id int, name string, age int) {
	p.name = name
	p.id = id
	p.age = age
}

// 只要接收者类型不一样，就算函数同名，也是不同的方法，不会冲突
// type long *int  // error 不能直接声明为指针类型，只能在函数使用时声明为指针
type long int

func (tmp *long) test() { // 函数使用时声明为指针

}

type char byte

func (tmp char) test() {

}




//pointer为接受者类型 他本身不能是指针类型
/*type pointer *int
func (temp pointer) testxx(){

}*/


//不支持重载   接受者类型不一样 下面会认为是两个不同方法 不是用一个函数 不会出现重复定义函数的错误
//func (tmp char) test() {}
//func (tmp long) test() {}
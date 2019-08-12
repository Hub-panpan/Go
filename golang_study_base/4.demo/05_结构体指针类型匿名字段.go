package main

import "fmt"

type Person struct {
	id   int
	name string
	sex  byte
}

type Student struct {
	*Person // 指针类型的结构体   初始化和之前就会不一样了
	age     int
	addr    string
}

func main() {
	var s Student = Student{&Person{1, "go", 'c'}, 33, "add"} // 指针接数据必须使用 & 符号
	fmt.Printf("s = %+v\n", s)
	fmt.Println(s.id, s.name, s.sex, s.age, s.addr)

	var s2 Student
	s2.Person = new(Person) // 申请内存地址，分配空间  因为是指针类型的有两种类型进行初始化操作 
	s2.id = 2
	s2.name = "c++"
	s2.age = 22
	fmt.Println(s2.id, s2.name, s2.age)
}

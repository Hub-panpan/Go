package main

//没有面向对象特性 封装 继承 多态 没有析构函数 啥的

//但是 我们可以实现以上的特性来实现面向对象编程

/*
封装：通过方法实现
继承：通过匿名函数实现
多态：通过接口实现
*/




//继承的主要工作是实现代码的复用
import "fmt"

type Person struct {
	id   int
	name string
	sex  byte
}

type Student struct {
	Person // 只有类型，没有名字，这个就是匿名字段，继承了 Person 的成员
	age    int
	addr   string
}

func main() {
	// 顺序初始化 必须一个不能少
	var s1 Student = Student{Person{1, "mike", 'v'}, 13, "cq"}
	fmt.Println("s1 = ", s1)

	// 自动推导类型
	s2 := Student{Person{2, "jack", 'e'}, 11, "cq"}
	fmt.Printf("s2 = %+v\n", s2) // %+v 详细打印

	// 指定成员初始化，没有初始化的成员默认为 0
	s3 := Student{age: 22}
	fmt.Printf("s3 = %+v\n", s3)

	s4 := Student{Person: Person{id: 3, name: "jane"}, addr: "cc"}
	fmt.Printf("s4 = %+v\n", s4)
}

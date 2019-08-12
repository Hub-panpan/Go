package main

import "fmt"

type Student struct { //结构体成员 是没有var 关键字的
	id   int
	name string
	age  int
	sex  byte
	addr string
}

func main() {
	// 顺序初始化，每个成员必须赋值 不然会报错
	var s1 Student = Student{1, "mike", 14, 'm', "地球"}
	fmt.Println("s1 = ", s1) //里面的m没有格式化的话 会打印的是ASCII码

	// 指定成员初始化，没有初始化的成员，自动赋值为0
	s2 := Student{name: "go", sex: 's'}
	fmt.Println("s2 = ", s2)
}

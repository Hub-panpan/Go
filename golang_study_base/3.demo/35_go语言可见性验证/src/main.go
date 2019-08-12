package main

import (
	"fmt"
	"test"
)

func main() {
	test.Test()

	var s test.Stu
	s.Id = 111
	s.Name = "goooo"
	fmt.Println("s = ", s)
}

//函数名字的大小写决定了访问方式
//如果想使用别的包的函数 结构体 结构体的成员
//函数名 类型吗 结构体成员变量名 首字母必须大写
//小写的话只能是在内部进行访问
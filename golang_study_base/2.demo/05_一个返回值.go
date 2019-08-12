package main

import "fmt"

func main() {
	var a int
	a = MyFunc()
	fmt.Println("a = ", a)

	b := MyFunc()
	fmt.Println("b = ", b)

	c := MyFunc01()
	fmt.Println("c = ", c)

	d := MyFunc02()
	fmt.Println("d = ", d)
}

func MyFunc() int { //只有一个返回值int类型的
	return 6666 //需要return中断函数 进行返回值
}

// 给返回值取一个变量名，go 推荐使用  缺点 代码冗余了
func MyFunc01() (result int) {
	return 01
}

//这种是常用的写法
func MyFunc02() (result int) {
	result = 2
	return
}

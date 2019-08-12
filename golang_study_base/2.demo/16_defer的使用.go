package main

import "fmt"

func main() {
	// defer 延迟调用， main 函数结束前调用  类似于析构函数做一些工作
	//比如客户端关闭前程序清理
	defer fmt.Println("111111111")

	fmt.Println("222222222222")
}

package main

import (
	"fmt"
	"time"
)

func main() {//主协程 
	// 在函数调用前添加 go 关键字，创建并发执行单元

	//两个携程是不相关的任务  因此打印起来是没有规律的

	go test() // 新建一个携程，新建一个任务    记得一定要放前面

	for {
		fmt.Println("this is main goroutine")
		time.Sleep(time.Second) // 延时执行1秒
	}
	//go test() //放在后面是不会被调用的
}

func test() {
	for {
		fmt.Println("this is test goroutine")
		time.Sleep(time.Second)
	}
}

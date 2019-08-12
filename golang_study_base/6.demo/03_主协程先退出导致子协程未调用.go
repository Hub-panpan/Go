package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			fmt.Print("child goroutine")
			time.Sleep(time.Second)
		}
	}() // 调用匿名函数
}

//没来得及打印 就结束了  确实是创建了两个携程
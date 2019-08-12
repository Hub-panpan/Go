package main

import (
	"fmt"
	"runtime"
)

func main() {
	//一定记得放在前面
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("child 01 ")
		}
	}()
    //一定记得放在前面
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("child 02")
		}
	}()

	for i := 0; i < 2; i++ {
		runtime.Gosched() // 让出时间片，让其他子协程先执行，执行完了，再执行此函数
	}
}

package main

import (
	"fmt"
	"runtime"
)

func main() {

	go func() {
		fmt.Println("aaaaaaaaa")

		test()

		fmt.Println("bbbbbbbbb")
	}()

	// 特地写一个循环，不让主协程结束
	for {

	}
}

func test() {
	defer fmt.Println("cccccccccccc")

	// return // 终止此函数   aaaccccbbb 
	runtime.Goexit() // 终止test() 所在的协程 只会打印aaaccc 

	fmt.Println("dddddddddddd")
}

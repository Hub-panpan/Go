package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := runtime.GOMAXPROCS(4) // 指定以4核运算
	fmt.Println(n)

	for {
		go fmt.Print("1")

		fmt.Print("0")
	}

}

//并行的话 是多个队伍 排队买多个咖啡店的咖啡
//并发的话 是时间片轮转

//设定 多核以后 效率就会很高
package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second)//新建一个定时器
	isOk := timer.Reset(1 * time.Second) // 重新设定时间为1秒
	fmt.Println(isOk)

	go func() { //开设一个携程 
		<-timer.C // 延时3S 结果在mian中被停止了 
		fmt.Println("子协程开始执行")
	}()

	//timer.Stop() // 停止定时器  一直会没有打印

	for {

	}

}

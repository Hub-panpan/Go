package main

import (
	"fmt"
	"math/rand" //首先需要导入这个包 调用当中的函数
	"time"
)

func main() {
	// 设置种子，只需一次
	// 如果种子参数一样，每次运行程序产生的随机数都一样
	// rand.Seed(555) 这样产生的随机数是相同的就是每次产生n个不同的数是相同的
	// 场景 验证码 随机数等等
	rand.Seed(time.Now().UnixNano()) //seed(整型的参数)
	for i := 0; i < 5; i++ {

		// 产生随机数
		fmt.Println("rand = ", rand.Int()) // 产水很大的随机数
		fmt.Println("rand = ", rand.Intn(10)) // 限制在10以内的数
	}
}

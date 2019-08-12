package main

import "fmt"
import "time"

func main() {
	i := 0
	for {
		i++
		time.Sleep(time.Second) // 延时1S

		if i == 5 {
			//break // 跳出循环，如果嵌套多个循环，跳出最近的那个循环 还没来得及打印5 最退出了
			continue // 跳过本次循环，没打印5，下一次继续
		}
		fmt.Println("i = ", i)
	}
}


// break 用在循环 选择（switch） selector中
//continue 
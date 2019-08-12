package main

import "fmt"

func main() {
	sum := 1
	//初始化条件 

	for i := 1; i <= 100; i++ { //初试条件 ；判断条件；条件变化
		sum += i
	}
	fmt.Println("sum = ", sum)
}



//没有while  
//没有do while
package main

import (
	"fmt"
	"math/rand"
	"time"
)
//只要是学习了数组 那么久必须涉及到排序
func main() {
	// 设置种子，只需一次 每次产生的种子是不一样的 所以得到的随机数也是不一样的 
	rand.Seed(time.Now().UnixNano())

	var a [10]int //声明一个长度为10的数组
	n := len(a) //得到数组a的长度 

	//准备给数组进行赋值
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(100) //产生随机数
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}
	//进行冒泡排序
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	fmt.Printf("\n冒泡排序，升序\n")

	for i := 0; i < n; i++ {
		fmt.Println(a[i])
	}

	fmt.Println(a)

}

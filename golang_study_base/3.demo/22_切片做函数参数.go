package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 切片作为函数参数传入到函数时，在函数中改变的值就是切片的值
// 切片不同于数组 数组是值传递 切片是引用传递 因此会直接改变切片的值
func main() {
	n := 10
	// s := []int{}
	// 使用 make 创建一个切片
	s := make([]int, n, n) //整型类型的切片

	InitData(s) //初始化数据

	fmt.Printf("初始化：%v\n", s)
	BubbleSort(s)
	fmt.Printf("排序后：%v\n", s)
}

// 产生随机数
func InitData(s []int) {
	n := len(s)
	//使用随机数 必须要导入相关的包
	rand.Seed(time.Now().UnixNano())
	//给每个元素进行赋值
	for i := 0; i < n; i++ {
		//产生的数要是有要求的  100以内
		s[i] = rand.Intn(100)
	}
}

// 冒泡排序
func BubbleSort(s []int) {
	n := len(s)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

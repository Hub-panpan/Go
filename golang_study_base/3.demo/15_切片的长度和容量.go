package main

import "fmt"

//本质上 切片不是数组 是利用一种结构实现动态数组的效果
func main() {
	a := [5]int{1, 2, 3, 0, 0}
	s := a[0:2:5] // 从下标 0 开始，到下标 2 结束，不包括下标 2 的值 （左闭右开）,
	fmt.Println("s = ", s)
	fmt.Println("len(s) = ", len(s)) // 数组的长度  2-0=2
	fmt.Println("cap(s) = ", cap(s)) // 数组的容量  5-0=5

	fmt.Println("===============================")

	s = a[1:3:5]
	fmt.Println("s = ", s)
	fmt.Println("len(s) = ", len(s))
	fmt.Println("cap(s) = ", cap(s))

}

/*
[low :high:maxNum] cap=maxNum-low指的是可以容纳多少个元素
实际元素的个数是 high-low 也就是长度

*/
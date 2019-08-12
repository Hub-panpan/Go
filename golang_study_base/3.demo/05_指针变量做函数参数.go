package main

import "fmt"

func main() {
	a, b := 1, 9
	swap(&a, &b)//地址传递 表示交换的时候传的地址 能够找到存放数值的房子 也就是说
	//是有指针的指向的 实际上指的是一块内存的地址
	fmt.Printf("a = %d, b = %d\n", a, b)
}

func swap(p1, p2 *int) {
	*p1, *p2 = *p2, *p1 // 操作指针，记得写 *  通过*才能够取得值
}

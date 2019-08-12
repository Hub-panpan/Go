package main

import "fmt"

func main() {
	//声明一个切片
	s := []int{}
	//打印的话 值都是0
	fmt.Println("s = ", s)
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))

	// 在原切片的末尾追加元素 如果容量不够的话 会自动的扩容
	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)
	//看看长度变化  通常是在元素的末尾添加元素 且容量大概以2倍的大学扩容
	fmt.Println("s = ", s)
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))

	s1 := []int{1, 1, 1}
	fmt.Println("s1 = ", s1)
	s1 = append(s1, 2)
	s1 = append(s1, 22)
	s1 = append(s1, 2222)
	s1 = append(s1, 22222)
	fmt.Println("s1 = ", s1)
	fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))

}

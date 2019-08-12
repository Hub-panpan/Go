package main

import "fmt"

// 切片会改变底层数组元素的值
func main() {
	//声明新的数组
	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("a = ", a)

	// 新切片
	s1 := a[1:5:6]
	// fmt.Println("s1 = ", s1)
	s1[0] = 111  //操作的切片实际上就是操作数组 所以会改变数组当中的值
	fmt.Println("s1 = ", s1)
	fmt.Println("a = ", a)

	// 另外新切片
	s2 := s1[1:4]
	s2[2] = 55555
	fmt.Println("s2 = ", s2)
	fmt.Println("a = ", a)
}

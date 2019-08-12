package main

import "fmt"

func main() {
	//有几个方括号 就是几维 就有几个大括号
	var a [3][4]int  //[3][4]int可以看成 是变量的类型
	k := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			k++
			a[i][j] = k
			fmt.Printf("a[%d][%d] = %d, ", i, j, k)
		}
		fmt.Printf("\n")
	}
		fmt.Println("直接打印a=",a)
	b := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println("b = ", b)

	//部分初始化 剩下的初始化的操作为0
	c := [3][4]int{1: {1, 2, 3, 4}}

	fmt.Println("c = ", c)
	

}

package main

import "fmt"
//加入我有个需求 表示我们班同学的50个学号
//那我就得写50   引入数组集合
//统一类型的集合  
func main() {
	var arr [50]int //有50个int类型的数字
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1 //给每一个元素赋值 并且进行打印
		fmt.Printf("arr[%d] = %d\n", i, arr[i])
	}
}

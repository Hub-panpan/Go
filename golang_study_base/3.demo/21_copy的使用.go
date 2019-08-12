package main

import "fmt"

func main() {
	//旧的切片
	oldSlice := []int{1, 2, 3}
	//新的切片
	newSlice := []int{0, 0, 0, 0, 0}

	//后面的是源  前面的石目的
	copy(oldSlice, newSlice)
	fmt.Println("oldSlice = ", oldSlice)
}

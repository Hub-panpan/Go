package main

import "fmt"
//当程序的错误致命时 例如 程序数组的访问越界 空指针引用等 因此使用panic函数就可以
func main() {
	TestA()
	TestB()
	TestC()
}

func TestA() {
	fmt.Println("aaaaaaaaaaaa")
}

func TestB() {
	// 显式调用 panic() 函数，导致程序中断
	panic("this is panic test")
	fmt.Println("bbbbbbbbbbbbbb")
}

func TestC() {
	fmt.Println("CCCCCCCCCCCCC")
}

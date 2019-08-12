package main

import "fmt"

func main() {
	// 数组做函数参数，它是值传递
	// 实参数组的每个元素都给形参数组拷贝一份 因此数组越大 运行的效率越低 因为拷贝非常的废内存
	// 形参数组是实参数组的复制品
	a := [5]int{1, 2, 3, 4, 5}
	Modify(a)
	fmt.Println("main ---》a = ", a)

}

func Modify(a [5]int) {
	a[0] = 666
	fmt.Println("Modify---》a = ", a)
}

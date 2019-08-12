package main


import "fmt"
//案例 求最大最小值
func main() {//程序入口
	max, min := MaxAndMin(1, 11)
	fmt.Printf("max = %d, min = %d\n", max, min)

	// _ 使用匿名变量丢弃某个返回值
	m, _ := MaxAndMin(111, 222)
	fmt.Println("m = ", m)
}

func MaxAndMin(a int, b int) (max int, min int) {//有参数 有返回值
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}

	return // 函数有返回值，一定要写 return
}

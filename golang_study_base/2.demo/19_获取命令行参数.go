package main

import "fmt"
import "os"

func main() {
	// 接受用户传递的参数，都是以字符串的形式传递
	// os.Args是string类型的数组 
	list := os.Args
	//获取在命令行 接受的参数
	n := len(list)

	fmt.Println("n = ", n)

	// for i := 0; i < n; i++ {
	// 	fmt.Printf("list[%d] = %s\n", i, list[i])
	// }

	for i, data := range list {
		fmt.Println("list[%d] = %s", i, data)
	}
}

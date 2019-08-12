package main

import (
	"fmt"
	"os"
)
//设备文件 屏幕输出设备fmt.println() 键盘分为输入设备fmt.Scan()
//磁盘文件 1.文本文件  2.二进制文件 
func main() {
	//os.Stdout.Close()        // 关闭输出   关闭后无法输出 os指的是当前系统  
	fmt.Println("are u ok?") //往标准输出设备(屏幕)写内容

	// os.Stdout 标准输出设备
	os.Stdout.WriteString("hello ?")

	// os.Stdin.Close() // 关闭输入
	var a int
	fmt.Println("请输入数字:")
	fmt.Scan(&a) //取地址 才可以改变值
	fmt.Println("a = ", a)
}
//为什么需要文件呢  内存数据是当程序结束运行的时候 数据丢失 文件的内容不会丢
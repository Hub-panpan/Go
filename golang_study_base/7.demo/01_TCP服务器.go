package main

import (
	"fmt"
	"net" //网络包
)

func main() {
	// 监听 第一个参数指的是network    本地ip可以省略
	//返回值 
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {//说明有问题不用玩下走了
		fmt.Println("err = ", err)
		return
	}

	defer listener.Close() //程序结束才关闭

	// 阻塞等待用户链接
	conn, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("err1 = ", err1)
		return
	}

	// 接收用户的请求
	buf := make([]byte, 1024) // 1024大小的缓冲区
	//conn接受 需要一个切片 数组
	n, err2 := conn.Read(buf)
	if err2 != nil {
		fmt.Println("err2 = ", err2)
		return
	}
	fmt.Println("buf = ", string(buf[:n])) // 切记字符转换为字符串

	defer conn.Close() //用完了 把用户的链接关闭了 
}


//通过netcat  nc 127.0.0.1 8000  发送的数据
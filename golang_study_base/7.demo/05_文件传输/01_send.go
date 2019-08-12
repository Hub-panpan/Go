package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

// 发送文件
func SendFile(path string, conn net.Conn) {
	// 以只读方式打开文件  拿到文件对象
	f, err := os.Open(path)
	//如果没有问题
	if err != nil {
		fmt.Println("os.Open err: ", err)
		return
	}
	//延时关闭一下
	defer f.Close()

	buf := make([]byte, 1024*4)
	// 读取文件内容，读多少发多少 核心代码
	for { //放在for循环中 不断的取数据
		n, err := f.Read(buf) // 从文件读取内容

		if err != nil {
			//到末尾了 结束成功
			if err == io.EOF {
				fmt.Println("文件发送完毕")
			} else {
				fmt.Println("f.Read err: ", err)
			}
			return
		}
		// 发送内容
		conn.Write(buf[:n]) // 给服务器发送内容
	}

}

func main() {
	fmt.Println("请输入需要传输的文件：")
	var path string
	fmt.Scan(&path)//取到当前的路径的值

	// 获取文件名 info.Name() 使用os.stat包进行获取文件信息
	info, err := os.Stat(path)
	//没有错误则说明获取文件成功
	if err != nil {
		fmt.Println("os.Stat err: ", err)
		return
	}

	// 主动连接到服务器 使用tcp 本机地址 通信端口为8000   返回一个连接对象
	conn, err1 := net.Dial("tcp", "127.0.0.1:8000")
	if err1 != nil {
		fmt.Println("net.Dial err: ", err1)
		return
	}

	//延迟 程序退出是 关闭链接对象
	defer conn.Close()

	// 给接收方，先发送文件名 文件名转换成切片  
	_, err = conn.Write([]byte(info.Name()))
	if err != nil {
		fmt.Println("conn.Write err: ", err)
		return
	}

	// 接收对方的回复，如果回复"ok",就发送文件
	var n int
	//声明一个空切片 来接受ok
	buf := make([]byte, 1024)
	//开始接受n
	n, err = conn.Read(buf)
	//如果没有问题
	if err != nil {
		fmt.Println("conn.Read err: ", err)
		return
	}

	if "ok" == string(buf[:n]) {
		// 发送文件内容 需要把路径和conn对象传进来
		SendFile(path, conn)
	}

}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	path := "./demo.txt" //操作的文件为 当前目录下的这个demo问价

	// writeFile(path)
	// readFile(path)
	readFileLine(path)
}

func writeFile(path string) {
	// 打开文件，新建文件  注意两个返回值
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	// 使用完毕，关闭文件  函数结束前的一瞬 调用
	defer f.Close()

	//往文件中写内容
	var buf string
	for i := 0; i < 10; i++ {
		// buf里面存储的是 i = 0  这个字符串存储在buf中
		buf = fmt.Sprintf("i = %d\n", i)//此函数是返回字符串的
		f.WriteString(buf)
	}

}

func readFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	defer f.Close()

	buf := make([]byte, 1024*2) // 创建一个2k的切片

	// n 代表从文件读取内容的长度
	n, err1 := f.Read(buf)
	if err1 != nil && err1 != io.EOF { // 文件出错，同时没有到结尾 io.eof代表结束符
		fmt.Println("err1 = ", err1)
		return
	}

	fmt.Println("buf = ", string(buf[:n])) // 必须转为string类型

}

// 读取一行数据  借助bufio 来实现按行读
func readFileLine(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	// 新建一个缓冲区，把内容先放在缓存区  然后从缓存区区内容
	r := bufio.NewReader(f)

	for {
		// 遇到 \n 结束读取，但是 \n 已经读取到了
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF { // 文件已经结束
				break
			}
			fmt.Println("err = ", err)
		}
		fmt.Printf("buf = #%s#\n", string(buf))
	}

}

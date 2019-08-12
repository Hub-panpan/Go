package main

import (
	"fmt"
	"strconv"
)

func main() {
	slice := make([]byte, 0, 1024) //字节类型的切片 
	slice = strconv.AppendBool(slice, true)
	// 第二个数为要追加的数，第3个为指定10进制方式追加
	slice = strconv.AppendInt(slice, 123, 10)
	slice = strconv.AppendQuote(slice, "abchellogoooo")

	fmt.Println("slice = ", string(slice)) // 转换为 string 后，在打印

	// 其他类型转换为字符串
	var str string
	str = strconv.FormatBool(false)
	fmt.Println("str = ", str)
	//以float64 -1指的是紧缩模式 指的是小数点的位数 建议查文档进行查看
	str = strconv.FormatFloat(3.14, 'f', -1, 64)
	fmt.Println("str = ", str)

	// 整型转字符串，常用
	str = strconv.Itoa(6666)
	fmt.Println("str = ", str)

	// 字符串转其他类型
	var flag bool
	var err error
	flag, err = strconv.ParseBool("t0rue") //注意此函数有两个返回值
	if err == nil {
		fmt.Println("flag = ", flag)
	} else {
		fmt.Println("err = ", err)
	}

	// 把字符串转换为整型
	a, _ := strconv.Atoi("675")
	fmt.Println("a = ", a)
}

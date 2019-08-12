package main

import (
	"fmt"
	"regexp" //正则表达式的包 
)
//模式匹配 文本操作 复杂而又强大 灵活 
func main() {
	buf := "  abc axxc  axc   a1c  a2c  go golang cc c++"
	reg1 := regexp.MustCompile(`a.c`)//反引号 .是任意字符
	fmt.Println("reg1 = ", reg1)
	reg2 := regexp.MustCompile(`a[0-9]c`)
	fmt.Println("reg2 = ", reg2)
	// 1) 解析规则，它会解析正则表达式，如果成功返回解释器 失败返回nil
	reg3 := regexp.MustCompile(`a[a-z]c`)
	fmt.Println(reg3)
	if reg3 == nil { // 解释失败，返回nil
		fmt.Println("rexexp error")
		return
	}

	//2) 根据规则提取关键信息
	result1 := reg1.FindAllStringSubmatch(buf, -1) // -1指的是 返回所有匹配的结果
	fmt.Println("result1 = ", result1)
	result2 := reg2.FindAllStringSubmatch(buf, -1) // 1指的是 返回所有匹配的结果的一个
	fmt.Println("result2 = ", result2)
	result3 := reg3.FindAllStringSubmatch(buf, -1) // -1指的是 返回所有匹配的结果
	fmt.Println("result3 = ", result3)

}

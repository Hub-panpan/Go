package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "11.2 .2 3. 333.3 56.4 4.4 553.4"

	// 解释正则表达式 写一个复杂的小案例 匹配所有的小数
	reg := regexp.MustCompile(`\d+\.\d+`) // esc 下面的 ` 匹配前一个字符的一次或者多次
	if reg == nil {//如果出现了 nil 表示出问题了
		fmt.Println("compile err")
		return
	}
	reg2 := regexp.MustCompile(`\d\.\d`) 
	// 提取关键信息
	res1 := reg.FindAllString(buf, -1) // -1 返回所有满足的结果
	fmt.Println(res1)
	res2 := reg2.FindAllStringSubmatch(buf, -1)
	fmt.Println(res2)
}

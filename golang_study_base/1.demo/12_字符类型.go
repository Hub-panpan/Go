package main

import "fmt"

func main() {
	//计算机只能存数 不能存字符 美国人就想  于是一个字母对应一个英文的字符就好啦 引入ASCII表
	//65  ---A
	//97 ----a
	//验证一下
	var ch byte // 声明字符类型
	ch = 97
	// 格式输出，%c 以字符方式打印，%d 以整型方式打印
	fmt.Printf("%c, %d\n", ch, ch)
	//现在反过来看看效果   
	ch = 'a' // 注意字符 ， 单引号
	fmt.Printf("%c, %d\n", ch, ch)

	// 大写转小写，小写转大写，看看他们之间的关系是:大小写相差32
	fmt.Printf("大写: %d, 小写: %d\n", 'A', 'a')
	fmt.Printf("大写转小写: %c\n", 'A'+32)
	fmt.Printf("小写转大写: %c\n", 'a'-32)

	// '\' 以反斜杠开头的字符是转义字符，'\n' 代表换行
	//打印字符 与转义字符  
	fmt.Printf("hello go%c", '\n')
	fmt.Printf("hello itcast")
}

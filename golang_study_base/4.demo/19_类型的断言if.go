package main

import "fmt"

type Student struct {
	id   int
	name string
}

// 类型查询，也叫 类型断言    反过来查询一个值得变量的类型
func main() {
	// 新建一个长度、容量为3的接口类型的切片  实际上 这是空接口变量
	//这个就很有意思 空间口里什么都能存 但是你却不知道存的是什么类型
	//go语言自动类型转换
	i := make([]interface{}, 3) 
	i[0] = "hello" //String
	i[1] = 123		//int类型	
	i[2] = Student{66666, "mike"} //结构体类型   随便存

	// 第一个返回的是元素的下标，第二个返回的是元素的值
	for index, data := range i { // 递归
		// 第一个返回的是值，第二个返回判断结果的真假
		// 依次判断是什么类型
		if value, ok := data.(int); ok == true {
			fmt.Printf("x[%d] 类型为int,内容为 %d\n", index, value)
		} else if value, ok := data.(string); ok == true {
			fmt.Printf("x[%d] 类型为 string, 内容为 %s\n", index, value)
		} else if value, ok := data.(Student); ok == true {
			fmt.Printf("x[%d] 类型为 Student{}, 内容为id = %d, name = %s\n", index, value.id, value.name)
		}
	}

}

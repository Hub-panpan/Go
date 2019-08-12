package main

import "fmt"

type Person struct {
	id   int
	name string
	age  int
}

func main() {
	var p1 Person = Person{1, "go", 11} // 初始化赋值
	fmt.Printf("声明初始化的地址为&p1 = %p\n", &p1)       // %p 打印内存地址
	
	 p1.setInfoValue(2, "c++", 22)	// 普通传值
	 fmt.Println("普通传值后 在main中 p1 = ", p1)
	(&p1).setInfoPointer(3, "ggggg", 33) // 取地址传值

	fmt.Println("取地址传值在main中 p1 = ", p1)
}

// 值语义 值传递 参数为普通变量 非指针 一份拷贝
func (tmp Person) setInfoValue(id int, name string, age int) {
	fmt.Printf("值传递拷贝进来的副本&tmp的地址为： = %p\n", &tmp)
	tmp.id = id
	tmp.name = name
	tmp.age = age
	fmt.Println("值传递改变内部值tmp = ", tmp)
}

// 引用语义 指针传递  会改变值得哦
func (p *Person) setInfoPointer(id int, name string, age int) {
	fmt.Printf("指针传递拷贝进来的副本&tmp的地址为： = %p\n", p)
	p.id = id
	p.name = name
	p.age = age
	fmt.Println("指针传递改变内部值p = ", p)
}

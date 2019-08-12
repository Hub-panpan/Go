//go语言以包为管理单位
//每个文件都必须声明包
//如果程序要运行的话 必须要有一个main包 程序编译会不通过
package main
// 调用函数 大部分都要导入包 
// 导入的包必须要使用
import "fmt"

//有且只有一个入口函数 
func main() {//括号不能放下面去左边括号不能换行 必须和函数名同行
	// 先声明后赋值 初始化操作 是一体的
	// 变量声明了也必须得用
	var i int = 1
	i = 10
	i = 20
	i = 30

	fmt.Println("i = ", i)

	// 上面的是不是很麻烦啊！所以推荐下面的这个写法  := 自动推导类型  先声明B的类型 在赋值
	b := 20 //同一个变量名字只能写一次
	fmt.Println("b = ", b)
	//区别Println 与printf的区别  
	fmt.Printf("b type is %T\n", b)
	//如果再次写b := 30 前面已经有变量b 注意=  ：=的区别

	//变量是可以改变的值
	b = 30//这叫做赋值  赋值可以赋值很多次
	//ln会自动换行 
	fmt.Println("b = ", b)


	//只是声明没有初始化的值  默认值为0
	var moren int
    fmt.Println("moren = ", moren)





}


//如果使用liteIDE 一个工程只能有一个main函数的入口 
//只能删除 因为他是组织成一个工程来看

//go build xx 编译生成一个可执行文件
//go run xx 其实也生成可执行程序.exe 但是对用户透明
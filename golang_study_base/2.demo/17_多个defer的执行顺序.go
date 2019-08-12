package main

import "fmt"

// 多个 defer 会先进后出    在mian函数结束一霎那的之前  调用4个defer
func main() {
	defer fmt.Println("aaaaaaaa")
	defer fmt.Println("bbbbbbbbb")
	// 调用一个函数，导致内存出问题 非法的分母
	defer test(0)
	defer fmt.Println("ccccccc")
}
func test(x int) {
	result := 100 / x
	fmt.Println("result = ", result)
}




/*func main() {
	//此函数运行的话 会先打印bbbbb-->aaaaaaa--->出错

	 fmt.Println("aaaaaaaa")
	 fmt.Println("bbbbbbbbb")
	// 调用一个函数，导致内存出问题 非法的分母
	 test(0)
	 fmt.Println("ccccccc")  //这句话都到不了就崩了 所以不会打印
}*/



/*func main() { 
	//此函数运行的话 会先bbbbb---->aaaaaa---->报错   ccccc不会显示因为没执行到
	defer fmt.Println("aaaaaaaa")
	defer fmt.Println("bbbbbbbbb")
	// 调用一个函数，导致内存出问题 非法的分母
	   test(0)
	defer fmt.Println("ccccccc")
}*/
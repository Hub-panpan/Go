package main

import "fmt"

//这种结构比if else的结构更加的清晰
func main() {

	var i int
	fmt.Print("请按下楼层：")
	fmt.Scan(&i)
	switch i { //swith放的是条件本身
	case 1:
		fmt.Println("按下的是1楼")
		// break  // go语言默认添加 break  go保留了break关键字 
		fallthrough // fallthrough 不跳出 switch 语句，后面的无条件执行
	case 2:
		fmt.Println("按下的是2楼")
	case 3:
		fmt.Println("按下的是3楼")
	case 4:
		fmt.Println("按下的是4楼")
	default://指的是其他的情况 
		fmt.Println("按下的是X楼")
	}

/*  
	var num int
	switch num {
	case 1:
		fmt.Println("按下的是1楼")
	case 2:
		fmt.Println("按下的是2楼")
	case 3:
		fmt.Println("按下的是3楼")
	case 4:
		fmt.Println("按下的是4楼")
	//其他情况 都不属于我的情况
	default:
		fmt.Println("按下的是X楼")
	}

*/




}

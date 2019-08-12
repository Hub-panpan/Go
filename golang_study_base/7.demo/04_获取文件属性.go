package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args
	//list的长度必须为2才可以往下走
	if len(list) != 2 {
		fmt.Println("useage: xxx file")
		return
	}

	fileName := list[1]
	//stat
	info, err := os.Stat(fileName)
	if err != nil {
		fmt.Println("os.Stat err: ", err)
		return
	}

	fmt.Println("fileName = ", info.Name())
	fmt.Println("fileSize = ", info.Size())
	

}

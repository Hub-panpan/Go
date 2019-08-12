package main

import (
	"fmt"
	"time"
)

func main() {
	//就像是闹钟一样  是有周期的
	ticker := time.NewTicker(1 * time.Second)

	i := 0
	for {
		<-ticker.C
		i++
		fmt.Println(i)

		if i == 5 {
			ticker.Stop()
			break
		}
	}

}

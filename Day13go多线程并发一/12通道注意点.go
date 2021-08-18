package main

import (
	"fmt"
)

func main() {
	/*
		通道注意点：

	*/
	for i := 1; i <= 3; i++ {
		go func(j int) {
			str := fmt.Sprint("第", j, "个goroutine\n")
			fmt.Println(str)
		}(i)
		//defer func(j int) {
		//	str := fmt.Sprint("第", j, "个goroutine\n")
		//	fmt.Println(str)
		//}(i)
		//time.Sleep(time.Second)
	}
	//time.Sleep(2*time.Second)
	fmt.Println("main...over...")
}

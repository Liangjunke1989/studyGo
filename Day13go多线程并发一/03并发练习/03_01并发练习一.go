package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		练习1：
			一条协程打印1000数字，另一个协程打印1000字母。

		运行结果：
			并发的程序运行的结果，每次都不一定相同。
			不同的计算机设备，结果也不一定相同。
	*/
	go printStr(1000)
	go printNum(1000)
	time.Sleep(21)
	fmt.Println("main...over...")

}
func printNum(num int) {
	for i := 0; i < num; i++ {
		fmt.Println("打印的数字是：", i)
		time.Sleep(20)
	}
}
func printStr(num int) {

	for i := 0; i < num; i++ {
		fmt.Printf("打印的字母是：%d,%c\n", num, num)
		time.Sleep(20)
	}
}

package main

import (
	"fmt"
)

func main() {
	/*
		练习一：
			创建并启动两个子goroutine,一个打印100个数字，另一个打印100个字母，要保证在main goroutine结束前结束。
			利用通道实现。
	*/
	ch1 := make(chan bool)
	go printNum(100, ch1)
	go printStr(100, ch1)
	<-ch1
	<-ch1
	fmt.Println("main...over...")
}
func printNum(num int, ch chan bool) {
	for i := 0; i < num; i++ {
		fmt.Println("打印的数字是：", i)
	}
	ch <- true
}
func printStr(num int, ch chan bool) {
	for i := 0; i < num; i++ {
		fmt.Printf("打印的字母是：%d,%c\n", i, i)
	}
	ch <- true
}

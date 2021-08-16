package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		通道阻塞：

	*/
	ch1 := make(chan string)
	go func() {
		fmt.Println("子goroutine执行！！！")
		time.Sleep(5 * time.Second)
		ch1 <- "ljk"
	}()
	time.Sleep(2 * time.Second)
	data := <-ch1
	fmt.Println("main读取到的数据是：", data) //主程序需要"5s"后读到数据。  2s时间内，两个goroutine同时进行，还需要3s才能继续。
}

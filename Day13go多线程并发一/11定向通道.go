package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		通道：
			默认是双向通道。
			可读可写
			能存能取。
			make(chan Type)
		定向通道：
			也叫单向通道，只读或只写
			只读
				make(<-chan Type)
				只能读取数据
			只写
				make(chan<- Type)
				只能写入数据

		创建通道时，采用单向通道，没有意义！
		传递参数的时候使用：
			函数，只有写入数据。
			函数，只有读取数据。
		语法级别保证通道的操作安全。
	*/
	var ch1 = make(chan string)
	fmt.Println("---------------------------------------01.双向通道------------------------------------------------")
	{
		//go func(ch chan string) {
		//	ch <- "我是小明！"
		//	//time.Sleep(2*time.Second)
		//	str := <-ch
		//	fmt.Println("读取到的数据是：", str)
		//}(ch1)
		//
		////time.Sleep(2*time.Second)
		//data := <-ch1
		//fmt.Println("获取到的数据是：", data)
		////time.Sleep(2*time.Second)
		//ch1 <- "你要上学吗？"
		//time.Sleep(2 * time.Second)
		//fmt.Println("main...over...")
	}
	fmt.Println("---------------------------------------02.定向通道------------------------------------------------")
	{
		//只写的，定向通道
		go func(ch chan<- string) {
			ch <- "我是小明！"
		}(ch1)
		data := <-ch1
		fmt.Println("\t只写通道完成！------", data)
		fmt.Println("main...over...")
		fmt.Println("-------------------------------------------------------------------------------------------")
		time.Sleep(2 * time.Second)
		//只读的，定向通道
		ch2 := make(chan string)
		go func(ch <-chan string) {
			data1 := <-ch
			fmt.Println("\t只读通道完成！------", data1)
		}(ch2)
		ch2 <- "你要上学吗?" //通道顺序需要"先出后进"
		time.Sleep(time.Second)
		fmt.Println("main...over...")
		fmt.Println("-------------------------------------------------------------------------------------------")
	}
}

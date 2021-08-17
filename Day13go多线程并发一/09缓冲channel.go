package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		非缓冲通道：
			读，写都是即时阻塞。
			默认创建的通道。
		缓冲通道：
			自带一块缓冲区，可以暂时存储数据。
			设置缓冲区的大小。缓冲区存满后会阻塞。
		通道的本质是队列结构。FIFO
	*/
	fmt.Println("-------------------------------------1.非缓冲通道-----------------------------------------------")
	{
		ch1 := make(chan int)
		fmt.Println("非缓冲通道:", len(ch1), cap(ch1))
		go func() {
			//time.Sleep(3*time.Second)
			<-ch1
		}()
		ch1 <- 100 //即时阻塞
		fmt.Println("写完了！！！")
	}
	fmt.Println("-------------------------------------2.缓冲通道-------------------------------------------------")
	//02.缓冲通道：缓冲区满了才会阻塞
	{
		//	ch2 := make(chan int, 3)
		//	go func() {
		//		data := <-ch2
		//		fmt.Println("获取到的数据为：", data)
		//	}()
		//	//fmt.Println("缓冲通道：", len(ch2), cap(ch2))
		//	ch2 <- 1 //写入
		//	fmt.Println("缓冲通道：", len(ch2), cap(ch2))
		//	ch2 <- 2 //写入
		//	fmt.Println("缓冲通道：", len(ch2), cap(ch2))
		//	ch2 <- 3 //写入
		//	fmt.Println("缓冲通道：", len(ch2), cap(ch2))
		//	ch2 <- 4 //写入
		//	fmt.Println("缓冲通道：", len(ch2), cap(ch2))
		//	fmt.Println("main...over...")
	}
	fmt.Println("-------------------------------------3.带缓冲区的通道--------------------------------------------")
	{
		ch3 := make(chan string, 5)
		go sendStrData02(ch3)
		//for v:=range ch3{
		//	fmt.Println("获取到的通道数据为：",v)
		//}
		for {
			time.Sleep(time.Second)
			data, ok := <-ch3

			if !ok {
				fmt.Println("---------------------------------------据读取完毕！退出！")
				break
			}
			fmt.Println("\t获取到通道3的数据为：", data)
		}
	}
}
func sendStrData02(ch chan string) {
	for i := 1; i <= 30; i++ {
		ch <- fmt.Sprint("数据为：", i)
		fmt.Println("---已完成写出的数据为：", i)
	}
	close(ch)
}

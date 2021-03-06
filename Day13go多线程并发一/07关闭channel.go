package main

import (
	"fmt"
	"time"
)

func main() {
	/*

			读取通道：data:=<-chan
					data,ok:=<-chan    ok为true,表示通道正常，读取到的data数据有效可用。
		                              ok为false,表示通道关闭，读取到的data是类型的零值（默认值）。
		                               与value,ok:=map[key]类似
			通道关闭：
					发送方如果数据写入完毕，可以关闭通道，用于通知接收方没有数据了，数据传输完毕。
					一般只有发送方关闭。
					g1--->写入数据
					g2--->读取数据

					发送方，不停的写入数据
						循环发送，循环结束，发送停止写入
						最后需要关闭通道，用于通知接收方没有数据了
					接收方，不停的读取数据
						死循环，读取数据，根据ok的值判断，是否已经结束。
					for{
						data,ok:=<-chan
						if !ok{
							break
						}
					}
					for data :=range chan{ //直到发送方关闭通道为止，for range关闭
						data
					}
	*/
	ch1 := make(chan int)
	//发送数据
	go sendData(ch1)
	//读取数据
	for {
		time.Sleep(time.Second) //1s的睡眠
		data, ok := <-ch1
		if !ok {
			fmt.Println("读取完毕，通道关闭了...", data, ok)
			break //break跳出for循环
		}
		fmt.Println("main中读取到的数据", data, ok)
	}
}
func sendData(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	fmt.Println("发送方，写入数据完毕！！！")
	close(ch) //发送方关闭通道，通知接收方没有数据了。
}

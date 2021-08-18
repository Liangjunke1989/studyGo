package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("--------------------------01.select+timer-------------------------------")
	{
		//ch1 := make(chan int)
		//ch2 := make(chan int)
		//go func() {
		//	time.Sleep(2*time.Second)
		//	<-ch1
		//}()
		//go func() {
		//	time.Sleep(3*time.Second)
		//	<-ch2
		//}()
		//select {
		//	case ch1 <- 100:
		//		fmt.Println("ch1写入数据")
		//	case ch2 <- 200:
		//		fmt.Println("ch2写入数据")
		//	case <-time.After(2*time.Second):
		//		fmt.Println("超时a...")
		//	case <-time.NewTimer(1*time.Second).C:
		//	fmt.Println("超时b...")
		//	//default:
		//	//	fmt.Println("还没来得及写入数据，就已经完成！")
		//}
		//time.Sleep(2*time.Second)
		//fmt.Println("main...over...")}
	}
	fmt.Println("--------------------------02for+select+timer----------------------------")
	{
		ch1 := make(chan int)
		ch2 := make(chan int)
		go func() {
			time.Sleep(1 * time.Second)
			ch1 <- 100
			//close(ch1)
		}()
		//go func() {
		//	time.Sleep(1*time.Second)
		//	ch2<-200
		//}()

		//结束循环的条件：通道结束，或者超时5次以上 结束循环
		count := 0
	out:
		for {
			select {
			case data, ok := <-ch1: //1s后，通道堵塞了，没有数据了
				if !ok {
					fmt.Println("ch1的通道关闭了！！！")
					break out
				}
				fmt.Println("ch1读取到到数据为:", data, ok)
			case data := <-ch2: //
				fmt.Println("ch1读取到到数据为:", data)
			case <-time.After(2 * time.Second):
				fmt.Println("超时2s了...")
				count++
				if count > 5 {
					break out
				}
				//default:
				//	fmt.Println("数据还未来得及写入.....")
			}
		}
	}
}

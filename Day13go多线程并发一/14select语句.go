package main

import (
	"fmt"
	"time"
)

func main() {
	/*
			select语句：
				select语句类似于switch语句，但是select会随机执行一个可运行的case。
				如果没有case可以运行，它将阻塞，直到有case可运行。

				select是一种分支语句，专门用于通道读写操作的。
				select{
					case chan读/写：
						分支1
					case chan读/写：
		                分支2
					case chan读/写：
						分支3
					...
					default:
				}
				执行流程：
					1.每个case后的通道都会被执行，如果多个都可以执行，那么select会随机执行一个可运行的case。
					2.如果没有case可以运行，如果有default执行default，
		                                 如果没有default它将阻塞，直到有case可运行。
	*/
	fmt.Println("-------------------------------------1.select分支语句---------------------------------------------")
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		ch1 <- 100
		time.Sleep(2 * time.Second)
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 200
	}()
	select {
	case data := <-ch1:
		fmt.Println("ch1中读取数据了！", data)
	case data := <-ch2:
		fmt.Println("ch2中读取数据了！", data)
	default:
		fmt.Println("目前一直都没有数据！！！！") //default直接执行，不等待了！
	}
}

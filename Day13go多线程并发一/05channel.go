package main

import (
	"fmt"
	"time"
)

func main() {
	/*
			并发程序之间的数据传递：
				其他语言：
					通过共享内存实现消息传递。
				go语言：
					通过数据传递来实现共享内存。
			队列结构：
				队列queue：
					是一种数据结构----FIFO（先进先出）
		       	数据结构：数据存储的特点。
					数组、切片————线性结构。
					栈stack---LIFO（后进先出）
			channel:
				通道概念：
					专门的goroutines通信的管道。类似于管道中的水从一端到另一端的流动，数据可以从一端发送到另一端，通过通道接收。
				通道目的：
					实现多个goroutine之间实现数据传递。
				通道的本质：
					通道是一种容器，通道是一种数据结构---队列结构
					容器，相当于"消息队列"。
				通道的语法：
					make()创建,也是引用类型的数据。
					需要关联一种相关的数据类型。
					通道的操作：
							一个goroutine可以从chan中读取数据，另一个goroutine可以从中写入数据。
							不能一个goroutine又读又写，会产生阻塞。
					从chan中读取数据：
						<- 特殊操作符
						从chan中读取数据，data := <-chan    从chan中读取数据
						向chan中写入数据，chan<-data        将data的数据写入chan
					nil chan,同map一样不能使用，需要make()创建
				通道阻塞：
					对于chan的读取和写入操作，都是阻塞式的。
					阻塞式：导致程序暂时不能执行！  直到解除阻塞！！

					从chan中读取数据：阻塞式，直到另一个goroutine向通道中写入数据，才能解除阻塞。
					向chan中写入数据：阻塞式，直到另一个goroutine向通道中读取出数据，才能解除阻塞。

					g1,g2-->chan写入数据
					main-->从通道中读取数据

	*/
	fmt.Println("---------------------------------1.通道channel的声明----------------------------------------------")
	{
		//var ch1 chan int
		//fmt.Println(ch1)  //nil
		//fmt.Printf("ch1的数据类型是：%T\n",ch1)

		//ch1=make(chan int,0)
		//fmt.Println(ch1) //0xc00008c060  引用类型}
	}
	fmt.Println("---------------------------------2.通道channel的使用----------------------------------------------")
	{
		ch2 := make(chan int, 0)
		ch3 := make(chan bool)
		go func() { //匿名函数，可以不需要做形参传递，直接调用变量
			fmt.Println("这是一个子goroutine")
			data := <-ch2 //现将ch2数据写出，再将新数据写入到ch2中
			time.Sleep(5 * time.Second)
			fmt.Println("从main传来的数据是：", data)
			ch3 <- true
		}()
		ch2 <- 100 //向通道ch2中写入数据            //仅有写入数据会报错。阻塞式，deadlock死锁
		//goroutine阻塞式，不能自救
		<-ch3 //如果ch3可以往外写数据，说明ch3已经先读取到了数据，也就说明子线程 go匿名函数已经执行完成了！
		fmt.Println("main...over...")
	}
}

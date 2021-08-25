package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		生产者/消费者模型： ------线程之间的通信
			生产者：生产产品
			消费者：消费产品
				比如：
					母鸡：生产者             -------线程t1   生产 	notify()唤醒和notifyAll()唤醒所有
					产品：鸡蛋

					吃货：消费者             -------线程t2   需要使用wait()进行阻塞，取数据 wait()等待;一般不用sleep()
														   notify()唤醒母鸡
					商店：存储一定容量的鸡蛋   --------容器    通道， 队列
	*/

	ch1 := make(chan int, 5)
	ch2 := make(chan bool) //判断结束
	rand.Seed(time.Now().UnixNano())
	//写出数据:生产者
	go func() {
		for i := 1; i <= 30; i++ {
			ch1 <- i
			fmt.Println("生产者 写出的数据是：", i)
		}
		close(ch1)
	}()
	//读取数据:消费者1
	go func() {
		for data := range ch1 {
			fmt.Println("\t消费者1 读取到的数据为：", data)
			//time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		ch2 <- true
	}()
	//读取数据：消费者2
	go func() {
		for data := range ch1 {
			fmt.Println("\t消费者2 读取到的数据为：", data)
			//time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		ch2 <- true
	}()
	<-ch2
	<-ch2
	fmt.Println("main...over...")
}

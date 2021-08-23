package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
		读写锁：sync.RWMutex
			变量的访问，set/get 读写操作
			成对的指针方法：
				Lock(),写锁定
				Unlock(),写解锁

				RLock(),读锁定
				RUnlock(),读解锁
			锁定规则：
				读写锁的使用中
					写操作都是互斥的。
					读和写是互斥的。
					读和读不互斥。   //读基本不存在不安全性
			理解为：
				可以多个goroutine同时读取数据，
				但是写只允许一个goroutine写数据。
	*/
	var rwm sync.RWMutex
	for i := 1; i <= 3; i++ {
		go func(j int) {
			fmt.Printf("goroutine %d,-----尝试读锁定...\n", j)
			rwm.RLock()
			fmt.Printf("goroutine %d,-----已经读锁定了...\n", j)
			time.Sleep(5 * time.Second)
			fmt.Printf("\tgoroutine %d,尝试读解锁...\n", j)
			rwm.RUnlock()
		}(i)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("main尝试写锁定...")
	rwm.Lock()
	fmt.Println("main已经写锁定了...")
	fmt.Println("111111")
	fmt.Println("222222")
	rwm.Unlock()
	fmt.Println("main已经写解锁了...")
}

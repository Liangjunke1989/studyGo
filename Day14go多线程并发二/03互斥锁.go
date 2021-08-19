package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
		互斥锁：
			锁头对象 （struct）
			互斥：锁定---解锁
			有两个指针方法：
				Lock(),上锁
					阻塞的：只能有一个goroutine上锁，其他的goroutine处于阻塞状态。
				Unlock(),开锁
	*/

	var mutex sync.Mutex
	fmt.Println("main,即将锁定mutex")
	mutex.Lock() //锁定产生互斥
	fmt.Println("main已经锁定mutex")

	for i := 1; i <= 3; i++ {
		go func(j int) {
			fmt.Println("子goroutine", j, "即将锁定mutex")
			mutex.Lock() //阻塞
			fmt.Println("子goroutine", j, "已经锁定！！！")
		}(i)
	}
	time.Sleep(3 * time.Second)
	fmt.Println("main,即将解锁mutex")
	mutex.Unlock()
	fmt.Println("main,已经解锁mutex")
	//mutex.Unlock()//报错了！unlock of unlocked mutex
	//fmt.Println("main,想再解锁mutex")
	time.Sleep(2 * time.Second)

	fmt.Println("main...over...")

}

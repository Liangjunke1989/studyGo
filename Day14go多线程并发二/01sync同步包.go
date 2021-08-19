package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	/*
		同步/异步：
			同步：sync
				one by one 一个接一个，串行
			异步：async
				并行，同时执行
			异步任务：
				同时执行，并发。
	*/
	/*
		WaitGroup: 同步等待组
			执行了wait的goroutine，要等待同步等待组中其他的goroutine执行完毕。
			内置计数器：
				初始值为：counter:0
				要执行的goroutine的数量
			Add(3),设置计数器的数量        设置counter的值
			Done(),将内置的计数器的数值减1。同add(-1)  counter的值-1

			Wait(),等待，导致执行wait的goroutine进入阻塞状态。同步等待组中的计数器的值为0，解除阻塞。

		Add(3),三个goroutine要执行
		...
			子goroutine：函数，最后一行Done() 计数器减1
		wait()主程序等待
	*/

	wg.Add(2) //等待n个执行完毕
	go printNum1(&wg)
	go printNum2(&wg)
	wg.Wait() //主程序main函数中执行了Wait(),进入阻塞状态。底层计数器为0，解除阻塞。
	fmt.Println("main...over...")
}
func printNum1(wg *sync.WaitGroup) { //值类型传递地址，不用副本拷贝
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 100; i++ {
		fmt.Println("子goroutine1中，i:", i)
		//time.Sleep(time.Duration(rand.Intn(1000)))
	}
	wg.Done()
}

func printNum2(wg *sync.WaitGroup) {
	for j := 1; j <= 100; j++ {
		fmt.Println("\t子goroutine2中，j:", j)
		//time.Sleep(time.Duration(rand.Intn(1000)))
	}
	wg.Done()
}

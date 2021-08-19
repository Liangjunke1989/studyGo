package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ticks02 int = 2000
var wg02 sync.WaitGroup
var mutex02 sync.Mutex //创建一个互斥锁
func main() {
	/*
		上了互斥锁之后，one by one执行，执行效率下降，但数据是安全的。
	*/
	/*
		售票：四个窗口一块售出100张票
	*/
	wg02.Add(4)
	go saleTicks02("售票厅1")
	go saleTicks02("售票厅2")
	go saleTicks02("售票厅3")
	go saleTicks02("售票厅4")
	wg02.Wait()
	fmt.Println("main...over...")
}
func saleTicks02(name string) {
	rand.Seed(time.Now().UnixNano())
	for {
		mutex02.Lock()
		if ticks02 > 0 {
			time.Sleep(time.Duration(rand.Intn(5000)))
			fmt.Println(name, ":", ticks02)
			ticks02--
		} else {
			fmt.Println(name, "结束买票了！！！")
			mutex02.Unlock() //跳出之前，要先解锁
			break
		}
		mutex02.Unlock()
	}
	wg02.Done()
}

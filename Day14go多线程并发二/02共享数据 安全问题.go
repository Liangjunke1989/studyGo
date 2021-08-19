package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ticks int = 1000
var _wg sync.WaitGroup

func main() {
	/*
		共享数据 安全问题：
			并发读取全局数据时,会出问题。
			利用"锁"解决问题。
			互斥锁：
				开 / 关
				"上厕所"的例子

			阻塞的几种方式：
				1.读键盘
				2.sleep
				3.通道
				4.上锁
				...
	*/

	/*
		售票：四个窗口一块售出100张票
	*/
	_wg.Add(4)
	go saleTicks("售票厅1")
	go saleTicks("售票厅2")
	go saleTicks("售票厅3")
	go saleTicks("售票厅4")
	_wg.Wait()
	fmt.Println("main...over...")
}
func saleTicks(name string) {
	rand.Seed(time.Now().UnixNano())
	for {
		if ticks > 0 {
			time.Sleep(time.Duration(rand.Intn(5000)))
			fmt.Println(name, ":", ticks)
			ticks--
		} else {
			fmt.Println(name, "结束买票了！！！")
			break
		}
	}
	_wg.Done()
}

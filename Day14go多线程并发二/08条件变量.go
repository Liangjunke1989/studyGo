package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
		条件变量：sync.Cond
			L:Locker 接口
			Cond条件变量，总是要和锁结合使用

			01.作用
				----条件
				多个goroutine等待或接收通知的集合地
			02.3个指针方法
				Wait(),等待goroutine等待接收通知，通过Single(),Broadcast(),解除阻塞。
				Single() 发送通知 ，发送一个
				Broadcast() 广播，发送给所有人
	*/
	/*
			goroutine中的阻塞方式：
				01.读取键盘
						---键盘输入，解除阻塞
				02.waitGroup
						---底层的cont=0，
		                   add(),add(-1) done()     wait()完成解除阻塞
				03.chan的读写
						读写--写读
				04.cond wait()
						条件变量的阻塞---通过通知解除阻塞
						Single() 发送通知 ，发送一个     wait()完成解除阻塞
						Broadcast() 广播，发送给所有人   wait()完成解除阻塞
	*/

	var mutex100 sync.Mutex
	cond := sync.Cond{L: &mutex100}
	cond.L.Lock()
	fmt.Println("main....已经上锁...")
	conValue := false
	fmt.Println("主程序设置了更改的目标值conValue:", conValue)
	go func() {
		time.Sleep(2 * time.Second)
		cond.L.Lock()
		fmt.Println("\t子goroutine已经锁定了...")
		conValue = true //更改数值
		fmt.Println("\t子goroutine已经更改了条件数值，并发送通知...")
		cond.Signal() //发送通知：一个goroutine
		time.Sleep(5 * time.Second)
		fmt.Println("\t一个子goroutine可以继续...")
		cond.L.Unlock()
		fmt.Println("\t一个子goroutine已经解锁...")
	}()

	if !conValue {
		fmt.Println("main即将等待...")
		//wait尝试解锁，等待-------------当前的goroutine进入了阻塞状态，等待被唤醒...single()或Broadcast()
		//一旦被唤醒后，又会被锁定
		cond.Wait()
		fmt.Println("main...再次被唤醒...")
	}
	fmt.Println("main...继续工作...")
	cond.L.Unlock()
	fmt.Println("main...已经解锁")
	fmt.Println("main...over...")
}

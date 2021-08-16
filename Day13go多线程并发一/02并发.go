package main

import (
	"fmt"
	"time"
)

func main() {
	/*
			并发：
				go关键字，启动一个goroutine,执行对应的函数。
				主程序结束所有子程序随之销毁。
			程序的执行过程：
				1.先创建一个主goroutine,设定每一个goroutine能申请到的最大栈空间是32位计算机是250MB,64位计算机1GB.
				2.创建一个特殊的defer语句，用于在主goroutine退出时，做必要的善后处理。因为主goroutine也可能非正常结束。
				3.启动专用于后台清扫内存垃圾的goroutine，bin个设置GC可用的标识
				4.执行main包中的init函数
				5.执行main函数
		          执行完main函数后，它还会检查主goroutine是否引发了运行时恐慌，并执行必要的处理。最后主goroutine会结束自己以及当前进行的运行。
			go语言的并发：
				利用go关键字
				系统自动创建并启动主goroutine，执行对应的函数main()
				用于自己创建并启动子goroutine，执行对应的函数。go关键字 + 函数调用
				go 函数() //go关键字创建并启动goroutine，然后执行对应的函数()
				子goroutine中执行的函数，往往没有返回值。如果有也会被舍弃。
	*/

	//1.启动子goroutine
	go hello()
	time.Sleep(10) //10ns
	//2.执行main()中的打印函数
	fmt.Println("main()中的打印函数...") //主函数先结束，子goroutine也会销毁。
	//子协程和主协程，执行顺序不一定！看谁先抢占资源！

}
func hello() {
	fmt.Println("我是一个hello()函数，在另外一个goroutine中执行的......")
}

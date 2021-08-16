package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	/*
			init()函数：
				同main()一样，特殊的函数，由系统自动调用执行。------> main goroutine 主协程中

		 	runtime包:
				设置程序所需要的进程数量。
	*/
	fmt.Println("主程序执行：HelloWorld!!!")
	go printNum2()
	go printStr2()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("main...over...")
}
func init() {
	//1.获取逻辑cpu的数量
	fmt.Println("逻辑cpu的核数：", runtime.NumCPU())
	//2.设置go程序执行的最大的并发程序[1-256]
	runtime.GOMAXPROCS(runtime.NumCPU()) //默认值为1，
}
func printNum2() {
	for i := 1; i <= 5; i++ {
		fmt.Println("打印的数字是：", i)
	}
}
func printStr2() {
	for i := 65; i <= 70; i++ { //字母对应的数字
		fmt.Printf("打印的字母是：%c\n", i)
	}
}

package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		练习2：
			子goroutine1,打印1-5数字，先睡眠，每睡眠250毫秒，打印一个数字。
			子goroutine2,打印A-E数字，先睡眠，每睡眠400毫秒，打印一个字母。
			主goroutine中睡眠3000毫秒
		观察运行结果：
			按顺序执行，需要时间控制。
	*/
	go printNum1()
	go printStr1()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main...over...")
}
func printNum1() {
	for i := 1; i <= 5; i++ {
		fmt.Println("打印的数字是：", i)
		time.Sleep(250 * time.Millisecond)
	}
}
func printStr1() {
	for i := 65; i <= 70; i++ { //字母对应的数字
		fmt.Printf("打印的字母是：%c\n", i)
		time.Sleep(400 * time.Millisecond)
	}
}

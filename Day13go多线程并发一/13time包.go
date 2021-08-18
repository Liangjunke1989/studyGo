package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		time包对chan的操作
			1.Timer(): 计时器
				time.NewTimer(duration)  ----->*Timer对象   中的属性C为：<-chan time.Time
			2.After(duration):  返回值为 <-chan time.Time 同属性C
	*/

	//1.创建一个计时器
	timer1 := time.NewTimer(3 * time.Second) //3s之后
	fmt.Printf("%T\n", timer1)               //*time.Timer
	chan1 := timer1.C
	fmt.Printf("%T\n", chan1) //<-chan time.Time
	fmt.Println(time.Now())
	time1 := <-chan1
	fmt.Println(time1) //3s后渠道了时间
	//2.使用After(),返回值 ,获取只读通道类型的值
	ch2 := time.After(5 * time.Second)
	fmt.Println(time.Now())
	data := <-ch2
	fmt.Println(data)
}

package main

import (
	"fmt"
)

func main() {
	/*
		try...catch...finally

		defer：延迟
		panic:恐慌
			让当前的一条协程进入恐慌，中断程序的执行。
		recover:恢复
			让程序恢复，必须在defer函数中执行。

		go语言没有异常机制，它不能抛出异常，而是使用了panic和recover机制。--->把它作为最后的手段来使用。
		go语言利用panic(),recover(),实现程序中的极特殊的异常处理。
			panic(),让当前程序进入恐慌，中断程序的执行。
			recover(),让程序恢复，必须在defer函数中执行。没有defer，recover()执行无效！
			当外围函数中的代码引发运行恐慌时，只有其中所有的延迟函数都执行完毕之后，该运行时恐慌才会真正被扩展至调用函数。
	*/
	funA()
	funB()
	funC()
	fmt.Println("main...over...")
}
func funA() {
	fmt.Println("我是一个A函数funA()...")
}
func funB() { //外围函数
	defer func() { //延迟函数恢复
		msg := recover() //msg拿到的返回值就是panic()传入的参数.
		if msg != nil {
			fmt.Println("利用recover()恢复了！") //存在恐慌了，程序最后执行
		}
	}()
	fmt.Println("我是一个A函数funB()...")
	for i := 0; i < 10; i++ {
		if i == 5 {
			//让程序中断
			panic("funB()程序中断了！出现恐慌错误！！！") //打断程序的执行... 主函数也不走了。
		}
		fmt.Println("i:", i)
	}
}
func funC() {
	defer func() {
		fmt.Println("funC的延迟函数！") //执行了
		msg := recover()
		if msg != nil {
			fmt.Println("利用recover()执行了恢复！")
		} else {
			fmt.Println("没有利用recover()执行恢复！")
		}

	}()
	fmt.Println("我是funC函数！")
	panic("funC程序存在恐慌了！！！！")
}

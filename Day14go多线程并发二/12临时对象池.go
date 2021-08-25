package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	/*
		临时对象池：
			看成复用某些数据的容器。
			可伸缩。
			同步安全。

		sync.Pool:
			大量对象产生并销毁。
			Put()存数据，向pool中存储数据。
			Get()读数据，从对象池中获取一个数据，如果pool没有数据，那么会调用new对应函数，创建一个并返回。
	*/
	fmt.Println("---------------------------------------------")
	var count1234 int64
	fun001 := func() interface{} {
		return atomic.AddInt64(&count1234, 1)
	}
	pool := sync.Pool{New: fun001}
	//获取数据
	fmt.Println(pool.Get())
	fmt.Println("---------------------------------------------")
	//向pool存储，获取数据
	pool.Put(10)
	pool.Put(12)
	pool.Put(13)
	pool.Put(11)
	pool.Put(15)
	fmt.Println(pool.Get()) //pool中如果有多个数据，任意获取一个，如果没有，就new对应的函数，产生一个并返回。
	fmt.Println("---------------------------------------------")
	//执行GC()
	//如果程序当中垃圾回收机制来收垃圾，那么临时对象池中的数据都被销毁。
	runtime.GC()
	pool.New = nil
	fmt.Println(pool.Get())
}

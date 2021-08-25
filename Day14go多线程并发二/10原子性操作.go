package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	/*
				原子性操作：sync.atomic
					上锁--->解锁的过程。
					针对于数值。
					操作：
						原子加/减
						交换
						比较并交换（CAS）
						存储
		                交换
				非原子性操作：

				数据安全：
					同步上锁-解锁
	*/

	//n123:=100
	//n123+=1
	var n111 int64 = 3
	fmt.Println("n的原始数值为：", n111)
	//原子加
	//newN := atomic.AddInt64(&n, 1)//执行原子操作时，不允许被别的goroutine访问，打断。
	//fmt.Println("n的新值为：",newN)
	fmt.Println("-------------------")
	//原子交换
	swapInt64 := atomic.SwapInt64(&n111, 9)
	fmt.Println(n111, swapInt64)
	swappedN := atomic.CompareAndSwapInt64(&n111, 9, 88)
	fmt.Println(swappedN, n111)

}

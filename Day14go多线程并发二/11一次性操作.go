package main

import (
	"fmt"
	"sync"
)

func main() {
	/*
		一次性操作：sync.Once{}

			有些操作只需要执行一次。
	*/
	count := 0
	once := sync.Once{}
	for i := 1; i <= 10; i++ {
		once.Do(func() { //设置多次，也只会执行一次
			count++
		})
	}
	fmt.Println(count)
}

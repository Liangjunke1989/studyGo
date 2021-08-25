package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var n int
var wg001 sync.WaitGroup
var rwm001 sync.RWMutex

func main() {
	wg001.Add(10)
	for i := 1; i <= 5; i++ {
		//read
		go func(k int) {
			defer wg001.Done()
			rwm001.RLock() //通过读锁定，来防止别人来写
			fmt.Println("读操作...", k, "即将读取数据...")
			v := n
			fmt.Println("读操作", k, "读取了数据...", n, v)
			rwm001.RUnlock()
		}(i)
	}

	for i := 6; i <= 10; i++ {
		//write
		go func(j int) {
			defer wg001.Done()
			rand.Seed(time.Now().UnixNano())
			rwm001.Lock()
			fmt.Println("写操作", j, "即将写入数据...")
			randNum := rand.Intn(100) + 1 //[1,100]
			n = randNum
			fmt.Println("写操作", j, "已经结束，写入了", randNum, n)
			rwm001.Unlock()
		}(i)
	}

	wg001.Wait()
}

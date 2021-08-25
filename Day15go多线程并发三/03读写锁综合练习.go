package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var srw = sync.RWMutex{}
var num int

func main() {
	//为了提高数据访问效率又兼顾安全性，建议使用读写锁完成同步
	/*
		读写锁：
			1.写锁定/写解锁：Lock(),Unlock()
			2.读锁定/读解锁：RLock(),RUnlock()
			3.尽量不要将读写锁/互斥锁与channel混用-----隐形死锁
	*/

	wg.Add(10)
	for i := 1; i <= 5; i++ {
		go read(i)
	}
	for i := 1; i <= 5; i++ {
		go write(i)
	}
	wg.Wait()
	fmt.Println()
	fmt.Println("main...over...")
}
func read(out int) {
	srw.RLock()
	data := num
	fmt.Printf("\t读go程...第%d个读go程-------读取数据为:%d\n", out, data)
	srw.RUnlock()
	wg.Done()

}
func write(in int) {

	srw.Lock()
	num = in * 1000
	fmt.Printf("写go程...写入第%d个go程-------写入的数据为:%d！\n", in, num)
	srw.Unlock()
	wg.Done()

}

package main

import (
	"fmt"
	"sync"
)

type myMap struct {
	map1 map[int]string
	rwm  sync.RWMutex
}

// 将键值对存入map中

func (m *myMap) Set(k int, v string) {
	m.rwm.Lock()
	m.map1[k] = v
	m.rwm.Unlock()
}

//从map中读取数据

func (m *myMap) Get(k int) string {
	m.rwm.Lock()
	defer m.rwm.Unlock()
	return m.map1[k]
}

func main() {
	/*
		并发读写map：

	*/
	fmt.Println("-------------------------------------01.测试-并发读写map------------------------------------------")
	{
		//var rwm99 sync.RWMutex
		//map1 := make(map[int]string)
		//for i := 1; i <= 20; i++ {
		//	go func(j int) {
		//		rwm99.Lock()
		//		map1[j] = fmt.Sprint("数据", j)
		//		rwm99.Unlock()
		//	}(i)
		//}
		//time.Sleep(2 * time.Second)
		//fmt.Println(map1)
	}
	fmt.Println("-------------------------------------02.测试-带读写锁的并发读写map----------------------------------")
	m1 := make(map[int]string)
	m2 := myMap{
		map1: m1,
		rwm:  sync.RWMutex{},
	}
	var myWg sync.WaitGroup
	myWg.Add(30)
	for i := 1; i <= 30; i++ { //map不允许并发写
		go func(j int) {
			m2.Set(j, fmt.Sprintf("数据的值为：%d", j*100))
			myWg.Done()
		}(i)
	}
	myWg.Wait()
	fmt.Println("main,开始读取数据...")
	for i := 1; i <= 20; i++ {
		fmt.Println(m2.Get(i))
	}
	fmt.Println("main...over...")
}

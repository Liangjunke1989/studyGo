package main

import "fmt"

func main() {
	/*
			for range:
				数组，切片，map,string,chan
					数组/切片/string------>返回index,value
								map------>返回key,value
		                       chan------>返回value
	*/
	ch1 := make(chan string)
	go sendStrData(ch1)
	for v := range ch1 { //循环的停止条件是通道的关闭,显示的调用Close()才可以         //for循环的另一种写法
		fmt.Println("从通道中读取的数据为：", v)
	}
}
func sendStrData(ch chan string) {
	for i := 1; i <= 10; i++ {
		ch <- fmt.Sprint("数据为：", i)
	}
	fmt.Println("写入数据完毕！")
	close(ch)
}

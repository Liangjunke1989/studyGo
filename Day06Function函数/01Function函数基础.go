package main

import "fmt"

func main() {
	/*
		函数基础部分：
			函数的定义
			函数的参数
			函数的返回值
			return关键字
			作用域

			可变参数
			参数的传递

			递归函数
	*/
	fmt.Println("从1-10的和为：", getSum(10))

}

//go语言不支持重载
//go语言支持一个函数有多个返回值。
func getSum(x int) int {
	sum := 0
	for i := 0; i < x; i++ {
		sum += i
	}
	return sum
}

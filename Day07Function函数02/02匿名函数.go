package main

import "fmt"

func main() {
	/*
		匿名函数：
			匿名对象/匿名变量
			匿名函数：没有名字的函数。

		通常只使用一次，定义的时候直接使用。
		匿名函数可以嵌套定义。
	*/
	func() {
		fmt.Println("我是匿名函数！！！")
	}() //匿名函数直接调用

	func(a, b int) {
		fmt.Println("我是匿名函数2！！！")
	}(1, 2) //带参数的匿名函数直接调用

	i := func(a, b, c int) int {
		sum := 0
		sum = a + b + c
		fmt.Println("我是匿名函数3,返回值是：", sum)
		return sum
	}(1, 2, 3) //带参数的带返回值的匿名函数直接调用
	fmt.Println(i)

	j := func() {
		fmt.Println("我也是匿名函数4！！！")
	}
	fmt.Println(j) //j为委托，即j为匿名函数类型的变量
	j()            //匿名函数的掉用，即委托实例的调用。

}

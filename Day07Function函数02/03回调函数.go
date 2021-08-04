package main

import "fmt"

func main() {
	/*
		回调函数：callback
			1.将一个函数 作为另一个函数的参数使用。
			2.将函数作为返回值。-----闭包结构closure
	*/
	/*
		根据go语言函数的数据类型特点，可以将函数作为另一个函数的参数。
		func1()
		func2()
		将func1函数作为func2函数的参数：
			func2函数：高阶函数 （接收一个函数作为参数的函数）
			func1函数：回调函数 （作为另一个函数的参数的函数）
	*/
	fmt.Println("------------------------------1. 方法作为参数使用，方法的内存地址-----------------------------")
	fmt.Println(operation1)
	operation1(4, 5, add) //内存地址

	fmt.Println("------------------------------2.1 方法作为参数使用加法--------------------------------------")
	fmt.Println(operation2)
	fmt.Println(operation2(4, 5, add))

	fmt.Println("------------------------------2.2 方法作为参数使用减法--------------------------------------")
	fmt.Println(operation2)
	fmt.Println(operation2(4, 5, sub))

	fmt.Println("------------------------------2.3 方法作为参数使用乘法，匿名函数的使用01------------------------")
	//匿名函数
	result00 := func(a, b int) int {
		//fmt.Println("--执行了sub回调函数")
		return a * b
	}
	fmt.Println(operation2) //指向了同一块内存地址
	fmt.Println(operation2(4, 5, result00))

	fmt.Println("------------------------------2.4 方法作为参数使用除法，匿名函数的使用02----------------------------------------")
	//var01:=operation2(78,5, func(a int, b int) int {
	//	if b==0 {
	//		return 0  //内存会溢出
	//	}
	//	return a/b
	//})
	//fmt.Println(var01)
	fmt.Println(operation2(100, 5, func(a int, b int) int { //调用operation2方法，往里面传参。得到结果是int值
		if b == 0 {
			fmt.Println("除数不能为0，内存会溢出")
			return 0 //内存会溢出
		}
		return a / b
	}))
	fmt.Println(operation2) //函数为引用类型，内存地址是不同的启动会改变的
	/*
		当operation2函数被调用的时候，匿名函数才会被调用。
		所以匿名函数称之为回调函数，operation2函数
	*/
}

//函数的类型: func(int,int)int
func add(a, b int) int {
	fmt.Println("--执行了add回调函数")
	return a + b
}
func sub(a, b int) int {
	fmt.Println("--执行了sub回调函数")
	return a - b
}

//函数的类型：func(int,int,func(int,int)int)
func operation1(m, n int, opera func(int, int) int) { //01 函数作为参数使用，go语言中函数类型不需要额外声明，即声明委托
	fmt.Println("--执行了operation1函数")
	i := opera(m, n)
	fmt.Println("operation1开辟了内存地址：", operation1)
	fmt.Println(i)
}

//函数的类型：func(int,int,func(int,int)int)int
func operation2(m, n int, opera func(int, int) int) int {
	fmt.Println("--执行了operation2函数")
	fmt.Println("operation2开辟了内存地址：", operation2)
	return opera(m, n)
}

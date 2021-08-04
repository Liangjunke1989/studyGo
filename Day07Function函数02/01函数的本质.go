package main

import "fmt"

func main() {
	/*
		function  属于引用类型
		struct 	  属于值类型
	*/
	/*
		函数的本质：
			是一个引用类型，定义了一块内存地址。
			函数可以作为变量使用，也可以作为其他函数的参数使用。

	*/
	fmt.Println("----------------------1.函数的类型---------------------------------------")
	//函数是引用类型的变量
	fmt.Printf("%T,%p\n", getNum, getNum)
	fmt.Println(getNum)
	fmt.Println("----------------------2.直接定义一个函数类型的变量--------------------------")
	//函数可以作为一种特殊类型的变量
	var mFunc func(int) // func(int)类似于c#中的委托，mFun类似于c#中的事件。
	fmt.Printf("%T,%p\n", mFunc, mFunc)
	fmt.Println(mFunc) //  没有new，内存堆里没有
	mFunc = getNum     //  函数赋值mFunc; 如果getNum()是将函数的值赋值给mFunc。
}
func getNum(i int) {
	fmt.Println(i)
}

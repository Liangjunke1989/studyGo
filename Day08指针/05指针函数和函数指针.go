package main

import "fmt"

func main() {
	/*
		函数指针：
			是一个指针，指向了一个函数的指针。
			在go语言中，function，默认看作一个指针，没有*标识而已。
			引用类型本身就是指针。  （指针---内存栈对应的内存堆。   指针有点类似于内存栈，对应的内存堆上存数据/值。）
		指针函数：
			是一个函数。

		使用指针可以节省一部分内存。
	*/
	//fun11()
	//fmt.Println(fun22())
	//fmt.Println(fun33())
	arr1 := fun22()
	arr2 := fun22()
	p3 := fun33()
	p4 := fun33()
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(p3)
	fmt.Println(p4)
}

//普通函数
func fun11() {
	fmt.Println("fun1函数！！")
}

//普通函数
func fun22() [4]int {
	arr := [4]int{1, 2, 3, 4}
	return arr
}

//指针函数
func fun33() *[4]int {
	arr := [4]int{1, 2, 3, 4}
	return &arr // *[4]int的指针类型
}

package main

import "fmt"

func main() {
	/*
		指针作为参数：
			传递参数：值传递（数据的副本）,引用传递（内存地址）
		指针一般仅针对值类型进行操作。
	*/
	fmt.Println("----------------------------------------------------------")
	a := 10
	fmt.Println("函数调用之前的a:", a)
	fun1(a)
	fmt.Println("函数调用之后的a:", a)

	fmt.Println("--------------------1.将指针作为参数使用---------------------------")
	fun2(&a) //将指针作为参数使用，传递的是a的地址，相当于传递的a的地址。

	arr1 := [4]int{1, 2, 3, 4}
	func3(arr1)
	func4(&arr1)
	fmt.Println(arr1)
}

func fun1(num int) {
	fmt.Println("函数中的num：", num)
	num = 100
	fmt.Println("函数中修改的num：", num)
}
func fun2(p1 *int) {
	fmt.Println("函数fun2中的p1：", *p1)
	*p1 = 100
	fmt.Println("函数fun2中修改的p1中的值：", *p1)
}
func func3(arr [4]int) {
	fmt.Println("函数func3中的数组是：", arr)
	arr[0] = 100
	fmt.Println("函数func3中修改后的数组是：", arr)
}
func func4(p2 *[4]int) {
	fmt.Println("函数func4中的数组是：", *p2)
	p2[0] = 100
	fmt.Println("函数func4中修改后的数组是：", *p2)
}

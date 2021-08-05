package main

import "fmt"

func main() {
	/*
		数组指针：
			首先是一个指针，一个数组的地址。
			存储的是数组的地址。
		指针数组：

	*/
	//1.数组的创建
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println(arr1)
	//2.创建一个指针，存储该数组的地址。         -->数组指针   本质：指针
	var p1 *[4]int
	p1 = &arr1
	fmt.Println(p1)
	fmt.Printf("%p\n", p1) //指针的地址
	//3.根据数组指针，操作数组
	//(*p1)[0]=100
	p1[0] = 100
	fmt.Println(arr1)
	//4.指针数组
	a := 1
	b := 2
	c := 3
	d := 4
	arr2 := [4]int{a, b, c, d}
	fmt.Println(arr2)
	arr3 := [4]*int{&a, &b, &c, &d} //-->指针数组     本质：数组
	fmt.Println(arr3)
	arr2[0] = 100
	fmt.Println(a) //a的值没有发生改变
	for _, v := range arr3 {
		fmt.Println(v)
	}
	fmt.Println("--------------")
	*arr3[0] = 200
	fmt.Println(arr3[0]) //内存地址没变。该数值不影响地址。
}

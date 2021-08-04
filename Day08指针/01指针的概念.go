package main

import "fmt"

func main() {
	/*
			指针：
				概念：指针是存储另一个变量的内存地址的变量。
		        指针往往针对值类型来说。
				1.指针的类型：*int,*float,*string,*array,*struct
				2.指针中存储的数据的类型：int,float,string ,array,struct
				3.指针中存储的数据的地址：(指针中存储的数值)    &int,&float,&string,&array,&struct
				4.指针自己的地址：& *int,& *string,& *array,& *struct
	*/
	//1.1定义一个int类型的变量
	a := 10
	fmt.Printf("%d,%T,%p\n", a, a, &a) //值类型的内存地址 0xc0000ae008 每次系统随机分配，可能变
	//1.2定义一个string类型的变量
	str := "中国人"
	fmt.Printf("%s,%T,%p\n", str, str, &str)
	//2.创建一个指针变量，用于存储变量a的地址
	//指针的声明：*T是指针变量的类型，它指向T类型的值。
	var p1 *int // 空指针 nil pointer
	p1 = &a
	fmt.Println("p1中数值是：", p1)
	fmt.Println("p1指针存储的变量地址的数据是：", *p1)
	fmt.Printf("p1自己的内存地址是：%p\n", &p1) //指针也是值类型？
	//3.操作变量，更改数值
	a = 100
	fmt.Println(a)
	fmt.Printf("%p\n", &a)
	//4.操作指针改变数值
	*p1 = 200 //更改指针存储的变量 地址 的数据
	fmt.Println(a)
	//5.指针的指针
	var p2 **int //**指针的指针，即指针的地址。
	p2 = &p1     //将p1的内存地址赋值给p2
	fmt.Printf("p1的变量地址是：%p\n", p1)
	fmt.Println("p1的变量地址的数据是：", *p1)
	fmt.Println("p1指针的自身内存地址是：", p2)
	fmt.Printf("p2指针自己的内存地址是：%p\n", &p2)    //  &值对应的内存地址
	fmt.Println("p2指针自己的内存地址 存储的数据是：", *p2) // *地址所对应的数据
	fmt.Println("p2指针自己的内存地址 存储的数据的数据是：", **p2)
	fmt.Printf("p1的类型是：%T\n", p1)
	fmt.Printf("p2的类型是：%T\n", p2)
}

package main

import "fmt"

type cat struct {
	name string
	age  int
}

func changeName1(c cat) {
	c.name = "花花"
}
func changeName2(c *cat) {
	c.name = "默默"
}
func main() {
	/*
		作为参数：
			结构体对象作为参数？
			结构体对象的指针作为参数？
	*/
	c1 := cat{"Tom", 1}
	fmt.Println("-------cat原结构数据：---------")
	fmt.Println(c1)
	fmt.Println("-------传入cat结构体对象的参数：---------")
	changeName1(c1) //将c1结构体对象为参数传递，值传递。c1的值没有被改变
	fmt.Println(c1.name)
	fmt.Println("-------传入cat结构体对象指针的参数：---------")
	changeName2(&c1) //将c1结构体对象的指针（内存地址）为参数传递，引用传递。c1的值发生了改变。
	fmt.Println(c1.name)
}

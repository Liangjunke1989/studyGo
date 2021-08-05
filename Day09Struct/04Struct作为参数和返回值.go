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
func getCat1() cat {
	c := cat{"露娜", 11}
	return c
}
func getCat2() *cat {
	c := cat{"加菲猫", 11}
	return &c //取c的内存地址
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
	changeName1(c1) //将c1结构体对象为参数传递，值传递。c1的值没有被改变。
	fmt.Println(c1.name)
	fmt.Println("-------传入cat结构体对象指针的参数：---------")
	changeName2(&c1) //将c1结构体对象的指针（内存地址）为参数传递，引用传递。c1的值发生了改变。  指针传递的是地址。
	fmt.Println(c1.name)
	/*
		作为返回值：
			结构体对象作为返回值？
			结构体对象的指针作为返回值？
	*/
	c2 := getCat1()
	c3 := getCat1()
	fmt.Println(c2, c3)
	c2.name = "jerry"
	fmt.Println(c2, c3)
	fmt.Println("------------------------------------------------------------------------------------------")
	c4 := getCat2()
	c5 := getCat2()
	fmt.Println(c4, c5)
	fmt.Println("------------------------------------------------------------------------------------------")
	c4.name = "花猫"
	fmt.Println(c4, c5) //c对象在方法内，重新创建

}

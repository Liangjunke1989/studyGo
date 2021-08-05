package main

import "fmt"

//1.定义一个类，就是定义一个结构体

type Person struct {
	name    string
	age     int
	sex     string
	address string
}

func main() {
	//第一种定义
	var person01 Person //根据类实例化对象（已经创建除了一个对象，不会报错。结构本质是值类型的数据。默认值是0）
	person01.name = "mm"
	fmt.Println(person01)

	//第二种定义
	person02 := Person{} //类似与new
	person02.name = "kk"
	person02.age = 12
	fmt.Printf("%p\n", &person02)
	fmt.Println(person02) //输出的是默认值

	//第三种定义
	person03 := Person{ //类似栈上创建
		name:    "dk",
		age:     29,
		sex:     "女",
		address: "八面神",
	}
	fmt.Println(person03)

	//第四种定义
	person04 := Person{"ljk", 32, "男", "万和城"} //注意赋值顺序
	fmt.Println(person04)

	//第五种定义
	//使用内置函数new(),go语言种专门用于创建某种类型的指针的函数。
	person05 := new(Person) //类似于堆上创建，是指针类型
	person05.name = "junke"
	fmt.Println(person05)      //&{junke 0  }   //打印的是Person类型的地址。即是Person类型的指针
	fmt.Printf("%p", person05) //指针

}

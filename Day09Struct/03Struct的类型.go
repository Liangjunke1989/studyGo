package main

import (
	"fmt"
)

type dog struct {
	color string
	name  string
	age   int
}

func main() {
	/*
		struct的数据类型：
			值类型
			默认是深拷贝
			如果使用浅拷贝，需要使用指针地址。节省内存。
	*/
	d1 := dog{"黄色", "大黄", 2}
	d2 := d1
	d2.name = "二黄"
	fmt.Println(d1.name)
	fmt.Println(d2.name)

	//实现结构体的浅拷贝
	d3 := new(dog)
	d3.color = "yellow"
	d3.name = "中华田园犬"
	d3.age = 12
	fmt.Println(d3)
	d4 := d3
	d4.name = "秋天犬"
	fmt.Println(d3.name, d4.name)
}

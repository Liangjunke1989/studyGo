package main

import "fmt"

type student struct {
	name string
	age  int
	int
	string
}

func main() {
	/*
		var 变量名 数据类型
		const 变量名 数据类型=赋值
		func 函数名 函数体{}
		type 结构体名 struct{}
		type:用于定义数据类型
	*/
	/*
		匿名函数：
			没有名字的函数，随着定义的时候直接进行调用。
		匿名结构体：
			 匿名类
		匿名字段：
			一个结构体的字段没有名字。
			如果一个字段没有名字，那么默认使用类型作为字段名。
			匿名字段的类型不能冲突。
	*/
	//匿名结构体
	p1 := struct {
		name string
		age  int
	}{"kk", 23} //new
	fmt.Println(p1)
	//匿名字段
	stu := student{"ljk", 32, 24, "1"}
	fmt.Println(stu)

	/*
		练习一：
			设计几个函数，分别接收结构体，结构体指针作为参数，返回值是结构体，结构体指针。
	*/

	/*
		练习二：
			设计一个匿名结构体，并创建对象。
			设计一个结构体中有匿名字段，创建对象使用。
	*/
}

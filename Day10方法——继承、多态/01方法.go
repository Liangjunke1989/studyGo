package main

import "fmt"

//1.1 定义一个工人类

type Worker struct {
	id   int
	name string
	age  int
	sex  string
}

//1.2 定义一个猫类

type Cat struct {
	name  string
	color string
	age   int
}

//2. 定义类的行为：方法
func (w Worker) work() {
	fmt.Println(w.name, "在工作！！！") //receiver是Worker类型的对象
}

func (c Cat) work() {
	fmt.Println(c.name, "在抓老鼠！！！") //receiver是Cat类型的对象
}
func main() {
	/*
		定义类：
			字段：field
				变量
				静态属性
			方法：method
				函数
				动态行为
	*/
	/*
		函数和方法：
			函数：function
				一段具有独立功能的代码，可以被反复多次调用。
			方法：method
				一个类的行为功能，只有该类的对象才能调用。
				同函数类似，但需要接收者。（也叫是调用者）
			方法和函数的不同：
				a.意义：
					方法：某个类别的行为功能。需要接收者调用。
					函数：一段独立功能的代码。可以直接调用。
				b.语法：
					函数：命名不能冲突。
					方法：方法名可以一致，只要接收者不同即可。 t Type
				func(t Type)methodName(parameter list){
				}
	*/
	/*
		go-->字段定义在struct里
			 方法定义在struct外
	*/

	//oop面向对象
	var worker01 = Worker{01, "小李", 18, "男"}
	worker01.work()
	var cat01 = Cat{
		name:  "小花",
		color: "黄色",
		age:   3,
	}
	cat01.work()
	car02 := &cat01
	car02.work()
}

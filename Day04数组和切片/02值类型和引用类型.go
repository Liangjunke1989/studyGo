package main

import "fmt"

func main() {
	/*
		值类型：
			存储的值的副本，
			开辟新空间，值和变量都在栈中
			将数据传递给其他变量，传递的是数据的副本。
			基本类型中的int, float,bool
			值类型在线程栈中(stack)
		引用类型：
			存储的数据的内存地址
			引用类型在托管堆上
			变量——>栈
			对象——>堆（heap）
	*/
	fmt.Println("-----------------------------------------------------------------")
	arrNum1 := [...]int{1, 2, 3, 4, 5}
	arrNum2 := arrNum1 //开辟新的内存空间
	arrNum2[3] = 2
	fmt.Println(arrNum1)
	fmt.Println(arrNum2)
}

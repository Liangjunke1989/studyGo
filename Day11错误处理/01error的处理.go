package main

import (
	"errors"
	"fmt"
)

func main() {
	/*
		go语言的异常处理：
			panic 恐慌
			recover 恢复
		go语言中错误也是一种类型。error类型

		error:
			go内置的error数据类型
			表示错误
			内置error接口类型
				Error()string
		使用go语言提供好的包：
			errors包下的函数New(),创建一个error对象。
	*/
	fmt.Println("----------------1. 创建一个error类型的对象的方法一-------------------------------------------")
	//1.创建一个error对象
	err1 := errors.New("主人，这里出现错误了！！！")
	fmt.Println(err1.Error())
	fmt.Printf("%T", err1)
	fmt.Println()
	fmt.Println("----------------1.1设计一个函数：验证年龄是否合法。如果为负数的时候，返回error--------------------")
	/*
		设计一个函数：验证年龄是否合法。如果为负数的时候，返回error。
	*/
	testFunc()

	fmt.Println("----------------2.创建一个error类型的对象的方法二--------------------------------------------")
	//age99:=99
	//err22 := fmt.Errorf("您给定的年龄是：%d,不合法",age99)
	//fmt.Println(err22)
	err3 := CheckAge(-22)
	if err3 != nil {
		fmt.Println(err3.Error())
		return //return 直接返回，以下都不会执行--------
	}
}
func CheckAge(age int) error {
	if age < 0 {
		//返回error对象
		//方法一：
		//return errors.New("年龄不合法！")
		//方法二：
		return fmt.Errorf("您给定的年龄是：%d,不合法", age)
	}
	fmt.Println("年龄正常输入，没有错误返回！")
	return nil
}
func testFunc() {
	err2 := CheckAge(-2)
	//fmt.Println(err2.Error())
	//err3 := CheckAge(2)
	//fmt.Println(err3.Error())         //nil.error 会发生恐慌错误，panic
	if err2 != nil {
		fmt.Println(err2.Error())
		fmt.Println()
		return //return 直接返回，以下都不会执行--------
	}
	fmt.Println("继续执行！！")
	fmt.Println()
}

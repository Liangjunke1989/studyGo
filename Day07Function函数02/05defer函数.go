package main

import "fmt"

func main() {
	/*
		defer函数：
			一个函数的执行被延迟了。
			多个函数被defer，先延迟的最后执行，后延迟的先执行。
			类似于栈stack的数据结构：先进后出。
		用途：
			defer 对象的关闭处理。
	*/
	fun1("ljk")
	defer fmt.Println("go!go!go!")
	defer fun1("dk") //被延迟，延迟到外层函数被执行完。
	fmt.Println("mmmmmmmmm")

}
func fun1(s string) {
	fmt.Println(s)
}

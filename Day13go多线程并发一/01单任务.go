package main

import "fmt"

func main() {
	/*
		单任务：
			也叫串行。
	*/
	fmt.Println("主函数main()调用......")
	fun1()
	fmt.Println("主函数的打印......")
	fun2()
	fun3()
}
func fun1() {
	fmt.Println("HelloWorld......")
}
func fun2() {
	for i := 1; i <= 5; i++ {
		fmt.Println("fun2()函数调用，i:", i)
	}
}
func fun3() {
	fmt.Println("fun3()函数调用......")
	fun2()
}

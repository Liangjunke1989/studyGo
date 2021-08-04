package main

import "fmt"

func main() {
	//a := 13
	//if a>10 {
	//	fmt.Println(a)
	//}

	//if a > 20 {
	//	fmt.Println(a) //满足了，就不执行下面
	//} else if a > 14 {
	//	fmt.Println(-a)
	//} else if a > 10 {
	//	return //跳出当前流程，下面的不再执行了！！！
	//}

	fmt.Println("—————————————————go if的初始化语句—————————————————————")
	//if 初始化语句;条件{
	//
	//}
	if b := 100; b > 90 { //b变量的作用域在条件内，条件结束，b变量销毁。
		fmt.Println("优秀！！")
	}
	//fmt.Println(b) //b的作用域是if条件语句内
	fmt.Println("main over!")
}

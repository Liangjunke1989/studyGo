//声明包 （程序在哪个包下,  Main特殊，其他的写"Day01环境配置和打印输出HelloWorld"）
package main

//导入包  fmt包是输入输出包
import (
	. "fmt"
)

var name = "kk"

//主函数，程序执行的入口
func main() {
	//使用fmt包下的打印输出功能
	Println("HelloWorld!!!")
	Println("ljk")
	Println("dk")
	Println("kooks")
	Printf("\n%q", name)
}

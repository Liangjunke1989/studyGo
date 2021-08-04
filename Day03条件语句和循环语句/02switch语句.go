package main

import "fmt"

func main() {
	var var1 = "春春"
	switch var1 { //switch后的变量类型必须和case后的数值类型匹配
	case "春":
		fmt.Println("这是第一季度！")
	case "夏":
		fmt.Println("这是第二季度！")
	case "秋":
		fmt.Println("这是第三季度！")
	case "冬":
		fmt.Println("这是第四季度！")
	default: //有无都可以
		fmt.Println("数据有误！")
	}
}

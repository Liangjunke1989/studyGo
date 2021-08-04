package main

import (
	"fmt"
)

func main() {
	/*
		双层for循环嵌套：
			外层i循环，值变化一次，执行一次循环体（里层循环完整的执行）。
			循环体是整个里层循环，那么j循环要完成的执行一边。
	*/
	{
		for i := 0; i < 5; i++ { //用于控制行数
			for j := 0; j < 5; j++ { //用于控制*的个数
				fmt.Print("*")
			}
			fmt.Println()
		}
	}
	fmt.Println("----------------------练习一：每行递增*--------------------------------------")
	{
		for m := 1; m <= 5; m++ {
			for n := 0; n < m; n++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	}

	fmt.Println("--------------------练习二：每行递增*，每行递减空格-------------------------------")
	{
		for m := 1; m <= 5; m++ { //用来表示行数
			//for k := 0; k<4-m; k++ {
			//	fmt.Print(" ")
			//}
			for k := 4; k >= m; k-- { //每行递减空格
				fmt.Print(" ")
			}
			for j := 0; j < m; j++ { //每行递增*
				fmt.Print("*")
			}
			fmt.Println()
		}
	}

	fmt.Println("---------------------平行四边形-------------------------------------")
	{
		for m := 0; m < 5; m++ { //用来表示行数
			for k := 0; k < 4-m; k++ {
				fmt.Print(" ")
			}
			for j := 0; j < 5; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	}
}

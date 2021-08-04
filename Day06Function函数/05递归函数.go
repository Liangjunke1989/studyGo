package main

import "fmt"

func main() {
	/*
		递归函数：
		  一个函数自己调用自己，就叫做递归函数。
		  递归函数要有出口，逐渐的向出口靠近。
	*/
	/*
		求1-5的和：
			getSum(5)=getSum(4)+5
			getSum(4)=getSum(3)+4
			getSum(3)=getSum(2)+3
			getSum(2)=getSum(1)+2
			getSum(1)=1
	*/

	fmt.Println(getSum1(5))
	fmt.Println(getSum2(5))
}

//方法一：for循环
func getSum1(x int) int {
	sum := 0
	for i := 1; i <= x; i++ {
		if i <= x {
			sum += i //正序增加
		}
	}
	return sum
}

//方法二：递归函数
func getSum2(x int) int {
	fmt.Printf("--------第%d次-----------\n", x)
	if x == 1 {
		return 1 // getSum(1)=1            //关键是有出口，不然死循环
	}
	return getSum2(x-1) + x //倒叙增加！！！     实质：慢慢向出口靠近的思想！！！
}

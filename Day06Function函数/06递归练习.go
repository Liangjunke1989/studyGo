package main

import "fmt"

func main() {
	/*
		递归练习：
			斐波那契数列：兔子生兔子。
						第1，2两项数值都为1，从第3项开始，是前两项之和。
			使用递归算法，求第n月兔子之和
	*/
	fmt.Println(getFibonacci(4))
}

//练习一：阶乘
func getFactorial(n int) int {
	if n == 1 {
		return 1
	}
	return getFactorial(n-1) * n
}

//练习二：斐波那契数列 (默认n>3）
func getFibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return getFibonacci(n-1) + getFibonacci(n-2)
}

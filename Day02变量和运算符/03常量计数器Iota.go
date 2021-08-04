package main

import "fmt"

func main() {
	/*
		    iota:
			作为个特殊的常量值，默认为0，go内置的计数器。
			每当有一个const，iota的值初始化为0，每当增加一个常量值，iota的值就累加1。
			iota清零:再次遇到const定义常量组的时候，iota的值又初始化为0。
	*/
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Println(a, b, c)
	const (
		m = iota
		n = iota
	)
	fmt.Println(m, n)
	const (
		l = iota
		p //p=iota
		q //q=iota
	)
	fmt.Println(l, p, q)
	const (
		MALE = iota
		FEMALE
		UNKNOWN
	)
	fmt.Println(MALE, FEMALE, UNKNOWN)
	const (
		A = iota
		B
		C
		D = "HAHA"
		E
		F = 100
		G
		H = iota
		I
	)
	const (
		J = iota
	)
	fmt.Println(A, B, C, D, E, F, G, H, I)
	fmt.Println(J)
}

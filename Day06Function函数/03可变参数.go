package main

import "fmt"

func main() {
	/*
		可变参数：
			指一个函数的参数，类型确定，但个数不定(0-多个)。
			如果有可变参，还有其他参数的话，可变参数放在最后。
			一个函数的列表中，最多只能一个可变参。
	*/
	//求n个整数的和。
	fmt.Println(getAdd())
	fmt.Println(getAdd(1, 2, 3))
	fmt.Println(getAdd(1, 2, 3, 4, 5, 6))
	//切片
	s1 := []int{5, 6, 7, 8}
	fmt.Println(getAdd(s1...))

}
func getAdd(params ...int) int {
	//fmt.Printf("%T",params)// []int 切片
	sum := 0
	for i := 0; i < len(params); i++ {
		sum += params[i]
	}
	return sum
}

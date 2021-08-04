package main

import "fmt"

func main() {
	/*
		因为切片是引用类型的数据，直接拷贝的是地址。
		浅拷贝/深拷贝
		浅拷贝：拷贝的是数据的地址。
			导致多个变量指向同一块内存。
			引用类型的数据默认都是浅拷贝。slice/map
		深拷贝：拷贝的是数据本身。
			值类型的数据，默认都是深拷贝。array/int/float/bool/string  /struct结构体也是值类型
	*/
	fmt.Println("----------------------------01手动创建实现对切片的深拷贝----------------------------------")
	{
		s1 := [...]int{1, 2, 3, 4}
		s2 := make([]int, 4)
		//fmt.Printf("%p\n",s1)      //数组为值类型不需要开辟内存地址
		fmt.Printf("%p\n", s2)
		for key, _ := range s2 {
			s2[key] = s1[key]
		}
		s2[1] = 98
		fmt.Println(s1, s2)
		//fmt.Printf("%p\n",s1)
		fmt.Printf("%p\n", s2)
	}

	fmt.Println("----------------------------02使用cope实现对切片的深拷贝----------------------------------")
	{
		s1 := []int{1, 2, 3}
		s2 := []int{7, 8, 9}
		//copy(s2,s1)             //dst目标地址,src源地址
		//fmt.Println(s1,s2)
		fmt.Println()
		copy(s2, s1[:2])
		fmt.Println(s2)
	}
}

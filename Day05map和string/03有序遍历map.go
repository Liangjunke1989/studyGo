package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("-------------------------1.设置源数据-----------------------------------")
	var map1 = make(map[int]string)
	{
		{
			map1[5] = "金角大王"
			map1[3] = "白骨精"
			map1[1] = "红孩儿"
			map1[2] = "小钻风"
		}
		//for _,v :=range map1{
		//	fmt.Println(v)
		//}
	}

	fmt.Println("-------------------------2.按key升序遍历-----------------------------------")
	{
		//fmt.Println(map1)
		keys := make([]int, 0, len(map1))
		for k, _ := range map1 {
			keys = append(keys, k)
		}
		//fmt.Println(keys)
		sort.Ints(keys) //升序
		//fmt.Println(keys)
		for _, kk := range keys { //keys数组中key为index，value为数组中的数
			fmt.Println(kk, map1[kk])
		}

	}
}

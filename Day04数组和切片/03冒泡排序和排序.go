package main

import "fmt"

func main() {
	/*
		数组的排序：
			------算法和数据结构------
			排序也是一种算法。（升序和降序）
				8种排序：
					冒泡排序
					插入排序
					选择排序
					堆排序
					快速排序...
			冒泡排序的思想：比较相邻的两个数，小的放在前面，大的放在后面。
				第一轮排序 找到第一个最大值的位置，
					第二轮排序，找到第二个最大值的位置，
						第三轮排序，找到第三个最大值的位置，
							第四轮排序，找到第四个最大值的位置...
			5个数，4轮可以排序完成。
	*/
	var arrNum = [...]int{111, 22, 3, 44, 55, 66, 7}
	for i := 0; i < len(arrNum); i++ {
		for j := 0; j < len(arrNum)-1; j++ {
			if arrNum[j] > arrNum[j+1] {
				arrNum[j+1], arrNum[j] = arrNum[j], arrNum[j+1] //go语言不用temp，直接可以换位赋值
			}
		}
	}
	fmt.Println(arrNum)

}

package main

import (
	"fmt"
)

func main() {
	/*
		数组：
			存储一组相同数据类型的数据结构。
		数据类型：
			复杂数据类型（数组、切片、map、结构体、指针、接口、函数...）
		数据结构：
			数组(Array)
			栈( Stack)
			队列(Queue)
			链表( Linked List)
			树( Tree)
			图(Graph)
			堆(Heap)
			散列表(Hash)
		语法：
			var 数组名 [长度]  数据类型
	*/
	fmt.Println("-------------------------数组的写法：-----------------------------")
	{
		//var num01 int =100
		//fmt.Println(num01)
		//写法一：
		var arrNum1 [5]int
		{
			arrNum1[0] = 1
			arrNum1[1] = 2
			arrNum1[2] = 3
			arrNum1[3] = 4
			arrNum1[4] = 5
		}
		//写法二：
		var arrNum2 = [5]int{1, 2, 3, 4, 5}
		//写法三：
		var arrNum3 = [5]int{0: 22, 1: 1, 4: 2}
		//写法四：
		arrNum4 := [4]int{6, 7, 8, 9}
		//写法五：
		arrNum5 := [...]int{6, 7, 8, 9, 10, 11, 12} //使用"..."进行长度推断，根据花括号的数据，来推断数据的长度
		//写法六：
		arrNum6 := [...]int{1: 2, 9: 0} //以最后的下标值为准
		//fmt.Println(arrNum[2])
		for i := 0; i < len(arrNum1); i++ {
			fmt.Println(arrNum1[i])
		}
		//内置函数 len/cap
		m := len(arrNum6) //长度：容器中实际上存储的数据量
		n := cap(arrNum6) //容量：容器中能够存储的最大的数量
		fmt.Println("数组的长度为：", m)
		fmt.Println("数组的容量为：", n)
		fmt.Println(arrNum2)
		fmt.Println(arrNum3[4])
		fmt.Println(arrNum3)
		fmt.Println(arrNum4)
		fmt.Println(arrNum5)
		fmt.Println(arrNum6)
	}

	fmt.Println("---------------------使用range访问数组----------------------------")
	{
		/*
			    		range: 范围
			    		for rangel类似于foreach
						使用range访问数组
		*/
		arrNum01 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		sum1 := 0
		sum2 := 0
		//写法一：
		for index, _ := range arrNum01 {
			//fmt.Println("下标是：",index,"数值是：",value)
			//fmt.Println(arrNum01[index])
			sum1 += arrNum01[index]
		}
		fmt.Println("1-10相加的和为：", sum1)

		//写法二：
		for _, value := range arrNum01 {
			sum2 += value
		}
		fmt.Println("1-10相加的和为：", sum2)
	}
}

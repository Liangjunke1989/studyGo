package main

import "fmt"

func main() {
	/*
		切片：
			切片是对数组功能的扩充。
			变长数组，动态数组。长度是不固定的。
			切片本身没有数据，它们只是对现有数组的引用。
			在[]中不用设定值。
			长度和容量不一样相同。

			是一个引用类型的容器，指向了一个底层数组。
			更像一个结构体。
	*/

	fmt.Println("---------------------01切片的声明-----------------------")
	{
		//写法一：
		var sliceNum1 []int
		fmt.Println(sliceNum1)
		//写法二：
		var sliceNum2 = []int{1, 2, 3, 4}
		fmt.Println(sliceNum2)
		//写法三：
		sliceNum3 := []int{5, 6, 7, 8}
		fmt.Println(sliceNum3)
	}

	fmt.Println("---------------------02创建切片-------------------------")
	{
		//写法四：
		//make内置函数
		//make(),创建slice/map/chan，都是引用类型
		sliceNum4 := make([]int, 3, 8) //长度为3，容量为8。
		fmt.Println(sliceNum4)
		//第一个参数为类型，第二个参数为len,第三个参数为容量。
		//第三个参数省略，默认与第二个参数相同。
		sliceNum5 := make([]int, 5)
		fmt.Println(sliceNum5)
		s11 := make([]int, 0, 5)
		fmt.Println("切片的长度为：", len(s11), "\n", "切片的容量为：", cap(s11))
	}

	fmt.Println("---------------------03切片的赋值,往切片中添加元素-----------")
	{
		//与数组，赋值不同
		var sliceNum1 []int
		sliceNum1 = append(sliceNum1, 1, 2, 3)
		fmt.Println("添加后的切片为：", sliceNum1)
		s1 := make([]int, 5)
		s1 = append(s1, sliceNum1...)
		fmt.Println(s1)
		fmt.Println(sliceNum1)
	}

	fmt.Println("---------------------04根据已有的数组创建切片----------------")
	{
		//获取数组中的一部分
		//s1:=make([]int,10)
		a := [...]int{1, 4, 5, 6, 7, 8, 9, 10}
		s1 := a[0:5]    //[start,end) start,end都是下标index
		fmt.Println(s1) //1-5
		//fmt.Println("s1的长度是：",len(s1),"s1的容量是：",cap(s1))       //容量>=长度
		//fmt.Println(a[3:8])//6-10
		//fmt.Println(a[:8]) //1-10
		//fmt.Println(a[5:]) //8-10
		//fmt.Println(a[:])  //1-10
	}

	fmt.Println("---------------------05 切片扩容前，数据变化----------------------")
	{
		fmt.Println("_____原始数组a的值:_____")
		a := [...]int{1, 4, 5, 6, 7, 8, 9, 10}
		fmt.Println(a)
		fmt.Println("_____更改为切片后s1的值为：_____")
		s1 := a[0:5]
		fmt.Println(s1)
		//更改数组，切片随着更改
		fmt.Println("_____01修改数组a中的值后_____")
		a[4] = 99
		fmt.Println("切片s1为：", s1)
		//更改切片，数组也发生变化
		fmt.Println("_____02修改切片s1中的值后_____")
		s1[4] = 88
		fmt.Println("数组a为：", a) // 7改为88
	}

	fmt.Println("---------------------06 切片的扩容后，数据不随之改变--------------------------")
	{
		fmt.Println("------原始数组-----")
		a := [...]int{1, 2, 3, 4, 0}
		fmt.Println(len(a), cap(a), a) //长度5，"容量5"
		//fmt.Printf("a原始数组的地址为：%p\n",a)
		fmt.Println("------将数组转成切片后-----")
		s1 := a[:]
		//fmt.Printf("a原始数组的地址为：%p\n",a)
		fmt.Printf("s1切片的地址为：%p\n", s1)
		s1 = append(s1, 99, 88)
		fmt.Println(len(s1), cap(s1), s1) //长度7，"容量10"
		fmt.Printf("s1数组的扩容到10后的新内存地址为：%p\n", s1)
		s1 = append(s1, 11, 22)
		fmt.Println(len(s1), cap(s1), s1) //长度9，"容量10"
		fmt.Printf("s1数组的扩容到10后的新内存地址为：%p\n", s1)
		s1 = append(s1, 33, 44)
		fmt.Println(len(s1), cap(s1), s1) //长度11，"容量20"
		fmt.Printf("s1数组的扩容到20后的新内存地址为：%p\n", s1)
		//s1=append(s1,55,66,77,88,99,10,11,18,19)
		//fmt.Println(len(s1),cap(s1),s1)//长度20，"容量20"
		s1 = append(s1, 55, 66, 77, 88, 99, 10, 11, 18, 19, 221)
		fmt.Println(len(s1), cap(s1), s1) //长度21，"容量40"
		fmt.Printf("s1数组的扩容到40后的新内存地址为：%p\n", s1)

		fmt.Println("_____________更改a数据后__________________________")
		a[2] = 5566
		fmt.Println("s1切片的值为：", s1)
		fmt.Printf("s1切片的地址为：%p\n", s1)

		fmt.Println("_____________更改s1数据后________________________")
		s1[2] = 999888
		fmt.Println("原始数组a为：", a)
		//fmt.Printf("a原始数组的地址为：%p\n",a)
		fmt.Println("s1为：", s1)

		/*
				切片中添加数据原则：
					当向切片中添加数据时，如果没有超过容量，直接添加；
			                      如果超过容量，自动扩容（成倍增长）原始容量成倍增长！
					不扩容的话，内存地址不变！
		*/
	}
}

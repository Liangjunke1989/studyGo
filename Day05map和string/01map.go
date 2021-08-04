package main

import "fmt"

func main() {
	/*
		map:
			映射，集合
			长度不固定的容器。
			make()创建，是一种引用类型。
			类似于c#中的字典dir{ key1,value1,key2,value2,key3,value3,key4,value4...}
			map是一种容器，存储的是键值对。
				a.键值对是无序的。
				b.键不能重复。修改/添加，同一语法。
				c.一个键对应一个值，也叫做一个映射项。

	*/

	fmt.Println("--------------------------1.声明创建map的方式---------------------------")
	{
		//第一种：
		var map1 map[int]string //只是声明map类型的变量map1，注意此时并没有初始化
		//第一种赋值方式：
		//map1[0] = "ljk"   //需要对map1进行初始化操作，不然会报错。
		//if map1==nil{
		//  map1[0] = "ljk"
		//	//map1=make(map[int]string,2)
		//	fmt.Println(map1)
		//}

		//第二种：
		var map2 = map[int]string{ //相当于new了一个map对象，构造器，赋值给map类型的变量map1
			98:  "go",
			99:  "python",
			100: "c#",
		}
		fmt.Println(map2)
		//第二种声明一并赋值方式：
		{
			map2[0] = "ljk"
			map2[2] = "koko"
			map2[1] = "dk"
			fmt.Println(map2)
		}

		//第三种：主要的声明创建方式
		map3 := make(map[int]string, 3) //make 创建引用类型的对象，进行初始化操作
		v1, ok := map1[0]
		if ok {
			fmt.Println("map1有值：", v1)
		} else {
			fmt.Println("没有赋值成功！！！")
		}
		fmt.Println(map2[0])
		v2, yes := map1[0]
		if yes {
			fmt.Println("map2有值：", v2)
		} else {
			fmt.Println("没有赋值成功！！！")
		}
		fmt.Println(map3[0]) //未赋值的情况下，值默认值
	}

	fmt.Println("--------------------------2.获取map的值--------------------------------")
	{
		// 使用ok-idiom获取值，可以知道key/value是否存在
		map0 := make(map[int]string, 3)
		{
			map0[0] = "kk"
			map0[1] = "ss"
			map0[9] = "dd"
			map0[10] = ""
		}

		v1, ok := map0[10]
		v2, yes := map0[11]
		//fmt.Println(v1,ok)//值确实存在，为""
		if ok {
			//如果key存在，获取对应的数值。
			fmt.Println("有值，值为：", v1)
		} else {
			//如果key存在，获取的是该值类型的默认值。
			fmt.Println("没有值，默认值为：", v1)
		}

		if yes {
			//如果key存在，获取对应的数值。
			fmt.Println("有值，值为：", v2)
		} else {
			//如果key存在，获取的是该值类型的默认值。
			fmt.Println("没有值，默认值为：", v2)
		}
	}

	fmt.Println("--------------------------3.删除map的值--------------------------------")
	{
		// 使用ok-idiom获取值，可以知道key/value是否存在
		map0 := make(map[int]string, 3)
		{
			map0[0] = "kk"
			map0[1] = "ss"
			map0[9] = "dd"
			map0[10] = ""
		}
		delete(map0, 1)
		//清空map的操作
		//直接make一个新的map，旧的map会被舍弃掉。
		//创建一个新的make，go使用GC比clear更快
	}
}

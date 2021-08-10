package main

import (
	"fmt"
)

//空接口

type A interface {
}
type Animal interface {
	eat()
	sleep()
}

type Cat01 struct {
	name string
	age  int
}

type Person01 struct {
	name string
	sex  string
}

//测试函数，空接口的使用:理解为代表了任意类型
func testFunc01(a A) {

}

//将匿名空接口作为参数使用，表示该函数可以接受任意类型的数据
func testFunc02(ani []interface{}) { //匿名空接口
	// ani []interface{}与ani ...interface{}的区别？
	for i := 0; i < len(ani); i++ {
		fmt.Println("第", i+1, "个数据")
		switch ins := ani[i].(type) {
		case Cat01:
			fmt.Println("\tcat对象：", ins.name, ins.age)
		case Person01:
			fmt.Println("\tperson对象", ins.name, ins.sex)
		case int:
			fmt.Println("\tint类型", ins)
		case string:
			fmt.Println("\tstring类型", ins)
		}
	}
}
func main() {
	/*
		空接口：
			也是一个接口，但是该接口中没有任何的方法。
			所以可以将任意类型作为该接口的实现。
			理解为代表了任意类型。
			匿名空接口-----类似于object类，一切都是它的子类
	*/
	var a1 A = Cat01{"小花", 12}
	fmt.Println(a1) //{小花 12}
	var a2 A = Person01{"LJK", "男"}
	fmt.Println(a2)

	var a3 A = "sss"
	var a4 A = 120
	fmt.Println(a3, a4)
	testFunc01("123")
	map1 := make(map[string]interface{})
	map1["ljk"] = 12
	map1["kk"] = "ll"
	fmt.Println(map1) //...可变参

	//定义一个切片，存储任意类型的数据
	slice1 := make([]interface{}, 0, 10) //容量为10，长度为0
	slice1 = append(slice1, a1, a2, a3, a4, a1)
	fmt.Println(slice1)
	fmt.Println()
	fmt.Println("----------------------接口的多态练习-------------------------------")
	/*
		定义一个接口：Animal
			eat(),sleep()
		定义两个实现类：
			Cat --属性name
			Dog --属性color

		创建一个数组，存储5个Animal，设计一个函数，接受这个数组作为参数。
		打印输出数组中每个动物的eat(),sleep(),如果是猫显示名字，如果是狗显示color。
	*/
	testFunc02(slice1)

}

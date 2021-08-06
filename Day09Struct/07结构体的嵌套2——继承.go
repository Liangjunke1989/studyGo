package main

import "fmt"

func main() {
	/*
			结构体的嵌套：
				"模拟" 面向对象的类与类的两种关系：
					聚合关系：一个类作为另一个类的属性。    不匿名
		        --> 继承关系：一个类作为另一个类的子类。    匿名
	*/
	/*
		继承：
			作为面向对象的第二个特征，用于描述两个类的关系。
			一个类（子类，派生类，subClass）继承另一个类（父类，超类，基类，superClass）
			1.子类可以直接访问父类已有的属性和方法。
			2.子类可以新增自己的属性和方法。
			3.子类可以重写父类已有的方法。
		继承的意义：
			a.可以避免重复代码。         ---子类角度
			b.子类在扩展父类的功能。      ---父类角度
	*/
	/*
		重写 override 和重载 overload ?
			重载 overload是多态的一种体现，也叫静态多态。（方法名相同，但是参数列表不同。）
			重写 override是一种继承结构，子类将父类已有的方法重新实现。
		注意：go语言中有重写override，没有重载 overload
	*/

	/*
		go语言的重写override,是通过结构体的嵌套实现的。
	*/
	p1 := Person001{001, "father", 50}
	fmt.Println(p1)
	var stu1 Student001
	stu1.name = "liang"
	stu1.age = 28
	fmt.Println(stu1)
	var stu2 = Student001{Person001{002, "yiyi", 21}, "清华大学"}
	fmt.Println(stu2.id)           //提升字段，其实子类直接访问的是父类的字段。匿名字段的字段。
	fmt.Println(stu2.Person001.id) //提升字段

}

//1.定义一个父类

type Person001 struct {
	id   int
	name string
	age  int
}

//2.定义一个子类

type Student001 struct {
	//p Person    嵌套
	Person001 //模拟继承结构 匿名字段。在结构体中属于匿名结构体的字段称为提升字段。
	school    string
}

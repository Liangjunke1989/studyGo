package main

import "fmt"

//1.定义一个父类

type Person101 struct {
	name string
	age  int
}

//2.定义一个子类

type Student101 struct {
	Person101 //继承父类
	school    string
}

func (p Person101) eat() {
	fmt.Printf("%s在吃饭！\n", p.name)
}
func (s Student101) eat() { //对父类方法的重写
	fmt.Printf("%s在%s学校的食堂吃饭！\n", s.name, s.school)
}
func (s Student101) study() {
	fmt.Printf("%s需要上%s学校来学习\n", s.name, s.school)
}

func main() {
	/*
		练习一：
			创建一个小汽车类：车的品牌，车的颜色，方法：跑，加速，停止...
	*/
	/*
		练习二：
			继承中方法的使用？
				a:子类可以直接访问 父类的属性和方法
				b:子类可以新增自己的属性和方法
				c:子类可以重新父类已有的方法重新实现（重写）
	*/
	p1 := Person101{"亚当", 33}
	fmt.Println(p1)
	s1 := Student101{Person101{"夏娃", 19}, "大神学院"}
	fmt.Println(s1.name, s1.age, s1.school)
	fmt.Println("----------------------------------------------------------------------------------------")
	p1.eat()
	s1.eat() //重写后，走的是子类重写的方法。不再走父类的方法。    有重写走重写。
	s1.study()
}

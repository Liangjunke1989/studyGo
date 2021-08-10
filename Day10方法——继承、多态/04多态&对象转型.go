package main

import (
	"fmt"
	"math"
)

//1.定义一个接口

type sharp interface {
	peri() float64
	area() float64
}

//2.1 定义实现类：矩形
type rect struct {
	name   string
	length float64
	width  float64
}

//2.2 定义实现类：圆
type circle struct {
	name string
	r    float64
}

//2.3 定义实现类：三角形
type triangle struct {
	name string
	l1   float64
	l2   float64
	l3   float64
}

//2.1 定义实现类的接口的实现
func (r rect) peri() float64 {
	return (r.width + r.length) * 2
}
func (r rect) area() float64 {
	return r.width * r.length
}

//2.2 定义实现类的接口的实现
func (c circle) peri() float64 {
	return 2 * c.r * math.Pi
}
func (c circle) area() float64 {
	return math.Pow(c.r, 2) * math.Pi
}

//2.3 定义实现类的接口的实现
func (t triangle) peri() float64 {
	return t.l1 + t.l2 + t.l3
}
func (t triangle) area() float64 {
	p := t.peri() / 2 //半周长
	s := math.Sqrt(p * (p - t.l1) * (p - t.l2) * (p - t.l3))
	return s
}

//3. 测试函数
func testInterface02(s sharp) (float64, float64) {
	return s.peri(), s.area()
}

//4.对象转型
//向下转型方法一：
func getType01(s sharp) {
	/*
		instance,ok:=接口对象.(实际类型),   instance转型之后的对象，OK的值为true
	*/

	if ins, ok := s.(triangle); ok {
		fmt.Println(ins.name, "是三角形， 三边是：", ins.l1, ins.l2, ins.l3)
	} else if ins, ok := s.(rect); ok {
		fmt.Println(ins.name, "是矩形， 矩形的长宽是：", ins.length, ins.width)
	} else if ins, ok := s.(circle); ok {
		fmt.Println(ins.name, "是圆形， 圆形的半径是：", ins.r)
	} else {
		fmt.Println("还未定义形状的公式！")
	}
}

//向下转型方法二：
func getType02(s sharp) {
	/*
		方法二：接口对象.(type),配合switch和case
	*/
	switch ins := s.(type) {
	case rect:
		fmt.Println(ins.name, "是矩形， 矩形的长宽是：", ins.length, ins.width)
	case circle:
		fmt.Println(ins.name, "是圆形， 圆形的半径是：", ins.r)
	case triangle:
		fmt.Println(ins.name, "是三角形， 三边是：", ins.l1, ins.l2, ins.l3)
	}

}
func main() {
	/*
			多态：
				一个事物的多种形态。
					静态多态：方法的重载。      go语言不支持多态。
					动态多态：程序运行时的多态。

					go语言：通过接口类型模拟多态。
						一个实现类的对象：
								看作是一个实现类类型：能够访问实现类中的方法和属性。
								还可以看作是对应的接口类型：只能够访问接口中定义的方法。
		    多态用法一：
				一个函数如果接收接口类型作为参数，那么实际上可以传入该接口的任意实现类对象作为参数。
			多态用法二：
				定义一个类型作为接口类型，那么实际上可赋值任意实现类的对象。
				如果定义了一个接口类型的容器，实际上该容器可以存储任意的实现类对象。
		    逆变/协变：

			对象转型：
				向上转型：实现类对象看做接口类型对象。--->会失去实现类的新增功能。
				向下转型：instance,ok:=接口对象.(实际类型),   instance转型之后的对象，OK的值为true


	*/
	fmt.Println("-----------------------------1.接口的练习---------------------------------------------------------")
	/*
		定义一个接口：
			shape形状
			方法：周长（）面积（）
		实现类：
			矩形：长宽
			圆：半径
			三角形：三条边
		测试接口的函数：
			testShape()
	*/
	fmt.Println()
	fmt.Println("------01.矩形的周长/面积---------------")
	r1 := rect{"矩形", 82.12, 123.12}
	peri, area := testInterface02(r1)
	//fmt.Println("长为：",rect01.length,"宽为：",rect01.width,"的",rect01.name,"面积为：",area,"周长为：",peri)
	fmt.Printf("长为%f\t宽为%f\t的%s,\n计算得到面积为：%f, 周长为：%f\n", r1.length, r1.width, r1.name, peri, area)
	fmt.Println()
	fmt.Println("------02.三角形的周长/面积---------------")
	tri1 := triangle{
		name: "三角形",
		l1:   3,
		l2:   4,
		l3:   5,
	}
	p02, a02 := testInterface02(tri1)
	fmt.Printf("长为%f,%f,%f\t的%s,\n计算得到面积为：%f, 周长为：%f", tri1.l1, tri1.l2, tri1.l3, tri1.name, a02, p02)
	fmt.Println()

	fmt.Println("------03.圆的周长/面积---------------")
	c1 := circle{"大圆", 5.23}
	k, v := testInterface02(c1)
	fmt.Println(k, v)

	fmt.Println("-----------------------------2.多态的练习---------------------------------------------------------")
	var t1 triangle
	t1 = triangle{"三角形", 3, 4, 5} //创建对象 new一个对象
	fmt.Println(t1.name)
	fmt.Println(t1.peri())
	fmt.Println(t1.area())
	var s1 sharp
	s1 = t1 //将s1指向t1, s1作为接口类型，只能访问接口中的方法。     s1为父，t1为子
	//s1.name
	fmt.Println(s1.peri())
	fmt.Println(s1.area())

	//定义一个接口类型的数组
	arr := [4]sharp{r1, c1, t1, s1}
	fmt.Println(arr)

	fmt.Println()
	fmt.Println("-----------------------------3.对象转型---------------------------------------------------------")
	fmt.Println("----------------接口类型的对象转成对应实现类类型-----------------------------")
	getType01(s1)
	getType01(r1)
	getType02(s1)
	getType02(c1)
}

package main

import "fmt"

//1.定义一个接口

type USB interface {
	start()
	end()
}

//2.实现类

type Mouse struct {
	name string
}

func (m Mouse) start() {
	fmt.Println(m.name, "准备就绪，可以左右键点击了！！")
}
func (m Mouse) end() {
	fmt.Println(m.name, "可以安全退出！！")
}

type FlashDisk struct {
	name string
}

func (f FlashDisk) start() {
	fmt.Println(f.name, "准备就绪，可以开始数据存储了！！")
}
func (f FlashDisk) end() {
	fmt.Println(f.name, "可以安全弹出！！")
}
func testInterface(usb USB) {
	usb.start()
	usb.end()
}

func main() {
	/*
		接口：
			interface
			面向对象世界中的接口，一般定义是"接口定义对象的行为"
			功能定义抽取出来，实现不要。
		类：
			属性-->变量
			行为——>方法
		接口：
			功能的描述——>方法声明
			功能的实现——>方法体

			接口是功能的描述的集合。
			接口定义具有哪些功能即可，不需要功能的实现。

			面向接口编程——>面向实现类编程  （定义——实现）
			在Go中，接口是一组方法签名。当类型为接口中的所有方法提供定义时，它被称为实现接口。他与OOP非常类似。
			接口指定了类型应该具有的方法，类型决定了如何实现这些方法。

			它把所有的具有共同的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。
			接口定义了一组方法，如果某个对象实现了某个接口的所有方法，则此对象就是实现了该接口。
	*/
	/*
		接口：
			方法的描述的集合。
			接口没有字段属性,只有方法!
		实现类：
			只要实现了该接口中的所有方法，那么该类就叫做该接口的实现类。
	*/

	fmt.Println("---------------------------1.接口实现的语法一：---------------------------------------------")
	m := Mouse{"罗技"}
	f := FlashDisk{"闪迪盘"}
	//1.定义一个接口类型的变量
	//var usb USB
	//2.创建接口的任意实现类对象
	//usb=m
	//fmt.Println(usb.name)  //接口对象，不能访问实现类的属性。    多态-----------------重点！！！
	//usb=f
	//usb.start()
	//usb.end()
	fmt.Println("---------------------------2.接口实现的语法二：利用多态---------------------------------------------")
	fmt.Println("-----------------")
	testInterface(m) //利用多态         逆变/协变深入了解加深
	fmt.Println("-----------------")
	testInterface(f) //利用多态

	/*
		注意点：
			1.当需要接口类型的对象时，那么可以使用任意实现类对象代替。
			2.
	*/

}

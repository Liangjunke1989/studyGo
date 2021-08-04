package main

import "fmt"

func main() {
	/*
		函数式编程：
			支持将函数作为另一个函数的参数使用，这个参数函数叫做回调函数。
			也支持将函数作为另一个函数的返回值。

		闭包：
		一个函数有内层函数，该内层函数中，还会操作外层函数的局部变量（外层函数中定义，外层函数的参数），
		并且该外层函数的返回值就是这个内层函数。
		局部变量不会随着外层函数的销毁而销毁，这是不同的地方。
		该内层函数和局部变量，统称为闭包结构。    //不会销毁所使用到的局部变量。
	*/
	fmt.Println(increment)
	m01 := increment()
	//fmt.Println(m01)
	//fmt.Printf("%T\n",m01)     //返回值是匿名函数
	//n:=increment()() //将返回值函数进行调用
	//fmt.Printf("%T\n",n)
	//fmt.Printf("%d\n",n)

	//返回值为func()int 类型的函数，即委托类型的函数
	fmt.Println(m01)   //   此时外层函数调用完成，但是i却没有被销毁
	fmt.Println(m01()) //1
	num := m01()       //2
	m02 := m01()       //3
	m03 := m01()       //4
	fmt.Println(m01()) //5
	m04 := m01()       //6
	fmt.Println(num, m02, m03, m04)
	fmt.Println(m01()) //7
	fmt.Println("------------------------------内存地址没变，但值重新赋值i，发生了改变-----------------------------")
	m01 = increment()      //重新开辟了空间
	fmt.Println(increment) //地址虽然一样，其实意义不同了。但是可以理解为go语言做了优化处理！
	fmt.Println(m01)
	fmt.Println(m01()) //1
	//函数作为参数使用时，

}
func increment() func() int { //increment():外层函数
	i := 0
	f := func() int { //f：内层函数
		i++
		return i
	}
	//注： f 与 f()不同
	//return f()  此时返回的是i
	return f // 此时返回的是函数，本质返回的是内层函数的内存地址
}

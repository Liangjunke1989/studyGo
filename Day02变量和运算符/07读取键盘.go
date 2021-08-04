package main

import (
	"bufio"
	io "fmt"
	"os"
) //I/O 输入输出,包的自定义名称

/*
	读取键盘：
		阻塞式，类似与c#中的 Console.ReadLine();
		Scanln()
		Scanf()
		reader:=bufio.NewReader(os.Stdin)
		reader.ReadLine()//拿到切片，可变数组
		reader.ReadString('\n')//拿到字符串

*/
func main() {
	//io.Println()
	//var a int
	//var b float64
	//var c string
	//io.Println("请输入一个整数，一个浮点数，一个字符串：")
	//io.Scanln( //阻塞式，类似与c#中的 Console.ReadLine();
	//	&a,
	//	&b,
	//	&c,
	//)
	//io.Scanf("%d:%f:%c",&a,&b,&c)
	//io.Printf("a的值为：%d\n",a)
	//io.Printf("b的值为：%f\n",b)
	//io.Printf("c的值为：%s\n",c)

	//var m string
	//io.Println("请输入一个字符串：")
	//io.Scanln(&m)//hello world
	//io.Println(m)//hello

	//go中函数和方法不同
	io.Println("请输入一个字符串：")
	reader := bufio.NewReader(os.Stdin)
	//方式一：
	//data, _, _ := reader.ReadLine()//拿到的结果是数组，切片（变长数组）
	//io.Printf("data的数据类型为：%T\n内容编码是：%v\n",data,data)
	//io.Println(string(data))
	//方式二：
	readString, _ := reader.ReadString('\n')
	io.Println(readString)
}

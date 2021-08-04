package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var str1 = 'a' //4 int32
	var str2 = '中' //4  int32

	var str3 = "a"         //16     占1个字节
	var str33 = "ab"       //16     占2个字节
	var str4 = "中"         //16      占3个字节
	var str5 = "中国"        //16    占6个字节
	var str6 = "中国人民共和国万岁" //16 go语言中每个中文占3个字节，占27个字节
	var int1 = 1000000000  //8

	/*
		字节（byte）是计量单位，表示数据量多少，通常情况下一个字节等于8个01码，8位。
		字符（Character）计算机中使用使用的字母、数字、字和符号。

		ASCII 码中，一个英文字母（不分大小写）为一个字节，一个中文汉字为两个字节。
		UTF-8 编码中，一个英文字为一个字节，一个中文为三个字节。  ---go采用utf-8编码---
		Unicode 编码中，一个英文为一个字节，一个中文为两个字节。
		符号：英文标点为一个字节，中文标点为两个字节。例如：英文句号 . 占1个字节的大小，中文句号 。占2个字节的大小。
		UTF-16 编码中，一个英文字母字符或一个汉字字符存储都需要 2 个字节（Unicode 扩展区的一些汉字存储需要 4 个字节）。
		UTF-32 编码中，世界上任何字符的存储都需要 4 个字节。
	*/

	//sizeof在编译时求值，和机器有关。 sizeof占用的内存空间的大小。
	fmt.Println("________占用的内存空间的大小________")
	fmt.Println(unsafe.Sizeof(str1))  //4,类型为int32, 一个字节byte为uint8    所以字节长度为4    C#中采用一个字节
	fmt.Println(unsafe.Sizeof(str2))  //4,类型为int32, 字节长度也为4                           C#中采用两个字节
	fmt.Println(unsafe.Sizeof(str3))  //16
	fmt.Println(unsafe.Sizeof(str33)) //16
	fmt.Println(unsafe.Sizeof(str4))  //16
	fmt.Println(unsafe.Sizeof(str5))  //16
	fmt.Println(unsafe.Sizeof(str6))  //16
	fmt.Println(unsafe.Sizeof(int1))  //8

	fmt.Println("________字符串的占有的字节数________")
	//fmt.Println(len(str1)) //1
	fmt.Println(len(str3))  //1      占一个字节    字符串，编码中，一个英文字为一个字节，一个中文为三个字节
	fmt.Println(len(str33)) //2      占两个个字节
	fmt.Println(len(str4))  //3
	fmt.Println(len(str5))  //6
	fmt.Println(len(str6))  //9

	fmt.Printf("%T\n", str1)                          //类型为int32,即rune.相当于c#语言中的char
	fmt.Printf("%T,%s\n", string(str1), string(str1)) //类型转换
	fmt.Printf("%T", str3)
}

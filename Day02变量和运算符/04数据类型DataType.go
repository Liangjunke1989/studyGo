package main

import "fmt"

func main() {

	/*
		数据类型：
			int32与rune相同，类似于char，但注意又是不一样的类型。
			unit8与byte相同
	*/
	var a int32 = 32
	var b int64 = 32
	var c int
	var d rune
	var e uint8 = 255
	var f byte
	//c=a//不能赋值
	//c=b//不能赋值
	d = a
	f = e
	fmt.Println(a, b, c, d, e, f)
}

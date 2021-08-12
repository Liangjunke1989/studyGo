package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		Seek:寻找
			用于设置指针光标的位置。
			第一个参数：偏移量
				距离文件开头offset个字节
			第二个参数：如何偏
				0：距离文件开头offset个字节
				1：距离光标当前位置多少个字节
				2：距离光标末尾位置多少个字节
		随机读取文件：
			可以设置指针光标的位置。
	*/
	file, _ := os.OpenFile("./TestDir/test.txt", os.O_RDWR, os.ModePerm)
	bs := []byte{0}
	file.Seek(10, 0) //从哪便宜
	_, err := file.Read(bs)
	if err != nil {
		fmt.Println("文件读取错误！", err.Error())
	} else {
		fmt.Println(string(bs))
	}
	_, err2 := file.Seek(0, 2)
	fmt.Println(err2.Error())
	file.WriteString("lolololo")
	fmt.Println("-----------------------------练习一：断点续传------------------------------------------------------")
}

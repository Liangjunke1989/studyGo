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
	fmt.Println("------------------------1.Seek练习一--------------------------------------")
	_, err := file.Seek(10, 0) //从哪偏移
	//文件读取
	_, err = file.Read(bs)
	if err != nil {
		fmt.Println("文件读取错误！", err.Error())
	} else {
		fmt.Println(string(bs)) //j
	}
	fmt.Println("------------------------2.Seek练习二--------------------------------------")
	_, err2 := file.Seek(0, 2)
	if err2 != nil {
		fmt.Println("文件光标设置有误！", err2.Error())
	} else {
		//文件写入
		_, err = file.WriteString("lolololo")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("文件写出完成！！！")
		}
	}
	//文件关闭
	err = file.Close()
	if err != nil {
		fmt.Println("文件关闭有误！！", err.Error())
	} else {
		fmt.Println("文件成功关闭！！")
	}
	fmt.Println("-----------------------------练习一：断点续传------------------------------------------------------")
}

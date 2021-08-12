package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		文件的写入：
			将程序中的数据写出到文件中。
			又称数据的持久化保存。
			程序中的数据称为瞬时性的数据。（内存中）
	*/
	//0.备用写入文件的数据
	bytes := []byte{65, 66, 67, 68, 69, 70}

	//1.打开文件，建立连接
	//1.1 os.open()
	//file, err := os.Open("./Day12io输入输出/aabbcc.txt")
	//if err!=nil{
	//	fmt.Println("文件打开有误！",err.Error())
	//}else {
	//	fmt.Println(file.Name())
	//}
	//1.2 os.OpenFile()
	file, err := os.OpenFile("./Day12io输入输出/aabbcc.txt", os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println("文件打开有误！", err.Error())
	} else {
		fmt.Println("------------------------------------2.文件写入操作---------------------------------------------")
		//01.   file.Write()
		{
			n, err := file.Write(bytes) //n 为写出的字节数
			if err != nil {
				fmt.Println("文件写入有误！", err.Error())
			} else {
				fmt.Println("写出文件成功！总共字节数为：", n)
			}
		}
		//02.   file.WriteString()
		{
			n, err := file.WriteString("LJKLJKLJKLJKLJKJLJKLJKLJKLJKLJKLJKLJKLJK") //n 为写出的字节数
			if err != nil {
				fmt.Println("文件写入有误！", err.Error())
			} else {
				fmt.Println("写出文件成功！总共字节数为：", n)
			}
		}
	}
	fmt.Println("------------------------------------3.关闭文件，断开连接--------------------------------------------")
	//3. 关闭文件，断开连接
	err = file.Close()
	if err != nil {
		fmt.Println("文件关闭错误！", err.Error())
	} else {
		fmt.Println("文件已经成功关闭！")
	}
	fmt.Println()
	fmt.Println("------------------------------------4.简写文件写出操作--------------------------------------------")
	file02, _ := os.OpenFile("./Day12io输入输出/aabbcc.txt", os.O_WRONLY, os.ModePerm)
	file02.WriteString("1234567890") // 光标一开始位于文件开头的位置，替换已有的字节
	file.Close()
	file03, _ := os.OpenFile("./Day12io输入输出/aabbcc.txt", os.O_WRONLY, os.ModePerm)
	file03.WriteString("hello面12") // 光标一开始位于文件开头的位置，替换已有的字节
	file.Close()
}

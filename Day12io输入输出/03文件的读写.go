package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		文件的读写：
			step01: 打开文件
				os.Open()
			step02: 读/写
			step03: 关闭文件
	*/
	//01 确定文件名&文件路径
	fileName := "./Day12io输入输出/aabbcc.txt"
	//02 打开文件，建立连接
	file01, err01 := os.Open(fileName)
	if err01 != nil {
		fmt.Println("文件打开有误！", err01.Error())
		return
	} else {
		fmt.Println("文件成功打开！", file01)
	}

	//03 对文件进行读/写
	bytes := make([]byte, 8, 8)
	//第一次读取
	//从file01对应的文件中读取最多len(bytes)个数据,存入到bytes数组中。
	n, err02 := file01.Read(bytes) //返回的是读取的个数和错误
	if err02 != nil {
		fmt.Println("文件读取有误！", err02.Error())
		return
	} else {
		fmt.Println("文件读取成功！", n) // 14, 为返回的字节的个数
		//fmt.Printf("%T\n",bytes)
		fmt.Println(string(bytes))
	}
	//第二次读取
	n, err03 := file01.Read(bytes) //返回的是读取的个数和错误
	if err03 != nil {
		fmt.Println("文件读取有误！", err03.Error())
		return
	} else {
		fmt.Println("文件读取成功！", n) // 14, 为返回的字节的个数
		//fmt.Printf("%T\n",bytes)
		fmt.Println(string(bytes))
	}
	//第三次读取
	n, err04 := file01.Read(bytes) //返回的是读取的个数和错误
	if err04 != nil {
		fmt.Println("文件读取有误！", err04.Error())
		return
	} else {
		fmt.Println("文件读取成功！", n) // 14, 为返回的字节的个数
		//fmt.Printf("%T\n",bytes)
		fmt.Println(string(bytes))
	}
	//第四次读取
	n, err05 := file01.Read(bytes) //返回的是读取的个数和错误
	if err05 != nil {
		fmt.Println("文件读取有误！", err05.Error())
		return
	} else {
		fmt.Println("文件读取成功！", n) // 14, 为返回的字节的个数
		//fmt.Printf("%T\n",bytes)
		fmt.Println(string(bytes))
	}
	//04 关闭文件，断开连接
	err99 := file01.Close()
	if err99 == nil {
		fmt.Println("文件已成功关闭！")
	}
}

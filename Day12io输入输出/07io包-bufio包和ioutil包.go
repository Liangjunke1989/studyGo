package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	/*
		io包：
			最基本的包，读写的常规操作。
			Reader接口，Writ接口
			io.copy()
		bufio包：																	——————————核心包————————————
			buffer缓存
			将io包下的Reader,Write对象进行封装，带缓存的封装，提供读写效率。
			将io包下的Reader和Write进一步封装，带缓存的Reader和Write，实现高效的读和写。
			ReadBytes()
			ReadString()
			ReadLine()
		ioutil包：
			工具包
			ReadFile()       //一次性读取所有
			WriteFile()      //一次性全部写入
			ReadDir()        //读取目录
			ReadAll()
	*/
	fmt.Println("-------------------------------1.io/ioutil包-----------------------------------------------------")
	{
		//1.ReadFile() 读取文件中所有的数据
		{
			//	bytes, err := ioutil.ReadFile("./TestDir/test.txt") //ReadFile()
			//	if err != nil {
			//		fmt.Println("文件读取有错！", err.Error())
			//	} else {
			//		fmt.Println("文件读取成功！", string(bytes))
			//	}
		}
		//2.WriteFile() 写出文件中所有的数据
		{
			//bytes=make([]byte,1024,1024)
			//	s1 := "helloWorld_ljk"
			//	bytes := []byte(s1)
			//	err := ioutil.WriteFile("./TestDir/test2.txt", bytes, os.ModePerm) //从文件头位置开始写入
			//	if err != nil {
			//		fmt.Println("文件写入有误！", err.Error())
			//	} else {
			//		fmt.Println("s1字符串成功写入到文件中！！")
			//	}
		}
		//3.复制文件
		{
			//	bytes, err := ioutil.ReadFile("./TestDir/test2.txt")
			//	err = ioutil.WriteFile("./TestDir/test3.txt", bytes, os.ModePerm)
			//	if err != nil {
			//		fmt.Println("文件写入有误！", err.Error())
			//	} else {
			//		fmt.Println("test2已经复制到test3中了！！")
			//	}
		}
		//4.ReadAll()
		{
			//	str := "ljkljkljklkjkljkljljkljljkljljkljkljlkjk"
			//	reader:= strings.NewReader(str)                               //strings类   类似于c#StringBuilder
			//	bytes, err := ioutil.ReadAll(reader)
			//	if err!=nil {
			//		fmt.Println("文件读取有误！",err.Error())
			//	}else {
			//		fmt.Println("文件成功读取",string(bytes))
			//	}
		}
		//5.ReadDir(), 读取一个目录下的子内容：子文件和子目录，但是仅一层
		{
			//	dirName := "./TestDir"
			//	fileInfos, err := ioutil.ReadDir(dirName) //获取到文件信息
			//	if err != nil {
			//		fmt.Println("文件读取有误！", err.Error())
			//	} else {
			//		fmt.Println("文件成功读取,总共读取到的文件数是：", len(fileInfos))
			//		fmt.Println("----------------------------")
			//		for _, v := range fileInfos {
			//			//fmt.Printf("切片中存储的元素的格式为%T\n",v)
			//			fmt.Println(v.Name())
			//			fmt.Printf("是否是一个目录：%t\n", v.IsDir())
			//
			//		}
			//	}
		}
	}
	fmt.Println("-------------------------------2.io/bufio包：执行更高效--------------------------------------------")
	{
		//1. ReadLine() 和 ReadBytes()
		{
			//file, _ := os.Open("./TestDir/test3.txt")        // file可以被看作io包下Reader和Write的实现。
			//rd := bufio.NewReaderSize(file, 1024)		       // 构建带缓存的reader对象：bufio.Reader
			//fmt.Printf("%T\n",rd)
			//bytes, isPrefix, _ := rd.ReadLine() //一次读一行,        如果是字符用rd.ReadLine()
			//														// 如果是图片数据，用rd.ReadBytes()
			//fmt.Println(string(bytes))
			//fmt.Println(isPrefix) //并不是前缀，可以读取完整，所以返回false
			//						// false能够读取完整，true不能读取完整

			//关闭文件，断开连接
			//err:=file.Close()
			//if err!=nil {
			//	fmt.Println("文件关闭有误！",err.Error())
			//}else {
			//	fmt.Println("文件已经成功关闭！！")
			//}
		}
		//2.ReadString
		{
			file, _ := os.Open("./TestDir/test3.txt") // file可以被看作io包下Reader和Write的实现。
			rd := bufio.NewReaderSize(file, 1024)     // 构建带缓存的reader对象：bufio.Reader

			for {
				readString, err := rd.ReadString('_')
				if err == io.EOF {
					fmt.Println("文件全部读取完毕！")
					break //打断，不然一直循环进行
				} else if err != nil {
					fmt.Println("文件读取有误！", err.Error())
					return
				} else {
					fmt.Println("文件读取成功！内容为：", readString)
				}
			}

			//关闭文件
			err := file.Close()
			if err != nil {
				fmt.Println("文件关闭有误！", err.Error())
			} else {
				fmt.Println("文件已经成功关闭！！")
			}
		}
	}
}

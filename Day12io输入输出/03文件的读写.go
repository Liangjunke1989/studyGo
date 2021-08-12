package main

import (
	"fmt"
	"io"
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
	fmt.Println("----------------------------01.确定文件名&文件路径------------------------------------------------")
	//01 确定文件名&文件路径
	fileName := "./Day12io输入输出/aabbcc.txt"
	fmt.Println("----------------------------02.打开文件，建立连接------------------------------------------------")
	//02 打开文件，建立连接
	file01, err01 := os.Open(fileName)
	if err01 != nil {
		fmt.Println("文件打开有误！", err01.Error())
		return
	} else {
		fmt.Println("文件成功打开！", file01)
	}
	fmt.Println("----------------------------03.对文件进行读取------------------------------------------------------")
	//03 对文件进行读/写
	bytes := make([]byte, 8, 8) //创建了一个长度为8的容器
	n := -1
	var err error
	//逐次推导过程：
	{
		////第一次读取
		////从file01对应的文件中读取最多len(bytes)个数据,存入到bytes数组中。
		//n, err02 := file01.Read(bytes) //返回的是读取的个数和错误           ----------------->文件读取的核心
		//if err02 != nil {
		//	fmt.Println("--------第一次读取---------")
		//	fmt.Println("文件读取有误！", err02.Error())
		//	return
		//} else {
		//	fmt.Println("--------第一次读取---------")
		//	fmt.Println("文件读取成功！", n) // 8, 为返回的字节的个数
		//	fmt.Println("bytes中缓存的数据为：")
		//	for _,v :=range bytes{
		//		fmt.Printf("%d \t",v)
		//	}
		//	fmt.Println()
		//	//fmt.Printf("%T\n",bytes)
		//	fmt.Println("解析到的字符串类型的数据为：",string(bytes[:n]))
		//}
		////第二次读取
		//n, err03 := file01.Read(bytes) //返回的是读取的个数和错误
		//if err03 != nil {
		//	fmt.Println("--------第二次读取---------")
		//	fmt.Println("文件读取有误！", err03.Error())
		//	return
		//} else {
		//	fmt.Println("--------第二次读取---------")
		//	fmt.Println("文件读取成功！", n) // 8, 为返回的字节的个数
		//	fmt.Println("bytes中缓存的数据为：")
		//	for _,v :=range bytes{
		//		fmt.Printf("%d \t",v)
		//	}
		//	fmt.Println()
		//	//fmt.Printf("%T\n",bytes)
		//	fmt.Println("解析到的字符串类型的数据为：",string(bytes[:n]))
		//}
		////第三次读取
		//n, err04 := file01.Read(bytes) //返回的是读取的个数和错误
		//if err04 != nil {
		//	fmt.Println("--------第三次读取---------")
		//	fmt.Println("文件读取有误！", err04.Error())
		//	return
		//} else {
		//	fmt.Println("--------第三次读取---------")
		//	fmt.Println("文件读取成功！", n) // 8, 为返回的字节的个数
		//	fmt.Println("bytes中缓存的数据为：")
		//	for _,v :=range bytes{
		//		fmt.Printf("%d \t",v)
		//	}
		//	fmt.Println()
		//	//fmt.Printf("%T\n",bytes)
		//	fmt.Println("解析到的字符串类型的数据为：",string(bytes[:n]))
		//}
		////第四次读取
		//n, err05 := file01.Read(bytes) //返回的是读取的个数和错误
		//if err05 != nil {
		//	fmt.Println("--------第四次读取---------")
		//	fmt.Println("文件读取有误！", err05.Error())
		//	return
		//} else {
		//	fmt.Println("--------第四次读取---------")
		//	fmt.Println("文件读取成功！", n) // 2, 为返回的字节的个数
		//	fmt.Println("bytes中缓存的数据为：")
		//	for _,v :=range bytes{
		//		fmt.Printf("%d \t",v)
		//	}
		//	fmt.Println()
		//	//fmt.Printf("%T\n",bytes)
		//	fmt.Println("解析到的字符串类型的数据为：",string(bytes[:n]))
		//}
		////第五次读取
		//n, err06 := file01.Read(bytes) //返回的是读取的个数和错误
		//if err06 != nil {
		//	fmt.Println("--------第五次读取---------")
		//	fmt.Println("文件读取有误！", err06.Error())  //EOF:end of file
		//	return
		//} else {
		//	fmt.Println("--------第五次读取---------")
		//	fmt.Println("文件读取成功！", n) // 0
		//	fmt.Println("bytes中缓存的数据为：")
		//	for _,v :=range bytes{
		//		fmt.Printf("%d \t",v)
		//	}
		//	fmt.Println()
		//	//fmt.Printf("%T\n",bytes)
		//	fmt.Println("解析到的字符串类型的数据为：",string(bytes[:n]))
		//}
	}
	//for优化写法：
	{
		for i := 0; ; {
			n, err = file01.Read(bytes) //返回的是读取的个数和错误                  //----------------->文件读取的核心
			i++
			if n == 0 || err == io.EOF {
				fmt.Printf("在第%d次，文件读取完毕！", i)
				//fmt.Println("文件读取有误！", err.Error())
				break //break跳出本次循环
			} else {
				fmt.Printf("第%d次解析到的字符串类型的数据为：%s\n", i, string(bytes[:n]))
			}
		}
	}
	fmt.Println("----------------------------04.关闭文件，断开连接------------------------------------------------------")
	//04 关闭文件，断开连接
	err99 := file01.Close()
	if err99 == nil {
		fmt.Println("文件已成功关闭！")
	}
}

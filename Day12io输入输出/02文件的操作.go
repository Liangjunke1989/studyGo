package main

import (
	"fmt"
	"os"
)

func main() {
	/*
			文件的操作：
				文件路径和文件名的区别？

				1.文件路径：
					相对路径：relative   都是相对于当前工程
						../aa.txt 表示上一层目录
						./aa.txt 表示当前目录
					绝对路径：absolute   从根盘符开始
						/Users/liangjunke/Desktop/三维生态景观设计软件开发计划.docx
				2.文件的信息。
				3.文件的常规操作：
					创建文件夹，创建文件，删除文件，打开文件，关闭文件... ...


		        4.文件的读/写。
				5.文件的复制。
				6.ioutil包
				7.bufio包

				  	作业1：文件的查找seek()  --->断点续传
					作业2：遍历一个目录，遍历给定的文件夹，包含子文件夹。
	*/
	fmt.Println("-------------------------------1.filepath文件路径：相对路径/绝对路径---------------------------------")
	{
		//fmt.Println("----------文件名/文件路径------------")
		//fileName1:="/Users/liangjunke/Desktop/三维生态景观设计软件开发计划.docx"
		//fileName2:="aa.txt"
		//fmt.Println(filepath.IsAbs(fileName1))//true
		//fmt.Println(filepath.IsAbs(fileName2))//false 只判断是否为根目录，不判断文件存在
		//fmt.Println("----------绝对路径------------")
		//abs, _ := filepath.Abs(fileName1)//
		//fmt.Println(abs)
		//fmt.Println(filepath.Abs(fileName2))
		//fmt.Println("----------上一层路径/当前路径------------")
		//fmt.Println("获取父目录：",path.Join(fileName1,".."))//获取文件的上一层路径         ——>技巧1
		//fmt.Println()
	}
	fmt.Println("-------------------------------2.fileInfo文件信息：根据文件名获取文件信息-----------------------------")
	{
		//fileInfo, err := os.Stat(fileName1)
		//if err != nil {
		//	fmt.Println(err.Error())
		//} else {
		//	fmt.Printf("%T\n", fileInfo)
		//}
		//fmt.Println(fileInfo.Name())
		//fmt.Println(fileInfo.Size())
		//fmt.Println(fileInfo.ModTime())
		//fmt.Println(fileInfo.Mode()) //-rw-r--r--   x可执行   也可以用八进制表示： r 4,  w 2,  x 1, - 0
		////fmt.Println(fileInfo.Sys())
		//fmt.Println()
	}
	fmt.Println("-------------------------------3.1文件的常规操作：创建文件夹-----------------------------------------")
	{
		////3.1 os.Mkdir()     只能创建一层目录
		//err2 := os.Mkdir("/Users/liangjunke/Desktop/LJKTest20210811", os.ModePerm) //os.ModePerm 0777 可读可写可执行
		////            0666 可读可写
		//if err2 != nil {
		//	fmt.Println(err2.Error())
		//} else {
		//	fmt.Println("文件夹创建成功了！！！！")
		//}
		////3.2 os.MkdirAll()  可以同时创建多层目录
		//err3 := os.MkdirAll("/Users/liangjunke/Desktop/LJKTest20210811_02/aa/bb/cc", os.ModePerm)
		//if err3 != nil {
		//	fmt.Println(err3.Error())
		//} else {
		//	fmt.Println("文件夹创建成功了！！！！")
		//}
		//fmt.Println()
	}
	fmt.Println("-------------------------------3.2文件的常规操作：创建文件-------------------------------------------")
	{
		////3.2 创建一个文件	使用的是0666模式（可读可写）
		//file1, err4:= os.Create("/Users/liangjunke/Desktop/LJKTest20210811/abc.txt")//如果已经存在的话，会被覆盖，重新创建
		//if err4!=nil {
		//	fmt.Println(err4.Error())
		//}else {
		//	fmt.Println(file1)
		//	fmt.Println("文件创建成功了！！！！")
		//}
		//file2, _:= os.Create("./Day12io输入输出/aabbcc.txt")
		//fmt.Println(file2.Name())
		//fmt.Println()
	}
	fmt.Println("-------------------------------3.3-3.4文件的常规操作：打开/关闭文件-----------------------------------")
	{
		fileName001 := "/Users/liangjunke/Desktop/LJKTest20210811/abc.txt"
		//fileName002:="./Day12io输入输出/aabbcc.txt"

		//3.3  打开文件，让当前的程序和指定的文件建立了一个连接
		//3.3.1 os.Open()
		file001, err001 := os.Open(fileName001) //打开，就是建立了一个连接，操作指针地址就是操作文件
		if err001 != nil {
			fmt.Println(err001.Error())
		} else {
			fmt.Println("文件打开成功！！", file001)
		}
		//3.3.2 os.OpenFile()
		/*
			第一个参数：文件的名称
			第二个参数：文件的打开方式
				O_RDONLY: 只读模式
				O_WRONLY: 只写模式
				O_RDWR:   可读可写模式
				O_APPEND: 追加模式
				O_CREATE: 文件不存在就创建
				O_EXCL:   在O_CREATE一起使用，构成一个新建文件的功能，他要求文件必须不存在。
				O_SYNC:   同步方式打开，即不使用缓存，追加写入硬盘
				O_TRUNC:  打开并清空文件
			第三个参数：文件的权限
				文件不存在，创建文件，需要指定权限
				文件存在，可以不写
		*/
		file002, err002 := os.OpenFile(fileName001, os.O_WRONLY|os.O_CREATE, os.ModePerm)
		if err002 != nil {
			fmt.Println(err002.Error())
		} else {
			fmt.Println("文件打开成功！！", file002)
		}
		//3.4  关闭文件，让当前的程序和指定的文件的连接断开

		err003 := file001.Close()
		if err003 != nil {
			fmt.Println(err003.Error())
		} else {
			fmt.Println("文件关闭成功了！！！")
		}
		err004 := file002.Close()
		if err004 != nil {
			fmt.Println(err004.Error())
		} else {
			fmt.Println("文件关闭成功了！！！")
		}
		fmt.Println()
	}
	fmt.Println("-------------------------------3.5文件的常规操作：删除文件夹/文件-------------------------------------")
	{
		//3.5.1 删除文件夹
		//os.Remove()  删除空文件和文件夹
		err005 := os.Remove("/Users/liangjunke/Desktop/LJKTest20210811_02/aa/bb")
		if err005 != nil {
			fmt.Println(err005.Error())
		} else {
			fmt.Println("文件删除成功了！！！")
		}
		//os.RemoveALl()  删除所有  -rs   循环删除
		err006 := os.RemoveAll("/Users/liangjunke/Desktop/LJKTest20210811")
		if err006 != nil {
			fmt.Println(err006.Error())
		} else {
			fmt.Println("文件删除成功了！！！")
		}
	}
}

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	/*
		复制文件：
			文件路径和文件名称？

	*/

	_srcFile := "./Day12io输入输出/aabbcc.txt"
	_destFile := "./TestDir/test.txt"
	//自定义函数一：
	{
		//total000, err000 := copyFile1(_srcFile, _destFile)
		//if err000 != nil {
		//	fmt.Println(err000.Error())
		//} else {
		//	fmt.Println("文件拷贝成功!!!", total000)
		//}
	}
	//自定义函数二：
	{
		total001, err001 := copyFile2(_srcFile, _destFile)
		if err001 != nil {
			fmt.Println(err001.Error())
		} else {
			fmt.Println("文件拷贝成功!!!", total001)
		}
	}
}

//自定义函数一：实现函数的拷贝，返回值是拷贝的字节的总数量和发生的错误
func copyFile1(src, dest string) (int, error) {
	srcfile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	destFile, err2 := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err2 != nil {
		return 0, err2
	}
	defer fmt.Println("所有文件关闭成功！")
	defer srcfile.Close() //defer的使用？在有return的情况下
	defer destFile.Close()
	defer fmt.Println("所有文件开始关闭！")
	//拷贝数据
	bytes := make([]byte, 1024*4, 1024*4) //4kb
	n := -1
	total := 0
	for {
		n, err = srcfile.Read(bytes)
		if err == io.EOF || n == 0 {
			fmt.Println("已从文件1拷贝完毕！！")
			break //break 和 return的区别？
		} else if err != nil {
			fmt.Println("拷贝异常！！")
			return total, err
		}
		total += n
		_, err := destFile.Write(bytes[:n])
		if err != nil {
			fmt.Println("往文件2中写出有误！", err.Error())
		} else {
			fmt.Println("往文件2中写出完毕！！！")
		}
	}
	return total, nil
}

//自定义函数二：实现函数的拷贝，返回值是拷贝的字节的总数量和发生的错误
func copyFile2(src, dest string) (int64, error) {
	srcfile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	destFile, err2 := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err2 != nil {
		return 0, err2
	}
	defer fmt.Println("所有文件关闭成功！")
	defer srcfile.Close() //defer的使用？在有return的情况下
	defer destFile.Close()
	defer fmt.Println("所有文件开始关闭！")
	n, err3 := io.Copy(destFile, srcfile)
	return n, err3
}

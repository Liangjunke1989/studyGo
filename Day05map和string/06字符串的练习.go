package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		练习一：给定一个路径名，获取文件名。
		pathName="Http://192.168.1.170:80/goStudyProject/Day05/kk.jpeg"
	*/
	pathName := "Http://192.168.1.170:80/goStudyProject/Day05/kk.jpeg"
	index := strings.LastIndex(pathName, "/")
	fileName := pathName[index+1:]
	fmt.Println(fileName)
	/*
		练习二：大写字母有多少个？
	*/
	//str:="adadDF312AFAFDd123fDFADFAFV132CdfaffafaDFAF"
	//count1:=0 //大写字母个数计数器[65,90]
	//count2:=0 //小写字母计数器[97,122]
	//count3:=0 //非字母的个数

}

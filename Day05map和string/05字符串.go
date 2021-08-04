package main

import (
	"fmt"
	"strings"
)

var s3 string

func main() {
	/*
		字符串：
			一个字节的切片。
			一些字节的集合。
			一个字符的序列。

	*/
	fmt.Println("-------------------strings包下的关于字符串的函数----------------------------")
	{
		str1 := "Hello"
		fmt.Println("-----------01查找：是否包含------------------")
		//包含
		b := strings.Contains(str1, "H")
		fmt.Println(b)
		c := strings.ContainsAny(str1, "Ha")
		fmt.Println(c)
		d := strings.Count(str1, "l")
		fmt.Println(d)
		fmt.Println("-----------02查找：索引------------------")
		//索引
		fmt.Println(strings.Index(str1, "l"))
		fmt.Println(strings.IndexAny(str1, "dal"))
		fmt.Println(strings.LastIndexAny(str1, "dal"))

	}
	{
		fmt.Println("-----------03拼接------------------")
		//拼接
		s1 := []string{"avc", "adc", "dda"}
		str := strings.Join(s1, "+")
		fmt.Println(str)
		fmt.Println("-----------04分割------------------")
		//分割
		s2 := strings.Split(str, "+")
		//fmt.Printf("%T",s2)
		for i := 0; i < len(s2); i++ {
			fmt.Println(s2[i])
			s3 += s2[i]
		}
		fmt.Println(s3)
		fmt.Println("-----------05按数量分割------------------")
		//按照数量分割
		splitN := strings.SplitN(str, "+", -1) //切成几个
		fmt.Println(splitN)
		fmt.Println("-----------06重复------------------")
		//重复repeat
		fmt.Println(strings.Repeat(str, 5))
		fmt.Println("-----------07替换------------------")
		//替换
		fmt.Println(strings.Replace(str, "a", "m", 2))
		fmt.Println("-----------08大小写替换------------------")
		//大小写转换
		fmt.Println(strings.ToLower(str))
		fmt.Println(strings.ToUpper(str))
		fmt.Println("-----------09去除首尾指定的内容------------")
		//Trim，去除首尾指定的内容
		str99 := "   *******3333*****45**%^$$***  "
		fmt.Println(strings.Trim(str99, "*33 "))
		fmt.Println(strings.TrimLeft(str99, "*33 "))
		fmt.Println(strings.TrimRight(str99, "*33 "))
		fmt.Println(strings.TrimSpace(str99))
		fmt.Println("-----------10截取字串------------")
		//截取字串
		s88 := "hello world"
		s77 := s88[6:]
		fmt.Println(s77)
	}

}

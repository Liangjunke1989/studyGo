package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

func main() {
	fmt.Println("-------------------------练习一：遍历文件夹pro，包含子目录--------------------------------------------")
	{
		//fileProName:="/Users/liangjunke/Desktop/20210729临时文件"
		////遍历文件夹
		//listDir(fileProName)
	}
	fmt.Println("-------------------------练习二：模拟聊天记录：从键盘接收数据------------------------------------------")
	{
		//chatFileName:="./TestDir/weChat.txt"
		////读取键盘输入
		//reader1 := bufio.NewReader(os.Stdin)                //创建一个bufio.Reader对象，读取键盘输入（标准输入）
		////打开要写出到的文件
		//openFile,err := os.OpenFile(chatFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
		//if err!=nil {
		//	fmt.Println("文件打开有误！",err.Error())
		//}else {
		//	//写出数据
		//	//1.创建一个bufio.Writer对象，写出数据到weChat.txt文件中
		//	writer1 := bufio.NewWriter(openFile)
		//	//2.延迟关闭
		//	defer openFile.Close()
		//	//3.1写的数据内容：日期
		//	writer1.WriteString("\n")
		//	writer1.WriteString(time.Now().Format("2006-01-02"))
		//	writer1.WriteString("\n")
		//	writer1.Flush() //刷新缓冲区：将缓冲区的数据，写入到目标文件中
		//	//3.2写的数据内容：时间+姓名+"输入的内容"
		//	str:="" //读取到的数据
		//	flag:=true
		//	name:=""
		//	content:=""
		//	for{
		//		//01.写出时间
		//		writer1.WriteString(time.Now().Format("15:04:05"))
		//		writer1.WriteString("\n")
		//		//02.写出"name：内容"
		//		if flag {
		//			name="小明"
		//		}else{
		//			name="小红"
		//		}
		//		flag=!flag
		//		//读取的键盘内容,以"\n"截取
		//		str, _= reader1.ReadString('\n')
		//		//拼接键盘输入的内容信息+姓名
		//		if	str=="over\n"{                             //手动设置退出标识"over"
		//			fmt.Println("程序即将退出！")
		//			//writer1.WriteString("over")
		//			writer1.WriteString(str)
		//			break
		//		}
		//		content=fmt.Sprint(name,":",str)
		//		fmt.Println(content)
		//		writer1.WriteString(content)
		//		//03.进行写入的数据更新
		//		writer1.Flush()
		//	}
		//}
	}
	fmt.Println("-------------------------练习三：断点续传-----------------------------------------------------------")
	{
		/*
				断点续传：
					文件传递：文件复制
					将"./Day12io输入输出/aabbcc.txt"   -----剪切/复制到 -----   "./TestDir/aabbcc.txt"中
				思路：
					1.边复制，边记录复制的总量。
					2.根据数据量，将源文件的光标移动到数据量的位置进行复制，对目标文件的光标移动到数据量的位置进行写入数据。
			        3.临时文件容器，作为参考，使用完后销毁！可读可写。
					内存：临时性数据（程序运行，开辟的缓存），程序关闭会销毁
					文件/硬盘：数据的永久化保存数据。（数据库也可以持久化保存数据）
		*/
		srcFileName := "./Day12io输入输出/head.jpg"  //源文件      可读
		destFileName := "./TestDir/copyHead.jpg" //目标文件    可写
		tempFileName := "./TestDir/tempHead.txt" //临时文件    可读可写
		openFile, _ := os.Open(srcFileName)
		copyDestFile, _ := os.OpenFile(destFileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
		tempFile, _ := os.OpenFile(tempFileName, os.O_CREATE|os.O_RDWR, os.ModePerm)

		//稍后关闭
		defer openFile.Close()
		defer tempFile.Close()
		defer copyDestFile.Close()

		//1.读取临时文件中的数据，根据seek
		tempFile.Seek(0, 0)
		bs := make([]byte, 100, 100) //只存文件数据量的大小
		n1, err := tempFile.Read(bs)
		fmt.Println("读取的数量是：", n1)
		conStr := string(bs[:n1])
		fmt.Println("临时文件获取到的字符串为：", conStr) //" "字符串
		//countTemp, _ := strconv.Atoi(conStr)//将字符串转成整型int64
		countTemp, _ := strconv.ParseInt(conStr, 10, 64) //将字符串转成整型int
		fmt.Println("转成的整型数据为：", countTemp)
		//fmt.Printf("%T",countTemp) //int64

		//2.设置读/写的偏移量
		openFile.Seek(countTemp, 0)
		copyDestFile.Seek(countTemp, 0)
		data := make([]byte, 512, 512) //内存的缓存数据
		n2 := -1                       //读入的数据量
		n3 := -1                       //写出的数据量
		total := int(countTemp)        //读取的总量
		for {
			//3.将源文件中的数据进行读取
			n2, err = openFile.Read(data) //返回读取的数据量
			if err == io.EOF {
				fmt.Println("文件复制完毕...")
				tempFile.Close()        //关闭缓存文件
				os.Remove(tempFileName) //删除缓存文件
				break
			}
			//4.往目标文件中写入数据
			n3, _ = copyDestFile.Write(data[:n2]) //将数据写入到目标文件

			//5.缓存文件中数据量更改
			total += n3
			tempFile.Seek(0, 0)                       //将复制总量，存储到临时文件中,覆盖掉原有的数据
			tempFile.WriteString(strconv.Itoa(total)) //将数值转成字符串

			//假装断电，出现异常
			//if total>12 {
			//	panic("出现断电情况了！！！！")
			//}
		}

	}
}

//利用递归，遍历文件夹
func listDir(dirName string) {
	//遍历当前文件夹名称下的文件目录
	fileInfos, err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Println("文件读取有误", err.Error())
	} else {
		//fmt.Println(len(fileInfos))
		for k, fileInfo := range fileInfos {
			fileName := path.Join(dirName, fileInfo.Name())
			fmt.Println("第", k+1, "个文件的名称为：", fileName)
			//如果是文件夹的话，继续遍历
			if fileInfo.IsDir() { //递归遍历的核心：考虑是什么时候跳出循环！！！！
				listDir(fileName)
			}
		}
	}
}

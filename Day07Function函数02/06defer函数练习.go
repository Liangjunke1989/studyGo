package main

import "fmt"

func main() {
	/*
		defer函数：
			1.当外围函数中的语句正常执行完毕时，只有其中所有的延迟函数都执行完毕，外围函数才会真正的结束。
			2.当执行外围函数中的return语句时，只有其中所有的延迟函数都执行完毕后，外围函数才会真正返回。
			3.当外围函数中的代码引发运行恐慌时，只有其中所有的延迟函数都执行完毕后，该运行时恐慌才会真正被扩展至
			  调用函数。
	*/

	fmt.Println("---------------练习一：根据defer的特点，倒叙打印字符串----------------------")
	{
		//str:="Hello World"
		//for k,_:=range str{
		//	defer fmt.Print(string(str[k]))
		//}
	}
	fmt.Println("---------------练习二：defer函数注意点----------------------")
	{
		fmt.Println("----------------")
		res1 := f1()
		fmt.Println(res1) //1
		fmt.Println("----------------")
		fmt.Println(f2()) //5
		fmt.Println("----------------")
		fmt.Println(f3())
	}
}
func f1() (result int) {
	fmt.Println("11111111")
	fmt.Println("22222222")
	defer func() {
		result++
	}()
	fmt.Println("33333333")
	//result=0
	//return result
	return 0 // 隐含动作 result=0
}
func f2() (r int) {
	//r:=0
	t := 5
	defer func() {
		t += 5
		fmt.Println("f2中的t：", t)
	}()
	//r=t
	return t
}
func f3() (r int) {
	//r:=0
	defer func(r int) { //此时的r为形参, 先走了一边函数，没执行而已（r=0）执行函数
		r += 5
		fmt.Println(r) //5
	}(r) //r传入的实参为1    //6
	//r=1
	//return r
	return 1
}

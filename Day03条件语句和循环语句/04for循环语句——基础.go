package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	/*
		01.循环语句：
			条件满足，某些代码反复多次执行。直到条件不成立为止。
	*/

	/*
		02.标准写法：
		for 表达式1；表达式2；表达式3{
			循环体
		}
		流程：
			step1: 先执行表达式1，只执行1次，习惯用于变量的初始化。
			step2: 再执行表达式2，用于循环的条件，必须是bool类型。
			step3: true，执行循环体，在执行表达式3，习惯用于变量的变化。
			step4: 在执行表达式2，看条件是否成立。直到不成立为止，循环结束。

		// 同时省略3个表达式
		for{ //相当于for true 永真
			循环体
		}
	*/
	// 表达式1，2，3都可以省略

	//随机数
	/*
			1. 获取种子数
		       rand.Seed(time.Now().UnixNano())
			2. 获取随机数
		       num4:=rand.Intn(100) //[0,100)

			   rand.Intn(n-m+1)+m                    //取随机数
	*/
	//获取当前系统的时间

	nowT1 := time.Now() //time.Time
	/*时间戳：
	指定时间，距离1970年1月1日0点0分0秒，之间的时间差值：s,ns。秒，纳秒。
	*/
	timeStamp1 := nowT1.Unix()     // 获取秒的时间戳
	timeStamp2 := nowT1.UnixNano() // 获取纳秒的时间戳
	rand.Seed(timeStamp1)          // 获取随机数的种子数：int64的数值即可。
	rand.Seed(timeStamp2)          // 获取随机数的种子数：int64的数值即可。

	num1 := rand.Intn(10) // 获取一个随机数：[0,n)
	fmt.Println(num1)
	num2 := rand.Intn(10)
	fmt.Println(num2)
}

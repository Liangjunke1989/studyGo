package main

import (
	"fmt"
)

func main() {
	/*
		循环结束：
			循环条件不满足时，循环结束。
			但是可以通过break和continue,强制结束循环，无论循环条件是否成立，循环都结束。

		break和continue关键字
			break： 打断
				在for语句中，强制结束循环。             终止   out
			continue：继续
				只结束当前的这次循环，下次继续。         中止

	*/
	{
		//for i := 0; i < 10; i++ {
		//	if i == 6 {
		//		break //强制结束循环，循环结束了
		//	}
		//	fmt.Println(i)
		//}
	}
	{
		//for i:=0;i<10;i++ {
		//	if i==6 {
		//		continue                              //只结束当前的这次循环，下次继续。
		//		                                      //continue之后的不会执行，直接执行表达式3
		//	}
		//	fmt.Println(i)
		//}
		//fmt.Println("main...over...")
	}

	fmt.Println("------------------------练习一：打印输出1-100，能被3整除，但不能被5整除的数,只打印输出前5个-----------------------")
	{
		//count:=0
		//for i:=1;i<=100;i++{
		//	if count==5 {
		//		break                           //break， 可以提高效率
		//	}
		//	if i%3==0&&i%5!=0 {
		//		count++
		//		fmt.Println(i)
		//	}
		//}
	}

	fmt.Println("------------------------练习二：打印素数中，提高效率-----------------------")
	{
		//out:for i := 2; i < 100; i++ {
		//	flag:=true
		//	//验证
		//	for j := 2; j < i; j++{
		//		if i%j==0 {
		//			flag=false
		//			break out                   //贴标签的形式，跳出外层循环
		//			                            //只要满足，结束就近的for循环
		//										//只满足一个就可以跳出
		//		}
		//	}
		//	//结果
		//	if flag {
		//		fmt.Println(i,"是素数")
		//	}
		//}
	}

	fmt.Println("------------------------练习三：比较三个数，写法一：--------------------------------")
	{
		/*
			比较三个数，从小到大输出：
		*/

		//var a, b, c int
		//var tempMax, tempMin, tempMid int
		//
		////输入三个数
		//fmt.Println("请输入第一个数：")
		//fmt.Scanf("%d", &a)
		//fmt.Println("请输入第二个数：")
		//fmt.Scanf("%d", &b)
		//fmt.Println("请输入第三个数：")
		//fmt.Scanf("%d", &c)
		//fmt.Printf("%d,%d,%d", a, b, c)
		////依次比较
		////先取出最大值
		//tempMax = a
		//if a > tempMax {
		//	tempMax = a
		//}
		//if b > tempMax {
		//	tempMax = b
		//}
		//if c > tempMax {
		//	tempMax = c
		//}
		////再取出最小值
		//tempMin = c
		//if a < tempMin {
		//	tempMin = a
		//}
		//if b < tempMin {
		//	tempMin = b
		//}
		//if c < tempMin {
		//	tempMin = c
		//}
		////  fmt.Printf("最大值为：%d,最小值为：%d\n", tempMax, tempMin)
		////然后用排除法 找到中间值
		//if a != tempMin && a != tempMax {
		//	tempMid = a
		//}
		//if b != tempMin && b != tempMax {
		//	tempMid = b
		//}
		//if c != tempMin && c != tempMax {
		//	tempMid = c
		//}
		//fmt.Printf("从大到小依次为：%d,%d,%d", tempMax, tempMid, tempMin)
	}

	fmt.Println("------------------------练习三：比较三个数，写法二：--------------------------------")
	{
		/*
			比较三个数，从小到大输出：
		*/

		//var a, b, c int
		//
		////输入三个数
		//fmt.Println("请输入第一个数：")
		//fmt.Scanf("%d", &a)
		//fmt.Println("请输入第二个数：")
		//fmt.Scanf("%d", &b)
		//fmt.Println("请输入第三个数：")
		//fmt.Scanf("%d", &c)
		////fmt.Printf("%d,%d,%d", a, b, c)
		////依次比较
		////让a为最大值1
		//if b > a {
		//	a,b = b,a  //需要交换，否则a=b  a的值需要存放
		//}
		//if c > a {
		//	a,c = c,a
		//}
		////在比较b，c的大小,让b最大
		//if c > b {
		//	b,c = c,b
		//}
		//fmt.Printf("\n%d>=%d>=%d", a, b, c)
	}

	fmt.Println("------------------------练习四：剪刀石头布游戏：--------------------------------")
	{
		/*
			剪刀：a-1，石头：b-2，布：c-3
			思路：
				1.自己输入一个[1,3]的数
			    2.电脑产生一个[1,3]的随机数
				3.进行比较
		*/
		//var playerNum, PCNum int
		//var playerStr, PCStr, resultStr string
		////玩家输入的随机数
		//fmt.Println("请出一个（剪刀1/石头2/布3）：")
		//fmt.Scanf("%d", &playerNum)
		////2fmt.Printf("此时的playerNum值为：%d\n", playerNum)
		//switch playerNum {
		//case 1:
		//	playerStr = "剪刀"
		//	break
		//case 2:
		//	playerStr = "石头"
		//	break
		//case 3:
		//	playerStr = "布"
		//	break
		//default:
		//	playerStr = "您输入的数字有误，不在游戏范围内！！"
		//}
		//fmt.Printf("您出的是：%s\n", playerStr)
		//
		////电脑产生的随机数
		//rand.Seed(time.Now().UnixNano())
		//PCNum = rand.Intn(3) + 1
		//switch PCNum {
		//case 1:
		//	PCStr = "剪刀"
		//	break
		//case 2:
		//	PCStr = "石头"
		//	break
		//case 3:
		//	PCStr = "布"
		//	break
		//default:
		//	PCStr = "电脑出现bug！！"
		//}
		//fmt.Printf("电脑出的是：%s\n", PCStr)
		//
		////算法，游戏规则
		//switch {
		//case PCNum == 1 && playerNum == 2 || PCNum == 2 && playerNum == 3|| PCNum == 3 && playerNum == 1:
		//	resultStr = "玩家胜利！"
		//	break
		//case PCNum == 1 && playerNum == 1 || PCNum == 2 && playerNum == 2 || PCNum == 3 && playerNum == 3:
		//	resultStr = "你们出得一样哦，平局了！"
		//	break
		//default:
		//	resultStr = "电脑胜利"
		//}
		//
		////显示结果
		//fmt.Printf("结果为：%s\n", resultStr)

	}

	fmt.Println("------------------------练习五：斐波那契数列fibonacci(递归)：--------------------------------")
}

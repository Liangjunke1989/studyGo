package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		学习包的方式：
			授之以鱼不如授之以渔。
		time时间日期包：
			time.Format()
			yyyy,mm,dd
	*/
	fmt.Println("--------------------------------------1.获取当前的时间------------------------------------------")
	t1 := time.Now()
	fmt.Println(t1)
	fmt.Println(time.Now().Format(time.RFC850))
	tFormat := time.Now().Format("2006年1月2日 15：04：05") //模板需要指定日期 6-1-2-3-4-5
	fmt.Println(tFormat)
	fmt.Println()
	fmt.Println("---------------------------------------2.将字符串转成时间类型-------------------------------------")
	str := "2008年8月8日"
	t2, err := time.Parse("2006年1月2日", str)
	if err != nil {
		fmt.Println("解析有误，请重试！", err.Error())
	} else {
		fmt.Printf("%s获取到的时间类型的值为：%v", str, t2)
	}
	fmt.Println()
	fmt.Println("---------------------------------------3.时间戳-------------------------------------------------")
	//距离1970年1月1日0点0时0分0秒的时间差值 称为时间戳。
	//世界协调时间U.T.C.
	//格林威治标准时间G.M.T.
	t3 := time.Date(2012, 12, 21, 0, 0, 0, 0, time.UTC)
	unix01 := t3.Unix() //秒的时间戳
	fmt.Println(unix01)
	unixNano01 := t3.UnixNano() //毫秒的时间戳
	fmt.Println(unixNano01)
	fmt.Println()
	fmt.Println("-------------------------------------4.获取当前时间的指定内容的年，月，日---------------------------")
	year, month, day := t1.Date()
	hour, min, sec := t1.Clock()
	fmt.Println(year, month, day, hour, min, sec)
	fmt.Println("---------------------")
	fmt.Println(t1.Year())
	fmt.Println(t1.Month())
	fmt.Println(t1.Weekday())
	fmt.Println(t1.Day())
	fmt.Println(t1.Hour())
	fmt.Println(t1.Minute())
	fmt.Println(t1.Second())
	fmt.Println("---------------------")
	fmt.Println(t1.YearDay())
	fmt.Println(t1.ISOWeek()) //当前年份和第几个星期
	fmt.Println()
	fmt.Println("-------------------------------------5.其他时间增减日期-------------------------------------------")
	fmt.Println(t1.Add(time.Hour * 24))  //增加几小时后
	fmt.Println(t1.Add(-time.Hour * 24)) //增加几小时后
	duration := t1.Sub(t3)               //两个时间之间的差值
	fmt.Printf("%T,%v", duration, duration)
	fmt.Println()
	fmt.Println("-------------------------------------6.睡眠-------------------------------------------")
	time.Sleep(3 * time.Second) //当前的程序进行睡眠状态！
	fmt.Println("程序休眠结束11！！！！，继续执行！！！！")
	rand.Seed(time.Now().UnixNano())
	i := time.Duration(rand.Intn(10) + 1)
	time.Sleep(i * time.Second)
	fmt.Println(i, "程序休眠结束22！！！！，继续执行！！！！")
}

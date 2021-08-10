package main

import "fmt"

//1.定义一个结构体，表示自定义错误的类型

type errorRect struct {
	msg string //错误信息
	wid float64
	len float64
}

//重写Error()方法
func (e *errorRect) Error() string {
	//return fmt.Sprintf("宽度是：%.2f,长度是%.2f，%s",e.wid, e.len,e.msg)
	return e.msg
}

//2.定义一个函数，用于求矩形的面积
func getArea(wid, len float64) (float64, error) {
	errorMsg := ""
	if wid < 0 || len < 0 {
		errorMsg = "长/宽存在负数，输入有误！"
	}
	if errorMsg != "" {
		return 0, &errorRect{errorMsg, wid, len}
	}
	area := wid * len
	return area, nil //正确的话，没有错误返回值
}
func main() {
	/*
		自定义error：
			error是一个内置的接口interface，定义了一个方法。
				Error()string
	*/
	area, err := getArea(2, 99.7777)
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf("面积是：%.2f", area)
	}
}

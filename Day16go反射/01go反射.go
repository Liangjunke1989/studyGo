package main

import (
	"fmt"
	"reflect"
)

type People struct {
	Name    string
	Address string `DiZhi:"精确到到社区"` //反单引号标记的写法. 只要有标记存在，都是使用反射来获取。
}

func main() {
	/*
		反射：
			根据变量获取对应的变量的类型和变量的值。
			反射可以操作属性，也可以操作方法。
			只要有标记存在，都是使用反射来获取。
	*/
	a := "name"
	typeOf := reflect.TypeOf(a)   //获取变量的类型
	valueOf := reflect.ValueOf(a) //获取变量的值
	fmt.Println(typeOf, valueOf)
	fmt.Println("----------------------------01、获取结构体的类型和结构体的值-----------------------------------------")
	{
		peo := People{"smallMing", "ZhengZhou"}
		//获取peo的值    获取实例的值
		v := reflect.ValueOf(peo)
		//获取属性个数，如果v不是结构体类型 返回panic
		fmt.Println(v.NumField())
		//获取第0个属性，并转成string类型
		fmt.Println(v.Field(0).String())               //  Property属性，Field字段
		fmt.Println(v.FieldByIndex([]int{1}).String()) //根据索引，获取第一个属性的值
		//根据名字获取类型，并把类型名称转换成string类型
		fieldName := v.FieldByName("name")
		fmt.Println(fieldName.Kind().String()) //字段的类型
	}
	fmt.Println("----------------------------02、设置结构体中的属性的值----------------------------------------------")
	{
		Addr := "Address"
		peo02 := People{
			Name:    "LJK",
			Address: "八面神",
		}
		v02 := reflect.ValueOf(&peo02).Elem() //Elem()获取阵阵指向地址的封装。地址的值必须调用Elem()才可以继续操作。
		fmt.Println(v02.FieldByName(Addr).CanSet())
		fmt.Println("----------------------")
		peo03 := new(People)                 //返回的是指针
		v03 := reflect.ValueOf(peo03).Elem() //反射时获取的是peo的地址，地址的值必须调用Elem()才可以继续操作。
		fmt.Println(v03.FieldByName(Addr).CanSet())
		//需要修改属性（字段）的内容时，要求结构体中属性名首字母大写才可以被设置。   -------------核心注意点
		v03.FieldByName(Addr).SetString("谦祥万和城")
		fmt.Println(peo03)
		fmt.Println(peo03.Address)
		v03.FieldByName("Name").SetString("KOKO")
		fmt.Println(peo03.Name)
	}
	fmt.Println("----------------------------03、设置结构体中的标记--------------------------------------------------")
	t := reflect.TypeOf(People{
		"juhnko",
		"Locaris",
	})
	fmt.Println(t.FieldByName("Address"))
	_field, _ := t.FieldByName("Address")
	fmt.Println(_field.Tag)
	fmt.Println(_field.Tag.Get("DiZhi"))
	//fmt.Println(_field.Index)
}

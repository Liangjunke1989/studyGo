package main

import "fmt"

//1.定义一个书的结构体

type Book struct {
	bookName string
	price    float64
	author   string
}

//2.定义一个人的结构体

type Person2 struct {
	name string
	age  int
	book Book //自定义的结构体作为类型
}
type Person3 struct {
	name string
	age  int
	book *Book //自定义的结构体指针作为类型（指针类型） 指针高效利用内存
}

func main() {
	/*
			结构体的嵌套：
				模拟面向对象的类与类的两种关系：
		         -->聚合关系：一个类作为另一个类的属性。
					继承关系：一个类作为另一个类的子类。    子类--父类
	*/
	fmt.Println("-------------------1.简单的结构嵌套-----------------------------------")
	p1 := Person3{"koko", 30, &Book{"花花世界", 65.33, "ljk"}} //聚合关系
	fmt.Println(p1.book.author)

	fmt.Println("-------------------2.1进阶版的结构嵌套1-----------------------------------")
	stu1 := Student{"dk", 123, 21, []*Book{
		{"a", 12, "l"},
		{"b", 10, "ll"},
		{"c", 13, "zz"},
	}}
	fmt.Println(*stu1.books[2])
	fmt.Println()
	fmt.Println("-------------------2.2进阶版的结构嵌套2,利用切片-----------------------------------")
	b1 := make([]*Book, 0)
	b1 = append(b1, &Book{"a", 12.23, "l"}, &Book{"b", 10.2, "ll"})
	stu2 := Student{"name", 1, 11, b1}
	fmt.Println(stu2.books)
	for k, v := range stu2.books {
		fmt.Printf("第%d本书的书名为%s,价格为%f,作者为%s\n", k+1, v.bookName, v.price, v.author)
	}
	fmt.Println()
	fmt.Println("-------------------2.3进阶版的结构嵌套3，利用map-----------------------------------")
	map1 := make(map[string]*Book, 0)
	{
		map1["三国"] = &Book{"三国", 23.33, "罗贯中"}
		map1["水浒"] = &Book{"水浒", 32, "施耐庵"}
	}
	stu3 := Student2{"zz", 123, 21, map1}
	fmt.Println(stu3.booksMap["水浒"].price)
	fmt.Println(stu3.booksMap["三国"].author)

}

/*
	练习一：创建一个student类，book类
			一个学生对象可以有多本书。
*/
type Student struct {
	name  string
	id    int
	age   int
	books []*Book
}
type Student2 struct {
	name     string
	id       int
	age      int
	booksMap map[string]*Book
}

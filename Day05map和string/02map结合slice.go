package main

import "fmt"

func main() {
	fmt.Println("------------------------------1.创建map，存放个人信息--------------------------")
	//创建数据源
	var map1, map2, map3 map[string]string
	{
		map1 = make(map[string]string)
		{
			map1["id"] = "1"
			map1["name"] = "ljk"
			map1["age"] = "32"
			map1["sex"] = "男"
			map1["address"] = "xxxx"
		}
		map2 = make(map[string]string)
		{
			map2["id"] = "2"
			map2["name"] = "juhnko"
			map2["age"] = "31"
			map2["sex"] = "男"
			map2["address"] = "xxxxxx"
		}
		map3 = map[string]string{
			"id":      "3",
			"name":    "dk",
			"age":     "30",
			"sex":     "女",
			"address": "yyyyy",
		}
	}

	fmt.Println("------------------------------2.将map存储到数组中--------------------------")
	{
		//arr1:=[...]int{1,2,3}
		arr1 := [...]map[string]string{map1, map2, map3}
		fmt.Println("map存储到数组中的结果为：", arr1)
	}

	fmt.Println("------------------------------3.将map存储到slice中,并遍历出来------- -------")
	{
		s1 := make([]map[string]string, 0, 3) //创建初始的切片
		s1 = append(s1, map1, map2, map3)
		//fmt.Println("map存储到切片中的结果为：",s1)

		//遍历切片
		for k, v := range s1 {
			fmt.Printf("第%d个人的信息:\n", k+1)
			if v != nil {
				fmt.Printf("\tid为：%s,姓名为：%s,年龄为：%s,性别为：%s,地址为：%s\n",
					v["id"], v["name"], v["age"], v["sex"], v["address"])
			}
		}
	}
}

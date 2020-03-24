package main

import (
	"fmt"
	"sort"
)

func main() {
	//map的使用
	var a map[string]string
	//在使用map前要声明make，key不能重复,value可以重复
	a = make(map[string]string, 10) //先写map的类型，再写分配多少个空间
	a["cmh01"] = "cmh"
	fmt.Println(a)

	student := make(map[string]map[string]string)
	student["student_1"] = make(map[string]string, 2)
	student["student_1"]["name"] = "cmh"
	student["student_1"]["sex"] = "man"

	student["student_2"] = make(map[string]string, 2)
	student["student_2"]["name"] = "cmh2"
	student["student_2"]["sex"] = "woman"

	student["student_3"] = make(map[string]string, 2)
	student["student_3"]["name"] = "cmh3"
	student["student_3"]["sex"] = "man"

	fmt.Println(student)
	fmt.Println()
	//map的增删查改
	student["student_2"]["name"] = "slb"
	//delete(student,"student_3")
	//一次性删除所有的key
	//1.使用遍历
	//2.直接make一个新空间
	//student = make(map[string]string)
	val, findRes := student["student_1"]
	if findRes {
		fmt.Printf("有 %v\n", val)
	} else {
		fmt.Printf("无 %v\n", val)
	}

	//map的遍历，只能用for-range来写
	for key, val1 := range student {
		fmt.Println("key:", key)
		for key1, val2 := range val1 {
			fmt.Println("遍历为：, 结果为：\n", key1, val2)
		}
	}

	//map切片的使用
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "齐天大圣"
		monsters[0]["age"] = "180000"
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "入海大圣"
		monsters[1]["age"] = "170000"
	}

	if monsters[2] == nil {
		monsters[2] = make(map[string]string, 2)
		monsters[2]["name"] = "平天大圣"
		monsters[2]["age"] = "160000"
	}

	if monsters[3] == nil {
		monsters[3] = make(map[string]string, 2)
		monsters[3]["name"] = "猕猴大圣"
		monsters[3]["age"] = "150000"
	}

	//这里可以使用切片的append函数，可以以动态的增加
	monster := map[string]string{
		"name": "冲云大圣",
		"age":  "130000",
	}
	monsters = append(monsters, monster)

	fmt.Println(monsters)
	fmt.Println()

	//map排序说明
	map1 := make(map[int]int, 10) //无序排序
	map1[0] = 100
	map1[1] = 26
	map1[5] = 52
	map1[3] = 25
	map1[4] = 65
	map1[6] = 54
	map1[7] = 87
	map1[9] = 98
	map1[8] = 50
	fmt.Println(map1)
	fmt.Println()
	//如何按照key的顺序进行排列输出
	var jk []int //定义切片
	for k1, _ := range map1 {
		jk = append(jk, k1) //将map1中的下标加到切片中
	}

	sort.Ints(jk) //排序，sort中的Ints方法，递增排序
	fmt.Println(jk)

	for _, k2 := range jk {
		fmt.Println("map = ", k2, map1[k2])
	}

	students := make(map[string]Stu, 10)
	stu_1 := Stu{"cmh", 20, "泰安"}
	stu_2 := Stu{"bllh", 19, "福州"}
	students["stu1"] = stu_1
	students["stu2"] = stu_2

	fmt.Println(students)

	//练习题
	users := make(map[string]map[string]string, 10)
	users["tom"] = make(map[string]string, 2)
	users["tom"]["pwd"] = "8888888"
	users["tom"]["nickname"] = "t"

	users["ross"] = make(map[string]string, 2)
	users["ross"]["pwd"] = "777777"
	users["ross"]["nickname"] = "r"

	users["cat"] = make(map[string]string, 2)
	users["cat"]["pwd"] = "9999999"
	users["cat"]["nickname"] = "c"

	modifyUser(users, "cmh", "c")
	modifyUser(users, "bllh", "b")
	modifyUser(users, "slb", "s")
	fmt.Println(users)

}

//map为引用类型
//定义一个学生结构体（体验一下）
type Stu struct {
	Name    string
	Age     int
	Address string
}

func modifyUser(users map[string]map[string]string, name string, nickname string) {
	if users[name] != nil {
		//有这个用户
		users[name]["pwd"] = "135672654"
	} else {
		//没有
		users[name] = make(map[string]string, 2)
		users[name]["pwd"] = "135672654"
		users[name]["nickname"] = "cmh"
	}
}

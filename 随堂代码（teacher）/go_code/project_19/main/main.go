package main

import (
	"fmt"
	//"encoding/json"
)

type Cat struct {
	Name  string
	Age   int
	Color string
}

//如果结构体的字段类型为指针，切片跟map，需要先make,才能使用

type Person struct {
	Name_1 string
	Age_1  int
	Scores [5]float64
	ptr    *int
	slice  []int
	map1   map[string]string
}

func main() {
	var cat_1 Cat
	cat_1.Name = "小白"
	cat_1.Age = 3
	cat_1.Color = "花色"

	fmt.Println("cat的信息为:", cat_1)

	var p_1 Person
	//使用
	p_1.slice = make([]int, 10)
	p_1.slice[0] = 170

	p_1.map1 = make(map[string]string)
	p_1.map1["name"] = "cmh"

	fmt.Println(p_1)

	var person_1 *Person = new(Person)
	(*person_1).Name_1 = "yyy"
	person_1.Name_1 = "yyy~"
	person_1.Age_1 = 20
	fmt.Println(*person_1)

	var person_2 *Person = &Person{}
	(*person_2).Name_1 = "yyy~~"
	person_2.Name_1 = "yyy~~~"
	person_2.Age_1 = 203
	fmt.Println(*person_2)

	var person_3 *Cat = &Cat{"bllh", 24, "red"}
	fmt.Println(*person_3)

	//	monster := Monster{"齐天大圣",888888,"金箍棒"}
	//	jsonMonster ,err := json.Marshal(monster)
	//	if err != nil{
	//		fmt.Println("错误")
	//	}
	//	fmt.Println("jsonMonster",string(jsonMonster))
}

//type Monster struct{
//	Name__1 string `json: "name"`//tag
//	Age__1 int `json: "age"`
//	Skill string `json: "skill"`//使用了反射
//}

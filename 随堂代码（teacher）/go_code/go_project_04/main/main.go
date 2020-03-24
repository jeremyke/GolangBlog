package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string `json:"name"` //放一个标签
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

func testStruct() {
	monster := Monster{
		Name:     "齐天大圣",
		Age:      520520,
		Birthday: "8520-2-23",
		Sal:      18000.0,
		Skill:    "千手浮屠",
	}

	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("序列化错误")
	}
	fmt.Println(string(data))
}

func testMap() {
	//定义一个map
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "杀心观音"
	a["age"] = 566666
	a["address"] = "光影天宫"

	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println("序列化错误")
	}
	fmt.Println(string(data))
}

func testSlice() {
	var silce []map[string]interface{}
	var m_1 map[string]interface{}
	m_1 = make(map[string]interface{})
	m_1["name"] = "如来"
	m_1["age"] = "566666"
	m_1["address"] = "光影天宫"
	silce = append(silce, m_1)

	var m_2 map[string]interface{}
	m_2 = make(map[string]interface{})
	m_2["name"] = "帝释天"
	m_2["age"] = "566666"
	m_2["address"] = "暗影天宫"
	silce = append(silce, m_2)

	data, err := json.Marshal(silce)
	if err != nil {
		fmt.Println("序列化错误")
	}
	fmt.Println(string(data))
}

func main() {
	//将map,结构体，切片进行序列化
	testStruct()
	testMap()
	testSlice()
}

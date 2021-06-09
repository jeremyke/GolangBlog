package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Birthday string
	Sal      float64
	skill    string
}

func testStruct() {
	monster := Monster{
		Name:     "蒋冬莲",
		Age:      24,
		Birthday: "0703",
		Sal:      8000.0,
		skill:    "测试",
	}
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化失败，%v", err)
	} else {
		fmt.Printf("序列化后结果为：%v\n", string(data))
	}
}

func testMap() {
	var a map[string]interface{} //表示k为string,v为任意类型
	a = make(map[string]interface{})
	a["name"] = "呵呵"
	a["age"] = 12
	a["address"] = "中国"
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化失败，%v", err)
	} else {
		fmt.Printf("序列化后结果为：%v\n", string(data))
	}

}

func testSlice() {
	var stringSlice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "蒋冬莲"
	m1["sex"] = "女"
	m1["address"] = "珠海"
	fmt.Printf("stringSlice的值为%v,类型为%T,地址是%p\n", stringSlice, stringSlice, stringSlice)
	//stringSlice[0] = m1 stringSlice没有make,只是申明了，但没有分配内存
	//但是为什么append能作用在stringSlice呢？因为append会在每次追加的时候重新分配内存
	stringSlice = append(stringSlice, m1)
	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "柯丽萍"
	m2["sex"] = "女"
	m2["address"] = "深圳"
	stringSlice = append(stringSlice, m2)
	data, err := json.Marshal(stringSlice)
	if err != nil {
		fmt.Printf("序列化失败，%v", err)
	} else {
		fmt.Printf("序列化后结果为：%v\n", string(data))
	}
}

func main() {
	//结构体，map,切片序列化
	testStruct()
	testMap()
	testSlice()

}

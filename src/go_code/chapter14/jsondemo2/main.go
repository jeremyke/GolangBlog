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
	str := "{\"name\":\"蒋冬莲\",\"age\":24,\"Birthday\":\"0703\",\"Sal\":8000}"
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("反序列化失败，%v", err)
	} else {
		fmt.Printf("struct反序列化后结果为：%v\n", monster)
	}
}

func testMap() {
	str := "{\"address\":\"中国\",\"age\":12,\"name\":\"呵呵\"}"
	var a map[string]interface{} //反序列化操作map不需要make操作,在反序列化是会自动make
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("map反序列化失败，%v", err)
	} else {
		fmt.Printf("map序列化后结果为：%v\n", a)
	}

}

func testSlice() {
	str := "[{\"address\":\"珠海\",\"name\":\"蒋冬莲\",\"sex\":\"女\"},{\"address\":\"深圳\",\"name\":\"柯丽萍\",\"sex\":\"女\"}]"

	var stringSlice []map[string]interface{}

	err := json.Unmarshal([]byte(str), &stringSlice)
	if err != nil {
		fmt.Printf("slice反序列化失败，%v", err)
	} else {
		fmt.Printf("slice反序列化后结果为：%v\n", stringSlice)
	}
}

func main() {
	//结构体，map,切片序列化
	testStruct()
	testMap()
	testSlice()

}

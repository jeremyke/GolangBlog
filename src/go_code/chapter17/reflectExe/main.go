package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Student struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
	Age  int    `json:"age"`
}

func (s Student) DoHomework(as string) {
	fmt.Println(s.Name + "is doing " + as)
}

func reflectTest(p interface{}) {
	//Go语言程序中对指针获取反射对象时，可以通过 reflect.Elem() 方法获取这个指针指向的元素类型

	//获取reflect.Type 类型
	typ := reflect.TypeOf(p)
	//获取reflect.Value 类型
	val := reflect.ValueOf(p)

	//字段个数
	fieldNum := val.Elem().NumField()
	fmt.Println("fieldNum=", fieldNum)

	for i := 0; i < fieldNum; i++ {
		fmt.Println(typ.Elem().Field(i).Name + "的标签是：" + typ.Elem().Field(i).Tag.Get("json"))
		fmt.Println(typ.Elem().Field(i).Name + "的值是：" + val.Elem().Field(i).String())
	}

	val.Elem().Field(0).SetString("Esther")
	val.Elem().Field(1).SetString("女")
	fmt.Println(typ.Elem().Field(0).Name + "的值是：" + val.Elem().Field(0).String())

	funcNum := val.Elem().NumMethod()
	fmt.Println("一共有" + strconv.Itoa(funcNum) + "个方法")

	var param []reflect.Value
	param = append(param, reflect.ValueOf("English"))
	val.Elem().Method(0).Call(param)
}

func main() {
	aa := Student{
		"keke",
		"男",
		20,
	}
	reflectTest(&aa)
	fmt.Println(aa)
}

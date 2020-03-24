package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string
	Age   int
	Score float32
	Sex   string
}

//方法，显示S的值
func (s Monster) Print() {
	fmt.Println("---start---")
	fmt.Println(s)
	fmt.Println("---end---")
}

//返回两个数的和
func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

//给monster赋值
func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func TestStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("退出")
		return
	}

	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num)
	for i := 0; i < num; i++ { //遍历字段
		tagVal := typ.Field(i).Tag.Get("json") //他的tag名称
		if tagVal != "" {
			fmt.Printf("Fsield %d: tage = %v\n", i, tagVal)
		}
	}
	numOfMethod := val.NumMethod()
	fmt.Printf("struct  has %d methods \n", numOfMethod)
	val.Method(1).Call(nil) //获取第二个方法
	//调用结构体的第一个方法Method[0]
	//按照函数的名字来排序的，所以自己写的第一个方法不一定是第一个方法
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params)
	fmt.Println("res= ", res[0].Int())
}

func main() {
	var a Monster = Monster{
		Name:  "齐天大圣",
		Age:   6500,
		Score: 99.99,
	}
	TestStruct(a)
}

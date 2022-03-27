package main

import (
	"fmt"
	"reflect"
)

type Stu struct {
	Name string
	Age  int
}

func reflectTest(b interface{}) {
	//获取reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType) //Stu

	//获取reflect.Value
	rValue := reflect.ValueOf(b)
	fmt.Printf("rValue=%v 类型为：%T\n", rValue, rValue) //rValue={tom 20} 类型为：reflect.Value

	//获取变量对应的类别常量rValue.kind() 或者 rType.kind()
	kind1 := rType.Kind()                           //struct
	kind2 := rValue.Kind()                          //struct
	fmt.Printf("kind1=%v kind2=%v\n", kind1, kind2) //struct

	//将reflect.Value-->interface{}
	iv := rValue.Interface()
	fmt.Printf("iv=%v 类型是：%T\n", iv, iv)

	//interface{}-->原来的类型
	student2 := iv.(Stu)
	fmt.Printf("name=%v,age=%v", student2.Name, student2.Age)
}

func main() {
	student := Stu{
		Name: "tom",
		Age:  20,
	}
	reflectTest(student)
}

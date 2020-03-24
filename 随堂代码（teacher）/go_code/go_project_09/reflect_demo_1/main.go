package main

import (
	"fmt"
	"reflect"
)

func reflectText_01(b interface{}) {
	//通过反射获取的传入变量的 type kink 值
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp:", rTyp)
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal:", rVal)
	//n_1 := 10
	n_2 := 1 + rVal.Int()
	fmt.Println("n_2:", n_2)
	//下面我们将rVal转成 interface{}
	iV := rVal.Interface()
	num_1 := iV.(int) //类型断言
	fmt.Println("num_1:", num_1)
}

func reflectText_02(b interface{}) {
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp:", rTyp)
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal:", rVal)
	//n_1 := 10
	//n_2 := 1 + rVal.Int()
	//fmt.Println("n_2:",n_2)
	//下面我们将rVal转成 interface{}
	iV := rVal.Interface()
	fmt.Printf("iV = %v, iV type = %T\n", iV, iV)
	//这里我们简单的使用一类带检测的类型断言
	stu, ok := iV.(Student)
	if ok {
		fmt.Println("stu.Name = ", stu.Name)
	}

	//获取变量对应的Kind
	//rVal.kind() 跟 rTyp.kind()是一样的，指向的值都是相同的
	fmt.Printf("kind = %v kind = %v", rVal.Kind(), rTyp.Kind())

}

type Student struct {
	Name string
	Age  int
}

type Monster struct {
	Name string
	Age  int
}

func main() {
	var num int = 100
	reflectText_01(num)

	stu := Student{
		Name: "tom",
		Age:  12,
	}
	reflectText_02(stu)
}

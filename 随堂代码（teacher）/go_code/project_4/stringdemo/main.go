package main

import "fmt"

func main() {
	var address string = "崔默涵，hello world"
	fmt.Println(address)
	var str = "hello"
	//str[0] = 'a'  在go中的字符串不可变
	str2 := "sd\nabd"
	fmt.Println(str)
	fmt.Println(str2)
	//反引号为`````
	var a int                                                                      //0
	var b float32                                                                  //0
	var c float64                                                                  //0
	var isMarryied bool                                                            //false
	var name string                                                                //""
	fmt.Println("a=%d,b=%f,c=%f,isMarryied=%v,name=%v", a, b, c, isMarryied, name) //%v表示变量的原始值输出
	//T（v）将值v转换为类型T
	var i int32 = 100
	var n1 float32 = float32(i)
	fmt.Println("i = %v n1 = %v", i, n1)
	var f int64
	var g int32 = 64
	//var h int8
	f = int64(g) + 20
	fmt.Println("f =", f)

}

package main

import "fmt"

//变量的三步曲（申明，赋值，使用）
/*func main()  {
	var i int
	i = 20
	fmt.Println("i=",i)
}*/

//多变量申明
var (
	name = "tom"
	sex  = "male"
	age  = 25
)

func main() {
	//第一种
	var a, b, c int
	//第二种
	var j, k, l = 100, "tom", 2.6
	//第三种
	m, n := "i", 600
	fmt.Println("a=", a, "b=", b, "c=", c)
	fmt.Println("j=", j, "k=", k, "l=", l)
	fmt.Println("m=", m, "n=", n)
	fmt.Println("name=", name, "sex=", sex, "age=", age)
}

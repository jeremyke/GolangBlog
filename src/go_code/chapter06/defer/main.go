package main

import "fmt"

func main() {
	res := sum(10, 20)
	fmt.Println("res=", res)
}

//ok3 res=30
//ok2 n2=20
//ok1 n1=10
//res res=30

func sum(n1, n2 int) int {
	//当执行defer时，其后面的语句会压入defer栈中；当函数执行完毕后再执行，先入后出执行。
	defer fmt.Println("ok1 n1=", n1)
	defer fmt.Println("ok2 n2=", n2)
	res := n1 + n2
	fmt.Println("ok3 res=", res)
	return res
}

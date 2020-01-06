package main

import "fmt"

func main() {
	//var ages map[string]int //申明变量为map类型为nil
	//ages := make(map[string]int) //申明变量为map类型且不为nil
	ages := map[string]int{} //申明变量为map类型且为空
	//ages["a"] = 77
	fmt.Println(len(ages))
	fmt.Println(ages)

}

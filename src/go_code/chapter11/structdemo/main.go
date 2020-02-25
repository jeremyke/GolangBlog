package main

import (
	"fmt"
)

type Person struct {
	Name   string
	Age    int
	Score  [5]float64
	ptr    *int
	slice1 []int
	map1   map[string]string
}

type Monster struct {
	Name string
	Age  int
}

func main() {
	var klp Person
	//使用slice需要先make
	klp.slice1 = make([]int, 10)
	//使用map也需要先make
	klp.map1 = make(map[string]string)
	klp.slice1[0] = 100
	klp.map1["name"] = "klp"
	fmt.Println(klp)

	var monster1 Monster
	monster1.Name = "牛魔王"
	monster1.Age = 200
	monster2 := &monster1
	monster2.Name = "ss"

	fmt.Println("monster1=", monster1)
	fmt.Println("monster2=", *monster2)
}

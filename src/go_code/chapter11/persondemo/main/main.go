package main

import (
	"fmt"
	"go_code/chapter11/persondemo/person"
)

func main() {
	p1 := person.NewPerson("tom", 25, 8000.0)
	fmt.Println(*p1)
	p1.SetAge(29)
	fmt.Println(*p1)
	fmt.Println(p1.GetAge())
}

package main

import (
	"fmt"
)

type A struct {
	Name string
	Age  int
}

type B struct {
	A
}

func (a *A) SayOk() {
	fmt.Println("A SayOk", a.Name)
}

func (a *A) hello() {
	fmt.Println("A hello", a.Name)
}

func main() {
	var b B
	b.A.Name = "cmh"
	b.A.Age = 19
	b.SayOk()
	b.A.hello()
}

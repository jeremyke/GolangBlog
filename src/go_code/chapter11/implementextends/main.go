package main

import (
	"fmt"
)

type Monkey struct {
	Name string
}

func (this *Monkey) climb() {
	fmt.Println(this.Name, "天生会爬树...")
}

type Bird interface {
	Fly()
}

type SonMonkey struct {
	Monkey
}

func (this *SonMonkey) Fly() {
	fmt.Println(this.Name, " 通过实现，会飞翔...")
}

func main() {
	xiaoMonkey := SonMonkey{Monkey{Name: "孙悟空"}}
	xiaoMonkey.climb()
	xiaoMonkey.Fly()
}

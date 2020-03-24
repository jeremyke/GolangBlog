package main

import (
	"fmt"
)

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
}

type Computer struct {
	Usb
}

func (p Phone) Start() {
	fmt.Println("手机接口开始工作")
}

func (p Phone) Stop() {
	fmt.Println("手机接口结束工作")
}

func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

func main() {
	var c = Computer{}
	var p = Phone{}

	c.Working(p)
}

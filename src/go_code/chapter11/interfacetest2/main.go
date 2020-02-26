package main

import (
	"fmt"
)

//声明/定义一个接口
type Usb interface {
	//声明了两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
	Name string
}

//让Phone 实现 Usb和usb2接口的方法(同时实现2个接口)
func (p Phone) Start() {
	fmt.Println("手机开始工作。。。")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作。。。")
}

type Camera struct {
	Name string
}

//让Camera 实现   Usb接口的方法
func (c Camera) Start() {
	fmt.Println("相机开始工作~~~。。。")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作。。。")
}

func main() {
	var usrArr [3]Usb //多态数组
	usrArr[0] = Phone{"vivo"}
	usrArr[1] = Phone{"mi"}
	usrArr[2] = Camera{"尼康"}
	fmt.Println(usrArr)

}

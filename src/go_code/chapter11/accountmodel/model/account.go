package model

import (
	"fmt"
)

type Account struct {
	name string
	yue  float64
	pwd  [6]byte
}

func (a *Account) SetName(name string) {
	if len(name) < 6 || len(name) > 10 {
		fmt.Println("账号名有误")
		return
	}
	a.name = name
}
func (a *Account) SetYue(yue float64) {
	if yue <= 20.0 {
		fmt.Println("余额有误。。")
		return
	}
	a.yue = yue
}
func (a *Account) SetPwd(pwd [6]byte) {
	a.pwd = pwd
}

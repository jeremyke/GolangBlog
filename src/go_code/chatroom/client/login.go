package main

import "fmt"

func login(userId int, userPwd string) (err error) {
	fmt.Printf("你输入的id是%v,密码%v\n", userId, userPwd)
	return nil
}

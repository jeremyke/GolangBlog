package main

import "fmt"

var userId int
var userPwd string

func main() {
	var key int
	var loop = true
	for loop {
		fmt.Println("--------欢迎登录多人聊天室-----------")
		fmt.Println("          1 登录聊天室")
		fmt.Println("          2 注册用户")
		fmt.Println("          3 退出系统")
		fmt.Println("请输入（1-3）：")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}

	if key == 1 {
		fmt.Println("请输入你的ID")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入你的密码")
		fmt.Scanf("%s\n", &userPwd)
		login(userId, userPwd)
		/*fmt.Println("--------欢迎登录多人聊天室-----------")
		fmt.Println("          1 登录聊天室")
		fmt.Println("          2 注册用户")
		fmt.Println("          3 退出系统")*/
	} else if key == 2 {
		fmt.Println("注册...")
	}
}

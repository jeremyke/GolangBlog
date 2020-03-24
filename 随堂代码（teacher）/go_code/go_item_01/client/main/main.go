package main

import (
	"fmt"
	"go_code/go_item_01/client/process"
)

func main() {
	//接收用户的选择
	var key int
	//判断是否显示菜单
	//var loop = true
	//定义用户的ID，表示用户的密码
	var user_ID int
	var user_Pwd string
	var user_Name string
	for {

		fmt.Println("欢迎使用多人PY聊天室 \n\n")
		fmt.Println("\t\t\t 1.登陆")
		fmt.Println("\t\t\t 2.注册")
		fmt.Println("\t\t\t 3.退出")
		fmt.Println("\t\t\t 请选择（1~3）")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			fmt.Println("请输入你的ID号：")
			fmt.Scanf("%d\n", &user_ID)
			fmt.Println("请输入密码：")
			fmt.Scanf("%s\n", &user_Pwd) //不带\n电脑会认为回车就是你所要输入的东西，就像回车就是密码
			up := &process.UserProcess{}
			up.Login(user_ID, user_Pwd)
			//loop = false
		case 2:
			fmt.Println("注册聊天室账号")
			fmt.Println("请输入用户ID:")
			fmt.Scanf("%d\n", &user_ID)
			fmt.Println("请输入用户密码:")
			fmt.Scanf("%s\n", &user_Pwd)
			fmt.Println("请输入用户的昵称:")
			fmt.Scanf("%s\n", &user_Name)
			//调用一个实例，完成注册请求
			up := &process.UserProcess{}
			up.Login(user_ID, user_Pwd)
			//loop = false
		case 3:
			fmt.Println("退出聊天室")
			//loop = false
			break
		default:
			fmt.Println("你的输入被怪兽吃掉了，请再PY一下吧")

		}
	}
	/*fmt.Println("loop 为：", loop)
	//根据用户输入显示二级菜单
	if key == 1 {
		//要登陆
		fmt.Println("请输入你的ID号：")
		fmt.Scanf("%d\n", &user_ID)
		fmt.Println("请输入密码：")
		fmt.Scanf("%s\n", &user_Pwd)//不带\n电脑会认为回车就是你所要输入的东西，就像回车就是密码
		err := login(user_ID, user_Pwd)
		if err != nil {
			fmt.Println("登陆失败")
		}else {
			fmt.Println("登陆成功")
		}
	}else if key ==2 {
		fmt.Println("进行注册")
	}	*/
}

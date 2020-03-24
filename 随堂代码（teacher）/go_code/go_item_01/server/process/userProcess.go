package process2

import (
	"encoding/json"
	"fmt"
	"go_code/go_item_01/common/message"
	"go_code/go_item_01/server/model"
	"go_code/go_item_01/server/utils"
	"net"
)

type UserProcess struct {
	Conn net.Conn
	//增加一个字段，表示该Conn是哪一个用户的
	UserId int
}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	//1.先从mes中取出data,并直接反序列化成LoginMes
	var registerMes message.RegisterMes
	json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err =", err)
		return
	}
	//先声明一个resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	//之后声明一个LoginResMes,并赋值
	var RegisterResMes message.RegisterResMes
	//我们需要到redis数据库去完成验证注册
	//使用model.MyUserDao 到redis去验证
	user, err := model.MyUserDao.Register(registerMes.User)
	if err != nil {
		if err == ERROR_USER_NOT_EXISTS {
			registerResMes.Code = 404
			registerResMes.Error = ERROR_USER_NOT_EXISTS.Error()
		} else {
			registerResMes.Code = 505
			registerResMes.Error = "注册发生未知错误"
		}
	} else {
		registerResMes.Code = 200
		//这里用户登陆成功，我们把登陆成功后的用户放入userMgr
		//将登陆成功的userId赋给this
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)
		//通知其他在线的用户
		this.NotifyOthersOnlineUser(loginMes.UserId)
		//将当前在线用户的ID放到loginResMes.UserId
		//遍历userMgr.onlineUsers
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}

		fmt.Println("登陆成功")
	}

	//3将其序列化
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	//4将data付给resMes
	resMes.Data = string(data)
	//5对resMes进行序列化,准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	//6发送data 我们将其封装
	//因为使用了分层模式，我们创建一个Transfer实力，然后读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}

//编写一个函数serverProcessLogin函数，专门处理登陆请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//1.先从mes中取出data,并直接反序列化成LoginMes
	var loginMes message.LoginMes
	json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err =", err)
		return
	}
	//先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//之后声明一个LoginResMes,并赋值
	var LoginResMes message.LoginResMes

	//我们需要到redis数据库去完成验证
	//使用model.MyUserDao 到redis去验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		LoginResMes.Code = 404
	} else {
		LoginResMes.Code = 200
	}

	//合法用户
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		//合法

		LoginResMes.Code = 200 //200状态码，表示用户合法
	} else {
		//不合法
		LoginResMes.Code = 404 //404状态码，表示用户不存在
	}
	//3将LoginResMes序列化
	data, err := json.Marshal(LoginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	//4将data付给resMes
	resMes.Data = string(data)
	//5对resMes进行序列化,准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	//6发送data 我们将其封装
	//因为使用了分层模式，我们创建一个Transfer实力，然后读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}

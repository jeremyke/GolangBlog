package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/go_item_01/common/message"
	"go_code/go_item_01/server/utils"
	"net"
	//"go_code/go_item_01/client/process"
	"go_code/go_item_01/client/model"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int //增加一个字段表示是哪一个用户
}

//这里我们编写以哦个通知所有在线的用户的方法
//userId表示通知其他人我上线了
func (this *UserProcess) NotifyOthersOnlineUser(userId int) {
	//遍历onlineUsers,然后一个一个的发送给NotifyUserStatusMes
	for id, up := range userMgr.olineUsers {
		//过滤掉自己
		if id == userId {
			continue
		}
		//开始通知，单独写一个方法
		up.NotifyMeOnline(userId)
	}
}

func (this *UserProcess) NotifyMeOnline(userId int) {
	//组装NotifyUserStatusMes
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOL
	//将notifyUserStatusMes序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("序列化出错！err :", err)
		return
	}
	//将序列化后的notifyUserStatusMes赋给mes.Data
	mes.Data = string(data)
	//对mes再次序列化，准备发送
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化出错！err :", err)
		return
	}
	//发送，创建一个transfer实例
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyUserStatusMes err = ", err)
		return
	}

}

func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	//开始制定协议
	//fmt.Printf("user_ID = %d user_Pwd = %s\n", user_ID, user_Pwd)
	//return nil

	//1.连接到服务器端
	conn, err := net.Dial("tcp", "localhost: 8889")
	if err != nil {
		fmt.Println("net.Dial失败, err:", err)
		return
	}
	//延时关闭
	defer conn.Close()
	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginResMesType
	//3.创建一个loginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	//4.将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}
	//5.把data赋给mes.Data字段
	mes.Data = string(data)
	//6.将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}
	//7.data就是发送的数据,先发送长度，获取到data的长度，转换成一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//现在发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("发送失败", err)
		return
	}
	fmt.Println("发送消息的长度成功")

	return

	n, err = conn.Write(data)
	if n != 4 || err != nil {
		fmt.Println("发送失败", err)
		return
	}
	//这里还需要处理服务器端返回的信息
	//创建一个transfer实例
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg is err :", err)
		return
	}
	//将mes.Data反序列化
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		//初始化CurUser
		CurUser.Conn = conn
		CurUser.UserId = userId

		//fmt.Println("登陆成功")
		//可以显示当前在线用户列表，遍历loginResMes.UsersId
		fmt.Println("当前在线用户列表如下：")
		for _, v := range loginResMes.UsersId {
			//如果我们要求不显示自己在线，增加如下代码：
			if v == userId {
				continue
			}
			fmt.Printf("用户Id :%t\n", v)
			//完成客户端的olineUsers的初始化
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOL,
			}
			onlineUsers[v] = user
		}
		//这时我们还要启动一个协程在客户端
		//该协程保持和服务器的通讯，如果服务器有数据推送给客户端
		//然后显示在客户端
		go ServerProcessMes(conn)
		for {
			ShowMenu()
		}
	} else if loginResMes.Code == 404 {
		fmt.Println(loginResMes.Error)
	}
	return
}

func (this *UserProcess) Register(user_ID int, user_Pwd string, user_Name string) (err error) {

	//1.连接到服务器端
	conn, err := net.Dial("tcp", "localhost: 8889")
	if err != nil {
		fmt.Println("net.Dial失败, err:", err)
		return
	}
	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.RegisterMesType
	//3.创建一个registerMes结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = user_ID
	registerMes.User.UserPwd = user_Pwd
	registerMes.User.UserName = user_Name
	//4.将registerMes序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}
	//5.把data赋给mes.Data字段
	mes.Data = string(data)
	//6.将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}
	//创建一个Transfer实例
	tf := &utils.Transfer{
		Conn: conn,
	}
	//发送data给服务器
	tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息错误 ", err)
		return
	}
	mes, err = tf.ReadPkg()

	if err != nil {
		fmt.Println("readPkg(conn) is err :", err)
		return
	}
	//将mes.Data反序列化
	var RegisterResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &RegisterResMes)
	if RegisterResMes.Code == 200 {
		fmt.Println("登陆成功")
		go ServerProcessMes(conn)
		for {
			ShowMenu()
		}
	} else if RegisterResMes.Code == 406 {
		fmt.Println(RegisterResMes.Error)
	}
	return
}

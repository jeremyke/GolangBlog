package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/go_item_01/common/message"
	"go_code/go_item_01/server/utils"
	"net"
)

func login(userId int, userPwd string) (err error) {
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
	tf := &utils.Transfer{
		Conn: conn,
	}
	//这里还需要处理服务器端返回的信息
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg is err :", err)
		return
	}
	//将mes.Data反序列化
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登陆成功")
	} else if loginResMes.Code == 404 {
		fmt.Println(loginResMes.Error)
	}
	return
}

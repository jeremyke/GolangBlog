package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	common "go_code/chatroom/common/message"
	"net"
)

func login(userId int, userPwd string) (err error) {
	//链接到服务器
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()
	//准备通过Conn发送消息
	var mes common.Message
	mes.Type = common.LoginMesType

	loginMes := common.LoginMes{
		UserId:  userId,
		UserPwd: userPwd,
	}
	//序列化data
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("Marshal err=", err)
		return
	}
	mes.Data = string(data)
	//将mes序列化
	mesData, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("Marshal err=", err)
		return
	}
	var pkgLen uint32
	pkgLen = uint32(len(mesData))
	var bytes [4]byte
	buf := bytes[0:4]
	binary.BigEndian.PutUint32(buf, pkgLen)
	//发送长度
	n, err := conn.Write(buf)
	if n != 4 || err != nil {
		fmt.Println("conn.Write err", err)
		return
	}
	//发送消息本身
	_, err = conn.Write(mesData)
	if err != nil {
		fmt.Println("conn.Write err", err)
		return
	}
	mes, err = readPkg(conn)
	if err != nil {
		fmt.Println("readPkg fail err=", err)
		return
	}
	var loginResMes common.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登录成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return
}

package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	common "go_code/chatroom/common/message"
	"io"
	"net"
)

//读包
func readPkg(conn net.Conn) (mes common.Message, err error) {
	buf := make([]byte, 8096)
	//conn.Read在conn没有关闭的情况下才会阻塞，如果客户端关闭了就不会阻塞
	_, err = conn.Read(buf[:4])
	if err != nil {
		//err = errors.New(fmt.Sprintln("conn.Read err=",err1))
		return
	}
	//根据buf长度，转为uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])
	//根据pkglen读取内容
	n, err := conn.Read(buf[:pkgLen])
	if uint32(n) != pkgLen || err != nil {
		//err = errors.New(fmt.Sprintln("conn.read失败：",err2))
		return
	}
	//把pkg反序列化
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		//err = errors.New(fmt.Sprintln("json.Unmarshal err=",err3))
		return
	}
	return
}

//写包
func writePkg(conn net.Conn, data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var bytes [4]byte
	buf := bytes[0:4]
	binary.BigEndian.PutUint32(buf, pkgLen)
	//发送长度
	n, err4 := conn.Write(buf)
	if n != 4 || err4 != nil {
		fmt.Println("conn.Write err", err4)
		return
	}
	//发送消息
	n1, err5 := conn.Write(data)
	if n1 != int(pkgLen) || err5 != nil {
		fmt.Println("conn.Write err", err5)
		return
	}
	return
}

//登录
func serverProcessLogin(conn net.Conn, mes *common.Message) (err error) {
	var loginMes common.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}
	var resMes common.Message
	resMes.Type = common.LoginResMesType
	var loginResMes common.LoginResMes
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		loginResMes.Code = 200
	} else {
		loginResMes.Code = 500
		loginResMes.Error = "该账号未注册，请先注册..."
	}
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail err=", err)
	}
	resMes.Data = string(data)
	resMesData, err := json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail err=", err)
	}
	//发送
	writePkg(conn, resMesData)
	return
}

//根据客户端消息种类不同调用不同函数处理
func serverProcessMes(conn net.Conn, mes *common.Message) (err error) {
	switch mes.Type {
	case common.LoginMesType:
		serverProcessLogin(conn, mes)
	case common.RegisterMesType:

	default:
		fmt.Println("消息类型不存在，无法处理！")
	}
	return
}

//处理与客户端的通讯
func process(conn net.Conn) {
	//读取客户端消息
	defer conn.Close()
	for {
		//读取数据包
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("服务器端退出...")
				return
			} else {
				fmt.Println("readPkg err=", err)
				return
			}

		}
		//fmt.Println("mes=",mes)
		err = serverProcessMes(conn, &mes)
		if err != nil {
			return
		}
	}
}

func main() {
	//监听端口
	fmt.Println("服务器监听8888端口")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err=", err)
	}
	//等待客户端链接
	for {
		fmt.Println("等待客户端的连接...")
		conn, err1 := listen.Accept()
		if err1 != nil {
			fmt.Println("Listen.Accept err=", err)
		}
		//链接成功，启动一个协程和客户端保持通讯

		go process(conn)
	}

}

package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	common "go_code/chatroom/common/message"
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

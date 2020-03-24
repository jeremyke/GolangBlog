package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/go_item_01/common/message"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {

	buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据")
	//conn.Read 在conn没有关闭的情况下才会阻塞
	//如果关闭就不会阻塞
	n, err := conn.Read(buf[:4])
	if n != 4 || err != nil {
		//fmt.Println("conn.Read err= ", err)

		//err = errors.New("read pkg header error!")//自定义错误
		return
	}
	//fmt.Println("读到的长度为： ", buf[:4])
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])
	//根据pakLen读取消息内容
	n, err = conn.Read(buf[:pkgLen]) //从conn中读[:pkgLen]这么多个字节扔到buf中
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read fail err= ", err)
		return
	}
	//把pkgLen 反序列化成message.Message
	//var mes message.Message
	err = json.Unmarshal(buf[:pkgLen], &mes) //必须加&，否则是空，因为是引用类型，不能是值类型
	if err != nil {
		fmt.Println("json.Unmarshal err :", err)
		return
	}
	return
}

func writePkg(conn net.Conn, data []byte) (err error) {
	//先发送一个长度给对方，可以双向发送
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
	//发送消息本身
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("发送失败", err)
		return
	}
	return
}

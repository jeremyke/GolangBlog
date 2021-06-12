package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("服务器开始监听...")
	//使用网络协议是tcp,监听本地8888端口
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err:", err)
	}
	defer listen.Close() //延迟关闭
	//循环等待客户端链接我
	for {
		fmt.Println("循环等待链接中...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		} else {
			fmt.Printf("conn=%v 客户端ip=%v\n", conn, conn.RemoteAddr().String())
		}
		//协程函数处理客户端链接
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		//创建切片
		buf := make([]byte, 1024)
		//fmt.Printf("服务器正在等待客户端%s 发送信息\n",conn.RemoteAddr().String())
		//等待客户端通过conn发送信息，如果客户端没有write,协程就阻塞在这里
		n, err := conn.Read(buf) //读取conn内容
		if err != nil {
			fmt.Println("客户端退出了，err=", err)
			return
		}
		recieveStr := string(buf[:n])
		fmt.Print(recieveStr)
	}
}

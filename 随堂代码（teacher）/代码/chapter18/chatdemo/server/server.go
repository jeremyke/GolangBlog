package main

import (
	"fmt"
	_ "io"
	"net" //做网络socket开发时,net包含有我们需要所有的方法和函数
)

func main() {

	//一、监听端口
	fmt.Println("服务器开始监听....")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	defer listen.Close() //延时关闭listen

	//二、等待链接
	for {
		//等待客户端链接
		fmt.Println("等待客户端来链接....")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)

		} else {
			fmt.Printf("Accept() suc con=%v 客户端ip=%v\n", conn, conn.RemoteAddr().String())
		}
		//这里准备其一个协程，为客户端服务
		go process(conn)
	}

	//fmt.Printf("listen suc=%v\n", listen)
}

//三、处理请求
func process(conn net.Conn) {

	//这里我们循环的接收客户端发送的数据
	defer conn.Close() //关闭conn

	for {
		//3.1 创建一个新的切片接受客户端数据
		buf := make([]byte, 1024)
		//conn.Read(buf)
		//1. 等待客户端通过conn发送信息
		//2. 如果客户端没有wrtie[发送]，那么协程就阻塞在这里
		//fmt.Printf("服务器在等待客户端%s 发送信息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf) //从conn读取
		if err != nil {

			fmt.Printf("客户端退出 err=%v", err)
			return //!!!
		}
		//3. 显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[:n]))
	}

}

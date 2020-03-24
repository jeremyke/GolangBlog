package main

import (
	"fmt"
	_ "io"
	"net"
)

func process(conn net.Conn) {
	//这里循环的接受客户端发送的数据
	defer conn.Close() //关闭conn，防止崩溃
	for {
		//创建一个新的切片
		buf := make([]byte, 1024)
		//等待客户端通过conn发送信息给我，如果客户端没有发送就一直等待，有可能超时
		//fmt.Printf("服务器在等待客户端发送的信息 %s\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("服务器的Read err =", err)
			return
		}
		//或者这样写：
		//if err == io.EOF{
		//fmt.Println("远程连接客户端退出")
		//return
		//}
		//显示客服端发送的内容发送到终端
		fmt.Print(string(buf[:n])) //n是从管道中真实取到的数据
	}

}

func main() {
	fmt.Println("服务器开始监听。。。")
	//tcp表示使用网络协议是tcp协议
	//0.0.0.0:8888表示再本地监听的是8888端口
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err = ", err)
		return
	}
	defer listen.Close() //延时关闭
	//循环连接等待客户端连接
	for {
		//等待连接
		fmt.Println("等待连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("出错:", err)
		} else {
			fmt.Println("Accept conn : \n 客户端的ip : ", conn, conn.RemoteAddr().String()) //不停等待客户连接
		}
		//起一个协程，为客户服务
		go process(conn)
	}
	fmt.Println("listen is ", listen)
}

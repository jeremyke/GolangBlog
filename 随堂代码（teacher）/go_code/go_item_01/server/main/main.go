package main

import (
	_ "encoding/binary"
	_ "encoding/json"
	_ "errors" //自定义错误的方法
	"fmt"
	_ "go_code/go_item_01/common/message"
	"go_code/go_item_01/server/model"
	_ "io"
	"net"
	"time"
)

/*func writePkg(conn net.Conn, data []byte) (err error){
	//先发送一个长度给对方，可以双向发送
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//现在发送长度
	n, err := conn.Write(buf[:4])
	if n!= 4||err != nil{
		fmt.Println("发送失败", err)
		return
	}
	//发送消息本身
	n, err = conn.Write(data)
	if n!= int(pkgLen) || err != nil{
		fmt.Println("发送失败", err)
		return
	}
	return
}*/

//编写一个函数serverProcessLogin函数，专门处理登陆请求
/*func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	//1.先从mes中取出data,并直接反序列化成LoginMes
	var loginMes message.LoginMes
	json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil{
		fmt.Println("json.Unmarshal fail err =", err)
		return
	}
	//先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//之后声明一个LoginResMes,并赋值
	var LoginResMes message.LoginResMes

	//合法用户
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		//合法

		LoginResMes.Code = 200//200状态码，表示用户合法
	}else {
		//不合法
		LoginResMes.Code = 404//404状态码，表示用户不存在
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
	err = writePkg(conn, data)
	return
}*/

/*//编写一个SersverProcessMes函数
//实现的功能是根据客户端发送消息种类不同，决定调用哪个函数来处理
func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type{
	case message.LoginResMesType:
		err = serverProcessLogin(conn, mes)
		//处理登陆
	case message.RegisterMesType:
		//处理注册
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}*/

/*func readPkg(conn net.Conn) (mes message.Message, err error){

	buf := make([]byte, 8096)
		fmt.Println("读取客户端发送的数据")
		//conn.Read 在conn没有关闭的情况下才会阻塞
		//如果关闭就不会阻塞
		n, err := conn.Read(buf[:4])
		if n!= 4 || err != nil{
			//fmt.Println("conn.Read err= ", err)

			//err = errors.New("read pkg header error!")//自定义错误
			return
		}
		//fmt.Println("读到的长度为： ", buf[:4])
		var pkgLen uint32
		pkgLen = binary.BigEndian.Uint32(buf[0:4])
		//根据pakLen读取消息内容
		n, err = conn.Read(buf[:pkgLen])//从conn中读[:pkgLen]这么多个字节扔到buf中
		if n != int(pkgLen) || err != nil {
			fmt.Println("conn.Read fail err= ",err)
			return
		}
		//把pkgLen 反序列化成message.Message
		//var mes message.Message
		err = json.Unmarshal(buf[:pkgLen], &mes)//必须加&，否则是空，因为是引用类型，不能是值类型
		if err != nil {
			fmt.Println("json.Unmarshal err :", err)
			return
		}
		return
}*/

//处理于客户端的通讯
func process(conn net.Conn) {
	//读取客户端发送的信息
	defer conn.Close()

	//创建一个总控
	processer := &Processer{
		Conn: conn,
	}
	err := processer.process_01()
	if err != nil {
		fmt.Println("客户端与服务器端出现问题 :", err)
		return
	}
	//读取客户端发送的信息
	/*for {

		//这里我们将读取数据包直接封装成一个函数readPkg()，返回Message,Err

		//buf := make([]byte, 8096)
		//fmt.Println("读取客户端发送的数据")
		//n, err := conn.Read(buf[:4])
		//if n!= 4 || err != nil{
			//fmt.Println("conn.Read err= ", err)
			//return
		//}
		//fmt.Println("读到的长度为： ", buf[:4])

		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF{
				fmt.Println("客户端退出")
				return
			} else {
			fmt.Println("conn err :", err)
			return
			}
		}
		fmt.Println("mes :", mes)
		err = serverProcessMes(conn, &mes)
		if err != nil{
			return
		}
	}*/
}

//这里我们写一个函数完成对UserDaod的初始化任务
func initUserDao() {
	//pool是全局变量，就可以初始化成功，这里需要注意一个初始化顺序问题
	//initpool在initUserDao里
	model.MyUaserDao = model.NewUserDao(pool)
}

func main() {
	//当服务器启动时，我们就去初始化reids
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
	//提示信息
	fmt.Println("服务器在8889端口监听。。。")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("失败, err:", err)
		return
	}
	//一旦监听成功，等待客户连接
	for {
		fmt.Println("等待链接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("出错, err:", err)
		}

		//一旦连接成功，启动一个协程于客户端保持通讯
		go process(conn)

	}
}

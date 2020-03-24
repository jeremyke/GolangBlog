package main

import (
	"fmt"
	"go_code/go_item_01/common/message"
	"go_code/go_item_01/server/process"
	"go_code/go_item_01/server/utils"
	"io"
	"net"
)

//先创建一个Processer的结构体
type Processer struct {
	Conn net.Conn
}

//编写一个SersverProcessMes函数
//实现的功能是根据客户端发送消息种类不同，决定调用哪个函数来处理
func (this *Processer) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginResMesType:
		//创建一个UserProcess实例
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
		//处理登陆
	case message.RegisterMesType:
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
		//处理注册
	case message.SmsMesType:
		//创建一个SmsProcess实例完成转发群聊
		smsProcess := &process2.SmsProcess{}
		smsProcess.SendGroupMes(&mes)

	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

func (this *Processer) process_01() (err error) {
	for {
		//创建一个Transfer实例完成读包任务
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()

		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出")
				return err
			} else {
				fmt.Println("conn err :", err)
				return err
			}
		}
		fmt.Println("mes :", mes)
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}

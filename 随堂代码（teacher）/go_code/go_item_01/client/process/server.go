package process

import (
	"encoding/json"
	"fmt"
	_ "go_code/go_item_01/client/utils"
	"go_code/go_item_01/common/message"
	"go_code/go_item_01/server/utils"
	"net"
)

//显示登陆成功后的界面
func ShowMenu() {
	fmt.Println("   恭喜登陆成功   ")
	fmt.Println("1.显示在线用户列表")
	fmt.Println("2.   寻找好友    ")
	fmt.Println("3.   添加好友    ")
	fmt.Println("4.   开始聊天    ")
	fmt.Println("5.   更换皮肤    ")
	fmt.Println("6.   离线消息    ")
	fmt.Println("7.   退出系统    ")
	fmt.Println("请选择数字")
	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("显示在线用户列表")
		//outputOnline()
	case 2:
		fmt.Println("寻找好友")
	case 3:
		fmt.Println("添加好友")
	case 4:
		fmt.Println("开始聊天")
	case 5:
		fmt.Println("更换皮肤")
	case 6:
		fmt.Println("离线消息")
	default:
		fmt.Println("感谢使用")
	}
}

//和服务器端保持通讯
func ServerProcessMes(conn net.Conn) {
	//创建一个transfer实例，不停的读取服务器发送的消息
	tf := &utils.Transfer{
		Conn: conn,
	}

	for {
		fmt.Println("客户端正在读取")
		mes, err := tf.ReadPkg() //_应为mes,是下一步所需要的信息，这里我没有打开是因为目前还未写道下一步
		if err != nil {
			fmt.Println(err)
			return
		}
		//如果读取到就是进行第二极界面
		switch mes.Type {
		case message.NotifyUserStatusMesType: //有人上线了
			//处理
			//取出NotifyUserStatusMes
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			//把这个用户的状态保存到客服map[int]User中
			updataUserStatus(&notifyUserStatusMes)
		default:
			fmt.Println("我处理不了")

		}
	}
}

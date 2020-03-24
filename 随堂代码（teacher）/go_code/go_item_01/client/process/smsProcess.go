package process

import (
	"encoding/json"
	"fmt"
	"go_code/go_item_01/common/message"
	"go_code/go_item_01/server/utils"
)

type SmsProcess struct {
}

//发送群聊天的消息

func (this *SmsProcess) SendGroupMes(content string) (err error) {
	//创建一个message
	var mes message.Message
	mes.Type = message.SmsMesType
	//创建一个SmsMes实例
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus
	//序列化
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("不玩了，你错了")
		return
	}
	mes.data = string(data)
	//对mes再次序列化
	data, err = json.Marshal(smsMes)
	if err != nil {
		fmt.Println("不玩了，你错了")
		return
	}
	//将mes发给服务器
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	//发送
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	return
}

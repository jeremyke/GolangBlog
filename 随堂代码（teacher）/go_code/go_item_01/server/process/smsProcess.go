package process2 //函数跟包同名会有问题
import (
	"encoding/json"
	"fmt"
	"go_code/go_item_01/common/message"
	"go_code/go_item_01/server/utils"
	"net"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMes(mes *message.Message) {
	//遍历服务器端的map,将消息转发出去
	//取出mes的内容
	var mes message.SmsMes
	err := json.Marshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("反序列化失败")
		return
	}
	data, err := json.Marshal(mes)
	for id, up := range userMgr.onlineUsers {
		//这里要过滤掉自己
		if id == smsMes.UserId {
			continue
		}
		this.SendMesToEachOlinrUser(data, up.Conn)
	}
}

func (this *SmsProcess) SendMesToEachOlinrUser(data []byte, conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败 err =", err)
	}
}

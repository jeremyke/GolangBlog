package process

import (
	"fmt"
	"go_code/go_item_01/client/model"
	"go_code/go_item_01/common/message"
)

//客户端要维护的map
var olineUsers map[int]message.User = make(map[int]message.User, 10)
var CueUser model.CurUser //用户登录成功后完成对CurUser的初始化
//在客户端显示当前用户
func outputOnlineUser() {
	//遍历
	fmt.Println("当前在线用户列表")
	for id, user := range olineUsers {
		fmt.Println("用户id: ", id)

	}
}

//编写一个处理返回的NotifyUserStatusMes
func updataUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	//适当优化
	user, ok := olineUsers[notifyUserStatusMes.UserId]
	if !ok { //原来没有
		user := &message.User{
			UserId: notifyUserStatusMes.UserId,
			//UserStatus : notifyUserStatusMes.UserStatus,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	olineUsers[notifyUserStatusMes.UserId] = user
	outputOnlineUser()
}

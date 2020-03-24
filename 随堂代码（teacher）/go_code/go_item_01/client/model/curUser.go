package model

import (
	"go_code/go_item_01/common/message"
	"net"
)

//因为在客户端很多地方使用到curUser，所以我们将他做成全局的
type CurUser struct {
	Conn net.Conn
	message.User
}

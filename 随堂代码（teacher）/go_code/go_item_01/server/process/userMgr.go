package process2

import (
	"fmt"
)

//因为UserMgr实例在服务器端有且只有一个
//因为在很多地方都会用到，因此我们将其定义为全局变量
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

//完成对UserMgr初始化工作
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

//完成对onlineUser
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

//删除一个功能
func (this *UserMgr) DelOnlineUser(UserId int) {
	delete(this.onlineUsers, UserId)
}

//返回所有在线的用户
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return this.onlineUsers
}

//根据ID返回一个对应的值
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	//如何从map中取出一个值，带检测
	up, ok := this.onlineUsers[userId]
	if !ok { //当前用户不在线
		err = fmt.Errorf("用户%d不在", userId)
		return
	}
	return
}

package message

//确定有哪些消息
const (
	loginMesType            = "lginMes"
	LoginResMesType         = "loginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

//这里我们定义几个用户状态的常量
const (
	UserOL = iota //在线
	UserOF        //离线
	UserBS        //忙碌
)

type Message struct {
	Type string //消息类型
	Data string //消息的类型
}

//定义两个消息，之后需要再增加
type LoginMes struct {
	UserId   int
	UserPwd  string
	UserName string
}

type LoginResMes struct {
	Code   int    //404表示用户没有注册，200表示登陆成功
	UserId []int  //增加字段，保存用户的ID
	Error  string //返回错误信息
}

type RegisterMes struct {
	User User //类型就是User结构体
}

type RegisterResMes struct {
	Code  int    //406表示用户已经占用，200表示注册成功
	Error string //返回错误信息
}

//为了配合服务器端推送通知（用户状态变化）
type NotifyUserStatusMes struct {
	UserId int `json : "userId"`
	Status int `json : "status"` //用户状态
}

//增加一个SmsMes这是发送的
type SmsMes struct {
	Content string `json : "contest"`
	User           //匿名结构体
}

//

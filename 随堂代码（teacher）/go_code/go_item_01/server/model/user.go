package model

//定义一个用户结构体
type User struct {
	UserName string `json : "userName"` //为了·序列化跟反序列化成功，要给个tag
	UserId   int    `json : "userId"`
	UserPwd  string `json : "userPwd"`
}

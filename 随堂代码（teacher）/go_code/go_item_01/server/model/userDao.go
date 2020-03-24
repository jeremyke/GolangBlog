package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"go_code/go_item_01/common/message"
)

//我们在服务器启动后，就初始化一个userDao实例，
//把他当成全局变量，在需要何使用时可直接使用
var (
	MyUserDao *UserDao
)

//定义一个UserDon结构体
//完成对User结构体的操作
type UserDao struct {
	pool *redis.Pool
}

//使用工厂模式，常见一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

//UserDao应该提供根据一个用户Id返回一个User实例
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	//通过给定的id去redis查询这个用户
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil { //表示没有找到对应的Id
			err = ERROR_USER_NOT_NOTEXISTS
		}
		return
	}
	user = &User{}
	//这里我们需要吧res反序列化为一个user实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return
	}
	return
}

//完成登陆校验
//Login 完成对用户的验证
//如果用户的ID跟密码都正确，则返回一个user实例
//如果用户的id 个别pwd都有错误，则返回错误信息
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	//先从UserDao 中链接池内去吃一根连接
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	//这时证明用户是获取到了，
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

func (this *UserDao) Register(user *message.User) (err error) {
	//先从UserDao 中链接池内去吃一根连接
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)
	if err != nil {
		err = ERROR_USER_NOT_EXISTS
		return
	}
	//这时证明用户是获取到了，
	//if user.UserPwd != userPwd {
	//err = ERROR_USER_PWD
	//return
	//}
	//return

	//这时，要序列化id 可以完成注册

	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	//入库
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存信息错误")
		return
	}
	return
}

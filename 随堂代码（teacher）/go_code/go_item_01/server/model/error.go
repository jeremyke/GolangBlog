package model

import (
	"errors"
)

//根据业务逻辑需要，自定义一些错误
var (
	ERROR_USER_NOT_NOTEXISTS = errors.New("用户不存在")
	ERROR_USER_NOT_EXISTS    = errors.New("用户已经存在")
	ERROR_USER_PWD           = errors.New("用户密码不正确")
)

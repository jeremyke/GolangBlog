package main

import (
	_ "fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

//定义一个全局变量
var pool *redis.Pool

func initPool(address string, maxIdle, maxActive int, IdleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle:     maxIdle,     //最大空闲链接数
		MaxActive:   maxActive,   //表示和数据库的最大连接数，0表示没有限制
		IdleTimeout: IdleTimeout, //最大空闲时间
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
}

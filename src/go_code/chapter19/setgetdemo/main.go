package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func connect() redis.Conn {
	//连接redis
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("redis connect err:", err)
	}
	return c
}

func main() {
	c := connect()
	defer c.Close()
	//set
	_, err2 := c.Do("Set", "key2", "七七团aa")
	if err2 != nil {
		fmt.Println("redis set err:", err2)
		return
	}
	//get
	r, err3 := redis.String(c.Do("Get", "key2"))
	if err3 != nil {
		fmt.Println("redis get err:", err3)
		return
	}
	fmt.Println(r)
}

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
	//hmset
	_, err2 := c.Do("HMSet", "user02", "name", "蒋东莲", "age", 18)
	if err2 != nil {
		fmt.Println("redis set err:", err2)
		return
	}
	//HMget
	info, err3 := redis.Strings(c.Do("HMGet", "user02", "name", "age"))
	if err3 != nil {
		fmt.Println("redis get err:", err3)
		return
	}

	for k, v := range info {
		fmt.Printf("info[%d]=%s\n", k, v)
	}

}

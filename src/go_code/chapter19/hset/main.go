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
	//hset
	_, err2 := c.Do("HSet", "user01", "name", "tom猫猫")
	_, err2 = c.Do("HSet", "user01", "age", 18)
	if err2 != nil {
		fmt.Println("redis set err:", err2)
		return
	}
	//get
	name, err3 := redis.String(c.Do("HGet", "user01", "name"))
	age, err4 := redis.Int(c.Do("HGet", "user01", "age"))
	if err3 != nil || err4 != nil {
		fmt.Println("redis get err:", err3)
		return
	}
	fmt.Printf("user01的name是：%v，age是：%d", name, age)
}

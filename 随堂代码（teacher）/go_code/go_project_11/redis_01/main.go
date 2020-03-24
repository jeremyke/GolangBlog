package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis" //引入redis包
)

func main() {
	//通过go，向redis 写入数据和读取数据
	//连接redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("错误，没有连接成功 错误是：", err)
		return
	}
	defer conn.Close()

	fmt.Println("conn succ...", conn)
	//写入数据
	_, err = conn.Do("HSet", "user_01", "name", "崔默涵")
	if err != nil {
		fmt.Println("错误，没有连接成功 错误是：", err)
		return
	}

	_, err = conn.Do("HSet", "user_01", "age", "21")
	if err != nil {
		fmt.Println("错误，没有连接成功 错误是：", err)
		return
	}

	fmt.Println("conn ok...", conn)
	//取出数据
	r_1, err := redis.String(conn.Do("HGet", "user_01", "name")) //内置方法进行转换类型
	if err != nil {
		fmt.Println("错误，没有连接成功 错误是：", err)
		return
	}

	r_2, err := redis.Int(conn.Do("HGet", "user_01", "name")) //内置方法进行转换类型
	if err != nil {
		fmt.Println("错误，没有连接成功 错误是：", err)
		return
	}
	//因为返回的r是空接口，
	//因为name对应的值是string, 因此我们需要转换
	//不能这样用类型断言：nameString := r.(string)
	fmt.Printf("提取的值为：%v\n %v", r_1, r_2)

	//一次性放多个
	_, err = conn.Do("HMSet", "user_02", "name", "百里凌涵", "age", "21")
	if err != nil {
		fmt.Println("错误，没有连接成功 错误是：", err)
		return
	}

	r_3, err := redis.Strings(conn.Do("HMGet", "user_02", "name", "age")) //内置方法进行转换类型
	if err != nil {
		fmt.Println("错误，没有连接成功 错误是：", err)
		return
	}
	for i, v := range r_3 {
		fmt.Printf("r[%d]=%s\n", i, v)
	}

	//Redis 链接池
	//pool = &redis.Pool{
	//Maxldle: 8,//最大空闲链接数
	//MaxActive: 0,//表示和数据库的最大连接数，0表示没有限制
	//IdleTimeout :100, //最大空闲时间
	//Dial : func()(redis.Conn, err){
	//return redis.Dial("tcp", "localhost : 6379")
	//}
	//}
	//先从pool取出一个连接
	conn = pool.Get()
	defer conn.Close() //别忘记关闭，否则不能再放回进行池里面
	_, err = conn.Do("Set", "name", "嘤嘤嘤")
	if err != nil {
		fmt.Println("错误，没有成功 错误是：", err)
		return
	}
	r_4, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("错误，没有成功 错误是：", err)
		return
	}
	fmt.Println("r_4", r_4)

}

//定义一个全局的pool
var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   //最大空闲链接数
		MaxActive:   0,   //表示和数据库的最大连接数，0表示没有限制
		IdleTimeout: 100, //最大空闲时间
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost : 6379")
		},
	}
}

package redis

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
	_, err = conn.Do("set", "name", "cmh")
	if err != nil {
		fmt.Println("错误，没有连接成功 错误是：", err)
		return
	}
	fmt.Println("conn ok...", conn)
	//取出数据
	r, err := redis.String(conn.Do("Get", "name")) //内置方法进行转换类型
	if err != nil {
		fmt.Println("错误，没有连接成功 错误是：", err)
		return
	}
	//因为返回的r是空接口，
	//因为name对应的值是string, 因此我们需要转换
	//不能这样用类型断言：nameString := r.(string)
	fmt.Println("提取的值为：", r)
}

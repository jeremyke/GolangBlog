package main

import (
	"fmt"
)

func main() {
	intChan := make(chan int, 10)
	for i := 0; i <= 9; i++ {
		intChan <- i
	}
	stringChan := make(chan string, 5)
	for j := 0; j <= 4; j++ {
		stringChan <- "cmh" + fmt.Sprintf("%d", j)
	}

	//可能不知道什么时候关闭，所以用select方法
	for {
		select { //注意：这里如果intChan没有关闭，也不会一直阻塞
		//而会自动的到下一个case匹配
		case v := <-intChan:
			fmt.Printf("从intChan 读取的数据 %d\n", v)
		case v := <-stringChan:
			fmt.Printf("从stringChan 读取的数据 %d\n", v)
		default:
			fmt.Println("娶不到") //这里程序员可以添加自己的逻辑
		}
	}
}

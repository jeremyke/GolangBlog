package main

import (
	"fmt"
)

type Monster struct {
	Name string
	Age  int
}

func main() {
	monster1 := Monster{"marry", 18}
	fmt.Println(monster1)

	monster2 := new(Monster)
	//monster2 是个指针
	(*monster2).Name = "smith"
	//上面的写法很繁琐，所以可以简化为：monster2.Name = "smith".底层会自动处理加上*
	monster2.Name = "jerry"
	(*monster2).Age = 20
	fmt.Println(*monster2)

	var monster3 *Monster = &Monster{"alitter", 19}
	//因为Monster3 是一个指针，因此标准访问字段的方法是：*monster3.Name = "tom",同样可以简化为monster3.Name = "tom"，原因同上
	(*monster3).Name = "tom"
	monster3.Age = 17
	fmt.Println(*monster3)
}

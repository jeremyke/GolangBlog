package main

import (
	"fmt"
)

func main() {
	monsters := make([]map[string]string, 1) //make切片,2指长度
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "孙悟空"
		monsters[0]["sex"] = "男"
	}
	newMonster := map[string]string{
		"name": "猪八戒",
		"sex":  "男",
	}

	monsters = append(monsters, newMonster) //动态增加map的个数
	fmt.Println(monsters)
}

package main

import (
	"fmt"
)

func main() {
	users := make(map[string]map[string]string, 10)
	modifyUser(users, "tom")
	modifyUser(users, "jerry")
	fmt.Println(users)
}

func modifyUser(users map[string]map[string]string, name string) {
	if users[name] != nil {
		users[name]["pws"] = "888888"
	} else {
		users[name] = make(map[string]string)
		users[name]["pwd"] = "888888"
		users[name]["nickname"] = "昵称为：" + name
	}
}

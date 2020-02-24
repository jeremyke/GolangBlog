package main

import (
	"fmt"
)

func main() {
	var studentMap = make(map[string]map[string]string)
	studentMap["stu01"] = make(map[string]string, 2)
	studentMap["stu01"]["name"] = "tom"
	studentMap["stu01"]["sex"] = "男"

	studentMap["stu02"] = make(map[string]string, 2)
	studentMap["stu02"]["name"] = "jerry"
	studentMap["stu02"]["sex"] = "女"

	fmt.Println(studentMap)
}

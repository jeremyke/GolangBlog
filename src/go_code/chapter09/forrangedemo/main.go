package main

import (
	"fmt"
)

func main() {
	cityMaps := make(map[string]string)
	cityMaps["no1"] = "北京"
	cityMaps["no2"] = "上海"
	cityMaps["no3"] = "深圳"
	cityMaps["no3"] = "广州"

	for _, value := range cityMaps {
		fmt.Println(value)
	}

	var studentMap = make(map[string]map[string]string)
	studentMap["stu01"] = make(map[string]string, 2)
	studentMap["stu01"]["name"] = "tom"
	studentMap["stu01"]["sex"] = "男"

	studentMap["stu02"] = make(map[string]string, 2)
	studentMap["stu02"]["name"] = "jerry"
	studentMap["stu02"]["sex"] = "女"

	for k1, value := range studentMap {
		fmt.Println("k1=", k1)
		for k2, value2 := range value {
			fmt.Printf("\t k2=%v  v2=%v\n", k2, value2)
		}
		fmt.Println()
	}
}

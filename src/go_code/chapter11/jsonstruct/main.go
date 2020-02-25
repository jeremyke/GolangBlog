package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string  `json:"name"` //tag
	Score float64 `json:"score"`
}

func main() {
	var Stu1 = Student{"小明", 68.0}
	json1, _ := json.Marshal(Stu1)
	fmt.Println("json", string(json1))
}

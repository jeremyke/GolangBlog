package unmaeshal
import (
	"fmt"
	"encoding/json"
)

type Monster struct{
	Name string 
	Age int
//	Birthday string
//	Sal float64
//	Skill string

func unmarshalStruct() {
	str := "{\"Age\":500,\"Name\":\"崔默涵\"}"

	var monster Monster
	data, err :=json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println("序列化错误")
	}
	fmt.Println(string(data))
}

func unmaeshal(){
	unmarshalStruct()
}


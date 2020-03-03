package testingdemo2

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (m *Monster) Store() bool {
	str, err := json.Marshal(m)
	if err != nil {
		fmt.Println("序列化失败：", err)
		return false
	}
	fileName := "/home/jeremyke/Desktop/b.txt"
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("打开文件失败：", err)
		return false
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString(string(str))
	writer.Flush()
	return true
}

func (m *Monster) Restore() bool {
	fileName := "/home/jeremyke/Desktop/a.txt"
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("读取失败：", err)
		return false
	}
	err = json.Unmarshal(content, m)
	if err != nil {
		fmt.Println("json反序列化失败：", err)
		return false
	}
	return true
}

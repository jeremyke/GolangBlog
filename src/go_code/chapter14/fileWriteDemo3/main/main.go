package main

import (
	"fmt"
	"io/ioutil"
)

/*func main()  {
	readFileName := "/home/jeremyke/Desktop/b.txt"
	writeFileName := "/home/jeremyke/Desktop/a.txt"
	read,err1 := ioutil.ReadFile(readFileName)
	writeFile,err2 := os.OpenFile(writeFileName,os.O_WRONLY | os.O_TRUNC,0666)
	if err1 !=nil {
		fmt.Println("打开文件失败：",err1)
	}
	if err2 != nil {
		fmt.Println("打开失败：",err2)
	}
	defer writeFile.Close()
	writer := bufio.NewWriter(writeFile)
	fmt.Printf("%v",string(read))
	writer.WriteString(string(read))
	writer.Flush()
}*/

func main() {
	readFileName := "/home/jeremyke/Desktop/b.txt"
	writeFileName := "/home/jeremyke/Desktop/a.txt"
	data, err := ioutil.ReadFile(readFileName)
	if err != nil {
		fmt.Println("打开文件失败：", err)
	}
	if err2 := ioutil.WriteFile(writeFileName, data, 0777); err2 != nil {
		fmt.Println("写入文件失败：", err2)
	}
}

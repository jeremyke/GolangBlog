## 14.文件操作
#### 14.1 文件的基本介绍

文件就是数据源，在程序中是以流的形式来操作的。golang程序<--->文件（中间就是流）

**流:**数据在数据源和程序之间经历的路径

- 输入流：数据从数据源（文件）到程序（内存）的路径【读文件】
- 输出流：数据在程序（内存）到数据源（文件）的路径【写文件】

os.File封装了所有文件的相关操作，File是一个结构体。这里经常会用到os.File

#### 14.2 文件的基本操作

1) 打开一个文件进行读操作：os.Open(name string)(*File,error)
2) 关闭一个文件：File.Close()

```go
package main

import (
	"fmt"
	"os"
)

func main()  {
	//打开一个文件
	file,err := os.Open("/home/jeremyke/Desktop/a1.txt")
	if err != nil {
		fmt.Println("打开失败：",err)
	}
	//输出文件
	fmt.Println(&file)//xo11121 说明：文件是一个指针
	//关闭文件
	file.Close()
}

```

#### 14.2 读取文件的应用实例

1) 读取文件的内容，并显示在终端（带缓冲区的方式）

```go
func main()  {
	//打开一个文件
	file,err := os.Open("/home/jeremyke/Desktop/a.txt")
	if err != nil {
		fmt.Println("打开失败：",err)
	}
	//关闭文件
	defer file.Close()

	//创建*Reader,是带缓冲的 默认大小4K
	reader := bufio.NewReader(file)
	for {
		str,err := reader.ReadString('\n')//读到一个换行就结束
		if err != io.EOF {//io.EOF文件的末尾
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取结束")

}
```

2) 读取文件的内容并且显示在终端（使用ioutil一次将整个文件读入内存中），这种方式适用于文件不大的情况。ioutil.ReadFile

```go
func main()  {
	fileName := "/home/jeremyke/Desktop/a.txt"
	content,err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("读取失败：",err)
	}
	fmt.Printf("%v",string(content))//这个方法不需要显示的打开和关闭文件
}
```

#### 14.3 写文件的应用实例

func OpenFile(name string,flag int,perm FileMode)(file *File,err error)

**说明：** os.OpenFile是一个更一般性的文件打开函数，他会使用指定的选项（如：O_RDONLY等），指定的模式（如0666等）打开指定名称的文件。
        操作成功，返回文件对象可用于I/O.如果出错，错误底层类型是*PathError
        
```go
func main()  {
	fileName := "/home/jeremyke/Desktop/b.txt"
	file,err := os.OpenFile(fileName,os.O_WRONLY | os.O_CREATE,0666)
	if err !=nil {
		fmt.Println("打开文件失败：",err)
	}
	defer file.Close()

	//写入缓冲流
	str := "hello klp\n"
	writer := bufio.NewWriter(file)
	for i := 5; i>0; i-- {
		writer.WriteString(str)
	}
	//刷入文件
	writer.Flush()
}
```
- 从一个文件读取内容写入到另外一个文件

```go
func main()  {
	readFileName := "/home/jeremyke/Desktop/b.txt"
	writeFileName := "/home/jeremyke/Desktop/a.txt"
	data,err := ioutil.ReadFile(readFileName)
	if err !=nil {
		fmt.Println("打开文件失败：",err)
	}
	if err2 := ioutil.WriteFile(writeFileName,data,0777);err2 != nil {
		fmt.Println("写入文件失败：",err2)
	}
}
```
- 判断文件是否存在
Golang中判断文件或者文件目录是否存在的方法是os.Stat()函数返回的错误值进行判断：
1) 返回nil,说明文件或者文件目录存在
2) 返回的错误类型为os.IsNotExist()判断为true，说明文件或者文件目录不存在
3) 返回的错误类型为其他类型，则不确定是否存在

```go
func PathExists(path string)(bool,error){
    _,err := os.Stat(path)
    if err ==nil {
        return true,nil
    }
    if os.IsNotExist(err) {
        return false,nil
    }
    return false,err
}
```
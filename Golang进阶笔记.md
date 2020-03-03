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

#### 14.3 读取文件的应用实例

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

#### 14.4 写文件的应用实例

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

#### 14.5 命令行参数

os.Args是一个string切片，用来存储所有的命令行参数。第一个是程序名称，后面的都是参数。

###### 14.5.1 应用案例

```go
func main()  {
	fmt.Println("命令行参数是：",len(os.Args))
	for k,v := range os.Args {
		fmt.Printf("args[%v]=%v\n",k,v)
	}
}
```

//输出：
命令行参数是： 4
args[0]=./main
args[1]=asasd
args[2]=fgfg
args[3]=aaa

###### 14.5.2 flag包用来解析命令行参数

这种方法参数顺序可以随意,例如：mysql -u root -p 123 -h localhost 

```go
func main()  {
	var user,pwd,host string
	var port int
	//这种方法参数顺序可以随意
	flag.StringVar(&user,"u","","用户名")
	flag.StringVar(&pwd,"pwd","","密码")
	flag.StringVar(&host,"h","localhost","主机名")
	flag.IntVar(&port,"port",3306,"端口号")

	flag.Parse()//转换，必须调用该方法。

	fmt.Printf("user=%v pwd=%v host=%v port=%v",user,pwd,host,port)
}
```

#### 14.6 json基本介绍

json是一种轻量级的数据交换格式，易于人阅读和编写，同时也易于机器解析和生成。

json有效的提升网络传输效率，通常程序在网络传输时会先将数据（结构体，map等）序列化成json字符串，到接收方得到json字符串时，再
反序列化恢复成原来的数据类型（结构体，map等），这种方式已然成为各个语言的标准。

Golang--（序列化）-->json字符串--（网络传输）-->程序--（反序列化）-->其他语言

###### 14.6.1 json序列化

- 结构体序列化

````go
type Monster struct {
	Name string `json:"name"`
	Age int	`json:"age"`
	Birthday string
	Sal float64
	skill string
}

func testStruct()  {
	monster := Monster{
		Name:     "蒋冬莲",
		Age:      24,
		Birthday: "0703",
		Sal:      8000.0,
		skill:    "测试",
	}
	data,err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化失败，%v",err)
	}else{
		fmt.Printf("序列化后结果为：%v\n",string(data))
	}
}
````

- map序列化

```go
func testMap()  {
	var a map[string]interface{}//表示k为string,v为任意类型
	a = make(map[string]interface{})
	a["name"] = "呵呵"
	a["age"] = 12
	a["address"] = "中国"
	data,err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化失败，%v",err)
	}else{
		fmt.Printf("序列化后结果为：%v\n",string(data))
	}

}
```
- slice序列化

```go
func testSlice()  {
	var stringSlice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "蒋冬莲"
	m1["sex"] = "女"
	m1["address"] = "珠海"
	stringSlice = append(stringSlice,m1)
	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "柯丽萍"
	m2["sex"] = "女"
	m2["address"] = "深圳"
	stringSlice = append(stringSlice,m2)
	data,err := json.Marshal(stringSlice)
	if err != nil {
		fmt.Printf("序列化失败，%v",err)
	}else{
		fmt.Printf("序列化后结果为：%v\n",string(data))
	}
}
```

###### 14.6.2 json反序列化

- 反序列化成结构体

```go
func testStruct()  {
	str := "{\"name\":\"蒋冬莲\",\"age\":24,\"Birthday\":\"0703\",\"Sal\":8000}"
	var monster Monster
	err := json.Unmarshal([]byte(str),&monster)
	if err != nil {
		fmt.Printf("反序列化失败，%v",err)
	}else{
		fmt.Printf("反序列化后结果为：%v\n",monster)
	}
}
```

- 反序列化成map

```go
func testMap()  {
	str := "{\"address\":\"中国\",\"age\":12,\"name\":\"呵呵\"}"
	var a map[string]interface{}//反序列化操作map不需要make操作,在反序列化是会自动make
	err := json.Unmarshal([]byte(str),&a)
	if err != nil {
		fmt.Printf("map反序列化失败，%v",err)
	}else{
		fmt.Printf("map序列化后结果为：%v\n",a)
	}

}
```

- 反序列化成slice

```go
func testSlice()  {
	str := "[{\"address\":\"珠海\",\"name\":\"蒋冬莲\",\"sex\":\"女\"},{\"address\":\"深圳\",\"name\":\"柯丽萍\",\"sex\":\"女\"}]"

	var stringSlice []map[string]interface{}

	err := json.Unmarshal([]byte(str),&stringSlice)
	if err != nil {
		fmt.Printf("slice反序列化失败，%v",err)
	}else{
		fmt.Printf("slice反序列化后结果为：%v\n",stringSlice)
	}
}
```



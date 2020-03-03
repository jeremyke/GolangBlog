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
>这里Unmarshal要将json数据解码写入一个指针。

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
## 15.单元测试

在工作中，我们常常需要确认一个函数或者一个模块的结果是否正确，往往需要做单元测试。

#### 15.1 基本介绍

Golang自带一个轻量级的测试框架testing和自带的go test命令来实现单元测试和性能测试，可以基于这个框架写针对相应函数的测试用例，
也可以基于这个框架写相应的压力测试用例。通过单元测试，解决如下问题：
1) 确保每个函数是可运行的，且运行结果正确
2) 确保写出来的代码性能良好
3) 单元测试能及时发现程序设计或者实现的逻辑错误，使得问题及早的暴露，便于问题定位解决，而性能测试的重点在于发现程序设计上的一些
    问题，使得程序能够在高并发的情况下保持稳定。

#### 15.2 快速入门

使用golang的单元测试对addUpper和sub函数进行测试

cal.go

```go
package main

import (

)

func AddUpper(n int) int {
	res := 0
	for i := 0; i<=n; i++ {
		res += i
	}
	return res
}

func getSub(n,m int) int {
	return n-m
}

```

cal_test.go

```go
package main

import (
	_"fmt"
	"testing"
)

func TestAddUpper(t *testing.T)  {
	res := AddUpper(10)
	if res != 55 {
		//fmt.Printf("AddUpper(10)执行错误，期望值=%v 实际值=%v\n",55,res)
		t.Fatalf("AddUpper(10)执行错误，期望值=%v 实际值=%v\n",55,res)
	}
	//正确就输出日志
	t.Logf("AddUpper(10)执行正确")
}

```

sub_test.go

```go
package main

import (
	"testing"
)

func TestGetSub(t *testing.T)  {
	res := getSub(10,2)
	if res != 8 {
		t.Fatalf("getSub(10,2)执行错误，期望值=%v 实际值=%v\n",8,res)
	}
	//正确就输出日志
	t.Logf("getSub(10,2)执行正确")
}

```

**单元测试的原理：**
1) testing框架先将xxx_test.go的文件引入
2) 在main函数里面调用TestXXX()函数


#### 15.3 单元测试细节说明

1) 测试用例文件名必须以_test.go结尾，比如cal_test.go,cal不是固定的。
2) 测试用例函数名必须以Test开头，一般是Test+测试函数名
3) TestAddUpper(t *Testing.T)的形参类型必须是*Testing.T
4) 一个测试用例文件中，可以有多个测试用例函数。
5) 运行测试用例指令：
    （1）cmd>go test(如果运行正确，无日志，错误时，输出日志)
    （2）cmd>go test -v (无论运行正确或者错误，都输出日志)
6) 当出现错误时，可以使用t.Fatalf来格式化输出错误信息，并退出程序
7) t.Logf方法输出相应的日志
8) 测试用例函数不需要main函数
9) Pass表示测试咏柳运行成功，fail表示失败，都成功才为成功。
10) 测试单个文件，一定要带上被测试的原文件：go test -v ./cal_test.go
11) 测试单个方法：go test -v -test.run TestAddUpper

## 16.goroutine(协程)和channel(管道)

#### 16.1 goroutine基本介绍

**并发和并行：**

1) 多线程程序在单核CPU上运行，微观时间点上只有一个线程在执行，就是并发。
2) 多线程程序在多核CPU上运行，微观时间点上有CPU核数个线程在执行，就是并行。

**Go协程和Go主线程：**

1) Go主线程：一个Go线程上，可以起多个协程，你可以这样理解，协程是轻量级的线程，编译器做的优化。

2) Go协程的特点：
    （1）有独立的栈空间
    （2）共享程序堆空间
    （3）调度由用户控制（协程的开启和关闭是程序员控制的）
    （4）协程是轻量级的线程
    
#### 16.2 goroutine快速入门

在主线程中开启一个goroutine,该协程每隔1秒输出“hello world”,在主线程中也每隔一秒输出“hello golang”,输出10次后，退出程序。

```go
package main

import (
	"fmt"
	"strconv"
	"time"
)

func test()  {
	for i:=0; i<10; i++ {
		fmt.Printf("hello world " + strconv.Itoa(i) + "\n")
		time.Sleep(time.Second)
	}
}

func main()  {
	go test()//开启一个协程执行test函数
	for i:=0; i<10; i++ {
		fmt.Printf("hello golang " + strconv.Itoa(i) + "\n")
		time.Sleep(time.Second)
	}
}

```

输出结果：

```text
hello golang 0
hello world 0
hello world 1
hello golang 1
hello golang 2
hello world 2
hello golang 3
hello world 3
hello golang 4
hello world 4
hello golang 5
hello world 5
hello world 6
hello golang 6
hello golang 7
hello world 7
hello world 8
hello golang 8
hello world 9
hello golang 9

```

**说明:**

1) 如果主线程退出了，则协程即使没有执行完也会退出

2) 当然协程也可以在主线程退出前，自己完成了任务就提前退出了。

**小结：**

1) 主线程是一个物理线程，直接作用在CPU上的。是重量级的，非常耗费cpu资源。
2) 协程是从主线程开启的，是轻量级的线程，是逻辑态。对资源消耗相对小。
3) Golang的协程机制是重要的特点，可以轻松开启上万个协程。其他编程语言的并发机制是一般基于线程的，开启过多的线程，
    资源耗费大，这里就突出了Golang的优势。
    

#### 16.3 goroutine的调度模型

###### 16.3.1 MPG基本介绍

M:操作系统主线程（物理线程）；P:协程执行需要的上下文；G:协程

###### 16.3.2 MPG运行的状态1：

1) 当前程序有3个M,如果3个M都在一个CPU上执行，就是并发，如果分别在不同CPU上执行就是并行。
2) M1，M2,M3正在执行一个G,M1的协程队列有3个，M2的协程队列有3个，M3的协程队列有2个
3) Go的协程是轻量级的线程，是逻辑态的，而C/java的多线程，往往是内核态的，比较重量级，几千个就可能耗光cpu

###### 16.3.3 MPG运行的状态2：

M0在执行一个G0协程，另外有3个协程在队列中等待，如果G0协程阻塞了（IO）,这时就会创建一个M1的主线程，并且将M0等待的
3个协程挂到M1上执行，M0继续等待G0的IO。等到M0不阻塞了，M0主线程就会被就会被放置到空闲的主线程继续执行。

#### 16.3 设置Golang运行的CPU个数

```go
func main()  {
	cpuMum := runtime.NumCPU()
	fmt.Println(cpuMum)

	//设置使用cpu个数
	runtime.GOMAXPROCS(cpuMum-1)
	fmt.Println("ok")
}
```



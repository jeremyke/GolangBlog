## 1.快速入门

#### 1.1 编译和运行
>编译成功会生成一个二进制码文件
```go
go bulid [-o 指定编译生成的可执行文件名] 源文件
go run //编译和运行合为一步
```
#### 1.2 go开发注意事项
```text
1. Go源文件必须以.go结尾
2. Go应用程序的执行文件入口是main函数
3. Go严格区分大小写
4. Go行末无需加‘;’,编译器会主动把特定符号后的换行符转换为分号，比如：
        （1）标识符、整数、浮点数、虚数、字符或字符串文字
        （2）关键字  break、continue、fallthrough或return
        （3）运算符  ++,--
        （4）分隔符  )、]或}
5. Golang定义的变量或者import的包在程序中没有用到就会编译失败
6. 括号成对出现 
```
#### 1.3 Golang转义字符
```text
1. \t 制表符
2. \n 换行符号
3. \\ 转义斜杠
4.\"  转义双引号
5. \r 回车（从当前行的最前面输出，替换原来的内容）
```
#### 1.4 golang注释
 - 行注释 //
 - 块注释 /**/
 
#### 1.5 Go规范的代码风格
 - Go官方推荐行注释来注释整个方法或者语句
 - tab缩进 （gofmt -w go源文件 格式化文件）
 - 运算符两边习惯加上空格
 - 换行：语句太长需要主动换行但又要避免编译器主动加分号，可以用“,”，连接上下两行
 
## 2.变量

#### 2.1 变量的介绍
```text
使用：
(1)变量申明
(2)赋值
(3)使用
```
#### 2.2 变量入门
```go
package main

import "fmt"

//变量的三步曲（申明，赋值，使用）
func main()  {
	var i int
	i = 20
	fmt.Println("i=",i)
}
```
#### 2.3 使用注意
- 指定变量类型，如果申明后不赋值，会使用默认值（int 默认 0）
- 根据值自行判断变量类型（类型推导）
- 省略var,注意 := 左侧的变量不应该是已经申明过的，否则导致编译失败
>i := 6 等价于 var i = 6

- 多变量申明：在编程中,golang可以支持一次性申明多个变量
```go
//多变量申明
package main

import "fmt"

func main()  {
	//第一种
	var a, b, c int
	//第二种
    var j, k, l = 100, "tom", 2.6
	//第三种
	m, n := "i",600
    fmt.Println("a=",a,"b=",b,"c=",c)
    fmt.Println("j=",j,"k=",k,"l=",l)
    fmt.Println("m=",m,"n=",n)
}
```
- 一次性生成多个全局变量（在Golang中，函数外的变量就是全局变量）
```go
package main

import "fmt"

var  (
	name = "tom"
	sex = "male"
	age = 25
	)
func main()  {
    fmt.Println("name=",name,"sex=",sex,"age=",age)
}
```
- 变量在所在区域的数据值可以在同一类型范围内不断变化
- 变量在同一作用域内不能重名
- Golang的变量如果没有赋初值，编译器会使用默认值，比如int默认为0，string默认为空字符串,小数也默认为0

#### 2.4 Golang中“+”的使用
- 左右两边为数值型时，则做加法运算
- 左右2边字符串，则做字符串拼接

```go
package main

import "fmt"

func main()  {
	var a = 10
	var b = 20
	var c = a + b
	var stra = "i love"
	var strb = " you"
	var strc = stra + strb
	fmt.Println("计算结果为",c)
	fmt.Println("计算结果为",strc)
}

```

## 3. 数据类型

![image](https://github.com/jeremyke/goStudy/blob/master/pic/18100630282833.png)


由于中间主要是通过看书来学习，到了接口这一张，感觉看不懂了，所以通过视频来学习吗，顺便接着记下笔记。但是中间的章节也会在后续的时间里慢慢补上。

## n. 接口（interface）

#### n.1 基本介绍

接口类型定义了一组方法，但是这些不需要实现。而且接口中不能包含任何变量。在自定义类型中需要实现的接口的时候，才写具体的方法。

```go
type 接口名称 interface{
    method1(参数列表)返回值列表
    method2(参数列表)返回值列表
}
```
实现接口所有方法：

```go
func (m 自定义类型) method1(参数列表)返回值列表{//method1(参数列表)返回值列表 这整个部分叫做函数的签名，需与接口完全一致
    //方法实现
}
func (m 自定义类型) method2(参数列表)返回值列表{
    //方法实现
}
```

- 总结：
```text
（1）接口里面的所有方法都没有方法体，即接口的方法是没有实现的，接口体现了程序设计的多态和高内聚低耦合的思想。
（2）Golang的接口不需要显示实现。只要一个变量，含有类型的所有方法，那么这个变量就实现了这个接口。因此接口中没有implement关键字。
```

#### n.2 注意事项和细节说明

 - （1）接口路本身不能创建实例，但是可以指向一个实现了该接口的自定义类型的变量（实例）。也可以这样描述：一个自定义类型，只有实现了某个
 接口，才能将该自定义类型的实例赋值给该接口类型
 ```go
package main

import "fmt"

type AInterface interface{
	Say()
}

type Stu struct{
	
}

func (stu Stu) Say(){
    fmt.Println("学生说")
}

func main()  {
    var s Stu//结构体变量，实现了AInterface接口
    var a AInterface = s//a接口指向一个实现了该接口的自定义类型的变量s（实例）
    a.say()    
}
```
- （2）接口中所以的方法都没有方法体，即都是没有实现的方法。

- （3）Golang中，一个自定义类型需要将某个接口的所有方法都实现，我们才能说这个自定义类型实现了该接口。

- （4）只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型。
```go
package main

import "fmt"

type AInterface interface{
	Say()
}

// int类型的自定义类型
type Integer int{
}

func (i Integer) Say(){
    fmt.Println("this number is = ",i)
}

func main()  {
    var i Integer
    var a AInterface = i
    a.say()    
}
```
- （5）一个自定义类型可以实现多个接口
```go
package main

import "fmt"

type AInterface interface{
	Say()
}
type BInterface interface{
    Hello()
} 

type Stu struct{
	
}

func (stu Stu) Say(){
    fmt.Println("学生say")
}
func (stu Stu) Hello(){
    fmt.Println("学生hello")
}

func main()  {//一个方法实现了2个接口
    var s Stu
    var a AInterface = s
    var b BInterface = s
    a.say()    
    b.hello()    
}
```
- （6）Golang接口中不能有任何变量或常量
- （7）一个接口（A）可以继承多个其他接口(B,C)，这时如果要实现A接口，也必须将B,C的接口全部实现。
```go
package main

import (
    "fmt"
)

type BInterface interface{
	test1()
}
type CInterface interface{
	test2()
}
//a接口继承人b,c两个接口
type AInterface interface{
    BInterface
    CInterface
	test3()
}

type Stu struct{}

func (stu Stu) test1()  {
 
}
func (stu Stu) test2()  {
 
}
func (stu Stu) test3()  {
 
}
func main()  {
    var stu Stu
    var a AInterface = stu
    a.test1()
}
```
- (8)interface 类型默认是一个指针（引用类型），如果没有对interface初始化就是有，就会输出nil

- (9)空接口interface{}（interface{}也是一种类型） 没有任何方法，所以所有类型都实现了空接口，即我们可以把任何一个变量赋值给空接口。
```go
package main

import (
    "fmt"
)

type T interface{
}
func main()  {
    var t T
    var t2 interface{} 
    var num float64 = 8.8
    t = num
    t2 = num
fmt.Println(t)
} 
//221
```


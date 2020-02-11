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

- 基本数据类型
```text
数值型，字符型，字符串型，布尔型
```

- 复杂数据类型
```text
指针，数组，结构体（struct，相当于php的class），管道（channel），函数，切片（slice。动态数组），接口，map（相当于集合）
```
#### 3.1 整数类型

- 注意说明
```text
int 占用空间以当前的操作系统为准：32位系统占用32位，64位系统占用64位，unit同理。
rune 同int32一样，表示一个unicode码，储存带有中文的字符串特别有用。
byte 等价于uint8,适合存储字符。
```

- 整数的使用细节
```text
1.Go的整型默认为int型（可用fmt.Printf("类型位 %T",str)查看数据类型）
2.查看程序中变量占用的字节大小：fmt.Printf("占用的字节数位 %d",unsafe.Sizeof(str))
3.Golang程序中，整形变量在使用时，遵守保小不保大的原则。即：在保证程序正确运行时，尽量使用占用空间小的数据类型。
4.bit计算机最小存储单位，byte计算机基本存储单位。1byte=8bit
```

#### 3.2 浮点数类型

###### 3.2.1 基本介绍
> 小数类型就是用来存放小数的，比如1.879

###### 3.2.1 浮点型的分类

|类型|占用空间|表数范围|
|-:|-|-|
|单精度float32|占用4个字节|-3.403E38~3.403E38|
|双精度float64|占用8个字节|-1.798E308~1.798E308|

**说明一下:**
1) 关于浮点数在机器中存放的简单说明。浮点数=符号位+指数位+尾数位，说明浮点数都是有符号的。
2) 尾数部分可能丢失，造成精度损失。

**浮点数使用细节：**
1) Golang浮点类型有固定的范围和字段长度，不受OS的影响。
2) Golang浮点型默认声明位float64
3) 浮点型常量有两种表示形式：
    十进制，例如：5.12， .512
    科学计算数：5.12e2(相当于5.12*10的2次方)
4) 开发中没有特殊指定，推荐使用float64。
 
#### 3.3 字符类型（char）

###### 3.3.1 基本介绍
> Golang中没有专门的字符类型，如果要存储单个字符（字母），一般使用byte来保存.
>字符串就是一串固定长度的字符连接起来的字符序列。GO的字符串是由单个字节连接起来的。也就是说对于传统的字符串是由字符组成的，而
>go的字符串不同，它由字节组成的。

```go
func main(){
    var c1 byte = 'a'
    var c2 byte = '0'
    //当我们直接输出byte值，就是输出了对应的ascll码值
    fmt.Println("c1=",c1)//97
    fmt.Println("c2=",c2)//48
    //如果希望输出对应的字符，需要格式化输出
    fmt.Printf("c1=%c c2=%c",c1,c2)
}
```
**说明一下:**
1) ASCll表的字符，直接用byte保存，因为只占8位。
2) 如果保存的字符对应的码值大于255，可以考虑用int
3) 需要按照字符的方式输出，需要格式化输出。

**字符使用细节：**

1)字符常量用单引号（‘’）括起来的单个字符，例如：var c1 byte = 'a'

2)Go中允许使用转义字符'\n'来将后面的字符转变位特字符型常量。

3)Golang字符使用utf-8编码（英文占1个字节，汉字占3个字节）

4)在Golang中，字符的本质是一个整数，直接输出时，对应该字符对应的UTF-8编码的码值。

5)可以直接给某个变量赋值一个数字，然后按照格式化输出时%c,会输出该数字对应的Unicode字符。

6)字符类型可以进行运算，相当于一个整数，因为他有对应的Unicode码。

###### 3.3.2 字符类型的本质讨论

1) 字符型存储到计算机中，需要将字符对应的码值（整数）找出来，
    存储：字符-->码值-->二进制-->存储
    读取：二进制-->码值-->字符-->读取
2) 字符和码值是通过uft-8编码表一一对应的。
3) Go语言编码统一采用utf-8


#### 3.4 布尔类型（bool）

###### 3.4.1 基本介绍

1) bool类型数据只允许取值true和false
2) bool类型占用1个字节
3) bool类型使用与逻辑运算，一般用于程序流程控制

#### 3.5 字符串类型（string）

###### 3.5.1 基本介绍
>字符串就是一串固定长度的字符连接起来的字符序列。Go的字符串是由单个字节连接起来的。

###### 3.5.2 使用细节

1) Golang的字符串统一使用了UTF-8编码标识Unicode文本，避免造成乱码。
2) 字符串一旦赋值，就不能修改。
3) 字符串的两种表示形式
    (1)双引号会识别转义字符
    (2)反引号，以字符串的原生形式输出，包括换行和特殊字符，可以实现防止攻击，输出源代码等效果。
4) 字符串拼接，使用“+”
5) 当一行字符串太长时，需要使用到多行字符串，可以如下处理
```go
str := "hello" + 
       " world!"

fmt.Println(str)
```
#### 3.6 基本数据类型的默认值

|数据类型|默认值（零值）|
|-:|:-:|
|整形|0|
|浮点型|0|
|字符串型|""|
|布尔型|false|

>fmt.Printf("a=%v",a)//%v指按变量的值输出

#### 3.7 基本数据类型的转换

###### 3.7.1 基本介绍

Golang中在不同类型的变量之间赋值时需要显示转换，也就是说：golang的数据类型不能自动转换。

###### 3.7.2 基本语法

表达式: T(V) 将值V转换为类型T
T：就是数据类型，比如int32,float64
V：需要转换的变量

```go
func main(){
  var i int32 = 34
  var j float32 = float32(i)
  fmt.Prinf("i=%v j=%v",i,j)
}
```
**细节说明**

1) Go中，数据类型的转换可以从表示小的范围-->表示大的范围，也可以从范围大的-->范围小的
2) 被转换的是变量存储的数据（即值），变量本身的数据类型并没有变化。
3) 在转换中，比如将int64转为int8，编译时不会报错，只是转换的结果按溢出处理，和我们期望的结果不一样，所以需要谨慎转换。

#### 3.8 基本数据类型和string的转换

###### 3.8.1 基本数据类型-->string类型

方式一：fmt.Sprintf("%参数"，表达式)（比较常用）
>1.参数需要和表达式的数据类型相匹配（可以参考手册插叙）
>2.fmt.Sprintf()会返回转换后的字符串

方式二：使用strconv包的函数

func FormatBool(b bool) string

func FormatFloat(f[被转化的变量] float64,fmt[格式] byte,prec[小数位数],bitSize[小数是float64] int) string

func FormatInt(i int64,base【进制】 int)string

func FormatUint(uint64,base int) string

func Itoa(a int) string


###### 3.8.2 string类型-->基本数据类型

使用strconv包的函数:

func ParseBool(str string)(value bool,err error)

func ParseFloat(s string,bitSize int)(f float64,err error)

func ParseInt(s string,base int,bitSize int)(i int64,err error)

func ParseUint(s string,b int,bitSize int)(n uint64,err error)

**注意：**

string转基本数据类型，要确保string类型能够转成有效数据，否则会转为零值。

```go
func main(){

    var str1 string = "hello"
    var str2 string = "123"
    int1, _ = strconv.ParseInt(str1,10,64)
    int2, _ = strconv.ParseInt(str2,10,64)
    fmt.Printf("int1's type %T int1=%v\n",int1,int1)// string 0:hello无法转为整数，只能设置位0值
    fmt.Printf("int2's type %T int2=%v\n",int1,int1)
}
```
#### 3.9 指针

###### 3.9.1 基本介绍

1) 基本数据类型，变量存的就是值，也叫值类型。
2) 获取变量的地址用&，比如：var n int,获取n的地址：&n
3) 指针类型，指针变量存的是一个地址，这个地址指向的空间存的才是值，比如：var ptr *int = &n
4) 获取指针类型所指向的值。使用：*,比如：var ptr *int 使用*ptr获取ptr指向的值

###### 3.9.2 指针细节说明

1) 值类型,都有对应的指针类型，形式为*数据类型。比如int数据类型，其指针类型位*int
2) 值类型包括：基本数据类型（int,float,bool,string）和数组，结构体

#### 3.10 值类型和引用类型

###### 3.10.1 说明
值类型： 基本数据类型（int系列，float系列，bool,string），数组，结构体

引用类型：指针，slice切片，map,channel，interface

###### 3.10.2 使用特点

1) 值类型，变量直接存储值，内存通常在栈中分配

2) 引用类型，变量存储的是一个地址，这个地址对应的空间才是真正的存储数据，内存通常在堆中分配。当没有任何变量引用这个地址时，该地址对应的数据
空间就成了一个垃圾，由GC来回收。


#### 3.11 标识符的命名规范

###### 3.11.1 概念

1) Golang对各种变量，方法等命名时使用的字符序列称为标识符。
2) 凡是自己起名字的地方都叫标识符

###### 3.11.2 命名规范

1) 由英文字母，阿拉伯数字，_组成
2) 数字不能开头
3) 严格区分大小写
4) 标识符不能包含空格
5)下划线“_”本身在Golang中是一个特殊的标识符，称为空标识符。可以代表任何其他的标识符。但是他对应的值会被忽略。所以仅能被作为占位符使用，
   不能作为表示服使用。
6)不能以系统保留关键字作为标识符号，比如break,if...

###### 3.11.3 标识符命名注意事项

1) 包名：保存package的名字和目录一致，尽量采用有意义的包名，不要和标准库冲突。

2) 变量名，函数名，常量名：采用驼峰法。

3) 如果变量名，函数名，常量名的首写字母大写，则可以被其他的包访问。如果首字母小写，则只能在本包中使用。

## 4. 运算符号

|运算符分类|举例|
|-|-|
|算术运算符|+，-，*,/,%,++,--|
|赋值运算符|=，+=，-=，*=,/=,%=;<<=,>>=,&=,^=,|=|
|比较（关系）运算符|==，！=，<,>,<=,>=|
|逻辑运算符|&&(逻辑与)，||(逻辑或)，！(逻辑非)|
|位运算符|-|
|其它运算符|-|

#### 4.1 算术运算符
>+，-，*,/,%,++,--

###### 4.1.1 基本使用
1)除法（/）：如果运算的数都是整数，结果会丢失小数部分，保留整数部分。
```go
fmt.Println(10/3) //3
```
2)求余（%）：结果的正负与被求余保持一致。
```go
fmt.Println(-10%3) //-1
```
###### 4.1.2 细节说明
1)对于“/”，整数之间的除法，结果为整数。
2)取模的结果正负与被取模数保持一致。
3)Golang的++，--只能当做一个独立的语句使用，不能赋值给其它变量，或者和其它变量比较等等操作。
4)++，--只能写在变量的后面，不能写在前面。

**课堂练习**
>1.假如还有97天放假，问：xx个星期零几天
>2.定义一个变量保存华氏温度，华氏温度-->摄氏温度：5/9**(华氏温度-100)，请求出华氏温度对应的摄氏温度
>代码见chapter04/suanshuexcise

#### 4.2 比较运算符
>==，！=，<,>,<=,>=

###### 4.2.1 基本介绍

1) 运算结果为bool类型，true/false
2) 经常用于if结构的条件中或循环结构的条件中

###### 4.2.2 细节说明

1)运算结果为bool类型，true/false
2)关系运算符组成的表达式我们称为关系表达式
3)比较运算符“==”不要误写位“=”

#### 4.3 逻辑运算符
>&&(逻辑与)，||(逻辑或)，！(逻辑非)

#### 4.4 赋值运算符 
>=，+=，-=，*=,/=,%=;<<=,>>=,&=,^=,|=

###### 4.4.1 赋值运算符特点
1) 运算顺序从右到左
2) 赋值运算的右边可以是变量，表达式【任何有值的都可以叫表达式】，常量值

#### 4.5 运算符的优先级

|分类|描述|关联性|
|-|-|-|
|后缀|(),[],->,++,--|从左到右|
|单目|+，-，！，～，（type）,*,&,sizeof|从右到左|
|乘法|*,/,%|从左到右|
|加法|+，-|从左到右|
|位移|<<,>>|从左到右|
|关系|<,<=,>,>=|从左到右|
|相等（关系）|==，！=|从左到右|
|按位AND|&|从左到右|
|按位XOR|^|从左到右|
|按位OR|"|"|从左到右|
|逻辑AND|&&|从左到右|
|逻辑OR|"||"|从左到右|
|赋值运算|”=，+=，-=，×=，/=,%=,>>=,<<=,&=,^=,|=“|从右到左|
|逗号|，|从左到右|

**说明**

1)上表优先级从上到下，依次降低。

2)只有单目运算符，赋值运算符从右到左运算，其他均为从左到右边。

#### 4.6 位运算符
>&,|,^,<<,>>

**特别注意**
在计算机中所有的运算，都是用的二进制的补码来运算的。

1) 右移 >> : 低位溢出，符号位不变，并用符号位补全溢出的高位。
2) 左移 << : 符号位不变，低位补0。



#### 4.7 其他运算符
>&（取地址）,*（取指针变量的值）


#### 4.8 键盘输入语句

###### 4.8.1 基本介绍

**步骤：**
1) 导入fmt包
2) 调用fmt包的fmt.Scanln()[备注：获取一行的输入信息]或者fmt.Scanf()[备注：一次性输入，空格作为间隔],推荐前一种

```go
package main

import (
	"fmt"
)
//从控制台接收用户信息【姓名，年龄，薪水，是否通过考试】
func main()  {
	var name string
	var age byte
	var sal float64
	var isPass bool
    //mt.Scanln()[备注：获取一行的输入信息]
	fmt.Println("请输入姓名")
	fmt.Scanln(&name)//输入的值直接传给了name的地址，引用传递，会改变原来的name值
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水")
	fmt.Scanln(&sal)
	fmt.Println("请输入是否 通过考试")
	fmt.Scanln(&isPass)
    //fmt.Scanf()[备注：一次性输入，空格作为间隔]
	fmt.Println("请依次输入姓名，年龄，薪水，是否通过考试")
	fmt.Scanf("%v %d %f %t",&name,&age,&sal,&isPass)

}
```









































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
（2）Golang的接口不需要显示实现。只要一个变量，含有接口的所有方法，那么这个变量就实现了这个接口。因此接口中没有implement关键字。
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

#### n.3 接口的实践
> 参考chapter11/sliceorder

#### n.3 接口实现和继承的比较
- 代码案例
```go
package main

import (
	"fmt"
)

type Monkey struct {
	Name string
}

func (this *Monkey)climb()  {
	fmt.Println(this.Name,"天生会爬树...")
}

type Bird interface {
	Fly()
}

type SonMonkey struct {
	Monkey
}

func (this *SonMonkey) Fly()  {
	fmt.Println(this.Name, " 通过实现，会飞翔...")
}


func main()  {
	xiaoMonkey := SonMonkey{Monkey{Name:"孙悟空"}}
	xiaoMonkey.climb()
	xiaoMonkey.Fly()
}
```
- 小结
```text
（1）当A结构体继承了B结构体，就自动继承了B结构体的字段和方法，并且可以直接使用。
（2）当A结构体需要扩展功能，同时不希望去破坏继承关系，则可以去实现某个接口即可。因此可以这么认为：实现接口是对继承接口的补充。
```

- 实现接口 vs 继承
```text
（1）实现接口和继承解决的问题不一样：
    继承的价值在于：解决代码的复用性和可维护性。
    实现接口的价值在于：设计好各种规范，让自定义的类型去实现具体方法。
（2）接口比继承更加灵活：
    继承满足is - a 关系，而接口只需满足 like -a 关系。
（3）接口在一定程度上实现代码解耦:
```

-------------------p225---------------------




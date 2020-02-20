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
        （2）关键字  break,continue,allthrough或return
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
int 占用空间以当前的操作系统为准：32位系统占用32位（4个字节），64位系统占用64位（8个字节），unit同理。
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
1)对于“/”，整数之间的除法，结果为整数。除数和被除数的数据类型要一致。
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

## 5 程序流程控制
> (1)顺序控制
> (2)分支控制
> (3)循环控制

#### 5.1 顺序控制
>从上到下依次执行

#### 5.2 if分支控制
>单分支，双分支，多分支

###### 5.2.1 单分支

```go
if 条件表达式 {
    执行代码块
}

func main()  {
    var age int
    fmt.Printf("请输入你的年龄")
    fmt.Scanln(&age)
    if age > 18 {
        fmt.Printf("你的年龄大于18，需要对自己行为负责！")
    }
}
```
**细节**

1)Golang支持在if条件语句中申明变量，且必须有条件表达式
```go
if var age = 8;age>18{
}
```
2)if后面的小括号可带可不带，推荐不写！
###### 5.2.2 多分支
```go
if 条件表达式 {
    执行代码块1
}else{
    执行代码块2
}
```
**细节：**
else不能换行

###### 5.2.3 多分支
```go
if 条件表达式1 {
    执行代码块1
}else if 条件表达式2 {//可以多个else if
    执行代码块2
}else {
    执行代码块n
}
```
**说明：**

1) else并不是必要的。
2) 多分支只有一个入口，从上到下依次执行。

#### 5.3 switch分支结构

###### 5.3.1 基本介绍

1) switch语句用于基于不同   条件执行不同动作，每一个case分支都是唯一的，从上到下逐一测试，直到匹配为止。

2) 匹配项后面不需要再加break

###### 5.3.2 基本语法

```go
switch 表达式 {
    case 表达式1,表达式2,...://多表达式，只要其中一个表达式匹配即可执行，也就是说是或的关系
        代码块1,
    
    case 表达式3,表达式4,...:
    	代码块2,
    
    //这里可以有多个case语句
    default:
        代码块n+1
}
```

**注意说明**

1) golang的case后面可以多个表达式

2) golang的case后面不需要加break

###### 5.3.3 switch使用细节

1) case后面是一个或多个表达式（即:常量，变量，有返回值的函数...）
2) case后面的各个表达式的值的数据类型必须和switch的表达式数据类型一致
3) case后面多个表达式用逗号隔开
4) case后面的表达式如果是常量值（字面值），则要求不能重复
5) case后面不加break
6) default语句不是必须的
7) switch后也可以不带表达式，类似if-else分支来使用。
```go
switch {
    case age==10 ://这里表达式也可以是范围
        fmt.Println("10.....")
    case age==20 :
    	fmt.Println("20.....")
    defaulf:
        fmt.Println("meiyoulaaaaaa")
}
```
8) switch后面也可以定义/申明一个变量分好结束，但是这里不推荐。
```go
switch age := 10;{
    case 10 :
        fmt.Println("10.....")
    case 20 :
    	fmt.Println("20.....")
    defaulf:
        fmt.Println("meiyoulaaaaaa")
}
```
9) switch穿透fallthrough,如果在case语句块后增加fallthrough，则会继续执行下一个case,这叫switch穿透(只能穿透一层)
```go
var age = 20
switch age {
    case 10 :
        fmt.Println("10.....")
        fallthrough
    case 20 :
    	fmt.Println("20.....")
    defaulf:
        fmt.Println("meiyoulaaaaaa")
}
```
10) Type Switch:switch语句还可以被用于type-switch来判断某个interface变量中实际指向的变量类型。
```go
var x interface{}
var y = 10.0
x = y
switch i := x.(type){//x的实际变量类型是float64
    case nil:
        fmt.Printf("x的类型：%d",i)
    case int:
    	fmt.Println("x是int类型")
    default:
        fmt.Println("都不是。。。")
}
```

**switch和if的比较：**
1) 如果判断的具体数值不多，而且符合整数，浮点数，字符，字符串这几种类型。建议使用switch语句，简洁高效。
2) 其他情况，对区间判断和结果为bool类型的判断，使用if,if的使用范围更广。

#### 5.4 for循环控制

###### 5.4.1 基本语法
```go
for 循环变量初始化;循环条件;循环变量迭代 {
    //循环体
}
```
**四大要素：**
变量初始化，循环条件，循环变量迭代，循环体

**执行顺序：**
循环变量初始化-->循环条件判断-->为真：循环体--->为假：退出循环-->循环变量迭代

###### 5.4.2 细节说明

1) for循环的第二种使用方式：将变量初始化和变量迭代写在其他地方
```go
for 循环判断条件 {
    //循环体
}
```
2) for循环的第三种使用方式：这个写法等价于for ; ; {},通常需要配合break语句使用。
```go
for {
    //循环体
}
```
3) Golang提供for-range的方式，可以方便遍历字符串和数组。
```go
//传统字符串遍历不能含有中文，会出现乱码。因为这是按字节遍历的，汉字占3个字节，以后再说解决方法：转为slice。
str := "hello,world"
//strslice := []rune(str)
for i := 0; i<len(str);i++ {
    fmt.Printf("%c",str[i])
}
//for-range,可以遍历中文,按字符方式遍历。
for index; val := range str {
    fmt.Printf("index=%d,val=%c",index,val)
}
```

###### 5.4.3 while和do...while的实现

Golang中没有while和do...while语法

- 用for实现while：
```go
//循环变量初始化
for{
    if 循环条件表达式{
        break
    }
    //循环体
    //循环变量迭代
}
```

- 用for实现do...while：
```go
//循环变量初始化
for{
    //循环体
    //循环变量迭代
    if 循环条件表达式{
        break
    }
}
```

#### 5.5 多重循环控制

###### 5.5.1 打印金字塔
```go
package main

import (
	"fmt"
)
func main()  {
	var num int
	fmt.Println("请输入金字塔层数")
	fmt.Scanln(&num)
	for i := 1 ; i<=num ;i++ {//层数
		for k := 1 ; k<=num-i ; k++ {//空格数
			fmt.Print(" ")
		}
		for j := 1 ; j<=(2*i-1) ; j++ {//星星数
			fmt.Print("*")
		}
		fmt.Println("\n")
	}
}
```
###### 5.5.2 打印九九乘法表
```go
package main

import (
	"fmt"
)

func main()  {
	for i := 1; i<10 ; i++ {
		for j := 1; j<=i ; j++ {
			fmt.Printf("%v * %v = %v\t",j,i,i*j)
		}
		fmt.Println("\n")
	}
}
```

#### 5.6 跳转控制-break

###### 5.6.1 基本介绍

break 语句用于终止某个语句块的执行，用于中断当前for循环或者switch语句。

###### 5.6.2 细节说明

1) break 语句出现在多层嵌套语句的语法块中时，可以通过标签指明要终止的是哪一层语法块。

2) 标签的基本使用

```go
label1: {
    label2: {
        break label1;//不写就是指当前循环
    }
}
//例子：
label1:
for i:=0;i<10;i++{
    label2:
    for j:=0;j<4;j++{
        break label1; 
   }
}
```
#### 5.7 跳转控制-continue

###### 5.7.1 基本介绍
1) continue语句用于结束本次循环，继续执行下一次循环。
2) continue语句出现在多层嵌套循环的循环语句体中时，可以通过标签指明要跳过的是哪一层循环，和break一样。

#### 5.8 跳转控制-goto

###### 5.8.1 基本介绍
1) Golang中goto可以无条件的转到程序中指定的行
2) goto语句通常与条件语句配合使用。可以用来实现条件转移，跳出循环体等功能。
3) Golang一般不主张使用goto语句，以免造成程序流程的混乱，使理解和调试程序都产生困难。

```go
goto label
...//中间不执行
label:statement
```
###### 5.8.2 案例

```go
func main() {
    fmt.Println("ok1")
    goto label1
    fmt.Println("ok2")
    fmt.Println("ok3")
    fmt.Println("ok4")
    label1:
    fmt.Println("ok5")
}
//ok1 ok5
```

#### 5.9 跳转控制-return

###### 5.9.1 基本结束

return使用在方法或者函数中，表示跳出所在的方法或函数，如果是main函数，表示终止main函数，也就是终止程序。


## 6 函数,包和错误处理

#### 6.1 函数
>为了完成某一项功能的程序指令的集合，统称为函数（自定义函数和系统函数）。

#### 6.1.1 基本语法
```go
func 函数名称(形参列表) (返回值列表) {
    执行语句
    return 返回值列表
}
```
#### 6.2 包
>包的本质就是创建不同的文件夹，来存放程序文件。

###### 6.2.1 包的基本概念

go的每一个文件都是属于一个包的，也就是说go是以包的形式来管理文件和项目目录结构的。

###### 6.2.2 包的三大作用

1) 区分相同名字的函数，变量等标识符
2) 当程序文件很多时，可以很好的管理项目
3) 控制程序，变量等访问范围，即作用域。

###### 6.2.3 包的相关说明

- 打包基本语法

package [包名]

- 引包的基本语法

import “包的路径”(从src后面的目录开始写)

###### 6.2.4 包的使用细节

1) 在给一个文件打包时，该包对应一个文件夹。比如utils目录对应的包名就是utils,文件的包名，通常和文件所在的目录保存一致，一般小写字母。
2) 当一个文件要使用其他包的函数或者变量时，需要引入对应的包。
3) 在import包时，路径从$GOPATH的src下开始，不用带src,编译器自动从src下开始引入。
4) 为了让其他包的文件可以访问到本包的函数，则该函数名的首字母需要大写,即可导出。
5) 在访问其他包函数时，其语法为：报名.函数名。
6) 如果包名比较长，Go支持给包取别名，注意，取别名之后，在该包里面，引入的原来的包名就不能使用了。实例如下：

```go
import (
    test "go_code/asdf/testadminqqq"
)
```

7) 在同一个包里面不能有相同名称的函数或者变量。
8) 如果要编译成可执行文件，就需要将该包声明位main,即package main.这是一个语法规范，如果你是写一个库，包名可以自定义。
```text
说明：
（1）编译的指令，在项目目录下，编译路径不需要带src,编译器会自动带
（2）编译时需要编译main包所在的文件目录。
（3）编译后生成一个默认名的可执行文件，在$GOPATH目录下，可以指定名字和目录，比如：
    放在bin目录下：D:\goproject>go build -o bin\my.exe go_code/project/main
```

#### 6.3 函数的调用机制

###### 6.3.1 return语句

GO函数支持返回多个值，这一点是其它编程语言没有的。
```go
func 函数名 (形参列表) (返回值类型列表) {
    语句...
    return 返回值列表//return str1,str2多值用,隔开
}
```
1) 如果返回多个值，在接收时，希望忽略某个值，则使用_符号站位忽略

2) 如果返回只有一个值，(返回值类型列表)中可以不写小括号()

#### 6.4 函数-递归调用

###### 6.4.1 基本介绍

一个函数在函数体内又调用了本身，我们称之为递归调用。

案例：
```go
func test(n int) {
    if n > 2 {
        n--
        test(n)
    }
    fmt.Println("n=",n)
}
//2 2 3
```

###### 6.4.2 递归调用的总结

1) 执行一个函数时，就创建一个新的受保护的独立空间（新的函数栈）。 

2) 函数的局部变量是独立的，不会相互影响。

3) 递归必须向退出递归的条件逼近，否则就是无限递归。

4) 当一个函数执行完毕，或者遇到return，就会返回，遵守谁调用就将结果返回给谁，同时当函数执行完毕或者返回时，该函数本身也会被系统销毁。

###### 6.4.3 课堂练习

- 题1：斐波那契数 请使用递归的方式，求出菲波拉契数1,1,2,3,5,8,13...给你一个整数求出他的值是多少？(代码见chapter06/feibolaqishu)

- 题2：求函数值，已知f(1)=3;f(n)=2*f(n-1)+1,使用递归的思想求出f(n)的值(代码见chapter06/hanshuzhi)

- 题3：猴子吃桃问题 有一堆桃子，猴子第一天吃了其中的一半，并再多吃了一个！以后每天猴子都吃了其中的一半，并再多吃一个。当到第十天，
  想在吃时（还没吃），发现只有一个桃子了，问一共有多少个桃子？
  
```go
package main

import (
	"fmt"
)

func main()  {
	res := getNum(1)
	fmt.Println("和为",res)
}

func getNum(n int) int{
	var num int
	if n==10 {
		num = 1
	}else{
		num = 2*(getNum(n+1)+1)
	}
	return num
}
```

###### 6.4.3 函数使用的注意事项和细节讨论

1) 函数中的变量是局部的，函数外不生效。
2) 基本数据类型和数组默认都是值传递，即进行值拷贝，在函数内修改，不会影响原来的值。
3) 如果希望函数内的变量能修改函数外的变量，可以传入函数的地址&，函数内以指针的方式操作变量。
4) Go函数不支持重载。
```go
func test (n1 int) int{
    //函数体
}
func test (n1 int,n2 int) int{
    //函数体
}
//这是不允许的！！！！！！

```s
5) 在GOlang中，函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量了。通过该变量可以对函数调用。
```go
func getSum(n1 int, n2 int) int {
    return n1+n2
}

func main()  {
    a := getSum
    fmt.Printf("a的类型%T,getSum类型是%T",a,getSum)//func(int,int) int func(int,int) int
}
```
6) 函数既然是一种数据类型，因此在Go中，函数可以作为形参，并且调用。
```go
func getSum(n1 int, n2 int) int {
    return n1+n2
}
func myFun(funvar func(int,int) int,num1 int,num2 int){
    return funvar(num1,num2)
}
```
7) 为了简化数据类型定义，Go支持自定义数据类型。
    基本语法：type 自定义数据类型 数据类型
```go
type myFunType func(int,int) int
```
8) 支持对函数返回值命名
```go
func cal(n1 int, n2 int) (sum int, sub int){
    sum = n1+n2
    sub = n1-n2
    return
}
```
9) Golang支持可变参数
```go
//支持0到多个参数
func sum(args... int) sum int{    
}
//支持1到多个参数
func sum(n1 int,args... int) sum int {
}
```
**说明：**
（1） args是slice,通过args[index]可以访问到各个值
（2） 案例：
```go
func getSum(n1 int,args... int) int {
    sum := n1
    for i:=0;i<len(args);i++{
        sum += args[i]
    }
    return sum
}
```
（3）如果一个函数的形参列表中有可变参数，则可变参数需要放在形参列表最后。

#### 6.5 init函数

###### 6.5.1 基本介绍
每一个源文件都可以包含一个init函数，该函数会在main函数执行前，被Go运行框架调用，也就是说init会在main函数前被调用。

###### 6.5.1 细节讨论

1) 如果一个文件同时包含全局变量定义，init函数和main函数，则执行的流程是：变量定义-->init函数-->main函数。
如果引入外包，则外包的初始化在引入的时候就被init了，即外包（引入的包）-->本包。
```go
package main
import "fmt"
var age = test()

func test()int {
    fmt.Println("test()")
    return 90
}

func init()  {
    fmt.Println("init()")
}

func main()  {
    fmt.Println("main(),age=",age)
}
//test()  init()    main(),age=90
```
2) init函数的最主要作用，就是完成一些初始化的工作。

#### 6.6 匿名函数

###### 6.6.1 基本介绍

Go支持匿名函数，如果某个函数只是希望使用一次，可以考虑使用匿名函数，匿名函数也可以实现多次调用。

- 使用方式：

（1）在定义匿名函数时直接调用，不用赋值给变量，它只能调用一次。
```go
package main

import "fmt"

func main(){
    res := func (n1 int,n2 int) int {
        return n1+n2
    }(12,45)
    fmt.Println("n1和n2的和为",res)
}
```
（2）将匿名函数赋给一个变量（函数变量），再通过该变量来调用匿名函数
```go
package main

import "fmt"

func main(){
    var a = func (n1 int,n2 int) int {
        return n1+n2
    }//a的数据类型就是函数类型，可以通过a完成调用
    res := a(10,30)
    fmt.Println("和为",res)
}
```

###### 6.6.2 全局匿名函数
如果将匿名函数赋值给一个全局变量，那么这个匿名函数就是一个全局变量，可以在程序有效。
```go
package main

import "fmt"

var (
    //Fun1就是全局匿名函数
    Fun1 = func(n1,n2 int) int {
        return n1*n2
    }
) 

func main(){
    res := Fun1(10,30)
    fmt.Println("积为",res)
}
```
#### 6.7 闭包

###### 6.7.1 基本介绍
闭包就是一个函数和与其相关的引用环境组合的一个整体。

```go
package main

import "fmt"
//累加器 返回的是一个函数。
func AddUpper() func(int) int {
	var n int = 10
	return func (x int) int {
		n = n + x
		return n
	}
}

func main(){
	f := AddUpper()
	fmt.Println(f(1))//11
	fmt.Println(f(2))//13
	fmt.Println(f(3))//16
}
```
**说明：**

（1）返回的是一个匿名函数，但是匿名函数引用到函数外的n,因此这个匿名函数就和n构成一个闭包。

（2）反复调用f函数，因为n只初始化一次，每一次调用就进行累计。

（3）闭包的关键，是返回的函数使用了哪些变量。

###### 6.7.2 闭包的实践
```go
package main

import (
    "strings"
)

func makeSuffix(suffix string) func (string) string {
    return func (name string) string {
        if !strings.HasSuffix(name,suffix) {
            return name + suffix
        }       
    }
}

f := makeSuffix(".jpg")

fmt.Println("文件名为=".f2("winter"))//winter.jpg
fmt.Println("文件名为=".f2("winter.jpg"))//winter.jpg
fmt.Println("文件名为=".f2("winter.avi"))//winter.avi.jpg
```
**总结：**

（1）返回的匿名函数和makeSuffix(suffix string)d的suffix变量组合成一个闭包，因为返回的函数引用到suffix变量。

（2）体会一下闭包的好处：如果使用传统的方法，亦可以实现功能。但是传统方法需要每次都传入后缀名，比如.jpg，而闭包因为可以
     保留上次引用的某个值，所以我们传入一次就可以反复使用。

#### 6.8 函数-defer

###### 6.8.1 基本介绍

函数中，程序员经常需要创建资源（比如：数据库连接，文件句柄，锁等）。为了在函数执行完毕后，及时释放资源，Go的设计者提供了defer(延时机制)

```go
package main

import (
    "fmt"
)

func sum(n1,n2 int) int {
    //当执行defer时，其后面的语句会压入defer栈中；当函数执行完毕后再执行，先入后出执行。
    defer fmt.Println("ok1 n1=",n1)
    defer fmt.Println("ok2 n2=",n2)
    res := n1 + n2 
    fmt.Println("ok3 res=",res)
    return res
}

func main()  {
    res := sum(10,20)
    fmt.Println("res=",res)
}
```
###### 6.8.2 细节说明

1) 当go执行到一个defer时，不会立即执行defer后的语句，而是将defer后的语句先压入到defer栈中，然后继续执行函数下一个语句。
2) 当函数执行完毕后，再从defer栈中，先入后出的方式执行语句。
3) 在defer将语句放入到栈中时，也会将相关的值拷贝同时入栈。
```go
package main

import (
    "fmt"
)

func sum(n1,n2 int) int {
    defer fmt.Println("ok1 n1=",n1)//10
    defer fmt.Println("ok2 n2=",n2)//20
    n1++//11
    n2++//21
    res := n1 + n2//32 
    fmt.Println("ok3 res=",res)
    return res
}

func main()  {
    res := sum(10,20)
    fmt.Println("res=",res)
}
```

###### 6.8.3 defer的实践

defer最主要的价值是在函数执行完毕后，可以及时释放函数创建的资源。

#### 6.9 变量的作用域

###### 6.9.1 说明

1) 局部变量：函数内部申明/定义的变量，作用域仅限于函数内部。
2) 全局变量：函数外部声明的变量，作用域在整个包中都有效，如果首字母大写，则作用域为整个程序有效。
3) 如果变量是一个代码块，比如for/if中，那么这个变量的作用域就在该代码块中。

#### 6.10 字符串的系统函数

1) len(str):内嵌函数，不用引入包，直接使用，对应手册里的builtin。字符串的字节长度。
2) 字符串遍历，同时处理有中文的问题r = []rune(str)
3) 字符串转整数： n,err := strconv.Atoi("12")
4) 整数转为字符串：str := strconv.Itoa("hello go")
5) 字符串转为[]byte:var bytes = []byte("hello go")//打印出每个字符的ascll码
6) []byte转为字符串:str = string([]byte{97,98,99})
7) 10进制转为2,8,16进制：str := strconv.FormatInt(123,2)//2-->8,16,返回对应字符串
8) 查找子串是否在指定的字符串中:strings.Contains("seafood","foo")//true
9) 统计一个字符串有几个指定的子串：strings.Count("ceheese","e")//4
10) 不区分大小写的字符串比较（==是区分字母大小写的）：fmt.Println(strings.EqualFold("abc","Abc"))
11) 返回子串在字符串中第一次出现的index值，如果没有，返回-1：strings.Index("NLT_abc","abc")
12) 返回子串在字符串最后一次出现的index，如果没有，返回-1：strings.LastIndex("go golang",""go")
13) 将指定的子串替换成另一个子串：strings.Replace("go go hello","go","go语言",n) n可以指定你希望替换的个数，如果n=-1表示全部替换。
14) 按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组：strings.Split("hello,world,ok",",")
15) 将字符串的字母进行大小写的转换：strings.ToLower("Go")//strings.ToUpper("Go")
16) 将字符串左右两边的空格去掉：strings.TrimSpace(" sdfds sdfd dsfdg    ")
17) 将字符串的左右两边指定的字符串去掉，：strings.Trim("! he！llo! "," !")//去掉左右两边的空格和！
18) 将字符串左边指定的字符去掉：strings.TrimLeft("!hello"," !")
19) 将字符串右边指定的字符去掉：strings.TrimRight("!hello"," !")
20) 判断字符串是否以指定上午字符串开头：strings.HasPrefix("ftp://192.168.0.0.1","ftp")
21) 判断字符串是否以指定上午字符串结尾：strings.HasPrefix("adf.jpg",".jpg")

#### 6.11 日期和时间相关函数

1) 需要引入time包
2) time.Time类型，用于表示时间//now := time.Now()
3) 通过now获取年月日时分秒:now.Year/Month/Day/Hour/Minute/Second
4) 格式化日期：
    （1）用上述的函数拼接
    （2）fmt.Printf(now.Format("2006/01/02 15:04:05"))//各个数字是固定的，必须这样写；但数字间的字符串间隔符号可以换。
         fmt.Printf(now.Format("2006-01-02"))
         fmt.Printf(now.Format("15:04:05"))
         fmt.Printf(now.Format("15"))

5) 时间和常量
    const(
        Nanosecond Duration = 1 //纳秒
        Microsecond = 1000*Nanosecond //微秒
        Millisecond = 1000*Microsecond //毫秒
        Second = 60*Millisecond //秒
        Minute = 60*Second //分钟
        Hour = 60*Minute //小时
    )
    常量的作用：在程序中可用于获取指定时间单位的时间，比如想获得100毫秒
    100*time.Millisecond
    
6) 休眠：结合time.Sleep()休眠来完成逻辑
7) 获取unix时间戳和unixnano时间戳

#### 6.12 Golang的内置函数（builtin）

1) len:用来求长度，比如string,array,slice,map,channel
2) new：用来分配内存，主要用来分配值类型，比如int,float32,struct,返回的是指针
3) make：用来分配内存，主要用来分配引用类型，比如channel,map,slice。

#### 6.13 错误处理

#### 6.13.1 基本说明

1) Go语言追求简洁优雅，所以不支持try...catch...finally
2) Go语言引入的处理方式为：defer,panic,recover
3) 这几个异常的使用场景可以简单描述为：Go中抛出一个panic异常,然后在defer中通过recover捕获这个异常，然后正常处理。

**defer+recover错误处理：**

```go
package main

import "fmt"

func test()  {
	//使用defer和recover来捕获和处理异常
	defer func(){
		err := recover()
		if err != nil {
			fmt.Println("err=",err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1/num2
	fmt.Println("res=",res)
}

func main()  {
	test()
	fmt.Println("main后面的代码...")
}
```
说明：进行错误处理后，程序不会轻易挂掉，如果加入预警代码，可以让程序更加更加健壮。

#### 6.13.2 自定义错误

Golang中支持自定义错误，使用errors.New和panic内置函数。

1) errors.New("错误说明")，会返回一个error类型的值，表示一个错误。
2) panic内置函数，接收一个interface{}类型的值（也就是任何值）作为参数，可以接收error类型的变量，输出错误信息，并退出程序。

```go
package main

import (
	"errors"
	"fmt"
)

func test()  {
	//使用defer和recover来捕获和处理异常
	defer func(){
		err := recover()
		if err != nil {
			fmt.Println("err=",err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1/num2
	fmt.Println("res=",res)
}

func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	}else{
		return errors.New("读取文件错误")
	}
}

func test2()  {
	err := readConf("config2.ini")
	if err != nil {
		//如果读取文件错误，就输出错误，并终止程序
		panic(err)
	}
	fmt.Println("程序继续执行...")
}

func main()  {
	test2()
	fmt.Println("main后面的代码...")
}

```
## 7 数组与切片

数组是可以存放多个同一类型数据的数据类型。数组是值类型。

#### 7.1 数组的定义和内存分布

###### 7.1.1 定义

```go
var 数组名 [数组大小]数据类型

var a [5]int

赋初值 a[0] = 1 a[1] = 3....
```
###### 7.1.2 数组内存图

1) 数组的地址可以通过&取出
2) 数组的地址就是数组第一个元素的地址，第二个元素的地址就是第一个元素地址加数组元素类型占用的字节数。

###### 7.1.3 四种初始化数组方式
```go
var numArr1 [3]int = [3]int{1,2,3}
var numArr1 = [3]int{1,2,3}
var numArr1 = [...]int{1,2,3}
//指定元素值对应的下标
var names = [3]string{1:"tom",2:"jack",0:"ke"}
```

#### 7.2 数组的遍历
方式一：for循环。
方式二：for-range结构遍历

###### 7.2.1 基本介绍
```go
for index, value := range array01 {
    //代码块
}
```

**说明：**

1) index是数组的下标
2) value是该下标的值
3) 上述两个变量都是for循环内部可见的变量
4) 遍历数组的时候，如果不想使用下标index，可以直接把下标index标为下划线_
5) index和value变量名随意，只是一般命名为index和value

```go
package main

import (
	"fmt"
)

func main()  {
	var arrtest1 = [5]string{"你好","地方","递归","改变","风格"}

	for index, value := range arrtest1 {
		fmt.Printf("数组的第%d个元素的值是%v\n",index+1,value)
	}
}
```

###### 7.2.2 数组的注意事项和细节

1) 数组是多个相同类型的数据的组合，一个数组一旦声明了，其长度是固定的，不能动态变化。
2) var arr []int ,这时arr就是一个slice切片。
3) 数组中的元素是任意数据类型，包括值类型和引用类型，但不能混用。
4) 数组创建以后，如果没有赋值，一般会默认0值。
5) 使用数组的步骤：1.声明数组开辟空间-->2.给数组各元素赋值-->3.使用数组
6) 数组下标从0开始
7) 数组下标必须在指定范围内使用，否则报panic异常，数组越界。
8) Go的数组属于值类型。
9) 如果想在其它函数内取修改原来的数组，可以使用引用传递（指针）。
10) 长度是数组类型数据的一部分，在传递函数参数时，需要考虑数组的长度。


#### 7.3 切片 slice

###### 7.3.1 切片的基本介绍

1) 切片是数组的引用，因此切片是引用类型。
2) 切片的使用和数组类似，遍历切片，访问切片的元素和求切片的长度len()都是一样的。
3) 切片的长度是变化的，可以理解为一个动态数组
4) 切片的基本语法：var 变量名 []类型 //var aslice []int//不用表明切片的长度

###### 7.3.2 切片的案例

```go
package main

import (
	"fmt"
)

func main()  {
	var intArr [5]int = [...]int{1,23,45,21,46}

	aslice := intArr[1:3]
	fmt.Println(aslice)
	fmt.Println("切片的长度是：",len(aslice))
	fmt.Println("切片的容量是：",cap(aslice))
}
```

###### 7.3.3 切片在内存中的形式

slice从底层来说，就是一个结构体Struct,包含引用数组引用的第一个元素的地址，和slice的长度以及容量。
```go
type slice struct {
        ptr *[2]int
        len int
        cap int
    }
```
 ###### 7.3.4 切片的使用   
- 方式一：定义一个切片，然后让切片去引用一个已经创建好的数组
```go
var arr [3]int{1,2,4}
var sliceA = arr[1:2]
```
-方式二：make创建切片

var 切片名 []类型 = make([]类型,len,[cap])//cap可选
```go
package main

import (
	"fmt"
)

func main() {
	var aslice []float64 = make([]float64,5)//对于slice必须先make再使用。var aslice []float64不行的
	fmt.Println(aslice)
	fmt.Println("切片的长度是：", len(aslice))
	fmt.Println("切片的容量是：", cap(aslice))
}
```
**总结：**

1) 通过make方式创建切片可以指定切片的大小和容量
2) 如果没有给切片的各个元素赋值，那么就会使用默认值。
3) 通过make创建的切片，其引用的数组是由make底层维护的，对外不可见，即只能通过slice去访问各个元素。


-方式三：定义一个切片，直接就指定具体数组，使用原理类似make

```go
func main() {
    var aslice []int = []int {1,3,5}
    fmt.Println(aslice)
} 
```
**方式一和方式二的区别：**

1) 方式1是直接引用数组，这个数组是事先存在的，程序员是可见的。
2) 方式2是通过make创建切片，make也会创建一个数组，是由切片在底层进行维护，程序员不可见。

###### 7.3.5 切片的遍历
 for循环和for-range循环

###### 7.3.6 切片的注意细节
1) 切片初始化时： var slice = arr[startIndex,endIndex],左闭右开
2) 切片初始化时，仍不能越界。范围在[0,len(arr)]之间，但是可以动态增长。
3) cap是内置函数，用于统计切片的容量，即最大可以存放多少个元素。
4) 切片定义完还不能使用，因为本身是一个空的，需要让其引用到一个数组或make一个空间供切片来使用。
5) 切片可以继续切片。
```go
var arr [5]int = [...]int{1,2,3,5,8}

var slice1 = arr[:]
var slice2 = slice1[1,2]//2
```
6) 使用append内置函数，可以对切片进行动态追加。
```go
var arr [5]int = [5]int{1,2,3,4,5}
var slice1 = arr[:]
slice1 = append(slice1,10,20,30)//追加具体元素
slice1 = append(slice1,slice1...)//追加切片，注意三个点
```
append操作的本质是对数组扩容，go底层会创建一个新的数组newArr,然后slice原来包含的元素拷贝到新的数组newArr，slice重新引用新的
数组newArr,newArr是在底层维护，不可见的。
7) 切片的拷贝copy
```go
copy(param1,param2)//param1和param2都是切片类型
```

###### 7.3.6 string和slice

1) string其实也是一个切片，底层是一个byte数组，因此string也可以进行切片处理
```go
package main

import (
	"fmt"
)

func main()  {
	str := "hello@abc"
	slice := str[6:]
	fmt.Println("slice=",slice)
}
```
2) string是不可变的，也就是说不通过str[0]='z'来修改字符串
3) 如果一定要修改字符串，可以先将string转为[]byte或者[]rune,然后再转为string
```go
str := "hello@abc"
	slice := str[6:]
	fmt.Println("slice=",slice)
	slice2 := []byte(str)
	slice2[0] = 'x'
	str2 := string(slice2)
	fmt.Println("str2",str2)
	slice3 := []rune(str)
	slice3[1] = '你'
	str3 := string(slice3)
	fmt.Println("str3",str3)
```





## 8 map p173

## 9 面向对象编程 p181-225



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




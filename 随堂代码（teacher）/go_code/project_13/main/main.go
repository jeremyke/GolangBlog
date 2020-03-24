//统计字符串的长度：按字节len(str)
//字符串遍历，可以解决中文：r:=[]rune(str)
//字符串转整数：n,err:=strconv.Atoi("12")
//整数转字符串：str = strconv.ltoa(123456)
//字符串转[]byte:var bytes=[]byte("hello")
//[]byte转字符串:str=string([]byte{98,54})
//10进制转二进制八进制十六进制：str=strconv.Formatlnt(123,2)//2->8,16
//查找字符串是否再制定的字符串中：string.Contains("seafood","foo")//true
//统计一个字符串有几个指定的子串：string.Count("chinese","e")//2
//不区分大小写的字符串比较（==是区分字母大小写的）:fmt.Println(string.EqualFold("abc","Abc"))//true
//返回子串再字符串中第一次出现的index值，如果没有返回-1：string.Index("NLT_abc","abc")//4
//返回子串再字符串最后一次出现的index，如果没有返回-1：string.LastIndex("go golang","go")
//将指定的子串替换为另一个子串：strings.Replace("go go hello","go","go语言"，n)n可以指定希望替换几个，如果n==-1表示全部替换
//按照指定的某哥字符为分割标志，将一个字符串拆分成字符串数组：string.Split("hello,world,ok",",")
//将字符串的字母进行大小写的转换：strings.ToLower("Go")//go strings.ToUpper("Go")//GO
//将字符串左右两边的空格去掉：strings.TrimSpace("tn a lone gopher ntm    ")
//将字符串左右两边指定的字符去掉：strings.Trim("!hello!","!")
//将字符串左边指定的字符去掉：trings.TrimLeft("!hello!","!")
//将字符串右边指定的字符去掉：trings.TrimRight("!hello!","!")
//判断字符串是否以指定的字符串开头：strings.HasPrefix("ftp://192.168.10.1","ftp")//true
//判断字符串是否以指定的字符串开头：strings.HasSuffix("NLT_abc.jpg","abc")//false

package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for {
		i++
		fmt.Println(i)
		//休眠
		time.sleep(100 * time.Millisecond)
		if i == 10 {
			break
		}
	}
	//unix获得时间戳
	//使用defer + recover 来捕获和处理异常
	defer func() {
		err := recover()
		if err != nil { //说明有错误
			fmt.Println("err = ", err)
			fmt.Println("发送邮件给管理员cuimohan123456@163.com")
		} //这里可以把代码错误发送给管理员
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res = ", res)
}

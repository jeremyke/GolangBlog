package main

import "fmt"

func main() {
	//fmt.Println("Tom\tJack")//转义字符
	//fmt.Println("Tom\nJack")//换行
	//fmt.Println("Tom\\Jack")//第一个转义符，第二个路径的斜杠输出
	//fmt.Println("Tom\"tJack")//原意想有的东西
	//fmt.Println("Tom\rJack")//从当前行的最前面开始输出，覆盖掉之前的内容
	///////下面为练习///////
	fmt.Println("姓名\t年龄\t籍贯\t地址\n崔默涵\t21\t山东\t青岛科技大学") //不能让两个main函数放在同一个包内
	//重新建一个包（新建一个文件夹）
}

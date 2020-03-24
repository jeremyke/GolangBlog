package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	//os.Args
)

func main() {
	//打开一个文件
	file, err := os.Open("...")
	if err != nil {
		fmt.Println("...", err)
	}
	//fmt.Printf("file = %v",file)
	defer file.Close() //要及时关闭file，否则会有内存泄漏
	//创建一个reader，自带缓冲
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //读到换行就结束一次
		if err == io.EOF {                  //表示文件的末尾
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取结束")

	//创建一个一次性读取所有文件的，不适合用于大型文件
	//因为ReadFile 没有显示的Open文件，
	//所以不需要打开或者关闭，
	//打开和关闭都被封装到函数本身里面了
	file_1 := "..."
	content, err := ioutil.ReadFile(file_1)
	if err != nil {
		//fmt.Println(...)
	}
	//fmt.Printf("%v", content)//[]byte
	fmt.Printf("%v", string(content)) //输出真实内容

	//写入文件
	//1.打开文件
	filePath := "d:/谷歌下载/嘤嘤嘤.txt"
	file_2, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("错了：", err)
		return
	}
	//及时环比file句柄
	defer file_2.Close()
	str := "hello 嘤嘤嘤怪\n"
	//写入时使用带缓存的*Writer
	writer := bufio.NewWriter(file_2)
	for i := 0; i <= 5; i++ {
		writer.WriteString(str)
	}
	//因为writer 是自带缓存的，因此在调用WriterString时
	//内容是先写到缓存中的,所以西药调用flush将缓存的内容
	//真正写入到缓存中，否则文件中会没有数据
	writer.Flush()

	//覆盖已有文件，换上新的修改
	//filePath := "d:/谷歌下载/嘤嘤嘤.txt"
	//file_2, err := os.OpenFile(filePath, os.O_WRONLY | os.O_TRUNC, 0666)
	//if err != nil{
	//fmt.Println("错了：", err)
	//return
	//}
	//defer file_2.Close()
	//str := "hello 嘤嘤嘤怪 我修改你了\r\n"
	//writer := bufio.NewWriter(file_2)
	//for i := 0; i <= 5; i++{
	//writer.WriteString(str)
	//}

	//writer.Flush()

	//在原来的内容上追加
	//filePath := "d:/谷歌下载/嘤嘤嘤.txt"
	//file_2, err := os.OpenFile(filePath, os.O_WRONLY | os.O_APPEND, 0666)
	//if err != nil{
	//fmt.Println("错了：", err)
	//return
	//}
	//defer file_2.Close()
	//str := "hello 嘤嘤嘤怪 我修改你了 我有追加你了\r\n"
	//writer := bufio.NewWriter(file_2)
	//for i := 0; i <= 5; i++{
	//writer.WriteString(str)
	//}

	//writer.Flush()

	//显示原先内容然后追加
	//filePath := "d:/谷歌下载/嘤嘤嘤.txt"
	//file_2, err := os.OpenFile(filePath, os.O_RDWR | os.O_APPEND, 0666)
	//if err != nil{
	//fmt.Println("错了：", err)
	//return
	//}
	//defer file_2.Close()
	//先读取，并显示在终端
	//reader := bufio.NewReader(file)
	//for {
	//str, err := reader.ReadString('\n')
	//if err == io.EOF {
	//break
	//}
	//继续读，显示到终端
	//fmt.Print()
	//}
	//str := "hello 嘤嘤嘤怪 我修改你了 还要读取追加你\r\n"
	//writer := bufio.NewWriter(file_2)
	//for i := 0; i <= 5; i++{
	//writer.WriteString(str)
	//}

	//writer.Flush()

	//将一个文件的内容写入到另一个文件中
	//file_5 := "..."
	//file_4 := "..."

	//ioutil.ReadFile(file_4)
	//if err != nil{
	//fmt.Println(...)
	//return
	//}
	//err = ioutil.WriteFile(file_5, data, 0666)
	//if err != nil{
	//fmt.Printf(...)
	//}

	//拷贝文件
	//file_3 := "..."
	//file_6 := "..."
	//CopyFile(file_3, file_6)
	//_,err := CopyFile(file_3, file_6)
	//if err == nil{
	//fmt.Println("拷贝完成")
	//}else {
	//fmt.Printf("拷贝失败 ，err =%v", err)
	//}

	//定义变量，接受命令行的参数
	var user string
	var pwd string
	var host string
	var port int
	//&user就是接受用户命令行中输入的 -u 后面的内容
	//”u“ 就是 -u 指定的参数
	//""，默认值
	//"用户说明" 说明
	flag.StringVar(&user, "u", "", "用户说明")
	flag.StringVar(&pwd, "pwd", "", "用户说明")
	flag.StringVar(&host, "h", "", "用户说明")
	flag.IntVar(&port, "port", 3306, "端口号")
	//重要操作，转换
	flag.Parse()
	fmt.Printf("user = %v pwd = %v host = %v port = %v",
		user, pwd, host, port)
}

//拷贝文件
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcfile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err = %v\n", err)
	}
	defer srcfile.Close()
	reader := bufio.NewReader(srcfile)
	//打开dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err = %v\n", err)
		return
	}
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()
	return io.Copy(writer, reader)
}

type HeroNode struct {
	no   int
	name string
	next *HeroNode //这里表示指向下一个节点
}

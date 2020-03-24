package main
import(
	"fmt"
)

func main(){
	var b = false//布尔类型
	fmt.Println("b = ",b)
	fmt.Println("b占用的空间"，unsafe.Sizeof(b))//不能像其他语言一样用0和1取代
}
package main()
import(
	"fmt"
	"math"
	"time"
)

func main(){
	var array [5]int
	//防止进行伪随机
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(array); i++{
		array[i] = rand.Intn(100)
	}

	fmt.Println("交换前为: ",array)
	temp := 0
	for i := 0; i< len(array) /2; i++{
		temp = array[len(array) - 1 - i]
		array[len(array) - 1 - i] = array[i]
		array[i] = temp
	}
	fmt.Println("交换后为: ",array)
}
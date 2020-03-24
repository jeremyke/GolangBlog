package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano()) //纳秒，Unix为秒
	n := rand.Intn(100) + 1
	fmt.Println("输出的数字：", n)
}

func cal(a float64, n2 float64, operator byte) float64 {

}

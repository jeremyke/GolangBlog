package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var arrtest1 = [5]string{"你好", "地方", "递归", "改变", "风格"}

	for index, value := range arrtest1 {
		fmt.Printf("数组的第%d个元素的值是%v\n", index+1, value)
	}

	zimou := zimou()
	fmt.Printf("%c\n", zimou)
	getMax([3]int{13, 2, 6})
	getTotal([3]int{13, 3, 6})
	fanzhuan()
}

func zimou() [26]byte {
	var zimou [26]byte
	for i := 0; i < len(zimou); i++ {
		zimou[i] = 'A' + byte(i)
	}
	return zimou
}

func getMax(arr [3]int) {
	var max, maxIndex int
	for index, value := range arr {
		if index == 0 {
			max = value
			maxIndex = index
		} else {
			if max < value {
				max = value
				maxIndex = index
			}
		}
	}
	fmt.Println("最大值为", max, "其索引为", maxIndex)
}

func getTotal(arr [3]int) {
	var total int
	for _, value := range arr {
		total += value
	}
	fmt.Println("和为", total, "平均值为", float64(total)/float64(len(arr)))
}

func fanzhuan() {
	var intArr [3]int
	for i := 0; i < len(intArr); i++ {
		rand.Seed(time.Now().UnixNano())
		intArr[i] = rand.Intn(100)
	}
	fmt.Println("随机生成的数组是,", intArr)
	var intArr2 [3]int
	for index, value := range intArr {
		intArr2[2-index] = value
	}
	fmt.Println("翻转数组是,", intArr2)
}

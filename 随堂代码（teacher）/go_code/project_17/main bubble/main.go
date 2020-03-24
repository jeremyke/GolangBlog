package main
import (
	"fmt"
)

//冒泡排序
func BubbleSort(arrary *[5]int){
		temp := 0
		for i := 0; i < 4; i++{
		for j:= 0; j < 4 - i; j++{
			if (*arrary)[j] > (*arrary)[j + 1]{
				temp = (*arrary)[j]
				(*arrary)[j] = *(arrary)[j + 1]
				(*arrary)[j + 1] = temp
			}
		}
	}
}


func main(){
	arrary := [5]int{15,23,13,84,55}
	BubbleSort(arrary)
	fmt.Println()

	name := [4]string{"草","木","花","语"}
	var heroname = " "
	fmt.Println("你要查找的物品是：")
	fmt.Scanln(&heroname)
	//顺序查找
	for i := 0; i < len(name); i++{
		if heroname == names[i]{
			fmt.Println("查询结果为",heroname)
			break}else if i == (len(name) - 1){
				fmt.Printf("%v查询结果为无",heroname)
			}
	}
	//二分查找
	var carrary := [6]int{1,2,3,4,5,6}
	BinaryFind(&carrary,0,(len(carrary)-1),1)
	
}

func BinaryFind(carrary *[6]int,leftIndex int,rightIndex int,findVal int){
	middle := (leftIndex + rightIndex)/2

	if leftIndex > rightIndex{
		fmt.Printf("没有找到")
		return 
	}

	if (*carrary)[middle] > findVal{
		BinaryFind(carrary,leftIndex,middle - 1,findVal)
	}else {
		fmt.Println("找到了")
	}

	//二维数组
	var varrary [5][6]int
	varrary[3][6] = 1

	fmt.Println("二维数组")

}
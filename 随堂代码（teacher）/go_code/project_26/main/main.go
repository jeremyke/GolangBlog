package main
import (
	"fmt"
	"sort"
	"math\rand"
)

type Hero struct{
	Name string
	Age int
}

type HeroSlice []Hero

func (hs HeroSlice) Len() int{
	return (hs)
}

func (hs HeroSlice) Less(i, j int) bool{
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice) Swep(i, j int) bool{
	temp := hs[i]
	hs[i] = hs[j]
	hs[j] = temp
	//hs[i], hs[j] = hs[j], hs[i]
}

func main(){
	var intSlice = []int{0,-1,20,30,93}
	sort.Ints(intSlice)
	fmt.Println(intSlice)


	var heros HeroSlice
	for i := 0; i < 11; i++{
		hero := Hero{
			Name : fmt.Sprintf("英雄~%d", rand.Intn(100)),
			Age : rand.Intn(100)
		}
		heros = append(heros, hero)
	}
	for _, v := range heros{//排序前
		fmt.Println(v)
	}
	sort.Sort(heros)
	for _, v := range heros{//排序后
		fmt.Println(v)
}
package main
import (
	"fmt"
	"go_code\project_22\model"
)



//使用工厂模式，在import中可以引进“路径”，可以任意使用，即使首字母小写
func main(){
	var stu = model.NewStudent(cmh,92.3)
	fmt.Println(stu)

	//p := model.NewPerson("cmh")
	//fmt.Println(p)
}
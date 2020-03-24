package main

import "fmt"

func main() {
	var num int = 9
	fmt.Println("num address=%v", &num)
	var ptr *int
	ptr = &num
	*ptr = 10 //这里修改的是num的值，会变化
	fmt.Println("num =", num)
}

//不能被用做变量名的有25个，他们分别是：
//break;default;func;interface;select;
//case;defer;go;map;struct;
//chan;else;goto;package;switch;
//continue;for;import;return;var
//%取余运算的公式：a % b = a - a / b * b

package main

import (
	"fmt"
)

type Goods struct {
	Name  string
	Price float64
}

type Brand struct {
	Name    string
	Address string
}

type Tv struct {
	Goods
	Brand
}

type Tv2 struct {
	*Goods
	*Brand
}

func main() {
	tv := Tv{
		Goods{
			Name:  "电视机",
			Price: 199.9,
		},
		Brand{
			Name:    "LG",
			Address: "韩国",
		},
	}
	tv2 := Tv2{
		&Goods{
			Name:  "手机",
			Price: 299.5,
		},
		&Brand{
			Name:    "小米",
			Address: "北京",
		},
	}

	fmt.Println(tv)
	fmt.Println(*tv2.Goods, *tv2.Brand)
}

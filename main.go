package main

import (
	"fmt"
	// "go_learning/calc"
	// "github.com/shopspring/decimal"
)

type Usber interface {
	start(string, string) string
	stop()
}

type Phone struct {
	Name string
}

func (p Phone) start() {
	fmt.Println(p.Name, "启动")
}
func (p Phone) stop() {
	fmt.Println(p.Name, "关机")
}
func main() {
	// price, err := decimal.NewFromString("136.02")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(price)
	// sum := calc.Add(10, 2)
	// fmt.Println(sum)
	p := Phone{
		Name: "huawei",
	}
	p.start()
}

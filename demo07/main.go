package main

import (
	"fmt"
	"reflect"
)

// 反射或取对象类型和属性
type myInt int
type Person struct {
	Name string
}

func reflectFn(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Println(v)
}
func main() {
	a := 10
	b := 23.4
	c := true
	d := "nihao"
	reflectFn(a)
	reflectFn(b)
	reflectFn(c)
	reflectFn(d)

	var e myInt = 34
	var f = Person{
		Name: "zhnagsna",
	}
	reflectFn(e)
	reflectFn(f)
}

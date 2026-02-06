package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("/Users/linghang/go_learning/demo02/main.go")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file)
	var tempSlice = make([]byte, 128)
	n, err := file.Read(tempSlice)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)
	fmt.Println(string(tempSlice))
}

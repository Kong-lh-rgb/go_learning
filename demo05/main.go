package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("ntestihao")
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done()
}

func main() {
	cpu := runtime.NumCPU()
	fmt.Println(cpu)
	wg.Add(1)
	go test() //表示开启一个携程
	for i := 0; i < 10; i++ {
		fmt.Println("nihao")
		time.Sleep(time.Millisecond * 50)
	}
	// time.Sleep()
	wg.Wait()
}

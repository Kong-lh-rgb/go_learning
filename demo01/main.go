package main

import "fmt"

type Usber interface {
	start()
	stop()
}

type Phone struct {
	Name string
}

func (p Phone) start() {
	fmt.Println(p.Name, "启动了")
}
func (p Phone) stop() {
	fmt.Println(p.Name, "关机了")
}

type Carmer struct {
	Name string
}

func (c Carmer) start() {
	fmt.Println(c.Name, "开机啦")
}
func (c Carmer) stop() {
	fmt.Println(c.Name, "关机啦")
}

func computerwork(device Usber) {
	device.start()
	device.stop()
}

func main() {
	p := Phone{Name: "huawei"}
	c := Carmer{Name: "sony"}
	computerwork(p)
	computerwork(c)
}

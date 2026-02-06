package main

import "fmt"

func main() {

	var userinfo = make(map[string]interface{})
	userinfo["username"] = "zds"
	userinfo["age"] = 20
	userinfo["hobby"] = []string{"sj", "cf"}
	fmt.Println(userinfo["age"])
}

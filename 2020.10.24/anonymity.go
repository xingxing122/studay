package main

import (
	"fmt"
	"time"
)

func main() {
	// 匿名结构体
	var   user  struct {
		id int
		name string
		tel string
		add string
		birthday  time.Time
	}

	fmt.Printf("%T, %#v\n", user,user)

	user.id = 10
	user.name = "kk"
	fmt.Printf("%T,%#v\n",user,user)
	// 字面量
	user  = struct {
		id       int
		name     string
		tel      string
		add      string
		birthday time.Time
	}{id:10,name:"kk"}
	fmt.Printf("%T, %#v\n", user,user)

}




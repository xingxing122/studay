package main

import "fmt"

func main() {

	var userinfo = make(map[string][]string)

	userinfo["hobby"] = []string{
		"吃饭",
		"睡觉",
		"敲代码",
		"跑步",
	}

	fmt.Println(userinfo)

	for _, v := range userinfo {
		for _, value := range v {
			fmt.Println(value)
		}
	}

}

package main

import "fmt"

func main() {

	var userinfo = make([]map[string]string, 3, 3)

	if userinfo[0] == nil {
		userinfo[0] = make(map[string]string)
		userinfo[0]["username"] = "张三"
		userinfo[0]["age"] = "20"
		userinfo[0]["height"] = "180cm"
	}
	fmt.Println(userinfo)

}

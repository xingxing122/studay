package main

import "fmt"

func main() {

	var userinfo = make(map[string]string)
	userinfo["username"] = "张三"
	userinfo["age"] = "20"
	userinfo["sex"] = "男"

	userinfo["username"] = "李四"
	fmt.Println(userinfo)

	delete(userinfo, "username")

	fmt.Println(userinfo)
}

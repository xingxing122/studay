package main

import "fmt"

func getUserinfo() (string, int) {

	return "zhangsan", 10
}

func main() {

	var username string
	var age = 20
	var sex = "男"
	fmt.Println(username, age, sex)

	var a1, a2 string

	a1 = "aaaa"
	a2 = "123"
	fmt.Println(a1, a2)

	a, b, c := 12, 13, "c"
	fmt.Println(a, b, c)
	fmt.Printf("a类型: %T b类型:%T c类型：%T", a, b, c)

	var _, ag = getUserinfo()

	fmt.Println(ag)
}

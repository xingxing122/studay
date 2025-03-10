package main

import "fmt"

func main() {

	var a = 'a'
	fmt.Printf("值： %v 类型：%T", a, a)

	var str = "this"

	fmt.Printf("值： %v 类型:%T", str[2], str[2])

	//6.通过循环输出字符串里面的字符

	//s := "golang"
	//for i := 0; i < len(s); i++ {
	//	fmt.Printf("%v(%c)", s[i], s[i])
	// }

	s := "你好golang"

	for _, v := range s {
		fmt.Printf("%v(%c)", v, v)

	}

}

package main

import "fmt"

func main() {

	s1 := "hello worold"

	fmt.Println(s1)

	s2 := `
	第一行
	第二行
	第三行
	`
	fmt.Println(s2)
	ss := "美国第一"
	ss_rune := []rune(ss)

	fmt.Println(ss_rune)
	fmt.Println(string(ss_rune[:2]))
	s_zh := "中国" + string(ss_rune[2:])
	fmt.Println(s_zh)

}

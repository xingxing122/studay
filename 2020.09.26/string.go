package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var   name string = "kk"
	var   desc  string = "i love china"

    fmt.Println(name,desc)

	fmt.Println(desc[1:5])
	var  txt = "我爱中国"
	for i,v := range  txt {
		fmt.Printf("%d, %q\n",i,v)
	}

	fmt.Println(utf8.RuneCountInString(txt))
	fmt.Println(utf8.RuneCountInString(desc))

	//字符串 ==> byte

	fmt.Println([]byte(desc))
	fmt.Println(string([]byte(desc)))

	fmt.Println([]byte(txt))



}

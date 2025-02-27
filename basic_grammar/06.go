package main

import (
	"fmt"
	"strings"
)

func main() {
	//var str1 string = "你好golang"
	//var str2 = "你好go"
	// str3 := "你好golang"

	//fmt.Println("%v--%T", str1, str1)

	//fmt.Println("%v--%T", str2, str2)
	//fmt.Println("%v--%T", str3, str3)

	//str1 := `this is str
	//this is str
	//`
	//fmt.Println(str1)

	var str1 = "123-456-789"
	arr := strings.Split(str1, "-")
	fmt.Println(arr)

}

package main

import "fmt"

func main() {

	var str string = "hello golang"

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])

	}

	for i, value := range str {
		fmt.Printf("索引为：%d，具体de值为：%c \n", i, value)

	}
}

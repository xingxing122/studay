package main

import "fmt"

func main() {
	letters := "abcdefghi"

	for i := 0; i < len(letters); i++ {

		fmt.Printf("%c\n", letters[i])

	}

	letters = "我爱中国"
	for _, v := range letters {
		fmt.Printf("%q\n", v)

	}

}

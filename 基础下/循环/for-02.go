package main

import "fmt"

func main() {

	s := "hello world"

	for index, val := range s {
		fmt.Println(index, val)

	}

}

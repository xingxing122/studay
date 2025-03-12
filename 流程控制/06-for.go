package main

import "fmt"

func main() {

	var row = 3
	var column = 4

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

}

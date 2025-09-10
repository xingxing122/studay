package main

import "fmt"

func main() {

	score := 66
	if score > 90 {

		fmt.Println("优秀")
	} else if score > 80 {
		fmt.Println("良好")
	} else {
		fmt.Println("一般")
	}

}

package main

import "fmt"

func main() {
	a := []string{"bj", "sh", "gz"}

	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}

	//
	for index, value := range a {

		fmt.Println(index, value)
	}
	
}

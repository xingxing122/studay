package main

import "fmt"

var sum = 1

func main() {
	for i := 1; i <= 5; i++ {
		sum *= i
	}
	fmt.Println(sum)

}

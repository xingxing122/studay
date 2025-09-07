package main

import "fmt"

func main() {
	var a int8 = 4
	var a2 int8 = 100

	fmt.Println(a, a2)

	fmt.Printf("%T: %v\n", a, a)

	n := 500
	b := 400
	fmt.Println(n, b)
	fmt.Println(n > b)

}

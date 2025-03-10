package main

import "fmt"

func main() {

	var a = 34
	var b = 10

	t := a
	a = b
	b = t

	fmt.Printf("a =%v b=%v", a, b)

}

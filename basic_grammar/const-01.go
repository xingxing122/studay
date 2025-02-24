package main

import "fmt"

func main() {

	const (
		n1 = iota
		n2
		n3
	)

	fmt.Println(n1, n2, n3)

	var a = 123
	fmt.Println(a)

}

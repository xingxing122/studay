package main

import "fmt"

func main() {

	const name string = "kk"
	fmt.Println(name)

	const (
		A = "test"
		B
		C
		D = "testD"
		E
		F
	)
	fmt.Println(A, B, C, D, E, F)

}

package main

import "fmt"

type myInt int

type myFn func(int, int) int

type myFloat = float64

func main() {

	var a myInt = 10
	fmt.Printf("%v %T", a, a)

	var b myFloat = 12.5

	fmt.Printf("%v %T", b, b)

}

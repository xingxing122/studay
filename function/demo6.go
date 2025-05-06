package main

import "fmt"

type calc func(int, int) int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func main() {
	var c calc
	c = add
	fmt.Printf("C的类型: %T", c)

	fmt.Println(c(10, 5))

	f := sub
	fmt.Printf("c的类型：%T", f)

}

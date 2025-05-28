package main

import "fmt"

func main() {
	a := 10
	p := &a
	*p = 30
	fmt.Println(p)
	fmt.Println(*p)

}

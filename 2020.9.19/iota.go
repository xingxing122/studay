package main

import "fmt"
func main() {

	const (
		c1 int = iota
		c2
		c3
		c4
	)
	fmt.Println(c1,c2,c3,c4)
}

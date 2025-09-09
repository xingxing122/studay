package main

import "fmt"

func main() {
	//生命一个变量a
	a := 100
	fmt.Println(&a)
	fmt.Printf("%d\n", &a)

	b := &a

	fmt.Println(*b)

}

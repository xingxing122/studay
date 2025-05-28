package main

import "fmt"

func main() {

	a := 10

	p := &a

	fmt.Printf("a的值: %v a的类型 %T a的地址%p\n", a, a, &a)
	fmt.Println(*p)

}

package main

import "fmt"

func main() {
	/*	var arr1 []int
		fmt.Printf("%v  -- %T  - 长度: %v", arr1, arr1, len(arr1))*/

	a := [6]int{55, 56, 57, 58, 59, 60}
	b := a[:]
	fmt.Printf("%v - %T\n", b, b)

	c := a[2:5]
	fmt.Println(c)

}

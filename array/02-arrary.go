package main

import "fmt"

func main() {
	var arr1 = []int{1, 2, 3}
	arr2 := arr1
	arr1[0] = 11
	fmt.Println(arr1)
	fmt.Println(arr2)
}

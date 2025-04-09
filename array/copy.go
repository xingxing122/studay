package main

import "fmt"

func main() {
	/*	sliceA := []int{1, 2, 3, 45}
		seliceB := sliceA
		seliceB[0] = 11
		fmt.Println(sliceA)
		fmt.Println(seliceB)*/

	sliceA := []int{1, 2, 3, 4}
	sliceB := make([]int, 4, 4)

	copy(sliceA, sliceB)
	fmt.Println(sliceA)
	fmt.Println(sliceB)

}

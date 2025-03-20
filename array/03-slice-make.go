package main

import "fmt"

func main() {

	var sliceA = make([]int, 4, 8)

	//赋值
	sliceA[0] = 10
	sliceA[1] = 20
	fmt.Println(sliceA)
	fmt.Println(len(sliceA), cap(sliceA))

}

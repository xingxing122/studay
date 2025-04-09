package main

import "fmt"

func main() {

	/*	var sliceC []int
		fmt.Println("%v -- %v---%v", sliceC, len(sliceC), cap(sliceC))

		sliceC[0] = 1
		fmt.Println(sliceC)*/

	sliceB := []string{"php", "java", "go"}

	sliceB[2] = "goland"
	fmt.Println(sliceB)

}

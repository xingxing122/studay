package main

import "fmt"

func main() {

	var   ms [3][2] int
	fmt.Printf("%T\n", ms)
	fmt.Println(ms)


	fmt.Printf("%T  %v\n", ms[0],ms[0])

	ms = [...][2]int{
		0:[2]int{1,2},
		1:[2]int{3,4},
		2:[2]int{5,6},
	}
   fmt.Println(ms)

}

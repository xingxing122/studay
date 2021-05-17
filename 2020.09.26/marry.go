package main

import "fmt"

func main() {

	var   ms [3][2] int
	fmt.Printf("%T\n", ms)
	fmt.Println(ms)


	fmt.Printf("%T  %v\n", ms[0],ms[0])

	ms = [...][2]int{
		1:[2]int{1,2},
		2:[2]int{3,4},
		0:[2]int{5,6},
	}
   fmt.Println(ms)

	ms[0][1] = 100
	fmt.Println(ms)
	ms[1] = [2]int{103,105 }
	fmt.Println(ms)

	for i, v := range  ms {
		fmt.Println(i,v)
	}


}

package main

import "fmt"

func main() {
	/*	var age int = 18
		fmt.Println(&age)
		var ptr *int = &age

		fmt.Println(ptr)
	*/
	
	var num int = 10
	fmt.Println(num)

	var ptr *int = &num
	*ptr = 20
	fmt.Println(num)

}

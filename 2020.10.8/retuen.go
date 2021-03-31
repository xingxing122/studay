package main

import "fmt"

func  test(stop bool){
	fmt.Println("A")
	if stop {
		fmt.Println("B")
		return
		fmt.Println("C")
	}

	fmt.Println("F")
}

func main(){
	test(true)
	fmt.Println("-----------")
	test(false)
}

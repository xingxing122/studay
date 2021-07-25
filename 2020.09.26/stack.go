package main

import "fmt"

func   main(){

	stack := []int{}

	stack = append(stack,1)
	stack = append(stack,2)
	stack = append(stack,3)


	for len(stack)  !=0 {
		fmt.Println(stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
}
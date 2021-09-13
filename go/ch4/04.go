package main

import "fmt"

func   operate(a,b  int,op string) int{
	switch op {
	case "+":
		return  a + b
	case "-":
		return  a - b
	case "*":
		return  a * b
	case "/":
		return  a / b
	default:
		panic(fmt.Sprintf("not support  operate: %s",op))

	}
}



func main() {
    a, b  := 10,5
	fmt.Println(operate(a,b,"+"))
}

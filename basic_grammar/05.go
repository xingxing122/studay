package main

import "fmt"

func main() {

	var a float32 = 3.12

	fmt.Println("值: %v----%f,类型%T\n", a, a, a)

	var c float64 = 3.1416927

	fmt.Printf("%.2f", c)

	var f2 = 3.14e2
	fmt.Printf("%f--%T", f2, f2)

	var s = false
	if s {
		fmt.Println("true")
	}

	var f1 = false

	if f1 {
		fmt.Println("true")
	} else {
		fmt.Printf("false")
	}

}

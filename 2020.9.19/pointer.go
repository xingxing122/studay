package main

import "fmt"

func main() {

	 A := 2
	 B := A
	 fmt.Println(A,B)

	 C := &A

	 fmt.Println(C)
	 fmt.Println(*C)

	  *C = 5
	  fmt.Println(A)

}

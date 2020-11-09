package main

import "fmt"

func main() {

	const  NAME  string  = "xingxing"

	const PI = 3.1415926

	const  C1,C2 int = 1,2

	const (
        C3 string = "silence"
        C4 int     = 1
        )
	fmt.Println(NAME)
	fmt.Println(PI)
	fmt.Println(C1,C2)
	fmt.Println(C3,C4)

}

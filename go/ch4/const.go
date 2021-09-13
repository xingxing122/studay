package main

import "fmt"

func constDemo(){
	const s string  = "Hello"
	const  a,b  = 3,4
	fmt.Println(s,a,b)

}

func  enumDemo(){
	const (
		Sunday  =  iota
		Monday
		Tuesday
		wednesday


	)
	fmt.Println(Sunday,Monday,Tuesday,wednesday)
}

func main() {
  constDemo()
  enumDemo()
}

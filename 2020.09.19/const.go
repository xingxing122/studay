package  main

import "fmt"

const  statusNew  int = 1
const  statusDeleted int = 2


func  main () {
	const (
		Monday = iota * 10
		Tuesday
		wedesday
	)

	fmt.Println(statusNew, statusDeleted)
	fmt.Println(Monday, Tuesday, wedesday)

}


package main

import "fmt"

func  test ()  (i int) {

	defer func(){
		fmt.Println("defer")
		i = 2
	}()
	return i
}

func  main(){
	fmt.Println(test())

}


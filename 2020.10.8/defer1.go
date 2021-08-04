package  main

import "fmt"


func main(){

    for i :=0;i < 2;i++ {
		defer func(i int ) {
			fmt.Println("i")
		}(i)
	}

	fmt.Println("end")
}

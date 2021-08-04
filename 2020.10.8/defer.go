package  main

import "fmt"


func main(){
	fmt.Println("start")

 	defer func(){
		fmt.Println("defer1")
	}()
	defer func(){
		fmt.Println("defer2")
	}()
	fmt.Println("end")
}

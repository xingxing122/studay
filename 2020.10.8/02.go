package main

import "fmt"

func sayHello(){
	fmt.Println("hello")
}


func sayHi(name string){
	fmt.Println("hi",name)
}

func add(a,b int) int{
	return  a + b
}

func test(a int,b string, c int){

}
func main(){
	fmt.Println(add(3,4))


}

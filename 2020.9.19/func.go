package main

import "fmt"

func   sayHello(){
	fmt.Println("hello")

}

func add (a,b int)  int{
	return  a + b
}
func  op( a,b int ) (int, int, int, int){
	return  a+ b , a-b, a*b, a/b
}


func main() {
       sayHello()
       add(1,2)
       c := add(3,4)
       fmt.Println(c)
       fmt.Println(op(4,2))

}


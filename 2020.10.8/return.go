package main

import "fmt"

func   sayHello(){
	 fmt.Println("hello")
}

func     add(a,b int) int {
    return   a+ b
}

func op(a,b int)(int,int,int,int){
	return  a + b , a - b , a * b  , a / b
}
func  opv2(a,b int) (sum int, sub int, mul int,div int ){
	sum =  a +b
	sub  = a -b
	div  = a/b
	mul  = a * b
	return

}


func main() {
	sayHello()
	add(1,2 )
	c := add(2,3)
	fmt.Println(c)
	fmt.Println(op(4,2))
	fmt.Println(opv2(4,2))


}

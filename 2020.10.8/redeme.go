package main

import "fmt"

//无参数
func sayHello(){
	fmt.Println("Hello")
}

//有参数
func sayHi(name string){
	fmt.Println("hi,",name)
}

//有返回值的

func add(a int,b int) int{
	 return  a + b
}
func main() {
     sayHello()
     sayHi("xingxing")
     c := add(2,3)
     fmt.Println(c)

}
package   main

import (
   "fmt"
)

func    sayHello(){
	fmt.Println("hello")

}

//有参数，无返回值
func sayHi(name string ){
	fmt.Println("hi",name)
}
//有返回值的  add
func  add(a int,b int) int{
	fmt.Println(a,b)
	return  a + b
}

func main(){
	sayHello()
	sayHello()
	sayHi("学习go编程")
	c := add(3,4)
    fmt.Println(c)


}
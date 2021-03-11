package main

import "fmt"

//无参函数
func sayHello(){
	fmt.Println("Hello")

}

//有参数，无返回值
func sayHi(name string){
    fmt.Println("hi,",name)
}

//有返回值  add
func  add(a int,b int) int{
	fmt.Println(a,b)
	return  a + b
}

func main(){
	sayHello()
	sayHi("中国")
	name :="北京"
	sayHi(name)

	c := add(1,2)
	fmt.Println(c)


}

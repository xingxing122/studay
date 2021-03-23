package  main

import "fmt"

// 无参函数
func    sayHello(){
	fmt.Println("Hello")

}

//有参数，无返回
func   sayHi(name string) {
	fmt.Println("hi,",name)
}

//有返回值 add

func  add(a int,b int) int{
	return  a + b
}

func main(){

	sayHello()
    sayHi("kk")
	name :="中国人"
	sayHi(name)
	c := add(1,2)
	fmt.Println(c)


}



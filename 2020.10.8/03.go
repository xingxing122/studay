package  main

import "fmt"

func sayHello(){
	fmt.Println("hello")
}

func add(a,b int) int{
	return  a + b
}

func op (a,b int)(int,int,int,int){
	return  a +b ,a -b, a* b, a/b
}

//命名返回值
func opv2(a,b int)(sum int,sub int,mul int,div int){
	sum = a + b
	sub = a - b
	mul = a * b
	div = a / b
	return
}

func main(){
	sayHello()
	add(1,2)
	c := add(3,4)
	fmt.Println(c)
	fmt.Println(op(4,2))
	a,b,c,d := op(4,2)
	fmt.Println("a=",a,"b=",b,"c=",c,"d=",d)
	fmt.Println(opv2(3,2))

}
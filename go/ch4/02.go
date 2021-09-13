package main

import "fmt"


func operatorDemo(){
	var a int = 21
	var b int = 10
	var c  int

	c = a + b
	fmt.Printf("a+b=%d\n",c)

	c = a - b
	fmt.Printf("a-b=%d\n",c)

	a++
	fmt.Println(a)

	a = 21
	a--
	fmt.Println(a)


}



func logicDemo(){
	var  a  bool = true
	var  b  bool  = false

	if  a && b{
		fmt.Println("a和b都是真")
	}else{
		fmt.Println("a和b有假")
	}

	if a || b {
		fmt.Println("a或者b至少有一个真")
	}else {
		fmt.Println("a和b都有假")
	}


}


func main() {
	operatorDemo()
	logicDemo()
}
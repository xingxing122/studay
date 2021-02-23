package main

import "fmt"

func add(a,b int)int {
	 return a + b
}

func mul(a,b int)int{
	return  a * b
}

func main() {
	var  f func(int,int) int
	var fs [] func(int,int)int
	fs = append(fs,add,mul)

	f = add
	c := f(2,3)

	fmt.Printf("%T\n",c)
	for _,f := range fs {
		fmt.Println(f(2,3))
	}
	
}

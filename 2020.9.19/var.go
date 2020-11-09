package main

import "fmt"

var name string

func main() {
	var me string
	me = "123"
	name = "xingxing"
	fmt.Printf("%s",me)
	fmt.Println(name,"\n")

	var   name,user  string  = "abc","我爱你祖国"

	fmt.Println(name,"\n",user,"\n")

	var   (
		age   int   = 31
		height   float64   = 1.68
	)

	fmt.Println(age,height)


	isBoay  := false
	fmt.Println(isBoay)



}





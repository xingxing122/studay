package main

import "fmt"

func  varInitValue(){
	var  a,b int = 3,4
	fmt.Println(a,b)
	var s string = "abcd"
	fmt.Println(s)

	var s1,s2,s3 = 5, "hello", true
	fmt.Println(s1,s2,s3)
}




func main() {
   var a int
   var b string
   fmt.Printf("%d %a\n",a,b)
   varInitValue()

}

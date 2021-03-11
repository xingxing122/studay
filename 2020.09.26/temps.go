package main

import "fmt"

func main(){
	a :=[]byte{'a','b','c','d','e','f','g','h','i','j','k'}

	 sa :=a[2:5]
	 sb := sa[1:3]
     fmt.Println(string(sb))



}


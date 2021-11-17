package main

import  "fmt"

// 现有的数据类型
type  Counter int

func main() {
	var counter Counter
	fmt.Printf("%T,%v\n", counter,counter)
	counter += 10
	fmt.Printf("%T,%v\n",counter,counter)
	var num int = 10
	c2 := Counter(num)+counter

	fmt.Printf("%T, %v\n",c2,c2)

   // var  user  User
	// fmt.Printf("%T,%v\n",user,user)

}

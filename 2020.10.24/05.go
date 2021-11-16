package main

import "fmt"

// 自定义类型

type Counter int

type  User map[string]string

func  main(){
	  var counter  Counter
      fmt.Printf("%T,%v\n",counter,counter)

	  counter  += 10
	  fmt.Printf("%T, %v\n",counter,counter)
	  var  num int = 10
	  c2 := Counter(num) + counter
	  fmt.Printf("%T,%v\n",c2,c2)


	  var use User = make(User)

	  fmt.Printf("%T,%v\n",use,use)


}
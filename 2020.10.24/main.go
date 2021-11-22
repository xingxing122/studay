package main

import  "fmt"

// 现有的数据类型
type  Counter int
type  User map[string]string

type  Callbacks   func() error


type  Counter2   Counter

func main() {
	var counter Counter
	fmt.Printf("%T,%v\n", counter,counter)
	counter += 10
	fmt.Printf("%T,%v\n",counter,counter)
	var num int = 10
	c2 := Counter(num)+counter

	fmt.Printf("%T, %v\n",c2,c2)

    var  user  User = make(User)
	fmt.Printf("%T,%v\n",user,user)
	user["id"] = "1"
	fmt.Println(user)

	callbacks  := map[string]Callbacks{}

	callbacks["add"] = func() error {
		fmt.Println("add")
		return nil
	}
	callbacks["add"]()


	var   v22 Counter2
	fmt.Printf("%T %v\n",v22,v22)


}

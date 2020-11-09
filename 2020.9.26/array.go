package main

import (
	"fmt"
)
func main() {
	var  num [2]int = [...]int{100,88}

	var scores = [...]int{100,88}

	fmt.Printf("%T\n",num)
	fmt.Println(scores)

	fmt.Println(num==scores)


	fmt.Println(num[0])
	fmt.Println(num[1])

	num[0]= 101
	fmt.Println(num)

	fmt.Println(len(num))

	// 遍历
	for  i := 0; i<len(num); i++{
		fmt.Println(i,num[i])
	}

	for  v := range  num{
		fmt.Println(v)
	}


}

package main

import "fmt"

func main(){
	name := "kk"
	nums := []int{}
	func(){
		fmt.Println(name,nums)
		name ="silence"
		nums =[]int{1,2,3}
		fmt.Println(name,nums)
	}()
      fmt.Println(name,nums)

}

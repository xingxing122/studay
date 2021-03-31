package main

import "fmt"

func main(){
	name := "kk"
	nums := [] int{1,2,3}

	func (name string, nums [] int ){
		fmt.Println(name,nums)
		name = "silence"
		nums[0] = 100
		fmt.Println(name,nums)
	}(name,nums)

	fmt.Println(name,nums)


}




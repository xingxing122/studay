package main

import "fmt"

func main() {

	nums :=[5]int{1,2,3,4,5}

	nums2 := nums
	fmt.Println(nums,nums2)

	nums2[0] = 100
	fmt.Println(nums,nums2)

     var   age = 1
     var  agePointer *int  = nil
     agePointer = &age

     fmt.Println(agePointer,age)

     fmt.Println(*agePointer)
}

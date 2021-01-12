package main

import "fmt"

func main() {
	nums := []int{1,2,3}
	nums2 :=[]int{3,4,5}

	nums = append(nums,100)

	fmt.Println(nums,nums2)

	



}

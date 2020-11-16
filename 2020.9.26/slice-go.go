package main

import "fmt"

func main() {
	//nums := []int{1,2,3,4,5}
	nums := make([]int,5,100)
	nums2 := nums[2:5]

	fmt.Println(nums2,len(nums2),cap(nums2))



}

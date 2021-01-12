package main

import "fmt"

func main() {
	nums := make([]int,3)
	fmt.Println(len(nums),cap(nums))

	//len 与cap
	nums2 := make([]int,2,5)
	fmt.Println(len(nums2),cap(nums2))

	nums3 := nums2
	nums3  = append(nums3,3)
	fmt.Println(nums2,nums3)
	nums2 = append(nums2,4)
	fmt.Println(nums3)
	fmt.Println(nums2)






}

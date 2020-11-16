package main

import "fmt"

func main(){
	//make
	// 2个参数: make(type,len)
	// 3个参数: make(type,len,cap)

	nums := make([]int,3)
	fmt.Println(len(nums),cap(nums))
	fmt.Println(nums)

	nums2 := make([]int,2,5)
	fmt.Println(len(nums2),cap(nums2))
	fmt.Println(nums2)

	nums3 := nums2
	nums3 = append(nums3,3)
	nums2 = append(nums2,4)
	fmt.Println(nums3)
	fmt.Println(nums2)


}



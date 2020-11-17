package main

import "fmt"

func main() {
	//nums := []int{1,2,3,4,5}
	nums := make([]int,5,100)
	// slice[start:end]
	// 0 <= start  <=  end  <= cap
	// len := end-start
	// cap := cap - start
	nums2 := nums[10:100]
	fmt.Println(nums2)
    fmt.Println(nums2,len(nums2),cap(nums2))

}

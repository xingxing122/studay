package main

import "fmt"

func main() {
    // 复制nums 中所有数据到nums2
	 //第一种方法
	//nums := []int{1,2,3,4,5}
	 //nums2 := []{}
	 //copy(nums2,nums)

	 /* 第二种方法
	nums  := []int{1,2,3,4,5}
	 nums2 := make([]int,len(nums))
	 for i, v := range nums {
	 	nums2[i] =v
	 }
     fmt.Println(nums2)*/

	nums := []int{1,2,3,4,5}
	nums2 := make ([]int,0,len(nums))
	for _,v := range nums {
	nums2 = append(nums2,v)
	}
	fmt.Println(nums2)

}

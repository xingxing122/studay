package main

import "fmt"

func main() {
	nums :=[]int{1,2,3}
	nums2 :=[]int{3,4,5}


	nums = append(nums,100)
	for _, v:=range nums2 {
		nums = append(nums,v)
	}
      fmt.Println(nums,nums2)

}

package main

import "fmt"

func main() {
	// 复制nums 中的所有数据到num2 中
	// v1
	//nums :=[]int{1,2,3,4,5}
	//nums2 :=[]int{}
	//copy(nums2,nums)
	//fmt.Println(nums2)
	//
	// v2
	nums :=[]int{1,2,3,4,5}
	nums2 := make([]int,len(nums))
	for i, v := range nums {
		nums2[i] = v
	}
	fmt.Println(nums2)
	//
	//v3
    //nums :=[]int{1,2,3,4,5}
    //nums2 :=make([]int,0,len(nums))
    //空白标识符
    //for _,v:=range nums{
    //	nums2 = append(nums2,v)
	//}
    // fmt.Println(nums2)
}

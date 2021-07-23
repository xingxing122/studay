package main

import "fmt"

func main() {
    nums := []int{1,2,3,4,5}
    fmt.Println(nums[0:3])
    fmt.Println(nums[:3])
    fmt.Println(nums[3:len(nums)])

    nums2 := []int{3,4,5}

    /*for _,v := range nums2 {
        nums = append(nums,v)
    }
     fmt.Println(nums,nums2)
      */

    nums = append(nums,nums2...)
    fmt.Println(nums,nums2)

}

package main

import "fmt"

func bubble(nums []int){
      for j := 0;j<len(nums)-1;j++{
            for i :=0 ; i<len(nums)-1-j;i++{
                  if nums[i] > nums[i+1] {
                        nums[i],nums[i+1]=nums[i+1],nums[i]
                  }
            }
      }
      fmt.Println(nums)

}

func main() {
      nums :=[]int{5,4,3,2}
      fmt.Println(nums)
      nums =[]int{100,90,88,70,1000,100001}
      bubble(nums)

}

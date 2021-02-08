package main

import "fmt"

func main() {

      nums  :=[]int{1,4,3,2}

      for  i := 0; i < len(nums)-1; i++{
            if nums[i] > nums[i+1]{
                  nums[i], nums[i+1]=nums[i+1],nums[i]
            }

      }
      fmt.Println(nums)

}

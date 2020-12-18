package main

import (
	"fmt"
)

func main(){
	 nums := []int{1,3,5,30,43,9,54}
	 maxNum := nums[0]

	 for i,v := range nums{
	  if v > maxNum {
	  	maxNum = v
	  }
	 	fmt.Println(i,":",maxNum)
	 }
	  fmt.Println(maxNum)
}
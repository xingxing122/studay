package main

import "fmt"

func main() {
      var  scores = [...]int{100,88}
      var  nums [2]int =[2]int{100,88}
      fmt.Println(nums == scores )
      fmt.Println(nums[0])
      fmt.Println(nums[1])

      nums[0] = 101
      nums[1]  = 102
      fmt.Println(nums[0])
      fmt.Println(nums[1])

      //遍历
      for i :=0; i < len(nums); i++ {
      	fmt.Println(i,nums[i])
	  }

	  for   v := range nums {
	  	fmt.Println(v,nums[v])
	  }

	    for  i, v := range nums{
	    	fmt.Println(i,v)
		}
}




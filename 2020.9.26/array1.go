package  main

import "fmt"

func main(){

	var names[10] string
	var scores = [...]int{100,88}  //数组赋值，字面量
	fmt.Println(names,scores)

	//运算
	var  nums [2]int = [...]int{88,100}
	fmt.Println(nums == scores)
	//访问
	fmt.Println(nums[0])

	nums[0] = 101
	nums[1] = 102
	fmt.Println(nums)

	 for i := 0; i<len(nums);i++{
	 	fmt.Println(i,nums[i])
	 }
	  for i ,v := range nums {
	  	fmt.Println(i,v)
	  }







}



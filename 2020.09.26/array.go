package  main

import "fmt"

func   main(){
	var names [55]string

	var  scores = [...]int{10000,90}
	fmt.Printf("%T\n",names )
	fmt.Printf("%T\n",scores)
	fmt.Println(names,scores)

	var  nums[2]int = [...]int{10000,90}

	fmt.Println(scores == nums)

	fmt.Println(nums[0])
	fmt.Println(nums[1])
	nums[0]=101
	fmt.Println(nums)
	nums[1]=103
	fmt.Println(nums)

	fmt.Println(len(nums))
	fmt.Println(cap(nums))

	for i:=0;i< len(nums);i++{
		fmt.Println(i,nums[i])

	}
	for  v := range nums {
		fmt.Println(v,nums[v])
	}






}
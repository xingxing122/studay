package main

import "fmt"

func  main(){
	//定义变量names 为元素类型
	var  names [55] string
	var  scores = [...]int{3:100,1:88,0:99}

	fmt.Printf("%T\n",names)
	fmt.Printf("%T\n",scores)

    fmt.Println(names,scores)


	var  nums [2] int =[...] int {100,88}
	fmt.Printf("%T\n",names)
	fmt.Printf("%q\n",names)
	fmt.Printf("%T\n",scores)

	fmt.Println(nums)
	nums[1] = 10000
	fmt.Println(nums[0])
	fmt.Println(nums[1])
	fmt.Println(len(nums))
	fmt.Println(cap(nums))


	for i := 0; i < len(nums);i++ {
		fmt.Println(i,nums[i])
	}

	for v:= range nums {
		fmt.Println(nums[v])
	}

    for  i, v := range  nums{
    	fmt.Println(i,v)
	}




}


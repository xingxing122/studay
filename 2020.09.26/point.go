package main

import "fmt"

func main() {
	//指针
	//值类型
	//赋值  原有的数据复制一份给新的变量
	//两个变量之间没有任何关系
	//对nums 进行修改都不会影响nums2
	//对nums2 进行任何修改也不会影响nums
	nums := [5]int{1,2,3,4,5}
	nums2 := nums
	fmt.Println(nums,nums2)

	nums2[0] = 100
	fmt.Println(nums,nums2)

	var  age = 1
	var  agePointer *int
	fmt.Println(age,agePointer)


	agePointer = &age
	fmt.Println(agePointer,age)

	fmt.Println(*agePointer)

	*agePointer = 33
	fmt.Println(age,*agePointer)

	var numsPoint   = &nums
	fmt.Printf("%T\n",numsPoint)

	numsPoint[0] = 100
	fmt.Println(nums,numsPoint)









}

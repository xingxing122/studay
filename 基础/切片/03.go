package main

import "fmt"

func main() {

	//定义切片

	nums := []int{1, 2, 3, 4}
	fmt.Println(nums)

	fmt.Printf("长度： %v,容量： %v\n", len(nums), cap(nums))
	nums = append(nums, 5)
	fmt.Println(nums)

	fmt.Printf("长度： %v,容量： %v\n", len(nums), cap(nums))

	//添加多个元素
	nums = append(nums, 11, 13, 14, 14)

	fmt.Println(nums)
	// 3.切片合并
	num1 := []int{1, 2, 3}
	num2 := []int{3, 4, 5}

	num1 = append(num1, num2...)

	fmt.Println(num1)

	//删除切片中的第二个元素
	num3 := []int{1, 2, 3, 4, 5}

	num3 = append(num3[2:], num3[3:]...)
	fmt.Println(num3)

}

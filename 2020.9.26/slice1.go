package main

import "fmt"
func main() {
	var   nums []int
	fmt.Printf("%T\n",nums)

	fmt.Println(nums==nil)

	// 赋值
	nums =[]int{1,2,3,4}
	fmt.Printf("%#v\n",nums)


	// 数组切片赋值
	var arrays  [10]int = [10]int{1,2,3,4,5,6,7,8,9}
	nums  = arrays[1:10]
	fmt.Printf("%#v\n",nums)
	fmt.Printf("%#v %d  %d\n",nums,len(nums),cap(nums))

    // make 函数
    nums = make ([]int,3)
    fmt.Printf("%#v %d %d\n",nums,len(nums),cap(nums))

    nums = make([]int,3,5)
	fmt.Printf("%#v %d %d\n",nums,len(nums),cap(nums))


    // 元素的  增   删 改  查
    fmt.Println(nums[0])
    fmt.Println(nums[1])
    fmt.Println(nums[2])
    nums[2]=10
    fmt.Println(nums)
    nums= append(nums,1)
	nums= append(nums,1)
	nums= append(nums,1)
	nums= append(nums,1)
	nums= append(nums,1)


	//遍历
	for i := 0; i<len(nums);i++{
		fmt.Println(nums)
	}

	// 切片的操作
	fmt.Printf( "%T\n",nums[1:3])

	n := nums[1:3:10]
    fmt.Printf("%T  %#v   %d  %d\n",n,n,len(n),cap(n))

	nums02 :=nums[1:3]
	fmt.Println(nums,nums02)

	nums02 = append(nums02,3)
	nums = append(nums,4)
	fmt.Println(nums,nums02)


	nums = arrays[:]
	fmt.Println(nums,arrays)


	nums[0] = 100
	fmt.Println(nums,arrays)


	//删除
	//copy
	nums04 := []int{1,2,3}
	nums05 := []int{10,20,30,40}

    copy(nums05,nums04)
	fmt.Println(nums05)

	nums05 = []int{10,20,30,40}
	copy(nums04,nums05)
	fmt.Println(nums04)

	//删除前一位和后一位
	nums06 :=[]int{1,2,3,4,5,}
	fmt.Println(nums06[1:])
	fmt.Println(nums06[:len(nums06)-1])

	//删除中间的数字
	copy(nums06[2:],nums06[3:])
	fmt.Println(nums06[:len(nums06)-1])

	//堆栈
	//队列
	queue :=[]int{}
	queue = append(queue,1)
	queue = append(queue,2)
	queue = append(queue,3)
	queue = append(queue,4)
	queue = append(queue,5)
	fmt.Println(queue[0])
	queue = queue[1:]
	fmt.Println(queue)
	fmt.Println(queue[0])
	queue=queue[1:]
	fmt.Println(queue)



}














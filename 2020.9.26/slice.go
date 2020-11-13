package main

import (
	"fmt"
)

func main(){

	var  names []string

	fmt.Printf("%T\n",names)
    fmt.Printf("%v\n",names)

	// 初始化值
	names = []string{"xingxing","123"}

	fmt.Printf("%T\n",names)
	fmt.Printf("%q\n",names)

	//字面量
	//[]type{}=> 空切片
	//[]type{v1,v2,v3......}
	//[]type{i1:v1,i2:v2....}
	names =[]string{1:"golang", 10:"123"}
	fmt.Printf("%T\n",names)
	fmt.Printf("%q\n",names)

	//访问  修改元素
	//索引
	fmt.Println(names[1])
	fmt.Println(names[10])
	fmt.Println(names[8])
	names[8]="学习go" //修改
	fmt.Println(names)


	fmt.Println(len(names))
	//长度，切片中已经存在元素的数量

	for i := 0;i<len(names);i++{
		fmt.Println(i,names[i])
	}

	for  i, v := range names{
		fmt.Println(i,v)
	}

	//添加元素
	names = append(names,"测试添加")
	fmt.Println(names)

	names = append(names,"继续添加")
	fmt.Println(names)


	//删除元素
	//切片操作
	//names[start:end] names 中从start开始到end-1 所有元素组成的切片
	fmt.Printf("%q\n",names[1:10])

	//删除索引为0，如果索引为len-1元素
	names = names[1:len(names)]
	fmt.Printf("%q\n",names)

	names = names[0:len(names)-1]
	fmt.Printf("%q\n",names)

     //删除中间的元素
     nums :=[]int{0,1,2,3,4,5,6,7,8,9}
     //删除3
     nums2 :=[]int{10,11,12,13,14,15,16,17}
     copy(nums,nums2)
     fmt.Println(nums,nums2)
     //nums[0:3],nums[4:5]








}
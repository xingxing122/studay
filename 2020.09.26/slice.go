package main

import (
	"fmt"
)

func main() {
     var  names []string

     fmt.Printf("%T\n",names)
     fmt.Printf("%v\n",names)

     names = []string{"我爱编程","abc"}
     fmt.Printf("%T\n",names)
     fmt.Printf("%q\n",names)

     //字面量
     names =[]string{1:"java",10:"go"}
     fmt.Printf("%T\n",names)
     fmt.Printf("%q\n",names)

     //访问元素
     fmt.Println(names[1])
     fmt.Println(names[3])
     fmt.Println(names[10])

     names[9]="学习go"
     fmt.Println(names)

     fmt.Println(len(names))

     for i:=0; i< len(names); i++{
     	fmt.Println(i,names[i])
	 }

     for v := range names{
     	fmt.Println(v,names[v])
	 }

     for i,v :=range  names {
     	fmt.Println(i,v)

     //添加元素
     names= append(names,"地理")
     	fmt.Println(names)
	 }

	 fmt.Println(len(names))
     fmt.Printf("%q\n",names[1:22])

     names = names[1:len(names)]
     fmt.Printf("%q\n",names)

     names = names[0:len(names)-1]
     fmt.Printf("%q\n",names)

     nums :=[]int{0,1,2,3,4,5}
     //删除3
     // copy(dst,src)
     nums2 :=[]int{10,11,12,13,14,15,16}
     copy(nums,nums2)
     fmt.Println(nums,nums2)

     copy (nums[3:len(nums)],nums[4:len(nums)])
    fmt.Println(nums)





}

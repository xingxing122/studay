package main

import "fmt"

func main() {

	 var  names []string
	 fmt.Printf("%T\n", names)

	 fmt.Println(names)

	 names = []string{"数字第一组","sen"}

	 fmt.Println(names)
	  names = []string{1:"数字第一组",10:"sen"}

	  fmt.Println(len(names))
	  fmt.Println(names[10])
	  names[9]="中国我的梦"

	  fmt.Println(names)

	  fmt.Println(len(names))

	  for i,v := range  names {
	  	fmt.Println(i,v)
	  }

	  //添加元素
     names=append(names,"paas")
     fmt.Println(names)
     fmt.Println(len(names))

     //删除元素
     fmt.Printf("%q\n",names[1:10] )
     names = names [1:len(names)]
     fmt.Printf("%q\n",names)

     names = names[0:len(names)-1]
     fmt.Printf("%q\n",names)

     // 删除中间元素
     nums :=[]int{0,1,2,3,4,5}

     nums2 :=[]int{10,11,12,13,14,15,16,17}

     copy(nums,nums2)
     fmt.Println(nums,nums2)




}

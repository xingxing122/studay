package main

import (
	"fmt"
)

func main() {
      var   names  [] string
      fmt.Printf("%T\n", names)
      fmt.Printf("%v\n",names)

      names  = []string {"go","paas"}
	  fmt.Printf("%T\n", names)
	  fmt.Printf("%v\n",names)
      // 字面量 是给某种数据类型初始化的
      // [] type{} ==》 空切片
      names =[] string{1:"k8s", 10: "sen"}
      fmt.Printf("%T\n", names)
      fmt.Printf("%q\n", names)

      names[9]="新奥paas"
      // 访问修改元素
      //索引
      fmt.Println(names[1])
      fmt.Println(names[10])
      fmt.Println(names[9])

      fmt.Println(len(names))
      fmt.Println(cap(names))

      for   i := 0; i < len(names); i++ {
      	fmt.Println(i,names[i])
	  }

	  for v := range  names {
	  	fmt.Println(v,names[v])
	  }

	  for i, v := range  names{
	  	fmt.Println(i,v)
	  }

      // 添加
	  names = append(names,"devops")
	  fmt.Println(names)

	  fmt.Println(len(names))
	  // 删除元素
	  // names[start:end]
	  // names [1:10]

	  fmt.Printf("%q\n",names[1:10])

	  nums  := []int{0,1,2,3,4,5}
	  //copy(dst,src)
	  nums2 :=[]int{10,11,12,13,14,15}
	  copy(nums,nums2)
	  fmt.Println(nums,nums2)











}

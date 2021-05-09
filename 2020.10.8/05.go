package main

import (
	"fmt"
)
func main(){

	var name[55] string
	// 定义变量name 为元素类型为字符串长度为55的数组

	var score  = [...]int{100,88}
	// 字面量


	var score1  = [...]int{100,88}
	fmt.Println(name)
	fmt.Printf("%T\n",name)
	fmt.Printf("%q\n",name)


	fmt.Printf("%T\n",score)
	fmt.Println(score)


	fmt.Println(score1==score)

	fmt.Println(score[0])
	fmt.Println(score[1])

	score[0]=101
	fmt.Println(score[0])

	 fmt.Println(len(score))


	 for i:=0;i<len(score);i++{
	 	fmt.Println(i,score[i])
	 }

	 for v := range  score {
	 	fmt.Println(v,score[v])
	 }

      for i, v := range score {
      	fmt.Println(i,v)
	  }



}

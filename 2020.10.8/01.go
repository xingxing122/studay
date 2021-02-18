package main

import (
	"fmt"
)

func  sayHi(name string){
	fmt.Println("hi",name)

}

func add(a int, b int) int {
	return a + b //】 // 参数中多个连续的参数类型相同

}

func  addV2(a,b int) int{
	return  a + b
}

func addAll(a,b int,args ... int)int{
	fmt.Println(a,b,args)
	fmt.Printf("%T\n",args)
	sum := a + b
	for _,v := range  args {
		sum +=v
	}
	return sum
}

func print (args ... int){
	for i,v := range args {
		fmt.Println(i,v)
	}
}


func main(){
	sayHi("KK")
	name := "我爱我家"
	sayHi(name)
	c := add(2,3)
	fmt.Println(c)
	fmt.Println(addV2(3,4))
	fmt.Println(addAll(1,2))
	fmt.Println(addAll(1,2,3,4))

	nums :=[]int{1,2,3,4,5}
	fmt.Println(addAll(1,2,nums...))
	print(nums...)


}


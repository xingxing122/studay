package  main

import "fmt"

func   sayHello(){
	fmt.Println("hello")
}

func sayHi(name string){
	fmt.Println("hi",name)
}

func  add(a int,b int) int{
	return  a + b
}
// 参数类型合并
func  addV2(a,b int) int{
	return  a + b
}
//可变参数,只能定义一个可变参数，并且在形参的最后
func addAll(a,b int,args...int)int{
	fmt.Println(a,b,args)
	fmt.Printf("%T\n",args)
	print(args...)
	sum := a + b
	for _, v := range args {
		sum += v
	}
	return sum

}

func  print(args ...int){
	for i , v := range  args {
		fmt.Println(i,v)
	}
}

func main(){
	fmt.Println(addV2(3,4))
	fmt.Println(addAll(1,2))
	fmt.Println(addAll(1,2))
	fmt.Println(addAll(1,2,3,4))
	fmt.Println(addAll(1,2,3,4,5,6,7))

	nums := []int{2,3,4,5}
	fmt.Println(addAll(1,2,nums...))
	print(nums...)


}

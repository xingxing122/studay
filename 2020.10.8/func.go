package main

import "fmt"



func sayHello(){
	fmt.Println("hello")
}

func sayHi(name string){
	fmt.Println("Hi,",name)
}

func   addAll(a,b int,args...int)int{
       fmt.Println(a,b,args)
       fmt.Printf("T\n",args)
       sum := a+b
       for _,v := range  args {
       	sum +=v
	   }
       return 0
}

func   testV2(a,b int) int {
	return  a + b
}

func  main(){
	sayHello()
	sayHi("xingxing")
	addAll(1,2)

	c := testV2(5,9)
	fmt.Println(c)
	d := addAll(1,2,3,4,5,6,7)
	fmt.Println(d)
	nums :=[]int{1,2,3,4,5}
	fmt.Println(addAll(1,2,nums...))


}
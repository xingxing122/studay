package main

import "fmt"

func  main(){
	//队列
	//先进先出

     queue :=[]int{}
     queue = append(queue,1)
     queue = append(queue,3)
     queue = append(queue,2)
     queue = append(queue,6)

     for len(queue)  != 0 {
     	fmt.Println(queue)
     	queue = queue[1:]
	 }

	 fmt.Println("over")
}

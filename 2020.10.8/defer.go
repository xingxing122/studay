package main

import "fmt"

func main() {
     fmt.Println("start")
     //defer 函数调用
     //defer 延迟执行
	defer func(){
		fmt.Println("defer1")
	}()

     defer func(){
     	fmt.Println("defer2")
	 }()

    fmt.Println("end")
}



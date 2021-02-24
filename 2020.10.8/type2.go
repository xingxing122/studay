package main

import "fmt"


func  main(){
c := func(){
	fmt.Println("我是匿名函数")
}

fmt.Printf("%T\n",c)
c()

func(){
	fmt.Println("啦啦啦啦啦")
}()

}

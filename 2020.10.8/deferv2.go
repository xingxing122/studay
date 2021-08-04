package main

import "fmt"

func test(){
	defer  func (){
		if err := recover(); err != nil{
			fmt.Println("错误了，进行恢复",err)
		}
	}()
	panic("我是一个错误")


}

func main() {
	 test()
}

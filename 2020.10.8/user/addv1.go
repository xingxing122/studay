package main

import (
	"fmt"
	"strconv"
)

users :=[]User{}

func id()string{


}

type User struct{
	ID,Name,Addr,Tel string
}


func main(){
	var (
		name, addr, tel string
	)

	fmt.Println("请输入用户名称:")
	fmt.Scan(&name)

	fmt.Println("请输入地址:")
	fmt.Scan(&addr)

	fmt.Println("请输入联系方式")
	fmt.Scan(&tel)

}






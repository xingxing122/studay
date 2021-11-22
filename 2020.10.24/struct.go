package main

import (
	"fmt"
	"time"
)

type  User  struct{
	 id  int
	 name string
	 addr string
	 tel  string
	 birthday time.Time

}

func main() {

	var  user User
	fmt.Printf("%T,%#v\n",user,user)

	//字面量
	// 必须严格按照结构体属性定义顺序指定
	//每个属性都必须指定对应的值
	user = User{10,"kk", "西安","xxxx",time.Now()}
	fmt.Printf("%#v\n",user)

	//按照属性定义的字面量
	user = User{id:10,name:"kk"}
	fmt.Printf("%#v\n",user)

	//属性的访问与修改
	fmt.Println(user.id)
	user.id= 100
	user.name="xingxing"
	fmt.Printf("%#v\n",user)

    u3 := NewUser(1,"kk","","",time.Now())
	fmt.Printf("%T, %#v\n",u3,u3)

}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

//定义存储用户的切片
//每个元素为一个用户
//var users []map[string]string{} =[]map[string]string{}


var users =[]map[string]string{}


func main(){
	var (
         name,addr,tel string
	)
	for{
		fmt.Print("请输入一个名称")
		fmt.Scan(&name)

		fmt.Print("请输入地址")
		fmt.Scan(&addr)

		fmt.Print("请输入联系方式")
		fmt.Scan(&tel)
		fmt.Println(name,addr,tel)
		user := map[string]string{
			 "name":name,
			 "addr":addr,
			 "tel":tel,
			 "id":generateID(),
		}
		users= append(users,user)
		usersBytes,_:=json.MarshalIndent(users,"  ","    ")
		fmt.Println(fmt.Sprintf("%s",usersBytes))
	}
	fmt.Println(users)
}

func generateID() string{
	for _,u := range users{
		if u["id"]!=""{
			id,err:=strconv.Atoi(u["id"]);
			if err!=nil{
				log.Infof("id format error: %s",err)
			}
			id++
			return fmt.Sprintf("%d",id)
		}
	}
	return "1"
}

func add(user map[string]string){
	users= append(users,user)
}
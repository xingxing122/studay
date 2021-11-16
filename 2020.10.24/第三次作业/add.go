package  main

import (
	"fmt"
	"strconv"
)

//var    users []map[string]string = []map[string]string{ }
var users = []map[string]string{}

// 获取增加用户的ID 值
func getID() string {

	id :=0
	for _, user := range users {

		if uid, err := strconv.Atoi(user["id"]); err == nil {
			if uid > id {
				id = uid
			}
		}
	}
	return   strconv.Itoa(id + 1)

}


func addUser(){
	var (
		name string
		addr string
		tel   string
	)
	fmt.Println("请输入名称: ")
	fmt.Scan(&name)
	fmt.Println("请输入联系地址:")
	fmt.Scan(&addr)
	fmt.Println("请输入联系方式")
	fmt.Scan(&tel)

	 user := map[string]string{
		 "id": getID(),
		 "name": name,
		 "addr": addr,
		 "tel": tel,
	 }
	 users = append(users,user)
}
func main(){
   addUser()
   fmt.Printf("%#v\n", users)

}
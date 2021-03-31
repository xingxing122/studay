package main

import (
	"fmt"
	"strconv"
)

// error 接口类型
func  div(left, right  int) (int, error) {
	 if   right ==0 {
	 	return 0 , fmt.Errorf("right num is zero")
	 }
	 return  left / right , nil

}




func  main(){
	v , err := strconv.Atoi("123abc")
	fmt.Println(v,err)
	if err == nil {
		fmt.Println("value是",v)
	} else {
		fmt.Println(err)
	}




	 fmt.Println(div(1,0))
	 fmt.Println(div(2,1))


}
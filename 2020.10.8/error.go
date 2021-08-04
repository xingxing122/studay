package  main

import (
	"fmt"
	"strconv"
)

// error 接口类型
func  div(left, right  int ) (int,error) {
	if right == 0 {
		return  0, fmt.Errorf("right num is zero")
	}
	return  left/ right, nil
}

func  divV2(left,right  int )(int, error ) {


}

func  main(){
      v,err := strconv.Atoi("123abc")
      if  err == nil {
      	fmt.Println("value是",v)
	  }else{
	  	fmt.Println(err)
	  }
	  //错误
	  //语法错误
	  //运行时错误
	  //
	  fmt.Println(div(1,0))
      fmt.Println(div(2,1))

}
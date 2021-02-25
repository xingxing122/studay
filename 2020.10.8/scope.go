package main

import "fmt"

func main() {
       name := "kk"
       nums := []int{}
       fmt.Println(name,nums)

       func(){
       	    name ="silence"
       	    nums = []int{1,2,3}
            fmt.Println(name,nums)
	   }()

       fmt.Println(name,nums)

}

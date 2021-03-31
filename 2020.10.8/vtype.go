package main

import "fmt"

func main(){
	name := "kk"
	nums := [] int{1,2,3}

	func (pname string, pnums [] int ){
		fmt.Println(pname,pnums)
		pname = "silence"
		nums = [] int{1,2}
		fmt.Println(pname,pnums)
	}(name,nums)

	fmt.Println(name,nums)


}
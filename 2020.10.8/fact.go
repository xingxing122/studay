package main

import "fmt"

func  factVn(n int) int{
	if n < 0 {
		return  -1
	}else if n == 0 {
		return  1
	}else {
		rt := 1
		for  i := 1; i<=n; i++{
			rt *= i
		}
		return  rt
	}

}


func fact(n int) int{
	if n  < 0 {
		return  - 1
	}
	if n ==0 {
		return  1
	}
	return n * fact(n-1)

}

func  main(){
	fmt.Println(fact(5))

}








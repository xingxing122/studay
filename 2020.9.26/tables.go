package main

import "fmt"

func main() {
	for i :=1; i<=9; i++{
		for b :=1; b<=i; b++{
			fmt.Printf("%d*%d=%2d\t",i,b,i*b)
		}
		fmt.Println()
	}

}

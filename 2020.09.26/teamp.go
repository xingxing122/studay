package main

import "fmt"

func main() {
	a := [...]int{99: 1}
	var p *[100]int = &a
	fmt.Println(p)

	v := [2]int{1, 2}
	c :=[2]int{2,3}
	fmt.Println(v==c)

	 s := new([10]int)
	 s[8] = 100
	 fmt.Println(s)

     fmt.Println("------------")
	 //多维数组
      f := [3][4]int{{2:1},{2,3}}
      fmt.Println(f)

}

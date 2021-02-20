package main

import "fmt"

func  tower(a,b,c string,layer int){
	if  layer ==1 {
		fmt.Println(a,"->",c)
		return
	}
	tower(a,c,b,layer-1)
	fmt.Println(a,"->",c)
	tower(b,a,c,layer-1)
}

func main() {
  tower("A","B","C",3)

}

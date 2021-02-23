package main

import "fmt"


func print(fm   func(string) string,   args ...string){
	for i, v := range args {
		fmt.Println(i,fm(v))
	}
}

func star(txt  string) string {
	return "*" + txt + "*"
}
func main() {
       names :=[]string{"赵长健","KK","17-赵"}
       print(star,names...)
}

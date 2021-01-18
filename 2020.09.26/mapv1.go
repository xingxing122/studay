package main

import "fmt"

func main() {
    m := map[int]string{1:"mike",2: "yoyo", 3:"go"}
    for key,  value := range  m {
    	fmt.Printf("%d=====> s%\n",key,value)

	}



}

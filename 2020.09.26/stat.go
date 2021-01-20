package main

import "fmt"

func main() {

    name := []string {
    	"小狗",
    	"二蛋",
    	"狗蛋",
    	"二蛋",
    	"狗蛋",
    	"狗蛋",
	}

	con := make(map[string]int)

	for _, stat := range name {
		v ,ok := con[stat]
		if !ok {
			con[stat] =1
		} else{
			con[stat] = v +1
		}
	}
fmt.Println(con)

}

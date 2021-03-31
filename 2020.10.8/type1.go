package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sayHi(){
	fmt.Println("hi")
}

func sayHello(){
	fmt.Println("Hello")
}

func   genFunc() func() {
	if rand.Int()%2  == 0 {
		return  sayHi
	}
	return  sayHello
}

func main(){
	rand.Seed(time.Now().Unix())
    f := genFunc()
    f()


}
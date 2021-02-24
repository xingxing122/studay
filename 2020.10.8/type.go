package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func   sayHi(){
	fmt.Println("hi")
}


func sayhello(){
	fmt.Println("Hello")
}

func  genFunc() func(){
	  if rand.Int()%2 == 0{
	  	return  sayHi
	  }
	  return  sayhello
}

func aFields(split rune) bool{
	if split == 'a' {
		return  true
	}
	return  false
}



func  main(){
      rand.Seed(time.Now().Unix())
      f := genFunc()
      f()
      fmt.Printf("%q\n",strings.FieldsFunc("abcdefabcabc",aFields))


}




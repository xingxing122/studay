package main

import "fmt"

func main() {
     var name string  = "\a"

     fmt.Printf("%T %s\n", name,name)


     name = "我爱中国"
     fmt.Println(name)
     fmt.Println(name[1])
     fmt.Println(len(name))
     fmt.Println(name[0:3])

}

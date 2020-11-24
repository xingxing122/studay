package main

import "fmt"

func main() {
    var  name   string = `xingxing`
    fmt.Printf("%T,%s\n",name,name)

    name += "hi"
    fmt.Println(name)
    name = "我爱中国"
    //索引
    fmt.Printf("%T\n",name[0])
    //长度
    fmt.Println(len(name))
    //切片
    fmt.Println(name[0:4])

}

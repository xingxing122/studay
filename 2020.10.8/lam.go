package main

import "fmt"

func   print(formatter  func(string) string, args... string){
	for i ,v := range args {
		fmt.Println(i,formatter(v))
	}
}


func   add()


func main() {
       names :=[]string{"我爱中国","kk","17-赵"}

      c := func(){
      	fmt.Println("我是匿名函数")
	  }
      fmt.Printf("%T\n",c)
      c()
      star := func(txt string) string{
      	return  "*" + txt + "*"
	  }
	  print(star,names...)



}

package main

import "fmt"

func  print(formatter func (string) string,args ... string ){
	for  i ,v := range args {
		fmt.Println(i,formatter(v))
	}
}
func star(txt string) string{
	return  "*" + txt + "*"
}

func line(txt string, end string)string{
	return  txt + end
}

func  main() {
	names := []string{"中国", "cc", "zz"}
	print(star,names...)
	//print(line,names...)
}
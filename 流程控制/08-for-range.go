package main

import "fmt"

func main() {

	/*	var str = "你好golang"
		for k, v := range str {
			fmt.Printf("key=%v val = %c\n", k, v)
		}*/

	var a = ".css"

	switch a {
	case ".html":
		fmt.Println("text/html")
	case ".css":
		fmt.Println("text/css")
	case ".js":
		fmt.Println("text/javascript")
		break
	default:
		fmt.Println("找不到后缀")
	}

}

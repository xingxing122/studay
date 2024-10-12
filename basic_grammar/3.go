package main

import "fmt"

func main() {

	var n1 int = 19
	var n2 float32 = 4.78
	var n3 bool = false
	var n4 byte = 'a'

	var s1 string = fmt.Sprint("%d", n1)
	fmt.Printf("s1对应的类型是： %T ,s1= %v ", s1, s1)

	var s2 string = fmt.Sprint("%d", n2)
	fmt.Printf("s2对应的类型是： %T ,s2= %q ", s2, s2)

	var s3 string = fmt.Sprint("%d", n3)
	fmt.Printf("s1对应的类型是： %T ,s3= %v ", s3, s3)

	var s4 string = fmt.Sprint("%d", n4)
	fmt.Printf("s4对应的类型是： %T ,s4= %v ", s4, s4)

}

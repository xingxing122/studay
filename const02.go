package main

import "fmt"

func main() {

	const (
		Mon = iota
		Thesd
		Web
		Thur = iota
		Fir
		Sat
		Sun
	)
	fmt.Println(Mon, Thesd, Web, Thur, Fir, Sat, Sun)

}

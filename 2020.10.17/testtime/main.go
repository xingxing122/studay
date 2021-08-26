package main

import (
	"fmt"
	"time"
)

func main() {
	now  := time.Now()
	fmt.Println(now)
	fmt.Println(now.Unix())
	format := "2006/01/02 15:04:05"
	fmt.Println(now.Format(format))

	 t1, err := time.Parse("2006-01-01 15:04","2008/01-02 16:01")

	fmt.Println(err,t1)

	// time.Unix()
}

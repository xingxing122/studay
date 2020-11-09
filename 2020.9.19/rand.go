package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	caishuzi := rand.Intn(100)
	
	fmt.Println(caishuzi)
}

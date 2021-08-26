package main

import (
	"crypto/md5"
	"fmt"
	"crypto/sha512"
)

func main() {
   md5 :=  fmt.Sprintf("%X",md5.Sum([]byte("i am kk")))
   fmt.Println(md5)

   sha512 :=  fmt.Sprintf("%X",sha512.Sum512([]byte(txt)))
   fmt.Println(sha512)
}

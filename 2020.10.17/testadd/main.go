package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    args := os.Args[1:]

    usage := func(){
        fmt.Println("usage: Add.exe   1  2 [3 4 .....]")
    }
    if len(args) < 2 {
        usage()
    }

    total := 0
    hasError := false
    for _, value := range args {
        if intValue, err := strconv.Atoi(value); err != nil {
            usage()
            hasError = true
            break
        } else {
            total += intValue
        }
    }
    if !hasError {
        fmt.Println("total:",total )
    }
}

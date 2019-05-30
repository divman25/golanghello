package main

import (
    "fmt"
    "test"
    "hello/testinside"
)

func main() {
    fmt.Println("Hello world")

    test.PrintMe()
    testinside.PrintMe()
    testinside.PrintMe2()
}
package main

import (
    "fmt"
    "test"
    "hello/testinside"
    "arrays"
)

func main() {
    fmt.Println("Hello world")

    test.PrintMe()
    testinside.PrintMe()
    testinside.PrintMe2()

    arrays.Execute()
}
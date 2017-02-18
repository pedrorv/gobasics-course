package main

import "fmt"

var x int = 42 // x is not capitalized so it's not exported
// var X int = 42 would be exported

func main() {
  fmt.Println(x)
  foo()
  y := 84 // block scoped variable
  fmt.Println(y)
}

func foo() {
  fmt.Println(x)
  // fmt.Println(y) // y is not accessible
}
package main

import "fmt"

func main() {
  // executes this piece of code at the end of the function
  // that surrounds it. good for closing files
  defer world()
  hello()
}

func hello() {
  fmt.Print("hello ")
}

func world() {
  fmt.Print("world \n")
}
